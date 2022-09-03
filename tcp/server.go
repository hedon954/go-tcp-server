package tcp

import (
	"context"
	"go-tcp-server/interface/tcp"
	"go-tcp-server/lib/logger"
	"net"
)

type Config struct {
	Address string
}

func ListenAndServeWithSignal(cfg *Config, handler tcp.Handler) error {
	listener, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		return err
	}
	logger.Info("start to listen")

	closeChan := make(chan struct{})
	ListenAndServe(listener, handler, closeChan)

	return nil
}

func ListenAndServe(listener net.Listener, handler tcp.Handler, closeChan <-chan struct{}) {

	ctx := context.Background()
	for {
		conn, err := listener.Accept()
		if err != nil {
			break
		}
		logger.Info("accepted connection")
		go func() {
			handler.Handle(ctx, conn)
		}()
	}
}
