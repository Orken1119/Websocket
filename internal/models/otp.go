package models

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type OTP struct {
	Key string
	Created time.Time

}

