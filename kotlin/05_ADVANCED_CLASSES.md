### Inheritance
By default kotlin classes can't be inherited. Koling classes  only support single inheritance
The parent class is called `parent`, all classes inherit from the `Any` parent class.

#### Abstract classes
Abstract classes can be inherited by default, they have a constructor but can't be instanced, the
child classes are expected to define the behavior of the parent's properties and functions by using the 
`override` keyword

```kotlin
abstract class Product(val name: String, val price: Double) {
    abstract val category: String

    fun productInfo(): String {
        return "Product: $name, Category: $category, Price: $price"
    }
}

class Electronic(name: String, price: Double, val warranty: Int): Product(name, price) {
    override val category = "Electronic"
}
```

#### Interfaces
Similar to classes but you:
- Can't create an instance of an interface
- The functions and properties are implicitly inheritable
- You don't need to mark their functions as abstract

```kotlin
interface PaymentMethod {
    fun initiatePayment(amount: Double): String
}
class CreditCardPayment(val cardNumber: String, val cardHolderName: String, val expiryDate: String): PaymentMethod {
    override fun initiatePayment(amount: Double): String {
        return "Payment of $$amount initiated using Credit card ending in ${cardNumber.takeLast(4)}"
    }
}
class SomeClass(): InterfaceClass1, InterfaceClass2 // multiple interface inheritance
```

#### Delegation
You can delegate the interface implementation to a class intance using the `by` keyword.
```kotlin
class CanvasSession(val tool: DrawingTool): DrawingTool by tool { // tool is a class instance that also implements the DrawingTool interface
    override val color: String = "blue" // we only re-implement whathever we need to change
}
```

