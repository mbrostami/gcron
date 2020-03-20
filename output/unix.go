package output

import (
	"bytes"
	"encoding/gob"
	"gcron/cron"
	"net"
	"os"
)

// SendOverUNIX send data over unix socket
func SendOverUNIX(path string, task cron.Task) bool {
	binaryBuff := new(bytes.Buffer)
	gobobj := gob.NewEncoder(binaryBuff)
	gobobj.Encode(task)
	conn := connectUNIX(path)
	if conn != nil {
		go func(bytes []byte) {
			conn.Write(bytes)
			conn.Close()
		}(binaryBuff.Bytes())
	}
	return true
}

func connectUNIX(path string) *net.UnixConn {
	if path != "" {
		unixAddr, err := net.ResolveUnixAddr("unix", path)
		if err != nil {
			println("ResolveUNIXAddr failed:", err.Error())
			os.Exit(1)
		}
		conn, err := net.DialUnix("unix", nil, unixAddr)
		if err != nil {
			println("Dial UNIX failed:", err.Error())
			os.Exit(1)
		}
		return conn
	}
	return nil
}
