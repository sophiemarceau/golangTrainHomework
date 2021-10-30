
第二周作业：
问题：我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
回答：
遇到一个 sql.ErrNoRows 我理解
if是确实查不到 导致的error是不需要 wrap这个error
if不是查不到 导致的 而是一些异常 或者代码错误到这个error 出现的问题 是要wrap 这个error 抛出去  防止漏掉代码 bug导致的 sql.ErrNoRows

第二周 代码 如下

package main

import (
	"database/sql"
	"strings"

	"github.com/pkg/errors"
)

type userInfo struct {
	id       int
	age      int
	name     string
	gender   string
	phoneNum string
	email    string
}

func main() {
	queryUserInfo("123232322") //userID
}

func queryUserInfo(userId string) (userInfo, error) {
	var user userInfo
	db, err := InitConnectDB()
	if err != nil {
		return user, err
	}
	return userDao(userId, db)
}

func InitConnectDB() (*sql.DB, error) {
	var userName = ""
	var password = ""
	var ip = ""
	var port = ""
	var dbName = ""
	connectAddress := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	db, _ := sql.Open("mysql", connectAddress)
	if err := db.Ping(); err != nil {
		return nil, errors.Wrap(err, "Database: Failed to connect to database.")
	}
	return db, nil
}

func userDao(userId string, db *sql.DB) (userInfo, error) {
	var user userInfo
	querySQLStr := strings.Join([]string{"select * from user where userId='", userId, "'"}, "")
	err := db.QueryRow(querySQLStr).Scan(&user.id, &user.name, &user.age)
	if err != nil {
		return user, errors.Wrap(err, "Database: "+userId+"'s information could not be found in the database.")
	}
	return user, nil
}
