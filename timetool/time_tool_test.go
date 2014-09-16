package timetool

import (
	"fmt"
	"testing"
)

var p = fmt.Println

func Test_time(t *testing.T) {
	p(GetDefaultTime())
	p(GetYear())
	p(GetMonth())
	p(GetDayOfMonth())
	p(GetHour())
	p(GetMinute())
	p(GetSecond())
}
