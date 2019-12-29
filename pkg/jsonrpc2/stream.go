package jsonrpc2

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	//headerLengthPrefix is the content length preamble
	headerLengthPrefix = "Content-Length"
)

//Stream is a transport abstraction free IO Stream for JSON-RPC 2.0
type Stream interface {
	//Read gets the next message from the stream
	Read() (data []byte, length int64, err error)
	Write(data []byte) (n int64, err error)
}

//NewStream communicates LSP over an io Reader and Writer
func NewStream(in io.Reader, out io.Writer) Stream {
	return &protocolStream{
		in:  bufio.NewReader(in),
		out: out,
	}
}

type protocolStream struct {
	in  *bufio.Reader
	out io.Writer
}

func (ps *protocolStream) Read() (data []byte, length int64, err error) {
	var total int64
	for {
		line, err := ps.in.ReadString('\n')
		total += int64(len(line))
		if err != nil {
			return data, total, fmt.Errorf("Error reading header %q", err)
		}
		line = strings.TrimSpace(line)
		if line == "" {
			break
		}
		fields := strings.Split(line, ":")
		if len(fields) != 2 {
			return data, total, fmt.Errorf("Invalid header line %q", line)
		}
		switch fields[0] {
		case headerLengthPrefix:
			value := strings.TrimSpace(fields[1])
			if length, err = strconv.ParseInt(value, 10, 32); err != nil {
				return data, total, fmt.Errorf("Error parsing Content-Length: %v", value)
			}
			if length < 1 {
				return data, total, fmt.Errorf("Invalid Content-Length: %v", length)
			}

		default: //skip everything else
		}
	}
	if length == 0 {
		return data, total, fmt.Errorf("No Content-Length header found")
	}
	data = make([]byte, length)
	if n, err := io.ReadFull(ps.in, data); err != nil {
		return data, total + int64(n), err
	}
	total += length
	return data, total, nil
}

func (ps *protocolStream) Write(data []byte) (int64, error) {
	n, err := fmt.Fprintf(ps.out, "%s: %v\r\n\r\n", headerLengthPrefix, len(data))
	total := int64(n)
	if err == nil {
		n, err = ps.out.Write(data)
		total += int64(n)
	}
	return total, err
}
