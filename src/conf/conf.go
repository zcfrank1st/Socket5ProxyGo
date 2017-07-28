package conf

import "flag"

var (
    User string
    Pass string
    Port string
)


func init () {
    flag.StringVar(&Port, "port", "9779", "socks5 proxy port")
    flag.StringVar(&User, "user", "nil", "auth user")
    flag.StringVar(&Pass, "pass", "nilnil", "auth pass")
    flag.Parse()
}