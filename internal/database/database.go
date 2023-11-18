package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

type Service interface {
	Todo(a string, c *gin.Context) (int, map[string]interface{})
	Health(a string, c *gin.Context) (int, map[string]interface{})
}

type service struct {
	db *sql.DB
}

var dburl = os.Getenv("DB_URL")

func New() Service {
	db, err := sql.Open("sqlite3", dburl)
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}
	s := &service{db: db}
	return s
}

func (s *service) Todo(a string, c *gin.Context) (int, map[string]interface{}) {
	type User struct {
		id         string
		created_at string
		updated_at string
		content    string
	}
	var user User
	uid := c.Request.Header.Get("X-User-Id")
	data := make(map[string]interface{})
	data["action"] = a
	data["user_id"] = uid
	content, err := c.GetRawData()
	if err != nil {
		content = []byte("")
	}
	content_string := string(content)

	if a == "GET" {
		row := s.db.QueryRow("SELECT id, created_at, updated_at, content FROM todo_users WHERE id = ?", uid)
		if err == nil {
			if err = row.Scan(&user.id, &user.created_at, &user.updated_at, &user.content); err != nil {
				_, err := s.db.Exec("INSERT INTO todo_users (id, updated_at, created_at, content) VALUES (?, ?, ?, ?)", uid, nil, nil, content_string)
				if err == nil {
					row := s.db.QueryRow("SELECT id, created_at, updated_at, content FROM todo_users WHERE id = ?", uid)
					row.Scan(&user.id, &user.created_at, &user.updated_at, &user.content)
				}
			}
			data["data"] = user
			return 200, data
		}
	} else if a == "SET" {
		s.db.Exec("UPDATE todo_users SET content = ? WHERE id = ?", content_string, uid)
		row := s.db.QueryRow("SELECT id, created_at, updated_at, content FROM todo_users WHERE id = ?", uid)
		row.Scan(&user.id, &user.created_at, &user.updated_at, &user.content)
		data["data"] = user
		return 200, data
	} else if a == "DEL" {
		s.db.Exec("DELETE FROM todo_users WHERE id = ?", uid)
		data["data"] = user
		return 200, data
	}

	return 500, data
}

func (s *service) Health(a string, c *gin.Context) (int, map[string]interface{}) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.db.PingContext(ctx)
	if err != nil {
		log.Fatalf(fmt.Sprintf("db down: %v", err))
	}

	return 200, map[string]interface{}{
		"message": "It's healthy",
	}
}
