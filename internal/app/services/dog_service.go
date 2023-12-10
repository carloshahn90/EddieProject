package services

import (
	"github.com/carloshahn90/EddieProject/internal/app/models"
	"github.com/carloshahn90/EddieProject/internal/app/repositories"
)

type DogService struct {
	dogRepo *repositories.DogRepository
}

func NewDogService() *DogService {
	return &DogService{
		dogRepo: repositories.NewDogRepository(),
	}
}

func (as *DogService) GetDog(userId int) []models.Dog {

	dogs, _ := as.dogRepo.GetDogsByUserID(userId)
	return dogs

}
