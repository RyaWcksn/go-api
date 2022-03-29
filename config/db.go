package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
    dsn := GetEnv("DB_USER") + ":" + GetEnv("DB_PASSWORD") + "@tcp(" + GetEnv("DB_HOST") + ":" + GetEnv("DB_PORT") + ")/" + GetEnv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
    db , err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic(err)
    }
    return db
}
