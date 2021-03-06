// +build windows

package execdrivers

import (
	"fmt"

	"github.com/docker/docker/daemon/execdriver"
	"github.com/docker/docker/daemon/execdriver/windows"
	"github.com/docker/docker/pkg/sysinfo"
)

func NewDriver(name, root, libPath, initPath string, sysInfo *sysinfo.SysInfo) (execdriver.Driver, error) {
	switch name {
	case "windows":
		return windows.NewDriver(root, initPath)
	}
	return nil, fmt.Errorf("unknown exec driver %s", name)
}
