package localapi

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	utils "example.com/utils"
)

var data map[string]utils.UserData = make(map[string]utils.UserData)

func loadData() {
	var err error
	data, err = utils.GetData()
	if err != nil {
		log.Fatal(err)
	}
}

func storeData() {
	err := utils.StoreData(data)
	if err != nil {
		fmt.Println("Error occured when storing the data in local storage")
	}
}

func Post(title, username, password string, passwordLength int) string {

	loadData()

	option, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}

	option = []byte(strings.Trim(string(option), "\n"))

	if title == "" {
		return "-t cannot be empty"
	} 

	if username == "" {
		return "-u cannot be empty"
	} 

	if password == "" && passwordLength < 8 {
		return "-p cannot be empty or -l cannot be less than 8"
	}

	if passwordLength >= 8 {
		password = utils.GetPassword(passwordLength)
	}


	userData := utils.UserData {
		Title: title,
		Username: username,
		PrevPassword: data[string(option)].Password,
		Password: password,
	}

	data[string(option)] = userData

	storeData()
	return Get("all")
}

func Get(option string) string {

	loadData()

	if(strings.ToUpper(option) == "ALL") {
		return utils.DisplayAll(data)
	}

	return utils.DisplayRequired(data, option)
}

func Put(option, title, username, password string, passwordLength int) string {

	loadData()

	if _, ok := data[option]; !ok  {
		return "id doesn't exist"
	}

	if title == "" {
		title = data[option].Title
	} 

	if username == "" {
		username = data[option].Username
	} 

	if password == "" && passwordLength < 8 {
		return "-p cannot be empty or -l cannot be less than 8"
	}

	if passwordLength >= 8 {
		password = utils.GetPassword(passwordLength)
	}

	userData := utils.UserData {
		Title: title,
		Username: username,
		PrevPassword: data[string(option)].Password,
		Password: password,
	}

	data[string(option)] = userData

	storeData()
	return Get("all")
}

func Delete(option string) string {

	loadData()

	delete(data, option)

	storeData()
	return Get("all")
}