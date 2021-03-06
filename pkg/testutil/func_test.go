package testutil

import (
	"context"

	. "gopkg.in/check.v1"

	"github.com/kanisterio/kanister/pkg/param"
)

type FuncSuite struct {
}

var _ = Suite(&FuncSuite{})

func (s *FuncSuite) SetUpSuite(c *C) {
}

func (s *FuncSuite) TearDownSuite(c *C) {
}

func (s *FuncSuite) TestFailFunc(c *C) {
	ctx := context.Background()
	err := failFunc(ctx, param.TemplateParams{}, nil)
	c.Assert(err, NotNil)
}

func (s *FuncSuite) TestWaitFunc(c *C) {
	ctx := context.Background()
	done := make(chan bool)
	go func() {
		err := waitFunc(ctx, param.TemplateParams{}, nil)
		c.Assert(err, IsNil)
		close(done)
	}()
	select {
	case <-done:
		c.FailNow()
	default:
	}
	ReleaseWaitFunc()
	<-done
}

func (s *FuncSuite) TestArgsFunc(c *C) {
	ctx := context.Background()
	args := map[string]interface{}{"arg1": []string{"foo", "bar"}}
	go func() {
		err := argsFunc(ctx, param.TemplateParams{}, args)
		c.Assert(err, IsNil)
	}()
	c.Assert(ArgFuncArgs(), DeepEquals, args)
}
