package controllers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin/binding"
	"github.com/wilsontamarozzi/bemobi-hire-me/api/helpers"
	"github.com/wilsontamarozzi/bemobi-hire-me/api/models"
	"github.com/wilsontamarozzi/bemobi-hire-me/api/repositories"
	"log"
)

type URLController struct {
	Repository repositories.URLRepositoryInterface
}

func (controller URLController) GetAllRanking(c *gin.Context) {
	params := c.Request.URL.Query()
	urls := controller.Repository.GetAllRanking(params)

	c.JSON(200, urls)
}

func (controller URLController) GetByAlias(c *gin.Context) {
	urlAlias := c.Param("alias")
	url := controller.Repository.GetByAlias(urlAlias)

	if url.IsEmpty() {
		c.JSON(404, gin.H{"ERR_CODE": "002", "Description": "SHORTENED URL NOT FOUND"})
		return
	}

	if controller.Repository.UpdateView(&url) == nil {
		c.JSON(200, gin.H{"url": url})
	}
}

func (controller URLController) Create(c *gin.Context) {
	startTime := helpers.MakeTimestamp()
	var url models.URL
	// Faz um bind da entrada com o modelo e válida bad request (400)
	if err := c.BindJSON(&url); err != nil {
		log.Print(err)
		c.JSON(400, gin.H{"ERR_CODE": "005", "Description": "INVALID INPUT"})
		return
	}

	// Valida entidade inválida (422)
	if err := url.Validate(); err != nil {
		log.Print(err)
		c.JSON(422, gin.H{"ERR_CODE": "004", "Description": err})
		return
	}

	// Analisa se foi passado um alias preenchido para saber se já existe
	if !url.AliasIsEmpty() {
		// Faz uma busca no repositório pelo Alias
		urlTemp := controller.Repository.GetByAlias(url.Alias)
		// Verifica se o retorno foi vazio
		if !urlTemp.IsEmpty() {
			c.JSON(409, gin.H{"ERR_CODE": "001", "Description": "CUSTOM ALIAS ALREADY EXISTS"})
			return
		}
	}

	// Valida se deu erro para cadastrar (500)
	if err := controller.Repository.Create(&url); err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"ERR_CODE": "003", "Description": "THERE WAS AN ERROR TRYING TO REGISTER"})
		return
	}

	endTime := helpers.MakeTimestamp()
	c.JSON(201, gin.H{
		"alias": url.Alias,
		"url":   url.Address,
		"statistics": gin.H{
			"time_taken": helpers.ConvTimestampToMillisecond(startTime, endTime),
		},
	})
}