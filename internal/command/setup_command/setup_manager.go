package setupcommand

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	configmanager "github.com/PullolilEzequiel/wizard-home/internal/config_manager"
	directory_management "github.com/PullolilEzequiel/wizard-home/internal/directory_management"
)

type setupManager struct {
	externalRepo string
	config       configmanager.Config
}

func SetupManager(urlConfig string) setupManager {
	config := configmanager.GetConfig()
	return setupManager{
		externalRepo: urlConfig,
		config:       config,
	}
}

/*
Replace the actual config state for the other user of wizard_home
*/
func (sm setupManager) SetupConfigState() error {
	return sm.config.CreateTemporalFolder("setup_folder", sm.replaceSystemFilesForRemote)
}

func (sm setupManager) replaceSystemFilesForRemote(folderPath string) error {
	if err := sm.cloneRemoteFiles(folderPath); err != nil {
		return err
	}
	if err := sm.replaceFilesAsociatedInConfig(path.Join(folderPath, sm.config.RepoName())); err != nil {
		return err
	}
	return nil
}

func (sm setupManager) cloneRemoteFiles(directoryPathDestiny string) error {
	os.Chdir(directoryPathDestiny)
	cmd := exec.Command("git", "clone", sm.externalRepo)
	if _, err := cmd.CombinedOutput(); err != nil {
		return err
	}
	configToCopy := path.Join(directoryPathDestiny, path.Base(sm.externalRepo), "config.json")
	if err := directory_management.ReplaceFile(sm.config.ConfigFilePath(), configToCopy); err != nil {
		return err
	}

	return os.Chdir(sm.config.HomeDir())
}
func (sm setupManager) replaceFilesAsociatedInConfig(pathLocation string) error {
	for _, fileOrFolder := range sm.config.ConfigPaths() {
		fmt.Printf("- Copy config \"%s\" from %s \n", path.Base(fileOrFolder), sm.externalRepo)
		folderName := path.Base(fileOrFolder)
		newFileOrDirectory := path.Join(pathLocation, folderName)
		if err := directory_management.ReplaceFileOrFolderFor(fileOrFolder, newFileOrDirectory); err != nil {
			return err
		}
	}
	return nil
}
