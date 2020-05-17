# Observer Pattern

In the series of the Gang of Four design patterns, I wrote about the Command, Chain of Responsibility, Iterator, Mediator, Interpreter, and Memento patterns. These are part of the Behavioral pattern family that address responsibilities of objects in an application and how they communicate between them at runtime. In this post, I will look at using the Observer Pattern in Java. The Observer Pattern is a very common pattern. You may already be using it without knowing it!

## Observer Pattern: Introduction

If you have done programming on messaging applications using JMS or GUI-based applications using Swing, you are likely aware of the publish-subscribe interaction model. In messaging applications, a publisher object publishes a message to a destination and subscriber objects that subscribe to the destination receive the message. Similarly in Swing, the whole concept of components having registered action listeners that fire up whenever the component’s state changes is based on the publish-subscribe model.

Even if you are not aware of the publish-subscribe model, don’t be anguished – It is very simple and intuitive. Consider this real life analogy of a magazine publisher that publishes magazines monthly. You, as a subscriber, subscribe to the publisher and you keep receiving a copy whenever a new edition of the magazine is published. At any time when you don’t want to receive the magazine anymore, you can unsubscribe, and the publisher stops sending you a copy.

Twitter is a great example of the Observer Pattern. You follow a user and when that user tweets, you along with all other followers of that user receive the tweet. Here, the user account you are following is the publisher and your twitter account along with the other followers are the subscribers. As with any other subscriber, when you unfollow the user, you stop receiving tweets of the user.

If you’ve understood the preceding scenarios involving publisher and subscriber, you are seeing the Observer Pattern. The only difference – In this pattern we will refer publisher as subject and subscriber as observer.

The Observer Pattern facilitates communication between decoupled objects. It defines relationship between objects – usually a single subject having a one-to-many relationship with multiple observers so that when the state of the subject changes, all its observers are notified accordingly. This figure illustrates what I just explained.

Subject and Observer objects in the Observer pattern,.

However, GoF doesn’t limit relationship between one subject with multiple registered observers. One observer can register itself with multiple subjects to receive notifications from them. Also, a subject can itself be an observer of another subject resulting in a chain of observers. However, this is not recommended as it can be the cause of debugging nightmares and result in memory leakage issues that Martin Flower very nicely explained here.

## Participants of the Observer Pattern

To identify the participants of the Observer pattern, consider a scenario of an online auction application that enables registered users to bid on auctioned products. The application enables bidders to optionally opt to receive notifications through email or SMS whenever another bidder places a higher bid amount on a product. A bidder who is not interested in bidding on the product can opt out from receiving further notifications. Recalling the concepts covered till now, we can identify a product under auction as the subject (its state (price) will change when a new bid is placed). As there can be different products under auction, we will go ahead and declare the contract of defining products through an interface. Let’s name the interface Subject. Let’s next model real products by creating a concrete class that implements Subject. For this example, we will name the concrete class Product.

Next, we will proceed to create the observers. For that, we will again create an interface, named Observer. We will model real observers as objects of a Bidder class that implements the Observer interface.

Let’s now summarize the participants of the Observer pattern.

- Subject (Subject interface): Provides an interface to attach and detach Observer objects.
- ConcreteSubject (Product class): Implements the Subject interface. A ConcreteSubject sends notification to Observer objects when its state change.
- Observer (Observer interface): Provides an interface for objects that should be notified of changes in a Subject.
- ConcreteObserver (Bidder class): Implements Observer to receive notifications from the Subject and keep its state consistent with the state of the Subject.

## Applying the Observer Pattern

Going ahead with the auction application, let’s start applying the Observer pattern by writing the Subject interface and the Product class.

Subject.java

```java
package guru.springframework.gof.observer.observerimpl;
import java.math.BigDecimal;
public interface Subject {
    public void registerObserver(Observer observer);
    public void removeObserver(Observer observer);
    public void notifyObservers();
    public void setBidAmount(Observer observer,BigDecimal newBidAmount);
}
```

Product.java

