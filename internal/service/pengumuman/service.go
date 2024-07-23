package pengumuman

import (
	"log"
	"time"

	"github.com/onainadapdap1/golang_kantin/internal/api"
	"github.com/onainadapdap1/golang_kantin/internal/repository/pengumuman"
	"github.com/onainadapdap1/golang_kantin/models"
)

type PengumumanService interface {
	CreatePengumuman(pengumuman models.Pengumuman) (models.Pengumuman, error)
	GetAllPengumuman() ([]models.Pengumuman, error)
	UpdatePengumuman(pengumumanID uint, inputData api.UpdatePengumumanInput) (models.Pengumuman, error)
	DeletePengumumanByID(pengumumanID uint) error
	// GetPengumumanByID(pengumumanID uint) (models.Pengumuman, error)
}

type pengumumanService struct {
	repo pengumuman.PengumumanRepository
}

func NewPengumumanService(repo pengumuman.PengumumanRepository) PengumumanService {
	return &pengumumanService{repo}
}

func (s *pengumumanService) CreatePengumuman(pengumuman models.Pengumuman) (models.Pengumuman, error) {
	pengumuman.TanggalPembuatan = time.Now()
	if err := s.repo.CreatePengumuman(&pengumuman); err != nil {
		return models.Pengumuman{}, err
	}
	return pengumuman, nil
}

func (s *pengumumanService) GetAllPengumuman() ([]models.Pengumuman, error) {
	// var pengumumans []models.Pengumuman
	// var err error
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	//     defer wg.Done()
	//     pengumumans, err = s.repo.GetAllPengumuman()
	// }()
	// wg.Wait()

	// return pengumumans, err

	pengumumans, err := s.repo.GetAllPengumuman()
	if err != nil {
		return nil, err
	}

	return pengumumans, nil
}

func (s *pengumumanService) UpdatePengumuman(pengumumanID uint, inputData api.UpdatePengumumanInput) (models.Pengumuman, error) {
	log.Println("error 1 service")
	pengumuman, err := s.repo.GetPengumumanByID(pengumumanID) // 183
	if err != nil {
		return pengumuman, err
	}

	parsedTime, _ := time.Parse("2006-01-02", inputData.TanggalBerakhir)

	// layout := "2006-01-02 15:04:05"
	// timeToString := pengumuman.TanggalBerakhir.String()

	// if inputData.TanggalBerakhir == "" {
	// 	inputData.TanggalBerakhir = timeToString
	// }
	// if inputData.Deskripsi == "" {
	// 	inputData.Deskripsi = pengumuman.Deskripsi
	// }

	// updatePengumuman := models.Pengumuman {
	// 	TanggalBerakhir: parsedTime,
	// 	Deskripsi: inputData.Deskripsi,
	// }
	pengumuman.TanggalBerakhir = parsedTime
	pengumuman.Deskripsi = inputData.Deskripsi
	log.Println("error 2 service")
	updatedPengumuman, err := s.repo.UpdatePengumuman(pengumuman)
	if err != nil {
		return updatedPengumuman, nil
	}
	log.Println("error 3 service")

	return updatedPengumuman, nil
}

func (s *pengumumanService) DeletePengumumanByID(pengumumanID uint) error {
	err := s.repo.DeletePengumumanByID(pengumumanID)
	if err != nil {
		return err
	}

	return nil
}