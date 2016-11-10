package main

import(
  "time"
  "sync"
  "strings"
)

var workid int = 1

func (w *Worker) init(name string, partner string, ip string) {
  w.campany.name = name
  w.campany.partner = partner
  w.id = workid
  w.ip = ip
  w.benchData = map[string]*profData{}
  workid += 1
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
    w.benchData[action] = new(profData)
    w.benchData[action].count = 0
    w.benchData[action].err_cnt = 0
    w.benchData[action].sum = 0.0
    w.benchData[action].max = 0.0
    w.benchData[action].histgram = map[string]float64{}
  }
  measureTime(w.benchData[action], func() {
    postJSON(w, createChainReq(action, msg))
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
    }
    wg.Done()
  }(w, endtime)
}
