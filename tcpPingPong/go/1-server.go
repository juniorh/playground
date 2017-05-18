package main 

import (
  "fmt"
  "net"
  "os"
  "runtime"
  "strconv"
  //"time"
)

func main() {
  var PORT string
  if len(os.Args) > 1 {
    PORT = os.Args[1]
  } else {
    PORT = "9000"
  }

  l, err := net.Listen("tcp", ":"+PORT)
  if err != nil {
    fmt.Println("Error listening:", err.Error())
    os.Exit(1)
  }
  // Close the listener when the application closes.
  defer l.Close()
  for {
    conn,err := l.Accept()
    if err != nil {
      fmt.Println("Error accepting: ", err.Error())
      os.Exit(1)
    }
    fmt.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())
    go handleRequest(conn)
  }
}

func handleRequest(conn net.Conn) {
  var msg string
  var reqLen int
  var err error
  buff := make([]byte, 1024)
  for {
    reqLen,err = conn.Read(buff)
    if err != nil {
      fmt.Println("Error reading:", err.Error(), " - Closed connection")
      return
    }
    msg = "recv "+strconv.Itoa(reqLen)+"bytes,\t"
    msg += "msg="+string(buff)
    fmt.Print("num routines: ", runtime.NumGoroutine(),"\t")
    fmt.Println(msg)
    conn.Write([]byte(msg));
    //time.Sleep(1000 * time.Millisecond)
  }
}
