package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginMiddlewareBuilder struct {
}

/*


在 Go 语言中，type LoginMiddlewareBuilder struct {} 定义了一个空的结构体类型 LoginMiddlewareBuilder。这个结构体没有包含任何字段，因此它本身不存储任何数据。
通常，这样的空结构体可能被用作以下几种场景之一：
标记类型：它可以作为某种标记或者标识，用于区分不同的类型，即使它们不携带任何状态信息。例如，在实现接口时，如果接口不需要任何状态，那么可以用空结构体作为实现该接口的类型。
构建模式（Builder Pattern）：虽然结构体本身是空的，但是可以通过在结构体上定义方法来实现构建模式。方法可以接收指向该结构体的指针，并允许逐步设置或修改其内部状态（即使状态不在结构体中）。这通常涉及到链式调用，最终返回一个配置好的对象。
占位符：在开发初期，可能还没有确定结构体应该包含哪些字段，但需要先定义其类型以便进行框架搭建或接口设计。

对于 LoginMiddlewareBuilder 这个特定的例子，它可能是用于构建登录中间件的构建器模式的一部分。通常，你会在 LoginMiddlewareBuilder 上定义一些方法，用于配置中间件的行为，然后有一个 Build 或者类似的函数来创建实际的中间


**/

func (m *LoginMiddlewareBuilder) CheckLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		if path == "/users/signup" || path == "/users/login" {
			// 不需要登录校验
			return
		}
		sess := sessions.Default(ctx)
		if sess.Get("userId") == nil {
			// 中断，不要往后执行，也就是不要执行后面的业务逻辑
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
