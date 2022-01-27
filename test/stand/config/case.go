package config

type Case struct {
	Protocol string
	Command  string
	Number   int
}

func NewCase(protocol string, command string, number int) Case {
	return Case{
		Protocol: protocol,
		Command:  command,
		Number:   number,
	}
}
