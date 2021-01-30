package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/role", Get)

	router.GET("/role/:id", GetOne)

	router.POST("/role", Post)

	router.PUT("/role/:id", Put)

	router.DELETE("/role/:id", Delete)

	router.Run(":8080")
}

// 取得全部資料
func Get(c *gin.Context) {
	c.JSON(http.StatusOK, Data)
}

// 取得單一筆資料
func GetOne(c *gin.Context) {
	queryId := c.Param("id")
	for _, value := range Data {
		dataId := fmt.Sprint(value.ID)
		if dataId == queryId {
			c.JSON(http.StatusOK, value)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "找不到資料！"})
}

// 新增資料
func Post(c *gin.Context) {
	var newRecord *Role
	err := c.Bind(&newRecord)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	for _, value := range Data {
		if value.ID == newRecord.ID {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Id 重複！"})
			return
		}
	}

	Data = append(Data, *newRecord)
	c.JSON(http.StatusOK, gin.H{"message": "新增成功!"})
}

type RoleVM struct {
	ID      uint   `json:"id"`      // Key
	Name    string `json:"name"`    // 角色名稱
	Summary string `json:"summary"` // 介紹
}

// 更新資料, 更新角色名稱與介紹
func Put(c *gin.Context) {
	targetId := c.Param("id")
	var updateData *Role
	err := c.Bind(&updateData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	for idx, value := range Data {
		dataId := fmt.Sprint(value.ID)
		if dataId == targetId {
			Data[idx].Name = updateData.Name
			Data[idx].Summary = updateData.Summary

			c.JSON(http.StatusOK, gin.H{"message": "更新成功!"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "找不到資料！"})
}

// 刪除資料
func Delete(c *gin.Context) {
	targetId := c.Param("id")

	for idx, value := range Data {
		dataId := fmt.Sprint(value.ID)
		if dataId == targetId {
			Data = removeDataOfIndex(Data, idx)
			c.JSON(http.StatusOK, gin.H{"message": "刪除成功!"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "找不到資料！"})
}

func removeDataOfIndex(r []Role, index int) []Role {
	d1 := r[0:index]
	d2 := r[index+1:]
	d3 := append(d1, d2...)
	return d3
}
