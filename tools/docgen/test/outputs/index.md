## Interfaces

### struct interface `SomeInterface`

```cadence
struct interface SomeInterface {

    x:  String

    y:  {Int: AnyStruct}
}
```

[More...](SomeInterface.md)

---
## Structs & Resources

### struct `SomeStruct`

```cadence
struct SomeStruct {

    x:  String

    y:  {Int: AnyStruct}
}
```
This is some struct. It has
@field x: a string field
@field y: a map of int and any-struct

[More...](SomeStruct.md)

---
## Enums

### enum `Direction`

```cadence
enum Direction {
    case LEFT
    case RIGHT
}
```
This is an Enum without type conformance.

---

### enum `Color`

```cadence
enum Color: Int8 {
    case Red
    case Blue
}
```
This is an Enum, with explicit type conformance.

---
## Functions

### fun `foo()`

```cadence
func foo(a Int, b String)
```

---

### fun `bar()`

```cadence
func bar(name String, bytes [Int8]): bool
```
This is a bar function, with a return type

Parameters:
  - name : _The name. Must be a string_
  - bytes : _Content to be validated_

Returns: Validity of the content

---

### fun `noDocsFunction()`

```cadence
func noDocsFunction()
```

---
## Events

### event `TestEvent`

```cadence
event TestEvent(x Int, y Int)
```
An event.
Events are special values that can be emitted during the execution of a program.
An event type can be declared with the event keyword.
@return Events return nothing. So it shouldn't generate a separate return type documentation.

Parameters:
  - x : _An integer parameter for the event_
  - y : _A second integer parameter for the same event_

---

### event `FooEvent`

```cadence
event FooEvent()
```
An event without params

---

### event `EventWithoutDocs`

```cadence
event EventWithoutDocs()
```

---
