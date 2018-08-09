package admin

import (
	"github.com/astaxie/beego"
	"vue-blog-api/models"
	"strings"
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
	"encoding/json"
)


// PostController operations for Post---admin
type PostController struct {
	beego.Controller
}

// URLMapping ...
func (c *PostController)URLMapping()  {
	c.Mapping("GetPostList",c.GetPostList)
	c.Mapping("GetPostById",c.GetPostById)
	c.Mapping("AddPost",c.AddPost)
	c.Mapping("UpdatePost",c.UpdatePost)
	c.Mapping("GetPostTotal",c.GetPostTotal)
	c.Mapping("OfflinePost",c.OfflinePost)
	c.Mapping("PublishPost",c.PublishPost)
	c.Mapping("DeletePost",c.DeletePost)
	c.Mapping("GetPostsByCatId",c.GetPostsByCatId)
	c.Mapping("GetPostsByCatId",c.GetPostsByTagId)

}


// GetAll ...
// @Title Get All
// @Description get Post
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Article
// @Failure 403
// @router /getPostList [get]
func (c *PostController)GetPostList()  {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64
	var result = make(map[string]interface{})
	//fmt.Println("GetPostList 执行了")
	page,err := c.GetInt64("page")
	if err !=nil{
		fmt.Println(err)
	}
	pageNum,err := c.GetInt64("pageNum")
	if err !=nil{
		fmt.Println(err)
	}

	fmt.Println("page:",page)
	fmt.Println("pageNum:",pageNum)

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllArticle(query, fields, sortby, order, offset, limit)
	if err != nil {
		result["success"] = 0
		result["message"]= "查询数据出错"
		result["posts"] = nil
	} else {
		result["success"] = 1
		result["message"] = nil
		result["posts"] = l
	}
	c.Data["json"] = result
	c.ServeJSON()
}

// @Title Put
// @Description get the Post
// @Param	id		path 	string	true		"The id you want to update"
// @Success 200 {object} models.Article
// @Failure 403 :id is not int
//@router /getPostById/:id [get]
func (c *PostController)GetPostById()  {
	
}

//@router /addPost [post]
func (c *PostController)AddPost()  {
	result := make(map[string]interface{})
	var v models.Article
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddArticle(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			result["success"] = 1
			result["id"] = v.Id
		} else {
			result["success"] = 0
			result["message"]= err.Error()
		}
	} else {
		result["success"] = 0
		result["message"]= err.Error()
	}
	c.Data["json"] = result
	c.ServeJSON()
}

// @Param	id		path 	string	true		"The id you want to update"
//@router /updatePost/:id [put]
func (c *PostController)UpdatePost()  {
	result := make(map[string]interface{})
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Article{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateArticleById(&v); err == nil {
			result["success"] = 1
			result["message"] = "文章更新成功"
		} else {
			result["success"] = 0
			result["message"]= "更新文章出错"
		}
	} else {
		result["success"] = 0
		result["message"]= "更新文章出错"
	}
	c.Data["json"] = result
	c.ServeJSON()
}

// @Title Get Post Total
// @Description get the Post
// @Success 200 {object} models.Article
// @Failure 403 :id is empty
//@router /getPostTotal [get]
func (c *PostController)GetPostTotal() {
	result := make(map[string]interface{})
	postTotal,err := orm.NewOrm().QueryTable("article").Count()
	if err == nil {
		result["success"] = 1
		result["message"] = nil
		result["total"] = postTotal | 0
	}else {
		result["success"] = 0
		result["message"]= "查询数据出错"
		result["total"] = 0
	}
	c.Data["json"] = result
	c.ServeJSON()
}

// @Param	id		path 	string	true		"The id you want to update"
//@router /offlinePost/:id [put]
func (c *PostController)OfflinePost()  {
	result := make(map[string]interface{})

	id := c.Ctx.Input.Param("id")
	_,err := orm.NewOrm().Raw("UPDATE article SET status = 'OFFLINE' WHERE id = "+id).Exec()
	if err == nil {
		result["success"] = 1
		result["message"] = nil
	}else {
		result["success"] = 0
		result["message"]= "文章下线出错"
	}
	c.Data["json"] = result
	c.ServeJSON()
}

// @Param	id		path 	string	true		"The id you want to update"
//@router /publishPost/:id [put]
func (c *PostController)PublishPost()  {
	result := make(map[string]interface{})

	id := c.Ctx.Input.Param("id")
	_,err := orm.NewOrm().Raw("UPDATE article SET status = 'PUBLISHED' WHERE id = "+id).Exec()
	if err == nil {
		result["success"] = 1
		result["message"] = nil
	}else {
		result["success"] = 0
		result["message"]= "文章发布出错"
	}
	c.Data["json"] = result
	c.ServeJSON()
}

// @Param	id		path 	string	true		"The id you want to update"
//@router /deletePost/:id [delete]
func (c *PostController)DeletePost()  {
	result := make(map[string]interface{})
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteArticle(id); err == nil {
		result["success"] = 1
		result["message"] = nil
	} else {
		result["success"] = 0
		result["message"]= "文章删除出错"
	}
	c.Data["json"] = result
	c.ServeJSON()
}

// @Param	id		path 	string	true		"The id you want to update"
//@router /getPostsByCatId/:id [get]
func (c *PostController)GetPostsByCatId()  {

}

// @Param	id		path 	string	true		"The id you want to update"
//@router /getPostsByTagId/:id [get]
func (c *PostController)GetPostsByTagId()  {

}