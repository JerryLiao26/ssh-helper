package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

type configStruct struct {
	PID  int    `json:"pid"`
	Sock string `json:"sock"`
}

var (
	// Config items
	PID  = 0
	Sock = ""

	// Paths
	dirPath  = getFullPath("~/.ssh-helper")
	fullPath = getFullPath("~/.ssh-helper/config.json")
)

func LoadConf() {
	// Check config exists
	_, err := os.Stat(dirPath)
	_, fileErr := os.Stat(fullPath)
	if err != nil && os.IsNotExist(err) {
		_ = os.Mkdir(dirPath, 0755)
	} else if fileErr != nil && os.IsNotExist(fileErr) {
		// Create with default
		SaveConf()
	} else {
		b, err := ioutil.ReadFile(fullPath)

		if err != nil {
			fmt.Println(err)
			fmt.Println("Error: Cannot read config file, try manually delete it at \"~/.ssh-helper/config.json\"")
		} else {
			var conf configStruct
			err = json.Unmarshal(b, &conf)

			if err != nil {
				fmt.Println("Error: Cannot parse config file, try manually delete it at \"~/.ssh-helper/config.json\"")
			} else {
				PID = conf.PID
				Sock = conf.Sock

				// Set env
				setEnv()
			}
		}
	}
}

func SaveConf() {
	conf := configStruct{
		PID:  PID,
		Sock: Sock,
	}
	s, _ := json.Marshal(conf)

	err := ioutil.WriteFile(fullPath, s, 0755)
	if err != nil {
		fmt.Println("Error: Cannot write config file, try manually create it at ~/.ssh-helper/config.json")
	} else {
		setEnv()
	}
}

func ValidateConf() {
	flag := checkAgentValid()

	if !flag {
		PID = 0
		Sock = ""

		// Save default
		SaveConf()
	}
}

func checkAgentValid() bool {
	// Check socket exists
	_, err := os.Stat(Sock)
	if err != nil && os.IsNotExist(err) {
		return false
	}

	// Check PID exists
	_, index := ListAgentProcess()

	if index == -1 {
		return false
	}

	// Pass
	return true
}

func setEnv() {
	sockValue := Sock
	pidValue := strconv.Itoa(PID)
	if pidValue == "0" {
		pidValue = ""
	}

	_ = os.Setenv("SSH_AUTH_SOCK", sockValue)
	_ = os.Setenv("SSH_AGENT_PID", pidValue)
}

func getFullPath(path string) string {
	if path[0:1] == "~" {
		path = os.Getenv("HOME") + path[1:]
	}

	fullPath, err := filepath.Abs(path)

	if err != nil {
		fmt.Println("error:", err)
	}

	return fullPath
}
