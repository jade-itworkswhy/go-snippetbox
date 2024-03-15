package models

import (
	"errors"
)

// encapsulate the sql error
var ErrNoRecord = errors.New("models: no matching record found")
