package migrations

import (
	"books/internal/pkg/configs"
	"database/sql"
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"
)

const (
	cutSet       = "file://"
	databaseName = "postgres"
)

type Migrator struct {
	pgDBMigrate *migrate.Migrate
}

func ProvideMigrator(config configs.DatabaseConfig, pgDB *gorm.DB) (*Migrator, error) {
	dbConn, err := pgDB.DB()
	if err != nil {
		return nil, err
	}

	pgDBMigrate, err := initMigrate(dbConn, config.MigrationPath)
	if err != nil {
		return nil, err
	}

	return &Migrator{
		pgDBMigrate: pgDBMigrate,
	}, nil
}

func (m Migrator) RunMigrations() {
	m.RunMigrationsWith(m.pgDBMigrate, "Postgres Database")
}

func (m Migrator) RunMigrationsWith(migrateInstance *migrate.Migrate, dbName string) {
	if err := migrateInstance.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Printf("No change detected after running migrations for %s", dbName)
			return
		}
		log.Println(fmt.Sprintf("Migration failed for %s", dbName), err)
	}
	log.Printf("Migrations applied successsfully to %s", dbName)
}

func initMigrate(dbConn *sql.DB, directory string) (*migrate.Migrate, error) {
	driver, err := postgres.WithInstance(dbConn, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	sourcePath, err := getSourcePath(directory)
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(sourcePath, databaseName, driver)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func getSourcePath(directory string) (string, error) {
	directory = strings.TrimPrefix(directory, cutSet)

	absPath, err := filepath.Abs(directory)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s", cutSet, absPath), nil
}
