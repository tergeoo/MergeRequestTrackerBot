package manager

import (
	"MRTrackerBot/model"
	"MRTrackerBot/utils"
	"fmt"
)

var devs = make(map[string]string)

func ProcessNote(webhook model.Webhook) {
	mr := webhook.CreateMergeRequestMessageFromComment()
	messageId := findMessageId(mr.Id)
	message := formatMessage(mr)
	utils.UpdateTelegramMessage(messageId, message)

	utils.SaveMessage(mr.Id, messageId)
}

func ProcessMergeRequest(webhook model.Webhook) {
	mr := webhook.CreateMergeRequestMessageFromMR()
	message := formatMessage(mr)
	messageResult := utils.SendMessage(message)

	utils.SaveMessage(mr.Id, messageResult.MessageID)
}

//[Ссылка на мр](https://gitlab)
//Есть коментарии 💩 @tergeo
//

//[Ссылка на мр](https://gitlab)
//апрув 👍 @tergeo
//

//[Ссылка на мр](https://gitlab)
//вмержен ✅

//[Ссылка на мр](https://gitlab)
//закрыт ❌

func formatMessage(mr model.MergeRequest) string {
	devs = utils.ReadConfig().Devs

	var messageBody string

	switch mr.State {
	case model.COMMENTED:
		messageBody = fmt.Sprintf("Есть коментарии 💩 @%s", devs[mr.AuthorUserName])
	case model.OPENED:
		messageBody = ""
	case model.APPROVED:
		messageBody = fmt.Sprintf("Апрув 👍 @%s", devs[mr.AuthorUserName])
	case model.MERGED:
		messageBody = "Вмержен ✅"
	case model.CLOSED:
		messageBody = "Закрыт ❌"
	}

	link := fmt.Sprintf("[%s](%s)\n", mr.Title, mr.Link)
	message := fmt.Sprintf("%s%s", link, messageBody)

	return message
}

func findMessageId(mrId int) int {
	saved := utils.ReadSavedMessages()
	return saved.MessageIds[mrId]
}
