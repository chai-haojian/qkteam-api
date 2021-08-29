package mysql

import (
	"database/sql"
	"go.uber.org/zap"
	"qkteam-api/models"
)

func GetAllPost() (PostList []*models.Register, err error) {
	sqlStr := `select name, gender, phone, qq, specialty, self_introduction from register`
	if err = db.Select(&PostList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no rows in db")
			err = nil
			return
		}
		zap.L().Error("db.Select(&PostList, sqlStr) failed", zap.Error(err))
		return nil, err
	}
	return
}
