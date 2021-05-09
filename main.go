package main

import (
	"GoBotTel/common/datastore"
	"fmt"
	"log"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	b, err := tb.NewBot(tb.Settings{
		// You can also set custom API URL.
		// If field is empty it equals to "https://api.telegram.org".
		URL: "",

		Token:  "1813786957:AAEE_fDJnuL02PLGjElBSd8gAawTxtNGHXA",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	fmt.Println("SVV_Support_bot has been authenticated")

	datastore.MongoConn()
	fmt.Println("Connection successful")

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle(tb.OnText, func(m *tb.Message) {

		inlineBtn1 := tb.InlineButton{
			Unique: "moon",
			Text:   "VASU 🌚",
		}

		inlineBtn2 := tb.InlineButton{
			Unique: "sun",
			Text:   "SHWETA 🌞",
		}

		b.Handle(&inlineBtn1, func(c *tb.Callback) {
			// Required for proper work
			b.Respond(c, &tb.CallbackResponse{
				ShowAlert: false,
			})
			// Send messages here
			b.Send(c.Sender, "VASU Saitan baccha hai")
		})

		b.Handle(&inlineBtn2, func(c *tb.Callback) {
			b.Respond(c, &tb.CallbackResponse{
				ShowAlert: false,
			})
			b.Send(c.Sender, "Shweta pyari Guddi hai")
		})
		inlineKeys := [][]tb.InlineButton{
			[]tb.InlineButton{inlineBtn1, inlineBtn2},
		}
		b.Handle("/Pingu", func(m *tb.Message) {
			b.Send(
				m.Sender,
				"Vasu or Shweta, you choose",
				&tb.ReplyMarkup{InlineKeyboard: inlineKeys})
		})

		/* 		fmt.Printf("Text received from ")
		   		b.Handle("/hello", func(m *tb.Message) {
		   			b.Send(m.Sender, "Hello World!")
		   		}) */
	})

	b.Handle(tb.OnPhoto, func(m *tb.Message) {
		// photos only
	})

	b.Handle(tb.OnChannelPost, func(m *tb.Message) {
		// channel posts only
	})

	b.Handle(tb.OnQuery, func(q *tb.Query) {
		// incoming inline queries
	})
	b.Handle(tb.OnSticker, func(m *tb.Message) {
		fmt.Printf("Sticker received from ")

	})

	b.Start()
}
