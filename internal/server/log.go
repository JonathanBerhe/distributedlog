package server

import (
	"fmt"
	"sync"
)

type Log struct {
	mutex   sync.Mutex
	records []Record
}

func NewLog() *Log {
	return &Log{}
}
func (c *Log) Append(record Record) (uint64, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	record.Offset = uint64(len(c.records))
	c.records = append(c.records, record)
	return record.Offset, nil
}

func (c *Log) Read(offset uint64) (Record, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if offset >= uint64(len(c.records)) {
		return Record{}, ErrOffsetNotFound
	}
	return c.records[offset], nil
}

type Record struct {
	Value  []byte `json:"value"`
	Offset uint64 `json:"offset"`
}

var ErrOffsetNotFound = fmt.Errorf("offset not found")
