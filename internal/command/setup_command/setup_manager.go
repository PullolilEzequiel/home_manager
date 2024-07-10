package setupcommand

import configmanager "github.com/PullolilEzequiel/wizard-home/internal/config_manager"

type setupManager struct {
	config configmanager.Config
}

func SetupManager() setupManager {
	config := configmanager.GetConfig()

	return setupManager{
		config: config,
	}
}

func (sm setupManager) SetupConfigState() error {
	return nil
}
