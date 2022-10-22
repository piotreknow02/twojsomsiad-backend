package controller

import (
	"net/http"
	"strconv"
	"twojsomsiad/model"
	"twojsomsiad/service"
	"twojsomsiad/utils"

	"github.com/gin-gonic/gin"
)

// GetAdverts godoc
// @Summary Get adverts
// @Description Get adverts
// @Tags advert
// @Accept json
// @Produce json
// @Param id path int true "advert ID"
// @Success 200 "Success"
// @Router /advert/ [get]
func (base *Controller) Adverts(c *gin.Context) {
	var args model.Args
	var err error

	args.Offset, err = strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		args.Offset = 0
	}
	args.Limit, err = strconv.Atoi(c.DefaultQuery("limit", "25"))
	if err != nil {
		args.Limit = 25
	}

	adverts, err := service.FindAdverts(base.DB, &args)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, &adverts)
}

// GetAdvert godoc
// @Summary Get advert
// @Description Get advert by id
// @Tags advert
// @Accept json
// @Produce json
// @Param id path int true "advert ID"
// @Success 200 "Success"
// @Router /advert/{id} [get]
func (base *Controller) Advert(c *gin.Context) {
	id := c.Param("id")
	advert, err := service.FindAdvertById(base.DB, id)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, &advert)
}

// Getadvert godoc
// @Summary Get advert
// @Description Get advert by ID
// @Tags advert
// @Accept json
// @Produce json
// @Param Advert body model.CreateAdvertDTO true "CreateAdvert"
// @Success 200 "Success"
// @Router /advert/ [post]
func (base *Controller) CreateAdvert(c *gin.Context) {
	var data model.CreateAdvertDTO
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.SendError(c, http.StatusBadRequest)
		return
	}
	sid, is := c.Get("id")
	id, ok := sid.(uint)
	if !is || !ok {
		utils.SendError(c, http.StatusNotAcceptable)
		return
	}
	advert, err := service.CreateAdvert(base.DB, id, &data)
	if err != nil {
		utils.SendError(c, http.StatusConflict)
		return
	}
	c.JSON(http.StatusOK, advert)
}

// Getadvert godoc
// @Summary Remove advert
// @Description Remove advert by ID
// @Tags advert
// @Accept json
// @Produce json
// @Param id path int true "advert ID"
// @Success 200 "Success"
// @Router /advert/{id} [delete]
func (base *Controller) RemoveAdvert(c *gin.Context) {
	id := c.Param("id")
	var advert model.Advert
	if err := base.DB.Delete(&advert, id).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, &advert)
}

// Apply godoc
// @Summary Apply
// @Description Apply for advert
// @Tags advert
// @Accept json
// @Produce json
// @Param id path int true "advert ID"
// @Success 200 "Success"
// @Router /advert/{id}/apply [get]
func (base *Controller) Apply(c *gin.Context) {
	sadvertId := c.Param("id")
	suserId, is := c.Get("id")
	if !is {
		utils.SendError(c, http.StatusNotAcceptable)
		return
	}
	advertId, err := strconv.Atoi(sadvertId)
	userId, ok := suserId.(uint)
	if err != nil || !ok {
		utils.SendError(c, http.StatusNotAcceptable)
		return
	}

	application, err := service.ApplyForEvent(base.DB, userId, uint(advertId))
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, application)
}

// GetApplication godoc
// @Summary Get aplication
// @Description Get application for advert
// @Tags advert
// @Accept json
// @Produce json
// @Param id path int true "advert ID"
// @Success 200 "Success"
// @Router /advert/{id}/application [get]
func (base *Controller) GetApplications(c *gin.Context) {
	sadvertId := c.Param("id")
	advertId, err := strconv.Atoi(sadvertId)
	if err != nil {
		utils.SendError(c, http.StatusNotAcceptable)
		return
	}

	adverts, err := service.GetApplicationsForAdvert(base.DB, uint(advertId))
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, adverts)
}

// VerifyApplication godoc
// @Summary Verify Application
// @Description Verify if user that applied to advert actually helped
// @Tags advert
// @Accept json
// @Produce json
// @Param id path int true "advert ID"
// @Param apid path int true "application ID"
// @Success 200 "Success"
// @Router /advert/{udvid}/application/{apid} [get]
func (base *Controller) VerifyApplication(c *gin.Context) {
	sadvertId := c.Param("id")
	sapplicationId := c.Param("apid")
	suserId, is := c.Get("id")
	if !is {
		utils.SendError(c, http.StatusNotAcceptable)
		return
	}
	advertId, err := strconv.Atoi(sadvertId)
	if err != nil {
		utils.SendError(c, http.StatusNotAcceptable)
		return
	}
	applicationId, err := strconv.Atoi(sapplicationId)
	userId, ok := suserId.(uint)
	if err != nil || !ok {
		utils.SendError(c, http.StatusNotAcceptable)
		return
	}

	appliction, err := service.ConfirmApplication(base.DB, uint(applicationId), userId, uint(advertId))
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, appliction)
}