```java
package guru.springframework.gof.observer.observerimpl;
import java.math.BigDecimal;
import java.util.ArrayList;
public class Product implements Subject{
    private ArrayList<Observer> observers = new ArrayList<>();
    private String productName;
    private BigDecimal bidAmount;
    private Observer observer;
    public Product(String productName, BigDecimal bidAmount){
        this.productName=productName;
        this.bidAmount=bidAmount;
    }
    @Override
    public void setBidAmount(Observer observer,BigDecimal newBidAmount){
        int res=bidAmount.compareTo(newBidAmount);
        if(res==-1){
            this.observer=observer;
            this.bidAmount=newBidAmount;
            notifyObservers();
        }
        else {
            System.out.println("New bid amount cannot be less or equal to current bid amount: "+this.bidAmount);
        }
    }
    @Override
    public void registerObserver(Observer observer) {
        observers.add(observer);
    }
    @Override
    public void removeObserver(Observer observer) {
        observers.remove(observer);
        System.out.println("-----------------"+observer+" has withdrawn from bidding----------------");
    }
    @Override
    public void notifyObservers() {
        System.out.println("-----------------New bid placed----------------");
        for (Observer ob : observers) {
            ob.update(this.observer,this.productName,this.bidAmount );
        }
    }
}
```

In the Subject interface above, we declared four methods that the Product class implements. In the overridden registerObserver() and removeObserver() methods of the Product class, we added and removed the Observer object passed to the methods to and from an ArrayList respectively. In the overridden setBidAmount() method, we assigned an Observer object (that represents a bidder) and the bid amount to the instance variables declared in the class. At runtime, when this method gets called, the state of the Product object changes. As the Product object needs to notify all registered observers about the change, we called the notifyObservers() method. In the notifyObservers() method, we iterated through the registered Observer objects stored in the ArrayList. For each iteration, we called the update() method on on the current Observer object passing the changed state of the subject (Product) and other information that the Observer object requires. We can even pass values that cause the Observer objects to react and rearrange their own state accordingly to be consistent with the new state of the subject.

Let’s next write the Observer interface and the Bidder class.

Observer.java

```java
package guru.springframework.gof.observer.observerimpl;
import java.math.BigDecimal;
public interface Observer {
    public void update(Observer observer,String productName, BigDecimal bidAmount);
  }
```

Bidder.java

```java
package guru.springframework.gof.observer.observerimpl;
import java.math.BigDecimal;
public class Bidder implements Observer{
    String bidderName;
    public Bidder(String bidderName) {
        this.bidderName = bidderName;
    }
    @Override
    public void update(Observer observer,String productName, BigDecimal bidAmount){
        if(observer.toString().equals(bidderName)){
            System.out.println("Hello "+bidderName+"! New bid of amount "+bidAmount+" has been placed on "+productName+" by you");
        }
        if(!observer.toString().equals(bidderName)) {
            System.out.println("Hello " + bidderName + "! New bid of amount " + bidAmount + " has been placed on " + productName + " by " + observer);
        }
    }
    @Override
    public String toString(){
        return bidderName;
    }
}
```

The Observer interface that we wrote declares a single update() method. For the sake of this example, we simply printed out the information about the changed state of the subject (Product) in the overridden update() method of the Bidder class. In a real-world enterprise application, you might perhaps need to write code to send out emails to the bidder’s accounts or send SMS to them or perform any other functions that the requirements demand.

If you have noticed, the most important thing here is that the concrete Bidder and Product classes don’t have any reference to each other. This is powerful decoupling – an important programming practice that the SOLID design principles advocate. What it means is that if we later need to add another observer, say the auction manager who wants notifications of high value bidding on a product, we simply write a AuctionManager class to implement the Observer interface and register its object with the subject (Product). The subject won’t require any modifications.

And how we achieved this? Recall the SOLID design principles that state “depend upon abstractions, not implementations”, which sums up Dependency Inversion principle. That is exactly what we did by writing the Subject and Observer interfaces.

Let’s write some test code to see the Observer pattern at work.

ObserverTest.java

```java
package guru.springframework.gof.observer.observerimpl;
import org.junit.Test;
import java.math.BigDecimal;
public class ObserverTest {
    @Test
    public void testObserver() throws Exception {
        Subject product=new Product("36 inch LED TV",new BigDecimal(350));
        Observer bidder1=new Bidder("Alex Parker");
        Observer bidder2=new Bidder("Henry Smith");
        Observer bidder3=new Bidder("Mary Peterson");
        product.registerObserver(bidder1);
        product.registerObserver(bidder2);
        product.registerObserver(bidder3);
        product.setBidAmount(bidder1, new BigDecimal(375));
        product.removeObserver(bidder2);
        product.setBidAmount(bidder3, new BigDecimal(400));
    }
}

```

