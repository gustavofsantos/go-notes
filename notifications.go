package main

import (
	"fmt"
	"log"
	"os/exec"
)

type Notification struct {
	text string
}

var HANDLERS = []func(n *Notification){notifyGnome, notifyTmux}

func HandleNotification(notification *Notification) {
  for _, handler := range HANDLERS {
    handler(notification)
  }
}

func notifyGnome(notification *Notification) {
	err := exec.Command("notify-send", "-u", "critical", "Atenção: Reunião", fmt.Sprintf("%s", notification.text)).Run()
	if err != nil {
		panic(err)
	}
}

func notifyTmux(notification *Notification) {
	err := exec.Command("tmux", "display-message", "-d", "5000", fmt.Sprintf("%s", notification.text)).Run()
	if err != nil {
		log.Println("Error displaying using tmux. Tmux is running?")
	}
}

