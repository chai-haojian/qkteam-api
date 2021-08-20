package mysql

import (
	"qkteam-api/models"
)

func Submit(p *models.Register) (err error) {
	sqlStr := `insert into register(name, gender, phone, qq, specialty, self_introduction) values(?,?,?,?,?,?)`
	_, err = db.Exec(sqlStr, p.Name, p.Gender, p.Phone, p.QQ, p.Specialty, p.SelfIntroduction)
	if err != nil {
		return err
	}
	return
}
