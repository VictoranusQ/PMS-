package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

//UserInfo Model
type UserInfo struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`     //姓名
	Gender   string `json:"gender"`   //性别
	Facaulty string `json:"facaulty"` //学院
	Major    string `json:"major"`    //专业
	Position string `json:"position"` //职位：新人 Leader HR 普通用户 行政财务
	Dept     string `json:"dept"`     //部门
	Status   bool   `json:"status"`   //状态：0：被冻结 1：未冻结
}

func initMySQL() (err error) {
	dsn := "root:985619714@(localhost)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}

	return DB.DB().Ping()
}

func main() {
	//创建数据库
	//sql:create database PMS
	//连接数据库

	//创建路由
	r := gin.Default()

	//访问人员管理系统操作主界面(未冻结用户)
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"err_code": "0000",
			"err_msg":  "处理成功",
			"data":     "any",
		})
	})

	//人员管理操作
	v1Group := r.Group("v1")
	{
		//todo 为待办事项
		//查看（check）
		//列表查看（checklist）
		v1Group.GET("/todo/check", func(c *gin.Context) {

		})
		//查看某一个用户详情(checkone)
		v1Group.GET("/todo/check/:id", func(c *gin.Context) {

		})
		//查看信息并导出为excel表格形式(checkexcel)
		v1Group.GET("/todo/excel", func(c *gin.Context) {

		})
		//修改或冻结某一位人员信息
		v1Group.PUT("/todo/:id", func(c *gin.Context) {

		})

	}

	r.Run()
}
