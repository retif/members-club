package main

import (
	"errors"
	"time"
)

type Club struct {
	Members []*Member
}

func (c *Club) addMember(m *Member) error {
	for _, elem := range c.Members {
		if elem.Email == m.Email {
			return errors.New("email already exists")
		}
	}

	c.Members = append(c.Members, m)
	return nil
}

func (c *Club) getMembers() []*Member {
	return c.Members
}

type Member struct {
	Name string
	Email string
	RegistrationDate time.Time
}
