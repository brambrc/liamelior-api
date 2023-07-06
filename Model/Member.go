package Model

import (
	"liamelior-api/Database"

	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	NamaLengkap     string `json:"nama_lengkap" binding:"required" gorm:"unique"`
	NamaPanggilan   string `json:"nama_panggilan" binding:"required" gorm:"unique"`
	JenisKelamin    string `json:"jenis_kelamin" binding:"required"`
	Domisili        string `json:"domisili" binding:"required"`
	UsernameTwitter string `json:"username_twitter" binding:"required"`
	IDLine          string `json:"id_line" binding:"required"`
	Reason          string `json:"reason" binding:"required"`
	ActiveAgrrement bool   `json:"active_agrrement" binding:"required"`
	CashAgrrement   bool   `json:"cash_agrrement" binding:"required"`
	Approved        bool   `json:"is_approved"`
}

func (m *Member) Save() (*Member, error) {

	var err error
	err = Database.Database.Create(&m).Error

	if err != nil {
		return &Member{}, err
	}

	return m, nil
}
