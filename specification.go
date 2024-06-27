package pgspecification

import (
	"fmt"
	"strings"
)

type Type int

const (
	TypeCondition Type = iota
	TypeLimiting
	TypeOrdering
)

type Specifier interface {
	Or(specifications ...Specification) Specification
	And(specifications ...Specification) Specification
}

type Specification interface {
	Query() string
	Value() []interface{}
	Type() Type
	IsValid() bool
}

type OrderingSpecification interface {
	Specification
}

const OrderingDirectionAsc OrderingDirection = "asc"
const OrderingDirectionDesc OrderingDirection = "desc"
const orderingDirectionZero OrderingDirection = ""

type OrderingDirection string

func (d OrderingDirection) IsZero() bool {
	return d == orderingDirectionZero
}

func (d OrderingDirection) String() string {
	return string(d)
}

type LimitSpecification interface {
	Specification
}

type OrSpecification struct {
	specifications []Specification
}

func (s *OrSpecification) IsValid() bool {
	for _, spec := range s.specifications {
		if !spec.IsValid() {
			return false
		}
	}

	return true
}

func (s *OrSpecification) Type() Type {
	return TypeCondition
}

func (s *OrSpecification) Query() string {
	var queries []string
	for _, spec := range s.specifications {
		if spec != nil {
			if s.Type() == spec.Type() {
				queries = append(queries, spec.Query())
			}
		}
	}

	query := strings.Join(queries, " or ")

	return fmt.Sprintf("(%s)", query)
}

func (s *OrSpecification) Value() []interface{} {
	var values []interface{}
	for _, spec := range s.specifications {
		if spec != nil {
			if s.Type() == spec.Type() {
				values = append(values, spec.Value()...)
			}
		}
	}
	return values
}

type AndSpecification struct {
	specifications []Specification
}

func (s *AndSpecification) IsValid() bool {
	for _, spec := range s.specifications {
		if !spec.IsValid() {
			return false
		}
	}

	return true
}

func (s *AndSpecification) Type() Type {
	return TypeCondition
}

func (s *AndSpecification) Query() string {
	var queries []string
	for _, spec := range s.specifications {
		if spec != nil {
			if s.Type() == spec.Type() {
				queries = append(queries, spec.Query())
			}
		}
	}

	query := strings.Join(queries, " and ")

	return fmt.Sprintf("(%s)", query)
}

func (s *AndSpecification) Value() []interface{} {
	var values []interface{}
	for _, spec := range s.specifications {
		if spec != nil {
			if s.Type() == spec.Type() {
				values = append(values, spec.Value()...)
			}
		}
	}
	return values
}
