package ip

import (
	"log"
	"net"
	"syscall"

	"github.com/davecgh/go-spew/spew"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

type ICMPHandleFunc func(*icmp.Message, string)

func (f ICMPHandleFunc) Handle(msg *icmp.Message, from string) {
	f(msg, from)
}

type ICMPHandler interface {
	Handle(*icmp.Message, string)
}

func ListenICMP(h ICMPHandler) error {
	// IPPROTO_RAW impolis enabled IP_HDRINCL
	// ICMP_FILTER
	//  Enable a special filter for raw sockets bound to the
	//  IPPROTO_ICMP protocol.  The value has a bit set for each
	//  ICMP message type which should be filtered out.  The
	//  default is to filter no ICMP messages.
	// From: https://man7.org/linux/man-pages/man7/raw.7.html
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_ICMP)
	if err != nil {
		return err
	}
	log.Printf("Socket created: %d", fd)

	buf := make([]byte, 128)
	for {
		n, from, err := syscall.Recvfrom(fd, buf, 0)
		if err != nil {
			log.Printf("recvfrom returns error: %v", err)
			break
		}

		iph, err := ipv4.ParseHeader(buf[:n])
		if err != nil {
			log.Printf("parse ipv4 header returns error: %v", err)
			break
		}

		log.Println("Received IPv4 header")
		log.Println(spew.Sdump(iph))

		if iph.Protocol == 1 {
			msg, err := icmp.ParseMessage(1 /* icmpv4 */, buf[20+len(iph.Options):n])
			if err != nil {
				log.Printf("icmp.Parse returns error: %v", err)
				break
			}

			// log.Printf("Received icmp from %v, data: %+v", sockAddress(from), msg)
			h.Handle(msg, sockAddress(from))
		}
	}

	return nil
}

func sockAddress(addr syscall.Sockaddr) string {
	var ret string
	switch addr := addr.(type) {
	case *syscall.SockaddrInet4:
		var ip net.IP = addr.Addr[0:]
		ret = ip.String()
	case *syscall.SockaddrInet6:
		var ip net.IP = addr.Addr[0:]
		ret = ip.String()
	case *syscall.SockaddrUnix:
		ret = addr.Name
	default:
		ret = "<unknown>"
	}
	return ret
}
