package validator

// ErrorResponse interface
type ErrorResponse interface {
	// AddError add new error
	AddError(field, tag, message string)
	// HasError check if response has error
	HasError() bool
	// Failed check if field has error
	// @example:
	// resp.Failed("firstname")
	Failed(field string) bool
	// FailedOn check if field has special error
	// @example:
	// resp.FailedOn("firstname", "required")
	FailedOn(field, err string) bool
	// Errors get errors list
	Errors() map[string]map[string]string
	// String convert to string
	String() string
	// Messages get error messages only without error key
	Messages() map[string][]string
	// Rules get error rules only without error message
	Rules() map[string][]string
	// MarshalJSON convert to json
	MarshalJSON() ([]byte, error)
}
