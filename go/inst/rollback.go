package inst

// Main use case:
// 1. User moves a group of slaves
// 2. This hangs because of an unrelated issue
// 3. User presses ^C to sent SIGINT
// 4. Orchestrator tries to cleanup: e.g. start slaves.

import (
	"github.com/openark/golib/log"
)

type rollbackAction struct {
	action   string
	argument string
}

var rollbackActions []rollbackAction

func AddRollbackAction(action string, argument string) {
	act := rollbackAction{action, argument}
	rollbackActions = append(rollbackActions, act)
}

func Rollback() error {
	if len(rollbackActions) == 0 {
		return nil
	}
	log.Infof("Starting rollback")
	for i, ra := range rollbackActions {
		log.Infof("RollbackAction %d/%d (type: %s)", i+1, len(rollbackActions), ra.action)
		switch ra.action {
		case "print_message":
			log.Infof("%s", ra.argument)
		case "start_slave":
			instanceKey, err := NewRawInstanceKey(ra.argument)
			if err == nil {
				log.Infof("Rollback: Starting replica %s", ra.argument)
				StartSlave(instanceKey)
			} else {
				log.Errorf("Failed to get instanceKey for %s", ra.argument)
			}
		default:
			return log.Errorf("Unknown rollback action: %s", ra.action)
		}
	}
	return nil
}
