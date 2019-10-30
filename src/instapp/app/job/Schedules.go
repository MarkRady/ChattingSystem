package job

import (
	"instapp/app/models"
	// "fmt"
)

type SaveMessagesToDB struct {
	//
}

func (s SaveMessagesToDB) Run() {
	// get apps
	apps, _ := models.SelectAllApps()

	for _, app := range apps {
		Token := app.Token
		// get chats
		chats, _ := models.SelectAllChats(Token)
		for _, chat := range chats {
			roomNumber := chat.Number
			cachedMsgs := models.GetMessagesFromCach(Token, roomNumber)
			for _,msg := range cachedMsgs {
				models.InsertMessage(Token, roomNumber, msg.Body)
			}
			models.ClearCachStorageAfterQueueForMessages(Token, roomNumber)
		}


	}
}




type SaveChatsToDB struct {
	//
}

func (s SaveChatsToDB) Run() {
	// get apps
	apps, _ := models.SelectAllApps()

	for _, app := range apps {
		Token := app.Token
		// get chats
		chatsInCach := models.GetRoomsFromCach(Token)

		for range chatsInCach {
			models.InsertChat(Token)
		}
		models.ClearCachStorageAfterQueueForChats(Token)

	

	}
}




