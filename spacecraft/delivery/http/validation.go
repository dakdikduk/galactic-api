package http

import (
	"errors"
	"net/url"

	"github.com/dakdikduk/galactic-api/domain"
)

func validateSpacecraft(spacecraft domain.Spacecraft) error {

	_, err := url.ParseRequestURI(spacecraft.Image)
	if err == nil {
		return errors.New("invalid image url")
	}

	return nil
}
