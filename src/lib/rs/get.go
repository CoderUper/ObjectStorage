package rs

import (
	"ObjectStorage/src/lib/objectstream"
	"fmt"
	"io"
)

type RSGetStream struct {
	*decoder
}

//get  every shards
func NewRSGetStream(locateInfo map[int]string, dataServers []string, hash string, size int64) (*RSGetStream, error) {
	if len(locateInfo)+len(dataServers) != ALL_SHARDS {
		return nil, fmt.Errorf("dataServers number is not match!")
	}
	readers := make([]io.Reader, ALL_SHARDS)
	for i := 0; i < ALL_SHARDS; i++ {
		server := locateInfo[i]
		if server == "" {
			locateInfo[i] = dataServers[0]
			dataServers = dataServers[1:]
			continue
		}
		//if server is not empty,read data from the server
		reader, err := objectstream.NewGetStream(locateInfo[i], fmt.Sprintf("%s.%d", hash, i))
		if err == nil {
			//get data shard i success
			readers[i] = reader
		}
	}
	writers := make([]io.Writer, ALL_SHARDS)
	perShards := (size + DATA_SHARDS - 1) / DATA_SHARDS
	var e error
	for i := range readers {
		if readers[i] == nil {
			writers[i], e = objectstream.NewTempPutStream(locateInfo[i], fmt.Sprintf("%s.%d", hash, i), perShards)
			if e != nil {
				return nil, e
			}

		}
	}
	dec := NewDecoder(readers, writers, size)
	return &RSGetStream{dec}, nil
}

func (s *RSGetStream) Close() {
	for i := range s.writers {
		if s.writers[i] != nil {
			s.writers[i].(*objectstream.TempPutStream).Commit(true)
		}
	}
}

func (s *RSGetStream) Seek(offset int64, whence int64) (int64, error) {
	if whence != io.SeekCurrent {
		panic("only support seekcurrent!")
	}
	if offset < 0 {
		panic("only support forward seek")
	}
	length := int64(BLOCK_SIZE)
	for offset != 0 {
		if offset > length {
			offset = length
		}
		buf := make([]byte, length)
		io.ReadFull(s, buf)
		offset -= length
	}
	return offset, nil
}
