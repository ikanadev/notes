### Classes
We can define the properties of a class in the constructor (arguments), and the methods in the body
```kotlin
class Contact(val id: Int, var email: String = "example@gmail.com") {
    val category: String = "work"

    fun printCategory() {
        println(category)
    }
}

val contact = Contact(1)
contact.printCategory() // 1
```

#### Data classes
Used to store data, and automatically implements the `equals`, `hashCode`, `toString`, `copy` and `componentN` functions
```kotlin
data class Contact(val id: Int, var name: String)

val contact = Contact(1, "Alice")
var contact2 = contact.copy(id = 2)
contact.toString() // Contact(id=1, name=Alice)
```

#### Extension functions
```kotlin
fun String.lastChar(): Char = this.get(this.length - 1)
"Hello".lastChar() // 'o'
```

### Scope functions

#### let
```kotlin
val someVar: String? = someFunctionThatReturnsStringOrNull()
someVar?.let { it -> println(it) } // it holds the not null value of someVar
someVar?.let { println(it) }
```

#### apply
It helps to initialize objects at the time of creation
```kotlin
class Contact() {
    var name: String? = null
    var email: String? = null
    fun printName() {
        name?.let { println(it) }
    }
}

val contact = Contact().apply {
    name = "Alice"
    email = "example@gmail.com"
    printName()
}
```

#### run
Allows to execute a block of code on an object and return something
```kotlin
class Contact() {
    var name: String? = null
    fun getName(): String {
        return name ?: "Unknown"
    }
}

val contact = Contact().apply {
    name = "Alice"
}
val bobName = contact.run {
    name = "Bob"
    getName()
}
```

#### also
```kotlin
Allows to do "something" with the object and return it
val medals: List<String> = listOf("Gold", "Silver", "Bronze")
val newMedals: List<String> =
    medals
        .map { it.uppercase() }
        .also { println(it) }
        .filter { it.length > 4 }
        .also { println(it) }
        .reversed()
println(newMedals)
```

#### with
Allow to call multiple functions on an object
```kotlin
class Messages() {
    fun hello() {
        println("Hello")
    }
    fun goodbye() {
        println("Goodbye")
    }
    fun welcome() {
        println("Welcome")
    }
}

val messages = Messages()
with(messages) {
    hello()
    goodbye()
    welcome()
}
```
