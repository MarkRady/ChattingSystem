package models

import (
	"errors"
	"log"
	// "github.com/revel/revel/cache"

)

type Chat struct {
    Id             int64  `json:"-"`   
    Number         int64     
    MessagesCount  int64  `db:"messages_count" json:"messages_count"`
    ApplicationId  int64  `json:"-"`   
}


/*
func getCachNameChat(token string, roomNumber int64, prefix string) string {
	return prefix+"_"+token+"_"+string(roomNumber)
}


func GetRoomsFromCach(token string, roomNumber int64) []Message {
	var chats []Chat
	cach_key := getCachNameChat(token, roomNumber, "messagesCach")
	cache.Get(cach_key, &chats);
	return chats
}

func getNewNumberCachForMsg(token string, roomNumber int64) int64 {
	msgs := GetRoomsFromCach(token, roomNumber)
	if msgs == nil && len(msgs) == 0 {
		lastMsg := Message{}
		Chat, _ := SelectChatRoomByNumber(token, roomNumber)
		lastMsg, _ = GetLastMessage(Chat.Id)
		return lastMsg.Number+1
	}

	lastMsg := msgs[len(msgs)-1]
	return lastMsg.Number + 1
}

func AddMessageToCach(body string, token string, roomNumber int64) (Message) {
	MessageModel := Message{}
	MessageModel.Body = body
	isWCach := getCachNameChat(token, roomNumber, "isWrite")

	for {
		var isWrite int;
		cache.Get(isWCach, &isWrite)
		if isWrite == 0 {
			break
		}
	}
	//start writing
    cache.Set(isWCach, 1, 60*time.Minute)

	MessageModel.Number = getNewNumberCachForMsg(token, roomNumber)

	msgs := []Message{}
	msgs = GetRoomsFromCach(token, roomNumber);
	msgs = append(msgs, MessageModel)
	cach_key := getCachNameChat(token, roomNumber, "messagesCach")
    cache.Set(cach_key, msgs, 60*time.Minute)
    // end of writeing
    cache.Set(isWCach, 0, 60*time.Minute)

    return MessageModel;
}

func ClearCachStorageAfterQueueForMessages(token string, roomNumber int64) {
	cach_key := "messagesCach_"+token+"_"+string(roomNumber)
	cache.Delete(cach_key)
}*/



func InsertChat(token string) (Chat, error) {
	Chat := &Chat{}
	// Fetch application
	app, err := SelectOneApp(token)
	if err != nil {
		log.Fatalln("ChatModel", err)
		return *Chat, errors.New("App not found")
	}
	// generate chat number 
	LastChat, _ := GetLastChat(app.Id)
	app.ChatCounts++
	//Create chat
	Chat.Number = LastChat.Number+1
	Chat.ApplicationId = app.Id

	err = dbmap.Insert(Chat)
	//Increase number of ChatCounts in application
	UpdateApp(app)

	if err != nil {
		log.Fatalln("ChatModel", err)
		return *Chat, errors.New("error insert")
	}
	return *Chat, nil
}

func UpdateChat(Chat Chat) (Chat, error) {
	_, err := dbmap.Update(&Chat)
	if err != nil {
		return Chat, errors.New("error update Chat")
	}
	return Chat, nil
}


func SelectAllChats(token string) ([]Chat, error) {
	Chats := []Chat{}
	//Get app 
	app, err := SelectOneApp(token)
	if err != nil {
		log.Fatalln("ChatModel", err)
		return Chats, errors.New("App not found")
	}
	// select all chats of this application
	_, err = dbmap.Select(&Chats, "SELECT * FROM `chats` WHERE `ApplicationId`=?", app.Id)
	if err != nil {
		log.Fatalln("ChatModel", err)
		return Chats, errors.New("error in select")
	}
	return Chats, nil
}

func SelectChatRoomByNumber(token string, number int64) (Chat, error) {
	Chat := &Chat{}
	// Fetch application
	app, err := SelectOneApp(token)
	if err != nil {
		log.Fatalln("ChatModel", err)
		return *Chat, errors.New("App not found")
	}
	// get chat room
	err = dbmap.SelectOne(&Chat, "SELECT * FROM `chats` WHERE `ApplicationId`=? AND `Number`=?;", app.Id, number)
	if err != nil {
		return *Chat, errors.New("error select")
	}
	return *Chat, nil
}

func GetLastChat(AppId int64) (Chat, error) {
	Chat := &Chat{}
	// get chat room
	err := dbmap.SelectOne(&Chat, "SELECT * FROM `chats` WHERE `ApplicationId`=? ORDER BY `Id` DESC limit 1;", AppId)
	if err != nil {
		return *Chat, errors.New("error select")
	}
	return *Chat, nil
}

func DeleteChatByApp(App Application) (error) {
	// delete chats
    _, err := dbmap.Exec("DELETE FROM `chats` WHERE ApplicationId=?", App.Id)
	if err != nil {
		return errors.New("error select")
	}

    // Delete Messages
    // 1) Get chats
    chats, _ := SelectAllChats(App.Token)
    // 2) loop through chats 
    for _,chat := range chats {
    	DeleteMessageByChat(chat)
    }


	return nil
}




func DeleteChatRoom(token string, number int64) (Chat, error) {
	Chat := &Chat{}
	// Fetch application
	app, err := SelectOneApp(token)
	if err != nil {
		log.Fatalln("ChatModel", err)
		return *Chat, errors.New("App not found")
	}
	// get chat room
	err = dbmap.SelectOne(&Chat, "SELECT * FROM `chats` WHERE `ApplicationId`=? AND `Number`=?;", app.Id, number)
	if err != nil {
		return *Chat, errors.New("error select")
	}
	// delete chatroom
	_, err = dbmap.Delete(Chat)
	if err != nil {
		return *Chat, errors.New("can not delete chat")
	}
	//decrease number of rooms
	app.ChatCounts--
	UpdateApp(app)
	DeleteMessageByChat(*Chat)


	return *Chat, nil
}



