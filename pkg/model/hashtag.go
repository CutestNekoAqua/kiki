package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// HashTag database model.
type HashTag struct {
	*gorm.Model
	FeedID uint
	Value  string
}

// PrefixedValue return with the value with # at the beginning.
func (ht HashTag) PrefixedValue() string {
	return fmt.Sprintf("#%s", ht.Value)
}
