package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	bot_r, bot_d *tgbotapi.BotAPI
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
					replyMsg.Text = "/service  <-- to show all service's items"
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
