package pikaq

import (
	"testing"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

type PikaqTestSuite struct {
	suite.Suite
	MQ_URL             string
	TestConsumer       *Consumer
	TestMessageHandler MessageHandler
	TestConsumerTag    *ConsumerTag
	TestUUID           uuid.UUID
	TestSession        *Session
	TestQueue          Queue
	TestExchangeName   string
	TestExchangeType   string
	TestQueueName      string
	TestRoutingKey     string
	TestConsumerName   string
}

func (suite *PikaqTestSuite) SetupTest() {
	log.SetLevel(log.ErrorLevel)
	suite.MQ_URL = "amqp://localhost:5672"
	suite.TestExchangeName = "test"
	suite.TestExchangeType = "headers"
	suite.TestQueueName = "test"
	suite.TestRoutingKey = "test"
	suite.TestConsumerName = "test"
	suite.TestUUID = uuid.New()
	suite.TestConsumerTag = &ConsumerTag{name: suite.TestConsumerName, id: suite.TestUUID}
	suite.TestConsumer = &Consumer{tag: suite.TestConsumerTag}
	suite.TestMessageHandler = func(msgs Messages, done chan error) { done <- nil }
	suite.TestSession, _ = NewSession(suite.MQ_URL)
	_, suite.TestQueue, _ = DeclareQueue(suite.TestSession, suite.TestQueueName)
}

func TestPikaqTestSuite(t *testing.T) {
	suite.Run(t, new(PikaqTestSuite))
}