In the test code above, we created a new Product object (subject) and three Bidder objects (observers). Next, we made calls to the registerObserver() method to register the observers with the subject. We then made a call to the setBidAmount() method of the subject. This will result in a change of state of the subject and the registered observers will get the notification. We then unregistered a bidder and again made a call to the setBidAmount() method to confirm that the unregistered bidder doesn’t receive notification this time. The output of the test is this.

---

## T E S T S

Running guru.springframework.gof.observer.observerimpl.ObserverTest
-----------------New bid placed----------------
Hello Alex Parker! New bid of amount 375 has been placed on 36 inch LED TV by you
Hello Henry Smith! New bid of amount 375 has been placed on 36 inch LED TV by Alex Parker
Hello Mary Peterson! New bid of amount 375 has been placed on 36 inch LED TV by Alex Parker
-----------------Henry Smith has withdrawn from bidding----------------
-----------------New bid placed----------------
Hello Alex Parker! New bid of amount 400 has been placed on 36 inch LED TV by Mary Peterson
Hello Mary Peterson! New bid of amount 400 has been placed on 36 inch LED TV by you
Tests run: 1, Failures: 0, Errors: 0, Skipped: 0, Time elapsed: 0.003 sec - in guru.springframework.gof.observer.observerimpl.ObserverTest
Observable Pattern in the Java API
Any discussion on the Observer pattern won’t be complete unless the Observable class and the Observer interface that Java provides out of the box for applying the Observer pattern are covered. Both the class and the interface are part of the java.util package. You can subclass Observable to represent a subject that observers wants to observe. Following is the modified Product class of our auction application that now extends Observable instead of our custom Subject interface that we wrote earlier.

Product.java

```java
package guru.springframework.gof.observer.javaapi;
import java.util.Observable;
import java.util.Observer;
import java.math.BigDecimal;
public class Product extends Observable{
    private String productName;
    private BigDecimal bidAmount;
    private Observer observer;
    public Observer getObserver() {
        return observer;
    }
    public BigDecimal getBidAmount() {
        return bidAmount;
    }
    public String getProductName() {
        return productName;
    }
    public Product(String productName, BigDecimal bidAmount){
        this.productName=productName;
        this.bidAmount=bidAmount;
    }
    public void setBidAmount(Observer observer,BigDecimal newBidAmount){
        System.out.println("-----------------New bid placed----------------");
        int res=bidAmount.compareTo(newBidAmount);
        if(res==-1){
            this.observer=observer;
            this.bidAmount=newBidAmount;
            setChanged();
            notifyObservers();
        }
        else {
            System.out.println("New bid amount cannot be less or equal to current bid amount: "+this.bidAmount);
        }
    }
}
```

As you can see the Product class above now have lesser amount of code. Now, the Product class doesn’t have the methods to register and remove observer objects; because by extending Observable, Product inherits the addObserver(Observer o) and deleteObserver(Observer o) methods of Observable. Things also changed in the setBidAmount() method. We now made calls to the setChanged() and notifyObservers() methods of the Observable class. The setChanged() method marks this Observable object as having been changed while the notifyObservers() method, as its name suggests, notifies all of its observers about the change. For the observers, we need to implement the Observer interface that have a single update(Observable o, Object arg) method. This method gets called whenever the subject (Product) changes. Following is the modified Bidder class implementing the Observer interface.

Bidder.java

```java
package guru.springframework.gof.observer.javaapi;
import java.util.Observable;
import java.util.Observer;
public class Bidder implements Observer {
    Product observable;
    String bidderName;
    public Bidder(String bidderName) {
        this.bidderName = bidderName;
    }
    @Override
    public void update(Observable observable, Object arg){
        this.observable = (Product) observable;
        String name = this.observable.getObserver().toString();
        if(name.equals(bidderName))
        {
            System.out.println("Hello "+bidderName+"! New bid of amount "+this.observable.getBidAmount()+" has been placed on "+this.observable.getProductName()+" by you");
        }
        if (!name.equals(bidderName))
            System.out.println("Hello "+bidderName+"! New bid of amount "+this.observable.getBidAmount()+" has been placed on "+this.observable.getProductName()+" by "+this.observable.getObserver());
    }
    @Override
    public String toString()
    {
        return bidderName;
    }
    }
```

