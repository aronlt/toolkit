package demo

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type LeaveType string

const (
	AnnualLeave LeaveType = "AnnualLeave"
	Sick                  = "Sick"
	BankHoliday           = "BankHoliday"
	Other                 = "Other"
)

func (lt *LeaveType) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	leaveType := LeaveType(s)
	switch leaveType {
	case AnnualLeave, Sick, BankHoliday, Other:
		*lt = leaveType
		return nil
	}
	return errors.New("Invalid leave type")
}

// MarshalJSON implements json.Marshaler.
func (lt LeaveType) MarshalJSON() ([]byte, error) {
	s := string(lt)
	return json.Marshal(s)
}

func (lt *LeaveType) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	return lt.UnmarshalJSON(bytes)
}

func (lt LeaveType) Value() (driver.Value, error) {
	return lt.MarshalJSON()
}
