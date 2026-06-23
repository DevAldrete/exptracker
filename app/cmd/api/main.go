package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/devaldrete/exptrack/app/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, "postgres://user:pass@localhost:5432/trackerdb")
	if err != nil {
		log.Fatalf("unable to create pool: %v", err)
	}

	defer pool.Close()

	queries := db.New(pool)

	route := gin.Default()

	route.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"health": "ok"})
	})

	err = route.Run(":8000")
	if err != nil {
		log.Fatal("server couldn't start!")
	}

	fmt.Println("Server running in http://localhost:8000")
}
