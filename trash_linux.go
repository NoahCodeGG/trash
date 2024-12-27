//go:build linux
// +build linux

package trash

import (
	"fmt"
	"os/exec"
)

// check command exists
func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

func MoveToTrash(filePath string) error {
	if commandExists("gio") {
		if err := exec.Command("gio", "trash", filePath).Run(); err == nil {
			return nil
		}
	}

	if commandExists("kioclient5") {
		if err := exec.Command("kioclient5", "move", filePath, "trash:/").Run(); err == nil {
			return nil
		}
	}

	if commandExists("xdg-trash") {
		if err := exec.Command("xdg-trash", filePath).Run(); err == nil {
			return nil
		}
	}

	if commandExists("trash-put") {
		if err := exec.Command("trash-put", filePath).Run(); err == nil {
			return nil
		}
	}

	return fmt.Errorf("failed to move file to trash: no supported method found")
}
