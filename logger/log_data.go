package logger

type LogData struct {
	datetime   string
	Level      string
	RemoteAddr string
	HttpStatus int
	Url        string
}
