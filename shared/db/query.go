package db

import (
//	"fmt"
//	"reflect"
//	"strings"
//
//	"github.com/go-gorp/gorp"
)

// Query - DB query
type Query struct {
	Filter map[string]interface{}
	Entity interface{}
}

// NewQuery - returns a new Query given some generic interface{}
func NewQuery(entity interface{}) *Query {
	return &Query{
		Filter: make(map[string]interface{}),
		Entity: entity,
	}
}
