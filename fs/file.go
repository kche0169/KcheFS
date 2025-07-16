package fs

import "sync"

type FileNode struct {
	Inode  *Inode
	Offset int64      // 当前偏移（适用于读写）
	mu     sync.Mutex // 并发安全
	// 可以扩展缓存、文件句柄、状态等
	// Cache   *FileCache
	// Flags   int
	// Extra   map[string]interface{}
}
