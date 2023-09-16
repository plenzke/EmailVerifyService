package exceptions

type ExceptionType struct {
	exception string
}

var (
	ExceptionTypeUnknown        = ExceptionType{"unknown"}
	ExceptionTypeIncorrectInput = ExceptionType{"incorrect-input"}
)

type SlugException struct {
	exception     string
	slug          string
	exceptionType ExceptionType
}

func (s SlugException) Exception() string {
	return s.exception
}

func (s SlugException) Slug() string {
	return s.slug
}

func (s SlugException) ExceptionType() ExceptionType {
	return s.exceptionType
}

func NewSlugException(exception string, slug string) SlugException {
	return SlugException{
		exception:     exception,
		slug:          slug,
		exceptionType: ExceptionTypeUnknown,
	}
}

func NewIncorrectInputException(exception string, slug string) SlugException {
	return SlugException{
		exception:     exception,
		slug:          slug,
		exceptionType: ExceptionTypeIncorrectInput,
	}
}
