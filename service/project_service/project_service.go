package project_service

import (
	"fmt"
	"go_server/models"
	"go_server/repository"
)

func GetProjectsTree() ([]models.ProjectTree, error) {

	projectTrees, err := repository.SelectProjectTree()

	if err != nil {
		return nil, fmt.Errorf("프로젝트 조회 실패: %w", err)
	}

	return projectTrees, nil
}
