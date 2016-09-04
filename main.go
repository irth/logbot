package main

import (
	"flag"
	"github.com/thoj/go-ircevent"
)

func main() {
	var addr = flag.String("addr", "chat.freenode.net:6667", "the irc server address")
	var nick = flag.String("nick", "[logbot]", "bot's nickname")
	var ident = flag.String("ident", "logbot", "bot's username")
	var channel = flag.String("channel", "#channel", "channel to join")
	//var filename = flag.String("filename", "/var/www/logs/{{.Year}}/{{.Year}}-{{.Month}}-{{.Day}}.txt", "where the logs should be stored")
	flag.Parse()

	// Connect to IRC
	conn := irc.IRC(*nick, *ident)

	conn.Connect(*addr)
	conn.Join(*channel)

	conn.Wait()
}
