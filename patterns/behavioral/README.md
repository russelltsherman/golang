# Behavioral Design Patterns

The entire family of behavior patterns revolves around the principle of “object composition rather than inheritance”. This principle is currently been widely adopted by the developer community because of the benefits it offers. What this principle states is – instead of extending from an existing class, design your class to refer an existing class that you want to use. In Java, it’s done by declaring an instance variable referencing the object of the existing class. Then, by initializing the object either through the constructor or a method, and finally using the referred (composed) object.

Behavioral patterns address responsibilities of objects in an application and how they communicate between them. The Mediator pattern allows the loose coupling between set of objects in an application by handling the interactions between the objects.

Chain of Responsibilit:. Delegates commands to a chain of processing objects.
Command: Creates objects which encapsulate actions and parameters.
Interpreter: Implements a specialized language.
Iterator: Accesses the elements of an object sequentially without exposing its underlying representation.
Mediator: Allows loose coupling between classes by being the only class that has detailed knowledge of their methods.
Memento: Provides the ability to restore an object to its previous state.
Observer: Is a publish/subscribe pattern which allows a number of observer objects to see an event.
State: Allows an object to alter its behavior when its internal state changes.
Strategy: Allows one of a family of algorithms to be selected on-the-fly at run-time.
Template Method: Defines the skeleton of an algorithm as an abstract class, allowing its sub-classes to provide concrete behavior.
Visitor: Separates an algorithm from an object structure by moving the hierarchy of methods into one object.
