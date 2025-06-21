type Inode struct {
	ID        uint64
	FileType  FileType // 枚举类型，普通文件/目录/符号链接等
	OwnerUID  uint32
	OwnerGID  uint32
	Permission uint16  // 例如 0644
	Size      int64
	Ctime     time.Time
	Mtime     time.Time
	Atime     time.Time
	Links     uint32   // 硬链接数
	Blocks    []uint64 // 指向数据块的列表
}

enum FileType {
	
}