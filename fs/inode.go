package fs

import (
	"time"
)

// 文件类型枚举
type FileType int

const (
	File FileType = iota
	Directory
)

type mode uint32

const (
	ModeNone    mode = iota
	ModeExecute mode = 1 << iota
	ModeWrite
	ModeRead
)

// Inode 结构体定义文件/目录元数据
type Inode struct {
	ID         uint64    // 唯一标识
	Name       string    // 文件或目录名
	Type       FileType  // 文件 or 目录
	Size       uint64    // 文件大小（字节，目录为0或统计所有子项）
	Mode       mode      // 权限，如Unix风格
	Owner      string    // 所有者
	Group      string    // 所属组
	CreatedAt  time.Time // 创建时间
	UpdatedAt  time.Time // 最近修改时间
	AccessedAt time.Time // 最近访问时间

	// 数据块指针（文件用），或者子目录项列表（目录用）
	Blocks   []uint64          // 数据块号列表
	Children map[string]*Inode // 目录项名 -> Inode指针，目录专用
	Parent   *Inode            // 父目录，根目录为nil
}

// NewFileInode 创建文件类型的inode
func NewFileInode(id uint64, name, owner, group string, mode mode) *Inode {
	now := time.Now()
	return &Inode{
		ID:         id,
		Name:       name,
		Type:       File,
		Mode:       mode,
		Owner:      owner,
		Group:      group,
		CreatedAt:  now,
		UpdatedAt:  now,
		AccessedAt: now,
		Blocks:     []uint64{},
	}
}

// NewDirInode 创建目录类型的inode
func NewDirInode(id uint64, name, owner, group string, mode mode) *Inode {
	now := time.Now()
	return &Inode{
		ID:         id,
		Name:       name,
		Type:       Directory,
		Mode:       mode,
		Owner:      owner,
		Group:      group,
		CreatedAt:  now,
		UpdatedAt:  now,
		AccessedAt: now,
		Children:   make(map[string]*Inode),
	}
}

// IsDir 判断是否为目录
func (i *Inode) IsDir() bool {
	return i.Type == Directory
}

// IsFile 判断是否为文件
func (i *Inode) IsFile() bool {
	return i.Type == File
}

// CanRead 判断是否可读
func (i *Inode) CanRead() bool {
	return i.Mode&ModeRead != 0
}

// CanWrite 判断是否可写
func (i *Inode) CanWrite() bool {
	return i.Mode&ModeWrite != 0
}

// CanExecute 判断是否可执行
func (i *Inode) CanExecute() bool {
	return i.Mode&ModeExecute != 0
}
