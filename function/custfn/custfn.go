package datetime

import (
	"time"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/support/log"

	"github.com/project-flogo/core/data/expression/function"
)

const DateTimeFormatDefault string = "2006-01-02T15:04:05-07:00"

type CustFn struct {
}

func init() {
	function.Register(&CustFn{})
}

func (s *CustFn) Name() string {
	return "custfn"
}

func (s *CustFn) GetCategory() string {
	return "custfns"
}

func (s *CustFn) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{}, false
}

func (s *CustFn) Eval(d ...interface{}) (interface{}, error) {
	log.RootLogger().Debugf("Returns the current datetime with timezone")
	var currentTime time.Time
	currentTime = time.Now().UTC()
	return currentTime.Format(DateTimeFormatDefault), nil
}
