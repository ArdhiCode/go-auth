package migrations

import (
	"fmt"

	"github.com/ArdhiCode/go-auth/internal/entity"
	mylog "github.com/ArdhiCode/go-auth/internal/pkg/logger"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	fmt.Println(mylog.ColorizeInfo("\n==========Start Migrating=========="))

	mylog.Infof("Migrating Tables . . .")

	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error; err != nil {
		return err
	}

	if err := db.AutoMigrate(
		&entity.User{},
	); err != nil {
		return err
	}

	return nil
}
