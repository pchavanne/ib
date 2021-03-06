package ib

import (
	"bufio"
	"bytes"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"

	//	"errors"
)

type Broker struct {
	ClientId  int64
	Conn      net.Conn
	Rid       int64
	OutStream *bytes.Buffer
	InStream  *bufio.Reader
}

func NextClientId() int64 {
	t := time.Now().UnixNano()
	rand.Seed(t)
	CLIENT_ID_INCR = int64(rand.Intn(9999))

	return CLIENT_ID_INCR
}

func (b *Broker) NextReqId() int64 {
	b.Rid += 1
	return b.Rid
}

func (b *Broker) Initialize() {
	b.ClientId = NextClientId()
	b.OutStream = bytes.NewBuffer(make([]byte, 0, 4096))
}

func (b *Broker) Connect(addr string, version int64) error {
	b.Initialize()

	conn, err := net.Dial("tcp", addr)

	if err != nil {
		return err
	}

	b.Conn = conn

	b.InStream = bufio.NewReader(b.Conn)

	err = b.ServerShake(version)

	return err
}

func (b *Broker) ServerShake(version int64) error {
	b.WriteInt(version)
	b.WriteInt(b.ClientId)

	_, err := b.SendRequest()

	return err
}

func (b *Broker) Disconnect() error {
	return b.Conn.Close()
}

func (b *Broker) SendRequest() (int, error) {
	b.WriteString(DELIM_STR)

	i, err := b.Conn.Write(b.OutStream.Bytes())

	b.OutStream.Reset()

	return i, err
}

func (b *Broker) SetServerLogLevel(i int64) {
	b.WriteInt(14)
	b.WriteInt(1)
	b.WriteInt(i)

	b.SendRequest()
}

func (b *Broker) WriteString(s string) (int, error) {
	return b.OutStream.WriteString(s + DELIM_STR)
}

func (b *Broker) WriteInt(i int64) (int, error) {
	return b.OutStream.WriteString(strconv.FormatInt(i, 10) + DELIM_STR)
}

func (b *Broker) WriteFloat(f float64) (int, error) {
	return b.OutStream.WriteString(strconv.FormatFloat(f, 'g', 10, 64) + DELIM_STR)
}

func (b *Broker) WriteBool(boo bool) (int, error) {
	if boo {
		return b.OutStream.WriteString("1" + DELIM_STR)
	}

	return b.OutStream.WriteString("0" + DELIM_STR)
}

func (b *Broker) ReadString() (string, error) {
	str, err := b.InStream.ReadString(DELIM_BYTE)

	if err != nil {
		return str, err
	}

	return strings.TrimRight(str, DELIM_STR), err
}

func (b *Broker) ReadInt() (int64, error) {
	str, err := b.ReadString()

	if err != nil {
		return 0, err
	}

	return strconv.ParseInt(strings.TrimRight(str, DELIM_STR), 10, 64)
}

func (b *Broker) ReadFloat() (float64, error) {
	str, err := b.ReadString()

	if err != nil {
		return 0, err
	}

	return strconv.ParseFloat(strings.TrimRight(str, DELIM_STR), 64)
}

func (b *Broker) ReadBool() (bool, error) {
	int, err := b.ReadInt()

	if err != nil {
		return false, err
	}

	if int != 0 {
		return true, err
	}

	return false, err
}
