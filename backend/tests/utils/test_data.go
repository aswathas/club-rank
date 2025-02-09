package utils

import (
	"backend/models"
)

func CreateTestDomain() *models.Domain {
    return &models.Domain{
        Name:   "Test Domain",
        ClubID: 1,
    }
}

func CreateTestClub() *models.Club {
    return &models.Club{
        Name: "Test Club",
    }
}