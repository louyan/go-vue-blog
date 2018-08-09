package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"vue-blog-api/models"

	"github.com/astaxie/beego"
	"fmt"
	"github.com/astaxie/beego/validation"
)

// UserController operations for User
type UserController struct {
	beego.Controller
}

// URLMapping ...
func (c *UserController) URLMapping() {
	c.Mapping("Register", c.Register)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Login", c.Login)
}

// Login ...
// @Title Login
// @Description Login
// @Param	username   query 	string	true		"username for User content"
// @Param	password   query 	string	true		"password for User content"
// @Success 201 {int} models.User
// @Failure 403 body is empty
// @router /login [post]
func (c *UserController) Login() {
	result := make(map[string]string)
	username := c.GetString("username")
	password := c.GetString("password")
	//password := c.Ctx.Input.Param("password")
	if password == ""{
		c.Ctx.WriteString("password is empty")
		return
	}
	fmt.Println("username:",username)
	fmt.Println("password:",password)
	user := models.User{Name:username,Password:password}
	if err := user.Read("name","password"); err== nil{
		c.Ctx.Output.SetStatus(201)
		result["sucsess"] = "1"
		result["message"] = "登陆成功"
	}else {
		result["code"] = "-1"
		result["message"] = "用户名或密码错误"
		result["error"] =   err.Error()
	}

	//v,err := models.GetUserByName(username,password)
	//if  err != nil{
	//	c.Data["json"] = err.Error()
	//}else {
	//	c.Ctx.Output.SetStatus(201)
	//	c.Data["json"] = v
	//}
	c.Data["json"] = result
	c.ServeJSON()
}

// Register ...
// @Title Register
// @Description create User
// @Param	body		body 	models.User	 true		"body for User content"
// @Success 201 {int} models.User
// @Failure 403 body is empty
// @router /register [post]
func (c *UserController) Register() {
	result := make(map[string]string)
	var u models.User

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &u); err == nil {
		username := strings.TrimSpace(c.GetString("username"))
		fmt.Println("username:",c.GetString("username"))
		password := strings.TrimSpace(c.GetString("password"))

		fmt.Println(&u)

		valid := validation.Validation{}
		if v := valid.Required(username,"username"); !v.Ok{
			result["username"] = "请输入用户名"

		}else {
			u.Name = username
			fmt.Println(u.Name)
		}
		if v := valid.Required(password,"password"); !v.Ok{
			result["password"] = "请输入密码"

		}else {
			u.Password = password
			fmt.Println("password:",u.Password)
		}
		//u.Name = c.Ctx.Input.Param("username")
		//fmt.Println(u.Name)
		//user.Password = util.Md5([]byte(c.Ctx.Input.Param("password")))
		//fmt.Println(user.Password)
		//if _, err := models.AddUser(&u); err == nil {
		//	c.Ctx.Output.SetStatus(201)
		//	c.Data["json"] = "Register ok"
		//} else {
		//	c.Data["json"] = err.Error()
		//}
		//u.Name = username
		//fmt.Println(u.Name)
		if err := u.Read("name");err ==nil {
			result["message"] = "用户名已存在"
		}else {
			if err := u.Insert();err !=nil {
				result["code"] = "0"
				result["message"] = "数据创建失败"
				result["debug"] = err.Error()
			}else {
				c.Ctx.Output.SetStatus(201)
				result["code"] = "1"
				result["message"] = "注册成功"
				//result["data"] =""
			}
		}


	} else {
		result["code"] = "failed"
		result["message"] = "请求数据解析失败"
		result["debug"] = err.Error()
	}
	c.Data["json"] = result
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get User by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :id is empty
// @router /:id [get]
func (c *UserController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetUserById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get User
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.User
// @Failure 403
// @router / [get]
func (c *UserController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

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

	l, err := models.GetAllUser(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the User
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.User	true		"body for User content"
// @Success 200 {object} models.User
// @Failure 403 :id is not int
// @router /:id [put]
func (c *UserController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.User{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateUserById(&v); err == nil {
			c.Data["json"] = "OK"
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
// @Description delete the User
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *UserController) Delete() {

	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteUser(id); err == nil {
		c.Data["json"] = "Delete Success!!!"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
