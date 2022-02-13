package httpcores

import (
	"bufio"
	"errors"
	"io"
)

type HttpEntity struct {
	R       io.Reader
	Content []byte
}

func (h *HttpEntity) HandleResponse() ([]byte, error) {
	if h.R == nil {
		panic(errors.New("reader is nil"))
	}

	reader := bufio.NewReader(h.R)

	var buffer = make([]byte, 10240)

	for {
		length, err := reader.Read(buffer)

		h.Content = append(h.Content, buffer[:length]...)

		if err == io.EOF {
			break
		}
	}

	var e error
	if h.Content == nil {
		e = errors.New("this request content is nil")
	}

	return h.Content, e
}
