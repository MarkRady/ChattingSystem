package models

import (
	"errors"
	"log"
	"context"
	"encoding/json"
    elasticapi "gopkg.in/olivere/elastic.v7"
	"github.com/revel/revel/cache"
	"time"

)


/**
* Caching layer
*/

func getCachNameMsg(token string, roomNumber int64, prefix string) string {
	return prefix+"_"+token+"_"+string(roomNumber)
}


func GetMessagesFromCach(token string, roomNumber int64) []Message {
	var Messages []Message
	cach_key := getCachNameMsg(token, roomNumber, "messagesCach")
	cache.Get(cach_key, &Messages);
	return Messages
}

func getNewNumberCachForMsg(token string, roomNumber int64) int64 {
	msgs := GetMessagesFromCach(token, roomNumber)
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
	// isWCach := "isWrite"+token+"_"+string(roomNumber)
	isWCach := getCachNameMsg(token, roomNumber, "isWriteMessage")

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
	msgs = GetMessagesFromCach(token, roomNumber);
	msgs = append(msgs, MessageModel)
	cach_key := getCachNameMsg(token, roomNumber, "messagesCach")

    cache.Set(cach_key, msgs, 60*time.Minute)
    // end of writeing
    cache.Set(isWCach, 0, 60*time.Minute)

    return MessageModel;
}

func ClearCachStorageAfterQueueForMessages(token string, roomNumber int64) {
	cach_key := getCachNameMsg(token, roomNumber, "messagesCach")
	cache.Delete(cach_key)
}



/**
*  Database layer
*/

type Message struct {
    Id             int64  `json:"-"`   
    Number         int64     
    ChatId  int64  `json:"-"`
    Body string 
}

// to save as json for elasticsearch
func (message Message) ToString() string {
	res, err := json.MarshalIndent(message, "", "")
	if err != nil {
		return "Data Is empty"
	}
	return string(res)
}



const (
	ELS_IndexMessage = "messages_index"
	ELS_DocMessage   = "message"
)

func InsertMessage(token string, chat_number int64, body string) (Message, error) {
	var err error
	MessageModel := &Message{}
	lastMsg := Message{}
	Chat := Chat{}
	// Fetch Chat room
	Chat, err = SelectChatRoomByNumber(token, chat_number)
	if err != nil {
		log.Fatalln("MessageModel", err)
		return *MessageModel, errors.New("Room not found")
	}
	// Start a db transaction
    // dbmap.Exec("START TRANSACTION;")
	lastMsg, _ = GetLastMessage(Chat.Id)
	//Create Message
	MessageModel.Body = body
	MessageModel.ChatId = Chat.Id
	MessageModel.Number = lastMsg.Number+1
	// Start a db transaction
	err = dbmap.Insert(MessageModel)
	if err != nil {
	    // dbmap.Exec("ROLLBACK;")
		log.Fatalln("MessageModel", err)
		return *MessageModel, errors.New("error insert")
	}
	// Commit the transaction
    // dbmap.Exec("COMMIT;")

	//Increase number of MessageCount in Chat
	Chat.MessagesCount++
	UpdateChat(Chat)
	// init background context
	ctx := context.Background()
	// init elastic connection
	elasticClient, err := InitElastic(ctx, ELASTICSEARCH_URI, false, -1)
	// imsert message to elasticsearch
	InsertMessageToElasticSearch(ctx, elasticClient, *MessageModel)


	
	return *MessageModel, nil
}

func GetLastMessage(ChatId int64) (Message, error) {
	Message := &Message{}
	// Start a db transaction
    // dbmap.Exec("START TRANSACTION;")
	// get Message room
	// TODO: Search in elasticsearch
	err := dbmap.SelectOne(&Message, "SELECT * FROM `messages` WHERE `ChatId`=? ORDER BY `Id` DESC limit 1;", ChatId)
	if err != nil {
		dbmap.Exec("ROLLBACK;")
		return *Message, errors.New("error select")
	}
	// Commit the transaction
    // dbmap.Exec("COMMIT;")
	return *Message, nil
}

func SearchMessages(body string) ([]Message) {
	ctx := context.Background()
	elasticClient, _ := InitElastic(context.Background(), ELASTICSEARCH_URI, false, -1)
	return SearchForMessages(ctx, elasticClient, body)
}

func GetSingleMessage(token string, chat_number int64, message_number int64) (Message, error){
	Message := &Message{}
	Chat, err := SelectChatRoomByNumber(token, chat_number)
	if err != nil {
		return *Message, errors.New("Chat room not found")
	}
	// TODO: Search in elasticsearch
	err = dbmap.SelectOne(&Message, "SELECT * FROM `messages` WHERE `ChatId`=? AND `Number`=? ORDER BY `Id` DESC limit 1;", Chat.Id, message_number)
	if err != nil {
		return *Message, errors.New("No message found")
	}
	return *Message, nil
}

