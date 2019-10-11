package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type stdMsg struct {
	Message string `json:"message"`
}

func ReadStdMsgFromFile(filename string) (message string, err error) {
	var bytes []byte
	fmt.Println("filename is" + filename)
	if filename == "-" {
		bytes, err = ioutil.ReadAll(os.Stdin)
	} else {
		bytes, err = ioutil.ReadFile(filename)
	}
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	stdMessage := stdMsg{}

	err = json.Unmarshal(bytes, &stdMessage)

	return stdMessage.Message, err
}

