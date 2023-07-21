package utils

import "fmt"


func DisplayAll(data map[string]UserData) string {

	message := ""

	for option := range data {
		message += fmt.Sprintf("option:%v\ntitle: %s\nusername: %s\n---------------------------\n", option, data[option].Title, data[option].Username)
	}

	return message
}

func DisplayRequired(data map[string]UserData, option string) string {
	for key := range data {
		if key == option {
			return fmt.Sprintf("username: %v\npassword: %v\nprev password: %v", data[option].Username, 
										data[option].Password, data[option].PrevPassword)
		}
	}

	return ""
}