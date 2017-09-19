package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

const DBPATH = `D:\dev\server\go\src\github.com\margin\server\db\`

var UserDb *sql.DB = func ()  {
	if db, err := sql.Open("sqlite3", DBPATH + "user.db"); err == nil {
		return db
	} else {
		return nil
	}
}

var DoctorDb *sql.DB = func ()  {
	if db, err := sql.Open("sqlite3", DBPATH + "doctor.db"); err == nil {
		return db
	} else {
		return nil
	}
}

func Close()  {
	if UserDb != nil {
		UserDb.Close()
		UserDb = nil
	}
	if DoctorDb != nil {
		DoctorDb.Close()
		DoctorDb = nil
	}
}