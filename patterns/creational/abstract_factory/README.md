# Abstract Factory Design Pattern

The abstract factory pattern is one of the classic Gang of Four creational design patterns used to create families of objects, where the objects of a family are designed to work together. In the abstract factory pattern you provide an interface to create families of related or dependent objects, but you do not specify the concrete classes of the objects to create. From the client point of view, it means that a client can create a family of related objects without knowing about the object definitions and their concrete class names.

It is easy to confuse the abstract factory pattern with the factory method pattern because both design patterns deal with object creation. Both the patterns advocates the Object Oriented Programming (OOP) principle “Program to an interface, not an implementation” to abstract how the objects are created. Both design patterns help in creating client code that is loosely-coupled with object creation code, but despite the similarities, and the fact that both the patterns are often used together, they do have distinct differences.

## Abstract Factory Pattern vs Factory Method Pattern

Abstract factory adds another level of abstraction to factory method. While factory method abstracts the way objects are created, abstract factory abstracts how the factories are created. The factories in turn abstracts the way objects are created. You will often hear the abstract factory design pattern referred to as a “factory of factories“.

From implementation point of view, the key difference between the factory method and abstract factory patterns is that factory method is just a method to create objects of a single type, while abstract factory is an object to create families of objects.

Another difference is that the factory method pattern uses inheritance while the abstract factory pattern uses composition. We say that that factory method uses inheritance because this pattern relies on a subclass for the required object instantiation. Recall in the Factory Method Design Pattern post where we created a `createPizza()` factory method in an abstract base class and implemented the factory method in a PizzaFactory subclass for the required Pizza object instantiation. On the other hand, the abstract factory pattern delegates responsibility to a separate object (abstract factory) dedicated to create a family of related objects. Then, through composition, the abstract factory object can be passed to the client who will use it (instead of factory method) to get the family of related objects.

## Participants in the Abstract Factory Pattern

To understand how the abstract factory pattern works, let us revisit the pizza store that we developed in the Factory Method Design Pattern. The store has seen a rapid increase in its customer base and now wants to serve their existing types of pizzas: cheese, pepperoni, and veggie in two different topping styles: Sicilian topping and Gourmet topping. Each topping style will require a different combination of products. Sicilian topping will have Goat Cheese with Tomato Sauce while Gourmet topping will have Mozzarella Cheese with California Oil Sauce. To model the new requirements of the application, we can create the concrete Cheeseproducts: GoatCheese and MozzarellaCheese and the concrete Sauce products: TomatoSauce and CaliforniaOilSauce. Now, as we do not want any client to directly instantiate the products, we will abstract the way the products are created by introducing an abstract factory. We will create a BaseToppingFactory abstract factory class and let its two concrete subclasses: SicillianToppingFactory and GourmetToppingFactory create our products. Here, it is important to note that an abstract factory should be designed to create families of products. So we can model SicillianToppingFactory to create the family of MozzarellaCheese and TomatoSauce products and GourmetToppingFactory to create the family of GoatCheese and CaliforniaOilSauce products.

Now, let us summarize the components of the abstract factory pattern in the context of the enhanced pizza store:

- AbstractProduct (Cheese and Sauce): Is an interface or an abstract class whose subclasses are instantiated by the abstract factory objects.
- ConcreteProduct (GoatCheese, MozzarellaCheese, TomatoSauce, and CaliforniaOilSauce): Are the concrete subclasses that implement/extend AbstractProduct. The abstract factory objects instantiate these subclasses.
- AbstractFactory (BaseToppingFactory): Is an interface or an abstract class whose subclasses instantiate a family of AbstractProduct objects.
- ConcreteFactory (SicillianToppingFactory and GourmetToppingFactory): Are the concrete subclasses that implement/extend AbstractFactory. An object of this subclass instantiates a family of AbstractProduct objects.
- Client: Uses AbstractFactory to get AbstractProduct objects.

## Applying the Abstract Factory Pattern

To apply the abstract factory pattern in the pizza store application, let us first create the products that the factories will produce.

Cheese.java

```java
package guru.springframework.gof.abstractFactory.topping;
public interface Cheese {
void prepareCheese();
}
```

GoatCheese.java

```java
package guru.springframework.gof.abstractFactory.topping;
public class GoatCheese implements Cheese {
public GoatCheese(){
prepareCheese();
}
@Override
public void prepareCheese(){
System.out.println("Preparing goat cheese...");
}
}
```

