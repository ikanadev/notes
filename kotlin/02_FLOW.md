### Conditionals
We can use `when` and `if` to conditionally execute code
```kotlin
val check = true
if (check) {
    //
} else {
    //
}
val result = if (check) 1 else 2
```

```kotlin
val greet = "Hello"
val result = when (greet) {
    "Hello" -> 1
    "Goodbye" -> 2
    else -> 3
}
println(result) // 1
```

### Loops
#### Ranges
```kotlin
1..4 // 1, 2, 3, 4
1..<4 // 1, 2, 3
4 downTo 1 // 4, 3, 2, 1
1..5 step 2 // 1, 3, 5
```
#### For
```kotlin
for (i in 1..5) {
    println(i) // 1, 2, 3, 4, 5
}
val names = listOf("Alice", "Bob", "Charlie")
for (name in names) {
    println(name)
}
```
#### While
```kotlin
var i = 1
while (i < 10) {
    println(i)
    i++
}
```
