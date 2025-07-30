package common

import (
	"errors"
	"time"
)

// AvroMapper ánh xạ dữ liệu Avro
type AvroMapper struct{}

// ToTicketStatus ánh xạ từ Avro statusEnum
func (m *AvroMapper) ToTicketStatus(avroStatus interface{}) (TicketStatus, error) {
	if avroStatus == nil {
		return "", nil
	}
	status, ok := avroStatus.(string)
	if !ok {
		return "", errors.New("invalid avro status type")
	}
	return FromString(status)
}

// ToTime ánh xạ từ Avro timestamp-millis
func (m *AvroMapper) ToTime(avroTimestamp interface{}) (*time.Time, error) {
	if avroTimestamp == nil {
		return nil, nil
	}
	timestamp, ok := avroTimestamp.(int64)
	if !ok {
		return nil, errors.New("invalid avro timestamp type")
	}
	t := time.UnixMilli(timestamp).UTC()
	return &t, nil
}
