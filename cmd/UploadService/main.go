package main

import (
	"fmt"
	"github.com/vds/amazon_scrapper/pkg/database/mysql_orp"
	"github.com/vds/amazon_scrapper/pkg/queue"
	"github.com/vds/amazon_scrapper/pkg/server"
	"log"
	"os"
	"os/signal"
)

func main(){
	port:=os.Getenv("SERVICE_PORT")
	fmt.Printf("port is %v",port)

	log.Println("*********************************")
	log.Println("Connecting To Database")
	log.Println("*********************************")
	db,err:=mysql_orp.NewDBmap()
	if err!=nil{
		fmt.Printf("error initiating DB map: %v",err)
		return
	}
	log.Println("*********************************")
	log.Println("Database Connected")
	log.Println("*********************************")
	server,err:=server.NewServer(db)
	if err!=nil{
		fmt.Printf("error initializing server: %v",err)
	}
	router,err:=server.Start()
	if err!=nil{
		fmt.Printf("error starting a router: %v",err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func(){
		for sig := range c {
			fmt.Printf("interrupt signal %v, closing connection",sig)
			queue.Close()
			os.Exit(0)
		}
	}()

	err=router.Run(":"+port)
	if err!=nil{
		fmt.Printf("could not start server on given port: %v",err)
	}

}


