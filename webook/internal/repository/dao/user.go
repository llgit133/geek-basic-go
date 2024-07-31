package dao

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
)

/*
1.errors.New("邮箱冲突") 是使用标准库 errors 包中的 New 函数创建一个新的错误对象
2.gorm.ErrRecordNotFound 是 GORM ORM 框架中预定义的一个错误，用于表示在数据库中没有找到匹配记录的情况。
*
*/
var (
	ErrDuplicateEmail = errors.New("邮箱冲突")
	ErrRecordNotFound = gorm.ErrRecordNotFound
)

type UserDAO struct {
	db *gorm.DB
}

// 初始化userDao 结构体
// 构造函数 用于创建并初始化结构体的实例
func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}

// UserDAO 参数插入到User表数据
// 尝试创建新用户时，如果遇到由于邮箱重复导致的数据库错误，就会捕获这个错误并返回一个专门的错误 ErrDuplicateEmail
func (dao *UserDAO) Insert(ctx context.Context, u User) error {
	now := time.Now().UnixMilli()
	u.Ctime = now
	u.Utime = now
	err := dao.db.WithContext(ctx).Create(&u).Error
	if me, ok := err.(*mysql.MySQLError); ok {
		const duplicateErr uint16 = 1062
		if me.Number == duplicateErr {
			// 用户冲突，邮箱冲突
			return ErrDuplicateEmail
		}
	}
	return err
}

// 查找用户，根据邮箱查找用户
// select * from user where email=?

/*
*

err := dao.db.WithContext(ctx).Where("email=?", email).First(&u).Error
1.使用 dao.db（即 gorm.DB 的实例）执行数据库操作。
2.WithContext(ctx) 方法将上下文 ctx 应用于数据库操作，使操作能够响应上下文的取消或超时。
3.Where("email=?", email) 方法指定查询条件，即查找 email 字段等于给定 email 参数的记录。
4.First(&u) 方法执行查询并获取第一个匹配的记录，结果将直接赋值给 u 变量。如果找不到匹配的记录，u 将保持为 User 类型的零值。
5. .Error 属性用于获取执行上述操作时发生的任何错误。

*
*/
func (dao *UserDAO) FindByEmail(ctx context.Context, email string) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Where("email=?", email).First(&u).Error
	return u, err
}

type User struct {
	Id       int64  `gorm:"primaryKey,autoIncrement"`
	Email    string `gorm:"unique"`
	Password string

	// 时区，UTC 0 的毫秒数
	// 创建时间
	Ctime int64
	// 更新时间
	Utime int64

	// json 存储
	//Addr string
}

//type Address struct {
//	Uid
//}
