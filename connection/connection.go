package connection

import (
	"archie/utils"
	"archie/utils/configer"
	"fmt"
	"github.com/garyburd/redigo/redis"
	_ "github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetDB() (*gorm.DB, error) {
	dbConfig := configer.LoadDBConfig()

	return gorm.
		Open(
			"postgres",
			fmt.Sprintf("host=%s dbname=%s port=%s user=%s password=%s sslmode=disable",
				dbConfig.Host,
				dbConfig.DBName,
				dbConfig.Port,
				dbConfig.User,
				dbConfig.Password,
			),
		)
}

func GetRedis() (redis.Conn, error) {
	redisConfig := configer.LoadRedisConfig()

	return redis.Dial("tcp", fmt.Sprintf("%s:%s", redisConfig.Bind, redisConfig.Port))
}

func GetRedisConnMust(cb func(conn redis.Conn)) {
	conn, err := GetRedis()
	utils.Check(err)

	defer conn.Close()
	cb(conn)
}
