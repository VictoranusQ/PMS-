package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB //初始化全局变量DB
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

//连接数据库
func initMySQL() (err error) {
	dsn := "root:985619714@tcp(localhost)/pms?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	sqlDB, err := DB.DB()
	return sqlDB.Ping()
}

//权限检测中间件(是否被冻结)(或者做在登陆界面更好？)
func freezeMiddleware() gin.HandlerFunc {
	//连接数据库
	err := initMySQL()
	if err != nil {
		panic(err)
	}
	//sqlDB, err := DB.DB()
	//defer sqlDB.Close() //程序退出关闭数据库连接
	//模型绑定
	DB.AutoMigrate(&UserInfo{})
	return func(c *gin.Context) {
		//从url获取参数userid
		userid, ok := c.Params.Get("userid")
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"err_code": "1001",
				"err_msg":  "处理失败",
				"data":     nil,
			})
			return
		}
		var user UserInfo //登录用户
		DB.First(&user, userid)
		//查询数据库中id=userid的用户职能，if status = 1(未冻结)
		if user.Status == true {
			c.Next()
		} else {
			c.Abort()
		}
	}
}

//修改中间件
func modifyMiddleware() gin.HandlerFunc {
	//连接数据库
	err := initMySQL()
	if err != nil {
		panic(err)
	}
	//sqlDB, err := DB.DB()
	//defer sqlDB.Close() //程序退出关闭数据库连接
	//模型绑定
	DB.AutoMigrate(&UserInfo{})
	return func(c *gin.Context) {
		//从url路径中获取参数
		id, ok := c.Params.Get("id")
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"err_code": "1001",
				"err_msg":  "处理失败",
				"data":     nil,
			})
			return
		}
		userid, ok := c.Params.Get("userid")
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"err_code": "1001",
				"err_msg":  "处理失败",
				"data":     nil,
			})
			return
		}
		//若为HR/若为Leader且所选成员为该部门
		var user1 UserInfo //已登录用户
		var user2 UserInfo //被修改用户
		DB.First(&user1, userid)
		DB.First(&user2, id)
		//查询数据库中id=userid的用户职能，if为HR则
		if user1.Position == "HR" {
			c.Next()
		} else if user1.Position == "Leader" && user2.Dept == user1.Dept { //查询数据库中id=userid，id=id的用户职能，若为Leader且dept相同则
			c.Next()
		} else {
			c.Abort()
		}
	}
}

func main() {
	//创建数据库
	//sql:create database PMS
	//连接数据库
	err := initMySQL()
	if err != nil {
		panic(err)
	}
	sqlDB, err := DB.DB()
	defer sqlDB.Close() //程序退出关闭数据库连接
	//模型绑定
	DB.AutoMigrate(&UserInfo{})
	//创建路由
	r := gin.Default()

	//访问人员管理系统操作主界面(未冻结用户)，中间件浏览权限(或者做在登陆界面更好？)
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"err_code": "0000",
			"err_msg":  "处理成功",
			"data":     nil,
		})
	})

	//人员管理操作
	v1Group := r.Group("v1") //v1为人员管理操作组
	{
		//userid 为已登录用户的id
		//列表查看（checklist）
		v1Group.GET("/:userid/check", func(c *gin.Context) {
			//前端页面点击 列表查看 发送请求到这里
			//查询UserInfo里的所有数据
			var userList []UserInfo
			if err = DB.Find(&userList).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"err_code": "1001",
					"err_msg":  "处理失败",
					"data":     nil,
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"err_code": "0000",
					"err_msg":  "处理成功",
					"data":     userList,
				})
			}
			//从数据库中调取所有成员信息

			//分页列表(前端？)返回所有成员信息
		})
		//查看某一个用户详情(checkone)
		v1Group.GET("/:userid/check/:id", func(c *gin.Context) {
			//前端页面点击 查看某个用户 发送请求到这里

			//从请求/url中读取数据(id)，在数据库中查找到该成员
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{
					"err_code": "1001",
					"err_msg":  "处理失败",
					"data":     nil,
				})
				return
			}
			//返回该成员信息
			var user UserInfo
			if err = DB.Find(&user, id).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"err_code": "1001",
					"err_msg":  "处理失败",
					"data":     nil,
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"err_code": "0000",
					"err_msg":  "处理成功",
					"data":     user,
				})
			}
		})
		//查看信息并导出为excel表格形式(checkexcel)
		v1Group.GET("/:userid/excel", func(c *gin.Context) {
			var userList []UserInfo
			if err = DB.Find(&userList).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"err_code": "1001",
					"err_msg":  "处理失败",
					"data":     nil,
				})
			} else {
				f := excelize.NewFile()
				index := 1
				for _, item := range *&userList {
					strIndex := strconv.Itoa(index)
					f.SetCellValue("Sheet1", "A"+strIndex, item.ID)
					f.SetCellValue("Sheet1", "B"+strIndex, item.Name)
					f.SetCellValue("Sheet1", "C"+strIndex, item.Gender)
					f.SetCellValue("Sheet1", "D"+strIndex, item.Major)
					f.SetCellValue("Sheet1", "E"+strIndex, item.Facaulty)
					f.SetCellValue("Sheet1", "F"+strIndex, item.Dept)
					f.SetCellValue("Sheet1", "G"+strIndex, item.Position)
					f.SetCellValue("Sheet1", "H"+strIndex, item.Status)
					index++
				}

				if err := f.SaveAs("易千.xlsx"); err != nil {
					fmt.Println(err)
				}
				c.Header("Content-Type", "application/octet-stream")
				c.Header("Content-Disposition", "attachment; filename="+"易千.xlsx")
				c.Header("Content-Transfer-Encoding", "binary")

				//回写到web 流媒体 形成下载
				_ = f.Write(c.Writer)
			}
		})
		//修改或冻结某一位人员信息
		v1Group.PUT("/:userid/:id", modifyMiddleware(), func(c *gin.Context) {
			//前端页面点击 修改 发送请求到这里

			//从请求/url中读取数据(userid)，中间件检测权限：若为HR，则放行；若为Leader且所选成员为该部门，则放行；else则拒绝

			//从请求/url中读取数据(id)，在数据库中找到该成员
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{
					"err_code": "1001",
					"err_msg":  "处理失败",
					"data":     nil,
				})
				return
			}

			var user UserInfo
			if err = DB.Where("id = ?", id).First(&user).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"err_code": "1001",
					"err_msg":  "处理失败",
					"data":     nil,
				})
			}
			//从请求中获取数据//修改该成员的相应信息，并保存到数据库中
			c.BindJSON(&user)
			if err = DB.Save(&user).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"err_code": "1001",
					"err_msg":  "处理失败",
					"data":     nil,
				})
			} else {
				c.JSON(http.StatusOK, user)
			}

		})
		v1Group.PUT("/:userid/:id/freeze", modifyMiddleware(), func(c *gin.Context) {
			//前端页面点击 冻结 发送请求到这里

			//从请求/url中读取数据(userid)，中间件检测权限：若为HR，则放行；若为Leader且所选成员为该部门，则放行；else则拒绝

			//从请求/url中读取数据(id)，在数据库中找到该成员
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{
					"err_code": "1001",
					"err_msg":  "处理失败",
					"data":     nil,
				})
				return
			}
			//修改该成员的status为0，表示已被冻结
			DB.Model(&UserInfo{}).Where("id = ?", id).Update("status", 0)
		})

	}

	r.Run() //默认为本机端口8080
}