MozzarellaCheese.java

```java
package guru.springframework.gof.abstractFactory.topping;
public class MozzarellaCheese implements Cheese{
public MozzarellaCheese(){prepareCheese();
}
@Override
public void prepareCheese() {
System.out.println("Preparing mozzarella cheese...");
}
}
```

Sauce.java

```java
package guru.springframework.gof.abstractFactory.topping;
public interface Sauce {
void prepareSauce();
}
```

TomatoSauce.java

```java
package guru.springframework.gof.abstractFactory.topping;
public class TomatoSauce implements Sauce {
public TomatoSauce(){
prepareSauce();
}
@Override
public void prepareSauce() {
System.out.println("Preparing tomato sauce..");
}
}
```

CaliforniaOilSauce.java

```java
package guru.springframework.gof.abstractFactory.topping;
public class CaliforniaOilSauce implements Sauce {
public CaliforniaOilSauce(){
prepareSauce();
}
@Override
public void prepareSauce() {
System.out.println("Preparing california oil sauce..");
}
}
```

In the above examples, we wrote the Cheese interface, which is an AbstractProduct. Then we wrote the GoatCheese and MozzarellaCheese classes, which are the ConcreteProduct to implement Cheese. Similarly for the pizza sauce, we wrote the Sauce interface and the TomatoSauce and CaliforniaOilSauce implementation classes.

Next, we will write the factories that will create the products. We will start with the abstract factory.

BaseToppingFactory.java

```java
package guru.springframework.gof.abstractFactory;
import guru.springframework.gof.abstractFactory.topping.Cheese;
import guru.springframework.gof.abstractFactory.topping.Sauce;
public abstract class BaseToppingFactory {
public abstract Cheese createCheese();
public abstract Sauce createSauce();
}
```

In the example above, we wrote the BaseToppingFactory abstract class, the abstract factory of our application. in the abstract factory, we declare the createCheese() and createSauce() abstract methods that return Cheese and Product objects respectively. As stated in the definition of abstract factory earlier, our abstract factory (BaseToppingFactory) is providing an “interface to create families of related or dependent objects“. The related objects here are Cheese and Sauce, both of which are together used to create toppings. The definition also states “…but you do not specify the concrete classes of the objects to create“. As you can notice in the BaseToppingFactory code, our abstract factory is not concerned with any of the concrete products: GoatCheese, MozzarellaCheese, TomatoSauce, or CaliforniaOilSauce. Let us now write the ConcreteFactory implementations.

SicillianToppingFactory.java

```java
package guru.springframework.gof.abstractFactory;
import guru.springframework.gof.abstractFactory.topping.Cheese;
import guru.springframework.gof.abstractFactory.topping.MozzarellaCheese;
import guru.springframework.gof.abstractFactory.topping.Sauce;
import guru.springframework.gof.abstractFactory.topping.TomatoSauce;
public class SicilianToppingFactory extends BaseToppingFactory{
@Override
public Cheese createCheese(){return new MozzarellaCheese();}
@Override
public Sauce createSauce(){return new TomatoSauce();}
}
```

GourmetToppingFactory.java

```java
package guru.springframework.gof.abstractFactory;
import guru.springframework.gof.abstractFactory.topping.CaliforniaOilSauce;
import guru.springframework.gof.abstractFactory.topping.Cheese;
import guru.springframework.gof.abstractFactory.topping.GoatCheese;
import guru.springframework.gof.abstractFactory.topping.Sauce;
public class GourmetToppingFactory extends BaseToppingFactory{
@Override
public Cheese createCheese(){return new GoatCheese();}
@Override
public Sauce createSauce(){return new CaliforniaOilSauce();}
}
```

In the above examples, we first wrote the SicillianToppingFactory class that provides implementation of the createCheese() method to create and return a MozzarellaCheese object and a createSauce() method to create and return a TomatoSauce object. Then, we wrote the GourmetToppingFactory class to create and return the GoatCheese and CaliforniaOilSauce objects. At this point let us understand the relationship of abstract factory with factory method. If you have noticed, the createCheese() and createSauce() are factory methods that we used in our abstract factory. In other words, we can say that an abstract factory object can use factory methods, one for each product to create. We are saying “can use” because, though this is the most common approach, it is not the only approach. Another approach is to use the Prototype pattern in an abstract factory to create products. Now that we have applied the abstract factory pattern to create the pizza topping factories, we will next update our pizza application to create pizzas using the pizza topping factories. First, let us create the pizza objects.

