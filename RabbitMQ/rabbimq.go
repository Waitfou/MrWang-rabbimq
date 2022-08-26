package RabbitMQ

//url格式是固定的，amqp://账号：密码@rabbitmq服务器地址:端口号/vhost
const MQURL = "amqp://imoocuser:imoocuser@127.0.0.1:5672/imooc"

type RabbitMQ struct {
	conn    *ampq.Connection
	channel *ampq.Channel
	//队列名
	QueueName string
	//交换机
	Exchange string
	//key
	Key string
	//连接信息
	Mqurl string
}

//创建结构体实例
func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	return &RabbitMQ{
		QueueName: queueName,
		Exchange:  exchange,
		Key:       key,
		Mqurl:     MQURL,
	}
}

//断开channel和connection，如果不断开的话会对我们的服务器资源产生浪费
func (r *RabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
}
