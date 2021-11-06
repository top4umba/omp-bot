package intern

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service/work/intern"
	"log"
	"strconv"
	"strings"
)

const newUsageMessage = "Usage: /new__work__intern INTERNSHIP_ID NAME\nTry /help__work__intern for more information"

func (c *WorkInternCommander) New(inputMessage *tgbotapi.Message) {
	var response string
	errMsg, newIntern := extractArgsFromNewCommand(inputMessage)
	if errMsg != "" {
		response = errMsg
	} else {
		internID, err := c.internService.Create(*newIntern)
		if err != nil {
			response = "Server error. Failed to create new intern"
		} else {
			response = fmt.Sprintf("Intern %s was successfully created and assigned with id=%d", newIntern.Name, internID)
		}
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, response)
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("WorkInternCommander.Get: error sending reply message to chat - %v", err)
	}
}

func extractArgsFromNewCommand(inputMessage *tgbotapi.Message) (string, *intern.Intern) {
	args := inputMessage.CommandArguments()

	if len(args) == 0 {
		return newUsageMessage, nil
	}

	splitArgs := strings.Split(args, " ")

	if len(splitArgs) < 2 {
		return newUsageMessage, nil
	}

	internshipId, err := strconv.Atoi(splitArgs[0])

	if err != nil {
		return "The first argument must be internship Id.\n" + newUsageMessage, nil
	}

	name := strings.Join(splitArgs[1:], " ")

	newIntern := intern.NewIntern(name, uint64(internshipId))
	return "", newIntern
}
