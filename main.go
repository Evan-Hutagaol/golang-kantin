package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/onainadapdap1/golang_kantin/config"
	pengumumanHandler "github.com/onainadapdap1/golang_kantin/internal/handler/pengumuman"
	"github.com/onainadapdap1/golang_kantin/internal/middleware"
	pengumumanRepo "github.com/onainadapdap1/golang_kantin/internal/repository/pengumuman"
	"github.com/onainadapdap1/golang_kantin/internal/service/auth"
	pengumumanServ "github.com/onainadapdap1/golang_kantin/internal/service/pengumuman"

	feedbackHandler "github.com/onainadapdap1/golang_kantin/internal/handler/feedback"
	feedbackRepo "github.com/onainadapdap1/golang_kantin/internal/repository/feedback"
	feedbackServ "github.com/onainadapdap1/golang_kantin/internal/service/feedback"

	barangHandler "github.com/onainadapdap1/golang_kantin/internal/handler/barang"
	barangRepo "github.com/onainadapdap1/golang_kantin/internal/repository/barang"
	barangServ "github.com/onainadapdap1/golang_kantin/internal/service/barang"

	menuMakanHandler "github.com/onainadapdap1/golang_kantin/internal/handler/menumakanan"
	menuMakanRepo "github.com/onainadapdap1/golang_kantin/internal/repository/menumakanan"
	menuMakanServ "github.com/onainadapdap1/golang_kantin/internal/service/menumakanan"

	allergyReportHandler "github.com/onainadapdap1/golang_kantin/internal/handler/allergyreport"
	allergyReportRepo "github.com/onainadapdap1/golang_kantin/internal/repository/allergyreport"
	allergyReportServ "github.com/onainadapdap1/golang_kantin/internal/service/allergyreport"

	userHandler "github.com/onainadapdap1/golang_kantin/internal/handler/user"
	userRepo "github.com/onainadapdap1/golang_kantin/internal/repository/user"
	userServ "github.com/onainadapdap1/golang_kantin/internal/service/user"

	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	// config.LoadEnv()
	DB = config.ConnectToDB()
}

func main() {
	authService := auth.NewAuthService()
	userRepo := userRepo.NewUserRepository(DB)
	userServ := userServ.NewUserService(userRepo)
	userHandler := userHandler.NewUserHandler(userServ, authService)

	pengumumanRepo := pengumumanRepo.NewPengumumanRepository(DB)
	pengumumanServ := pengumumanServ.NewPengumumanService(pengumumanRepo)
	pengumumanHandler := pengumumanHandler.NewPengumumanHandler(pengumumanServ)

	feedbackRepo := feedbackRepo.NewFeedbackRepository(DB)
	feedbackServ := feedbackServ.NewFeedbackService(feedbackRepo)
	feedbackHandler := feedbackHandler.NewFeedbackHandler(feedbackServ)

	barangRepo := barangRepo.NewBarangRepository(DB)
	barangServ := barangServ.NewBarangService(barangRepo)
	barangHandler := barangHandler.NewBarangHandler(barangServ)

	menuMakanRepo := menuMakanRepo.NewMenuMakananRepo(DB)
	menuMakanServ := menuMakanServ.NewMenuMakananServ(menuMakanRepo)
	menuMakanHandler := menuMakanHandler.NewMenuMakananHandler(menuMakanServ)

	allergyReportRepo := allergyReportRepo.NewAllergyReportRepo(DB)
	allergyReportServ := allergyReportServ.NewAllergyReportServ(allergyReportRepo)
	allergyReportHandler := allergyReportHandler.NewAllergyReportHandler(allergyReportServ)

	router := gin.Default()
	api := router.Group("/api/v1")
	// Membuat pengumuman baru
	// newPengumuman := models.Pengumuman{
	//     TanggalBerakhir:  time.Now(),
	//     TanggalPembuatan: time.Now(),
	//     Deskripsi:        "Contoh pengumuman",
	//     Published:        true,
	// }

	// // Menyimpan pengumuman ke dalam database
	// DB.Create(&newPengumuman)

	// 1
	api.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hi sarah meilani butar butar!",
		})
	})
	api.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "testing endpoint!",
		})
	})
	api.POST("/login", userHandler.Login)
	// admin
	// 2
	api.POST("/pengumuman", middleware.AuthAdminMiddleware(authService, userServ), pengumumanHandler.CreatePengumuman)
	// 3
	api.GET("/pengumuman", middleware.AuthBothMiddleware(authService, userServ), pengumumanHandler.GetAllPengumuman)
	// new request update
	api.PUT("/pengumuman/:id", middleware.AuthAdminMiddleware(authService, userServ), pengumumanHandler.UpdatedPengumuman)
	api.DELETE("/pengumuman/:id", middleware.AuthAdminMiddleware(authService, userServ), pengumumanHandler.DeletePengumumanByID)
	// 4
	api.GET("/feedback", middleware.AuthAdminMiddleware(authService, userServ), feedbackHandler.GetAllFeedback)
	// 5
	api.POST("/menu-makanans", middleware.AuthAdminMiddleware(authService, userServ), menuMakanHandler.CreateMenuMakanan)
	api.PUT("/menu-makanans/:id", middleware.AuthAdminMiddleware(authService, userServ), menuMakanHandler.UpdateMenuMakanan)
	// 10
	api.GET("/menu-makanans", middleware.AuthAdminMiddleware(authService, userServ), menuMakanHandler.GetAllMenuMakanan)
	// new request delete
	api.DELETE("/menu-makanans/:id", middleware.AuthAdminMiddleware(authService, userServ), menuMakanHandler.DeleteMenuMakanan)
	// 11
	api.GET("/show-barangs/:id", middleware.AuthAdminMiddleware(authService, userServ), barangHandler.ShowBarang)
	// 12
	api.GET("/hide-barangs/:id", middleware.AuthAdminMiddleware(authService, userServ), barangHandler.HideBarang)

	// user
	// 6
	api.POST("/feedback", middleware.AuthUserMiddleware(authService, userServ), feedbackHandler.CreateFeedback)
	// 7 (INI BELUM DI DOC)
	api.GET("/my-feedback", middleware.AuthUserMiddleware(authService, userServ), feedbackHandler.GetAllMyFeedback)
	// 8
	api.POST("/barangs", middleware.AuthUserMiddleware(authService, userServ), barangHandler.CreateBarang)
	// 9
	api.GET("/barangs", middleware.AuthUserMiddleware(authService, userServ), barangHandler.GetPengumuman)
	// (INI BELUM DI DOC)
	api.POST("/allergy-reports", middleware.AuthUserMiddleware(authService, userServ), allergyReportHandler.CreateAllergyReport)
	// (INI BELUM DI DOC)
	api.GET("/allergy-reports", middleware.AuthUserMiddleware(authService, userServ), allergyReportHandler.GetAllAllergyReportByUserId)

	router.Run(envPortOr("8080"))
}

func envPortOr(port string) string {
	envPort := os.Getenv("PORT")
	if envPort != "" {
		return ":" + envPort
	}
	return ":" + port
}
