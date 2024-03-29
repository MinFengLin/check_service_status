package main

import (
	"fmt"
	"strconv"

	/*
	 * env
	 */
	"log"
	"os"
	"time"

	bot "github.com/MinFengLin/check_service_status/bot"
	service "github.com/MinFengLin/check_service_status/service"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

var (
	service_data service.Services_slice
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	chatID, _ := strconv.ParseInt(os.Getenv("chatID"), 10, 64)
	yourToken := os.Getenv("yourToken")
	Env_Time, _ := strconv.ParseInt(os.Getenv("Timer_Minutes"), 10, 64)
	crontab_time := os.Getenv("Crontime")

	go bot.Telegram_reply_run(&chatID, &yourToken)
	if len(crontab_time) > 0 {
		fmt.Printf("chatID: %d, yourToken: %s, Use crontab: %s \n", chatID, yourToken, crontab_time)
		c := cron.New()
		_, _ = c.AddFunc(crontab_time, func() {
			failed_data := service.Check_service_status()
			bot.Telegram_bot_run(&chatID, &yourToken, failed_data)
		})
		failed_data := service.Check_service_status()
		bot.Telegram_bot_run(&chatID, &yourToken, failed_data)
		c.Start()
		select {}
	} else {
		fmt.Printf("chatID: %d, yourToken: %s, How often to run: %d (Minutes)\n", chatID, yourToken, Env_Time)
		tChannel := time.NewTimer(time.Duration(Env_Time) * time.Minute)
		for {
			failed_data := service.Check_service_status()
			bot.Telegram_bot_run(&chatID, &yourToken, failed_data)
			tChannel.Reset(time.Duration(Env_Time) * time.Minute)
			<-tChannel.C
		}
	}
}
