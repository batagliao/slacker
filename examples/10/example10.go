package main

import (
	"context"
	"log"

	"github.com/shomali11/slacker"
	"github.com/slack-go/slack"
)

func main() {
	bot := slacker.NewClient("<YOUR SLACK BOT TOKEN>")

	definition := &slacker.CommandDefinition{
		Description: "Echo a word!",
		Handler: func(request slacker.Request, response slacker.ResponseWriter) {
			word := request.Param("word")

			attachments := []slack.Block{}
			attachments = append(attachments, slack.NewContextBlock("1",
				slack.NewTextBlockObject("mrkdwn", "Hi!", false, false)),
			)

			response.Reply(word, slacker.WithBlocks(attachments))
		},
	}

	bot.Command("echo <word>", definition)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
