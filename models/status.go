package models

type Status struct {
	Finished   bool `json:"finished"`
	Notstarted bool `json:"notstarted"`
	Skipped    bool `json:"skipped"`
}
