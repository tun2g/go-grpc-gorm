package database

import (
	"app/src/lib/logger"

	userModel "app/src/api/user/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/urfave/cli"
	"gorm.io/gorm"
)

var _logger = logger.NewLogger("database")

func GetMigrations() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		{
			ID: "20240526000003",
			Migrate: func(tx *gorm.DB) error {
				tx.AutoMigrate(&userModel.User{})
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				tx.Migrator().DropTable(&userModel.User{})
				return nil
			},
		},
	}
}

func Migration() cli.Command {
	cli := cli.Command{
		Name:  "migration:run",
		Usage: "Apply all migrations",
		Action: func(c *cli.Context) error {
			db := InitDB()
			m := gormigrate.New(db, gormigrate.DefaultOptions, GetMigrations())
			if err := m.Migrate(); err != nil {
				_logger.Errorf("could not migrate: %s", err)
				return nil
			}
			_logger.Info("Migrations applied successfully!")
			return nil
		},
	}
	return cli
}

func Rollback() cli.Command {
	cli := cli.Command{
		Name:  "migration:rollback",
		Usage: "Rollback the last migration",
		Action: func(c *cli.Context) error {
			db := InitDB()
			m := gormigrate.New(db, gormigrate.DefaultOptions, GetMigrations())
			if err := m.RollbackLast(); err != nil {
				_logger.Errorf("could not rollback: %s", err)
				return nil
			}
			_logger.Info("Migration rolled back successfully!")
			return nil
		},
	}
	return cli
}

func DropDatabase() cli.Command {
	return cli.Command{
		Name:  "schema:drop",
		Usage: "Drop the entire database",
		Action: func(c *cli.Context) error {
			db := InitDB()
			_logger.Warn("Dropping the entire database...")

			err := db.Migrator().DropTable(&userModel.User{})
			if err != nil {
				_logger.Errorf("could not drop database: %s", err)
				return err
			}

			_logger.Info("Database dropped successfully!")
			return nil
		},
	}
}
