package cmd

import (
	"fmt"
	"os"

	"github.com/ArdhiCode/go-auth/db"
	"github.com/ArdhiCode/go-auth/db/migrations"
	mylog "github.com/ArdhiCode/go-auth/internal/pkg/logger"
	"gorm.io/gorm"
)

func Commands() error {
	db := db.New()
	if err := getParams(db); err != nil {
		return err
	}
	return nil
}

func getParams(db *gorm.DB) error {
	migrate := false

	for _, arg := range os.Args[1:] {
		if arg == "--migrate" {
			migrate = true
		}
	}

	if migrate {
		if err := migrations.Migrate(db); err != nil {
			return fmt.Errorf("migration failed: %w", err)
		}
		mylog.Infof("Migration completed successfully")
	}
	return nil
}
