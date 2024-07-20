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

	for _, fileOrFolder := range sm.config.ConfigPaths() {
		fmt.Println("Copy")
		fmt.Printf("*%s to %s \n", fileOrFolder, folderName)
		if err := directory_management.CopyFolderOrFile(fileOrFolder, folderName); err != nil {
			return err
		}
	}
	return nil
}

func (sm saveManager) pushConfigStateToRepository(folderName string) error {
	fmt.Println(folderName)
	return nil
}
