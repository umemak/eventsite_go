/*
 * eventsite
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

type Event struct {

	Id int64 `json:"id,omitempty"`

	Title string `json:"title,omitempty"`

	Start time.Time `json:"start,omitempty"`

	Place string `json:"place,omitempty"`

	Open time.Time `json:"open,omitempty"`

	Close time.Time `json:"close,omitempty"`

	Author int64 `json:"author,omitempty"`
}

// AssertEventRequired checks if the required fields are not zero-ed
func AssertEventRequired(obj Event) error {
	return nil
}

// AssertRecurseEventRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Event (e.g. [][]Event), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseEventRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aEvent, ok := obj.(Event)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertEventRequired(aEvent)
	})
}
