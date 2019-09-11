package main

import (
	"fmt"
	"github.com/vds/amazon_scrapper/pkg/database/mysql_orp"
	"log"
	"time"
)

func main(){
	log.Println("*********************************")
	log.Println("Initializing  Database")
	log.Println("*********************************")
	db,err:=mysql_orp.NewDBmap()
	if err!=nil{
		fmt.Printf("error initiating DB map: %v",err)
		return
	}
	log.Println("*********************************")
	log.Println("Database Connected")
	log.Println("*********************************")
	now:=time.Now().Unix()
	_,err=db.Exec("update products set deleted=2 where ?-updated > 60",now)
	if err!=nil{
		fmt.Printf("error updating the database:%v",err)
	}
}
