package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/thoj/go-ircevent"
	"os"
	"path"
	"text/template"
	"time"
)

type FilenameData struct {
	Year    string
	Month   string
	Day     string
	Channel string
}

func Write(filenameTemplate string, channel string, text string) {
	var t = time.Now()
	var d = FilenameData{
		Year:    t.Format("2006"),
		Month:   t.Format("01"),
		Day:     t.Format("02"),
		Channel: channel,
	}
	var tmpl, err = template.New("filename").Parse(filenameTemplate)
	if err != nil {
		panic(err)
	}

	var filenameBytes bytes.Buffer
	err = tmpl.Execute(&filenameBytes, d)
	if err != nil {
		panic(err)
	}
	var filename = filenameBytes.String()

	err = os.MkdirAll(path.Dir(filename), 0777)
	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	_, err = file.WriteString(text + "\n")
	if err != nil {
		panic(err)
	}
}

func main() {
	var addr = flag.String("addr", "chat.freenode.net:6667", "the irc server address")
	var nick = flag.String("nick", "[logbot]", "bot's nickname")
	var ident = flag.String("ident", "logbot", "bot's username")
	var channel = flag.String("channel", "#channel", "channel to join")
	var filename = flag.String("filename", "/var/www/logs/{{.Channel}}/{{.Year}}/{{.Year}}-{{.Month}}-{{.Day}}.txt", "where the logs should be stored")
	flag.Parse()

	// Connect to IRC
	conn := irc.IRC(*nick, *ident)

	conn.Connect(*addr)
	conn.Join(*channel)

	conn.AddCallback("PRIVMSG", func(e *irc.Event) {
		var timestamp = time.Now().Format("15:04:05")
		var msg = fmt.Sprintf("%s <%s> %s", timestamp, e.Nick, e.Message())
		Write(*filename, e.Arguments[0], msg)
	})

	conn.AddCallback("CTCP_ACTION", func(e *irc.Event) {
		var timestamp = time.Now().Format("15:04:05")
		var msg = fmt.Sprintf("%s * %s %s", timestamp, e.Nick, e.Message())
		Write(*filename, e.Arguments[0], msg)
	})

	conn.Wait()
}
