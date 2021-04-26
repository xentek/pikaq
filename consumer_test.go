package pikaq

func (suite *PikaqTestSuite) TestConsumerType() {
	suite.IsType(&Consumer{}, suite.TestConsumer, "expects an instance of pikaq.Consumer")
}

func (suite *PikaqTestSuite) TestNewConsumer() {
	c, err := NewConsumer(suite.MQ_URL, suite.TestExchangeName, suite.TestExchangeType, suite.TestQueueName, suite.TestRoutingKey, suite.TestConsumerName, suite.TestConsumerPrefetch, suite.TestMessageHandler)
	suite.NotNil(c)
	if suite.NoError(err) {
		suite.IsType(&Consumer{}, c, "expects an instance of Consumer")
		suite.IsType(&ConsumerTag{}, c.tag, "expects an instance of ConsumerTag")
		suite.IsType(Queue{}, c.queue, "expects an instance of ConsumerTag")
		suite.IsType((chan error)(nil), c.done, "expects an error channel")
	}
}
