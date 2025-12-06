Learning the basics of Kotlin

### Entry point
Kotlin entry point is the `main` function

```kotlin
fun main() {
    println("Hello, world!")
}
```

### Variables
```kotlin
val a = 1 // immutable
var b = 2 // mutable
b = 3
```

### String templates
```kotlin
var a = 1
var b = 2
println("$a + $b = ${a + b}")
```

### Basic types
Kotlin can infer the type of a variable, but we can set it explicitly
```kotlin
var a: Int = 1
var b = 2
var c: String = "hello"
var d = "world"
val e: Float // uninitialized
```
The following table summarizes the basic types
| Category | Basic types |
| -------- | ----------- |
| Integers | `Byte`, `Short`, `Int`, `Long` |
| Unsigned Integers | `UByte`, `UShort`, `UInt`, `ULong` |
| Floating point | `Float`, `Double` |
| Character | `Char` |
| Boolean | `Boolean` |
| String | `String` |
- The unsigned version of an integer type has a prefix `U` `200u`
- The floating point version has a prefix `F` `3.14f`
- Characters use single quotes `'` `'a'`

### Complex types
#### Collections
Kotlin has the following collections to manage group of items: `Lists` (ordered), `Sets` (unordered) and `Maps` (key-value pairs), each one with *Mutable* and *Immutable* versions.

#### List
```kotlin
val names: List<String> = listOf("Alice", "Bob", "Charlie") // Immutable
// Mutable versions are MutableList and mutableListOf
println("First name: ${names[0]}")
// Popular methods
names.first()
names.last()
names.count()
names.add("Dave")
names.remove("Bob")
if ("Alice" in names) {
    //
}
```

#### Set
It can contain only unique values, so duplicates are ignored
```kotlin
val names: Set<String> = setOf("Alice", "Bob", "Charlie") // Immutable
// Mutable versions are MutableSet and mutableSetOf
// Popular methods
names.count()
names.add("Dave")
names.remove("Bob")
if ("Alice" in names) {
    //
}
```

#### Map
```kotlin
val names: Map<String, Int> = mapOf("Alice" to 1, "Bob" to 2, "Charlie" to 3) // Immutable
// Mutable versions are MutableMap and mutableMapOf
// Popular methods
names["Alice"] // 1
names["Bob"] = 4 // Assign
names.remove("Bob")
names.count()
names.containsKey("Bravo") // false
names.keys // List of keys
names.values // List of values
if ("Alice" in names.keys) {
    //
}
if (2 in names.values) {
    //
}
```
