package controller

import (
	"app/lesson2/dao"
	"app/lesson2/model"
	"app/lesson2/query"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"sync"
)

func CreateUser(c *gin.Context) {
	// 1. 从请求中把数据拿出来
	var user model.SysUser
	er := c.BindJSON(&user)
	if er != nil {
		return
	}
	// 2. 存入数据库
	err := dao.DB.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	err := dao.DB.Where("id=?", id).Delete(&model.SysUser{}).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, id)
	}
}

func UpdateUsers(c *gin.Context) {
	var user model.SysUser
	er := c.BindJSON(&user)
	if er != nil {
		return
	}
	if user.ID == 0 {
		c.JSON(http.StatusOK, gin.H{"error": "ID empty"})
	}

	err := dao.DB.Save(&user).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func PageList(c *gin.Context) {
	var p query.UserQuery
	er := c.ShouldBindQuery(&p)
	if er != nil {
		return
	}
	if p.PageNum <= 0 {
		p.PageNum = 1
	}
	var users []model.SysUser
	tx := dao.DB.Model(&model.SysUser{}).Session(&gorm.Session{})
	if len(p.NameLike) > 0 {
		tx.Where(" name like ?", "%"+p.NameLike+"%")
	}

	if p.AgeStart != 0 {
		tx.Where(" age >= ?", p.AgeStart)
	}

	if p.AgeEnd != 0 {
		tx.Where(" age <= ?", p.AgeEnd)
	}
	var wg sync.WaitGroup
	wg.Add(2)

	var total int64
	go func() {
		defer wg.Add(-1)
		tx.Count(&total)
	}()

	go func() {
		defer wg.Add(-1)
		tx.Order(" ID desc ").Offset(p.PageSize * (p.PageNum - 1)).Limit(p.PageSize).Find(&users)
	}()

	wg.Wait()

	//if total > 0 {
	//	if err := tx.Order(" ID desc ").Offset(p.PageSize * (p.PageNum - 1)).Limit(p.PageSize).Find(&users).Error; err != nil {
	//		c.JSON(http.StatusUnauthorized, gin.H{
	//			"code": 500,
	//			"msg":  err.Error(),
	//		})
	//	} else {
	c.JSON(http.StatusOK, gin.H{"code": 200,
		"msg":      "查询成功",
		"data":     users,
		"total":    total,
		"page_num": p.PageNum,
	})
	//	}
	//} else {
	//	c.JSON(http.StatusOK, gin.H{"code": 200,
	//		"msg":      "查询成功",
	//		"data":     users,
	//		"total":    total,
	//		"page_num": p.PageNum,
	//	})
	//}
}

func GetUserList(c *gin.Context) {
	// 查询user这个表里的所有数据
	var userList []*model.SysUser
	err := dao.DB.Find(&userList).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, userList)
	}
}
