# Mediator Pattern

The Mediator Pattern, which is similar to the Command, Chain of Responsibility, and Iterator patterns, are part of the Behavioral pattern family of the Gang of Four design patterns.

## Behavioral Patterns: Introduction

Well-designed enterprise applications are composed of lightweight objects with specific responsibilities divided between them in accordance to the Single Responsibility principle, one of the SOLID design principles. However, the benefits of an application composed of a large number of small objects comes with a challenge – the communication between them. Objects need to communicate to perform business requirements, and as the number of objects grow, connections between the objects required can quickly become unmanageable. In addition objects communicating between them needs to know about each other – they are all tightly coupled that violates the basics of the SOLID principles. The Mediator pattern is a proven solution to address such recurring challenges and the problems arising out of them.

The Mediator pattern says that instead of allowing a set of objects to directly interact between them, define an object (mediator) that will handle the interactions. What the mediator essentially says to such set of objects is “talk with me instead of talking among yourselves”. This figure conceptually shows how objects interact without and with a mediator.

The Mediator pattern, as shown in the above figure employs a mediator object to enable other objects of the application to interact without knowing each other’s identities. The mediator also encapsulates a protocol that objects can follow.

It is the mediator object that encapsulates all interaction logic between objects in a system. Therefore, if an existing object is updated with new interaction rules or a new object is added to the system, it is only the mediator object that you need to update. In the absence of the mediator, you would need to update all the corresponding objects that which to interact. Through the use of the Mediator Pattern your code becomes more encapsulated, thus changes are not as extensive.

Let’s now identify the participants of the Mediator Pattern.

## Participants of the Mediator Pattern

Imagine a war zone where armed units are moving into the enemy’s territory. Armed units can include soldier, tank, grenadier, and sniper units. The strategy being employed is that whenever one unit attacks, other units should stop attacking and take cover. To do so, the unit that is currently attacking needs to notify the other units.

In the programming world, you can easily model this requirement by creating classes for each armed unit. In each class, whenever its object is about to start attacking, you can implement the logic to notify objects of the other classes. Now, imagine that a new unit joins in. The consequence – all the existing classes need to be updated. In the worst case, imagine that the battle tactics change so that both the soldier and sniper units can now attack simultaneously. Again, the consequence is a lot of changes to the code base. We can address such challenges, similarly as done in real life. Place a Commander in a base camp that will act as the mediator. All units, instead of communicating between themselves will communicate with the mediator. The mediator based on the notifications received from some units can send request to one or more other units to perform actions as requirements demand.

Let’s model the mediator with a Commander interface and a concrete CommanderImpl subclass of the interface. In the Commander interface, we can declare methods to send messages to objects representing armed units and also methods that armed unit’s objects can communicate with the Commander. In the CommanderImpl subclass, we will maintain reference to the objects representing armed units and override the methods of Commander. For our example, let’s create an interface as ArmedUnit whose implementing classes will represent specific armed units. We will create two such concrete classes: SoldierUnit and TankUnit, and each of them will hold references to Commander.

With our example, let’s look at the participants of the Mediator pattern.

- Mediator (Commander): Is an interface that declares methods for communicating with Colleague objects.
- ConcreteMediator (CommanderImpl): Implements Mediator. This class maintains and coordinates Colleague objects.
- Colleague(SoldierUnit and TankUnit): Communicates with its Mediator when their state changes and responds to requests from the Mediator.

## Applying the Mediator Pattern

Going ahead with our example on armed units, let’s apply the Mediator pattern to it. We will start with the Mediator interface followed by the ConcreteMediator classes.

Commander.java

```java
package guru.springframework.gof.mediator.mediator;
import guru.springframework.gof.mediator.colleague.ArmedUnit;
import guru.springframework.gof.mediator.colleague.SoldierUnit;
import guru.springframework.gof.mediator.colleague.TankUnit;
public interface Commander {
void registerArmedUnits(ArmedUnit soldierUnit, ArmedUnit tankUnit);
void setAttackStatus(boolean attackStatus);
boolean canAttack();
void startAttack(ArmedUnit armedUnit);
void ceaseAttack(ArmedUnit armedUnit);
}
```

