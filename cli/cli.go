package main

import (
	"flag"
	"fmt"
	"strings"

	localapi "example.com/api"
	"example.com/utils"
)

func main() {

	// Functionalities
	// ---------------
	// ./cli GET all (or) ./cli
	// ./cli GET <option> (username / pass / prev_pass)
	// ./cli POST username {data} pass {data}
	// ./cli POST username {data} pass_len {data}
	// ./cli PUT username {data} pass {data}
	// ./cli PUT username {data} pass_len {data}
	// ./cli DELETE <option>

	operation := flag.String(utils.OPERATION_FLAG, utils.DEFAULT_OPERATION, "specify your operation")
	show := flag.Bool(utils.SHOW_FLAG, utils.DEFAULT_SHOW, "show the password in display")
	option := flag.String(utils.OPTION_FLAG, utils.DEFAULT_OPTION, "specify your password UUID")
	title := flag.String(utils.TITLE_FLAG, utils.DEFAULT_TITLE, "enter a new title")
	username := flag.String(utils.USERNAME_FLAG, utils.DEFAULT_USERNAME, "enter a new unsername")
	password := flag.String(utils.PASSWORD_FLAG, utils.DEFAULT_PASSWORD, "enter a new password")
	passwordLength := flag.Int(utils.PASSWORD_LENGTH_FLAG, utils.DEFAULT_PASSWORD_LENGTH, "enter your desired password length")

	flag.Parse()
	
	message := ""
	switch strings.ToUpper(*operation) {
		case utils.GET_OPERATION:
			message = localapi.Get(*option, *show)
		case utils.POST_OPERATION:
			message = localapi.Post(*title, *username, *password, *passwordLength)
		case utils.PUT_OPERATION:
			message = localapi.Put(*option, *title, *username, *password, *passwordLength)
		case utils.DELETE_OPERATION:
			message = localapi.Delete(*option)
		default:
			message = fmt.Sprintf("%v isn't a valid option", strings.ToUpper(*operation))
	}

	fmt.Println(message)
}
	