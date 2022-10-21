package controller

import (
	"fmt"
	"net/http"
	"twojsomsiad/model"
	"twojsomsiad/service"
	"twojsomsiad/utils"

	"github.com/gin-gonic/gin"
)

func (base *Controller) Adverts(c *gin.Context) {
	var adverts []model.Advert
	c.BindJSON(&adverts)
	base.DB.Find(&adverts, "deleted_at IS NULL")
	c.JSON(http.StatusOK, &adverts)

}

func (base *Controller) Advert(c *gin.Context) {
	id := c.Param("id")
	fmt.Print("this is id: 				", id)

	advert, err := service.FindAdvertById(base.DB, id)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, &advert)
}

func (base *Controller) CreateAdvert(c *gin.Context) {
	var data model.CreateAdvertDTO
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.SendError(c, http.StatusBadRequest)
		return
	}
	err := service.CreateAdvert(base.DB, &data)
	if err != nil {
		utils.SendError(c, http.StatusConflict)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Advert created"})
}

func (base *Controller) RemoveAdvert(c *gin.Context) {
	id := c.Param("id")
	var advert model.Advert
	base.DB.Delete(&advert, id)
	c.JSON(http.StatusOK, &advert)
}
