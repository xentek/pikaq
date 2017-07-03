package pikaq

func (suite *PikaqTestSuite) TestConsumerTagType() {
	suite.IsType(&ConsumerTag{}, suite.TestConsumerTag, "expects instance of ConsumerTag")
}

func (suite *PikaqTestSuite) TestConsumerTagValue() {
	suite.Equal(suite.TestConsumerTag.Tag(), suite.TestConsumerName+"_"+suite.TestUUID.String())
}
