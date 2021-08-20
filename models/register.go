package models

type Register struct {
	Name             string `json:"name" db:"name"`
	Gender           int64  `json:"gender" db:"gender"`
	Phone            string `json:"phone" db:"phone"`
	QQ               string `json:"qq" db:"qq"`
	Specialty        string `json:"specialty" db:"specialty"`
	SelfIntroduction string `json:"self_introduction" db:"self_introduction"`
}
