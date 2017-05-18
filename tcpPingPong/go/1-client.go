package main

import (
  "fmt"
  "net"
  "time"
  "os"
  "strconv"
)

func main() {
  i := int(0)
  var msg string

  conn,err := net.Dial("tcp", ":9000")

  if err != nil {
    fmt.Println("Error connection:", err.Error())
    os.Exit(1)
  }
  // Close the listener when the application closes.
  defer conn.Close()

  for {
    i += 1
    msg = "MESSAGE-"+strconv.Itoa(i)
    conn.Write([]byte(msg))
    time.Sleep(1000 * time.Millisecond)
  }
}
