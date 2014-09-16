package logtool

//import (
//	"fmt"
//	"log"
//	"os"
//	"strings"
//)

//// Level
//type Level uint32

//const (
//	//DEBUG < INFO < WARN < ERROR
//	DEBUG Level = iota
//	INFO
//	WARN
//	ERROR
//)

////Appender
//type Appender string

//const (
//	CONSOLE     Appender = "console"
//	DailyFile   Appender = "dailyFile"
//	RollingFile Appender = "rollingFile"
//)

//var Logger map[Appender]interface{} = make(map[Appender]int)

//func InitConsoleLog(level Level) {
//	var temp_func = func() {
//		log.Println()
//	}()
//	Logger[CONSOLE] = temp_func
//}

////func InitDailyFileLog(level Level, fileName string) {
////	file, err := os.create(fileName)

////}

////func Debug(arg0 interface{}, args ...interface{}) {

////	str:= fmt.Sprint(arg0)+strings.Repeat(" %v", len(args)), args)
////}
