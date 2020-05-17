# Iterator Pattern

The Behavioral pattern family of the Gang of Four design patterns addresses responsibilities of objects in an application and how they communicate between them. The Iterator Pattern is one of the patterns in this family. We have already learned about the Command and Chain of Responsibility Behavioral patterns. In this post we will learn about the Iterator pattern.

## Iterator Pattern: Introduction

As a Java programmer, it’s likely you have worked with aggregate objects, which are referred as collections in Java. A collection is an object that acts as a container to store other objects, which are referred as elements. Based on how a collection stores elements, there are different types of collections. For example, some collections may allow duplicate elements while others may not. ArrayList, LinkedList, HashMap, LinkedHashSet are few of the collections that are part of the Java Collection Framework. Like any other objects, a collection can be passed around the application or stored in other collection objects.

A collection at the minimum needs to provide clients methods to add and remove elements from it. But more importantly, it needs to allow clients to traverse through the elements it stores. A collection class can itself can implement the functionality required to provide access and allow traversal through its elements. But by doing so, the collection’s underlying structure and implementation will be exposed to clients. This is a bad object-oriented design principle that does not follow encapsulation. In addition, if you go back and revisit the SOLID design principles, it will be apparent that implementing element access and traversal operations in the collection itself is a clear violation of the Single Responsibility Principle. A collection should be only responsible to act as a container for storing elements, and any other responsibility such as element access and traversal should not be part of it.

The Iterator pattern addresses such recurring problems when dealing with aggregate objects. What this pattern says is that aggregate objects should provide a way to access its elements sequentially without exposing its internal structure. To accomplish the intent of the pattern, we need to separate the responsibility for access and traversal of the elements stored in the aggregate object and assign it to another object, which is referred as the iterator. The iterator keeps track of the elements and can perform different types of traversals sequentially based on what you want to accomplish. It is important to decouple the iterator object from the aggregate object so it can be reused for traversing other types of aggregate objects. You can design your program in such a way that the client asks the aggregate objects for its iterator, and the aggregate object returns the iterator initialized with the data structure holding the elements. The client then uses the iterator to traverse through the elements.

## Participants of the Iterator Pattern

Consider a situation where you need to store information of the GoF design patterns in an application that other components will be using either to print out information about a pattern or publish it on a Web page. For this requirement we can model a DesignPattern class to represent a design pattern. Next, we will start with an interface for the collection that will store DesignPattern objects. Let’s, name it PatternAggregate. We can then model a concrete class, named as PatternAggregateImpl that will implement PatternAggregate. Objects of this class will be responsible for storing DesignPattern elements and creating its iterator. Next, we create an interface for the iterator as PatternIterator to declare the methods for traversing the elements. We will implement the methods in the concrete sub class PatternIteratorImpl.

In the context of our example, let’s summarize the participants of the Iterator pattern.

- Aggregate (PatternAggregate): Is an interface that declares the methods to create and return an iterator.
- ConcreteAggregate (PatternAggregateImpl): Is a concrete class that implements the Aggregate interface to create and return an iterator.
- Iterator (PatternIterator): Is an interface with methods to allow clients to access and traverse elements.
- ConcreteIterator (PatternIteratorImpl): Is a concrete class that implements the Iterator interface. Objects of this class keeps track of the elements and implements access and traversal operations on the elements.

## Applying the Iterator Pattern

Let’s work on a rudimentary implementation of the Iterator Pattern. We will start with the DesignPattern class whose objects will be the elements that we will store in a ConcreteAggregate and iterate through them.

DesignPattern.java

```java
package guru.springframework.gof.iterator;
public class DesignPattern {
private String patternType;
private String patternName;
public DesignPattern(String patternType, String patternName){
this.patternType=patternType;
this.patternName=patternName;
}
public String getPatternType() {
return patternType;
}
public String getPatternName() {
return patternName;
}
}
```

