package main

import (
	//"net/http"
	//"strconv"
	//"crypto/md5"
	//"github.com/emicklei/go-restful"
	"crypto/rand"
	"encoding/base64"
	"io"

	"code.google.com/p/gorilla/sessions"
	"github.com/nu7hatch/gouuid"
)

var store sessions.Store

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

// func storeUser() {
// 	//store = sessions.NewFilesystemStore(path string, keyPairs ...[]byte)
// }
