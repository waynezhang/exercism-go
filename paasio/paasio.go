package paasio

import (
	"io"
	"sync"
)

type counter struct {
	nops  int
	bytes int64
	mutex sync.Mutex
}

type readCounter struct {
	reader io.Reader
	counter
}

type writeCounter struct {
	writer io.Writer
	counter
}

type readWriteCounter struct {
	readCounter
	writeCounter
}

func (c *counter) process(n int, err error) {
	if err != nil {
		return
	}
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.nops++
	c.bytes += int64(n)
}

func (c *counter) count() (n int64, nops int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.bytes, c.nops
}

func (c *readCounter) Read(p []byte) (n int, err error) {
	n, err = c.reader.Read(p)
	c.process(n, err)
	return n, err
}

func (c *readCounter) ReadCount() (n int64, nops int) {
	return c.count()
}

func (c *writeCounter) Write(p []byte) (n int, err error) {
	n, err = c.writer.Write(p)
	c.process(n, err)
	return n, err
}

func (c *writeCounter) WriteCount() (n int64, nops int) {
	return c.count()
}

// NewReadCounter doc here
func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{reader: reader}
}

// NewWriteCounter doc here
func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{writer: writer}
}

// NewReadWriteCounter doc here
func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{readCounter: readCounter{reader: rw}, writeCounter: writeCounter{writer: rw}}
}
