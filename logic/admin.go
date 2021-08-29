package logic

import (
	"qkteam-api/dao/mysql"
	"qkteam-api/models"
)

func GetAllPost() ([]*models.Register, error) {
	return mysql.GetAllPost()
}
