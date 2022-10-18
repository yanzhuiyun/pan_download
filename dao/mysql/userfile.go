package mysql

import (
	"fmt"
	"pandownload/model"
)

func SaveuserFile(file *model.UserFile) error {
	sqlStr := fmt.Sprintf("INSERT INTO userfiles(username,filename,filehash) VALUES(?,?,?)")
	stmt, _ := db.Prepare(sqlStr)
	_, err := stmt.Exec(file.Username, file.Filename, file.Filehash)
	return err
}

func Getfiles(username string) ([]string, error) {
	sqlStr := fmt.Sprintf("SELECT filename FROM userfiles where username=?")
	stmt, _ := db.Prepare(sqlStr)
	rows, err := stmt.Query(username)
	if err != nil {
		return nil, err
	}
	var filenames []string
	var filename string
	for rows.Next() {
		rows.Scan(&filename)
		filenames = append(filenames, filename)

	}
	return filenames, nil
}

func DeleteFile(filename string) error {
	sqlStr := fmt.Sprintf("DELETE FROM userfiles where filename=?")
	stmt, err := db.Prepare(sqlStr)
	_, err = stmt.Exec(filename)
	return err
}
