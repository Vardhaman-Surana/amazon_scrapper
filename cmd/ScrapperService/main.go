package main

import (
	"encoding/json"
	"fmt"
	"github.com/vds/amazon_scrapper/pkg/database/mysql_orp"
	"github.com/vds/amazon_scrapper/pkg/processor"
	"github.com/vds/amazon_scrapper/pkg/queue"
	"log"
	"os"
	"os/signal"
)


func main(){
	log.Println("*********************************")
	log.Println("Initializing the Queue")
	log.Println("*********************************")
	qu:=queue.InitializeQueue()
	if qu.Connection==nil{
		log.Printf("error getting environment variables")
		return
	}
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
	msgs, err := qu.Ch.Consume(
		qu.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	LinkStruct:=struct{
		Links []string `json:"links"`
	}{}
	forever := make(chan bool)
	queue.FailOnError(err,"Failed to register a consumer")
	go func() {
		for d := range msgs {
			err=json.Unmarshal(d.Body,&LinkStruct)
			if err!=nil{
				log.Printf("err is %v",err)
			}else{
				log.Println("*********************************")
				log.Println("Got a Message")
				log.Println("*********************************")
				processor.LinkProcessor(LinkStruct.Links,db)
			}
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func(){
		for sig := range c {
			fmt.Printf("interrupt signal %v, closing connection",sig)
			queue.Close()
			fmt.Printf("queue closed")
			os.Exit(0)
		}
	}()
	fmt.Println("*********************************")
	fmt.Println("Waiting for getting a message in the queue")
	fmt.Println("*********************************")
	<-forever
}
