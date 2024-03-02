package src

import "context"

// Transform ...
type Transform[T, U any] func(T) U

// WithContext ...
func WithContext[T, MONAD any](ctx context.Context, f func(ctx context.Context, v T) MONAD) Transform[T, MONAD] {
	return func(v T) MONAD {
		return f(ctx, v)
	}
}

// Either represents a computation that may return a value or an error.
type Either[T any] struct {
	Value T
	Error error
}

// Right creates an Either with a value.
func Right[T any](value T) Either[T] {
	return Either[T]{Value: value}
}

// Left creates an Either with an error.
func Left[T any](err error) Either[T] {
	return Either[T]{Error: err}
}

// Bind ...
func (e Either[T]) Bind(trans Transform[T, Either[T]]) Either[T] {
	if e.Error != nil {
		return e
	}
	return trans(e.Value)
}

// Bind ...
func Bind[T, U any](e Either[T], trans Transform[T, Either[U]]) Either[U] {
	if e.Error != nil {
		return Left[U](e.Error)
	}
	return trans(e.Value)
}

// Unwrap ...
func (e Either[T]) Unwrap() (T, error) {
	return e.Value, e.Error
}
