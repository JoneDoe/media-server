package models

type ReqFile struct {
	Uuid string `uri:"uuid" binding:"required,uuid"`
}
