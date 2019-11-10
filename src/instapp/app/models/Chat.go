package models

import (
	"errors"
	"log"
	"github.com/revel/revel/cache"
	"time"
	"encoding/json"
	// "fmt"
)

type Chat struct {
    Id             int64  `json:"-"`   
    Number         int64     
    MessagesCount  int64  `db:"messages_count" json:"messages_count"`
    ApplicationId  int64  `json:"-"`   
}


func (chat Chat) ToString() string {
	res, err := json.Marshal(chat)
	if err != nil {
		return "Data Is empty"
	}
	return string(res)
}

func jsonToChat(data string) Chat {
	chat := Chat{}
	json.Unmarshal([]byte(data), &chat)
	return chat
}




func getCachNameChat(token string, prefix string) string {
	return prefix+"_"+token
}


func GetRoomsFromCach(token string) []Chat {
	cachKey := getCachNameChat(token, "chatCash") 
	var chats []Chat
	var chatsAsJson []byte // declare variable to contain data from caching as bytes
	cache.Get(cachKey, &chatsAsJson) // get data from cach and assign to chatsAsJson
	json.Unmarshal(chatsAsJson, &chats) // convert bytes to struct of Chat
	return chats
}

func GetRoomsFromCachAsJson(token string) []string {
	var chatsAsJson []string
	cachKey := getCachNameChat(token, "chatCash")
	cache.Get(cachKey, &chatsAsJson);
	
	return chatsAsJson
}



func getNewNumberCachForChats(token string) int64 {
	chats := GetRoomsFromCach(token)
	if chats == nil && len(chats) == 0 {
		app, _ := SelectOneApp(token)
		LastChat, _ := GetLastChat(app.Id)
		return LastChat.Number+1
	}
	lastChat := chats[len(chats)-1]
	return lastChat.Number + 1
}


func AddChatToCach(token string) Chat {
	ChatModel := Chat{}
	isWCach := getCachNameChat(token, "isWriteChat")

	for {
		var isWrite int;
		cache.Get(isWCach, &isWrite)
		if isWrite == 0 {
			break
		}
	}
	//start writing
    cache.Set(isWCach, 1, 60*time.Minute)

	ChatModel.Number = getNewNumberCachForChats(token)
	addToCachAsJson(ChatModel, token)
	
    // end of writeing
    cache.Set(isWCach, 0, 60*time.Minute)

    return ChatModel;
}

func addToCachAsJson(chat Chat, token string) {

	chatsInCach := []Chat{}
	chatsInCach = GetRoomsFromCach(token);
	chatsInCach = append(chatsInCach, chat)
	
	convertChatToJsonArray,_ := json.MarshalIndent(&chatsInCach, "", "\t") // convert array of chats to json array and with json intendents

	cach_key := getCachNameChat(token, "chatCash")
    cache.Set(cach_key, convertChatToJsonArray, 60*time.Minute) // save to cach

}

func ClearCachStorageAfterQueueForChats(token string) {
	cach_key := getCachNameChat(token, "chatCash")
	cache.Delete(cach_key)
}



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
		logErrDb(err, "db err")
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



