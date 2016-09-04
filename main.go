package main

import (
	"flag"
	"fmt"
	_ "github.com/thoj/go-ircevent"
)

func main() {
	var addr = flag.String("addr", "chat.freenode.net:6667", "the irc server address")
	var nick = flag.String("nick", "[logbot]", "bot's nickname")
	var ident = flag.String("ident", "logbot", "bot's username")
	var realname = flag.String("realname", "beep boop imma bot", "bot's realname")
	var channel = flag.String("channel", "#channel", "channel to join")
	var filename = flag.String("filename", "/var/www/logs/{{.Year}}/{{.Year}}-{{.Month}}-{{.Day}}.txt", "where the logs should be stored")
	flag.Parse()
	fmt.Printf("%s, %s, %s, %s, %s, %s\n", *addr, *nick, *ident, *realname, *channel, *filename)
}
