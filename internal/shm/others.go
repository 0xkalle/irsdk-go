// +build !windows

package shm

import (
	ipcShm "bitbucket.org/avd/go-ipc/shm"
	"fmt"
)

func NewWindowsNativeMemoryObject(_ string, _ int, _ int) (*ipcShm.MemoryObject, error) {
	return nil, fmt.Errorf("only implemented on windows")
}
