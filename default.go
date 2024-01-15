package pgspecification

func NewSpecifier() *DefaultSpecifier {
	return &DefaultSpecifier{}
}

type DefaultSpecifier struct{}

func (*DefaultSpecifier) Or(specifications ...Specification) Specification {
	ss := Filter(TypeCondition, specifications...)

	return &OrSpecification{
		specifications: ss,
	}
}

func (*DefaultSpecifier) And(specifications ...Specification) Specification {
	ss := Filter(TypeCondition, specifications...)

	return &AndSpecification{
		specifications: ss,
	}
}

func (s *DefaultSpecifier) Limit(value int) Specification {
	return &limit{value: value}
}

type limit struct {
	value int
}

func (h *limit) IsValid() bool {
	return h.value > 0
}

func (h *limit) Type() Type {
	return TypeLimiting
}

func (h *limit) Query() string {
	return "limit ?"
}

func (h *limit) Value() []interface{} {
	return []interface{}{h.value}
}
