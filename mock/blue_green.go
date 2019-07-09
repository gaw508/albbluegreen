package mock

import "github.com/gaw508/albbluegreen"

var _ albbluegreen.BlueGreenService = &BlueGreenService{}

type BlueGreenService struct {
	StatusFn      func() (status albbluegreen.BlueGreenStatus, err error)
	StatusInvoked bool

	SetStatusFn      func(status albbluegreen.BlueGreenStatus) error
	SetStatusInvoked bool

	ToggleFn func() (newStatus albbluegreen.BlueGreenStatus, err error)
	ToggleInvoked bool
}

func (s *BlueGreenService) Status() (status albbluegreen.BlueGreenStatus, err error) {
	s.StatusInvoked = true
	return s.StatusFn()
}

func (s *BlueGreenService) SetStatus(status albbluegreen.BlueGreenStatus) error {
	s.SetStatusInvoked = true
	return s.SetStatusFn(status)
}

func (s *BlueGreenService) Toggle() (newStatus albbluegreen.BlueGreenStatus, err error) {
	s.ToggleInvoked = true
	return s.ToggleFn()
}
