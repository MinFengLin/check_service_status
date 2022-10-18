package main

import (
	bot_service "bot"
	service "service"
	"fmt"
	"strconv"

	/*
	 * env
	 */
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

var (
	service_data service.Services_slice
)

func check(chatid *int64, token *string) {
	service_data = service.Parser_services()
	failed_data := ""
	for ii := range service_data.Services {
		service.Check_service_status(ii, 500, &service_data, &failed_data)
	}
	if len(failed_data) > 0 {
		failed_data = "↻ Check Status ...... 🔴FAILED! \n - \n" + failed_data + "-"
		bot_service.Telegram_bot_run(*chatid, *token, failed_data)
	} else {
		failed_data = "↻ Check Status ...... 🟢PASS! \n" + failed_data
		bot_service.Telegram_bot_run(*chatid, *token, failed_data)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	chatID, _   := strconv.ParseInt(os.Getenv("chatID"), 10, 64)
	yourToken   := os.Getenv("yourToken")
	Env_Time, _ := strconv.ParseInt(os.Getenv("Timer_Minutes"), 10, 64)
	crontab_time := os.Getenv("Crontime")

	if len(crontab_time) > 0 {
		fmt.Printf("chatID: %d, yourToken: %s, Use crontab: %s \n", chatID, yourToken, crontab_time)
		c := cron.New()
		c.AddFunc(crontab_time, func() {
			check(&chatID, &yourToken)
		})
		check(&chatID, &yourToken)
		c.Start()
		select{}
	} else {
		fmt.Printf("chatID: %d, yourToken: %s, How often to run: %d (Minutes)\n", chatID, yourToken, Env_Time)
		tChannel := time.NewTimer(time.Duration(Env_Time) * time.Minute)
		for {
			check(&chatID, &yourToken)
			tChannel.Reset(time.Duration(Env_Time) * time.Minute)
			<-tChannel.C
		 }
	}
}
