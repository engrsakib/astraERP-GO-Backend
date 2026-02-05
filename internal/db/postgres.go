package db

import (
    "log"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

func NewPostgres(dsn string) *gorm.DB {
    
   
    db, err := gorm.Open(postgres.New(postgres.Config{
        DSN:                  dsn,
        PreferSimpleProtocol: true, 
    }), &gorm.Config{
        
        PrepareStmt: false,
        
        Logger: logger.Default.LogMode(logger.Info),
    })

    if err != nil {
        log.Fatalf("failed to connect postgres: %v", err)
    }
    return db
}