package main
import (
	Main "https/gin_router"
	api "https/gin_api"
	"github.com/gin-gonic/gin"
	. "https/data"
	Func "https/func_judge"
	"net/http"
	)
func main() {
	Main.InitRouter()
}

type User struct {
	Id     int
	Name   string
	Passwd string
}
var Slice []User
var State = make(map[string]interface{})

package gin_router
import (
)
func InitRouter() {
	gin.SetMode(gin.ReleaseMode)
	//使用gin的Default方法建立一个路由handler
	router := gin.Default()
	//访问一个错误网站时返回
	router.NoRoute(api.NotFound)
	//使用gin中Group进行分组
	v1 := router.Group("admin")
	{
		v1.GET("/register", api.Register)
		v1.GET("/login", api.Login)
	}
}

func Register(c *gin.Context) {
	//获取姓名和密码
	name := c.Request.FormValue("Name")
	passwd := c.Request.FormValue("Passwd")
	//判断账户是否存在，若不存在则建立，保存姓名和密码
	Bool := Func.IsExist(name)
	if Bool {
		State["state"] = 1
		State["text"] = "此使用者已存在！"
	} else {
		AddStruct(name, passwd)
		State["state"] = 1
		State["text"] = "註冊成功！"
	}
	//返回给服务端
	c.String(http.StatusOK, "%v", State)
}
func Login(c *gin.Context) {
	name := c.Request.FormValue("Name")
	passwd := c.Request.FormValue("Passwd")
	//先判断用户是否存在，再判断密码是否正确
	Bool := Func.IsExist(name)
	if Bool {
		Bool_Pwd := Func.IsRight(name, passwd)
		if Bool_Pwd {
			State["state"] = 1
			State["text"] = "登入成功！"
		} else {
			State["state"] = 0
			State["text"] = "密碼錯誤！"
		}
	} else {
		State["state"] = 2
		State["text"] = "登入失敗！此使用者尚未註冊！"
	}

	c.String(http.StatusOK, "%v", State)
}
//添加账户
func AddStruct(name string, passwd string) {
	var user User
	user.Name = name
	user.Passwd = passwd
	user.Id = len(Slice) + 1
	Slice = append(Slice, user)
}
isExist.go:

package func_judge
import (
. "https/data"
)

func IsExist(user string) bool {
	//如果长度为0说明没有人注册
	if len(Slice) == 0 {
		return false
	} else {
		for _, v := range Slice {
			// return v.Name == user 此时和第一个比较
			if v.Name == user {
				return true
			}
		}
	}
	return false
}
//判断密码是否正确
func IsRight(user string, passwd string) bool {
	for _, v := range Slice {
		if v.Name == user {

			return v.Passwd == passwd
		}
	}
	return false
}