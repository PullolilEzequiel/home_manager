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


	return nil
}
