package main

import (
  "sync"
  "flag"
  "fmt"
  "time"
  //"regexp"
  //"strconv"
)

func main() {
  var a_ip = flag.String("a", "localhost", "ipアドレス")
  var b_ip = flag.String("b", "localhost", "ipアドレス")
  var c_ip = flag.String("c", "localhost", "ipアドレス")
  var d_ip = flag.String("d", "localhost", "ipアドレス")
  var work_num = flag.Int("w", 1, "work数")
  var interval_time = flag.String("time", "1s", "ベンチ時間")
  flag.Parse()

  var wg *sync.WaitGroup
  wg = &sync.WaitGroup{}

  // init
  var a_w Worker
  a_w.init("a", "b", *a_ip, "jim")
  var b_w Worker
  b_w.init("b", "c", *b_ip, "lukas")
  var c_w Worker
  c_w.init("c", "d", *c_ip, "diego")
  var d_w Worker
  d_w.init("d", "a", *d_ip, "binhn")


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
    d_w.work(endtime, wg)
  }

  wg.Wait()

  a_w.resultPrintf()
  b_w.resultPrintf()
  c_w.resultPrintf()
  d_w.resultPrintf()

//  r := regexp.MustCompile(`[\d]+`)
//
//  result := r.FindAllStringSubmatch(a_w.res, -1)
//  a_fin, _ := strconv.Atoi(result[2][0])
//  result = r.FindAllStringSubmatch(b_w.res, -1)
//  b_fin, _ := strconv.Atoi(result[2][0])
//  result = r.FindAllStringSubmatch(c_w.res, -1)
//  c_fin, _ := strconv.Atoi(result[2][0])
//  result = r.FindAllStringSubmatch(d_w.res, -1)
//  d_fin, _ := strconv.Atoi(result[2][0])
//  a_err := (10000 + (int(d_w.benchData["invoke"].count) - int(a_w.benchData["invoke"].count))) - a_fin
//  b_err := (10000 + (int(a_w.benchData["invoke"].count) - int(b_w.benchData["invoke"].count))) - b_fin
//  c_err := (10000 + (int(b_w.benchData["invoke"].count) - int(c_w.benchData["invoke"].count))) - c_fin
//  d_err := (10000 + (int(c_w.benchData["invoke"].count) - int(d_w.benchData["invoke"].count))) - d_fin
//  fmt.Println("a:", a_fin, " b:", b_fin, " c:", c_fin, " d:", d_fin)
//  fmt.Println("a_err:", a_err, " b_err:", b_err, " c_err:", c_err, " d_err:", d_err)
}

