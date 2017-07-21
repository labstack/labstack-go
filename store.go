package labstack

import (
	"fmt"
	"time"

	"github.com/dghubble/sling"
	"github.com/labstack/gommon/log"
)

type (
	// Store defines the LabStack store service.
	Store struct {
		sling  *sling.Sling
		logger *log.Logger
	}

	// StoreEntry defines the store entry
	StoreEntry struct {
		Key       string      `json:"key"`
		Value     interface{} `json:"value"`
		CreatedAt time.Time   `json:"created_at"`
		UpdatedAt time.Time   `json:"updated_at"`
	}

	// StoreQueryParams defines the query parameters for find entries.
	StoreQueryParams struct {
		Filters string `url:"filters"`
		Limit   int    `url:"limit"`
		Offset  int    `url:"offset"`
	}

	// StoreQueryResponse defines the query response.
	StoreQueryResponse struct {
		Total   int64         `json:"total"`
		Entries []*StoreEntry `json:"entries"`
	}

	// StoreError defines the store error.
	StoreError struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
)

func (s *Store) Insert(key string, value interface{}) (*StoreEntry, error) {
	e := &StoreEntry{
		Key:   key,
		Value: value,
	}
	se := new(StoreError)
	_, err := s.sling.Post("/store").BodyJSON(e).Receive(e, se)
	if err != nil {
		return nil, err
	}
	if se.Code == 0 {
		return e, nil
	}
	return nil, se
}

func (s *Store) Get(key string) (*StoreEntry, error) {
	e := new(StoreEntry)
	se := new(StoreError)
	_, err := s.sling.Get("/store/"+key).Receive(e, se)
	if err != nil {
		return nil, err
	}
	if se.Code == 0 {
		return e, nil
	}
	return nil, se
}

func (s *Store) Query() (*StoreQueryResponse, error) {
	return s.QueryWithParams(StoreQueryParams{})
}

func (s *Store) QueryWithParams(params StoreQueryParams) (*StoreQueryResponse, error) {
	qr := new(StoreQueryResponse)
	se := new(StoreError)
	_, err := s.sling.Get("/store").QueryStruct(&params).Receive(qr, se)
	if err != nil {
		return nil, err
	}
	if se.Code == 0 {
		return qr, nil
	}
	return nil, se
}

func (s *Store) Update(key string, value interface{}) error {
	se := new(StoreError)
	_, err := s.sling.Put("/store/"+key).BodyJSON(&StoreEntry{
		Value: value,
	}).Receive(nil, se)
	if err != nil {
		return err
	}
	if se.Code == 0 {
		return nil
	}
	return se
}

func (s *Store) Delete(key string) error {
	se := new(StoreError)
	_, err := s.sling.Delete("/store/"+key).Receive(nil, se)
	if err != nil {
		return err
	}
	if se.Code == 0 {
		return nil
	}
	return se
}

func (e *StoreError) Error() string {
	return fmt.Sprintf("store error, code=%d, message=%s", e.Code, e.Message)
}
