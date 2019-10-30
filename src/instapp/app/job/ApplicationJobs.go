package job

import (
	"instapp/app/models"

)

type UpdateAppJob struct {
    App models.Application
}

func (m UpdateAppJob) Run() {
	models.UpdateApp(m.App)
}



type DeleteAppJob struct {
    App models.Application
}

func (m DeleteAppJob) Run() {
	models.DeleteApp(m.App)
}



