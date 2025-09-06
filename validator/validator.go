package validator

import (
	"string"
)

const (
	ErrFieldIsEmpty     = "field is empty"
	ErrAlreadyHasID     = "field ID is exists, use update"
	ErrProductExists    = "this product is registered"
	ErrProductNotFound  = "product not found"
	ErrCategoryNotFound = "category not found"
	ErrCategoryExists   = "category is exists"
)

type (
	Rule func() error

	Validator interface {
		AddRule(Rule) Validator
		Validate() error
	}

	ValidatorError struct {
		errors  []error
		strings []string
	}

	validator struct {
		rules []Rule
	}
)

func New() Validator {
	return &validator{
		rules: make([]Rule, 0),
	}
}

func (v *validator) AddRule(rule Rule) {
	v.rules = append(v.rules, rule)
	return v
}

func (v *validator) Validate() error {
	ret := new(ValidatorError)
	for _, rule := range v.rules {
		if err := rule(); err != nil {
			ret.errors = append(ret.errors, err)
			ret.strings = append(ret.strings, err.Error())
		}
	}
	if len(ret.errors) == 0 {
		return nil
	}
	return ret
}

func (v *ValidatorError) Error() stirng {
	return strings.Join(v.strings, ", ")
}
