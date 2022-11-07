package bot

import (
	"log"
	service "github.com/MinFengLin/check_service_status/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	bot_r, bot_d *tgbotapi.BotAPI
	service_data service.Services_slice
)

func sendMsg(msg string, chatID int64) {
	NewMsg := tgbotapi.NewMessage(chatID, msg)
	// NewMsg.ParseMode = tgbotapi.ModeHTML   //傳送html格式的訊息
	_, err := bot_d.Send(NewMsg)
	if err == nil {
		log.Printf("Send telegram message success")
	} else {
		log.Printf("Send telegram message error")
	}
}

func replyMsg(msg string, chatID int64) {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	updates, _ := bot_r.GetUpdatesChan(updateConfig)
	for update_i := range updates {
		update := update_i
		go func() {
			cmd_text := update.Message.Text
			chatID := update.Message.Chat.ID
			replyMsg := tgbotapi.NewMessage(chatID, cmd_text)
			replyMsg.ReplyToMessageID = update.Message.MessageID
			if update.Message.IsCommand() {
				switch update.Message.Command() {
				case "service":
					replyMsg.Text = msg
				case "help":
					replyMsg.Text = "/service        <-- to show all service's items\n"
					replyMsg.Text = replyMsg.Text + "/service_debug  <-- execute immediately check service"
				case "service_debug":
					service_data = service.Parser_services()
					failed_data := ""
					for ii := range service_data.Services {
						service.Check_service_status(ii, 500, &service_data, &failed_data)
					}
					if len(failed_data) > 0 {
						failed_data = "↻ Check Status ...... 🔴FAILED! \n - \n" + failed_data + "-"
					} else {
						failed_data = "↻ Check Status ...... 🟢PASS! \n" + failed_data
					}
					replyMsg.Text = failed_data
				default:
					replyMsg.Text = ""
				}
			} else {
				replyMsg.Text = ""
			}
			if len(replyMsg.Text) > 0 {
				_, _ = bot_r.Send(replyMsg)
			}
		}()
	}
}

func Tgbot_cmd(chatid *int64, token *string) {
	service_data = service.Parser_services()
	services_info := "-\n"

	for ii := range service_data.Services {
		services_info = services_info + service_data.Services[ii].Ip+":"+service_data.Services[ii].Port + " - (" +service_data.Services[ii].Service + ")" + "\n"
	}
	services_info = services_info + "-\n"
	Telegram_reply_run(*chatid, *token, services_info)
}

func Telegram_reply_run(chatID int64, yourToken string, msg string) {
	var err error
	bot_r, err = tgbotapi.NewBotAPI(yourToken)
	if err != nil {
		log.Fatal(err)
	}

	bot_r.Debug = false

	replyMsg(msg, chatID)
}

func Telegram_bot_run(chatID int64, yourToken string, msg string) {
	var err error
	bot_d, err = tgbotapi.NewBotAPI(yourToken)
	if err != nil {
		log.Fatal(err)
	}

	bot_d.Debug = false

	sendMsg(msg, chatID)
}
