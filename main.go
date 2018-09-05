package main

import (
    "net"
    "fmt"
)

func main() {
    //监听
    port := "60000"
    ip := "192.168.3.113"//
    //port := "37371"
    //ip := "124.78.94.72"
    address := ip + ":" + port
    fmt.Println("tcp is listening: ", address)

    listener, err := net.Listen("tcp", address)

    if err != nil {
        fmt.Println("err = ", err)
        return
    }

    defer listener.Close()

    //接收多个用户
    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("err = ", err)
            return
        }

        //处理用户请求, 新建一个协程
        go HandleConn(conn)
    }

}
