package coinflip_test

import (
	. "chatterbet/coinflip"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"testing"
)

type NewCoinflipTestSuite struct {
	suite.Suite
	cf *Coinflip
}

func TestNewCoinflipTestSuite(t *testing.T) {
	suite.Run(t, new(NewCoinflipTestSuite))
}

func (suite *NewCoinflipTestSuite) SetupTest() {
	suite.cf = NewCoinflip()
}

func (suite *NewCoinflipTestSuite) TestShouldNotBeStarted() {
	assert.False(suite.T(), suite.cf.IsStarted())
}

func (suite *NewCoinflipTestSuite) TestCanBeStarted() {
	suite.cf.Start()
	assert.True(suite.T(), suite.cf.IsStarted())
}

func (suite *NewCoinflipTestSuite) TestShouldStartWhenSomeoneSendPlayMessage() {
	_ = suite.cf.Process(&Message{
		User:    "user#1",
		Content: "!play coinflip",
	})
	assert.True(suite.T(), suite.cf.IsStarted())
}

type StartedCoinflipTestSuite struct {
	suite.Suite
	cf *Coinflip
}

func TestStartedCoinflipTestSuite(t *testing.T) {
	suite.Run(t, new(StartedCoinflipTestSuite))
}

func (suite *StartedCoinflipTestSuite) SetupTest() {
	suite.cf = NewCoinflip()
	suite.cf.Start()
}

func (suite *StartedCoinflipTestSuite) TestShouldRegisterABetOnHead() {
	_ = suite.cf.Process(&Message{
		Content: "!bet head",
		User:    "user#12",
	})
	assert.Contains(suite.T(), suite.cf.AllBets(), Bet{
		Outcome: "head",
		User:    "user#12",
	})
}

func (suite *StartedCoinflipTestSuite) TestShouldRegisterABetOnTail() {
	_ = suite.cf.Process(&Message{
		Content: "!bet tail",
		User:    "user#12",
	})
	assert.Contains(suite.T(), suite.cf.AllBets(), Bet{
		Outcome: "tail",
		User:    "user#12",
	})
}

func (suite *StartedCoinflipTestSuite) TestShouldNotHaveAnOutcomeBeforeItsFlipped() {
	_, has := suite.cf.Outcome.Value()
	assert.False(suite.T(), has)
}

func (suite *StartedCoinflipTestSuite) TestShouldHaveAnOutcomeAfterItsFlipped() {
	suite.cf.Flip()
	v, has := suite.cf.Outcome.Value()
	assert.True(suite.T(), has)
	assert.Contains(suite.T(), []string{"head", "tail"}, v)
}

type CoinflipWithSomeBetsTestSuite struct {
	suite.Suite
	cf *Coinflip
}

func TestCoinflipWithSomeBetsTestSuite(t *testing.T) {
	suite.Run(t, new(CoinflipWithSomeBetsTestSuite))
}

func (suite *CoinflipWithSomeBetsTestSuite) SetupTest() {
	suite.cf = NewCoinflip()
	suite.cf.Start()
	suite.cf.Process(&Message{
		User:    "user#3",
		Content: "!bet head",
	})
	suite.cf.Process(&Message{
		User:    "user#4",
		Content: "!bet tail",
	})
}

//
//var _ = Describe("a coinflip game with bets", func() {
//	var coinflip *Coinflip
//
//	Context("that was flipped on head", func() {
//		BeforeEach(func() {
//			coinflip.Flip()
//		})
//
//		It("should have bets on head as wins", func() {
//			bs := coinflip.WonBets()
//			Expect(bs).To(HaveExactElements(
//				Bet{
//					Outcome: "head",
//					User:    "user#3",
//				},
//			))
//		})
//
//		It("should have bets on tail as loses", func() {
//			bs := coinflip.LostBets()
//			Expect(bs).To(HaveExactElements(
//				Bet{
//					Outcome: "tail",
//					User:    "user#4",
//				},
//			))
//		})
//	})
//})