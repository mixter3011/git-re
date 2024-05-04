package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Git add
	cmd := exec.Command("git", "add", ".")
	if err := cmd.Run(); err != nil {
		fmt.Println("Error executing git add:", err)
		return
	}

	// Get commit message from user
	commitMsg, err := getCommitMessage()
	if err != nil {
		fmt.Println("Error getting commit message:", err)
		return
	}

	// Git commit
	cmd = exec.Command("git", "commit", "-m", commitMsg)
	if err := cmd.Run(); err != nil {
		fmt.Println("Error executing git commit:", err)
		return
	}

	// Git push
	cmd = exec.Command("git", "push")
	if err := cmd.Run(); err != nil {
		fmt.Println("Error executing git push:", err)
		return
	}

	fmt.Println("Successfully added, committed, and pushed changes!")
}

func getCommitMessage() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter commit message: ")
	commitMsg, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	commitMsg = strings.TrimSpace(commitMsg)
	return commitMsg, nil
}
