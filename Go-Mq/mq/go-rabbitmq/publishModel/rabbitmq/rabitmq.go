package rabbitmq

//mq的封装
import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

//连接信息amqp://fws:giuest@127.0.0.1:5672/fws这个信息是固定不变的amqp://事固定参数后面两个是用户名密码ip地址端口号Virtual Host
const MQURL = "amqp://fws:guest@127.0.0.1:5672/fws"

//rabbitMQ结构体
type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	//队列名称
	QueueName string
	//交换机名称
	Exchange string
	//bind Key 名称
	Key string
	//连接信息
	Mqurl string
}

//创建结构体实例
func NewRabbitMQ(queueName string,exchange string,key string) *RabbitMQ {
	return &RabbitMQ { QueueName: queueName,Exchange: exchange,Key: key,Mqurl: MQURL }
}

//断开channel 和connection
func (r *RabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
}

//错误处理函数
func (r *RabbitMQ) failOnErr(err error,message string) {
	if err !=nil {
		log.Fatalf("%s:%s",message,err)
		panic(fmt.Sprintf("%s:%s", message,err))
	}
}

//订阅模式下创建Rabbitmq实例
func NewRabbitMQPubSub (exchangeName string) *RabbitMQ {
	//创建RabbitMQ实例
	rabbitmq :=NewRabbitMQ("",exchangeName,"")
	var err error
	//获取connection
	rabbitmq.conn,err=amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err,"failed to connect rabb"+"itmq!")
	//获取channel
	rabbitmq.channel,err=rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err,"failed to open a channel")
	return rabbitmq
}

//订阅模式生产
func (r *RabbitMQ) PublishPub(message string) {
	//1.尝试创建交换机
		err :=r.channel.ExchangeDeclare(
			r.Exchange,
			//交换机内型
			"fanout",
			//是否持久化
			true,
			//是否自动删除
			false,
			//true表示这个exchange不可以被client用来推送消息，仅用来进行exchange和exchange之间的绑定
			false,
			//是否阻塞处理，
			false,
			//额外的属性
			nil,
			)
		r.failOnErr(err,"failed to declare an excha"+ "nge")

		//发送消息到队列中
		r.channel.Publish(
			r.Exchange,
			"",
			// //如果为true，根据自身exchange类型和routekey规则无法找到符合条件的队列会把消息返还给发送者
			false,
			//如果为true，当exchange发送消息到队列后发现队列上没有消费者，则会把消息返还给发送者
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body: []byte(message),
			})
}

//pub/sub 模式下消费者
func (r *RabbitMQ) ReceiveSub() {
	//1.试探性创建交换机
	 err := r.channel.ExchangeDeclare(
		r.Exchange,
		//交换机内型
		"fanout",
		//是否持久化
		true,
		//是否自动删除
		false,
		//true表示这个exchange不可以被client用来推送消息，仅用来进行exchange和exchange之间的绑定
		false,
		//是否阻塞处理，
		false,
		//额外的属性
		nil,
	)
	r.failOnErr(err, "Failed to declare an exch"+
		"ange")

	//2.试探性创建队列，这里注意队列名称不要写
	q, err := r.channel.QueueDeclare(
		"", //随机生产队列名称
		false,
		false,
		true,
		false,
		nil,
	)
	r.failOnErr(err,"failed to declare a queue")

	//绑定队列到exchange中
	err = r.channel.QueueBind(
		q.Name,
		"",
		r.Exchange,
		false,
		nil,
		)
	//消费消息
	messages, err :=r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
		)

	forever := make(chan bool)
	//启用协程处理消息
	go func() {
		for d := range messages {
			//消息逻辑处理，可以自行设计逻辑
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}







