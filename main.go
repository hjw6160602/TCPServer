package main

import (
    "fmt"
    "net"
)

func main() {
    //addListen()
    internetgateway()
    //Example_WANPPPConnection1_GetExternalIPAddress()
    //Example_WANIPConnection_GetExternalIPAddress()
    //Example_ReuseDiscoveredDevice()
    //Example_WANCommonInterfaceConfig1_GetBytesTransferred()
    //Example_WANCommonInterfaceConfig2_GetBytesTransferred()

}



func addListen()  {
    //监听
    port := "60001"
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
        } else {
            fmt.Println("find a new request, accept successful.")
        }

        //处理用户请求, 新建一个协程
        go HandleConn(conn)
    }
}
