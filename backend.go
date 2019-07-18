package main

import "github.com/nxshock/signaller/matrix"

type Backend interface {
	Register(username, password, device string) (token string, error *matrix.ApiError)
	Login(username, password, device string) (token string, err *matrix.ApiError)
	Logout(token string) *matrix.ApiError
	Sync(token string, request matrix.SyncRequest) (response *matrix.SyncReply, err *matrix.ApiError)
}
