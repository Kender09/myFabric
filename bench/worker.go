package main

import(
  "time"
  "sync"
  "strings"
  "fmt"
  "encoding/csv"
  "os"
)

type Worker struct{
  campany Campany
  id int
  ip string
  benchData map[string]*profData
  req_cnt int
  res string
  res_err error
  m *sync.Mutex
}

type Campany struct{
  name string
  partner string
}

type profData struct{
  count float64
  err_cnt float64
  sum float64
  max float64
  histgram map[string]string
}

var workid int = 1

func (w *Worker) init(name string, partner string, ip string) {
  w.campany.name = name
  w.campany.partner = partner
  w.id = workid
  w.ip = ip
  w.req_cnt = 0
  w.benchData = map[string]*profData{}
  w.m = new(sync.Mutex)
  workid += 1
}

func writeCsv(filename string, outmap map[string]string) {
  file, err := os.OpenFile("logs/" + filename + ".csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
  if err != nil {
    fmt.Println(err)
  }
  defer file.Close()

  writer := csv.NewWriter(file)
  for k, v := range outmap {
    writer.Write([]string{k, v})
  }
  writer.Flush()
}

func (w *Worker) resultPrintf() {
  fmt.Println(w.campany.name, " : ", w.campany.partner)
  for k, v := range w.benchData {
    fmt.Println(k, " count:", v.count, "err_count:", v.err_cnt, " sum:", v.sum, " ave:", v.sum/v.count, " max:", v.max)
    writeCsv(w.campany.name + "_" + k, v.histgram)
  }
}

func (w *Worker) chainReqController(action string) {
  var msg CtorMsg
  switch action {
    case "invoke": {
      msg.Args = []string{"invoke", w.campany.name, w.campany.partner, "1"}
    }
    case "query": {
      msg.Args = []string{"query", w.campany.name}
    }
    case "deploy": {
      msg.Args = []string{"init", "a", "10000", "b", "10000", "c", "10000"}
    }
  }
  _, ok := w.benchData[action]
  if !ok {
    w.m.Lock()
    w.benchData[action] = new(profData)
    w.benchData[action].count = 0.0
    w.benchData[action].err_cnt = 0.0
    w.benchData[action].sum = 0.0
    w.benchData[action].max = 0.0
    w.benchData[action].histgram = map[string]string{}
    w.m.Unlock()
  }
  measureTime(w.benchData[action], w.m, func() {
    postJSON(w, createChainReq(action, msg, w.req_cnt))
  })
  status := strings.Contains(w.res, "OK")
  if w.res_err != nil || !status {
    w.benchData[action].err_cnt +=1
  }
}

func (w *Worker) work(endtime time.Time, wg *sync.WaitGroup) {
  wg.Add(1)
  go func(w *Worker, endtime time.Time) {
    var nowtime time.Time
    for true {
      w.chainReqController("invoke")
      nowtime = time.Now()
      if nowtime.After(endtime) { break }
      //time.Sleep(100 * time.Millisecond)
    }
    wg.Done()
  }(w, endtime)
}
