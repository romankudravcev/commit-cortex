package components

import "time"

type Report struct {
	Repository  Repo
	ReportItems []ReportItem
}

type ReportItem struct {
	Branch string
	Commit string
	Time   time.Time
	Author string
}
