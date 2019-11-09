package controllers

import (
	"github.com/revel/revel"
	// "encoding/json"
	"instapp/app/models"
    // "github.com/revel/modules/jobs/app/jobs"
	// "instapp/app/job"
)

type ChatController struct {
	*revel.Controller
}

/**
* Fetch all resources
*/
func (c ChatController) Index(token string) revel.Result {
	// st1 := `
	// [{"Number": 1,"messages_count": 0}]`
	// chat := models.Chat{}
	// json.Unmarshal([]byte(st1), &chat)
	// return c.RenderJSON(chat)
	
	chatsFromDB, _ := models.SelectAllChats(token)
	chatsFromCash := models.GetRoomsFromCach(token)
	var chats []models.Chat
	chats = append(chats, chatsFromDB...)
	chats = append(chats, chatsFromCash...)
	return c.RenderJSON(chats)
}



/**
* Create resource
*/
func (c ChatController) Create(token string) revel.Result {
	Chat := models.AddChatToCach(token)
	return c.RenderJSON(Chat)
}

/**
* get single resource
*/
func (c ChatController) Show(token string, number int64) revel.Result {
	chat, err := models.SelectChatRoomByNumber(token, number)
	if err != nil {
		c.Response.Status = 404
		return c.RenderJSON(ErrorResponse{Message:"Resource not found"})
	}
	return c.RenderJSON(chat)
}


/**
* Delete resource
*/
func (c ChatController) Delete(token string, number int64) revel.Result {
	chat, err := models.DeleteChatRoom(token, number)
	if err != nil {
		c.Response.Status = 404
		return c.RenderJSON(ErrorResponse{Message:"Resource not found"})
	}
	return c.RenderJSON(chat)
}

