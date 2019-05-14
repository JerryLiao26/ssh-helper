package cli

import (
	"fmt"
	"github.com/JerryLiao26/ssh-helper/helper"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func EnvHandler() {
	sock := os.Getenv("SSH_AUTH_SOCK")
	pid := os.Getenv("SSH_AGENT_PID")

	// Print
	fmt.Print("SSH_AUTH_SOCK=")
	fmt.Println(sock)
	fmt.Print("SSH_AGENT_PID=")
	fmt.Println(pid)

	// Print check result
	agentGroup, index := helper.ListAgentProcess()
	fmt.Print("Total " + strconv.Itoa(len(agentGroup)) + " ssh-agent running, ")
	if index != -1 {
		fmt.Println("1 managed.")
	} else {
		fmt.Println("none is managed.")
	}
}

func KillHandler() {
	agentGroup, _ := helper.ListAgentProcess()

	for _, agent := range agentGroup {
		cmd := exec.Command("kill", strconv.Itoa(agent.PID))
		_ = cmd.Run()
	}

	// Reset config
	helper.PID = 0
	helper.Sock = ""

	helper.SaveConf()
}

func TidyHandler() {
	agentGroup, _ := helper.ListAgentProcess()

	for _, agent := range agentGroup {
		if agent.PID != helper.PID {
			cmd := exec.Command("kill", strconv.Itoa(agent.PID))
			_ = cmd.Run()
		}
	}
}

func StartHandler() {
	if helper.PID == 0 && helper.Sock == "" {
		cmd := exec.Command("ssh-agent")
		output, err := cmd.CombinedOutput()

		if err != nil {
			fmt.Println("Error: Your system does not support \"ssh-agent\" command")
		} else {
			stream := strings.Replace(string(output), "\n", "", -1)
			commands := strings.Split(stream, ";")
			for _, c := range commands[:(len(commands) - 1)] {
				c = strings.TrimSpace(c)
				if strings.HasPrefix(c, "SSH_AUTH_SOCK=") {
					helper.Sock = c[14:]
				} else if strings.HasPrefix(c, "SSH_AGENT_PID=") {
					helper.PID, _ = strconv.Atoi(c[14:])
				}
			}

			// Save
			helper.SaveConf()
		}
	} else {
		fmt.Println("There's already a managed ssh-agent with PID " + strconv.Itoa(helper.PID) + " running")
	}
}

func AddHandler(command string) {
	if helper.PID != 0 && helper.Sock != "" {
		cmd := exec.Command("ssh-add", command)
		output, _ := cmd.CombinedOutput()

		// Output anyway
		fmt.Print("[ssh-add]", string(output))
	} else {
		fmt.Println("No managed ssh-agent is running. Run ssh-helper -s first")
	}
}
