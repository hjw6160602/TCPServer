package main

import (
    "fmt"
    "net"
    "os"
)

func main() {
    port := "60001"
    ip := "192.168.3.113"//
    UDPListen(ip, port)
    //TCPListen(ip, port)
    //internetgateway()
    //Example_WANPPPConnection1_GetExternalIPAddress()
    //Example_WANIPConnection_GetExternalIPAddress()
    //Example_ReuseDiscoveredDevice()
    //Example_WANCommonInterfaceConfig1_GetBytesTransferred()
    //Example_WANCommonInterfaceConfig2_GetBytesTransferred()
}



func TCPListen(ip string, port string)  {
    //监听
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

func UDPListen(ip string, port string) {
    address := ip + ":" + port
    fmt.Println("UDP is ready for resolve UDP address...")
    addr, err := net.ResolveUDPAddr("udp", address)

    if err != nil {
        fmt.Println("Can't resolve address: ", err)
        os.Exit(1)
    }
    fmt.Println("UDP is listening: ", address)
    conn, err := net.ListenUDP("udp", addr)
    if err != nil {
        fmt.Println("Error listening:", err)
        os.Exit(1)
    }

    HandleUDPConn(conn)
}

func HandleUDPConn(conn *net.UDPConn) {
    buf := make([]byte, 1024)
    n, remoteAddr, err := conn.ReadFromUDP(buf)

    fmt.Println("Find a new request, accept successful.")
    fmt.Println("Client ip address is: ",remoteAddr.IP)
    fmt.Println("Client port is: ",remoteAddr.Port)
    data := string(buf[:n])
    fmt.Println("Receive data from server: ", data)

    if err != nil {
        fmt.Println("failed to read UDP msg because of ", err.Error())
        return
    }
    fmt.Println(remoteAddr)
    conn.Write([]byte("Hello"))
    conn.Close()
}