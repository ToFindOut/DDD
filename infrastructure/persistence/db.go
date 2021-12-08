package persistence

import (
	"food-app/domain/entity"
	"food-app/domain/repository"
	"github.com/go-sql-driver/mysql"
	"time"
	//"gorm.io/driver/mysql"
	"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Repositories struct {
	User repository.UserRepository
	Food repository.FoodRepository
	db   *gorm.DB
}

func NewRepositories(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) (*Repositories, error) {
	cfg := mysql.Config{
		User:                 DbUser,
		Passwd:               DbPassword,
		Net:                  "tcp",
		Addr:                 DbHost + ":" + DbPort,
		DBName:               DbName,
		Collation:            "utf8mb4_general_ci",
		Loc:                  time.FixedZone("Asia/Shanghai", 8*60*60),
		Timeout:              time.Second,
		ReadTimeout:          30 * time.Second,
		WriteTimeout:         30 * time.Second,
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	db, err := gorm.Open(Dbdriver, cfg.FormatDSN())

	if err != nil {
		return nil, err
	}
	db.LogMode(true)

	return &Repositories{
		User: NewUserRepository(db),
		Food: NewFoodRepository(db),
		db:   db,
	}, nil
}

//closes the  database connection
func (s *Repositories) Close() error {
	return s.db.Close()
}

//This migrate all tables
func (s *Repositories) Automigrate() error {
	return s.db.AutoMigrate(&entity.User{}, &entity.Food{}).Error
}
