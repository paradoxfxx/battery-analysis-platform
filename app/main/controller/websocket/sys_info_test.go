package websocket

import (
	"battery-anlysis-platform/app/main/model"
	"github.com/gorilla/websocket"
	"testing"
)

func TestSysInfo(t *testing.T) {
	url := "ws://localhost:8080/websocket/v1/sys-info"
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatal(err)
	}
	var jd model.SysInfo
	if err = conn.ReadJSON(&jd); err != nil {
		t.Fatal(err)
	} else {
		t.Log(jd)
	}
}
