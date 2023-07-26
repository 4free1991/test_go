package controller

import (
	"app/lesson4/models"
	"app/lesson4/service/personService"
	"errors"
	"github.com/gin-gonic/gin"
)

func Result(c *gin.Context, data any, err error) {
	if err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "error:" + err.Error(),
		})
		c.Abort()
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "ok",
			"data": data,
		})
		c.Abort()
	}
}

func Init(r *gin.Engine) {

	r.GET("/health", func(c *gin.Context) {
		Result(c, nil, nil)
	})

	{
		personApi := r.Group("/api/person")
		personApi.GET("/get", func(c *gin.Context) {

			name := c.Query("name")
			if len(name) == 0 {
				Result(c, nil, errors.New("参数错误"))
				return
			}
			person, err := personService.GetByName(name)
			Result(c, person, err)
		})
		personApi.POST("/add", func(c *gin.Context) {

			var person = new(models.UserInfo)
			if err := c.ShouldBindJSON(&person); err != nil {
				c.JSON(200, gin.H{
					"code": 400,
					"msg":  "add person error",
				})
				return
			}
			err := personService.Add(person)
			Result(c, nil, err)

		})
		personApi.POST("/delete", func(c *gin.Context) {

			name := c.Query("name")

			if len(name) == 0 {
				Result(c, nil, errors.New("name不能为空"))
				return
			}
			deleteCnt, err := personService.DeleteByName(name)
			Result(c, deleteCnt, err)
		})

	}

}
