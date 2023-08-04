package migration

import (
	"github.com/devops-codegpt/server/container"
	"github.com/devops-codegpt/server/models"
)

// CreateDatabase creates the tables used in this application
func CreateDatabase(c container.Container) error {
	if c.GetConfig().Database.Migration {
		db := c.GetRepository()

		err := db.AutoMigrate(
			&models.User{},
			&models.Role{},
			&models.Url{},
			&models.Conversation{},
			&models.Message{},
			&models.LLM{},
		)
		return err
	}
	return nil
}
