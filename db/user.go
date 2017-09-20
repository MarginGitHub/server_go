package db

import (
	// "database/sql"
)

type User struct {
	UserId string `json:"user_id"`
	Mobile string `form:"mobile" json:"mobile"`
	Password string `form:"password" json:"password"`
	NickName string `form:"nick_name" json:"nick_name"`
	HeadImgUrl string `json:"head_img_url"`
	Sex string `json:"sex"`
	RefUserId string `json:"ref_user_id"`
	LoginCount int `json:"login_count"`
	LastLoginTime int64 `json:"last_login_time"`
	CreateTime string `json:"create_time"`
	Token string `json:"token"`
	WilddogToken string `json:"wilddog_token"`
	Amount float32 `json:"amount"`
}

func (this *User)Add() (int64, error)  {
	sql := "insert into user(mobile, password, nick_name, " +  
		"head_img_url, sex, ref_user_id, login_count, last_login_time, create_time, token, wilddog_token, amount) " +
		 "values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	ret, err := UserDb.Exec(sql, this.Mobile, this.Password, this.NickName, this.HeadImgUrl, this.Sex, 
		this.RefUserId, this.LoginCount, this.LastLoginTime, this.CreateTime, this.Token, this.WilddogToken, this.Amount)
	if err != nil {
		return -1, err
	}
	return ret.LastInsertId()
}

func (this *User)Query() error {
	sql := "select * from user where user_id=?"
	rows, err := UserDb.Query(sql, this.UserId)
	if err != nil {
		return err
	}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&this.UserId, &this.Mobile, &this.Password, &this.NickName, &this.HeadImgUrl, &this.Sex, 
			&this.RefUserId, &this.LoginCount, &this.LastLoginTime, &this.CreateTime, 
			&this.Token, &this.WilddogToken, &this.Amount)
	}
	return nil
}

func QueryAllUsers() ([]User, error) {
	sql := "select * from user"
	rows, err := UserDb.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	ret := []User{}
	for rows.Next() {
		user := new(User)
		rows.Scan(&user.UserId, &user.Mobile, &user.Password, &user.NickName, &user.HeadImgUrl, &user.Sex, 
			&user.RefUserId, &user.LoginCount, &user.LastLoginTime, &user.CreateTime, &user.Token, &user.WilddogToken, &user.Amount)
		ret = append(ret, *user)
	}
	return ret, nil
}

func QueryUsersWithRange(start, size int) ([]User, error) {
	sql := "select * from user order by user_id limit ? offset ?"
	rows, err := UserDb.Query(sql, start, size)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	ret := []User{}
	for rows.Next() {
		user := new(User)
		rows.Scan(&user.UserId, &user.Mobile, &user.Password, &user.NickName, &user.HeadImgUrl, &user.Sex, 
			&user.RefUserId, &user.LoginCount, &user.LastLoginTime, &user.CreateTime, &user.Token, &user.WilddogToken, &user.Amount)
		ret = append(ret, *user)
	}
	return ret, nil
}

