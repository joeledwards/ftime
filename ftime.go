package main

import (
  "fmt"
  "os"
  "strconv"
  "time"
)

const NS int64 = 1
const US int64 = 1000 * NS
const MS int64 = 1000 * US
const S int64 = 1000 * MS
const M int64 = 60 * S
const H int64 = 60 * M
const D int64 = 24 * H

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
  return now().UTC().Format(time.RFC3339Nano)
}

func toIso(t time.Time) string {
  return t.UTC().Format(time.RFC3339Nano)
}

func nanoTime(ns int64) time.Time {
  return time.Unix(ns / S, ns % S)
}

func parseInt(s string) int64 {
  v, _ := strconv.ParseInt(s, 10, 64)
  return v
}

func main() {
  args := os.Args
  argc := len(args)

  if (argc == 2) {
    switch args[1] {
      case "ns": fmt.Println(nowNano())
      case "us": fmt.Println(nowNano() / US)
      case "ms": fmt.Println(nowNano() / MS)
      case "s": fmt.Println(nowNano() / S)
      case "m": fmt.Println(nowNano() / M)
      case "h": fmt.Println(nowNano() / H)
      case "d": fmt.Println(nowNano() / D)
      case "iso": fmt.Println(nowIso())
      default: usage()
    }
  } else if (argc == 3) {
    switch args[1] {
      case "format": fmt.Println()
      case "ns-iso": fmt.Println(toIso(nanoTime(parseInt(args[2]))))
      case "us-iso": fmt.Println(toIso(nanoTime(parseInt(args[2]) * US)))
      case "ms-iso": fmt.Println(toIso(nanoTime(parseInt(args[2]) * MS)))
      case "s-iso": fmt.Println(toIso(nanoTime(parseInt(args[2]) * S)))
      default: usage()
    }
  } else {
    usage()
  }
}

