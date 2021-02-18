package smokesignal

import (
	"context"
	"os/exec"
	"testing"
	"time"
)

type used string

var (
	localhost string = "localhost"

	USED    used = "USED"
	NOTUSED used = "NOT USED"
)

func CheckIfPortIsInUse(port string, timeout time.Duration, t *testing.T) (used, error) {
	t.Helper()

	timeoutContext, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	return checkIfPortIsInUseWithTimeout(port, timeoutContext)
}

func checkIfPortIsInUseWithTimeout(port string, timeoutContext context.Context) (used, error) {
	for {
		select {
		case <-timeoutContext.Done():
			return runLsofForPort(port)
		default:
			if useState, err := runLsofForPort(port); useState == USED {
				return useState, err
			}
		}
		time.Sleep(1 * time.Second)
	}
}

func runLsofForPort(port string) (used, error) {
	cmd := exec.Command("lsof", "-i:"+port)
	if err := cmd.Run(); err != nil {
		return checkIfLsofExitedWithNonzeroCode(err)
	}

	return USED, nil
}

func checkIfLsofExitedWithNonzeroCode(err error) (used, error) {
	if _, ok := err.(*exec.ExitError); ok {
		return NOTUSED, nil
	}

	return "", err
}
