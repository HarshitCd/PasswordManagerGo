package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type UserData struct {
	Password      string `yaml:"password"`
	PrevPassword  string `yaml:"prev_password"`
	Title         string `yaml:"title"`
	Username      string `yaml:"username"`
}

func GetData() (map[string]UserData, error) {
	yamlData, err := ioutil.ReadFile("/Users/harshitcd/.config/passswdmanagergo/password_details.yml")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return nil, err
	}

	data := make(map[string]UserData)
	err = yaml.Unmarshal(yamlData, &data)
	if err != nil {
		fmt.Printf("Error unmarshaling YAML: %v\n", err)
		return nil, err
	}

	return data, nil
}

func StoreData(data map[string]UserData) error {

	yamlData, err := yaml.Marshal(data)
	if err != nil {
		return err
	}

	// Write the YAML data to a file
	err = ioutil.WriteFile("/Users/harshitcd/.config/passswdmanagergo/password_details.yml", yamlData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func GetDataJson() (map[string]UserData, error) {
	jsonData, err := ioutil.ReadFile("/Users/harshitcd/.config/passswdmanagergo/password_details.yml")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return nil, err
	}

	data := make(map[string]UserData)
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Printf("Error unmarshaling JSON: %v\n", err)
		return make(map[string]UserData), nil
	}

	return data, nil
}

func StoreDataJson(data map[string]UserData) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Write the JSON data to a file
	err = ioutil.WriteFile("/Users/harshitcd/.config/passswdmanagergo/password_details.yml", jsonData, 0644)
	if err != nil {
		return err
	}

	fmt.Println("JSON data successfully written to output.json")
	return nil
}