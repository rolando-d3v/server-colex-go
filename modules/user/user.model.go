package modules

import (
	"time"
	"github.com/google/uuid"
)

// User modelo para la tabla usuarios
type User struct {
	Id        uuid.UUID `json:"id"`
	PersonaID string    `json:"persona_id"`
	ColegioID int16     `json:"colegio_id"`
	Password  string    `json:"password,omitempty"`
	Email     string    `json:"email"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}
