package fs

type FileSystem struct {
	root  *Inode      // 根目录 inode
	super *SuperBlock // 超级块

	blocks     BlockManager // 块管理器
	inodes     InodeManager // inode 管理器
	inodetable InodeTable   // inode 表
	users      map[string]User
	// 其他全局资源
}

func (fs *FileSystem) CreateFile(path, owner, group string, mode mode) error
func (fs *FileSystem) ReadFile(path string, offset, size int) ([]byte, error)
func (fs *FileSystem) ListDir(path string) ([]string, error)
func (fs *FileSystem) DeleteFile(path string) error
