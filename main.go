package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/charmbracelet/huh"
)
var (
	newSessionOption = huh.NewOption("[start a new session]", "new")
)

func main() {
	var options []huh.Option[string]
	for _, session := range getCurrentSessions() {
		options = append(options, huh.NewOption(session, session))
	}
	options = append(options, newSessionOption)

	var chosen string
	form := huh.NewSelect[string]().
		Title("Choose session").
		Options(options...).
		Value(&chosen)
	
	if err := form.Run(); err != nil {
		panic(err)
	}

	if chosen != newSessionOption.Value {
		attachToSession(chosen)
	} else {
		var sessionName string
		huh.NewInput().Title("Provide name for new session").Value(&sessionName).Run()
		if len(sessionName) == 0 {
			os.Exit(1)
		}
		createNewSession(sessionName)
	}
}

func getCurrentSessions() []string {
	cmd := exec.Command("tmux", "ls")
	stdout, err := cmd.Output()

	if err != nil {
		panic(err)
	}
	sessions := make([]string, 0)
	for _, l := range strings.Split(string(stdout), "\n") {
		if len(l) == 0 {
			continue
		}
		names := strings.SplitN(l, ":", 2)
		sessions = append(sessions, names[0])
	}
	return sessions
}

func attachToSession(sessionName string) {
	cmd := exec.Command("tmux", "at", fmt.Sprintf("-t %v", sessionName))
	fmt.Println(cmd.String())
}

func createNewSession(sessionName string) {
	cmd := exec.Command("tmux", "new", fmt.Sprintf("-s %v", sessionName))
	fmt.Println(cmd.String())
}
