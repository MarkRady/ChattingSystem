package controllers

import (
	"github.com/revel/revel"
    "github.com/revel/modules/jobs/app/jobs"
	"instapp/app/models"
	"instapp/app/job"
)

type ApplicationController struct {
	*revel.Controller
}


func (c ApplicationController) Home() revel.Result {
    return c.RenderTemplate("App/Home.html")
}



/**
* Fetch all resources
*/
func (c ApplicationController) Index() revel.Result {
	apps, err := models.SelectAllApps();
	if err != nil {
		c.Response.Status = 500
		return c.RenderJSON(ErrorResponse{Message:"Internal server error"})
	}
	return c.RenderJSON(apps)
}


/**
* Create resource
*/
func (c ApplicationController) Create() revel.Result {
	var Name = c.Params.Form.Get("Name")
	app, err := models.InsertApplication(Name)
	if err != nil {
		c.Response.Status = 500
		return c.RenderJSON(ErrorResponse{Message:"Internal server error"})
	}
	return c.RenderJSON(app)
}

/**
* get single resource
*/
func (c ApplicationController) Show(token string) revel.Result {
	app, err := models.SelectOneApp(token)
	if err != nil {
		c.Response.Status = 404
		return c.RenderJSON(ErrorResponse{Message:"Resource not found"})
	}
	return c.RenderJSON(app)
}


/**
* Update resource
*/
func (c ApplicationController) Update(token string) revel.Result {
	var Name = c.Params.Form.Get("Name")
	app, err := models.SelectOneApp(token)
	app.Name = Name
	jobs.Now(job.UpdateAppJob{App: app})
	if err != nil {
		c.Response.Status = 404
		return c.RenderJSON(ErrorResponse{Message:"Resource not found"})
	}

	return c.RenderJSON(app)
}

/**
* Delete resource
*/
func (c ApplicationController) Delete(token string) revel.Result {
	app, err := models.SelectOneApp(token)
	jobs.Now(job.DeleteAppJob{App: app})
	
	if err != nil {
		c.Response.Status = 404
		return c.RenderJSON(ErrorResponse{Message:"Resource not found"})
	}

	return c.RenderJSON(app)
}


