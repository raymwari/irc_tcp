package main

type event int

const (
	ev_ping   event = iota /*ping*/
	ev_join                /*join channel*/
	ev_onmsg               /*on message*/
	ev_onpmsg              /*on private message*/
)

var ev_stat = map[event]string{
	ev_ping:   "/ping",
	ev_join:   "/join",
	ev_onmsg:  "/message",
	ev_onpmsg: "/private_message",
}

type opts struct {
	net string /*network*/
	add string /*address*/
}
