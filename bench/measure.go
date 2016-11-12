package main

import(
  "time"
  "fmt"
  "sync"
)

func measureTime(data *profData, m *sync.Mutex, fn func()) {
  start := time.Now()
  defer recordTime(start, data, m)
  fn()
}

func recordTime(start time.Time, data *profData, m *sync.Mutex) {
  end := time.Now()
  measure_time := (end.Sub(start)).Seconds()
  data.count += 1
  data.sum += measure_time
  m.Lock()
  data.histgram[end.String()] = fmt.Sprint(measure_time)
  m.Unlock()
  if data.max < measure_time { data.max = measure_time }
}