In the overridden update() method of Bidder, we again printed out the information about the changed state of the subject (Product).

Let’s now write the test code.

ObservableJavaAPITest.java

```java
package guru.springframework.gof.observer.javaapi;
import org.junit.Test;
import java.math.BigDecimal;
public class ObservableJavaAPITest {
    @Test
    public void testObserver() throws Exception {
        Product product=new Product("L340 Digital Camera",new BigDecimal(325));
        Bidder bidder1=new Bidder("Shally Ferguson");
        Bidder bidder2=new Bidder("Dwayne Bravo");
        Bidder bidder3=new Bidder("Craig Dawson");
        product.addObserver(bidder1);
        product.addObserver(bidder2);
        product.addObserver(bidder3);
        product.setBidAmount(bidder1, new BigDecimal(350));
        product.deleteObserver(bidder2);
        product.setBidAmount(bidder3, new BigDecimal(375));
    }
}
```

As you can see, the test code above is similar to the code we wrote earlier for our custom implementation but with a different set of data. Also, notice that instead of calling our custom registerObserver() and removeObserver() methods, here we made calls to the addObserver() and deleteObserver() methods of the Observable class. The output on running the test is this.

---

## T E S T S

Running guru.springframework.gof.observer.javaapi.ObservableJavaAPITest
-----------------New bid placed----------------
Hello Craig Dawson! New bid of amount 350 has been placed on L340 Digital Camera by Shally Ferguson
Hello Dwayne Bravo! New bid of amount 350 has been placed on L340 Digital Camera by Shally Ferguson
Hello Shally Ferguson! New bid of amount 350 has been placed on L340 Digital Camera by you
-----------------New bid placed----------------
Hello Craig Dawson! New bid of amount 375 has been placed on L340 Digital Camera by you
Hello Shally Ferguson! New bid of amount 375 has been placed on L340 Digital Camera by Craig Dawson
Tests run: 1, Failures: 0, Errors: 0, Skipped: 0, Time elapsed: 0.158 sec - in guru.springframework.gof.observer.javaapi.ObservableJavaAPITest
Observer Pattern in the Spring Framework
In Spring, the Observer Pattern is used in the ApplicationContext's event mechanism. The ApplicationEvent class and ApplicationListener interface of Spring enables event handling in Spring ApplicationContext. When you deploy a bean that implements the ApplicationListener interface, it will receive an ApplicationEvent every time the ApplicationEvent is published by an event publisher. Here, the event publisher is the subject and the bean that implements ApplicationListener is the observer.

If you are creating your own custom event, your event publisher (subject) must implement the ApplicationEventPublisherAware interface. This interface has a setter method named setApplicationEventPublisher() that provides an ApplicationEventPublisher object for using in your class. The subject can then publish an event by calling the publishEvent() method of ApplicationEventPublisher. The subject can publish any event that extends ApplicationEvent and when the subject does so, the bean implementing ApplicationListener (observer) receives the event. If you’re interested in creating custom events in Spring, I have a section on event publishing and consuming custom application events in my Spring Core online course.

## Summary

In this post I have shown how to apply the Observer pattern with your own implementation and also using the in-built Observable class and Observer interface. So the natural question is – which approach should we use? I personally prefer my own implementation. The primary reason is, Observable being a class, my subject after extending Observable won’t be able to extend another class. Often, I find classes designed to be a subject already being part of an inheritance hierarchy. This out rightly prevents extending Observable. In addition, I like to have more liberty on how my subject stores its observers. As, you can see in the example code, I used an ArrayList – In fact I could have used any collection. But, by extending Observable, my subject is forced to store observers in a Vector because it is what Observable internally uses. I am not trying to push any point that my ArrayList implementation is better than the Vector implementation of Observable. But I can say that ArrayList being faster than Vector (because ArrayList have no synchronization overhead like Vector), my custom implementation will be the preferred choice in a performance critical application when synchronization is not an issue. Even if synchronization becomes a factor, I can always add the required synchronization code or even choose Vector instead of ArrayList. In short, writing your own implementation gives you flexibility. You just pay the cost of writing those extra lines of code.

At the end, irrespective of the implementation approach, the point is – During Enterprise Applications using the Spring Framework whenever you are in a situation where you have an object that needs to share its state with other objects, without knowing who those objects are, the Observer pattern is exactly what you need.
