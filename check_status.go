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
)

var (
	service_data service.Services_slice
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	chatID, _   := strconv.ParseInt(os.Getenv("chatID"), 10, 64)
	yourToken   := os.Getenv("yourToken")
	Env_Time, _ := strconv.ParseInt(os.Getenv("Timer_Minutes"), 10, 64)

	fmt.Printf("chatID: %d, yourToken: %s, How often to run: %d (Minutes)\n", chatID, yourToken, Env_Time)
	tChannel := time.NewTimer(time.Duration(Env_Time) * time.Minute)
	for {
		service_data = service.Parser_services()
		failed_data := ""
		for ii := range service_data.Services {
			service.Check_service_status(ii, 500, &service_data, &failed_data)
		}
		if len(failed_data) > 0 {
			failed_data = "↻ Check Status ...... 🔴FAILED! \n - \n" + failed_data + "-"
			bot_service.Telegram_bot_run(chatID, yourToken, failed_data)
		} else {
			failed_data = "↻ Check Status ...... 🟢PASS! \n" + failed_data
			bot_service.Telegram_bot_run(chatID, yourToken, failed_data)
		}
		tChannel.Reset(time.Duration(Env_Time) * time.Minute)
		<-tChannel.C
	 }
}
