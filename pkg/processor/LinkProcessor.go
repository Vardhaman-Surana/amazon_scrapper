package processor

import (
	"fmt"
	"github.com/vds/amazon_scrapper/pkg/scrapper"
	"gopkg.in/gorp.v1"
	"time"
)

func LinkProcessor(links []string,db *gorp.DbMap){
	updatedTime:=time.Now().Unix()
	for _,link:=range links{
		_,err:=db.Exec("Update products set Updated=?,Status=? where URL=?",updatedTime,2,link)
		if err!=nil{
			fmt.Printf("error updating links fields and status: %v",err)
			return
		}
	}
	for _,link:=range links{
		title,company,price,status:= scrapper.ScrapeLink(link)
		if status==0{
			status=7	//to indicate successful scrapping
		}
		_,err:=db.Exec(`Update products set Title=?,Price=?,CompanyName=?,Status=? where URL=?`,title,price,company,status,link)
		if err!=nil{
			fmt.Printf("error updating the scrapped data: %v for link:%v",err,link)
		}
	}
}
