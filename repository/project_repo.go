package repository

import (
	"go_server/database"
	"go_server/models"
)

const recursiveQuery = `
WITH RECURSIVE project_tree AS (
    -- 앵커 멤버: 시작 지점, parent_id가 NULL인 최상위 프로젝트
    SELECT 
        id, 
        name, 
        parent_id,
        is_public,
        description,
        owner_id,
        1 AS level -- 계층 레벨을 나타내는 컬럼
    FROM 
        projects
    WHERE 
        parent_id IS NULL

    UNION ALL

    -- 재귀 멤버: parent_id를 사용하여 자신을 조인
    SELECT 
        p.id, 
        p.name, 
        p.parent_id,
        p.is_public,
        p.description,
        p.owner_id,
        pt.level + 1
    FROM 
        projects AS p
    JOIN 
        project_tree AS pt ON p.parent_id = pt.id
)
SELECT * FROM project_tree ORDER BY level;
`

func SelectProjectTree() ([]models.ProjectTree, error) {
	var projectsTree []models.ProjectTree

	result := database.DB.Raw(recursiveQuery).Scan(&projectsTree)

	if result.Error != nil {
		return nil, result.Error
	}

	return projectsTree, nil
}
