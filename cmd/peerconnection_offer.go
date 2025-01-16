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

	// 2. 데이터 채널 생성 (옵션)
	dataChannel, err := peerConnection.CreateDataChannel("data", nil)
	if err != nil {
		log.Fatal(err)
	}

	dataChannel.OnOpen(func() {
		fmt.Println("Data channel opened")
	})

	dataChannel.OnMessage(func(msg webrtc.DataChannelMessage) {
		fmt.Printf("Received message: %s\n", string(msg.Data))
	})

	// 3. 로컬 SDP 생성 (Offer)
	offer, err := peerConnection.CreateOffer(nil)
	if err != nil {
		log.Fatal(err)
	}

	// 4. 로컬 SDP 설정
	err = peerConnection.SetLocalDescription(offer)
	if err != nil {
		log.Fatal(err)
	}

	// 5. 로컬 SDP 출력 (다른 피어로 전송해야 함)
	fmt.Println("=== Copy the below SDP to the remote peer ===")
	fmt.Println(offer.SDP)
	fmt.Println("============================================")

	// 6. 리모트 SDP 입력 받기
	fmt.Println("Paste the remote peer's SDP and press Enter:")
	remoteSDP := readSDP()

	// 7. 리모트 SDP 설정
	err = peerConnection.SetRemoteDescription(webrtc.SessionDescription{
		Type: webrtc.SDPTypeAnswer,
		SDP:  remoteSDP,
	})
	if err != nil {
		log.Fatal(err)
	}

	// 8. 연결 대기
	select {}
}

// 사용자로부터 SDP를 입력 받는 함수
func readSDP() string {
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
