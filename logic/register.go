package logic

import (
	"qkteam-api/dao/mysql"
	"qkteam-api/models"
)

func Submit(p *models.Register) error {
	return mysql.Submit(p)
}
