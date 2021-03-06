package main

import (
    "fmt"
    "net"
    "strings"
)

//处理用户请求
func HandleConn(conn net.Conn) {
    //函数调用完毕，自动关闭conn
    defer conn.Close()

    //获取客户端的网络地址信息
    fmt.Println("start parse connection ip address.")
    addr := conn.RemoteAddr().String()

    fmt.Println(addr, " conncet sucessful")

    buf := make([]byte, 2048)

    for {
        //读取用户数据
        fmt.Println("start reading data from client...")
        n, err := conn.Read(buf)
        if err != nil {
            fmt.Println("err = ", err)
            return
        }
        fmt.Println("read data successcul. data:")
        fmt.Printf("[%s]: %s\n", addr, string(buf[:n]))
        fmt.Println("len = ", len(string(buf[:n])))

        //if "exit" == string(buf[:n-1]) { //nc测试
        if "exit" == string(buf[:n-2]) { //自己写的客户端测试, 发送时，多了2个字符, "\r\n"
            fmt.Println(addr, " exit")
            return
        }

        //把数据转换为大写，再给用户发送
        conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
    }
}