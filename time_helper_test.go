package main

import (
	"strconv"
	"testing"
	"time"
)

func TestParsingSomeUnixTimestamps(t *testing.T) {
	timeExample := time.Now().Truncate(time.Second)

	var SecondsTests = []struct {
		in   string
		out  time.Time
		err  bool
		name string
	}{
		{strconv.FormatInt(timeExample.Unix(), 10), timeExample, false, "Time should be converted"},
		{strconv.FormatInt(timeExample.Add(time.Hour).Unix(), 10), timeExample.Add(time.Hour), false, "Time should be converted"},
		{"abc", timeExample, true, "Invalid time containing string"},
		{"900000000000000000000", timeExample, true, "Overflow timestamp"},
	}

	for _, tt := range SecondsTests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := parseUnixTime(tt.in)

			if err == nil && s != tt.out {
				t.Errorf("got %v, want %v", s, tt.out)
			}

			if tt.err && err == nil || !tt.err && err != nil {
				t.Errorf("got nil  error but was expecting one")
			}
		})
	}
}

func TestMapToTimestampFromStringArray(t *testing.T) {
	timeExample := time.Now().Truncate(time.Second)
	validStringArray := []string{strconv.FormatInt(timeExample.Unix(), 10), strconv.FormatInt(timeExample.Add(time.Hour).Unix(), 10), strconv.FormatInt(timeExample.Add(time.Minute).Unix(), 10)}

	t.Run("Expect to convert all the values correctly", func(t *testing.T) {
		values, err := MapToTimestamps(validStringArray)
		if len(values) != len(validStringArray) {
			t.Errorf("got [%v] for lenthg, but expected [%v]", len(values), len(validStringArray))
		}

		if err != nil {
			t.Errorf("got unexpected error in parsing [%v]", err)
		}
	})
}

func TestMapToTimestampFromStringArrayWithUnexpectedValue(t *testing.T) {
	timeExample := time.Now().Truncate(time.Second)
	validStringArray := []string{strconv.FormatInt(timeExample.Unix(), 10), strconv.FormatInt(timeExample.Add(time.Hour).Unix(), 10), strconv.FormatInt(timeExample.Add(time.Minute).Unix(), 10), "abcd"}

	t.Run("Last value is errored; we should get an error.", func(t *testing.T) {
		values, err := MapToTimestamps(validStringArray)

		if err == nil {
			t.Errorf("Got no errors, but expected one. ")
		}

		if len(values) != 0 {
			t.Errorf("Got non empty array on error [%v]", values)
		}
	})
}
