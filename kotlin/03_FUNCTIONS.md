### Functions
Simple Functions
```kotlin
fun sum(x: Int, y: Int): Int {
    return x + y
}
func sum(x: Int, y: Int) = x + y // Single expression function
```
Named arguments
```kotlin
fun sum(x: Int, y: Int): Int {
    return x + y
}
sum(y = 1, x = 2)
```

Default parameter values
```kotlin
fun sum(x: Int, y: Int = 1): Int {
    return x + y
}
sum(1) // 2
sum(1, 2) // 3
```

Functions without return type (Unit)
```kotlin
fun sum(x: Int, y: Int) {
    println(x + y)
}
```

Lambda functions
```kotlin
val sum = { x: Int, y: Int -> x + y }
sum(1, 2) // 3
```

Function type for passing, returning functions
```kotlin
(Int, Int) -> Int // Function type

```

Example
```kotlin
fun toSeconds(time: String): (Int) -> Int = when (time) {
    "hour" -> { value -> value * 60 * 60 }
    "minute" -> { value -> value * 60 }
    "second" -> { value -> value }
    else -> { value -> value }
}

fun main() {
    val timesInMinutes = listOf(2, 10, 15, 1)
    val min2sec = toSeconds("minute")
    val totalTimeInSeconds = timesInMinutes.map(min2sec).sum()
    println("Total time is $totalTimeInSeconds secs")
}
```
