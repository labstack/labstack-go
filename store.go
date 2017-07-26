package labstack

import (
	"fmt"

	"github.com/dghubble/sling"
	"github.com/labstack/gommon/log"
)

type (
	// Store defines the LabStack store service.
	Store struct {
		sling  *sling.Sling
		logger *log.Logger
	}

	Document Fields

	// StoreSearchResponse defines the query response.
	StoreSearchResponse struct {
		Total     int64      `json:"total"`
		Documents []Document `json:"documents"`
	}

	// StoreError defines the store error.
	StoreError struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
)

func (s *Store) Insert(collection string, document Document) (Document, error) {
	se := new(StoreError)
	_, err := s.sling.Post("/store/"+collection).
		BodyJSON(&document).
		Receive(&document, se)
	if err != nil {
		return nil, err
	}
	if se.Code == 0 {
		return document, nil
	}
	return nil, se
}

func (s *Store) Get(collection, id string) (Document, error) {
	doc := Document{}
	se := new(StoreError)
	_, err := s.sling.Get(fmt.Sprintf("/store/%s/%s", collection, id)).
		Receive(&doc, se)
	if err != nil {
		return nil, err
	}
	if se.Code == 0 {
		return doc, nil
	}
	return nil, se
}

func (s *Store) Search(collection string, parameters *SearchParameters) (*StoreSearchResponse, error) {
	sr := new(StoreSearchResponse)
	se := new(StoreError)
	_, err := s.sling.Post(fmt.Sprintf("/store/%s/query", collection)).
		BodyJSON(parameters).
		Receive(sr, se)
	if err != nil {
		return nil, err
	}
	if se.Code == 0 {
		return sr, nil
	}
	return nil, se
}

func (s *Store) Update(collection string, id string, document Document) error {
	se := new(StoreError)
	_, err := s.sling.Patch(fmt.Sprintf("/store/%s/%s", collection, id)).
		BodyJSON(&document).
		Receive(nil, se)
	if err != nil {
		return err
	}
	if se.Code == 0 {
		return nil
	}
	return se
}

func (s *Store) Delete(collection string, id string) error {
	se := new(StoreError)
	_, err := s.sling.Delete(fmt.Sprintf("/store/%s/%s", collection, id)).
		Receive(nil, se)
	if err != nil {
		return err
	}
	if se.Code == 0 {
		return nil
	}
	return se
}

func (e *StoreError) Error() string {
	return e.Message
}
