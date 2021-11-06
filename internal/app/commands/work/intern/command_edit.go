package intern

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service/work/intern"
	"log"
	"strconv"
	"strings"
)

const editUsageMessage = "Usage: /edit__work__intern INTERN_ID INTERNSHIP_ID NAME\nTry /help__work__intern for more information"

func (c *WorkInternCommander) Edit(inputMessage *tgbotapi.Message) {
	errMsg, internToUpdate := extractArgsFromEditCommand(inputMessage)
	var response string
	if errMsg != "" {
		response = errMsg
	} else {
		err := c.internService.Update(internToUpdate.UniqueKey, *internToUpdate)
		if err != nil {
			response = err.Error()
		} else {
			response = fmt.Sprintf("Intern with id %d was successfully updated", internToUpdate.UniqueKey)
		}
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, response)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("WorkInternCommander.Get: error sending reply message to chat - %v", err)
	}
}

func extractArgsFromEditCommand(inputMessage *tgbotapi.Message) (string, *intern.Intern) {
	args := inputMessage.CommandArguments()

	if len(args) == 0 {
		return editUsageMessage, nil
	}

	splitArgs := strings.Split(args, " ")

	if len(splitArgs) < 3 {
		return editUsageMessage, nil
	}

	internId, err := strconv.Atoi(splitArgs[0])

	if err != nil {
		return "The first argument must be intern Id.\n" + editUsageMessage, nil
	}

	internshipId, err := strconv.Atoi(splitArgs[1])

	if err != nil {
		return "The first argument must be internship Id.\n" + editUsageMessage, nil
	}

	name := strings.Join(splitArgs[2:], " ")

	newIntern := intern.NewIntern(name, uint64(internshipId))
	newIntern.UniqueKey = uint64(internId)
	return "", newIntern
}
