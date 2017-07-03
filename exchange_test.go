package pikaq

func (suite *PikaqTestSuite) TestDeclareExchange() {
	var err error
	suite.NotNil(suite.TestSession)
	suite.TestSession, err = DeclareExchange(suite.TestSession, suite.TestExchangeName, suite.TestExchangeType)
	if suite.NoError(err) {
		suite.IsType(&Session{}, suite.TestSession, "expects an instance of Session")
	}
}
