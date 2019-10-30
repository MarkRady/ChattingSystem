package controllers

import (
	"github.com/revel/revel"
	"instapp/app/models"
    // "sync"
	// "time"
	// "fmt"
)


type MessageController struct {
	*revel.Controller
}



/**
* Fetch all resources
*/
func (c MessageController) Index(token string, chat_number int64) revel.Result {
	messagesFromCash := models.GetMessagesFromCach(token, chat_number)
	messagesFromDB, _ := models.SelectAllMessages(token, chat_number)
	var msgs []models.Message
	msgs = append(msgs, messagesFromDB...)
	msgs = append(msgs, messagesFromCash...)
	return c.RenderJSON(msgs)
	// return c.Render();
}


/**
* Create resource
*/
func (c MessageController) Create(token string, chat_number int64) revel.Result {
	var Body = c.Params.Form.Get("Body")
	Message := models.AddMessageToCach(Body, token, chat_number)
	return c.RenderJSON(Message)
}



/**
* Create resource
*/
/*func (c MessageController) Create(token string, chat_number int64) revel.Result {
	var NewNumber int64
	go GetLastId(&NewNumber, token, chat_number)
	var Body = c.Params.Form.Get("Body")
	Message, err := models.InsertMessage(token, chat_number, Body, NewNumber)
	Lock.Unlock()
	if err != nil {
		c.Response.Status = 404
		return c.RenderJSON(ErrorResponse{Message:"Resource not found"})
	}
	return c.RenderJSON(Message)
}*/

/**
* Search resource
*/
func (c MessageController) Search() revel.Result {
	var Body = c.Params.Form.Get("Body")
	Messages := models.SearchMessages(Body)
	return c.RenderJSON(Messages)
}



/**
* get single resource
*/
func (c MessageController) Show(token string, chat_number int64, message_number int64) revel.Result {
	Message, err := models.GetSingleMessage(token, chat_number, message_number)
	if err != nil {
		c.Response.Status = 404
		return c.RenderJSON(ErrorResponse{Message:"Resource not found"})
	}
	return c.RenderJSON(Message)
}

/**
* get single resource
*/
func (c MessageController) Update(token string, chat_number int64, message_number int64) revel.Result {
	var Body = c.Params.Form.Get("Body")
	Message, err := models.UpdateMessage(token, chat_number, message_number, Body)
	if err != nil {
		c.Response.Status = 404
		return c.RenderJSON(ErrorResponse{Message:"Resource not found"})
	}
	return c.RenderJSON(Message)
}

func (c MessageController) Delete(token string, chat_number int64, message_number int64) revel.Result {
	Message, err := models.DeleteMessage(token, chat_number, message_number)
	if err != nil {
		c.Response.Status = 404
		return c.RenderJSON(ErrorResponse{Message:"Resource not found"})
	}
	return c.RenderJSON(Message)
}



/**
* Delete resource
*/
/*func (c MessageController) Delete(token string, number int64) revel.Result {
	chat, err := models.DeleteChatRoom(token, number)
	if err != nil {
		c.Response.Status = 404
		return c.RenderJSON(ErrorResponse{Message:"Resource not found"})
	}
	return c.RenderJSON(chat)
}
*/

/*func (c MessageController) Create(token string, chat_number int64) revel.Result {
	var NewNumber int64
    Wait.Add(1)
    	Lock.Lock();
		go GetLastId(&NewNumber, token, chat_number)
		// fmt.Println("hi:", NewNumber)
    Wait.Wait()
    // go GetLastId(&NewNumber, token, chat_number)

	var Body = c.Params.Form.Get("Body")
	Message, err := models.InsertMessage(token, chat_number, Body, NewNumber)
	Lock.Unlock()
	if err != nil {
		c.Response.Status = 404
		return c.RenderJSON(ErrorResponse{Message:"Resource not found"})
	}
	return c.RenderJSON(Message)
}*/