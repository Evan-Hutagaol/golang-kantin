package menumakanan

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/onainadapdap1/golang_kantin/internal/api"
	"github.com/onainadapdap1/golang_kantin/internal/service/menumakanan"
	"github.com/onainadapdap1/golang_kantin/models"
)

type MenuMakananHandler interface {
	CreateMenuMakanan(c *gin.Context)
	GetAllMenuMakanan(c *gin.Context)
	DeleteMenuMakanan(c *gin.Context)
	UpdateMenuMakanan(c *gin.Context)
}

type menuMakananHandler struct {
	serv menumakanan.MenuMakananService
}

func NewMenuMakananHandler(serv menumakanan.MenuMakananService) MenuMakananHandler {
	return &menuMakananHandler{serv: serv}
}

func (h *menuMakananHandler) DeleteMenuMakanan(c *gin.Context) {
	menumakananID, _ := strconv.Atoi(c.Param("id"))

	err := h.serv.DeleteMenuMakanan(uint(menumakananID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete menu makanan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "menu makanan successfully deleted",
	})

}

func (h *menuMakananHandler) CreateMenuMakanan(c *gin.Context) {
	var menuMakanInput api.CreateMenuMakananInput
	if err := c.ShouldBind(&menuMakanInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error brewuu": err.Error()})
		return
	}

	parsedTime, err := time.Parse("2006-01-02", menuMakanInput.TanggalMakan)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	menuMakan := models.MenuMakanan{
		TanggalMakan: parsedTime,
		MenuPagi:     menuMakanInput.MenuPagi,
		MenuSiang:    menuMakanInput.MenuSiang,
		MenuMalam:    menuMakanInput.MenuMalam,
		Foto1:        menuMakanInput.Foto1,
		Foto2:        menuMakanInput.Foto2,
		Foto3:        menuMakanInput.Foto3,
	}

	if err := h.serv.CreateMenuMakanan(menuMakan); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create menu makanan"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": "Successfully to create menu makanan"})
}

func (h *menuMakananHandler) GetAllMenuMakanan(c *gin.Context) {
	menus, err := h.serv.GetAllMenuMakanan()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no data is found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    menus,
	})
}

func (h *menuMakananHandler) UpdateMenuMakanan(c *gin.Context) {
	menuMakananID, _ := strconv.Atoi(c.Param("id"))
	log.Println("handler 111")
	var inputMenuMakanan api.UpdateMenuMakananInput
	if err := c.ShouldBindJSON(&inputMenuMakanan); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	updatedMenuMakanan, err := h.serv.UpdateMenuMakanan(uint(menuMakananID), inputMenuMakanan)
	if err != nil {
		response := gin.H{
			"success": false,
			"message": "failed to update menu makanan",
			"error":   err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	log.Println("handler 2")
	successResp := gin.H{
		"success": true,
		"message": "successfully updated menu makanan",
		"data":    updatedMenuMakanan,
	}

	c.JSON(http.StatusOK, successResp)
}
