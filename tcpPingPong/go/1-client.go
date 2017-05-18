package main

import (
  "fmt"
  "net"
  "time"
  "strconv"
  "flag"
)


func main() {
  hostPtr := flag.String("host", "localhost", "destination host")
  portPtr := flag.String("port", "9000", "destination port")
  procPtr := flag.Int("proc", 1, "number connection created")
  flag.Parse()
  
  for j:=0; j<= *procPtr; j++ {
    go spawnConn(*hostPtr,*portPtr)
  }
  time.Sleep(10000 * time.Millisecond)
}

func spawnConn(host,port string) {
  var msg string

  conn,err := net.Dial("tcp", host+":"+port)

  if err != nil {
    fmt.Println("Error connection:", err.Error())
    //os.Exit(1)
  }
  // Close the listener when the application closes.
  defer conn.Close()
  
  i := int(0)
  for {
    i += 1
    msg = "MESSAGE-"+strconv.Itoa(i)
    conn.Write([]byte(msg))
    time.Sleep(1000 * time.Millisecond)
  }
}
