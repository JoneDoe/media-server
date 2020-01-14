package models

type RequestToken struct {
	Uuid string `uri:"uuid" binding:"required,uuid"`
}
