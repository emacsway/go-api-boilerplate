package identity

import (
	"encoding/json"
	"errors"

	"github.com/google/uuid"
)

// Identity data to be encode in auth token
type Identity struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	Roles []string  `json:"roles"`
}

// FromGoogleData sets *i to a copy of data.
func (i *Identity) FromGoogleData(data json.RawMessage) error {
	if i == nil {
		return errors.New("auth.Identity: FromGoogleData on nil pointer")
	}

	err := json.Unmarshal(data, i)
	if err != nil {
		return err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	var defaultRoles []string
	i.ID = id
	i.Roles = defaultRoles

	return nil
}

// FromFacebookData sets *i to a copy of data.
func (i *Identity) FromFacebookData(data json.RawMessage) error {
	if i == nil {
		return errors.New("auth.Identity: FromFacebookData on nil pointer")
	}

	err := json.Unmarshal(data, i)
	if err != nil {
		return err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	var defaultRoles []string
	i.ID = id
	i.Roles = defaultRoles

	return nil
}

// New returns a new Identity
func New(id uuid.UUID, email string, roles []string) *Identity {
	return &Identity{id, email, roles}
}
