package main

import (
	"context"
	"net/http"
	"net/mail"
	"strconv"
	"time"

	"github.com/devaldrete/exptrack/app/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

const TIMEOUT = 30 * time.Second

// Expenses

func getExpenses(queries *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), TIMEOUT)
		defer cancel()

		var req db.GetExpensesParams

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		res, err := queries.GetExpenses(ctx, req)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, res)
	}
}

func getExpenseByID(queries *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), TIMEOUT)

		defer cancel()

		req := c.Param("id")

		id, err := strconv.Atoi(req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		res, err := queries.GetExpenseById(ctx, int64(id))
		if err != nil {
			c.JSON(http.StatusBadGateway, res)
		}

		c.JSON(http.StatusOK, res)
	}
}

func getExpensesByUserEmail(queries *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), TIMEOUT)
		defer cancel()

		req := c.Param("email")

		email, err := mail.ParseAddress(req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		res, err := queries.GetExpensesByUserEmail(ctx, email.String())
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, res)
	}
}

func getExpensesByUserID(queries *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), TIMEOUT)
		defer cancel()

		req := c.Param("id")

		idRaw, err := strconv.Atoi(req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		id := pgtype.Int8{
			Int64: int64(idRaw),
			Valid: true,
		}

		res, err := queries.GetExpensesByUserId(ctx, id)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, res)
	}
}

func createExpense(queries *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), TIMEOUT)
		defer cancel()

		var req db.CreateExpenseParams

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		res, err := queries.CreateExpense(ctx, req)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusCreated, res)
	}
}

// Users

func getUsers(queries *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), TIMEOUT)
		defer cancel()

		var req db.GetUsersParams

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		res, err := queries.GetUsers(ctx, req)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, res)
	}
}
