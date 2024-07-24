# Builder Pattern

The Builder pattern is a classic Gang of Four creational design pattern. This pattern, similar to the other creational patterns, such as factory method and abstract factory, is concerned with the creation of objects. But why do we need another pattern to create objects? To answer this, let us first look at a problem scenario.

## The Problem

Consider that you need to build a house for a customer. Now, building a house will consist a series of steps. You will start with the foundation, then the structure, and finally the roof. But, how will you do this in the programming world? This should be simple in Java- Right? You create a House class with the required fields and initialize them through a constructor, like this.

```java
public House(String foundation, String structure, String roof){
    this.foundation = foundation;
    this.structure = structure;
    this.roof = roof;
}
```

This builds the house and your customer is happy. But the next customer wants the house to be painted and furnished while another wants it only painted. You now need to revisit the House class with a set of overloaded constructors, something like this.

```java
public House(String foundation, String structure, String roof) {...}
public House(String foundation, String structure, String roof, boolean painted){...}
public House(String foundation, String structure, String roof, boolean painted, boolean furnished){...}
```

Although this will work, we have a flawed design. What we applied here is the telescopic constructor pattern, which is considered an anti-pattern. Though this pattern works for simple object creation with a limited number of fields, it becomes unmanageable for complex object creation. Imagine the number of constructors you will need to create a more complex house with options for plumbing, lightning, interiors, and so on. Another major problem in this design is related to constructor calls that clients need to make. It is hard and error prone to write client code when there are multiple constructors, each with a large set of parameters. In addition, readability is a major issue with such client code.

While writing such client code, you will often end up with questions, such as:

- Which constructor should I invoke?
- What will be the default values of the parameters if I don’t provide?
- Does the first boolean in the constructor represents painting or furnishing?

One solution to the telescopic constructor pattern is to follow JavaBeans conventions by writing setter methods instead of a set of constructors to initialize the fields.

```java
public void setFoundation(String foundation) {
        this.foundation = foundation;
    }
    public void setStructure(String structure) {
        this.structure = structure;
    }
    public void setRoof(String roof) {
        this.roof = roof;
    }
    public void setFurnished(boolean furnished) {
        this.furnished = furnished;
    }
    public void setPainted(boolean painted) {
        this.painted = painted;
    }
```

Clients can now call the setter methods with appropriate parameters to create House objects. Also, client code are now more readable and therefore have lesser chances of errors creeping in.

Your house building business is growing and everything is going fine until a customer calls up and complains that his house collapsed during construction. On examining, you found this particular client code.

```java
House house=new House();
house.setBasement("Concrete, brick, and stone");
house.setRoof("Concrete and reinforced steel");
house.setStructure("Concrete, mortar, brick, and reinforced steel");
house.setFurnished(true);
house.setPainted(true);
```

As you can see, the client code tried building the roof before the structure was in place, which means that the steps to build a house was not in the correct order. Another problem is the client having an instance of the House class in an inconsistent state. This means, if a client wants to create a House object with values for all its fields then the object will not have a complete state until all the setter methods have been called. As a result, some part of the client application might see and use a House object assuming that is already constructed while that is actually not the case.

While you might be still pondering over the existing problems on hand, imagine that a customer calls up with a requirement for a prefabricated house, another customer for a tree house, and yet another for an Igloo (a snow house). Now, here is a whole new set of problems to solve.

At this point you should consider yourself lucky because other people have faced similar problems and have come up with proven solutions. It is time to learn the classic GoF Builder pattern.

## Introduction to the Builder Pattern

The builder pattern allows you to enforce a step-by-step process to construct a complex object as a finished product. In this pattern, the step-by-step construction process remains same but the finished products can have different representations. In the context of the house building example, the step-by-step process includes the steps to create the foundation, structure, and roof followed by the steps to paint and furnish a house and these steps remain the same irrespective of the type of house to build. The finished product, which is a house, can have different representations. That is, it can be a concrete house, a prefabricated house, or a tree house.

## Participants in the Builder Pattern

To understand how the builder pattern works, let us solve the problems of our house building example. The main problem was that we expected the clients to perform the steps to construct a house and that too in the correct order. So, how will we address this in real life? We will hire a construction engineer who knows the process to construct houses.

The second problem was that we require different types of houses, such as concrete, prefabricated, tree house, and even Igloos. So next, we will hire builders (contractors) who specializes in building specific types of houses. A builder knows how to put things together with actual building materials, components, and parts to build a particular type of house. For example, a concrete house builder knows how to build the structure of a concrete house using concrete, brick, and stone. Similarly, a prefabricated house builder knows how to build the structure of a prefabricated house using structural steels and wooden wall panels. So from now on, whenever we need a house, the construction engineer will direct a builder to build the house.

