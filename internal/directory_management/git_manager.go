package directorymanagement

import (
	"fmt"
	"os"
	"os/exec"
)

/*
* Start a repo, add remote url and push all the files inside the dir passed for parameter
@pathDir -string- the path of the directory to push changes
*/
func PushChanges(pathDir string, repoUrl string) error {
	if err := os.Chdir(pathDir); err != nil {
		return err
	}

	if err := start_repository(repoUrl); err != nil {
		return err
	}
	if err := stageAndPushChanges(); err != nil {
		return err
	}

	return nil
}

func stageAndPushChanges() error {
	c := exec.Command("git", "add", ".")

	if _, err := c.CombinedOutput(); err != nil {
		return err
	}

	c = exec.Command("git", "commit", "-m", "changes")
	if _, err := c.CombinedOutput(); err != nil {
		return err
	}

	c = exec.Command("git", "push", "-u", "--force", "origin", "main")
	if std, err := c.CombinedOutput(); err != nil {
		fmt.Println(string(std))
		return err
	}

	return nil
}
func start_repository(repoUrl string) error {
	c := exec.Command("git", "init")
	if _, err := c.CombinedOutput(); err != nil {
		return err
	}
	c = exec.Command("git", "branch", "-M", "main")
	if _, err := c.CombinedOutput(); err != nil {
		return err
	}
	c = exec.Command("git", "remote", "add", "origin", repoUrl)
	if _, err := c.CombinedOutput(); err != nil {
		return err
	}

	return nil
}