CommanderImpl.java

```java
package guru.springframework.gof.mediator.mediator;
import guru.springframework.gof.mediator.colleague.ArmedUnit;
import guru.springframework.gof.mediator.colleague.SoldierUnit;
import guru.springframework.gof.mediator.colleague.TankUnit;
public class CommanderImpl implements Commander {
ArmedUnit soldierUnit, tankUnit;
boolean attackStatus = true;
@Override
public void registerArmedUnits(ArmedUnit soldierUnit, ArmedUnit tankUnit) {
this.soldierUnit = soldierUnit;
this.tankUnit = tankUnit;
}
@Override
public void setAttackStatus(boolean attackStatus) {
this.attackStatus = attackStatus;
}
@Override
public boolean canAttack() {
return attackStatus;
}
@Override
public void startAttack(ArmedUnit armedUnit) {
armedUnit.attack();
}
@Override
public void ceaseAttack(ArmedUnit armedUnit) {
armedUnit.stopAttack();
}
}
```

In the Commander interface (Mediator) above, we declared five methods that we implemented in the CommanderImpl class (ConcreteMediator). In the CommanderImpl class, we added two reference to ArmedUnit (Colleague) and initialized them through the registerArmedUnits() method from Line 12 – Line 16. In this class, we also declared the attackStatus boolean variable in Line 10. This variable will hold the state whether any ArmedUnit is currently attacking. Colleague objects can set this state by calling the setAttackStatus() method that we wrote from Line 18 – Line 21. The canAttack() method that we wrote from Line 23 – Line 26 returns the current state of the attackStatus variable. The startAttack() and ceaseAttack() methods make calls to the attack() and stopAttack() methods on the ArmedUnit object passed as method parameter.

We will next move ahead and create the Colleague interface and the implementation classes.

ArmedUnit.java

```java
package guru.springframework.gof.mediator.colleague;
public interface ArmedUnit {
void attack();
void stopAttack();
}
```

SoldierUnit.java

```java
package guru.springframework.gof.mediator.colleague;
import guru.springframework.gof.mediator.mediator.Commander;
public class SoldierUnit implements ArmedUnit{
private Commander commander;
public SoldierUnit(Commander commander){
this.commander=commander;
}
@Override
public void attack()
{
if(commander.canAttack())
{
System.out.println("SoldierUnit: Attacking.....");
commander.setAttackStatus(false);
}
else{
System.out.println("SoldierUnit: Cannot attack now. Other units attacking....");
}
}
@Override
public void stopAttack(){
System.out.println("SoldierUnit: Stopped Attacking.....");
commander.setAttackStatus(true);
}
}
```

TankUnit.java

```java
package guru.springframework.gof.mediator.colleague;
import guru.springframework.gof.mediator.mediator.Commander;
public class TankUnit implements ArmedUnit{
private Commander commander;
public TankUnit(Commander commander){
this.commander=commander;
}
@Override
public void attack()
{
if(commander.canAttack())
{
System.out.println("TankUnit: Attacking.....");
commander.setAttackStatus(false);
}
else{
System.out.println("TankUnit: Cannot attack now. Other units attacking....");
}
}
@Override
public void stopAttack(){
System.out.println("TankUnit: Stopped attacking.....");
commander.setAttackStatus(true);
}
}
```

