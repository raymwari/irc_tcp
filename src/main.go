package main

import (
	"bufio"
	"fmt"
	"net"
	"regexp"
	"strings"
)

var guest string

const (
	nick  = "nixngc"
	chnl  = "fallen8192"
	email = "fallen8192@gmail.com"
)

func r_callback(ev_str string, con net.Conn) {
	switch ev_str {
	case "/ping":
		fmt.Fprintf(con, "PONG\r\n")
	case "/join":
		if guest != nick {
			cmd := fmt.Sprintf("PRIVMSG #%s :Hello %s, thanks for joining. Kindly check your inbox for more information!\r\n",
				chnl,
				guest)
			fmt.Fprintf(con, "%s", cmd)
			cmd = fmt.Sprintf("PRIVMSG %s :Hello %s, issue command: \x02whois %s\x02 then leave a message if active, or leave a message at \x02%s!\x02 Thanks again...\r\n",
				guest,
				guest,
				chnl,
				email)
			fmt.Fprintf(con, "%s", cmd)
		}
	case "/message":
		/*handle for on message received*/
	case "/private_message":
		/*handle for on private message received*/
	default:
		panic("encountered an undefined event")
	}
}

/*response handler:*/
func r_handler(fn func(string, net.Conn), resp string, con net.Conn) {
	fmt.Println(resp)
	if strings.HasPrefix(resp, "PING") {
		fn(ev_stat[ev_ping], con)
	} else if strings.Contains(resp, "JOIN") {
		r := regexp.MustCompile("^:([^!]+)!")
		match := r.FindString(resp)
		match = strings.Trim(match, ":")
		match = strings.Trim(match, "!")
		guest = match
		fn(ev_stat[ev_join], con)
	}
}

func connect(opts opts, fn func(con net.Conn)) {
	con, err := net.Dial(opts.net, opts.add)
	if err != nil {
		panic(err.Error())
	}
	fn(con)
}

func main() {
	options := opts{}
	options.net = "tcp"
	options.add = "irc.libera.chat:6667"

	connect(options, func(con net.Conn) {
		defer con.Close()
		/*registration:*/
		fmt.Fprintf(con, "USER %s 0 * :IRC BOT\r\n", nick)
		fmt.Fprintf(con, "NICK %s\r\n", nick)
		/*join channel:*/
		fmt.Fprintf(con, "JOIN #%s\r\n", chnl)
		scnr := bufio.NewScanner(con)
		for scnr.Scan() {
			r_handler(r_callback, scnr.Text(), con)
		}
	})
}
