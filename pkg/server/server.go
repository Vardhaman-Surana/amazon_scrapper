package server

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gopkg.in/gorp.v1"
)

type Server struct{
	DB *gorp.DbMap
}


func NewServer(database *gorp.DbMap)(*Server,error){
	if database == nil {
		return nil, errors.New("server expects a valid database instance")
	}
	return &Server{DB:database}, nil
}

func(server *Server)Start()(*gin.Engine,error) {
	router,err:=NewRouter(server.DB)
	if err!=nil{
		return nil,err
	}
	r := router.Create()
	return r,nil
}
