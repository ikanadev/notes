### Objects
Classes where you create the class and the instance at the same time

#### Declaration
```kotlin
// data object MyObject { <- we can have a data object (that comes with .equals() and .toString() methods)
object MyObject {
    var myProperty: Int = 0
    fun myFunction() {}
}
MyObject.myFunction()
```

#### Companion Objects
Objects that are inside a class
```kotlin
class MyClass {
    companion object {
        fun myFunction() {}
    }
}
MyClass.myFunction()
```
