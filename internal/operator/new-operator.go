package operator

func Eq() Operator {
	return Operator{Operator: "="}
}

func Neq() Operator {
	return Operator{Operator: "<>"}
}

func Gt() Operator {
	return Operator{Operator: ">"}
}

func Gte() Operator {
	return Operator{Operator: ">="}
}

func Lt() Operator {
	return Operator{Operator: "<"}
}

func Lte() Operator {
	return Operator{Operator: "<="}
}

func Like() Operator {
	return Operator{Operator: "LIKE"}
}

func NotLike() Operator {
	return Operator{Operator: "NOT LIKE"}
}

func Similar() Operator {
	return Operator{Operator: "SIMILAR"}
}

func NotSimilar() Operator {
	return Operator{Operator: "NOT SIMILAR"}
}

func In() Operator {
	return Operator{Operator: "IN"}
}

func NotIn() Operator {
	return Operator{Operator: "NOT IN"}
}

func IsNull() Operator {
	return Operator{Operator: "IS NULL"}
}

func IsNotNull() Operator {
	return Operator{Operator: "IS NOT NULL"}
}

func And() Operator {
	return Operator{Operator: "AND"}
}

func Or() Operator {
	return Operator{Operator: "OR"}
}
