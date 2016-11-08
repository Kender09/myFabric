package main

import (
  "sync"
  "flag"
  "fmt"
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

func main() {
  var a_ip = flag.String("a", "localhost", "ipアドレス")
  //var b_ip = flag.String("b", "0.0.0.0", "ipアドレス")
  //var c_ip = flag.String("c", "0.0.0.0", "ipアドレス")
  flag.Parse()

  // init
  var a_w Worker
  a_w.campany.name = "a"
  a_w.campany.partner = "b"
  a_w.id = 1
  a_w.ip = *a_ip
  a_w.benchData.data = map[string]*profData{}

  a_w.chainReqController("invoke")
  fmt.Printf("%+v", a_w.benchData.data["invoke"])
}

