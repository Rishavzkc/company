package controller

import (
	"company/model"
	"company/repository"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CompanyServer interface {
	CreateCompany(c *gin.Context)
	GetAllCompanies(c *gin.Context)
	GetCompany(c *gin.Context)
	UpdateCompany(c *gin.Context)
	DeleteCompany(c *gin.Context)
}

type CompanyController struct {
	database repository.CompanyRepository
}

func NewCompanyController(db repository.CompanyRepository) CompanyServer {
	return &CompanyController{
		database: db,
	}
}

func (cs *CompanyController) CreateCompany(c *gin.Context) {
	var company model.Company
	if err := c.BindJSON(&company); err != nil {
		companyErr := DBInsertionFailure
		companyErr.Error = fmt.Sprintf("failed to parse record from input: %v", err)
		c.JSON(http.StatusInternalServerError, companyErr)
		return
	}

	if company.ID == "" {
		companyErr := MissingMandatoryFields
		companyErr.Error = fmt.Sprintf("missing mandatory field: 'id'")
		return
	}

	if err := cs.database.CreateCompany(&company); err != nil {
		companyErr := DBInsertionFailure
		companyErr.Error = err.Error()
		c.JSON(http.StatusInternalServerError, companyErr)
		return
	}

	resp := SuccessStatus
	resp.Message = "Companies database creation successful"
	c.JSON(http.StatusOK, resp)
}

func (cs *CompanyController) GetAllCompanies(c *gin.Context) {
	companies, err := cs.database.GetAllCompanies(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, DBRetrievalFailure)
		return
	}
	c.JSON(http.StatusOK, companies)
}

func (cs *CompanyController) GetCompany(c *gin.Context) {
	id := c.Param("id")
	company, err := cs.database.GetCompanyByID(id)
	if err != nil {
		companyErr := DBRetrievalFailure
		companyErr.Error = err.Error()
		c.JSON(http.StatusInternalServerError, companyErr)
		return
	}
	c.JSON(http.StatusOK, company)
}

func (cs *CompanyController) UpdateCompany(c *gin.Context) {
	var company model.Company
	if err := c.BindJSON(&company); err != nil {
		companyErr := DBUpdateFailure
		companyErr.Error = fmt.Sprintf("failed to parse record from input: %v", err)
		c.JSON(http.StatusInternalServerError, companyErr)
		return
	}

	if company.ID == "" {
		companyErr := MissingMandatoryFields
		companyErr.Error = fmt.Sprintf("missing mandatory field: 'id'")
		return
	}

	if err := cs.database.UpdateCompany(&company); err != nil {
		companyErr := DBUpdateFailure
		companyErr.Error = err.Error()
		c.JSON(http.StatusInternalServerError, companyErr)
		return
	}

	resp := SuccessStatus
	resp.Message = fmt.Sprintf("Record with id '%s' updated successfully", company.ID)
	c.JSON(http.StatusOK, resp)
}

func (cs *CompanyController) DeleteCompany(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		companyErr := MissingMandatoryFields
		companyErr.Error = fmt.Sprintf("missing mandatory field: 'id'")
		return
	}

	err := cs.database.DeleteCompanyByID(id)
	if err != nil {
		companyErr := DBDeleteFailure
		companyErr.Error = err.Error()
		c.JSON(http.StatusInternalServerError, companyErr)
		return
	}

	successMessage := SuccessStatus
	successMessage.Message = fmt.Sprintf("Successfully deleted record with id '%s'", id)
	c.JSON(http.StatusOK, successMessage)
}