func UpdateMessage(token string, chat_number int64, message_number int64, Body string) (Message, error){
	Message := &Message{}
	Chat, err := SelectChatRoomByNumber(token, chat_number)
	if err != nil {
		return *Message, errors.New("Chat room not found")
	}
	err = dbmap.SelectOne(&Message, "SELECT * FROM `messages` WHERE `ChatId`=? AND `Number`=? ORDER BY `Id` DESC limit 1;", Chat.Id, message_number)
	if err != nil {
		return *Message, errors.New("No message found")
	}
	Message.Body = Body
	_, err = dbmap.Update(Message)
	if err != nil {
		return *Message, errors.New("No message found")
	}
	// TODO: Update message in elasticsearch
	return *Message, nil
}



func SelectAllMessages(token string, chat_number int64) ([]Message, error) {
	Messages := []Message{}
	//Get Chat 
	Chat, err := SelectChatRoomByNumber(token, chat_number)
	if err != nil {
		return Messages, errors.New("Chat Room not found")
	}
	// select all messages of this application
	_, err = dbmap.Select(&Messages, "SELECT * FROM `messages` WHERE `ChatId`=?", Chat.Id)
	if err != nil {
		return Messages, errors.New("error in select messages")
	}
	return Messages, nil
}

func DeleteMessage(token string, chat_number int64, message_number int64) (Message, error){
	Message := &Message{}
	Chat, err := SelectChatRoomByNumber(token, chat_number)
	if err != nil {
		return *Message, errors.New("Chat room not found")
	}
	err = dbmap.SelectOne(&Message, "SELECT * FROM `messages` WHERE `ChatId`=? AND `Number`=? ORDER BY `Id` DESC limit 1;", Chat.Id, message_number)
	if err != nil {
		return *Message, errors.New("No message found")
	}
	// Delete message
	_, err = dbmap.Delete(Message)
	if err != nil {
		return *Message, errors.New("can not delete message")
	}
	// decrease count messages
	Chat.MessagesCount--
	UpdateChat(Chat)
	// TODO: Delete message from elasticsearch
	return *Message, nil
}


func DeleteMessageByChat(Chat Chat) (error){
    _, err := dbmap.Exec("DELETE FROM `messages` WHERE ChatId=?", Chat.Id)
	// TODO: Delete message from elasticsearch
	if err != nil {
		return errors.New("Can not delete messages")
	}
	
	return nil
}



// CreateIndexIfDoesNotExist ...
func CreateIndexIfDoesNotExistForMessages(ctx context.Context, client *elasticapi.Client) error {
	exists, err := client.IndexExists(ELS_IndexMessage).Do(ctx)
	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	res, err := client.CreateIndex(ELS_IndexMessage).Do(ctx)

	if err != nil {
		return err
	}

	if !res.Acknowledged {
		return errors.New("CreateIndex was not acknowledged. Check that timeout value is correct.")
	}

	return nil
}




func InsertMessageToElasticSearch(ctx context.Context, elasticClient *elasticapi.Client, Message Message) {
	CreateIndexIfDoesNotExistForMessages(ctx, elasticClient)
	// insert data in elasticsearch
	_, err := elasticClient.Index().Index(ELS_IndexMessage).Type(ELS_DocMessage).BodyJson(Message).Do(ctx)
		if err != nil {
			log.Fatalln("MessageModel", err)
		}
	

	// Flush data (need for refreshing data in index) after this command possible to do get.
	elasticClient.Flush().Index(ELS_IndexMessage).Do(ctx)
}

// search messages 
func SearchForMessages(ctx context.Context, elasticClient *elasticapi.Client, Body string) []Message {
	// query := elasticapi.NewBoolQuery()
	// musts := []elasticapi.Query{elasticapi.NewTermQuery("Id", 26)}
	// query = query.Must(musts...)
    // query := elasticapi.MatchAllQuery{}
    multiQuery := elasticapi.NewMultiMatchQuery(
        Body,
        "Body",
    ).Type("phrase_prefix")

    // matchQuery := elastic.NewMatchQuery("id", key)
    query := elasticapi.NewBoolQuery().Must(multiQuery)

 
	searchResult, err := elasticClient.Search().
		Index(ELS_IndexMessage). // search in index
		Query(query).     // specify the query
		Do(ctx)           // execute
	if err != nil {
		log.Fatalln("Error during execution SearchForMessages : %s", err.Error())
	}

	return convertSearchResultToMessages(searchResult)
}

// Convert messages from elastic entity to struct array
func convertSearchResultToMessages(searchResult *elasticapi.SearchResult) []Message {
	var Messages []Message
	for _, hit := range searchResult.Hits.Hits {
		var Message Message
		err := json.Unmarshal(hit.Source, &Message)
		if err != nil {
			log.Fatalln("Can't deserialize 'message' object : %s", err.Error())
			continue
		}
		Messages = append(Messages, Message)
	}
	return Messages
}
