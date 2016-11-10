package main

import(
  "time"
  "fmt"
)

func measureTime(data *profData, fn func()) {
  start := time.Now()
  defer recordTime(start, data)
  fn()
}

func recordTime(start time.Time, data *profData) {
  end := time.Now()
  measure_time := (end.Sub(start)).Seconds()
  data.count += 1
  data.sum += measure_time
  data.histgram[end.String()] = fmt.Sprint(measure_time)
  if data.max < measure_time { data.max = measure_time }
}
