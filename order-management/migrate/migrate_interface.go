package migrate

import "gorm.io/gorm"

type Migrateable interface {
	Migrate(DatabaseConnection *gorm.DB) error
}
