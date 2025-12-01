package traceback_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/ysuzuki19/collections-go/traceback"
)

type Suite struct {
	suite.Suite
	*require.Assertions
}

func (s *Suite) SetupTest() {
	s.Assertions = require.New(s.T())
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestNew() {
	// testdoc begin New
	err := traceback.New("something went wrong")
	fmt.Println(err.Error())
	// testdoc end
	s.Error(err)
}

func (s *Suite) TestErrorf() {
	// testdoc begin Errorf
	err := traceback.Errorf("failed to process item %d", 42)
	fmt.Println(err.Error())
	// testdoc end
	s.Error(err)
}

func (s *Suite) TestFrom() {
	// testdoc begin From
	someOperation := func() (any, error) { return nil, fmt.Errorf("some error") }
	result, err := someOperation()
	if err != nil {
		err = traceback.From(err)
	}
	// testdoc end
	s.Nil(result)
	s.Error(err)

	s.Nil(traceback.From(nil))
}

func (s *Suite) TestWrap() {
	// testdoc begin Wrap
	someOperation := func() (any, error) { return nil, fmt.Errorf("some error") }
	result, err := someOperation()
	if err != nil {
		err = traceback.Wrap(err, "failed to complete operation")
	}
	// testdoc end
	s.Nil(result)
	s.Error(err)

	s.Nil(traceback.Wrap(nil, "message"))
}

func (s *Suite) TestWrapf() {
	// testdoc begin Wrapf
	fetchUser := func(userID int) (any, error) {
		return nil, fmt.Errorf("user not found userID=%d", userID)
	}
	userID := 123
	result, err := fetchUser(userID)
	if err != nil {
		err = traceback.Wrapf(err, "failed to fetch user %d", userID)
	}
	// testdoc end
	s.Nil(result)
	s.Error(err)

	s.Nil(traceback.Wrapf(nil, "format %d", 1))
}

func (s *Suite) TestError() {
	// testdoc begin Error
	err := traceback.New("something went wrong")
	fmt.Println(err.Error())
	// testdoc end
	s.Error(err)
}

func (s *Suite) TestFrames() {
	// testdoc begin Error.Frames
	err := traceback.New("something went wrong")
	frames := err.Frames()
	fmt.Println(frames.String())
	// testdoc end
	s.Error(err)
}

func (s *Suite) TestString() {
	// testdoc begin Error.String
	err := traceback.New("something went wrong")
	fmt.Println(err.String())
	// testdoc end
	s.Error(err)
}
