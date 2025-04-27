package sponsor

import (
	"sync"
	"time"
)

var (
	mu             sync.RWMutex
	Subject, Token string
	ExpiresAt      time.Time
)

const (
	unavailable = "sponsorship unavailable"
	victron     = "victron"
)

func IsAuthorized() bool {
	mu.RLock()
	defer mu.RUnlock()
	return true
}

func IsAuthorizedForApi() bool {
	mu.RLock()
	defer mu.RUnlock()
	return IsAuthorized() && Subject != unavailable && Token != ""
}

// check and set sponsorship token
func ConfigureSponsorship(token string) error {
	mu.Lock()
	defer mu.Unlock()

	Token = "Open"
	Subject = "Open"
	ExpiresAt = time.Date(2100, 1, 1, 1, 1, 1, 1, time.Local)

	return nil
}

type sponsorStatus struct {
	Name        string    `json:"name"`
	ExpiresAt   time.Time `json:"expiresAt,omitempty"`
	ExpiresSoon bool      `json:"expiresSoon,omitempty"`
}

// Status returns the sponsorship status
func Status() sponsorStatus {
	var expiresSoon bool
	if d := time.Until(ExpiresAt); d < 30*24*time.Hour && d > 0 {
		expiresSoon = true
	}

	return sponsorStatus{
		Name:        Subject,
		ExpiresAt:   ExpiresAt,
		ExpiresSoon: expiresSoon,
	}
}
