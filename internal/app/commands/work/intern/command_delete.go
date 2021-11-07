package intern

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c *WorkInternCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	internId, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	ok, err := c.internService.Remove(uint64(internId))
	var msgText string
	if err != nil {
		log.Printf("Server error. Fail to delete intern with id %d", internId)
		return
	} else {
		if ok {
			msgText = fmt.Sprintf("Intern with id %d was removed", internId)
		} else {
			msgText = fmt.Sprintf("There is no intern with id %d ", internId)
		}
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		msgText,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("WorkInternCommander.Get: error sending reply message to chat - %v", err)
	}
}
