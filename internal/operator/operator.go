package operator

type Operator struct {
	Operator string
}

func (o *Operator) IsEq() bool {
	return o.Operator == "="
}

func (o *Operator) IsNeq() bool {
	return o.Operator == "<>"
}

func (o *Operator) IsGt() bool {
	return o.Operator == ">"
}

func (o *Operator) IsGte() bool {
	return o.Operator == ">="
}

func (o *Operator) IsLt() bool {
	return o.Operator == "<"
}

func (o *Operator) IsLte() bool {
	return o.Operator == "<="
}

func (o *Operator) IsLike() bool {
	return o.Operator == "LIKE"
}

func (o *Operator) IsNotLike() bool {
	return o.Operator == "NOT LIKE"
}

func (o *Operator) IsSimilar() bool {
	return o.Operator == "SIMILAR"
}

func (o *Operator) IsNotSimilar() bool {
	return o.Operator == "NOT SIMILAR"
}

func (o *Operator) IsIn() bool {
	return o.Operator == "IN"
}

func (o *Operator) IsNotIn() bool {
	return o.Operator == "NOT IN"
}

func (o *Operator) IsNull() bool {
	return o.Operator == "IS NULL"
}

func (o *Operator) IsNotNull() bool {
	return o.Operator == "IS NOT NULL"
}

func (o *Operator) IsAnd() bool {
	return o.Operator == "AND"
}

func (o *Operator) IsOr() bool {
	return o.Operator == "OR"
}
