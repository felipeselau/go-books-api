package controllers

import (
	"strconv"

	"github.com/felipeselau/go-books-api/database"
	"github.com/felipeselau/go-books-api/models"
	"github.com/gin-gonic/gin"
)

func ShowBook(c *gin.Context) {
	id := c.Param("id")

	newId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID precisa ser um inteiro"})
		return
	}

	db := database.GetDatabase()

	var book models.Book
	err = db.First(&book, newId).Error
	if err != nil {
		c.JSON(400, gin.H{"error": "Livro não encontrado " + err.Error()})
		return
	}

	c.JSON(200, book)
}

func ShowBooks(c *gin.Context) {
	db := database.GetDatabase()

	var books []models.Book
	err := db.Find(&books).Error
	if err != nil {
		c.JSON(400, gin.H{"error": "Erro ao buscar livros " + err.Error()})
		return
	}

	c.JSON(200, books)
}

func CreateBook(c *gin.Context) {
	db := database.GetDatabase()

	var book models.Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{"error": "Erro ao criar livro, JSON Inválido " + err.Error()})
		return
	}

	err = db.Create(&book).Error
	if err != nil {
		c.JSON(400, gin.H{"error": "Erro ao criar livro" + err.Error()})
		return
	}

	c.JSON(200, book)
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")

	newId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID precisa ser um inteiro"})
		return
	}

	db := database.GetDatabase()
	var book models.Book

	err = db.First(&book, newId).Error
	if err != nil {
		c.JSON(400, gin.H{"error": "Livro não encontrado " + err.Error()})
		return
	}

	err = c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{"error": "Erro ao atualizar livro, JSON Inválido " + err.Error()})
		return
	}

	err = db.Save(&book).Error
	if err != nil {
		c.JSON(400, gin.H{"error": "Erro ao atualizar livro" + err.Error()})
		return
	}

	c.JSON(200, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	newId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID precisa ser um inteiro"})
		return
	}

	db := database.GetDatabase()
	var book models.Book

	err = db.First(&book, newId).Error
	if err != nil {
		c.JSON(400, gin.H{"error": "Livro não encontrado " + err.Error()})
		return
	}

	err = db.Delete(&book).Error
	if err != nil {
		c.JSON(400, gin.H{"error": "Erro ao deletar livro" + err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Livro deletado com sucesso"})

}
