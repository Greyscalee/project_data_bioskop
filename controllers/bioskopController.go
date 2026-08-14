package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"data-bioskop/config"

	"github.com/gin-gonic/gin"
)

const (
	errInvalidID = "Invalid ID"
	errNotFound  = "Data Not Found"
	errFetchData = "Failed to Fetch Data"
)

type Bioskop struct {
	BioskopID int    `json:"id"`
	Nama      string `json:"name"`
	Lokasi    string `json:"lokasi"`
	Rating    int    `json:"rating"`
}

func CreateBioskop(ctx *gin.Context) {
	var newBioskop Bioskop

	if err := ctx.ShouldBindJSON(&newBioskop); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	query := `INSERT INTO bioskop (name, lokasi, rating) VALUES ($1, $2, $3) RETURNING id`
	err := config.DB.QueryRow(query, newBioskop.Nama, newBioskop.Lokasi, newBioskop.Rating).Scan(&newBioskop.BioskopID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error_status":  "Failed to Create Data",
			"error_message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"bioskop": newBioskop,
	})
}

func GetBioskops(ctx *gin.Context) {
	query := `SELECT id, name, lokasi, rating FROM bioskop ORDER BY id`
	rows, err := config.DB.Query(query)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error_status":  errFetchData,
			"error_message": err.Error(),
		})
		return
	}
	defer rows.Close()

	bioskopDatas := []Bioskop{}
	for rows.Next() {
		var bioskop Bioskop
		if err := rows.Scan(&bioskop.BioskopID, &bioskop.Nama, &bioskop.Lokasi, &bioskop.Rating); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error_status":  "Failed to Scan Data",
				"error_message": err.Error(),
			})
			return
		}
		bioskopDatas = append(bioskopDatas, bioskop)
	}

	if err := rows.Err(); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error_status":  errFetchData,
			"error_message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"bioskops": bioskopDatas,
	})
}

func GetBioskop(ctx *gin.Context) {
	bioskopID, err := strconv.Atoi(ctx.Param("bioskopID"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  errInvalidID,
			"error_message": err.Error(),
		})
		return
	}

	var bioskopData Bioskop
	query := `SELECT id, name, lokasi, rating FROM bioskop WHERE id = $1`
	err = config.DB.QueryRow(query, bioskopID).Scan(&bioskopData.BioskopID, &bioskopData.Nama, &bioskopData.Lokasi, &bioskopData.Rating)
	if err == sql.ErrNoRows {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  errNotFound,
			"error_message": fmt.Sprintf("bioskop with id %v not found", bioskopID),
		})
		return
	} else if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error_status":  errFetchData,
			"error_message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"bioskop": bioskopData,
	})
}

func UpdateBioskop(ctx *gin.Context) {
	bioskopID, err := strconv.Atoi(ctx.Param("bioskopID"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  errInvalidID,
			"error_message": err.Error(),
		})
		return
	}

	var updatedBioskop Bioskop
	if err := ctx.ShouldBindJSON(&updatedBioskop); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	query := `UPDATE bioskop SET name = $1, lokasi = $2, rating = $3 WHERE id = $4`
	result, err := config.DB.Exec(query, updatedBioskop.Nama, updatedBioskop.Lokasi, updatedBioskop.Rating, bioskopID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error_status":  "Failed to Update Data",
			"error_message": err.Error(),
		})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  errNotFound,
			"error_message": fmt.Sprintf("bioskop with id %v not found", bioskopID),
		})
		return
	}

	updatedBioskop.BioskopID = bioskopID
	ctx.JSON(http.StatusOK, gin.H{
		"bioskop": updatedBioskop,
	})
}

func DeleteBioskop(ctx *gin.Context) {
	bioskopID, err := strconv.Atoi(ctx.Param("bioskopID"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  errInvalidID,
			"error_message": err.Error(),
		})
		return
	}

	query := `DELETE FROM bioskop WHERE id = $1`
	result, err := config.DB.Exec(query, bioskopID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error_status":  "Failed to Delete Data",
			"error_message": err.Error(),
		})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  errNotFound,
			"error_message": fmt.Sprintf("bioskop with id %v not found", bioskopID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("bioskop with id %v successfully deleted", bioskopID),
	})
}
