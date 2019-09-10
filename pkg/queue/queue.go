package queue

import (
	"fmt"
	"github.com/vds/amazon_scrapper/pkg/getenv"
	"log"
	"github.com/streadway/amqp"
	"sync"
	"time"
)

type Queue struct{
	Name string
	Ch *amqp.Channel
	Connection *amqp.Connection
}
var(
	UploadingQueue Queue
	once     sync.Once
)

func InitializeQueue() *Queue {
	UploadingQueue.Name = "UploadingQueue"
	once.Do(func() {
		fmt.Println("*********************************")
		fmt.Println("Inside Once")
		fmt.Println("*********************************")
		conn := rConnect()
		if conn==nil{
			UploadingQueue.Connection=nil
		}else {
			fmt.Println("*********************************")
			fmt.Println("Connection Created")
			fmt.Println("*********************************")
			UploadingQueue.Connection = conn
			ch, err := conn.Channel()
			FailOnError(err, "Failed to open a channel")
			fmt.Println("*********************************")
			fmt.Println("Channel Created")
			fmt.Println("*********************************")
			_, err = ch.QueueDeclare(
				UploadingQueue.Name, // name
				true,
				false,
				false,
				false,
				nil,
			)
			FailOnError(err, "Failed to declare a queue")
			fmt.Println("rabbitmq connected")
			UploadingQueue.Ch = ch
		}
	})
	return &UploadingQueue
}
func rConnect() *amqp.Connection{
	fmt.Println("*********************************")
	fmt.Println(" Creating Connection")
	fmt.Println("*********************************")
	env,err:=getenv.GetRabbitEnv()
	if err!=nil{
		return nil
	}else {
		fmt.Printf("%v",env.URL)
		conn, err := amqp.Dial(env.URL)
		if err != nil {
			fmt.Printf("trying to reconnect")
			time.Sleep(5 * time.Second)
			return rConnect()
		}
		return conn
	}
}
func(q *Queue)PublishData(data []byte){
	err := q.Ch.Publish(
		"",     // exchange
		q.Name, 		// routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        data,
			DeliveryMode:amqp.Persistent,
		})
	log.Println(" Sent ")
	FailOnError(err, "Failed to publish a message")
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
func Close(){
	if UploadingQueue.Name!=""{
		UploadingQueue.Connection.Close()
		UploadingQueue.Ch.Close()
		fmt.Println("queue closed")
	}
}