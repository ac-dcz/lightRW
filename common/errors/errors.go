package errors

import (
	"errors"
	"fmt"
	"github.com/ac-dcz/lightRW/common/codes"
	gcodes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Message struct {
	Code codes.Code
	Msg  string
}

func New(code codes.Code, msg string) *Message {
	return &Message{Code: code, Msg: msg}
}

func (m *Message) Error() string {
	if m == nil {
		return "nil"
	}
	return fmt.Sprintf("code: %d msg: %s", m.Code, m.Msg)
}

func FromError(err error) *Message {
	if err == nil {
		return nil
	}
	m := &Message{}
	if errors.As(err, &m) {
		return m
	}
	if s, ok := status.FromError(err); ok {
		return &Message{
			Code: codes.Code(s.Code()),
			Msg:  s.Message(),
		}
	}
	return &Message{codes.UnKnown, err.Error()}
}

func ToStatus(err error) *status.Status {
	if err == nil {
		return nil
	}
	m := &Message{}
	if errors.As(err, &m) {
		return status.New(gcodes.Code(m.Code), m.Msg)
	}
	return status.Convert(err)
}
