package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-mysql/mysql"
	"net/http"
	"strconv"
	"time"
)

// Query account result
type Go struct {
	Id         uint32 `json:"id"`
	Name       string `json:"name"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}

var engine = mysql.GetDB()

func (s Server) get(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, "invalid parameter name")
		return
	}
	var data = &Go{}
	_, err := engine.Table("go").Where("name = ?", name).Get(data)
	if err != nil {
		fmt.Println("get sql err: ", err)
	}
	if data.Id > 0 {
		c.JSON(http.StatusOK, data)
		return
	}
	c.JSON(http.StatusNotFound, "not found")
	return
}

func (s Server) list(c *gin.Context) {

	var (
		data []*Go
	)

	err := engine.OrderBy("id").Limit(10, 0).Find(&data)
	if err != nil {
		fmt.Println("get list err: ", err)
	}
	c.JSON(http.StatusOK, data)
}

func (s Server) create(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, "invalid parameter")
		return
	}
	now := time.Now().Unix()
	resp := Go{
		Name:       name,
		CreateTime: now,
		UpdateTime: now,
	}
	_, err := engine.Insert(&resp)
	if err != nil {
		fmt.Println("create sql err: ", err)
	}
	c.JSON(http.StatusOK, "success create")
}

func (s Server) update(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, "invalid parameter name")
		return
	}
	Id := c.Query("id")
	if Id == "" {
		c.JSON(http.StatusBadRequest, "invalid parameter id")
		return
	}
	idNum, err := strconv.Atoi(Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid parameter id")
		return
	}
	now := time.Now().Unix()
	resp := Go{
		Name:       name,
		UpdateTime: now,
	}
	_, err = engine.Update(&resp, &Go{Id: uint32(idNum)})
	if err != nil {
		fmt.Println("Update sql err: ", err)
	}
	c.JSON(http.StatusOK, "success Update")
}

func (s Server) del(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, "invalid parameter id")
		return
	}
	idNum, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid parameter id")
		return
	}
	var table = Go{}
	_, err = engine.Where("id=?", uint32(idNum)).Delete(&table)
	if err != nil {
		fmt.Println("del sql err: ", err)
	}
	c.JSON(http.StatusOK, "success del")
}