In our application, we can model the construction engineer by creating a ConstructionEngineer class. Then we can model the builders by first creating a HouseBuilder interface and then builder classes, such as ConcreteHouseBuilder and PrefabricatedHouseBuilder that implement the HouseBuilder interface. Here, notice that we have added a layer of abstraction by providing an interface (HouseBuilder). This is because we do not want our construction engineer to be tied with a particular builder. The construction engineer should be able to direct any builder that implements the HouseBuilder interface to build a house. This will also allow us to later add new builders without making changes to the existing application code.

We can now summarize the components of the builder pattern in the context of the house building example as:

- Product (House): A class that represents the product to create.
- Builder (HouseBuilder): Is an interface to build the parts of a product.
- ConcreteBuilder(ConcreteHouseBuilder and PrefabricatedHouseBuilder): Are concrete classes that implement Builder to construct and assemble parts of the product and return the finished product.
- Director (ConstructionEngineer): A class that directs a builder to perform the steps in the order that is required to build the product.

## Applying the Builder Pattern

To apply the builder pattern to the house building example, let us first create the product that the builders will construct.

House.java

```java
package guru.springframework.gof.builder.product;
public class House {
private String foundation;
private String structure;
private String roof;
private boolean furnished;
private boolean painted;
public void setFoundation(String foundation) {
this. foundation = foundation;
}
public void setStructure(String structure) {
this.structure = structure;
}
public void setRoof(String roof) {
this.roof = roof;
}
public void setFurnished(boolean furnished) {
this.furnished = furnished;
}
public void setPainted(boolean painted) {
this.painted = painted;
}
@Override
public String toString() {
return "Foundation - " + foundation + " Structure - " + structure + " Roof - " + roof +" Is Furnished? "+furnished+" Is Painted? "+painted;
}
}
```

In the example above, we wrote a House class with five fields and their corresponding setter methods. Next, we will create the HouseBuilder interface, which is the Builder in the application. HouseBuilder.java

```java
package guru.springframework.gof.builder.builders;
import guru.springframework.gof.builder.product.House;
public interface HouseBuilder {
void buildFoundation();
void buildStructure();
void buildRoof();
void paintHouse();
void furnishHouse();
House getHouse();
}
```

In the example above, we wrote the HouseBuilder interface to declare five methods to create the parts of the product (House). We also declared a getHouse() method that returns the finished product. We will provide the implementation of the methods in the concrete subclasses: ConcreteHouseBuilder and PrefabricatedHouseBuilder, which are the ConcreteBuilder components in the application.

ConcreteHouseBuilder.java

```java
package guru.springframework.gof.builder.builders;
import guru.springframework.gof.builder.product.House;
public class ConcreteHouseBuilder implements HouseBuilder{
private House house;
public ConcreteHouseBuilder() {
this.house = new House();
}
@Override
public void buildFoundation() {
house.setFoundation("Concrete, brick, and stone");
System.out.println("ConcreteHouseBuilder: Foundation complete...");
}
@Override
public void buildStructure(){
house.setStructure("Concrete, mortar, brick, and reinforced steel");
System.out.println("ConcreteHouseBuilder: Structure complete...");
}
@Override
public void buildRoof(){
house.setRoof("Concrete and reinforced steel");
System.out.println("ConcreteHouseBuilder: Roof complete...");
}
@Override
public void paintHouse(){
house.setPainted(true);
System.out.println("ConcreteHouseBuilder: Painting complete...");
}
@Override
public void furnishHouse(){
house.setFurnished(true);
System.out.println("ConcreteHouseBuilder: Furnishing complete...");
}
public House getHouse() {
System.out.println("ConcreteHouseBuilder: Concrete house complete...");
return this.house;
}
}
```

PrefabricatedHouseBuilder.java

```java
package guru.springframework.gof.builder.builders;
import guru.springframework.gof.builder.product.House;
public class PrefabricatedHouseBuilder implements HouseBuilder{
private House house;
public PrefabricatedHouseBuilder() {
this.house = new House();
}
@Override
public void buildFoundation() {
house.setFoundation("Wood, laminate, and PVC flooring");
System.out.println("PrefabricatedHouseBuilder: Foundation complete...");
}
@Override
public void buildStructure(){
house.setStructure("Structural steels and wooden wall panels");
System.out.println("PrefabricatedHouseBuilder: Structure complete...");
}
@Override
public void buildRoof(){
house.setRoof("Roofing sheets");
System.out.println("PrefabricatedHouseBuilder: Roof complete...");
}
@Override
public void paintHouse(){
house.setPainted(false);
System.out.println("PrefabricatedHouseBuilder: Painting not required...");
}
@Override
public void furnishHouse(){
house.setFurnished(true);
System.out.println("PrefabricatedHouseBuilder: Furnishing complete...");
}
public House getHouse() {
System.out.println("PrefabricatedHouseBuilder: Prefabricated house complete...");
return this.house;
}
}
```

