package model

import "github.com/jinzhu/gorm"

type Auth struct {
	*Model
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
}

func (a Auth) TableName() string {
	return "blog_auth"
}

//解释一下这个表的作用，其实就是需要这个表里有你的发过来的内容，才会给你token 给你
//token之后 你才能使用功能呢 至于表里怎么有发过来的内容 估计是注册的时候插入到数据库的
//2 git rebase test
//3 git rebase test
//git rebase test 4
//git rebase test 5
func (a Auth) Get(db *gorm.DB) (Auth, error) {
	var auth Auth
	db = db.Where("app_key = ? AND app_secret = ? AND is_del = ?", a.AppKey, a.AppSecret, 0)
	err := db.First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return auth, err
	}

	return auth, nil
}
