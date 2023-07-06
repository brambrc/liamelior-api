package Controller

import (
	"liamelior-api/Model"
	"net/http"

	"github.com/gin-gonic/gin"
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

func RegisterMember(context *gin.Context) {
	var input Model.Member

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	member := Model.Member{
		NamaLengkap:     input.NamaLengkap,
		NamaPanggilan:   input.NamaPanggilan,
		JenisKelamin:    input.JenisKelamin,
		Domisili:        input.Domisili,
		UsernameTwitter: input.UsernameTwitter,
		IDLine:          input.IDLine,
		Reason:          input.Reason,
		ActiveAgrrement: input.ActiveAgrrement,
		CashAgrrement:   input.CashAgrrement,
	}

	_, err := member.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Data kamu berhasil disubmit, tunggu kontak admin kami untuk validasi dan seleksi ya ! "})

}
