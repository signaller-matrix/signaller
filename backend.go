package main

import "github.com/nxshock/signaller/models"

type Backend interface {
	Register(username, password, device string) (token string, error *models.ApiError)
	Login(username, password, device string) (token string, err *models.ApiError)
	Logout(token string) *models.ApiError
	Sync(token string, request models.SyncRequest) (response *models.SyncReply, err *models.ApiError)
}
