package article
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wuyan94zl/api/app/models/article"
	"github.com/wuyan94zl/api/pkg/orm"
	"github.com/wuyan94zl/api/pkg/utils"
	"strconv"
)

func Create(c *gin.Context) {
	// 验证参数
	data := make(map[string][]string)

	data["title"] = []string{"required","min:10","max:50"} 
	data["description"] = []string{"required","min:10","max:200"} 
	data["content"] = []string{"required"} 
	data["admin_id"] = []string{"required","numeric"} 

	validate := utils.Validator(c.Request, data)
	if validate != nil{
		utils.SuccessErr(c,403,validate)
		return
	}
	var Article article.Article
	Article.Title = c.PostForm("title")
	Article.Description = c.PostForm("description")
	Article.Content = c.PostForm("content")
	View,_ := strconv.Atoi(c.PostForm("view"))
	Article.View = uint64(View)
	AdminId,_ := strconv.Atoi(c.PostForm("admin_id"))
	Article.AdminId = uint64(AdminId)

	orm.GetInstance().Create(&Article)
	utils.SuccessData(c, Article) // 返回创建成功的信息
}
func Update(c *gin.Context) {
	// 验证参数
	data := make(map[string][]string)

	data["title"] = []string{"required","min:10","max:50"} 
	data["description"] = []string{"required","min:10","max:200"} 
	data["content"] = []string{"required"} 
	data["admin_id"] = []string{"required","numeric"} 

	validate := utils.Validator(c.Request, data)
	if validate != nil{
		utils.SuccessErr(c,403,validate)
		return
	}
	id, _ := strconv.Atoi(c.Query("id"))
	var Article article.Article
	orm.GetInstance().First(&Article,id,"Admin")
	if Article.Id == 0 {
		utils.SuccessErr(c, -1000, "数据不存在")
		return
	}

	Article.Title = c.PostForm("title")
	Article.Description = c.PostForm("description")
	Article.Content = c.PostForm("content")
	View,_ := strconv.Atoi(c.PostForm("view"))
	Article.View = uint64(View)
	AdminId,_ := strconv.Atoi(c.PostForm("admin_id"))
	Article.AdminId = uint64(AdminId)

	orm.GetInstance().Save(Article)
	utils.SuccessData(c, Article) // 返回创建成功的信息
}
func Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	var Article article.Article

	orm.GetInstance().First(&Article,id)
	if Article.Id == 0 {
		utils.SuccessErr(c, -1000, "数据不存在")
		return
	}
	orm.GetInstance().Delete(&Article)
	utils.SuccessData(c, "删除成功")
}
func Info(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	var Article article.Article
	orm.GetInstance().First(&Article,id,"Admin")

	utils.SuccessData(c, Article)
}	
func Paginate(c *gin.Context) {
	where := make(map[string]interface{})
	Title := c.PostForm("title")
	if Title != "" {
		where["title"] = orm.Where{Way: "LIKE",Value:fmt.Sprintf("%s%s", Title, "%")}
	}

	var Article []article.Article
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	lists := orm.SetPageList(&Article, int64(page))
	orm.GetInstance().Where(where).Paginate(lists,"Admin")
	utils.SuccessData(c, lists)
}