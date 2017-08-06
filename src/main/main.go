package main

import (
    "net"
    "../proxy"
    . "../conf"
    "github.com/logrusorgru/aurora"
    "log"
)

func main()  {

    socket, err := net.Listen("tcp", ":" + Port)
    if err != nil {
        return
    }
    log.Printf("socks5 proxy server running on port [:%s], listening ...\n", aurora.Green(Port))

    for {
        client, err := socket.Accept()

        if err != nil {
            return
        }


        var handler proxy.Handler = new(proxy.Socks5ProxyHandler)

        go handler.Handle(client)

        log.Println(aurora.Blue(client), " request handling...")
    }

}
