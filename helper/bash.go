package helper

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

type processItem struct {
	PID     int
	Command string
}

func ListAgentProcess() ([]processItem, int) {
	var mark int
	var processes []processItem

	// Execute
	cmd := exec.Command("ps", "-ux")
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("Error: Your system does not support \"ps -aux\" command")
	} else {
		// Init mark
		mark = -1

		// Search user processes
		pGroup := strings.Split(string(output), "\n")
		for i, p := range pGroup[1:(len(pGroup) - 1)] {
			reg := regexp.MustCompile(`\x20+`)
			pInfoGroup := reg.Split(p, -1)
			name := strings.TrimSpace(pInfoGroup[10])
			if name == "ssh-agent" {
				pid, _ := strconv.Atoi(strings.TrimSpace(pInfoGroup[1]))

				item := processItem{
					PID: pid,
					Command: name,
				}

				// Append to ssh-agent processes
				processes = append(processes, item)

				// Process is managed
				if pid == PID {
					mark = i
				}
			}
		}
	}

	return processes, mark
}
