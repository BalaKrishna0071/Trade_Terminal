package loggers
import(
	"os"
	"path/filepath"
	"fmt"
	"log"
	"time"
	"runtime"
	"github.com/sirupsen/logrus"
	"github.com/lestrrat-go/file-rotatelogs"
)

var Logger *logrus.Logger

func init(){
	InitLogger()
}

func InitLogger(){

	Logger = logrus.New()

	file, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Logger.SetOutput(file)
	} else {
		Logger.SetOutput(os.Stdout)
	}

	Logger.SetReportCaller(true)

	logFile := "tmp/build-error.log"
	rotationPattern := logFile + ".%Y%m%d%H%M%S"
	writer, err := rotatelogs.New(
		rotationPattern,
		rotatelogs.WithLinkName(logFile),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(5*time.Minute),
	)

	if err != nil{
		log.Fatalf("Failed to create log file : %v", err)
	}

	Logger.SetOutput(writer)

	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		TimestampFormat: "2006-02-06 15:04:05",
		ForceColors: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
		
			_, filename := filepath.Split(f.File)
			line := f.Line
			return "", fmt.Sprintf("[%s:%d]", filename, line) 
		},
	})

	Logger.SetLevel(logrus.DebugLevel)
}