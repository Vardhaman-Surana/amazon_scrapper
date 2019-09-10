package router

import (
	"github.com/vds/amazon_scrapper/pkg/database"
	"github.com/gin-gonic/gin"
)
type Router struct{
	db database.Database
}

func NewRouter(db database.Database)(*Router,error){
	router := new(Router)
	router.db = db
	return router,nil
}
func (r *Router)Create() *gin.Engine {

}