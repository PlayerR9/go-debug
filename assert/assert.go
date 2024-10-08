package assert

import (
	"fmt"
	"strconv"
)

// Assert panics if cond is false. It is intended to be used for debugging.
//
// Parameters:
//   - cond: the condition to check.
//   - msg: the message to print if the condition is false.
//
// Example:
//
//	foo := "foo"
//	Assert(foo == "bar", "foo is not \"bar\"") // panics: "[ASSERT FAILED]: foo is not \"bar\""
func Assert(cond bool, msg string) {
	if cond {
		return
	}

	panic(NewErrAssertFailed(msg))
}

// AssertF panics if cond is false. It is intended to be used for debugging.
//
// Parameters:
//   - cond: the condition to check.
//   - format: the format of the message.
//   - args: the arguments of the format.
//
// Example:
//
//	foo := "foo"
//	bar := "bar"
//	AssertF(foo == bar, "%q is not %q", foo, bar) // panics: "[ASSERT FAILED]: \"foo\"" is not \"bar\""
func AssertF(cond bool, format string, args ...any) {
	if cond {
		return
	}

	msg := fmt.Sprintf(format, args...)

	panic(NewErrAssertFailed(msg))
}

// AssertOk panics if v is false. It is intended to be used for debugging.
//
// Parameters:
//   - ok: the value to check.
//   - format: the format of the message.
//   - args: the arguments of the format.
//
// Example:
//
//	foo := "foo"
//	ok := my_function(foo, "bar")
//	AssertOk(ok, "my_function(%s, %q)", foo, "bar")
//	// panics: "[ASSERT FAILED]: my_function(foo, \"bar\"") = false"
func AssertOk(ok bool, format string, args ...any) {
	if ok {
		return
	}

	msg := fmt.Sprintf(format, args...) + " = false"

	panic(NewErrAssertFailed(msg))
}

// AssertOk panics if v is true. It is intended to be used for debugging.
//
// Parameters:
//   - ok: the value to check.
//   - format: the format of the message.
//   - args: the arguments of the format.
//
// Example:
//
//	foo := "foo"
//	ok := my_function(foo, "bar")
//	AssertNotOk(!ok, "my_function(%s, %q)", foo, "bar")
//	// panics: "[ASSERT FAILED]: my_function(foo, \"bar\"") = true"
func AssertNotOk(ok bool, format string, args ...any) {
	if !ok {
		return
	}

	msg := fmt.Sprintf(format, args...) + " = true"

	panic(NewErrAssertFailed(msg))
}

// AssertErr panics if err is not nil. It is intended to be used for debugging.
//
// Parameters:
//   - err: the error to check.
//   - format: the format of the message.
//   - args: the arguments of the format.
//
// Example:
//
//	foo := "foo"
//	err := my_function(foo, "bar")
//	AssertErr(err, "my_function(%s, %q)", foo, "bar")
//	// panics: "[ASSERT FAILED]: my_function(foo, \"bar\"") = <err>"
func AssertErr(err error, format string, args ...any) {
	if err == nil {
		return
	}

	msg := fmt.Sprintf(format, args...) + " = " + err.Error()

	panic(NewErrAssertFailed(msg))
}

// AssertNotNil panics if v is nil. It is intended to be used for debugging.
//
// Parameters:
//   - v: the value to check.
//   - name: the name of the value.
func AssertNotNil(v any, name string) {
	if v != nil {
		return
	}

	panic(NewErrAssertFailed(strconv.Quote(name) + " must not be nil"))
}

// TODO writes a panic message indicating that a case has not been handled yet.
//
// This function is intended to be used as a placeholder until the case is handled.
func TODO() {
	panic("TODO: Handle this case")
}

// AssertDeref tries to dereference an element and panics if it is nil.
//
// Parameters:
//   - elem: the element to dereference.
//   - param_name: the name of the parameter.
//
// Returns:
//   - T: the dereferenced element.
func AssertDeref[T any](elem *T, is_param bool, name string) T {
	if elem != nil {
		return *elem
	}

	var msg string

	if is_param {
		msg = "parameter (" + name + ")"
	} else {
		msg = "variable (" + name + ")"
	}

	msg += " expected to not be nil"

	panic(NewErrAssertFailed(msg))
}

// AssertTypeOf panics if the element is not of the expected type.
//
// Parameters:
//   - elem: the element to check.
//   - var_name: the name of the variable.
//   - allow_nil: if the element can be nil.
func AssertTypeOf[T any](elem any, target string, allow_nil bool) {
	if elem == nil {
		if !allow_nil {
			msg := fmt.Sprintf("expected %q to be of type %T, got nil instead", target, *new(T))

			panic(NewErrAssertFailed(msg))
		}

		return
	}

	_, ok := elem.(T)
	if !ok {
		msg := fmt.Sprintf("expected %q to be of type %T, got %T instead", target, *new(T), elem)

		panic(NewErrAssertFailed(msg))
	}
}

// AssertConv tries to convert an element to the expected type and panics if it is not possible.
//
// Parameters:
//   - elem: the element to check.
//   - var_name: the name of the variable.
//
// Returns:
//   - T: the converted element.
func AssertConv[T any](elem any, target string) T {
	if elem == nil {
		msg := fmt.Sprintf("expected %q to be of type %T, got nil instead", target, *new(T))

		panic(NewErrAssertFailed(msg))
	}

	res, ok := elem.(T)
	if !ok {
		msg := fmt.Sprintf("expected %q to be of type %T, got %T instead", target, *new(T), elem)

		panic(NewErrAssertFailed(msg))
	}

	return res
}
