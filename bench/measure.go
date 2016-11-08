package bench

import(
  "time"
)

func measureTime(data *profData, fn func()) {
  start := time.Now()
  defer recordTime(start, data)
}

func recordTime(start time.Time, data *profData) {
  end := time.Now()
  measure_time := (end.Sub(start)).Seconds()
  data.count += 1
  data.sum += measure_time
  if data.max < measure_time { data.max = measure_time }
}
