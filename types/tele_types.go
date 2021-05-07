package types

import "os"

const telegramAPIBaseURL string = "https://api.telegram.org/bot"
const telegramAPISendMessage string = "/sendMessage"
const telegramTokenEnv string = "TELEGRAM_BOT_TOKEN"
const wordnikToken string = "WORDNIK_TOKEN"

// TelegramAPI is the api to which we should send the message to
var TelegramAPI string = telegramAPIBaseURL + os.Getenv(telegramTokenEnv) + telegramAPISendMessage

// WordnikAPI is the base url for getting words
var WordnikAPI string = "https://api.wordnik.com/v4/words.json/" + "wordOfTheDay" + "?api_key=" + os.Getenv(wordnikToken)

// Update is the type of request that telegram sends once u send message to the bot
type Update struct {
	UpdateID      int           `json:"update_id"`
	Message       Message       `json:"message"`
	CallbackQuery CallbackQuery `json:"callback_query"`
}

// Message is the structure of the message sent to the bot
type Message struct {
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
	Date int    `json:"date"`
}

// Chat indicates the conversation to which the message belongs.
type Chat struct {
	ID int `json:"id"`
}

// User is a telegram user
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"username"`
}

// CallbackQuery gives the structure of the callback that is received once user clicks on a button
type CallbackQuery struct {
	ID   string `json:"id"`
	From User   `json:"from"`
	Data string `json:"data"`
}

// Buttons is the structure for sending buttons with chat in telegram
type Buttons struct {
	InlineKeyboard [][]struct {
		Text         string `json:"text"`
		CallbackData string `json:"callback_data"`
	} `json:"inline_keyboard"`
}

// CreateInlineButtons creates inline buttons
func (but *Buttons) CreateInlineButtons(cols, rows int, arguments ...string) {
	but.InlineKeyboard = make([][]struct {
		Text         string "json:\"text\""
		CallbackData string "json:\"callback_data\""
	}, cols)

	but.InlineKeyboard[0] = make([]struct {
		Text         string "json:\"text\""
		CallbackData string "json:\"callback_data\""
	}, rows)

	iterator := 0
	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {
			but.InlineKeyboard[col][row].Text = arguments[iterator]
			but.InlineKeyboard[col][row].CallbackData = arguments[iterator+1]
			iterator = iterator + 2
		}
	}
}
