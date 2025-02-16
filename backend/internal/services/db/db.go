package db

import (
	"database/sql"
	"log"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
  db    *sql.DB
  once  sync.Once
)

func InitDB() (*sql.DB, error) {
  var err error

  err = godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  pg_addr := os.Getenv("PG_ADDR")
  if pg_addr == "" {
    log.Fatal("Unable to set PG_ADDR from .env")
  }

  once.Do(func() {
    db, err = sql.Open("postgres", pg_addr)
    if err != nil {
      log.Fatalf("Failed to open database: %v", err)
    }


    // Set connection pool params
    db.SetMaxOpenConns(30)
    db.SetMaxIdleConns(30)
    db.SetConnMaxLifetime(5 * time.Minute)

    // Test Connection
    if err = db.Ping(); err != nil {
      log.Fatalf("Failed to connect to database: %v", err)
    }

    log.Println("Database conencted!")
  })

  return db, err
}

func GetDB() *sql.DB {
  if db == nil {
    log.Fatal("Database not initialized. Call InitDB first.")
  }
  return db
}

func CloseDB() {
  if db != nil {
    db.Close()
  }
}
