package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vds/amazon_scrapper/pkg/models"
	"gopkg.in/gorp.v1"
	"net/http"
)
const(
	TitleNotFound = "Unable to find title for the product"
	CompanyNotFound ="Unable to find company for the product"
	PriceNotFound ="Unable to find Price for the product"
	TitleAndCompanyNotFound="Unable to find title,company for the product"
	TitleAndPriceNotFound="Unable to find title,price for the product"
	CompanyAndPriceNotFound="Unable to find company,price for the product"
	AllNotFound="Unable to find title company and price for the product"
)

type ScrappingStatusController struct{
	Db *gorp.DbMap
}

func NewScrappingStatusController(db *gorp.DbMap)*ScrappingStatusController{
	sc:=new(ScrappingStatusController)
	sc.Db=db
	return sc
}

func (sc *ScrappingStatusController)CheckStatus(c *gin.Context){

	count, err := sc.Db.SelectInt("select count(*) from products where Updated=(select max(Updated) from products) and Status=0")
	if err!=nil{
		fmt.Printf("error in selecting count for status:%v",err)
		JSONResponse:=models.NewJsonResponse("Internal Server error",err)
		c.JSON(http.StatusInternalServerError, JSONResponse)
		return
	}
	if count!=0{
		JSONResponse:=models.NewJsonResponse("Ready for scrapping",nil)
		c.JSON(http.StatusOK, JSONResponse)
		return
	}

	count, err = sc.Db.SelectInt("select count(*) from products where Status=2 and Updated=(select max(Updated) from products)")
	if err!=nil{
		fmt.Printf("error in selecting count for status:%v",err)
		JSONResponse:=models.NewJsonResponse("Internal Server error",err)
		c.JSON(http.StatusInternalServerError, JSONResponse)
		return
	}
	if count!=0{
		JSONResponse:=models.NewJsonResponse("Scrapping Running",nil)
		c.JSON(http.StatusOK, JSONResponse)
		return
	}

	LinksWithStatus:=[]struct{
		URL string
		Status int
	}{}

	_,err=sc.Db.Select(&LinksWithStatus,"select URL,Status from products where Updated=(select max(Updated) from products) and Status<>7")
	if err!=nil{
		fmt.Printf("error in selecting Links with Status for status:%v",err)
		JSONResponse:=models.NewJsonResponse("Internal Server error",err)
		c.JSON(http.StatusInternalServerError, JSONResponse)
		return
	}
	if len(LinksWithStatus)==0{
		JSONResponse:=models.NewJsonResponse("Scrapping Successful",nil)
		c.JSON(http.StatusOK, JSONResponse)
		return
	}
	type Status struct{
		URL string `json:"url"`
		Err string `json:"err"`
	}
	StatusResponse:=[]Status{}

	for _,item:=range LinksWithStatus{
		status:=getStringStatus(item.Status)
		StatusResponse=append(StatusResponse,Status{item.URL,status})
	}
	type Response struct{
		Msg string `json:"msg"`
		Links []Status `json:"links"`
	}
	response:=Response{"error In scrapping following product url's",StatusResponse}
	c.JSON(http.StatusOK, response)
}

func getStringStatus(status int)string{
	statusString:=""
	switch status {
	case 1:
		statusString=TitleNotFound
	case 3:
		statusString=CompanyNotFound
	case 5:
		statusString=PriceNotFound
	case 4:
		statusString=TitleAndCompanyNotFound
	case 6:
		statusString=TitleAndPriceNotFound
	case 8:
		statusString=CompanyAndPriceNotFound
	case 9:
		statusString=AllNotFound
	}
	return statusString
}
func (sc *ScrappingStatusController)GetArchived(c *gin.Context){
	var output []models.Product
	_,err:=sc.Db.Select(&output,"select * from products where deleted=2")
	if err!=nil{
		fmt.Printf("error in selecting archived links:%v",err)
		JSONResponse:=models.NewJsonResponse("Internal Server error",err)
		c.JSON(http.StatusInternalServerError, JSONResponse)
		return
	}
	c.JSON(http.StatusOK, output)
}
