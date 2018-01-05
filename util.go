package cuckoofilter

import (
	"encoding/binary"
	"bytes"
	"log"
)
func Fingerprint2Bytes(fingerprint uint16) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, fingerprint)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}