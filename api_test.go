package webos

import (
	"crypto/tls"
	"log"
	"net"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

func TestLaunch(*testing.T) {
	dialer := websocket.Dialer{
		HandshakeTimeout: 10 * time.Second,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		NetDial: (&net.Dialer{
			Timeout: time.Second * 5,
		}).Dial,
	}

	tv, err := NewTV(&dialer, "192.168.1.130")
	if err != nil {
		log.Fatalf("could not dial: %v", err)
	}
	defer tv.Close()

	go tv.MessageHandler()

	if err = tv.AuthoriseClientKey("9bf6e831b96bcc80bf89464e1bb93ce9"); err != nil {
		log.Fatalf("could not authoise using client key: %v", err)
	}

	/*	tv.LaunchAppWithContentId("youtube.leanback.v4", "v=din4SiHB-nc")
		time.Sleep(10 * time.Second)
		tv.KeyOK()
	*/

	//tv.LaunchAppWithContentId("netflix", "m=https://www.netflix.com/watch/81298832?trackId=252072904")

	//tv.LaunchAppWithContentId("netflix", "m=https://www.netflix.com/browse/genre/83")
	tv.LaunchAppWithContentId("amazon", "")
}
