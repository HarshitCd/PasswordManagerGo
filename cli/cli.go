package main

import (
	"flag"
	"fmt"
	"strings"

	localapi "example.com/api"
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

	operation := flag.String("op", "GET", "specify your operation")
	option := flag.String("o", "all", "specify your option")
	title := flag.String("t", "", "enter a title")
	username := flag.String("u", "", "enter a unsername")
	password := flag.String("p", "", "enter a password")
	passwordLength := flag.Int("l", 8, "enter your desired password length")

	flag.Parse()
	
	message := ""
	switch strings.ToUpper(*operation) {
		case "GET":
			message = localapi.Get(*option)
		case "POST":
			message = localapi.Post(*title, *username, *password, *passwordLength)
		case "PUT":
			message = localapi.Put(*option, *title, *username, *password, *passwordLength)
		case "DELETE":
			message = localapi.Delete(*option)
		default:
			message = fmt.Sprintf("%v isn't a valid option", strings.ToUpper(*operation))
	}

	fmt.Println(message)
}
	