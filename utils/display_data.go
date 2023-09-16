package utils

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
)


func DisplayAll(data map[string]UserData) string {

    t := table.NewWriter()

	t.AppendHeader(table.Row{"Option", "Title", "Username"})
	for option := range data {
		t.AppendRow(table.Row{option, data[option].Title, data[option].Username})
	}

	return fmt.Sprintf(t.Render())
}

func DisplayAllWithPasswd(data map[string]UserData) string {

    t := table.NewWriter()

	t.AppendHeader(table.Row{"Option", "Title", "Username", "Password"})
	for option := range data {
		t.AppendRow(table.Row{option, data[option].Title, data[option].Username, data[option].Password})
	}

	return fmt.Sprintf(t.Render())
}

func DisplayRequired(data map[string]UserData, option string) string {
	
	if _, ok := data[option]; !ok {
		return "The id does not exist"
	}

	t := table.NewWriter()

	t.AppendHeader(table.Row{"Option", "Title", "Username", "Password"})
	t.AppendRow(table.Row{option, data[option].Title, data[option].Username, data[option].Password})

	return fmt.Sprintf(t.Render())
}