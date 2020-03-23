package socket

import (
	transportHTTPGmf "github.com/ValentinEncinasRojas/ava/pkg/transport"

	errorGmf "github.com/ValentinEncinasRojas/ava/errors"
)

// Socket is our pseudo socket for transport.Socket
type Socket struct {
	id string
	// closed
	closed chan bool
	// remote addr
	remote string
	// local addr
	local string
	// send chan
	send chan *transportHTTPGmf.Message
	// recv chan
	recv chan *transportHTTPGmf.Message
}

// Accept passes a message to the socket which will be processed by the call to Recv
func (s *Socket) Accept(m *transportHTTPGmf.Message) errorGmf.ErrorGmfI {
	select {
	case s.recv <- m:
		return nil
	case <-s.closed:
		return nil
	}
}

// Process takes the next message off the send queue created by a call to Send
func (s *Socket) Process(m *transportHTTPGmf.Message) errorGmf.ErrorGmfI {
	select {
	case msg := <-s.send:
		*m = *msg
	case <-s.closed:
		// see if we need to drain
		select {
		case msg := <-s.send:
			*m = *msg
			return nil
		default:
			return nil
		}
	}
	return nil
}

func (s *Socket) Send(m *transportHTTPGmf.Message) errorGmf.ErrorGmfI {
	// make copy
	msg := &transportHTTPGmf.Message{
		Header: make(map[string]string),
		Body:   make([]byte, len(m.Body)),
	}

	// copy headers
	for k, v := range m.Header {
		msg.Header[k] = v
	}

	// copy body
	copy(msg.Body, m.Body)

	// send a message
	select {
	case s.send <- msg:
	case <-s.closed:
		return nil
	}

	return nil
}

func (s *Socket) Recv(m *transportHTTPGmf.Message) errorGmf.ErrorGmfI {
	// receive a message
	select {
	case msg := <-s.recv:
		// set message
		*m = *msg
	case <-s.closed:
		return nil
	}

	// return nil
	return nil
}

// Close closes the socket
func (s *Socket) Close() errorGmf.ErrorGmfI {
	select {
	case <-s.closed:
		// no op
	default:
		close(s.closed)
	}
	return nil
}

// New returns a new pseudo socket which can be used in the place of a transport socket.
// Messages are sent to the socket via Accept and receives from the socket via Process.
// SetLocal/SetRemote should be called before using the socket.
func New(id string) *Socket {
	return &Socket{
		id:     id,
		closed: make(chan bool),
		local:  "local",
		remote: "remote",
		send:   make(chan *transportHTTPGmf.Message, 128),
		recv:   make(chan *transportHTTPGmf.Message, 128),
	}
}
