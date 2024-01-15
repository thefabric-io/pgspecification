package pgspecification

import (
	"strings"
)

func ComputeSpecifications(ss ...Specification) (string, []interface{}) {
	b := strings.Builder{}
	var values []interface{}

	qConditions, vConditions := ComputeConditions(ss...)
	b.WriteString(qConditions)
	values = append(values, vConditions...)

	qOrdering, vOrderings := ComputeOrderingsOnly(ss...)
	b.WriteString(qOrdering)
	values = append(values, vOrderings...)

	qLimiting, vLimitings := ComputeLimitingsOnly(ss...)
	b.WriteString(qLimiting)
	values = append(values, vLimitings...)

	return b.String(), values
}

func Filter(t Type, ss ...Specification) []Specification {
	result := make([]Specification, 0)
	for _, s := range ss {
		if s == nil {
			continue
		}
		if s.Type() == t {
			result = append(result, s)
		}
	}

	return result
}

func ConditionsOnly(ss ...Specification) []Specification {
	return Filter(TypeCondition, ss...)
}

func ComputeConditions(ss ...Specification) (string, []interface{}) {
	b := strings.Builder{}
	var values []interface{}

	for _, c := range ConditionsOnly(ss...) {
		if c.IsValid() {
			b.WriteString(" and ")
			b.WriteString(c.Query())
			values = append(values, c.Value()...)
		}
	}

	return b.String(), values
}

func OrderingsOnly(ss ...Specification) []Specification {
	return Filter(TypeOrdering, ss...)
}

func ComputeOrderingsOnly(ss ...Specification) (string, []interface{}) {
	b := strings.Builder{}
	var values []interface{}

	for _, c := range OrderingsOnly(ss...) {
		b.WriteString(" ")
		b.WriteString(c.Query())
		values = append(values, c.Value()...)
	}

	return b.String(), values
}

func LimitingsOnly(ss ...Specification) []Specification {
	return Filter(TypeLimiting, ss...)
}

func ComputeLimitingsOnly(ss ...Specification) (string, []interface{}) {
	b := strings.Builder{}
	var values []interface{}

	for _, c := range LimitingsOnly(ss...) {
		b.WriteString(" ")
		b.WriteString(c.Query())
		values = append(values, c.Value()...)
	}

	return b.String(), values
}
