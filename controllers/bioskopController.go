package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Bioskop struct {
	BioskopID string `json:"bioskop_id"`
	Nama string `json:"nama"`
	Lokasi string `json:"lokasi"`
	Rating float32 `json:"rating"`
}

var BioskopDatas = []Bioskop{}

func CreateBioskop(ctx *gin.Context) {
	var newBioskop Bioskop

	if err := ctx.ShouldBindJSON(&newBioskop); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newBioskop.BioskopID = fmt.Sprint("c%d", len(BioskopDatas)+1)
	BioskopDatas = append(BioskopDatas, newBioskop)

	ctx.JSON(http.StatusCreated, gin.H{
		"bioskop": newBioskop,
	})
}

func GetBioskop(ctx *gin.Context) {
	bioskopID := ctx.Param("bioskopID")
	condition := false
	var bioskopData Bioskop

	for i, Bioskop := range BioskopDatas {
		if bioskopID == Bioskop.BioskopID {
			condition = true
			bioskopData = BioskopDatas[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_atatus": "Data Not Found",
			"error_message": fmt.Sprintf("bioskop with id %v not found", bioskopID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"bioskop": bioskopData,
	})
}
