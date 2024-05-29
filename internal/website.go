package internal

import (
	requests "github.com/TheTNB/panel/app/http/requests/website"
	"github.com/TheTNB/panel/app/models"
	"github.com/TheTNB/panel/types"
)

type Website interface {
	List(page int, limit int) (int64, []models.Website, error)
	Add(website types.Website) (models.Website, error)
	SaveConfig(config requests.SaveConfig) error
	Delete(id uint) error
	GetConfig(id uint) (types.WebsiteSetting, error)
	GetConfigByName(name string) (types.WebsiteSetting, error)
	GetIDByName(name string) (uint, error)
}