Pizza.java

```java
package guru.springframework.gof.abstractFactory.product;
public abstract class Pizza {
public abstract void addIngredients();
public void bakePizza() {
System.out.println("Pizza baked at 400 for 20 minutes.");
}
}
```

CheesePizza.java

```java
package guru.springframework.gof.abstractFactory.product;
import guru.springframework.gof.abstractFactory.BaseToppingFactory;
public class CheesePizza extends Pizza {
BaseToppingFactory toppingFactory;
public CheesePizza(BaseToppingFactory toppingFactory){
this.toppingFactory=toppingFactory;
}
@Override
public void addIngredients() {
System.out.println("Preparing ingredients for cheese pizza.");
toppingFactory.createCheese();
toppingFactory.createSauce();
}
}
```

PepperoniPizza.java

```java
package guru.springframework.gof.abstractFactory.product;
import guru.springframework.gof.abstractFactory.BaseToppingFactory;
public class PepperoniPizza extends Pizza {
BaseToppingFactory toppingFactory;
public PepperoniPizza(BaseToppingFactory toppingFactory)
{
this.toppingFactory=toppingFactory;
}
@Override
public void addIngredients() {
System.out.println("Preparing ingredients for pepperoni pizza.");
toppingFactory.createCheese();
toppingFactory.createSauce();
}
}
```

VeggiePizza.java

```java
package guru.springframework.gof.abstractFactory.product;
import guru.springframework.gof.abstractFactory.BaseToppingFactory;
public class VeggiePizza extends Pizza {
BaseToppingFactory toppingFactory;
public VeggiePizza(BaseToppingFactory toppingFactory)
{
this.toppingFactory=toppingFactory;
}
@Override
public void addIngredients() {
System.out.println("Preparing ingredients for veggie pizza.");
toppingFactory.createCheese();
toppingFactory.createSauce();
}
}
```

In the examples above, we modified the CheesePizza, PepperoniPizza, and VeggiePizza classes. Each pizza class is now composed of the abstract topping factory, BaseToppingFactory. The constructor of each class will initialize a pizza object with an abstract topping factory object at run time. Also, the concrete pizza classes override the preparePizza() method of the base Pizza class. In the preparePizza() method, the classes use the topping factory object (provided at run time via the constructor) to create the toppings. The concrete pizza classes are not tied to any concrete topping factory or any concrete topping product to use. This is because we have followed the basic principle to “Program to an interface, not an implementation“. Therefore, if we later introduce a new topping factory, say NeapolitanToppingFactory to produce toppings of FontinaCheese with ItalianPlumTomatoSauce, we do not require changing our pizza classes to use the new variety. Now that our pizza classes are ready, we will need the factories to create them.

BasePizzaFactory.java

```java
package guru.springframework.gof.abstractFactory;
import guru.springframework.gof.abstractFactory.product.Pizza;
public abstract class BasePizzaFactory {

    public abstract Pizza createPizza(String type);

}
```

SicilianPizzaFactory.java

```java
package guru.springframework.gof.abstractFactory;
import guru.springframework.gof.abstractFactory.product.CheesePizza;
import guru.springframework.gof.abstractFactory.product.PepperoniPizza;
import guru.springframework.gof.abstractFactory.product.Pizza;
import guru.springframework.gof.abstractFactory.product.VeggiePizza;
public class SicilianPizzaFactory extends BasePizzaFactory {
@Override
public Pizza createPizza(String type){
Pizza pizza;
BaseToppingFactory toppingFactory= new SicilianToppingFactory();
switch (type.toLowerCase())
{
case "cheese":
pizza = new CheesePizza(toppingFactory);
break;
case "pepperoni":
pizza = new PepperoniPizza(toppingFactory);
break;
case "veggie":
pizza = new VeggiePizza(toppingFactory);
break;
default: throw new IllegalArgumentException("No such pizza.");
}
pizza.addIngredients();
pizza.bakePizza();
return pizza;
}
}
```

GourmetPizzaFactory.java

