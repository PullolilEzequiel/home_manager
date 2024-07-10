package reversecommand

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	configmanager "github.com/PullolilEzequiel/wizard-home/internal/config_manager"
	directory_management "github.com/PullolilEzequiel/wizard-home/internal/directory_management"
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
	if err := rm.replaceFilesAsociatedInConfig(path.Join(folderPath, rm.config.RepoName())); err != nil {
		return err
	}
	return nil
}

/*
Clones the remote repository of the wizard_home config

@directoryPathDestiny string : the folder to clone the repository
*/
func (rm reverseManager) cloneRemoteFiles(directoryPathDestiny string) error {
	fmt.Println("Cloning repository")
	os.Chdir(directoryPathDestiny)
	cmd := exec.Command("git", "clone", rm.config.RepoUrl())

	if _, err := cmd.CombinedOutput(); err != nil {
		return err
	}

	return os.Chdir(rm.config.HomeDir())
}

/*
Replace the files asociated in the wizard_home config.json with the files of the location path passed for parameter

@pathLocation string: The directory where the files and folders are located
*/
func (rm reverseManager) replaceFilesAsociatedInConfig(pathLocation string) error {
	for _, fileOrFolder := range rm.config.ConfigPaths() {
		folderName := path.Base(fileOrFolder)
		newFileOrDirectory := path.Join(pathLocation, folderName)
		if err := directory_management.ReplaceFileOrFolderFor(fileOrFolder, newFileOrDirectory); err != nil {
			return err
		}
	}
	return nil
}
