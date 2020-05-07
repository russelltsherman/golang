# Open/closed principle

So, what causes the second test to fail?

To uncover the reason behind the failed test, we need to dig a bit deeper into how Go methods work under the hood.

Go methods are nothing more than syntactic sugar for invoking a function with an object instance as an argument (also known as a receiver).

In the sword.go file, String is always invoked with a Sword receiver and as a result, the call to the Damage method always gets dispatched to the implementation that's been defined by Sword type.

This is a prime example of the closed principle in action: Sword is not aware of any type that may embed it and its set of methods cannot be altered by objects it is embedded into. It is important to point out that while the EnchantedSword type cannot modify the implementation of the methods that have been defined on the embedded Sword instance, it can still access and mutate any fields defined by it (including private ones if both types are defined in the same package).
