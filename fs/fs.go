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

func InitFileSystem(fs *FileSystem, blockSize int, inodeSize int) error {
	// 初始化超级块
	fs.super = &SuperBlock{
		BlockSize: blockSize,
		InodeSize: inodeSize,
	}

	// 初始化块管理器
	fs.blocks = NewBlockManager(blockSize)

	// 初始化 inode 管理器
	fs.inodes = NewInodeManager(inodeSize)

	// 初始化 inode 表
	fs.inodetable = NewInodeTable()

	// 创建根目录 inode
	rootInode := fs.inodes.AllocInode()
	rootInode.Mode = Dir
	rootInode.Owner = "root"
	rootInode.Group = "root"
	fs.root = rootInode

	return nil

}

func (fs *FileSystem) CreateFile(path, owner, group string, mode mode) error
func (fs *FileSystem) ReadFile(path string, offset, size int) ([]byte, error)
func (fs *FileSystem) ListDir(path string) ([]string, error)
func (fs *FileSystem) DeleteFile(path string) error
