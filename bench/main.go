package main

import (
  "sync"
  "flag"
  "fmt"
  "time"
)

func main() {
  var a_ip = flag.String("a", "localhost", "ipアドレス")
  var b_ip = flag.String("b", "localhost", "ipアドレス")
  var c_ip = flag.String("c", "localhost", "ipアドレス")
  var work_num = flag.Int("w", 1, "work数")
  var interval_time = flag.String("time", "1s", "ベンチ時間")
  flag.Parse()

  var wg *sync.WaitGroup
  wg = &sync.WaitGroup{}

  // init
  var a_w Worker
  a_w.init("a", "b", *a_ip)
  var b_w Worker
  b_w.init("b", "c", *b_ip)
  var c_w Worker
  c_w.init("c", "a", *c_ip)


  duration, err := time.ParseDuration(*interval_time)
  if err != nil {
    fmt.Println(err)
    panic(err)
  }
  endtime := time.Now().Add(duration)

  for i := 0; i < *work_num; i++ {
    a_w.work(endtime, wg)
    b_w.work(endtime, wg)
    c_w.work(endtime, wg)
  }

  wg.Wait()

  a_w.resultPrintf()
  b_w.resultPrintf()
  c_w.resultPrintf()
}

