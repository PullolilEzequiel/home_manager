package savecommand

import (
	"fmt"

	configmanager "github.com/PullolilEzequiel/wizard-home/internal/config_manager"
	directory_management "github.com/PullolilEzequiel/wizard-home/internal/directory_management"
)

type saveManager struct {
	config configmanager.Config
}

func SaveManager() saveManager {
	config := configmanager.GetConfig()
	return saveManager{
		config: config,
	}
}

func (sm saveManager) SaveConfigState() error {
	return sm.config.CreateTemporalFolder("save_folder", sm.copyConfigFilesAndSaveInRemoteRepository)
}

func (sm saveManager) copyConfigFilesAndSaveInRemoteRepository(folderName string) error {
	if err := sm.copyConfigFilesAndFolder(folderName); err != nil {
		return err
	}
	if err := sm.pushConfigStateToRepository(folderName); err != nil {
		return err
	}
	return nil
}

func (sm saveManager) copyConfigFilesAndFolder(folderName string) error {
	fmt.Println("Saving folders and files ")
	for _, fileOrFolder := range sm.config.ConfigPaths() {
		fmt.Println("Saving " + fileOrFolder)
		if err := directory_management.CopyFolderOrFile(fileOrFolder, folderName); err != nil {
			return err
		}
	}

	return sm.copyConfigFile(folderName)
}

func (sm saveManager) copyConfigFile(folderName string) error {
	return directory_management.CopyFolderOrFile(sm.config.ConfigFilePath(), folderName)
}

func (sm saveManager) pushConfigStateToRepository(folderName string) error {
	fmt.Println(folderName)

	return directory_management.PushChanges(folderName, sm.config.RepoUrl())
}
