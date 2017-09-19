package db

import (
	"database/sql"
)

type User struct {
	UserId string `json:"user_id"`
	Mobile string `json:"mobile"`
	NickName string `json:"nick_name"`
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

func (this *User)add() (sql.Result, error)  {
	sql := "insert into user(user_id, mobile, nick_name, 
		head_img_url, sex, ref_user_id, login_count, last_login_time, create_time, token, wilddog_token, amount)
		 values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	return UserDb.Exec(sql, this.UserId, this.Mobile, this.NickName, this.HeadImgUrl, this.Sex, 
		this.RefUserId, this.LoginCount, this.LastLoginTime, this.CreateTime, this.Token, this.WilddogToken, this.Amount)
}

func (this *User)query() error {
	sql := "select * from user where user_id=?"
	rows, err := UserDb.Query(sql, this.UserId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&this.Mobile, &this.NickName, &this.HeadImgUrl, &this.Sex, 
			&this.RefUserId, &this.LoginCount, &this.LastLoginTime, &this.CreateTime, 
			&this.Token, &this.WilddogToken, &this.Amount)
	}
	return nil
}