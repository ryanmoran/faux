# Type Alias and Enhanced Generic Support

This document describes the enhanced type alias and generic support added to faux.

## Type Alias Support

Faux now supports Go type aliases defined with the `=` syntax:

```go
// Basic type aliases
type StringAlias = string
type IntAlias = int
type MapAlias = map[string]interface{}
type SliceAlias = []string
type ErrorAlias = error

// Function type aliases
type HandlerFunc = func(http.ResponseWriter, *http.Request)
type MiddlewareFunc = func(HandlerFunc) HandlerFunc

// Generic type aliases
type GenericAlias[T any] = []T
type GenericMapAlias[K comparable, V any] = map[K]V

// Interface using type aliases
type AliasInterface interface {
    ProcessString(StringAlias) IntAlias
    ProcessMap(MapAlias) SliceAlias
    HandleError() ErrorAlias
    HandleRequest(HandlerFunc) MiddlewareFunc
}
```

When generating fakes for interfaces that use type aliases, faux will:
- Preserve the alias names in the generated fake types
- Properly handle imports for aliased types from other packages
- Support both simple and complex (function, generic) type aliases

## Enhanced Generic Support

The existing generic support has been enhanced to better handle:

### Interface-level Generics
```go
type GenericInterface[T, S comparable] interface {
    SomeMethod(map[T]S) Result[int, error]
}
```

Generates:
```go
type GenericInterface[T comparable, S comparable] struct {
    // ... fake implementation with proper type parameters
}
```

### Simple Generic Type Aliases in Interfaces
```go
type SimpleGenericAliasInterface interface {
    ProcessString(StringAlias) IntAlias
    ProcessMap(MapAlias) ErrorAlias
}
```

### Simple Nested Generics
```go
type SimpleNestedGenericInterface[T any] interface {
    SimpleMethod(value T) Result[T, error]
    ProcessValue(Result[T, error]) T
}
```

Generates fakes that properly handle type parameters in straightforward scenarios.

## Implementation Details

### Type Alias Handling
- Added support for `*types.Alias` in the `NewType` function in `rendering/type.go`
- Enhanced `NewArgument` in `parsing/argument.go` to handle type aliases with type parameters
- Type aliases are resolved to maintain their original names while preserving semantic correctness

### Testing
- Added comprehensive test cases for type alias interfaces (`AliasInterface`)
- Added test cases for simple generic alias interfaces (`SimpleGenericAliasInterface`)
- Added test cases for simple nested generic interfaces (`SimpleNestedGenericInterface`)
- All acceptance tests validate the generated fake code compiles and works correctly

## Known Limitations

1. **Complex Nested Generic Type Parameters**: Complex nested generics where type parameters are used as arguments to other generic types within composite types (like `map[string]Result[T, error]`) may not always preserve type parameters correctly in all contexts. This is a complex issue in Go's type system that requires further investigation.

   For example, this interface has limitations:
   ```go
   type NestedGenericInterface[T any] interface {
       ComplexMethod(map[string]Result[T, error]) ([]Result[T, error], error)  // Type params may be lost in composite types
   }
   ```

2. **Complex Generic Type Aliases**: Generic type aliases with type parameters may not be fully resolved in complex scenarios:
   ```go
   type GenericAliasInterface[T any, K comparable] interface {
       ProcessGeneric(GenericAlias[T]) GenericMapAlias[K, T]  // Type params may not be preserved
   }
   ```

3. **Method-level Generics**: Go interfaces do not support method-level type parameters, so this is not supported by faux (this is a language limitation, not a faux limitation).

## Working Examples

The following patterns work well with the current implementation:

### Basic Type Aliases
```go
type AliasInterface interface {
    ProcessString(StringAlias) IntAlias
    HandleError() ErrorAlias
}
```

### Simple Generic Interfaces
```go
type SimpleNestedGenericInterface[T any] interface {
    SimpleMethod(value T) Result[T, error]
    ProcessValue(Result[T, error]) T
}
```

## Usage Examples

Generate a fake for an interface with type aliases:
```bash
faux --file ./interfaces.go --interface AliasInterface --output ./fakes/alias_interface.go
```

Generate a fake for a simple generic interface:
```bash
faux --file ./interfaces.go --interface SimpleNestedGenericInterface --output ./fakes/simple_nested_generic_interface.go
```

The generated fakes will correctly preserve type alias names and handle simple generic type parameters as expected.

## Test Coverage

The implementation includes comprehensive test coverage:
- `AliasInterface` - Tests basic type alias support
- `SimpleGenericAliasInterface` - Tests type aliases in non-generic interfaces
- `SimpleNestedGenericInterface` - Tests simple generic interfaces with proper type parameter handling
- All tests validate that generated code compiles and functions correctly