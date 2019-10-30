package controllers

import (
	"github.com/revel/revel"
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
	apps, err := models.SelectAllChats(token)
	if err != nil {
		c.Response.Status = 404
		return c.RenderJSON(ErrorResponse{Message:"Resource not found"})
	}
	return c.RenderJSON(apps)
}



/**
* Create resource
*/
func (c ChatController) Create(token string) revel.Result {
	// Chat := models.Chat{}
	// jobs.Now(job.InsertChatJob{Token: token, Chat: &Chat})
	Chat, err := models.InsertChat(token)
	if err != nil {
		c.Response.Status = 404
		return c.RenderJSON(ErrorResponse{Message:"Resource not found"})
	}
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

