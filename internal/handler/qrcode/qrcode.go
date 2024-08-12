package qrcode

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/onainadapdap1/golang_kantin/models"
	"github.com/onainadapdap1/golang_kantin/utils"
	"gorm.io/gorm"
)

type QrcodeHandler interface {
	GenerateQR(c *gin.Context)
	ScanQR(c *gin.Context)
}

type qrcodeHandler struct {
	database *gorm.DB
}

func NewQrcodehandler(db *gorm.DB) QrcodeHandler {
	return &qrcodeHandler{database: db}
}

func (db *qrcodeHandler) GenerateQR(c *gin.Context) {
	// role := c.MustGet("role").(string)
	// if role != "admin" {
	//     c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
	//     return
	// }
	type RequestData struct {
		AbsensiType string `json:"absensi_type" form:"absensi_type"`
	}
	// type AllergyReportInput struct {
	// 	// UserID    uint   `json:"user_id" form:"user_id"`
	// 	Allergies string `json:"allergies" form:"allergies"`
	// 	File      string `json:"file" form:"file"`
	// }

	var req RequestData
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	absensiType := req.AbsensiType
	fmt.Println("absensiType : ", absensiType)
	in := "masuk"
	isSame := absensiType == in
	fmt.Println("is true ? ", isSame)
	// absensiType := c.PostForm("absensi_type")
	if absensiType != "masuk" && absensiType != "keluar" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid absensi type"})
		return
	}
    // zones := []string{
    //     "Asia/Jakarta",
    //     "Asia/Bangkok",
    //     "Asia/Singapore",
    //     // Add more time zones to check
    // }

	// loc, err := time.LoadLocation("Asia/Jakarta")
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"success2": false, "error": err.Error()})
	// 	return
	// }
    // loc, err := time.LoadLocation("Asia/Jakarta")
    // if err != nil {
    //     // Attempt using an alternative time zone if the preferred one fails
    //     loc, err = time.LoadLocation("Asia/Bangkok")
    //     if err != nil {
    //         c.JSON(http.StatusInternalServerError, gin.H{"success2": false, "error": "Failed to load time zone: " + err.Error()})
    //         return
    //     }
    // }

    // now := time.Now().In(loc)
	// var validFrom, validTo time.Time

	// if absensiType == "masuk" {
	// 	validFrom = time.Date(now.Year(), now.Month(), now.Day(), 8, 0, 0, 0, loc)
	// 	validTo = time.Date(now.Year(), now.Month(), now.Day(), 8, 30, 0, 0, loc)
	// } else if absensiType == "keluar" {
	// 	validFrom = time.Date(now.Year(), now.Month(), now.Day(), 23, 0, 0, 0, loc)
	// 	validTo = time.Date(now.Year(), now.Month(), now.Day(), 23, 50, 0, 0, loc)
	// }
    
    loc, err := time.LoadLocation("Asia/Jakarta")
    if err != nil {
        loc = time.UTC // Fallback to UTC if the location fails to load
    }

    now := time.Now().In(loc)
    var validFrom, validTo time.Time

    if absensiType == "masuk" {
        validFrom = time.Date(now.Year(), now.Month(), now.Day(), 8, 0, 0, 0, loc)
        validTo = time.Date(now.Year(), now.Month(), now.Day(), 8, 30, 0, 0, loc)
    } else if absensiType == "keluar" {
        validFrom = time.Date(now.Year(), now.Month(), now.Day(), 23, 0, 0, 0, loc)
        validTo = time.Date(now.Year(), now.Month(), now.Day(), 23, 50, 0, 0, loc)
    }


	qrCode := absensiType + "-" + now.Format("20060102150405")
	filename := qrCode + ".png"

	err = utils.GenerateQRCode(qrCode, filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate QR code"})
		return
	}

	result := db.database.Create(&models.QRCode{
		Code:        qrCode,
		AbsensiType: absensiType,
		ValidFrom:   validFrom,
		ValidTo:     validTo,
		GeneratedAt: now,
	})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save QR code"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "QR code generated", "filename": filename})
}

func (db *qrcodeHandler) ScanQR(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(models.User)
	qrCode := c.PostForm("qr_code")
	// role := c.MustGet("role").(string)

	// if role != "karyawan" {
	//     c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
	//     return
	// }

	var qrCodeRecord models.QRCode

	result := db.database.Where("code = ?", qrCode).First(&qrCodeRecord)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "QR code not found"})
		return
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)
	validFrom := qrCodeRecord.ValidFrom.In(loc)
	validTo := qrCodeRecord.ValidTo.In(loc)
	fmt.Println("now date is : ", now)
	fmt.Println("valid from : ", validFrom)
	fmt.Println("valid to : ", validTo)

	if !(now.After(validFrom) && now.Before(validTo)) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "QR code is not valid at this time"})
		return
	}

	// userId := c.MustGet("user_id").(int)
	// userIDFloat, ok := c.MustGet("user_id").(float64)
	// if !ok {
	//     c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
	//     return
	// }
	// userID := int(userIDFloat)
	db.database.Create(&models.Absensi{
		UserID:      currentUser.ID,
		AbsensiType: qrCodeRecord.AbsensiType,
		CreatedAt:   now,
	})

	c.JSON(http.StatusOK, gin.H{"message": "Attendance recorded"})
}
