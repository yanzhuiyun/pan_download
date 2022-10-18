package mysql

import (
	"fmt"
	"pandownload/model"
)

func SaveFile(file *model.File) error {
	sqlStr := fmt.Sprintf("INSERT INTO files(filehash,filepath) VALUES(?,?)")
	stmt, _ := db.Prepare(sqlStr)
	_, err := stmt.Exec(file.Hash, file.Path)
	return err
}

func GetPath(hash string) (path string) {
	sqlStr := fmt.Sprintf("SELECT filepath FROM files WHERE filehash=?")
	stmt, _ := db.Prepare(sqlStr)
	result := stmt.QueryRow(hash)
	result.Scan(&path)
	return
}

func Gethash(filename string) (hash string) {
	sqlStr := fmt.Sprintf("SELECT filehash FROM userfiles WHERE filename=?")
	stmt, _ := db.Prepare(sqlStr)
	result := stmt.QueryRow(filename)
	result.Scan(&hash)
	return
}
