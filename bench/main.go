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
  benchData map[string]*profData
  res string
  res_err error
  mu sync.Mutex
}

type Campany struct{
  name string
  partner string
}

type profData struct{
  count int
  err_cnt int
  sum float64
  max float64
  histgram map[string]float64
}

type VPs struct{
  ips []string
}

var vps VPs

func main() {
  var a_ip = flag.String("a", "localhost", "ipアドレス")
  //var b_ip = flag.String("b", "0.0.0.0", "ipアドレス")
  //var c_ip = flag.String("c", "0.0.0.0", "ipアドレス")
  var interval_time = flag.String("time", "1s", "ベンチ時間")
  flag.Parse()

  var wg *sync.WaitGroup
  wg = &sync.WaitGroup{}

  // init
  var a_w Worker
  a_w.init("a", "b", *a_ip)

  duration, err := time.ParseDuration(*interval_time)
  if err != nil {
    fmt.Println(err)
    panic(err)
  }
  endtime := time.Now().Add(duration)
  a_w.work(endtime, wg)
  wg.Wait()

  fmt.Printf("%+v", a_w.benchData["invoke"])
}

