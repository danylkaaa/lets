package config

import (
	"fmt"

	"github.com/lets-cli/lets/commands/command"
	"github.com/lets-cli/lets/util"
)

const (
	NoticeColor = "\033[1;36m%s\033[0m"
)

// Validate loaded config
func Validate(config *Config) error {
	return validateCircularDepends(config)
}

// if any two commands have each other command in deps, raise error
// TODO not optimal function but works for now
func validateCircularDepends(cfg *Config) error {
	for _, cmdA := range cfg.Commands {
		for _, cmdB := range cfg.Commands {
			if cmdA.Name == cmdB.Name {
				continue
			}

			if yes := depsIntersect(cmdA, cmdB); yes {
				return fmt.Errorf(
					"command '%s' have circular depends on command '%s'",
					fmt.Sprintf(NoticeColor, cmdA.Name),
					fmt.Sprintf(NoticeColor, cmdB.Name),
				)
			}
		}
	}

	return nil
}

func depsIntersect(cmdA command.Command, cmdB command.Command) bool {
	return util.IsStringInList(cmdA.Name, cmdB.Depends) && util.IsStringInList(cmdB.Name, cmdA.Depends)
}

func validateTopLevelFields(rawKeyValue map[string]interface{}, validFields []string) error {
	for key := range rawKeyValue {
		if !util.IsStringInList(key, validFields) {
			return fmt.Errorf("unknown top-level field '%s'", key)
		}
	}

	return nil
}
