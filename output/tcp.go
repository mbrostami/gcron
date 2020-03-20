package output

import (
	"bytes"
	"encoding/gob"
	"gcron/cron"
	"net"
	"os"
)

// SendOverTCP send data over tcp
func SendOverTCP(host string, port string, task cron.Task) bool {
	binaryBuff := new(bytes.Buffer)
	gobobj := gob.NewEncoder(binaryBuff)
	gobobj.Encode(task)
	conn := connectTCP(host, port)
	if conn != nil {
		go func(bytes []byte) {
			conn.Write(bytes)
			conn.Close()
		}(binaryBuff.Bytes())
	}
	return true
}

func connectTCP(host string, port string) *net.TCPConn {
	if host != "" {
		servAddr := host + ":" + port
		tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
		if err != nil {
			println("ResolveTCPAddr failed:", err.Error())
			os.Exit(1)
		}
		conn, err := net.DialTCP("tcp", nil, tcpAddr)
		if err != nil {
			println("Dial failed:", err.Error())
			os.Exit(1)
		}
		return conn
	}
	return nil
}
