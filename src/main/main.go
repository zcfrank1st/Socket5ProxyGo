package main

import (
    "net"
    "../proxy"
    "flag"
    "github.com/logrusorgru/aurora"
    "log"
)

var port string

func init() {
    flag.StringVar(&port, "port", "9779", "socks5 proxy port")
    flag.Parse()
}

func main()  {
    socket, err := net.Listen("tcp", ":" + port)
    if err != nil {
        return
    }
    log.Printf("socks5 proxy server running on port [:%s], listening ...\n", aurora.Green(port))

    for {
        client, err := socket.Accept()

        if err != nil {
            return
        }


        var handler proxy.ProxyHandler = new(proxy.Socks5ProxyHandler)

        go handler.Handle(client)

        log.Println(aurora.Blue(client), " request handling...")
    }

}
