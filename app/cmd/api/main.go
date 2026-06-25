package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/devaldrete/exptrack/app/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func setupRouter(queries *db.Queries) *gin.Engine {
	router := gin.Default()

	// Expenses
	expgrp := router.Group("/expenses", CORSMiddleware())

	expgrp.GET("/", getExpenses(queries))
	expgrp.GET("/:id", getExpenseByID(queries))
	expgrp.POST("/", createExpense(queries))
	expgrp.PATCH("/:id", updateExpenseByID(queries))
	expgrp.DELETE("/:id", deleteExpenseByID(queries))

	// Users
	usrgrp := router.Group("/users", CORSMiddleware())

	usrgrp.GET("/", getUsers(queries))
	usrgrp.POST("/", createUser(queries))
	usrgrp.GET("/:id", getUserByID(queries))
	usrgrp.GET("/role/:id", getUsersByRole(queries))
	usrgrp.PATCH("/:id", updateUserByID(queries))
	usrgrp.DELETE("/:id", deleteUserByID(queries))

	// Admins
	// admgrp := router.Group("/admin", CORSMiddleware())

	// RBAC

	return router
}

func main() {
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("unable to create pool: %v", err)
	}

	defer pool.Close()

	queries := db.New(pool)

	r := setupRouter(queries)

	r.GET("/health", func(c *gin.Context) {
		conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		ctxWithTimeout, cancel := context.WithTimeout(ctx, 15*time.Second)
		defer cancel()
		if err = conn.Ping(ctxWithTimeout); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"health": "ok"})
	})

	err = r.Run(":8000")
	if err != nil {
		log.Fatal("server couldn't start!")
	}

	fmt.Println("Server running in http://localhost:8000")
}
