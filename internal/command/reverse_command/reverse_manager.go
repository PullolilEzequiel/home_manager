package reversecommand

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	configmanager "github.com/PullolilEzequiel/wizard-home/internal/config_manager"
)

type reverseManager struct {
	config configmanager.Config
}

func ReverseManager() reverseManager {
	config := configmanager.GetConfig()

	return reverseManager{
		config: config,
	}
}

func (rm reverseManager) ReverseConfigState() error {
	return rm.config.CreateTemporalFolder("reverse_folder", rm.replaceSystemFilesForRemote)
}

/*
Replace the system files asociated in the config for those that exist in the remote repository

@folderPath string : the folder path to pull al remote files
*/
func (rm reverseManager) replaceSystemFilesForRemote(folderPath string) error {
	if err := rm.cloneRemoteFiles(folderPath); err != nil {
		return err
	}
	if err := rm.replaceFilesAsociatedInConfig(folderPath); err != nil {
		return err
	}
	return nil
}

func (rm reverseManager) cloneRemoteFiles(temporalPath string) error {
	fmt.Println("Cloning repository")
	os.Chdir(temporalPath)
	cmd := exec.Command("git", "clone", rm.config.RepoUrl())

	if _, err := cmd.CombinedOutput(); err != nil {
		return err
	}

	return os.Chdir(rm.config.HomeDir())
}

func (rm reverseManager) replaceFilesAsociatedInConfig(temporalPath string) error {
	dir := temporalPath + rm.config.RepoName()

	for _, fileDir := range rm.config.ConfigPaths() {
		d := path.Base(fileDir)
		fmt.Println(fileDir)
		fmt.Println(dir + d)
	}
	return nil
}
