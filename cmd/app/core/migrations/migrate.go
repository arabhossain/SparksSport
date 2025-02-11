package migrations

import (
	"fmt"
	"log"

	"SparksSport/cmd/app/core/config"
	"SparksSport/cmd/app/core/database"
	adminModel "SparksSport/pkg/admin/models"
)

// Migrate handles database migrations
func Migrate(cfg core_config.Config) error {
	fmt.Println("Starting database migration...")

	db, err := database.ConnectDB(cfg.DatabaseURL)
	if err != nil {
		log.Println("Database connection failed:", err)
		return err
	}

	// Run migrations
	if err := db.AutoMigrate(
		&adminModel.Admin{},
		// Add more models as needed
	); err != nil {
		log.Println("Migration failed:", err)
		return err
	}

	fmt.Println("Database migration completed successfully")
	return nil
}
