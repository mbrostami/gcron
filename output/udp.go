package output

import (
	"bytes"
	"encoding/gob"
	"gcron/cron"
	"net"
	"os"
)

// SendOverUPD send data over udp
func SendOverUPD(host string, port string, task cron.Task) bool {
	binaryBuff := new(bytes.Buffer)
	gobobj := gob.NewEncoder(binaryBuff)
	gobobj.Encode(task)
	conn := connectUDP(host, port)
	if conn != nil {
		go func(bytes []byte) {
			conn.Write(bytes)
			conn.Close()
		}(binaryBuff.Bytes())
	}
	return true
}

func connectUDP(host string, port string) *net.UDPConn {
	if host != "" {
		servAddr := host + ":" + port
		udpAddr, err := net.ResolveUDPAddr("udp", servAddr)
		if err != nil {
			println("ResolveUDPAddr failed:", err.Error())
			os.Exit(1)
		}
		conn, err := net.DialUDP("udp", nil, udpAddr)
		if err != nil {
			println("Dial UDP failed:", err.Error())
			os.Exit(1)
		}
		return conn
	}
	return nil
}
