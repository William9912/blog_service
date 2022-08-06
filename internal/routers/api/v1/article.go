package v1

import (
	"blog-service/pkg/app"
	"blog-service/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}

func (a Article) Get(c *gin.Context) {
	//下列函数返回的是gin.Context.H 本质上和直接
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"message": "pong"})
	// })
	//没什么区别 就是套了一层
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
}
func (a Article) List(c *gin.Context)   {}
func (a Article) Create(c *gin.Context) {}
func (a Article) Update(c *gin.Context) {}
func (a Article) Delete(c *gin.Context) {}
