//go:build (linux || aix || zos || sylixos) && !js && !wasi
// +build linux aix zos sylixos
// +build !js
// +build !wasi

package logrus

import "golang.org/x/sys/unix"

const ioctlReadTermios = unix.TCGETS

func isTerminal(fd int) bool {
	_, err := unix.IoctlGetTermios(fd, ioctlReadTermios)
	return err == nil
}