```java
package guru.springframework.gof.abstractFactory;
import guru.springframework.gof.abstractFactory.product.CheesePizza;
import guru.springframework.gof.abstractFactory.product.PepperoniPizza;
import guru.springframework.gof.abstractFactory.product.Pizza;
import guru.springframework.gof.abstractFactory.product.VeggiePizza;
public class GourmetPizzaFactory extends BasePizzaFactory {
@Override
public Pizza createPizza(String type){
Pizza pizza;
BaseToppingFactory toppingFactory= new GourmetToppingFactory();
switch (type.toLowerCase())
{
case "cheese":
pizza = new CheesePizza(toppingFactory);
break;
case "pepperoni":
pizza = new PepperoniPizza(toppingFactory);
break;
case "veggie":
pizza = new VeggiePizza(toppingFactory);
break;
default: throw new IllegalArgumentException("No such pizza.");
}
pizza.addIngredients();
pizza.bakePizza();
return pizza;
}
}
```

In the above examples, we wrote the SicilianPizzaFactory and GourmetPizzaFactory subclasses of the abstract BasePizzaFactory class. In the subclasses we wrote the code to override the createPizza() factory method declared in BasePizzaFactory. In the createPizza() method, we first created the BaseToppingFactory object that a particular pizza factory will use for topping. We then used a switch statement to create a Pizza object based on the parameter passed to the method. Notice that while creating a Pizza object we initialized it by passing the BaseToppingFactory object to the constructor. We then made calls to the addIngredients() and bakePizza() methods on the Pizza object before returning it to the caller.

Now that our enhanced pizza store is ready for use, let us write a couple of unit tests to observe the abstract factory pattern at work.

GourmetPizzaFactoryTest.java

```java
package guru.springframework.gof.abstractFactory;
import guru.springframework.gof.abstractFactory.product.Pizza;
import org.junit.Test;
public class GourmetPizzaFactoryTest {
@Test
public void testCreatePizza() throws Exception {
BasePizzaFactory pizzaFactory=new GourmetPizzaFactory();
Pizza cheesePizza= pizzaFactory.createPizza("cheese");
Pizza veggiePizza=pizzaFactory.createPizza("veggie");
}
}
```

SicilianPizzaFactoryTest.java

```java
package guru.springframework.gof.abstractFactory;
import guru.springframework.gof.abstractFactory.product.Pizza;
import org.junit.Test;
public class SicilianPizzaFactoryTest {
@Test
public void testCreatePizza() throws Exception {
BasePizzaFactory pizzaFactory=new SicilianPizzaFactory();
Pizza cheesePizza=pizzaFactory.createPizza("cheese");
Pizza pepperoniPizza =pizzaFactory.createPizza("pepperoni");
}
}
```

When you run the above unit test code examples, you will see the following output:

---

## T E S T S

Running guru.springframework.gof.abstractFactory.GourmetPizzaFactoryTest
Preparing ingredients for cheese pizza.
Preparing goat cheese...
Preparing california oil sauce..
Pizza baked at 400 for 20 minutes.
Preparing ingredients for veggie pizza.
Preparing goat cheese...
Preparing california oil sauce..
Pizza baked at 400 for 20 minutes.
Tests run: 1, Failures: 0, Errors: 0, Skipped: 0, Time elapsed: 0.46 sec - in guru.springframework.gof.abstractFactory.GourmetPizzaFactoryTest
Running guru.springframework.gof.abstractFactory.SicilianPizzaFactoryTest
Preparing ingredients for cheese pizza.
Preparing mozzarella cheese...
Preparing tomato sauce..
Pizza baked at 400 for 20 minutes.
Preparing ingredients for pepperoni pizza.
Preparing mozzarella cheese...
Preparing tomato sauce..
Pizza baked at 400 for 20 minutes.
Tests run: 1, Failures: 0, Errors: 0, Skipped: 0, Time elapsed: 0.001 sec - in guru.springframework.gof.abstractFactory.SicilianPizzaFactoryTest
Running guru.springframework.gof.factoryMethod.PizzaFactoryTest
Preparing ingredients for cheese pizza.
Pizza baked at 400 for 20 minutes.
Preparing ingredients for veggie pizza.
Pizza baked at 400 for 20 minutes.
Tests run: 1, Failures: 0, Errors: 0, Skipped: 0, Time elapsed: 0.001 sec - in guru.springframework.gof.factoryMethod.PizzaFactoryTest

## Conclusion

As you get further into enterprise application development, you will you will encounter use cases for the abstract factory pattern, especially as the objects you’re creating become more complex. Its not uncommon to start off using the factory method design pattern and have your code evolve into using the abstract factory design pattern. Often you’ll find you only need one instance of the factory. If this is the case, you should consider implementing the concrete factory as a Singleton.
