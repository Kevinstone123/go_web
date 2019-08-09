package dbops

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func AddUserCredential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("")
	if err != nil {
		return err
	}
	stmtIns.Exec()
	defer stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("")
	if err != nil {
		log.Printf("%s", err)
		return  "" , err
	}

	var pwd string
	stmtOut.QueryRow(loginName).Scan(&pwd)
	defer stmtOut.Close()

	return pwd, nil
}