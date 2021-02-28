package smokesignal

import (
	"context"
	"fmt"
	"os/exec"
	"testing"
	"time"
)

type PortStatus string

var (
	used    PortStatus = "used"
	notUsed PortStatus = "NOT used"
)

func CheckIfPortIsInUse(port int, timeout time.Duration, t *testing.T) (PortStatus, error) {
	t.Helper()

	timeoutContext, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	return checkIfPortIsInUseWithTimeout(port, timeoutContext)
}

func checkIfPortIsInUseWithTimeout(port int, timeoutContext context.Context) (PortStatus, error) {
	for {
		select {
		case <-timeoutContext.Done():
			return runLsofForPort(port)
		default:
			if useState, err := runLsofForPort(port); useState == used {
				return useState, err
			}
		}
		time.Sleep(1 * time.Second)
	}
}

func runLsofForPort(port int) (PortStatus, error) {
	cmd := exec.Command( "lsof", fmt.Sprintf("-i:%d", port))
	if err := cmd.Run(); err != nil {
		return checkIfLsofExitedWithNonzeroCode(err)
	}

	return used, nil
}

func checkIfLsofExitedWithNonzeroCode(err error) (PortStatus, error) {
	if _, ok := err.(*exec.ExitError); ok {
		return notUsed, nil
	}

	return "", err
}
