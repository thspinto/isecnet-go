package errors

type ErrorType struct {
	t string
}

var (
	ErrorTypeUnknown      = ErrorType{"unknown"}
	ErrorTypeAlarmCentral = ErrorType{"alarm-central-error"}
)

type SlugError struct {
	error     string
	slug      string
	errorType ErrorType
}

func (s SlugError) Error() string {
	return s.error
}

func (s SlugError) Slug() string {
	return s.slug
}

func (s SlugError) ErrorType() ErrorType {
	return s.errorType
}

func NewAlarmCentralError(error string, slug string) SlugError {
	return SlugError{
		error:     error,
		slug:      slug,
		errorType: ErrorTypeAlarmCentral,
	}
}

func NewSlugError(error string, slug string) SlugError {
	return SlugError{
		error:     error,
		slug:      slug,
		errorType: ErrorTypeUnknown,
	}
}
