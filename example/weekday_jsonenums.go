// generated by jsonenums -type=WeekDay; DO NOT EDIT

package main

import (
	"encoding/json"
	"fmt"
)

func (r WeekDay) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := map[WeekDay]string{
		Monday: "Monday", Tuesday: "Tuesday", Wednesday: "Wednesday", Thursday: "Thursday", Friday: "Friday", Saturday: "Saturday", Sunday: "Sunday",
	}[r]
	if !ok {
		return nil, fmt.Errorf("invalid WeekDay: %d", r)
	}
	return json.Marshal(s)
}

var _WeekDayNameToValue = map[string]WeekDay{
	"Monday": Monday, "Tuesday": Tuesday, "Wednesday": Wednesday, "Thursday": Thursday, "Friday": Friday, "Saturday": Saturday, "Sunday": Sunday,
}

func init() {
	var v WeekDay
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_WeekDayNameToValue = map[string]WeekDay{
			interface{}(Monday).(fmt.Stringer).String(): Monday, interface{}(Tuesday).(fmt.Stringer).String(): Tuesday, interface{}(Wednesday).(fmt.Stringer).String(): Wednesday, interface{}(Thursday).(fmt.Stringer).String(): Thursday, interface{}(Friday).(fmt.Stringer).String(): Friday, interface{}(Saturday).(fmt.Stringer).String(): Saturday, interface{}(Sunday).(fmt.Stringer).String(): Sunday,
		}
	}
}

func (r *WeekDay) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("WeekDay should be a string, got %s", data)
	}
	v, ok := _WeekDayNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid WeekDay %q", s)
	}
	*r = v
	return nil
}
