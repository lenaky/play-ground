package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/pion/webrtc/v3"
)

func main() {
	// 1. 피어 연결 설정
	peerConnection, err := webrtc.NewPeerConnection(webrtc.Configuration{})
	if err != nil {
		log.Fatal(err)
	}

	// 2. 데이터 채널 이벤트 핸들러 설정
	peerConnection.OnDataChannel(func(d *webrtc.DataChannel) {
		d.OnOpen(func() {
			fmt.Println("Data channel opened")
		})

		d.OnMessage(func(msg webrtc.DataChannelMessage) {
			fmt.Printf("Received message: %s\n", string(msg.Data))
		})
	})

	// 3. 리모트 SDP 입력 받기 (Offer)
	fmt.Println("Paste the remote peer's SDP and press Enter:")
	remoteSDP := readSDP()

	// 4. 리모트 SDP 설정
	err = peerConnection.SetRemoteDescription(webrtc.SessionDescription{
		Type: webrtc.SDPTypeOffer,
		SDP:  remoteSDP,
	})
	if err != nil {
		log.Fatal(err)
	}

	// 5. 로컬 SDP 생성 (Answer)
	answer, err := peerConnection.CreateAnswer(nil)
	if err != nil {
		log.Fatal(err)
	}

	// 6. 로컬 SDP 설정
	err = peerConnection.SetLocalDescription(answer)
	if err != nil {
		log.Fatal(err)
	}

	// 7. 로컬 SDP 출력 (Offer를 생성한 피어로 전송해야 함)
	fmt.Println("=== Copy the below SDP to the remote peer ===")
	fmt.Println(answer.SDP)
	fmt.Println("============================================")

	// 8. 연결 대기
	select {}
}

// 사용자로부터 SDP를 입력 받는 함수
func _readSDP() string {
	scanner := bufio.NewScanner(os.Stdin)
	sdp := ""
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		sdp += line + "\n"
	}
	return sdp
}
