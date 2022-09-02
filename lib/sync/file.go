package sync

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
)

func checkNotExist(src string) bool {
	_, err := os.Stat(src)
	return os.IsNotExist(err)
}

func checkPermission(src string) bool {
	_, err := os.Stat(src)
	return os.IsPermission(err)
}

func isNotExistMkdir(src string) error {
	if notExist := checkNotExist(src); notExist {
		if err := mkDir(src); err != nil {
			return err
		}
	}
	return nil
}

func mkDir(src string) error {
	syscall.Umask(0)
	return os.MkdirAll(src, os.ModePerm)
}

func mustOpen(filename, dir string) (*os.File, error) {
	if checkPermission(dir) {
		return nil, fmt.Errorf("permission denied dir: %s", dir)
	}

	if err := isNotExistMkdir(dir); err != nil {
		return nil, fmt.Errorf("error during make dir: %s, err: %s", dir, err)
	}
	syscall.Umask(0)
	f, err := os.OpenFile(filepath.Join(dir, filename), os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("fail to open file, err: %s", err)
	}
	return f, nil
}
