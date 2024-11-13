//go:build linux

package main

func GetKubeCTLInstallCommand() (string, []string) {
	cmd := "curl"
	args := []string{"-LO", "https://dl.k8s.io/release/v1.31.2/bin/linux/amd64/kubectl"}
	return cmd, args
}
