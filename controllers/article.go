package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"vue-blog-api/models"

	"github.com/astaxie/beego"
	"fmt"
	"github.com/astaxie/beego/orm"
)

// ArticleController operations for Article
type ArticleController struct {
	beego.Controller
}

// URLMapping ...
func (c *ArticleController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Article
// @Param	body		body 	models.Article	true		"body for Article content"
// @Success 201 {int} models.Article
// @Failure 403 body is empty
// @router / [post]
func (c *ArticleController) Post() {
	var v models.Article
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddArticle(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Article by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Article
// @Failure 403 :id is empty
// @router /getOne/:id [get]
func (c *ArticleController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	result := make(map[string]interface{})
	artcile := []models.Article{}

	//pre := make(map[string]string)
	//next := make(map[string]string)

	prevNextResults,err := orm.NewOrm().Raw("(select id,title from article where id < ? AND status = 'PUBLISHED' order by id desc limit 0,1) union all "+
		"(select id,title from article where id > ? AND status = 'PUBLISHED' order by id asc limit 0,1)",id,id).QueryRows(&artcile)
	if err ==nil {
		fmt.Printf("从article表中共查询到记录:%d条\n", prevNextResults)
		fmt.Println(artcile)
		prev := artcile[0]
		next := artcile[1]

		fmt.Println(prev)
		fmt.Println(next.Title)
		result["prev"] = prev
		result["next"] = next

	}

	v, err := models.GetArticleById(id)
	if err != nil {
		result["success"] = 0
		result["message"]= "查询数据出错"
		result["post"] = nil
	} else {
		result["success"] = 1
		result["message"] = nil
		result["post"] = v

	}
	c.Data["json"] = result
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Article
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Article
// @Failure 403
// @router /getAll [get]
func (c *ArticleController) GetAll() {
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

// Put ...
// @Title Put
// @Description update the Article
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Article	true		"body for Article content"
// @Success 200 {object} models.Article
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ArticleController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Article{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateArticleById(&v); err == nil {
			//c.Ctx.Output.Body([]byte("更新成功"))
			//c.Data["json"] = "Put Success!!!"
			c.Data["json"] = "{message:更新成功,}"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Article
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ArticleController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteArticle(id); err == nil {
		c.Data["json"] = "Delete Success!!!"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
