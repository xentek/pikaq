package pikaq

func (suite *PikaqTestSuite) TestDeclareQueue() {
	s, q, err := DeclareQueue(suite.TestSession, suite.TestQueueName)
	suite.NotNil(s)
	suite.NotNil(q)
	if suite.NoError(err) {
		suite.IsType(Queue{}, q, "expects instance of Queue")
	}
}

func (suite *PikaqTestSuite) TestBindQueue() {
	s, err := BindQueue(suite.TestSession, suite.TestQueue, suite.TestRoutingKey, suite.TestExchangeName)
	suite.NotNil(s)
	suite.NoError(err)
}
