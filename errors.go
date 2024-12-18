package timephrase

import "errors"

var (
	ErrEmptyString     = errors.New("date string cannot be empty")
	ErrInvalidFormat   = errors.New("date string format is not recognized")
	ErrInvalidKeyworkd = errors.New("invalid natural language keyword")

	ErrAbsoluteParse = errors.New("could not parse absolute date")
	ErrRelativeParse = errors.New("could not parse relative date")
)
