package os

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func PathExists(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	if !errors.Is(err, fs.ErrNotExist) {
		return fmt.Errorf("path exists check failed: %w", err)
	}
	return fmt.Errorf("path does not exist: %s", path)

}
