// +build !windows

package shm

import (
	"fmt"

	ipcShm "bitbucket.org/avd/go-ipc/shm"
)

func NewWindowsNativeMemoryObject(_ string, _ int, _ int) (*ipcShm.MemoryObject, error) {
	return nil, fmt.Errorf("only implemented on windows")
}
