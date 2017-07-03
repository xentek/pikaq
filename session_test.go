package pikaq

func (suite *PikaqTestSuite) TestNewSessionType() {
	s, err := NewSession(suite.MQ_URL)
	suite.NotNil(s)
	if suite.NoError(err) {
		suite.IsType(&Session{}, s, "expects instance of Session")
	}
}
