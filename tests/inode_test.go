package fs_test

import (
	"github.com/kche0169/KcheFS/fs"
	"testing"
)

func TestNewFileInode(t *testing.T) {
	inode := fs.NewFileInode(1, "test.txt", "user", "group", fs.ModeRead|fs.ModeWrite)
	if inode.Name != "test.txt" || !inode.IsFile() {
		t.Errorf("File inode creation failed")
	}
	if !inode.CanRead() || !inode.CanWrite() || inode.CanExecute() {
		t.Errorf("File inode mode check failed")
	}
}

func TestNewDirInode(t *testing.T) {
	inode := fs.NewDirInode(2, "docs", "user", "group", fs.ModeRead|fs.ModeExecute)
	if inode.Name != "docs" || !inode.IsDir() {
		t.Errorf("Dir inode creation failed")
	}
	if !inode.CanRead() || !inode.CanExecute() || inode.CanWrite() {
		t.Errorf("Dir inode mode check failed")
	}
}
