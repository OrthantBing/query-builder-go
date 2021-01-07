package utils

import "time"

var (
	CurrentTime      func() time.Time
	CurrentTimeInUTC func() time.Time
	CurrentTimeInIST func() time.Time
	ParseTime        func(string) (time.Time, error)
	ParseTimeInUTC   func(string) (time.Time, error)
	ParseTimeInIST   func(string) (time.Time, error)
	ConvertUTCtoIST  func(time.Time) (time.Time, error)
)

func init() {
	CurrentTime = func() time.Time {
		return time.Now()
	}

	CurrentTimeInUTC = func() time.Time {
		return CurrentTime().UTC()
	}

	CurrentTimeInIST = func() time.Time {
		location, _ := time.LoadLocation("Asia/Kolkata")
		return CurrentTime().In(location)
	}

	ParseTime = func(st string) (time.Time, error) {
		return time.Parse(time.RFC3339, st)
	}

	ParseTimeInUTC = func(st string) (time.Time, error) {
		t, err := ParseTime(st)
		if err != nil {
			return t, err
		}

		return t.UTC(), nil
	}

	ParseTimeInIST = func(st string) (time.Time, error) {
		t, err := ParseTime(st)
		if err != nil {
			return t, err
		}
		location, _ := time.LoadLocation("Asia/Kolkata")
		return t.In(location), nil
	}

	ConvertUTCtoIST = func(utctime time.Time) (time.Time, error) {

		location, _ := time.LoadLocation("Asia/Kolkata")
		return utctime.In(location), nil
		// zone, _ := utctime.Zone()
		// if zone == "IST" {
		// 	return utctime, nil
		// }
		// return ParseTimeInIST(utctime.String())
	}

}
