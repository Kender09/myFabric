package main

import (
  "sync"
)

type worker struct{
  campany string
  id int
  data benchdata
  mu sync.Mutex
}

// まだ考え中
type benchdata struct{
  r_cnt int
  w_cnt int
  err_cnt int
  err_strs []string
}

type VPs struct{
  ips []string
}

var vps VPs
