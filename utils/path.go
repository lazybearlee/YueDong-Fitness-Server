package utils

import (
	"errors"
	"os"
	"path/filepath"
)

func DirectoryExists(path string) bool {
	// 检查INode是否存在
	inode, err := os.Stat(path)
	if err == nil {
		// 检查是否是目录
		if inode.IsDir() {
			return true
		}
		return false
	}
	return false
}

func FileExists(path string) bool {
	// 检查INode是否存在
	inode, err := os.Stat(path)
	if err == nil {
		// 检查是否是文件
		if !inode.IsDir() {
			return true
		}
		return false
	}
	return false
}

func CreateDir(dirs ...string) error {
	for _, v := range dirs {
		// 检查目录是否存在
		if !DirectoryExists(v) {
			// 创建目录
			err := os.MkdirAll(v, os.ModePerm)
			if err != nil {
				return err
			}
		} else {
			return errors.New("存在同名文件")
		}
	}
	return nil
}

func MoveFile(src string, dst string) error {
	if dst == "" {
		return nil
	}
	src, err := filepath.Abs(src)
	if err != nil {
		return err
	}
	dst, err = filepath.Abs(dst)
	if err != nil {
		return err
	}
	// 尝试创建目录
	err = CreateDir(filepath.Dir(dst))
	if err != nil {
		return err
	}
	return os.Rename(src, dst)
}

func DeleteFile(path string) error {
	return os.RemoveAll(path)
}
