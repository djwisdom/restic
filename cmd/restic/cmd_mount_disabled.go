//go:build !darwin && !freebsd && !linux
// +build !darwin,!freebsd,!linux

package main

import "github.com/spf13/cobra"

func registerMountCommand(_ *cobra.Command) {
	// Mount command not supported on these platforms
}
