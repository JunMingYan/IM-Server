package mysql

import (
	"server/internal/models"
)

func SaveBrowserSetting(b models.BrowserSetting) error {
	err := DB.Save(b).Error
	if err != nil {
		return err
	}
	return nil
}
