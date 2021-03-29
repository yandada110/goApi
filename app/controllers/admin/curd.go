package admin
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wuyan94zl/api/app/models/admin"
	"github.com/wuyan94zl/api/pkg/orm"
	"github.com/wuyan94zl/api/pkg/utils"
	"strconv"
	"golang.org/x/crypto/bcrypt"
)

func Create(c *gin.Context) {
	// 验证参数
	data := make(map[string][]string)

	data["email"] = []string{"required","min:6","email"} 
	data["password"] = []string{"min:6"} 
	data["name"] = []string{"required","min:6"} 
	data["phone"] = []string{"required","min:11","max:11"} 

	validate := utils.Validator(c.Request, data)
	if validate != nil{
		utils.SuccessErr(c,403,validate)
		return
	}
	var Admin admin.Admin
	pwd, _ := bcrypt.GenerateFromPassword([]byte(c.PostForm("password")), bcrypt.DefaultCost)
	Admin.Email = c.PostForm("email")
	Admin.Password = string(pwd)
	Admin.Name = c.PostForm("name")

	Admin.Phone = c.PostForm("phone")
	orm.GetInstance().Create(&Admin)
	utils.SuccessData(c, Admin) // 返回创建成功的信息
}
func Update(c *gin.Context) {
	// 验证参数
	data := make(map[string][]string)

	data["email"] = []string{"required","min:6","email"} 
	data["password"] = []string{"min:6"} 
	data["name"] = []string{"required","min:6"} 
	data["phone"] = []string{"required","min:11","max:11"} 

	validate := utils.Validator(c.Request, data)
	if validate != nil{
		utils.SuccessErr(c,403,validate)
		return
	}
	id, _ := strconv.Atoi(c.Query("id"))
	var Admin admin.Admin
	orm.GetInstance().First(&Admin,id,"Roles")
	if Admin.Id == 0 {
		utils.SuccessErr(c, -1000, "数据不存在")
		return
	}

	pwd, _ := bcrypt.GenerateFromPassword([]byte(c.PostForm("password")), bcrypt.DefaultCost)
	Admin.Email = c.PostForm("email")
	Admin.Password = string(pwd)
	Admin.Name = c.PostForm("name")

	Admin.Phone = c.PostForm("phone")
	orm.GetInstance().Save(Admin)
	utils.SuccessData(c, Admin) // 返回创建成功的信息
}
func Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	var Admin admin.Admin

	orm.GetInstance().First(&Admin,id)
	if Admin.Id == 0 {
		utils.SuccessErr(c, -1000, "数据不存在")
		return
	}
	orm.GetInstance().Delete(&Admin)
	utils.SuccessData(c, "删除成功")
}
func Info(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	var Admin admin.Admin
	orm.GetInstance().First(&Admin,id,"Roles")

	utils.SuccessData(c, Admin)
}	
func Paginate(c *gin.Context) {
	where := make(map[string]interface{})
	Email := c.PostForm("email")
	if Email != "" {
		where["email"] = orm.Where{Way: "LIKE",Value:fmt.Sprintf("%s%s", Email, "%")}
	}
	Name := c.PostForm("name")
	if Name != "" {
		where["name"] = orm.Where{Way: "LIKE",Value:fmt.Sprintf("%s%s", Name, "%")}
	}
	Phone := c.PostForm("phone")
	if Phone != "" {
		where["phone"] = Phone
	}

	var Admin []admin.Admin
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	lists := orm.SetPageList(&Admin, int64(page))
	orm.GetInstance().Where(where).Paginate(lists,"Roles")
	utils.SuccessData(c, lists)
}