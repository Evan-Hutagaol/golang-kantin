package menumakanan

import (
	"log"
	"time"

	"github.com/onainadapdap1/golang_kantin/internal/api"
	"github.com/onainadapdap1/golang_kantin/internal/repository/menumakanan"
	"github.com/onainadapdap1/golang_kantin/models"
)

type MenuMakananService interface {
	CreateMenuMakanan(menu models.MenuMakanan) error
	GetAllMenuMakanan() ([]models.MenuMakanan, error)
	DeleteMenuMakanan(id uint) error
	UpdateMenuMakanan(id uint, inputData api.UpdateMenuMakananInput) (models.MenuMakanan, error)
}

type menuMakananService struct {
	repo menumakanan.MenuMakananRepository
}

func NewMenuMakananServ(repo menumakanan.MenuMakananRepository) MenuMakananService {
	return &menuMakananService{repo: repo}
}

func (s *menuMakananService) DeleteMenuMakanan(id uint) error {
	err := s.repo.DeleteMenuMakanan(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *menuMakananService) CreateMenuMakanan(menu models.MenuMakanan) error {
	return s.repo.CreateMenuMakanan(menu)
}

func (s *menuMakananService) GetAllMenuMakanan() ([]models.MenuMakanan, error) {
	menusMakanans, err := s.repo.GetAllMenuMakanan()
	if err != nil {
		return nil, err
	}
	return menusMakanans, nil
}

func (s *menuMakananService) UpdateMenuMakanan(id uint, inputData api.UpdateMenuMakananInput) (models.MenuMakanan, error) {
	log.Println("service 1 menu makanan")
	menuMakanan, err := s.repo.GetMenuMakananByID(id)
	if err != nil {
		return menuMakanan, err
	}
	log.Println("service 2")
	parsedTime, _ := time.Parse("2006-01-02", inputData.TanggalMakan)
	menuMakanan.TanggalMakan = parsedTime
	menuMakanan.MenuPagi = inputData.MenuPagi
	menuMakanan.MenuSiang = inputData.MenuSiang
	menuMakanan.MenuMalam = inputData.MenuMalam
	menuMakanan.Foto1 = inputData.Foto1
	menuMakanan.Foto2 = inputData.Foto2
	menuMakanan.Foto3 = inputData.Foto3
	updatedMenuMakanan, err := s.repo.UpdateMenuMakanan(menuMakanan)
	if err != nil {
		return updatedMenuMakanan, nil
	}
	log.Println("service 2")

	return updatedMenuMakanan, nil
}
