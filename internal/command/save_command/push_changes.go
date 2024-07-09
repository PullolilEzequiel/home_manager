package savecommand

import (
	"fmt"
	"os"
	"os/exec"
)

func PushChanges(dir string, repoUrl string) {
	os.Chdir(dir)
	init_repo()
	add_remote(repoUrl)
	stage_changes()
	push_changes()
}

func push_changes() {
	c := exec.Command("git", "push", "-u", "origin", "main")
	if std, err := c.CombinedOutput(); err != nil {
		fmt.Println(string(std))
		panic(err)
	}
}

func init_repo() {
	c := exec.Command("git", "init")
	if _, err := c.CombinedOutput(); err != nil {
		panic(err)
	}
	c = exec.Command("git", "branch", "-M", "main")
	if _, err := c.CombinedOutput(); err != nil {
		panic(err)
	}
}

func add_remote(url string) {
	c := exec.Command("git", "remote", "add", "origin", url)
	if _, err := c.CombinedOutput(); err != nil {
		panic(err)
	}
}

func stage_changes() {
	c := exec.Command("git", "add", ".")

	if _, err := c.CombinedOutput(); err != nil {
		panic(err)
	}

	c = exec.Command("git", "commit", "-m", "changes")
	if _, err := c.CombinedOutput(); err != nil {
		panic(err)
	}
}
