package mysql

import (
	"server/internal/models"
)

func SaveProvince(p models.Province) error {
	err := DB.Create(p).Error
	if err != nil {
		return err
	}
	return nil
}

func SaveCity(c models.City) error {
	err := DB.Create(c).Error
	if err != nil {
		return err
	}
	return nil
}

func SaveTown(t models.Town) error {
	err := DB.Create(t).Error
	if err != nil {
		return err
	}
	return nil
}
