package pingy

import (
	"net"
	"os"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

func Ping(host string, timeout time.Duration, address string) (*icmp.Message, time.Duration, error) {

	addr := &net.IPAddr{IP: net.ParseIP(address)}

	// Resolve the IP address of the host
	if len(address) <= 0 {
		addr, _ = net.ResolveIPAddr("ip", host)
	}

	// Create an ICMP message
	message := icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 0,
		Body: &icmp.Echo{
			ID:   os.Getpid() & 0xffff,
			Seq:  1,
			Data: []byte("hello"),
		},
	}

	// Serialize the ICMP message
	messageBytes, err := message.Marshal(nil)
	if err != nil {
		return nil, 0, err
	}

	// Open a raw socket
	conn, err := net.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		return nil, 0, err
	}
	defer conn.Close()

	// Send the ICMP message to the remote host
	start := time.Now()
	_, err = conn.WriteTo(messageBytes, addr)
	if err != nil {
		return nil, 0, err
	}

	// Wait for the response from the remote host
	buf := make([]byte, 1500)
	err = conn.SetReadDeadline(time.Now().Add(timeout))
	if err != nil {
		return nil, 0, err
	}
	n, _, err := conn.ReadFrom(buf)
	if err != nil {
		return nil, 0, err
	}
	elapsed := time.Since(start)

	// Parse the ICMP message from the response
	responseMessage, err := icmp.ParseMessage(ipv4.ICMPTypeEchoReply.Protocol(), buf[:n])
	if err != nil {
		return nil, 0, err
	}

	// Check if the response is an ICMP Echo Reply
	if responseMessage.Type == ipv4.ICMPTypeEchoReply {
		return responseMessage, elapsed, nil
	} else {
		return nil, 0, nil
	}
}
