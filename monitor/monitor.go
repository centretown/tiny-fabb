package monitor

import (
	"bufio"
	"io"
	"time"
)

// MonitorIO -
type MonitorIO interface {
	Open() (rdr io.Reader, wrt io.Writer, err error)
	Close() (err error)
}

// Monitor I/O via Bus
func Monitor(mio MonitorIO, bus *Bus) {

	rdr, wrt, err := mio.Open()
	if err != nil {
		bus.Err <- err
		return
	}

	defer mio.Close()

	var (
		bufReader = bufio.NewReaderSize(rdr, 256)
		bufWriter = bufio.NewWriterSize(wrt, 64)
		line      []byte
	)

	for {
		select {

		case <-bus.Done:
			return

		case request := <-bus.Out:
			_, err = bufWriter.WriteString(request + "\r")
			if err == nil {
				err = bufWriter.Flush()
			}
			if err != nil {
				bus.Err <- err
				// glog.Warningln(err)
			}

		default:
			_, err = bufReader.ReadByte()
			if err != nil {
				if err != io.EOF {
					bus.Err <- err
				}
				break
			}

			err = bufReader.UnreadByte()
			if err != nil {
				bus.Err <- err
				break
			}

			line, err = bufReader.ReadBytes('\n')
			if err != nil {
				if err != io.EOF {
					bus.Err <- err
				}
				break
			}

			bus.In <- string(line)
		}
		time.Sleep(time.Millisecond)
	}
}
