package fs

import "errors"

// 通用错误定义
var (
	ErrFileNotFound     = errors.New("File not found")
	ErrPermissionDenied = errors.New("Permission denied")
	ErrInvalidPath      = errors.New("Invalid path")
)

// 自定义错误类型（如需更复杂的错误信息）
type FileError struct {
	Path string
	Msg  string
}

func (e *FileError) Error() string {
	return e.Msg + ": " + e.Path
}
