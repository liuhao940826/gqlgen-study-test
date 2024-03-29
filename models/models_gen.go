// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

type Person interface {
	IsPerson()
}

// 自己测试的玩家用户实体对象
type Player struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int64  `json:"age"`
	// 下面是玩家自己的属性,上面是继承Person的
	Game string `json:"game"`
}

func (Player) IsPerson() {}

// 用户注册入参实体对象
type RegisterPlayerReq struct {
	// 注册用户名字
	Name string `json:"name"`
	// 注册用户年龄
	Age int64 `json:"age"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int64  `json:"age"`
	// 用户的地址
	Address *string `json:"address"`
}

func (User) IsPerson() {}
