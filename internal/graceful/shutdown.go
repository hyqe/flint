package graceful

import (
	"os"
	"os/signal"
	"syscall"
)

type Stopper interface {
	Stop() error
}

type Starter interface {
	Start() error
}

type StartStoper interface {
	Starter
	Stopper
}

func Run(x StartStoper) error {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	var startErr chan error
	go func() {
		startErr <- x.Start()
	}()

	defer close(startErr)
	defer close(stop)

	select {
	case err := <-startErr:
		return err
	case <-stop:
		return x.Stop()
	}
}
