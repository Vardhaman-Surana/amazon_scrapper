package controller

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/vds/amazon_scrapper/pkg/models"
	"github.com/vds/amazon_scrapper/pkg/queue"
	"gopkg.in/gorp.v1"
	"io"
	"net/http"
	"time"
)

type FileUploadController struct{
	Db *gorp.DbMap
}

func NewFileUploadController(db *gorp.DbMap)*FileUploadController{
	fc:=new(FileUploadController)
	fc.Db=db
	return fc
}

func (fc *FileUploadController) UploadCSV(c *gin.Context){
	if err := c.Request.ParseMultipartForm(32 << 20); err != nil {
		fmt.Printf("error parsing multipartform: %v",err)
		JSONResponse:=models.NewJsonResponse("There was a problem with the JSON payload",err)
		c.JSON(http.StatusBadRequest,JSONResponse)
		return
	}
	defer c.Request.MultipartForm.RemoveAll()


	file, _, err := c.Request.FormFile("file")
	if err != nil {
		fmt.Printf("error while reading file: %v",err)
		JSONResponse:=models.NewJsonResponse("There was a problem while reading file",err)
		c.JSON(http.StatusBadRequest,JSONResponse)
		return
	}
	defer file.Close()
	reader := csv.NewReader(bufio.NewReader(file))
	first:=true
	LinkStruct:=struct{
		Links []string `json:"links"`
	}{}
	for{
		record, err := reader.Read()
		if first {
			first = false
			continue
		}
		if err == io.EOF {
			break
		}
		if record[0]==""{
			continue
		}
		LinkStruct.Links=appendIfMissing(LinkStruct.Links,record[0])
	}
	trans, err := fc.Db.Begin()
	defer func() {
		if err != nil {
			if err = trans.Rollback(); err != nil {
				fmt.Printf("error while rolling back transaction: %v", err)
			}
		} else {
			if err = trans.Commit(); err != nil {
				fmt.Printf("unable to commit link insert transaction: %v",err)
			}
		}
	}()
	createdTime:=time.Now().Unix()
	updatedTime:=time.Now().Unix()
	for _,url:=range LinkStruct.Links{

		product:=models.Product{
			BaseModel:    models.BaseModel{
				Created:createdTime,
				Updated:updatedTime,
			},
			URL:          url,
		}

		if err = trans.Insert(&product); err != nil {
			duplicateError, ok := err.(*mysql.MySQLError)
			if !ok {
				fmt.Printf("error while inserting link: %v", err)
				JSONResponse:=models.NewJsonResponse("There was a problem while inserting the links in the database",err)
				c.JSON(http.StatusInternalServerError,JSONResponse)
				return
			}
			if duplicateError.Number == 1062 {
				err:=trans.SelectOne(&product,fmt.Sprintf(`select * from products where URL="%s"`,url))
				if err!=nil{
					JSONResponse:=models.NewJsonResponse("Error While selecting data for duplicate link",err)
					c.JSON(http.StatusInternalServerError,JSONResponse)
					return
				}
				product.Updated=updatedTime
				product.Status=0
				_,err=trans.Update(&product)
				if err!=nil {
					JSONResponse:=models.NewJsonResponse("Error While updating data for duplicate link",err)
					c.JSON(http.StatusInternalServerError,JSONResponse)
					return
				}
				continue
			} else {
				fmt.Printf("error while inserting link: %v", err)
				JSONResponse:=models.NewJsonResponse("There was a problem while inserting the links in the database",err)
				c.JSON(http.StatusInternalServerError, JSONResponse)
				return
			}
		}
	}
	qu:=queue.InitializeQueue()
	if qu.Connection==nil{
		fmt.Print("error while getting environment variables")
		JSONResponse:=models.NewJsonResponse("error while getting environment variables",errors.New("server error"))
		c.JSON(http.StatusInternalServerError, JSONResponse)
	}
	fmt.Println("*********************************")
	fmt.Println("Initializing the Queue")
	fmt.Println("*********************************")
	JsonLinks:=new(bytes.Buffer)
	err=json.NewEncoder(JsonLinks).Encode(LinkStruct)
	if err!=nil{
		fmt.Printf("error while marshalling links to json: %v", err)
		JSONResponse:=models.NewJsonResponse("Error Converting links To json",err)
		c.JSON(http.StatusInternalServerError, JSONResponse)
	}
	qu.PublishData(JsonLinks.Bytes())
	JSONResponse:=models.NewJsonResponse("CSV uploaded successfully",nil)
	c.JSON(http.StatusOK, JSONResponse)
}


func appendIfMissing(slice []string, i string)[]string{
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}