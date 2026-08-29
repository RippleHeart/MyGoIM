package WS

import (
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"log"
)

func (c *Client) ConsumePrivate() error {

	msgs, err := c.MQCh.Consume(c.Name, "", false, false, false, false, nil)
	if err != nil {
		log.Println(" c.MQCh.Consume:", err)
		return err
	}
	var msg amqp091.Delivery
	for {

		select {
		case msg = <-msgs:
			c.Send <- msg.Body
			fmt.Println("接收一次")
			msg.Ack(false) // 处理成功，确认删除
		case <-c.Close:
			fmt.Println("下线通知")
			return nil
		default:
			
			if msg.DeliveryTag == 0 {
				msg.Nack(false, true) // 处理失败，重新入队
			}
			continue
		}

	}
	return nil
}
