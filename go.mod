module main

replace (
	bot => ./bot
	service => ./service
)

go 1.19

require (
	bot v0.0.0-00010101000000-000000000000
	github.com/joho/godotenv v1.4.0
	service v0.0.0-00010101000000-000000000000
)

require (
	github.com/go-telegram-bot-api/telegram-bot-api v4.6.4+incompatible // indirect
	github.com/technoweenie/multipartstreamer v1.0.1 // indirect
)
