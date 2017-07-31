package conf

import "flag"

var (
    Auth bool
    User string
    Pass string
    Port string
)


func init () {
    flag.BoolVar(&Auth, "auth", false, "if use auth")
    flag.StringVar(&Port, "port", "9779", "socks5 proxy port")
    flag.StringVar(&User, "user", "hello", "auth user")
    flag.StringVar(&Pass, "pass", "world", "auth pass")
    flag.Parse()
}