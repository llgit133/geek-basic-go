package main

import (
	"gitee.com/geekbang/basic-go/webook/internal/repository"
	"gitee.com/geekbang/basic-go/webook/internal/repository/dao"
	"gitee.com/geekbang/basic-go/webook/internal/service"
	"gitee.com/geekbang/basic-go/webook/internal/web"
	"gitee.com/geekbang/basic-go/webook/internal/web/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

func main() {

	// 1. 初始化数据库连接
	db := initDB()

	// 2. 初始化路由
	server := initWebServer()

	// 3. 初始化业务逻辑
	initUserHdl(db, server)

	// 4. 启动服务
	server.Run(":8080")
}

// 初始化数据库连接
func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/webook"))
	if err != nil {
		panic(err)
	}

	err = dao.InitTables(db)
	if err != nil {
		panic(err)
	}
	return db
}

// 初始化路由
func initWebServer() *gin.Engine {

	// 1.使用 gin.Default() 创建一个默认的 Gin 服务器。
	server := gin.Default()

	// 2.使用 cors.New() 中间件来配置跨域资源共享（CORS）的规则。
	// 使用一个匿名函数作为中间件，打印出 "这是我的 Middleware"。
	server.Use(cors.New(cors.Config{
		//AllowAllOrigins: true,
		//AllowOrigins:     []string{"http://localhost:3000"},
		AllowCredentials: true,

		AllowHeaders: []string{"Content-Type"},
		//AllowHeaders: []string{"content-type"},
		//AllowMethods: []string{"POST"},
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				//if strings.Contains(origin, "localhost") {
				return true
			}
			return strings.Contains(origin, "your_company.com")
		},
		MaxAge: 12 * time.Hour,
	}), func(ctx *gin.Context) {
		println("这是我的 Middleware")
	})

	// 3.创建一个 LoginMiddlewareBuilder 对象
	login := &middleware.LoginMiddlewareBuilder{}

	// 4.使用 cookie.NewStore() 创建一个 cookie 存储对象。
	// 存储数据的，也就是你 userId 存哪里
	// 直接存 cookie
	store := cookie.NewStore([]byte("secret"))

	// 5.使用 sessions.Sessions() 和 login.CheckLogin() 中间件来处理用户登录状态，并返回初始化好的服务器对象。
	server.Use(sessions.Sessions("ssid", store), login.CheckLogin())
	return server
}

// 初始化业务逻辑
func initUserHdl(db *gorm.DB, server *gin.Engine) {

	// 1.UserDAO 实例，负责数据库操作层的业务
	ud := dao.NewUserDAO(db)

	// 2.封装对数据库的访问逻辑，提供更高级别的数据访问服务
	ur := repository.NewUserRepository(ud)

	// 3.封装对业务逻辑的访问逻辑，提供更高级别的业务逻辑服务
	us := service.NewUserService(ur)

	// 4.负责处理 HTTP 请求和响应，将业务逻辑和前端界面连接起来。
	hdl := web.NewUserHandler(us)

	// 5.注册路由
	hdl.RegisterRoutes(server)
}
