package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mahfuzan/localswiftoms-v2/helpers"
	"github.com/mahfuzan/localswiftoms-v2/initializers"
	"github.com/mahfuzan/localswiftoms-v2/models"
)

func GetCompanyList(c *gin.Context) {
	var companyList []models.OmsCompany
	if err := initializers.DB.Find(&companyList).Error; err != nil {
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
		"message":        "Company data successfully returned",
		"result":         companyList,
	})
}

func GetCompanyById(c *gin.Context) {
	var company models.OmsCompany
	if err := initializers.DB.First(&company, c.Param("id")).Error; err != nil {
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
		"message":        "Company data successfully returned",
		"result":         company,
	})
}

func CreateCompany(c *gin.Context) {
	// get post data
	var json models.CompanyJson
	c.BindJSON(&json)
	company := models.OmsCompany{
		CompanyCode:      json.Company.CompanyCode,
		CompanyName:      json.Company.CompanyName,
		Email:            json.Company.Email,
		NoTelephone:      json.Company.NoTelephone,
		CompanyStreet:    json.Company.CompanyStreet,
		CompanyCountryId: json.Company.CompanyCountryId,
		CompanyRegion:    json.Company.CompanyRegion,
		CompanyCity:      json.Company.CompanyCity,
	}

	result := initializers.DB.Create(&company)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"transaction_id": helpers.GenerateUUID(),
			"success":        false,
			"message":        "Failed to save company data",
			"result":         []string{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"transaction_id": helpers.GenerateUUID(),
		"success":        true,
		"message":        "New company data has been saved",
		"result":         company,
	})
}

func UpdateCompanyById(c *gin.Context) {
	// get post data
	var json models.CompanyJson
	c.BindJSON(&json)
	companyPost := models.OmsCompany{
		CompanyCode:      json.Company.CompanyCode,
		CompanyName:      json.Company.CompanyName,
		Email:            json.Company.Email,
		NoTelephone:      json.Company.NoTelephone,
		CompanyStreet:    json.Company.CompanyStreet,
		CompanyCountryId: json.Company.CompanyCountryId,
		CompanyRegion:    json.Company.CompanyRegion,
		CompanyCity:      json.Company.CompanyCity,
	}

	// get data to update by id
	var company models.OmsCompany
	if err := initializers.DB.First(&company, json.Company.CompanyId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"transaction_id": helpers.GenerateUUID(),
			"success":        false,
			"message":        "Record not found",
			"result":         []string{},
		})
		return
	}

	if err := initializers.DB.Model(&company).Where("company_id = ?", company.CompanyId).Updates(companyPost).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"transaction_id": helpers.GenerateUUID(),
			"success":        false,
			"message":        "Failed to update company data",
			"result":         []string{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"transaction_id": helpers.GenerateUUID(),
		"success":        true,
		"message":        "Company data has been updated",
		"result":         company,
	})
}

func DeleteCompanyById(c *gin.Context) {
	if err := initializers.DB.Delete(&models.OmsCompany{}, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"transaction_id": helpers.GenerateUUID(),
			"success":        false,
			"message":        "Failed to delete company data",
			"result":         []string{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"transaction_id": helpers.GenerateUUID(),
		"success":        true,
		"message":        "Company data has been deleted",
		"result":         []string{},
	})
}
