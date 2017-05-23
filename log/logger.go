package log


import (
	"fmt"
	"os"
	"github.com/kardianos/osext"
	"path/filepath"
	"gopkg.in/natefinch/lumberjack.v2"
	"github.com/go-kit/kit/log"
)

var Logger log.Logger

func init() {
	folderPath, err := osext.ExecutableFolder()
	if err != nil {
		fmt.Println("error when initializing logger:", err)
		os.Exit(1)
	}
	base := filepath.Base(folderPath)

	lumberOut := &lumberjack.Logger{
		Filename:   fmt.Sprint("./logs/", base, ".log"),
		MaxSize:    100, // megabytes
		MaxBackups: 10,
		MaxAge:     30, //days
	}

	Logger = log.NewLogfmtLogger(lumberOut)
	Logger = log.With(Logger, "ts", log.DefaultTimestamp)
}

func Log(keyValues... interface{})  {
	err :=Logger.Log(keyValues...)
	if err!=nil {
		fmt.Println(err)
	}
}