In the above examples, we first wrote the ConcreteHouseBuilder class. In the constructor of this class, we created a House object. We then implemented the methods declared in the HouseBuilder interface to create the parts of a concrete house through calls to the setter methods of the House object. Finally, we implemented the getHouse() method to return the final House object that represents a concrete house. Similarly, we wrote the PrefabricatedHouseBuilder class to create the parts of a prefabricated house and return the final House object that represents a prefabricated house.

With these two classes in place, we are almost ready to “create different representations” of a house: concrete and prefabricated. But, we are yet to define the “same construction process“. We will do it next in the ConstructionEngineer class, which is the Director in the application.

ConstructionEngineer.java

```java
package guru.springframework.gof.builder.director;
import guru.springframework.gof.builder.builders.HouseBuilder;
import guru.springframework.gof.builder.product.House;
public class ConstructionEngineer {
private HouseBuilder houseBuilder;
public ConstructionEngineer(HouseBuilder houseBuilder){
this.houseBuilder = houseBuilder;
}
public House constructHouse() {
this.houseBuilder.buildFoundation();
this.houseBuilder.buildStructure();
this.houseBuilder.buildRoof();
this.houseBuilder.paintHouse();
this.houseBuilder.furnishHouse();
return this.houseBuilder.getHouse();
}
}
```

In the above example, we wrote the ConstructionEngineer class with a constructor that accepts a HouseBuilder object. In the constructHouse() method, we made a series of calls on the HouseBuilder object in a certain order and returned the final House object to the caller. Notice that the ConstructionEngineer class is not tied to any concrete builder. Also, this class uses the same construction process in the constructHouse() method irrespective of the type of concrete builder provided to it at run time. This allows us to add new concrete builder classes without making any changes to the construction process. Now that our house building example is ready, let us write a unit test to observe the builder pattern at work. ConstructionEngineerTest.java

```java
package guru.springframework.gof.builder.director;
import guru.springframework.gof.builder.builders.HouseBuilder;
import guru.springframework.gof.builder.builders.ConcreteHouseBuilder;
import guru.springframework.gof.builder.builders.PrefabricatedHouseBuilder;
import guru.springframework.gof.builder.product.House;
import org.junit.Test;
public class ConstructionEngineerTest {
@Test
public void testConstructHouse() throws Exception {
HouseBuilder concreteHouseBuilder = new ConcreteHouseBuilder();
ConstructionEngineer engineerA = new ConstructionEngineer(concreteHouseBuilder);
House houseA = engineerA.constructHouse();
System.out.println("House is: "+houseA);
PrefabricatedHouseBuilder prefabricatedHouseBuilder = new PrefabricatedHouseBuilder();
ConstructionEngineer engineerB = new ConstructionEngineer(prefabricatedHouseBuilder);
House houseB = engineerB.constructHouse();
System.out.println("House is: "+houseB);
}
}
```

As you can see in the example above, a client is now insulated from the object creation process. A client only needs to provide the Director a ConcreteBuilder to use. It is the responsibility of the Director to instruct the ConcreteBuilder on the construction process and the ConcreteBuilder in turn will create the finished product. Finally, the client receives the finished product from the Director.
When you run the code above, you will see this output:

## T E S T S

Running guru.springframework.gof.builder.director.ConstructionEngineerTest
ConcreteHouseBuilder: Foundation complete...
ConcreteHouseBuilder: Structure complete...
ConcreteHouseBuilder: Roof complete...
ConcreteHouseBuilder: Painting complete...
ConcreteHouseBuilder: Furnishing complete...
ConcreteHouseBuilder: Concrete house complete...
House is: Foundation - Concrete, brick, and stone Structure - Concrete, mortar, brick, and reinforced steel Roof - Concrete and reinforced steel Is Furnished? true Is Painted? true
PrefabricatedHouseBuilder: Foundation complete...
PrefabricatedHouseBuilder: Structure complete...
PrefabricatedHouseBuilder: Roof complete...
PrefabricatedHouseBuilder: Painting not required...
PrefabricatedHouseBuilder: Furnishing complete...
PrefabricatedHouseBuilder: Prefabricated house complete...
House is: Foundation - Wood, laminate, and PVC flooring Structure - Structural steels and wooden wall panels Roof - Roofing sheets Is Furnished? true Is Painted? false
Tests run: 1, Failures: 0, Errors: 0, Skipped: 0, Time elapsed: 2.119 sec - in guru.springframework.gof.builder.director.ConstructionEngineerTest

## Conclusion

If you are familiar with the abstract factory pattern, you might have observed that both the abstract factory and builder patterns are similar, as both can be used to abstract object creation. But there are distinct differences between the two. While abstract factory emphasizes on creating a family of related objects in one go, builder is about creating an object through a step-by-step construction process and returning the object as the final step. In short abstract factory is concerned with what is made, while the builder with how it is made. So as you go further into enterprise application development, whenever you need to create complex objects independently of the construction algorithm turn to the classic GoF Builder Pattern!