We have kept the DesignPattern class simple with two instance variables to hold the type and name of a pattern. We initialized the variables in the constructor and provided public getter methods to access them. Next, we will write the Aggregate – the PatternAggregate interface.

PatternAggregate.java

```java
package guru.springframework.gof.iterator;
public interface PatternAggregate {
void addPattern(DesignPattern designPattern);
void removePattern(DesignPattern designPattern);
PatternIterator getPatternIterator();
}
```

In the interface we wrote above, we declared three methods. The addPattern() and removePattern() methods are self-explanatory. Their implementations add and remove DesignPattern elements to and from the ConcreteAggregate. The method we are interested in is getPatternIterator(). We will write its implementations along with the other two methods in the ConcreteAggregate – the PatternAggregateImpl class.

PatternAggregateImpl.java

```java
package guru.springframework.gof.iterator;
import java.util.ArrayList;
import java.util.List;
public class PatternAggregateImpl implements PatternAggregate{
List patternList;
public PatternAggregateImpl(){
patternList=new ArrayList();
}
@Override
public void addPattern(DesignPattern designPattern){
patternList.add(designPattern);
}
@Override
public void removePattern(DesignPattern designPattern){
patternList.remove(designPattern);
}
@Override
public PatternIterator getPatternIterator(){
return new PatternIteratorImpl(patternList);
}
}
```

In the PatternAggregateImpl class above, we internally used a List implementation of type ArrayList. In the overridden addPattern() and removePattern() methods, we added and removed the DesignPattern object passed to the methods to the List implementation. In the overridden getPatternIterator() method we instantiated the PatternIteratorImpl class initialized with the List implementation, and returned the PatternIteratorImpl object to the caller. At this point, we are yet to implement the iterator, which we are coming to next. We will implement the iterator code starting with the Iterator – the PatternIterator interface.

PatternIterator.java

```java
package guru.springframework.gof.iterator;
public interface PatternIterator {
DesignPattern nextPattern();
boolean isLastPattern();
}
```

In the PatternIterator interface above, we declared two methods – nextPattern() and isLastPattern() to implement forward traversal of elements sequentially.

Let’s now write the ConcreteIterator – the PatternIteratorImpl class.

PatternIteratorImpl.java

```java
package guru.springframework.gof.iterator;
import java.util.List;
public class PatternIteratorImpl implements PatternIterator{
public List patternList;
int position;
DesignPattern designPattern;
public PatternIteratorImpl(List patternList){
this.patternList=patternList;
}
@Override
public DesignPattern nextPattern(){
System.out.println("Returning pattern at Position: "+position);
designPattern=(DesignPattern)patternList.get(position);
position++;
return designPattern;
}
@Override
public boolean isLastPattern(){
if(position< patternList.size()){
return false;
}
return true;
}
}
```

In the overridden nextPattern() method of the PatternIteratorImpl class above, we retrieved an element (DesignPattern object) from the List implementation by providing an index stored in the position variable and then incremented the position by one. We also implemented the overridden isLastPattern() method to check whether the iterator has reached the last element in the List implementation. Client will typically call this method before calling nextPattern() to avoid exception of type IndexOutOfBoundsException. Let’s now write some test code and observe how the Iterator pattern works.

PatternAggregateImplTest.java

