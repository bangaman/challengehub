package hasuragraphql

import (
	"time"
)


func GetDate(date string) string {
	dt,_ := time.Parse("2006-01-02", date)
	dformatlast := dt.Format("Jan 2 ")

	return dformatlast
}