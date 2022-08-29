package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mahfuzan/localswiftoms-v2/helpers"
	"github.com/mahfuzan/localswiftoms-v2/initializers"
	"github.com/mahfuzan/localswiftoms-v2/models"
)

func GetLocationList(c *gin.Context) {
	var locationList []models.OmsLocation
	if err := initializers.DB.Table("oms_location").
		Select("oms_location.*, oms_company.company_name").
		Joins("LEFT JOIN oms_company ON oms_company.company_id = oms_location.company_id").
		Scan(&locationList).
		Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"transaction_id": helpers.GenerateUUID(),
			"success":        false,
			"message":        "Record not found",
			"result":         []string{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"transaction_id": helpers.GenerateUUID(),
		"success":        true,
		"message":        "Location data successfully returned",
		"result":         locationList,
	})
}

func GetLocationById(c *gin.Context) {
	var location models.OmsLocation
	if err := initializers.DB.Table("oms_location").
		Select("oms_location.*, oms_company.company_name").
		Joins("LEFT JOIN oms_company ON oms_company.company_id = oms_location.company_id").
		Where("loc_id = ?", c.Param("id")).
		Scan(&location).
		Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"transaction_id": helpers.GenerateUUID(),
			"success":        false,
			"message":        "Record not found",
			"result":         []string{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"transaction_id": helpers.GenerateUUID(),
		"success":        true,
		"message":        "Location data successfully returned",
		"result":         location,
	})
}
