package fixtures

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/pivotal-cf/jhanda"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
)

type VariadicInterface interface {
	VariadicMethod(someParams ...int)
}

type ChanInterface interface {
	ChanMethod(chan int, <-chan string) chan<- bool
}

type ModuleInterface interface {
	SomeMethod(usage jhanda.Usage)
}

type DuplicateArgumentInterface interface {
	Duplicates(string, string, int) (string, int, int)
}

type FunctionInterface interface {
	FuncMethod(func(string) error) func(int) bool
}

type NamedInterface interface {
	SomeMethod(someParam *bytes.Buffer) (someResult io.Reader)
}

type Result[T, E any] struct {
	Value T
	Error E
}

type GenericInterface[T, S comparable] interface {
	SomeMethod(map[T]S) Result[int, error]
}

type BurntSushiParser struct {
	Key toml.Key
}

type PackageConflictInterface interface {
	ListNodes(ctx context.Context) (*v1.NodeList, error)
	ListDaemonsets(ctx context.Context) (*appsv1.DaemonSetList, error)
}

// Type aliases
type StringAlias = string
type IntAlias = int
type MapAlias = map[string]interface{}
type SliceAlias = []string
type ErrorAlias = error

// More complex type aliases
type HandlerFunc = func(http.ResponseWriter, *http.Request)
type MiddlewareFunc = func(HandlerFunc) HandlerFunc

// Generic type aliases
type GenericAlias[T any] = []T
type GenericMapAlias[K comparable, V any] = map[K]V

// Interface with type aliases
type AliasInterface interface {
	ProcessString(StringAlias) IntAlias
	ProcessMap(MapAlias) SliceAlias
	HandleError() ErrorAlias
	HandleRequest(HandlerFunc) MiddlewareFunc
}

// Interface with generic aliases
type GenericAliasInterface[T any, K comparable] interface {
	ProcessGeneric(GenericAlias[T]) GenericMapAlias[K, T]
}

// Interface with complex nested generics
type NestedGenericInterface[T any] interface {
	NestedMethod(Result[T, error]) Result[[]T, map[string]error]
	ComplexMethod(map[string]Result[T, error]) ([]Result[T, error], error)
}

// Simplified interfaces that work with current implementation
type SimpleGenericAliasInterface interface {
	ProcessString(StringAlias) IntAlias
	ProcessMap(MapAlias) ErrorAlias
}

type SimpleNestedGenericInterface[T any] interface {
	SimpleMethod(value T) Result[T, error]
	ProcessValue(Result[T, error]) T
}
