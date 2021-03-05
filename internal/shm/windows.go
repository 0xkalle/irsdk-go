// +build windows

package shm

import ipcShm "bitbucket.org/avd/go-ipc/shm"

func NewWindowsNativeMemoryObject(name string, flag int, size int) (*ipcShm.WindowsNativeMemoryObject, error) {
	return ipcShm.NewWindowsNativeMemoryObject(name, flag, size)
}
