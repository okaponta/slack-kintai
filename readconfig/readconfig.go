package readconfig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Token    string
	Channels []string
	Comment  string
}

func ReadConfig(filename string) (Config, error) {
	var c Config
	jsonString, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ERROR", err)
		return c, err
	}
	err = json.Unmarshal(jsonString, &c)
	if err != nil {
		fmt.Println("ERROR", err)
		return c, err
	}
	return c, nil
}
