package postgres

import (
	"github.com/vishal1132/bootstrap-go-web/config"
	"github.com/vishal1132/bootstrap-go-web/utils"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(databaseConfig *config.DatabaseConfig) *gorm.DB {
	zap.L().Info("Connecting to postgres")
	db := utils.Must(gorm.Open(postgres.Open(databaseConfig.ConnectionString), &gorm.Config{}))
	zap.L().Info("Connected to postgres")
	return db
}
