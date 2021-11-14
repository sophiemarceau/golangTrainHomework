package data

import "go_week04_app/internal/biz"

type UserPO struct {
	Id string
	Name string
	Age int
	UserName string
	Country string
}

type userRepo struct {

}

func InitUserRepo() biz.UserRepo {
	return &userRepo{}
}

func (u *userRepo) Get(id string) biz.UserDO  {
	user := dbGetUser(id)
	return biz.UserDO{
		Id: user.Id,
		Name: user.Name,
		Age: user.Age,
		UserName: user.UserName,
	}
}

func dbGetUser(id string) UserPO {
	return UserPO{
		Id: id,
		Name: "Name" + id,
		Age: 35,
		UserName: "user_" + id,
		Country: "china",
	}
}