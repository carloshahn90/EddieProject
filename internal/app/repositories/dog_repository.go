package repositories

import (
	"errors"

	"github.com/carloshahn90/EddieProject/internal/app/models"
	"github.com/carloshahn90/EddieProject/pkg/db"
)

type DogRepository struct{}

func NewDogRepository() *DogRepository {
	return &DogRepository{}
}

// GetDogsByUserID retorna informações dos cachorros associados a um user_id
func (dr *DogRepository) GetDogsByUserID(userID int) ([]models.Dog, error) {
	query := "SELECT id, user_id, name, breed FROM dogs WHERE user_id = $1"
	rows, err := db.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dogs []models.Dog
	for rows.Next() {
		var dog models.Dog
		err := rows.Scan(&dog.Id, &dog.UserId, &dog.Name, &dog.Breed)
		if err != nil {
			return nil, err
		}
		dogs = append(dogs, dog)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(dogs) == 0 {
		return nil, errors.New("nenhum cachorro encontrado para o user_id fornecido")
	}

	return dogs, nil
}
