package handler

import (
	"go_server/service/project_service"

	"github.com/gin-gonic/gin"
)

func ProjectTree(c *gin.Context) {

	// 프로젝트 서비스 호출
	result, err := project_service.GetProjectsTree()

	if err != nil {
		c.JSON(500, gin.H{"error": "Fail", "details": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": result})
}
