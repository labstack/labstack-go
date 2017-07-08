package labstack

import (
	"fmt"
	"net/http"
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
		Filters string
		Limit   int
		Offset  int
	}

	// StoreQueryResponse defines the query response.
	StoreQueryResponse struct {
		Total   int64         `json:"total"`
		Entries []*StoreEntry `json:"entries"`
	}
)

func (s *Store) Insert(key string, value interface{}) (e *StoreEntry, err error) {
	e = &StoreEntry{
		Key:   key,
		Value: value,
	}
	res, err := s.sling.Post("/store").BodyJSON(e).Receive(e, nil)
	if err != nil {
		return
	}
	if res.StatusCode != http.StatusCreated {
		err = fmt.Errorf("store: error inserting entry, status=%d, message=%v", res.StatusCode, err)
	}
	return
}

func (s *Store) Get(key string) (e *StoreEntry, err error) {
	e = new(StoreEntry)
	res, err := s.sling.Get("/store/"+key).Receive(e, nil)
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("store: error getting entry, status=%d, message=%v", res.StatusCode, err)
	}
	return
}

func (s *Store) Query() (entries []*StoreEntry, err error) {
	qr := new(StoreQueryResponse)
	res, err := s.sling.Get("/store").Receive(qr, nil)
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("store: error getting entries, status=%d, message=%v", res.StatusCode, err)
	}
	entries = qr.Entries
	return
}

func (s *Store) QueryWithParams(params *StoreQueryParams) (entries []*StoreEntry, err error) {
	qr := new(StoreQueryResponse)
	res, err := s.sling.Get("/store").QueryStruct(params).Receive(qr, nil)
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("store: error getting entries, status=%d, message=%v", res.StatusCode, err)
	}
	return
}

func (s *Store) Update(key string, value interface{}) (err error) {
	res, err := s.sling.Put("/store/"+key).BodyJSON(&StoreEntry{
		Value: value,
	}).Receive(nil, nil)
	if res.StatusCode != http.StatusNoContent {
		err = fmt.Errorf("store: error updating entry, status=%d, message=%v", res.StatusCode, err)
	}
	return
}

func (s *Store) Delete(key string) (err error) {
	res, err := s.sling.Delete("/store/"+key).Receive(nil, nil)
	if res.StatusCode != http.StatusNoContent {
		err = fmt.Errorf("store: error deleting entry, status=%d, message=%v", res.StatusCode, err)
	}
	return
}
