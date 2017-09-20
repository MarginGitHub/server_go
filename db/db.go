package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
)

const DBPATH = `D:\dev\server\go\src\github.com\margin\server\db\`

var UserDb *sql.DB

var DoctorDb *sql.DB

func init()  {
	userDb, err := sql.Open("sqlite3", DBPATH + "user.db")
	if err == nil {
		UserDb = userDb
		_, err := UserDb.Exec(`
			create table if not exists user(
				user_id integer primary key autoincrement,
				mobile varchar(11) not null,
				password text not null,
				nick_name text not null,
				head_img_url text,
				sex varchar(1),
				ref_user_id text,
				login_count integer,
				last_login_time integer,
				create_time integer,
				token text,
				wilddog_token text,
				amount float
				)`)
		if err != nil {
			fmt.Printf("err : %v\n", err)
		} else {
			fmt.Println("创建用户表成功")
		}
	}

	doctorDb, err := sql.Open("sqlite3", DBPATH + "doctor.db")
	if err == nil {
		DoctorDb = doctorDb
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