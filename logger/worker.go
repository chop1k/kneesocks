package logger

type Worker interface {
	Work()
	Send(log Log)
}

type Log struct {
	Type       string
	Parameters map[string]string
}

type BaseWorker struct {
	channel chan Log
}

func NewWorker(length int) BaseWorker {
	return BaseWorker{
		channel: make(chan Log, length),
	}
}

func (w BaseWorker) Work() {
	for {
		err := w.receive()

		if err != nil {
			break
		}
	}
}

func (w BaseWorker) receive() error {
	log, ok := <-w.channel

	if !ok {

	}
}

func (w BaseWorker) Send(log Log) {
	w.channel <- log
}
