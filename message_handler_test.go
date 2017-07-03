package pikaq

func (suite *PikaqTestSuite) TestMessageHandlerType() {
	suite.IsType((MessageHandler)(nil), suite.TestMessageHandler, "expects an instance of pikaq.MessageHandler")
}

func (suite *PikaqTestSuite) TestMessageLoggerType() {
	suite.IsType((MessageHandler)(nil), MessageLogger, "expects an instance of pikaq.MessageHandler")
}
