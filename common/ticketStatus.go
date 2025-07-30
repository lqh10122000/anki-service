package common

import "errors"

// TicketStatus là enum cho trạng thái ticket
type TicketStatus string

const (
	New               TicketStatus = "NEW"
	Open              TicketStatus = "OPEN"
	Pending           TicketStatus = "PENDING"
	Solved            TicketStatus = "SOLVED"
	SubmitAsDuplicate TicketStatus = "SUBMIT_AS_DUPLICATE_SOLVED"
)

// Validate kiểm tra tính hợp lệ
func (ts TicketStatus) Validate() error {
	switch ts {
	case New, Open, Pending, Solved, SubmitAsDuplicate:
		return nil
	}
	return errors.New("invalid ticket status")
}

// ToString chuyển sang string
func (ts TicketStatus) ToString() string {
	return string(ts)
}

// FromString chuyển từ string sang TicketStatus
func FromString(s string) (TicketStatus, error) {
	status := TicketStatus(s)
	if err := status.Validate(); err != nil {
		return "", err
	}
	return status, nil
}
