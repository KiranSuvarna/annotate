package annotate

import (
	"fmt"
	"time"

	"github.com/twinj/uuid"
)



type RFC3339 struct {
	time.Time
}

func (t RFC3339) MarshalJSON() ([]byte, error) {
    return []byte(`"` + t.Format(time.RFC3339) + `"`), nil
}

type Annotation struct {
	Id           string
	Message      string
	StartDate    RFC3339
	EndDate      RFC3339
	CreationUser string
	Url          string
	Source       string
	Host         string
	Owner        string
	Category     string
}

const (
	Message      = "Message"
	StartDate    = "StartDate"
	EndDate      = "EndDate"
	Source       = "Source"
	Host         = "Host"
	CreationUser = "CreationUser"
	Owner        = "Owner"
	Category     = "Category"
)

type Annotations []Annotation

func (a *Annotation) SetGUID() error {
	if a.Id != "" {
		return fmt.Errorf("GUID already set: %v", a.Id)
	}
	a.Id = uuid.NewV4().String()
	return nil
}

func (a *Annotation) SetNow() {
	a.StartDate.Time = time.Now()
	a.EndDate = a.StartDate
}

func (a *Annotation) IsTimeNotSet() bool {
	t := time.Time{}
	return a.StartDate.Equal(t) || a.EndDate.Equal(t)
}

func (a *Annotation) IsOneTimeSet() bool {
	t := time.Time{}
	return (a.StartDate.Equal(t) && !a.EndDate.Equal(t)) || (!a.StartDate.Equal(t) && a.EndDate.Equal(t))
}

// Match Times Sets Both times to the greater of the two times
func (a *Annotation) MatchTimes() {
	if a.StartDate.After(a.EndDate.Time) {
		a.EndDate = a.StartDate
		return
	}
	a.StartDate = a.EndDate
}

func (a *Annotation) ValidateTime() error {
	t := time.Time{}
	if a.StartDate.Equal(t) {
		return fmt.Errorf("StartDate is not set")
	}
	if a.EndDate.Equal(t) {
		return fmt.Errorf("StartDate is not set")
	}
	if a.EndDate.Before(a.StartDate.Time) {
		return fmt.Errorf("EndDate is before StartDate")
	}
	return nil
}
