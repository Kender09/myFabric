package main

import (
  "sync"
  "flag"
  "fmt"
  "time"
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

func main() {
  var a_ip = flag.String("a", "localhost", "ipアドレス")
  //var b_ip = flag.String("b", "0.0.0.0", "ipアドレス")
  //var c_ip = flag.String("c", "0.0.0.0", "ipアドレス")
  flag.Parse()

  var wg *sync.WaitGroup
  wg = &sync.WaitGroup{}

  // init
  var a_w Worker
  a_w.campany.name = "a"
  a_w.campany.partner = "b"
  a_w.id = 1
  a_w.ip = *a_ip
  a_w.benchData.data = map[string]*profData{}

  endtime := time.Now().Add(1*time.Second)
  a_w.work(endtime, wg)
  wg.Wait()

  fmt.Printf("%+v", a_w.benchData.data["invoke"])
}

