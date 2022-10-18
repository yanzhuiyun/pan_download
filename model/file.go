package model

type File struct {
	Hash   string //文件hash值
	Path   string //文件路径
	Status int    //文件状态
}

// UserFile 用户文件表
type UserFile struct {
	Username string //用户名
	Filename string //文件名
	Filehash string //文件hash
}
