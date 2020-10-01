package main

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
	"strconv"
	"time"
)

var lastStatus string
var chatID, _ = strconv.ParseInt(os.Getenv("CHAT_ID"), 10, 64)

func main() {
	tg, err := tb.NewBot(tb.Settings{
		Token: os.Getenv("TG_TOKEN"),
	})
	if err != nil {
		log.Fatalln("Telegram", err)
	}

	for {
		status, err := getStatus(os.Getenv("JOB_NUMBER"))
		if err == nil {
			if status != lastStatus {
				lastStatus = status
				tg.Send(&tb.Chat{ID: chatID}, "*Новый статус B2X*\n\n" + status, tb.ModeMarkdown)
			}
		}
		time.Sleep(time.Minute * 5)
	}
}

func getStatus(jobNumber string) (string, error) {
	sessionID, err := globAuthenticate(os.Getenv("USER_ID"), os.Getenv("PASSWORD"))
	if err != nil {
		return "", err
	}
	return repairSummaryLookup(os.Getenv("AUTH_TOKEN"), sessionID, jobNumber)
}
