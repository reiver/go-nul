package nul

import (
	"fmt"
)

type Nullable[T any] struct {
	value T
	isnull bool
	something bool
}

func Nothing[T any]() Nullable[T] {
	var nothing Nullable[T]

	return nothing
}

func Null[T any]() Nullable[T] {
	return Nullable[T]{
		isnull:true,
	}
}

func Something[T any](value T) Nullable[T] {
	return Nullable[T]{
		something:true,
		value:value,
	}
}

func (receiver Nullable[T]) isnothing() bool {
	return !receiver.something && !receiver.isnull
}

func (receiver Nullable[T]) Filter(fn func(T)bool) Nullable[T] {
	if receiver.isnothing() {
		return Nothing[T]()
	}
	if receiver.isnull {
		return Nothing[T]()
	}

	if !fn(receiver.value) {
		return Nothing[T]()
	}

	return receiver
}

func (receiver Nullable[T]) Get() (T, bool) {
	return receiver.value, receiver.something
}

func (receiver Nullable[T]) GoString() string {
	if receiver.isnothing() {
		return fmt.Sprintf("nul.Nothing[%T]()", receiver.value)
	}
	if receiver.isnull {
		return fmt.Sprintf("nul.Null[%T]()", receiver.value)
	}

	return fmt.Sprintf("nul.Something[%T](%#v)", receiver.value, receiver.value)
}

func (receiver Nullable[T]) WhenNothing(fn func()) {
	if receiver.isnothing() {
		fn()
	}
}

func (receiver Nullable[T]) WhenNull(fn func()) {
	if receiver.isnull {
		fn()
	}
}

func (receiver Nullable[T]) WhenSomething(fn func(T)) {
	if receiver.something {
		fn(receiver.value)
	}
}
