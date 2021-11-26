package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/Penjumlahan", jumlahOperation)
	router.POST("/Pengurangan", kurangOperation)
	router.POST("/Perkalian", kaliOperation)
	router.POST("/Pembagian", bagiOperation)

	router.Run()
}

type numberInput struct {
	Angka1 int
	Angka2 int
	Hasil  int
}

func jumlahOperation(c *gin.Context) {
	var numberInput numberInput

	err := c.ShouldBindJSON(&numberInput)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"Angka1": numberInput.Angka1,
		"Angka2": numberInput.Angka2,
		"Hasil":  numberInput.Angka1 + numberInput.Angka2,
	})
}

func kurangOperation(c *gin.Context) {
	var numberInput numberInput

	err := c.ShouldBindJSON(&numberInput)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"Angka1": numberInput.Angka1,
		"Angka2": numberInput.Angka2,
		"Hasil":  numberInput.Angka1 - numberInput.Angka2,
	})
}

func kaliOperation(c *gin.Context) {
	var numberInput numberInput

	err := c.ShouldBindJSON(&numberInput)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"Angka1": numberInput.Angka1,
		"Angka2": numberInput.Angka2,
		"Hasil":  numberInput.Angka1 * numberInput.Angka2,
	})
}

func bagiOperation(c *gin.Context) {
	var numberInput numberInput

	err := c.ShouldBindJSON(&numberInput)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"Angka1": numberInput.Angka1,
		"Angka2": numberInput.Angka2,
		"Hasil":  numberInput.Angka1 / numberInput.Angka2,
	})
}
