package main

import (
	"crypto/rand"
	"encoding/base64"
	"io"

	"github.com/nu7hatch/gouuid"
)

//Currently experimenting with ways to generate a session ID

func GenerateID() string {
	s, err := uuid.NewV4()
	if err != nil {
		return err.Error()
	}
	return s.String()
}

func sessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}
