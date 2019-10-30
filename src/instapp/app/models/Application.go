package models

import (
	"errors"
	"time"
	"crypto/md5"
    "encoding/hex"
	"strings"
	"strconv"
	"log"
)


type Application struct {
    Id          int64  `json:"-"`   
    Name        string     
    ChatCounts  int64 `db:"chat_count" json:"chat_count"`
    Token  		string
}


func InsertApplication(Name string) (Application, error) {
	var App = &Application{
		Name: Name, 
		Token : GenerateUniqId(),
	}

	err := dbmap.Insert(App)
	if err != nil {
		log.Fatalln("ApplicationModel", err)
		return *App, errors.New("error insert")
	}
	return *App, nil
}

func SelectAllApps() ([]Application, error) {
	Apps := []Application{}
	_, err := dbmap.Select(&Apps, "SELECT * FROM `applications`")
	if err != nil {
		log.Fatalln("ApplicationModel", err)
		return Apps, errors.New("error in select")
	}
	return Apps, nil
}

func SelectOneApp(token string) (Application, error) {
	var App = &Application{}
	err := dbmap.SelectOne(&App, "SELECT * FROM `applications` WHERE `token`=?;", token)
	if err != nil {
		return *App, errors.New("error select")
	}
	return *App, nil
}

func UpdateApp(App Application) (Application, error) {
	_, err := dbmap.Update(&App)
	if err != nil {
		return App, errors.New("error update app")
	}
	return App, nil
}

func DeleteApp(App Application) (Application, error) {
	_, err := dbmap.Delete(&App)
	// Delete chats
	DeleteChatByApp(App)
	if err != nil {
		return App, errors.New("error delete app")
	}
	return App, nil
}





func GenerateUniqId() string {
    now := time.Now()
	var string = strconv.FormatInt(int64(now.Nanosecond()),10)
    h := md5.New()
    h.Write([]byte(strings.ToLower(string)))
    return hex.EncodeToString(h.Sum(nil))
 }