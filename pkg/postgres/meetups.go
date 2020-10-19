package postgres

import (
	"github.com/go-pg/pg/v9"
	"yogan.dev/meetmeup/pkg/models"
)

// MeetupsRepo holds a pointer to the db reference
type MeetupsRepo struct {
	DB *pg.DB
}

// GetMeetups returns all meetups
func (m *MeetupsRepo) GetMeetups() ([]*models.Meetup, error) {
	var meetups []*models.Meetup

	err := m.DB.Model(&meetups).Select()
	if err != nil {
		return nil, err
	}

	return meetups, nil
}

// CreateMeetup takes input and creates a new entry
func (m *MeetupsRepo) CreateMeetup(meetup *models.Meetup) (*models.Meetup, error) {
	_, err := m.DB.Model(meetup).Returning("*").Insert()
	if err != nil {
		return nil, err
	}

	return meetup, nil
}