```java
package guru.springframework.gof.iterator;
import org.junit.Test;
import static org.junit.Assert.\*;
public class PatternAggregateImplTest {
@Test
public void testPatternIterator() throws Exception {
DesignPattern pattern1 = new DesignPattern("Creational", "Factory Method");
DesignPattern pattern2 = new DesignPattern("Creational", "Abstract Factory");
DesignPattern pattern3 = new DesignPattern("Structural", "Adapter");
DesignPattern pattern4 = new DesignPattern("Structural", "Bridge");
DesignPattern pattern5 = new DesignPattern("Behavioral", "Chain of Responsibility");
DesignPattern pattern6 = new DesignPattern("Behavioral", "Iterator");
PatternAggregate patternAggregate = new PatternAggregateImpl();
patternAggregate.addPattern(pattern1);
patternAggregate.addPattern(pattern2);
patternAggregate.addPattern(pattern3);
patternAggregate.addPattern(pattern4);
patternAggregate.addPattern(pattern5);
patternAggregate.addPattern(pattern6);
System.out.println("-----Pattern list-----");
printPatterns(patternAggregate);
patternAggregate.removePattern(pattern1);
patternAggregate.removePattern(pattern2);
System.out.println("-----Pattern list after removal operation-----");
printPatterns(patternAggregate);
}
public void printPatterns(PatternAggregate patternAggregate){
PatternIterator patternIterator= patternAggregate.getPatternIterator();
while(!patternIterator.isLastPattern()){
DesignPattern designPattern=patternIterator.nextPattern();
System.out.println(designPattern.getPatternType() + " : " + designPattern.getPatternName());
}
}
}
```

In the test class above, From Line 11 – Line 16, we created six DesignPattern objects. In Line 18 we created a PatternAggregateImpl object and then from Line 19 – Line 24 we made calls to the addPattern() method to add each of our DesignPattern object to PatternAggregateImpl. In Line 27 we called the printPatterns() method passing the PatternAggregate object filled with DesignPattern elements. We will come to the printPatterns() method where the actual traversal happens a bit later.
Next, from Line 29 – Line 30 we called the removePattern() method twice to remove two DesignPattern elements from PatternAggregateImpl. In Line 33, we again called the printPatterns() method. In this method, we started by calling the getPatternIterator() method on the PatternAggregate implementation passed as a method parameter to printPatterns(). The getPatternIterator() method returns the PatternIterator object. We then used a while loop from Line 37 – Line 41 to check whether the iterator is positioned on the last element. For each pass of the while loop, we called the nextPattern() method to retrieve the next DesignPattern element and printed out its type and name. The while loop continues to execute until the last element is reached on which the while loop evaluates to false.

The output of the test is this.

---

## T E S T S

Running guru.springframework.gof.iterator.PatternAggregateImplTest
-----Pattern list-----
Returning pattern at Position: 0
Creational : Factory Method
Returning pattern at Position: 1
Creational : Abstract Factory
Returning pattern at Position: 2
Structural : Adapter
Returning pattern at Position: 3
Structural : Bridge
Returning pattern at Position: 4
Behavioral : Chain of Responsibility
Returning pattern at Position: 5
Behavioral : Iterator
-----Pattern list after removal operation-----
Returning pattern at Position: 0
Structural : Adapter
Returning pattern at Position: 1
Structural : Bridge
Returning pattern at Position: 2
Behavioral : Chain of Responsibility
Returning pattern at Position: 3
Behavioral : Iterator
Tests run: 1, Failures: 0, Errors: 0, Skipped: 0, Time elapsed: 0.346 sec - in guru.springframework.gof.iterator.PatternAggregateImplTest

## Summary

The Collection Framework of the Java SE platform makes use of the Iterator pattern. It provides the Iterable interface that the collection classes in the java.util package implements to return an Iterator object for traversing collection elements. In most cases, the built in collection classes with their capabilities to return an iterator will get your work done. But, during enterprise application development, there might be requirements to implement custom collections. For example, in a Spring MVC application, you might need a collection to store Employee objects. Naturally, you’ll use the built in collection classes of Java. But, imagine that the requirement also states that the collection should only store Employee objects having Manager as the designation. We can create a collection class for this requirement. While creating the class, remember that some other components of the application will require traversing the elements of your collection class. By now, as your design toolkit contains the Iterator pattern, put it at work to step through the elements without exposing the internal representation of your collection. The Spring Framework also extends the Iterator pattern through the CompositeIterator class. This class implements the Java SE Iterator interface to maintain multiple other iterators which are invoked in sequence until all iterators have completed their operations.
