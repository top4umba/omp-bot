package intern

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

const getUsageMessage = "Usage: /get__work__intern INTERN_ID\nTry /help__work__intern for more information"

func (c *WorkInternCommander) Get(inputMessage *tgbotapi.Message) {
	errMsg, internId := extractArgsFromGetCommand(inputMessage)
	var response string
	if errMsg != "" {
		response = errMsg
	} else {
		result, err := c.internService.Describe(uint64(internId))
		if err != nil {
			response = fmt.Sprintf("Cannot find intern with id %d", internId)
		} else {
			response = fmt.Sprint(result)
		}
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, response)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("WorkInternCommander.Get: error sending reply message to chat - %v", err)
	}
}

func extractArgsFromGetCommand(inputMessage *tgbotapi.Message) (string, int) {
	arg := inputMessage.CommandArguments()

	if len(arg) == 0 {
		return getUsageMessage, -1
	}

	internId, err := strconv.Atoi(arg)

	if err != nil {
		return "An argument must be intern Id.\n" + getUsageMessage, -1
	}

	return "", internId
}
