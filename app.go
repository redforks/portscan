package main

import (
  "fmt"
  "os"
  "net"
  "time"
  "sync"
)

var (
  host string // The host address to scan
)

func init() {
  if len(os.Args) != 2 {
    fmt.Fprintf(os.Stderr, "Usage: %s host\n", os.Args[0])
    os.Exit(1)
  }
  host = os.Args[1]
}

func main() {
  d := net.Dialer{Timeout: 10 * time.Second}
  p := make(chan bool, 500) // make 500 parallel connection
  wg := sync.WaitGroup{}

  c := func(port int) {
    conn, err := d.Dial(`tcp`, fmt.Sprintf(`%s:%d`, host, port))
    if err == nil {
      conn.Close()
      fmt.Printf("%d passed\n", port)
    }
    <-p
    wg.Done()
  }

  wg.Add(65536)
  for i:=0; i < 65536; i++ {
    p<-true
    go c(i)
  }

  wg.Wait()
}
