package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rahul.com/secure-exam/internal/model"
	"rahul.com/secure-exam/internal/repository"
)

var organisationRepo = repository.NewMongoRepository("secure-exam", "organisations")

func CreateOrganisation(ctx *gin.Context) {
	var organisation model.Organisation
	err := ctx.ShouldBindJSON(&organisation)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request", "error": err.Error()})
		return
	}

	if err := organisationRepo.CreateOrganisation(&organisation); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": organisation})
}

func GetOrganisations(ctx *gin.Context) {
	organisations, err := organisationRepo.GetAllOrganisations()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": organisations})
}
