package database

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"local-media-gallery/models"
	"log"
	"time"
)

var DB *gorm.DB

var sqlDB *sql.DB

var allModels = []interface{}{
	&models.Directory{},
	&models.Image{},
	&models.Thumbnail{},
	&models.Video{},
}

func InIt() error {
	var err error

	DB, err = connectDatabase()
	if err != nil {
		return fmt.Errorf("error connecting to database: \n %v \n", err)
	}

	return nil
}

func connectDatabase() (*gorm.DB, error) {
	dbUsername := "root"
	dbPassword := ""
	dbName := "media_gallery"
	dbHost := "localhost"
	dbPort := "3306"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbPort, dbName)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: \n %v \n", err)
	}

	sqlDB, err = database.DB()
	if err != nil {
		return nil, fmt.Errorf("error getting sqlDB: \n %v \n", err)
	}

	maxIdleConns := 10
	maxOpenConns := 10
	connMaxLifetime := time.Minute * 3

	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetConnMaxLifetime(connMaxLifetime)

	return database, nil
}

func DropAllTables() error {
	err := DB.Migrator().DropTable(allModels...)
	if err != nil {
		return fmt.Errorf("error dropping table: \n %v \n", err)
	}

	return nil
}

func PerformMigration() error {
	DB.Config.DisableForeignKeyConstraintWhenMigrating = true
	err := DB.AutoMigrate(allModels...)
	if err != nil {
		return fmt.Errorf("failed initial AutoMigrate (table creation): %w", err)
	}

	DB.Config.DisableForeignKeyConstraintWhenMigrating = false
	err = DB.AutoMigrate(allModels...)
	if err != nil {
		return fmt.Errorf("failed final AutoMigrate (constraint creation): %w", err)
	}

	return nil
}

func CloseConnection() {
	if sqlDB != nil {
		err := sqlDB.Close()
		if err != nil {
			log.Printf("Error closing database connection: %v \n", err)
		} else {
			log.Println("Database connection closed successfully.")
		}
	}
}
