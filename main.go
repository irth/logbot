package main

import (
	"flag"
	"fmt"
	"github.com/thoj/go-ircevent"
	"time"
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

	conn.AddCallback("PRIVMSG", func(e *irc.Event) {
		var timestamp = time.Now().Format("15:04:05")
		var msg = fmt.Sprintf("%s <%s> %s", timestamp, e.Nick, e.Message())
		fmt.Println(msg)
	})

	conn.Wait()
}
