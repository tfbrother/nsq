package protocol

import (
	"net"
	"runtime"
	"strings"

	"github.com/tfbrother/nsq/internal/lg"
)

type TCPHandler interface {
	Handle(net.Conn)
}

func TCPServer(listener net.Listener, handler TCPHandler, logf lg.AppLogFunc) {
	logf(lg.INFO, "TCP: listening on %s", listener.Addr())

	for {
		clientConn, err := listener.Accept()
		if err != nil {
			if nerr, ok := err.(net.Error); ok && nerr.Temporary() {
				logf(lg.WARN, "temporary Accept() failure - %s", err)
				runtime.Gosched()
				continue
			}
			// theres no direct way to detect this error because it is not exposed
			if !strings.Contains(err.Error(), "use of closed network connection") {
				logf(lg.ERROR, "listener.Accept() - %s", err)
			}
			break
		}

		// 启动一个线程, 交给 handler 处理, 这里使用的是 one connect per thread 模式
		// 因为golang的特性, one connect per thread 模式 实际上是  one connect per goroutine
		// 再加上golang将io操作都做了封装, 那么实际上这里是 one connect per event loop 模式
		go handler.Handle(clientConn)
	}

	logf(lg.INFO, "TCP: closing %s", listener.Addr())
}