We first wrote the ArmedUnit interface with methods that the implementing SoldierUnit and TankUnit (Colleague) classes override. In both the implementing classes, we maintained a reference to the Commander object that we initialized through the constructor. Also, in both the classes we implemented the overridden attack() and stopAttack() methods declared in the ArmedUnit interface. In the attack() method, we communicated with the mediator by calling the canAttack() method. If the method returns true, the Colleague object is confirmed that it can attack. We also wrote a stopAttack() method. The Mediator calls this method whenever it wants a unit to stop attacking. Notice that the Colleague classes are not aware of each other. Programmatically we are not maintaining any reference between Colleague objects. They only communicate with the Mediator and the Mediator communicates with them. This is the essence of the Mediator pattern. Let’s write some test code so we can observe the Mediator Pattern in action.

CommanderImplTest.java

```java
package guru.springframework.gof.mediator.mediator;
import guru.springframework.gof.mediator.colleague.ArmedUnit;
import guru.springframework.gof.mediator.colleague.SoldierUnit;
import guru.springframework.gof.mediator.colleague.TankUnit;
import org.junit.Test;
import static org.junit.Assert.\*;
public class CommanderImplTest {
@Test
public void testMediator() throws Exception {
Commander commander= new CommanderImpl();
ArmedUnit soldierUnit=new SoldierUnit(commander);
ArmedUnit tankUnit=new TankUnit(commander);
commander.registerArmedUnits(soldierUnit, tankUnit);
commander.startAttack(soldierUnit);
commander.startAttack(tankUnit);
commander.ceaseAttack(soldierUnit);
commander.startAttack(tankUnit);
}
}
```

In the test method of the CommanderImplTest class above, we created an object of type Commander and then objects, one each of SoldierUnit and TankUnit (Colleague objects). While creating the Colleague objects, we passed the Commander object as constructor argument. At runtime, both the SoldierUnit and TankUnit objects will have reference to the mediator (Commander).

Next, we called the registerArmedUnits() of Commander passing the instantiated SoldierUnit and TankUnit objects. Now, the Commander have references to both the SoldierUnit and TankUnit (Colleague objects). Then, we called called the startAttack() method passing the SoldierUnit and TankUnit objects one after the other. We expect that the TankUnit object on communicating back with the mediator (Commander) will find that another unit (which in our example is SoldierUnit) is already attacking and therefore the TankUnit object will not attack. It is only when the test method calls the ceaseAttck() method on Commander passing the currently attacking SoldierUnit object, the TankUnit should be able to attack. We verified this by calling the startAttack() method of Commander passing the TankUnit object.

The output of the test code is this.

---

## T E S T S

Running guru.springframework.gof.mediator.mediator.CommanderImplTest
SoldierUnit: Attacking.....
TankUnit: Cannot attack now. Other units attacking....
SoldierUnit: Stopped Attacking.....
TankUnit: Attacking.....
Tests run: 1, Failures: 0, Errors: 0, Skipped: 0, Time elapsed: 0.075 sec - in
The Mediator Pattern in the Spring Framework
In Spring MVC, there is a great example of the Mediator Pattern in action with how Spring MVC uses the Dispatcher Servlet in conjunction with the controllers.

Consider this diagram from the official Spring MVC documentation:

Spring MVC and the Mediator PatternHere you can see how the Front Controller (aka Dispatcher Servlet) acts as a mediator between web request, the controller objects, and the view objects. The individual controllers are unaware of each other. Each controller is also blissfully unaware of the context of the web request. The view templates are unaware of the each other, and the controllers. It is the responsibility of the Front Controller to decide with controller and which view template to utilize when building a response to a web request.

Because of the usage of the Mediator pattern, changes in the individual controllers do not affect other controllers. Changes in the web request object, will not cause downstream changes in the controllers, nor views.

## Summary

The Mediator pattern is considered one of the most important and widely adopted patterns. It is mainly because of its ability to encapsulate communication logic between sets of objects to fulfill some business requirements. A disadvantage of this pattern often being discussed in the developer community is that the mediator object can gradually get inflated and finally become overly complex. But, by following SOLID design principles, specifically the Single Responsibility principle, you can segregate the responsibilities of a mediator into separate classes.
