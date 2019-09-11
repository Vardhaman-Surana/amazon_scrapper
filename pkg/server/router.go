package server

import (
	"github.com/gin-gonic/gin"
	"github.com/vds/amazon_scrapper/pkg/controller"
	"gopkg.in/gorp.v1"
)
type Router struct{
	DB *gorp.DbMap
}

func NewRouter(db *gorp.DbMap)(*Router,error){
	router := new(Router)
	router.DB = db
	return router,nil
}
func (r *Router)Create() *gin.Engine {
	ginRouter:=gin.Default()
	fc:=controller.NewFileUploadController(r.DB)
	sc:=controller.NewScrappingStatusController(r.DB)

	ginRouter.POST("/uploadProductLinksFile",fc.UploadCSV)
	ginRouter.GET("/status",sc.CheckStatus)
	ginRouter.GET("/getArchived",sc.GetArchived)
	return ginRouter
}