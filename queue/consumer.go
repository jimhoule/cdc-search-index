package queue

import (
	"context"

	"github.com/IBM/sarama"
)

type ConsumerGroup = sarama.ConsumerGroup
type ConsumerGroupSession = sarama.ConsumerGroupSession

type ConsumerGroupHandler struct {
    ConsumerGroup ConsumerGroup
}

func NewConsumerGroupHandler(addresses []string, id string) (ConsumerGroupHandler, error) {
    consumerGroupHandler := ConsumerGroupHandler{}

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true


    consumerGroup, err := sarama.NewConsumerGroup(addresses, id, config)
    if err != nil {
        return consumerGroupHandler, err
    }
	defer consumerGroup.Close()

    consumerGroupHandler.ConsumerGroup = consumerGroup

    return consumerGroupHandler, nil
}

/**
 * NOTES:
 *  - Satisfies the sarama ConsumerGroupHandler interface
 *  - Runs at the beginning of a new session, before ConsumeClaim
 */
func (*ConsumerGroupHandler) Setup(ConsumerGroupSession) error   { 
	return nil
}

/**
 * NOTES:
 *  - Satisfies the sarama ConsumerGroupHandler interface
 *  - Runs at the end of a session, once all ConsumeClaim goroutines have exited but before the 
 *    offsets are committed for the very last time
 */
func (*ConsumerGroupHandler) Cleanup(ConsumerGroupSession) error {
	return nil
}

/**
 * NOTES:
 *  - Satisfies the sarama ConsumerGroupHandler interface
 *  - Starts a consumer loop of ConsumerGroupClaim's Messages()
 *  - Once the Messages() channel is closed, the Handler must finish its processing
 *    loop and exit
 */
func (cgh *ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
    for message := range claim.Messages() {
        // userID := string(msg.Key)
        // var notification models.Notification
        // err := json.Unmarshal(msg.Value, &notification)
        // if err != nil {
        //     log.Printf("failed to unmarshal notification: %v", err)
        //     continue
        // }
        // consumer.store.Add(userID, notification)

		// Marks message as consumed
        session.MarkMessage(message, "")
    }

    return nil
}

/**
 * NOTES:
 *  - Custom add
 *  - Starts listening for incoming messages
 */
func (cgh *ConsumerGroupHandler) Listen(topics []string) error {
    for {
        err := cgh.ConsumerGroup.Consume(context.Background(), topics, cgh)
        if err != nil {
            return err
        }
    }
}