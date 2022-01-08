package logger

type Output interface {
	Log(format string, parameters map[string]string)
}
