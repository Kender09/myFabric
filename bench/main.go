package bench

import (
  "sync"
)

type Worker struct{
  campany Campany
  id int
  ip string
  chName string
  benchData benchData
  mu sync.Mutex
}

type Campany struct{
  name string
  partner string
}

// まだ考え中
type benchData struct{
  data map[string]*profData
  err_cnt int
  err_strs []string
}

type profData struct{
  count int
  sum float64
  max float64
}

type VPs struct{
  ips []string
}

var vps VPs

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
  measureTime(w.benchData.data[action], func() {
    postJSON(w.ip, createChainReq(action, msg))
  })
}


