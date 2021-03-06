package main

import "fmt"

type channelWriter struct {
	Channel chan byte
}

func NewCannelWriter() *channelWriter {
	return &channelWriter{Channel: make(chan byte, 1024)}
}

func (c *channelWriter) Write(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	go func() {
		defer close(c.Channel)
		for _, b := range p {
			c.Channel <- b
		}
	}()

	return len(p), nil
}

func main() {
	cw := NewCannelWriter()
	go func() {
		fmt.Fprintf(cw,"Stream me")
	}()

	for c:=range cw.Channel{
		fmt.Printf("%c\n",c)
	}

}
