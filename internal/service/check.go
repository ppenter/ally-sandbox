package service

import (
	"errors"

	"github.com/ppenter/ally-sandbox/internal/core/runner/types"
	"github.com/ppenter/ally-sandbox/internal/static"
)

var (
	ErrNetworkDisabled = errors.New("network is disabled, please enable it in the configuration")
)

func checkOptions(options *types.RunnerOptions) error {
	configuration := static.GetDifySandboxGlobalConfigurations()

	if options.EnableNetwork && !configuration.EnableNetwork {
		return ErrNetworkDisabled
	}

	return nil
}
