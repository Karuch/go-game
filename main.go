package main

import (
	"main/internal/authserver"
	"main/internal/authserver/jwthandling/creator"
)

func main(){
    creator.Creator(2)
    authserver.AuthServer()
    //server.Server()
}