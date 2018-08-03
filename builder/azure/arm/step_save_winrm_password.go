package arm

import (
	"context"

	"github.com/hashicorp/packer/helper/multistep"
)

type StepSaveWinRMPassword struct {
	Password  string
	BuildName string
}

func (s *StepSaveWinRMPassword) Run(_ context.Context, state multistep.StateBag) multistep.StepAction {
	// store so that we can access this later during provisioning
	state.Put("winrm_password", s.Password)
	return multistep.ActionContinue
}

func (s *StepSaveWinRMPassword) Cleanup(multistep.StateBag) {
}
