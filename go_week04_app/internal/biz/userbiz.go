package biz

type UserDO struct {
	Id string
	Name string
	Age int
	UserName string
}

type UserRepo interface {
	Get(string) UserDO
}

type UserBiz struct {
	repo UserRepo
}

func initUserBiz(repo UserRepo) *UserBiz {
	return &UserBiz{repo: repo}
}

func (u *UserBiz) Get(id string) UserDO {
	return u.repo.Get(id)
}