package models

import "gorm.io/gorm"

type Club struct {
    gorm.Model
    Name  string `json:"name"`
    Owner string `json:"owner"`
}