// +build !windows

package main

import (
	"context"
)

const defaultConfigPath = "/etc/litestream.yml"

func isWindowsService() (bool, error) {
	return false, nil
}

func runWindowsService(ctx context.Context) error {
	panic("cannot run windows service as unix process")
}