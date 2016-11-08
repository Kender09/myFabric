package main

import(
  "time"
  "sync"
)

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
  _, ok := w.benchData.data[action]
  if !ok {
    w.benchData.data[action] = new(profData)
    w.benchData.data[action].count = 0
    w.benchData.data[action].sum = 0.0
    w.benchData.data[action].max = 0.0
  }
  measureTime(w.benchData.data[action], func() {
    postJSON(w.ip, createChainReq(action, msg))
  })
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
