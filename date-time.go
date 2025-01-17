package apipgu

import (
	"fmt"
	"strings"
	"time"
)

// "date": "2023-11-02T07:27:22.586+0300"
const apipguLayout = "2006-01-02T15:04:05.000-0700"
const apipguLayoutWithoutOffset = "2006-01-02T15:04:05.000"

// DateTime - дата и время в формате API ЕПГУ.
//
//	2023-11-02T07:27:22.586+0300
type DateTime struct {
	time.Time
}

func (d *DateTime) UnmarshalJSON(b []byte) (err error) {
	s := string(b)
	if s == "null" {
		d.Time = time.Time{}
		return
	}
	s = strings.Trim(string(b), `"`)
	d.Time, err = time.Parse(apipguLayout, s)
	return
}

func (d DateTime) MarshalJSON() ([]byte, error) {
	if d.Time.IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, d.Time.Format(apipguLayout))), nil
}

func (d DateTime) GoString() string {
	return d.Time.Format(apipguLayoutWithoutOffset)
}
