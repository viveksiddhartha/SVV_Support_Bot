package main

import (
	"GoBotTel/types"
	"GoBotTel/utils"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func handler(res http.ResponseWriter, r *http.Request) {
	update, err := utils.ParseTelegramUpdate(r)
	if err != nil {
		log.Printf("error parsing update, %s", err.Error())
		return
	}
	log.Printf(update.Message.Text)
	log.Println(update.CallbackQuery.Data)

	userCmd, err := types.ParseCommand(update.Message.Text)
	callbackData := update.CallbackQuery.Data

	log.Println(userCmd)
	// TODO Parse Arguments of the command too

	var output string
	var keyboard []byte = nil

	if err != nil && callbackData == "" {
		output = "Sorry you entered the wrong command. Here are the list of supported commands \n"
		for command, desc := range types.Commands {
			output += command + " : " + desc + "\n"
		}
	}

	// Check if there is a callback from an inline button
	if callbackData == "GN" {
		output, err = utils.GetNewsForResponse("the-times-of-india")
		log.Println(output)

		if err != nil {
			resp, err := utils.SendTextToTelegram(update.CallbackQuery.From.ID, output, keyboard)
			if err != nil {
				log.Printf("got error %s from telegram, response body is %s", err.Error(), resp)
			} else {
				log.Printf("punchline %s successfully distributed to chat id %d", output, update.CallbackQuery.From.ID)
			}
			return
		}
	} else if callbackData == "BN" {
		output, err = utils.GetNewsForResponse("business-insider")
		if err != nil {
			resp, err := utils.SendTextToTelegram(update.CallbackQuery.From.ID, output, keyboard)
			if err != nil {
				log.Printf("got error %s from telegram, response body is %s", err.Error(), resp)
			} else {
				log.Printf("punchline %s successfully distributed to chat id %d", output, update.CallbackQuery.From.ID)
			}
			return
		}
	} else if callbackData == "TN" {
		output, err = utils.GetNewsForResponse("techcrunch")
		if err != nil {
			resp, err := utils.SendTextToTelegram(update.CallbackQuery.From.ID, output, keyboard)
			if err != nil {
				log.Printf("got error %s from telegram, response body is %s", err.Error(), resp)
			} else {
				log.Printf("punchline %s successfully distributed to chat id %d", output, update.CallbackQuery.From.ID)
			}
			return
		}
	} else {
		// No callbacks hence start looking for commands in user inputs
		if userCmd == "/start" {
			output = "Hello I'm " + types.BotName + " I can do the following things for you \n\n"
			for command, desc := range types.Commands {
				output += command + " : " + desc + "\n"
			}
		} else if userCmd == "/news" {
			but := types.Buttons{}
			but.CreateInlineButtons(1, 3, "General News", "GN", "Business News", "BN", "Tech News", "TN")

			keyboard, err = json.Marshal(but)
			if err != nil {
				log.Printf(err.Error())
				return
			}

			output += "Great.. Almost there.. Please choose which kind of news you want\n"
		} else if userCmd == "/word" {
			output, err = utils.GetWordOfTheDay()
			if err != nil {
				resp, err := utils.SendTextToTelegram(update.Message.Chat.ID, output, keyboard)
				if err != nil {
					log.Printf("got error %s from telegram, response body is %s", err.Error(), resp)
				} else {
					log.Printf("punchline %s successfully distributed to chat id %d", output, update.Message.Chat.ID)
				}
				return
			}
		}
		// else if userCmd == "/port" {
		// 	output = "Sorry... This command is under implementation..."
		// }
	}

	var resp string
	if update.Message.Chat.ID != 0 {
		resp, err = utils.SendTextToTelegram(update.Message.Chat.ID, output, keyboard)
	} else if update.CallbackQuery.From.ID != 0 {
		resp, err = utils.SendTextToTelegram(update.CallbackQuery.From.ID, output, keyboard)
	}

	if err != nil {
		log.Printf("got error %s from telegram, response body is %s", err.Error(), resp)
	} else {
		log.Printf("punchline %s successfully distributed to chat id %d", output, update.Message.Chat.ID)
	}
}

func main() {
	log.Printf(os.Getenv("PORT"))
	http.ListenAndServe("0.0.0.0:"+os.Getenv("PORT"), http.HandlerFunc(handler))
}
