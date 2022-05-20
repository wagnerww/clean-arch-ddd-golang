package orm

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() *gorm.DB {
	dbURL := "postgres://postgre:postgre@localhost:5432/orders"

	//dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			ti, _ := time.LoadLocation("America/Sao_Paulo")
			return time.Now().In(ti)
		},
	})

	if err != nil {
		log.Println("falha na comunicação com a base de dados")
		panic("falha na comunicação com o DB")
	}

	log.Println("Conexão com a DB OK")
	return db
}
