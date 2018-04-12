package main

import (
  "fmt"
  "os"
  "strconv"
  "time"
)

const NANOSECOND int64  = 1
const MICROSECOND int64 = 1000 * NANOSECOND
const MILLISECOND int64 = 1000 * MICROSECOND
const SECOND int64      = 1000 * MILLISECOND
const MINUTE int64      = 60 * SECOND
const HOUR int64        = 60 * MINUTE
const DAY int64         = 24 * HOUR

const COLOR_OFF string     = "\x1B[0m"
const COLOR_RED string     = "\x1B[31mm"
const COLOR_GREEN string   = "\x1B[32m"
const COLOR_YELLOW string  = "\x1B[33m"
const COLOR_BLUE string    = "\x1B[34m"
const COLOR_MAGENTA string = "\x1B[35m"
const COLOR_CYAN string    = "\x1B[36m"
const COLOR_WHITE string   = "\x1B[37m"

func usage() {
  fmt.Println("Usage: ftime format <nano_duration>")
  fmt.Println("       ftime ns|us|ms|s|m|h|d")
  fmt.Println("       ftime ns-iso|us-iso|ms-iso|s-iso <timestamp>")
  fmt.Println("       ftime iso|iso-bash")
  fmt.Println("       ftime bash")
}

func now() time.Time {
  return time.Now()
}

func nowNano() int64 {
  return now().UnixNano()
}

func nowIso() string {
  return toIso(now())
}

func toIso(t time.Time) string {
  utc := t.UTC()
  return fmt.Sprintf("%04d-%02d-%2dT%02d:%02d:%02d.%03dZ",
    utc.Year(),
    utc.Month(),
    utc.Day(),
    utc.Hour(),
    utc.Minute(),
    utc.Second(),
    utc.UnixNano() % SECOND / MILLISECOND)
}

func bashIso(t time.Time) string {
  utc := t.UTC()
  return fmt.Sprintf("%s%04d%s-%s%02d%s-%s%2d%sT%s%02d%s:%s%02d%s:%s%02d%s.%s%03d%sZ",
    COLOR_GREEN, utc.Year(), COLOR_OFF,
    COLOR_GREEN, utc.Month(), COLOR_OFF,
    COLOR_GREEN, utc.Day(), COLOR_OFF,
    COLOR_YELLOW, utc.Hour(), COLOR_OFF,
    COLOR_YELLOW, utc.Minute(), COLOR_OFF,
    COLOR_YELLOW, utc.Second(), COLOR_OFF,
    COLOR_YELLOW, utc.UnixNano() % SECOND / MILLISECOND, COLOR_OFF)
}

func bashTime(t time.Time) string {
  return fmt.Sprintf("%s%04d%s-%s%02d%s-%s%2d%s %s%02d%s:%s%02d%s:%s%02d%s",
    COLOR_GREEN, t.Year(), COLOR_OFF,
    COLOR_GREEN, t.Month(), COLOR_OFF,
    COLOR_GREEN, t.Day(), COLOR_OFF,
    COLOR_YELLOW, t.Hour(), COLOR_OFF,
    COLOR_YELLOW, t.Minute(), COLOR_OFF,
    COLOR_YELLOW, t.Second(), COLOR_OFF)
}

func nanoTime(ns int64) time.Time {
  return time.Unix(ns / SECOND, ns % SECOND)
}

func parseInt(s string) int64 {
  v, _ := strconv.ParseInt(s, 10, 64)
  return v
}

func duration(nanos int64) string {
  if (nanos >= HOUR) {
    high := nanos / HOUR
    low := nanos % HOUR / MINUTE
    return fmt.Sprintf("%d h, %d m", high, low)
  } else if (nanos >= MINUTE) {
    high := nanos / MINUTE
    low := nanos % MINUTE / SECOND
    return fmt.Sprintf("%d m, %d s", high, low)
  } else if (nanos >= SECOND) {
    high := nanos / SECOND
    low := nanos % SECOND / MILLISECOND
    return fmt.Sprintf("%d.%03d s", high, low)
  } else if (nanos >= MILLISECOND) {
    high := nanos / MILLISECOND
    low := nanos % MILLISECOND / MICROSECOND
    return fmt.Sprintf("%d.%03d ms", high, low)
  } else {
    high := nanos / MICROSECOND
    low := nanos % MICROSECOND
    return fmt.Sprintf("%d.%03d us", high, low)
  }
}

func main() {
  args := os.Args
  argc := len(args)

  if (argc == 2) {
    switch args[1] {
      case "ns": fmt.Println(nowNano())
      case "us": fmt.Println(nowNano() / MICROSECOND)
      case "ms": fmt.Println(nowNano() / MILLISECOND)
      case "s": fmt.Println(nowNano() / SECOND)
      case "m": fmt.Println(nowNano() / MINUTE)
      case "h": fmt.Println(nowNano() / HOUR)
      case "d": fmt.Println(nowNano() / DAY)
      case "iso": fmt.Println(nowIso())
      case "iso-bash": fmt.Println(bashIso(now()))
      case "bash": fmt.Println(bashTime(now()))
      default: usage()
    }
  } else if (argc == 3) {
    switch args[1] {
      case "format": fmt.Println(duration(parseInt(args[2])))
      case "ns-iso": fmt.Println(toIso(nanoTime(parseInt(args[2]))))
      case "us-iso": fmt.Println(toIso(nanoTime(parseInt(args[2]) * MICROSECOND)))
      case "ms-iso": fmt.Println(toIso(nanoTime(parseInt(args[2]) * MILLISECOND)))
      case "s-iso": fmt.Println(toIso(nanoTime(parseInt(args[2]) * SECOND)))
      default: usage()
    }
  } else {
    usage()
  }
}

