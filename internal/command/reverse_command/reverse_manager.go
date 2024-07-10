package reversecommand

import (
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

	return nil
}
