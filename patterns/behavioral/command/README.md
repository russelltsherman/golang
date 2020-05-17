# Command Pattern

The Command Pattern is one of the 11 design patterns of the Gang of Four Behavioral pattern family. We’ve already learned about the Chain of Responsibility pattern, and similar to it and the other patterns of the family, the Command Pattern addresses responsibilities of objects in an application and how they communicate between them.

## Command Pattern: Introduction

Communication between objects in enterprise applications starts from an object sending a request to another object. The object receiving the request can either process the request and send back a response or forward the request to another object for processing. Typically requests are made by the invoker (the object making the request) through method calls on the object that processes the request, which we will refer as the receiver. As a result, the invoker and the receiver are tightly coupled. This violates the SOLID design principles that advocates loosely coupled components to ensure that changes made to one component does not affect the other components of the application.

The Command Pattern is a proven solution that addresses such recurring problems of tightly coupled invoker and receiver components in applications. This pattern states that requests should be encapsulated as objects that like any other objects can be stored and passed around the application. Requests encapsulated as objects are known as commands.

In the command pattern, the invoker issues command without knowing anything about the receiver. In fact, the invoker issuing the command doesn’t even know what operation will be carried out on issuing a command. Let’s look at it from a programming point of view.

A command object basically has a execute() method and an optional undo() method. The execute() method executes some operation on the receiver while the undo() method reverses the current operation. The implementation of the operation is done by the receiver. The invoker only sets itself up with a command object and invokes a command by calling the execute() method. The invoker does not know what the execute() method will do.

Imagine that the invoker is a switch, When the invoker invokes a command by calling the execute() method of a command object, it might turn on a light or even start a generator. What this essentially means is that the same invoker (switch) can be used to invoke commands to perform actions on different receivers (devices). We only need to create the appropriate command object and set it on the invoker. As apparent, by applying the Command Pattern we can take reusability to another level and this is made possible due to the loose coupling between the invoker and receiver and the command object that acts as the interface between them.

## Participants of the Command Pattern

Before drilling into the participants of the Command Pattern and their roles, let’s start with an analogy of remote-controlled toys.

Imagine a toy manufacturing company, which recently decided to manufacture remote-controlled toys. They started with a remote-controlled car. The remote control has three buttons: On, Off, and Undo. The remote control is programmed so that the On button moves the car, the Off button stops it, and the Undo button reverses the current action of the car. The remote-controlled car was a success and the management next decided to manufacture a remote controlled rotating top. Again, along with the top, a new remote control programmed to start and stop rotating and reverse the current action of the top was manufactured. Riding on the success of remote-controlled toys, the management decided to manufacture 200 varieties of remote controlled toys each with a remote control programmed differently to operate its corresponding toy. Imagine yourself in the programming team. You need to create 200 different programs for the remote controls each designed to operate on a specific toy. Apparently a bad design decision because of the inability to reuse code. The reason is also apparent. Each remote control is tightly coupled with the toy it operates on. What we need is a remote control that is programmed to operate on all the toys, and here is where the Command Pattern comes to the rescue.

Let’s look how we can apply the pattern as a solution. The toys are the receivers and perform actions. Let us model two receivers as Car and RotatingTop. These classes have behaviors in the form of methods to perform actions. For example, the Car receiver will have methods to move and stop and the RotatingTop will have methods to start and stop rotating.

We will next model the commands that will trigger the actions on the receivers. We start with a CommandBase interface with two methods execute() and undo(). Each concrete sub classes of CommandBase represents a command to trigger an action on the receiver, and to do so it maintains a reference to the receiver whose action it will trigger. For the Car and RotatingTop receivers, let us model the concrete command classes as CarMoveCommand, CarStopCommand, TopRotateCommand, and TopStopRotateCommand. Finally we will model the invoker as a RemoteControl class. This class is not aware of any concrete command class. Whenever one of its button is pressed, the method handling the press event is injected with an appropriate command object at runtime and the invoker calls the execute() method on it.

In the context of the remote controlled toys example, let’s summarize the participants of the Command Pattern.

- Command (CommandBase): Is an interface for executing an action.
- ConcreteCommand (CarMoveCommand, CarStopCommand, TopRotateCommand, and TopStopRotateCommand): Are concrete classes that implements Command and defines the execute() and undo() methods to communicate with receivers for performing an action and undoing it respectively.
- Invoker(RemoteControl): Asks Command to carry out the action.
- Receiver (Car and RotatingTop): Performs the action based on the command it receives.
- Client: Creates a ConcreteCommand object and sets its receiver.

## Applying the Command Pattern

Let’s implement the remote controlled toys example in Java code. We will start with the receivers.

Car.java

```java
package guru.springframework.gof.command.receiver;
public class Car {
public void move()
{
System.out.println("Car is moving");
}
public void stop()
{
System.out.println("Car has stopped");
}
}
RotatingTop.java
package guru.springframework.gof.command.receiver;
public class RotatingTop {
public void startRotating(){
System.out.println("Top has start rotating");
}
public void stopRotating(){
System.out.println("Top has stopped rotating");
}
}
```

Both the receiver classes above define methods, whose names itself are self-explanatory, to perform actions.

Now, let’s write the commands. We will start with the CommandBase interface.

CommandBase.java

```java
package guru.springframework.gof.command.commandobjects;
public interface CommandBase {
void execute();
void undo();
}
```

In the CommandBase interface above, we declared the execute() and undo() methods that the ConcreteCommand classes will implement to invoke the actions on the receivers. We have four actions, two actions for each receiver. So we will write four ConcreteCommand classes to invoke the four actions.

CarMoveCommand.java

```java
package guru.springframework.gof.command.commandobjects;
import guru.springframework.gof.command.receiver.Car;
public class CarMoveCommand implements CommandBase {
private Car car;
public CarMoveCommand(Car car){
this.car=car;
}
@Override
public void execute(){
System.out.println("CarMoveCommand.execute(): Invoking move() on Car");
car.move();
}
@Override
public void undo(){
System.out.println("CarMoveCommand.undo(): Undoing previous action->Invoking stop() on Car");
car.stop();
}
}
```

CarStopCommand.java

```java
package guru.springframework.gof.command.commandobjects;
import guru.springframework.gof.command.receiver.Car;
public class CarStopCommand implements CommandBase{
private Car car;
public CarStopCommand(Car car){
this.car=car;
}
@Override
public void execute(){
System.out.println("CarStopCommand.execute(): Invoking stop() on Car");
car.stop();
}
@Override
public void undo()
{
System.out.println("CarStopCommand.undo(): Undoing previous action-> Invoking move() on Car");
car.move();
}
}
```

TopRotateCommand.java

```java
package guru.springframework.gof.command.commandobjects;
import guru.springframework.gof.command.receiver.RotatingTop;
public class TopRotateCommand implements CommandBase{
RotatingTop rotatingTop;
public TopRotateCommand(RotatingTop rotatingTop){
this.rotatingTop=rotatingTop;
}
@Override
public void execute(){
System.out.println("TopRotateCommand.execute(): Invoking startRotating() on RotatingTop");
rotatingTop.startRotating();
}
@Override
public void undo(){
System.out.println("TopRotateCommand.undo(): Undoing previous action->Invoking stopRotating() on RotatingTop");
rotatingTop.stopRotating();
}
}
```

TopStopRotateCommand.java

```java
package guru.springframework.gof.command.commandobjects;
import guru.springframework.gof.command.receiver.RotatingTop;
public class TopStopRotateCommand implements CommandBase{
RotatingTop rotatingTop;
public TopStopRotateCommand(RotatingTop rotatingTop){
this.rotatingTop=rotatingTop;
}
@Override
public void execute(){
System.out.println("TopStopRotateCommand.execute(): Invoking stopRotating() on RotatingTop");
rotatingTop.stopRotating();
}
@Override
public void undo(){
System.out.println("TopStopRotateCommand.undo(): Undoing previous action->Invoking startRotating() on RotatingTop");
rotatingTop.startRotating();
}
}
```

Each of the preceding ConcreteCommand classes maintains a reference to its receiver and the reference is initialized through the constructor. In the overridden execute() method, we called the corresponding method of the receiver that implements the action to be performed. In the undo() method we reversed the action performed through execute(). With our commands and receivers set, we will next write the invoker – the RemoteControl class.

RemoteControl.java

```java
package guru.springframework.gof.command.invoker;
import guru.springframework.gof.command.commandobjects.CommandBase;
public class RemoteControl {
CommandBase onCommand, offCommand, undoCommand;
public void onButtonPressed(CommandBase onCommand){
this.onCommand=onCommand;
onCommand.execute();
undoCommand=onCommand;
}
public void offButtonPressed(CommandBase offCommand){
this.offCommand=offCommand;
offCommand.execute();
undoCommand=offCommand;
}
public void undoButtonPressed(){
undoCommand.undo();
}
}
```

In the RemoteControl class above, we wrote the onButtonPressed() method that accepts a CommandBase object. The ConcreteCommand object will be passed to the method at runtime. The method then calls the execute() method on the ConcreteCommand object.

As you can observe, the onButtonPressed() method does not know what ConcreteCommand object will be passed to it neither which receiver will perform the action. It can be the Car, the RotatingTop, or any other receiver we add later. In Line 12 we assigned the command object passed to onButtonPressed() to the undoCommand variable, which is of type CommandBase. This will help us track the current command being issued in order to undo the action carried out by the command whenever required.

The offButtonPressed() method is the same as the onButtonPressed() method.

In the undoButtonPressed() method, we called the undo() method on undoCommand. At runtime, the undo() method of the ConcreteCommand object that undoCommand is currently assigned to will get invoked.

Let’s write a test class and see how our remote control works. We will first create the commands for the Car receiver and load the RemoteControl object with the commands to test how the remote control operates the car. We will next load the RemoteControl object with commands for the RotatingTop receiver.

RemoteControlTest.java

```java
package guru.springframework.gof.command.invoker;
import guru.springframework.gof.command.commandobjects._;
import guru.springframework.gof.command.receiver.Car;
import guru.springframework.gof.command.receiver.RotatingTop;
import org.junit.Test;
import static org.junit.Assert._;
public class RemoteControlTest {
@Test
public void testRemoteControlButtonPressed() throws Exception {
RemoteControl remoteControl=new RemoteControl();
System.out.println("-----Testing onButtonPressed on RemoteControl for Car-----");
Car car=new Car();
CommandBase carMoveCommand=new CarMoveCommand(car);
remoteControl.onButtonPressed(carMoveCommand);
System.out.println("-----Testing offButtonPressed on RemoteControl for Car-----");
CommandBase carStopCommand=new CarStopCommand(car);
remoteControl.offButtonPressed(carStopCommand);
System.out.println("-----Testing undoButtonPressed() on RemoteControl for Car-----");
remoteControl.undoButtonPressed();
System.out.println("-----Testing onButtonPressed on RemoteControl for RotatingTop-----");
RotatingTop top=new RotatingTop();
CommandBase topRotateCommand=new TopRotateCommand(top);
remoteControl.onButtonPressed(topRotateCommand);
System.out.println("-----Testing offButtonPressed on RemoteControl for RotatingTop-----");
CommandBase topStopRotateCommand=new TopStopRotateCommand(top);
remoteControl.offButtonPressed(topStopRotateCommand);
System.out.println("-----Testing undoButtonPressed on RemoteControl for RotatingTop-----");
remoteControl.undoButtonPressed();
}
}
```

In the test method, we created a RemoteControl object. From Line 15 – Line 17 we created a Car object, then created a CarMoveCommand object initialized with the Car object. We then called the onButtonPressed() method of RemoteControl passing the CarMoveCommand object. Similarly, from Line 19 – Line 20, we created a CarStopCommand object initialized with the same Car object we created earlier, and then called the offButtonPressed() method of RemoteControl passing the CarStopCommand object. To test the undo functionality, we called the undo() method of RemoteControl in Line 22. Similarly, from Line 25 -31 we wrote the code to create a RotatingTop object, initialized the corresponding ConcreteCommand objects (TopRotateCommand and TopStopRotateCommand) with RotatingToy, and then called the onButtonPressed() and offButtonPressed() methods of RemoteControl, passing the appropriate ConcreteCommand objects. We also called the undo() method of RemoteControl in Line 35 to test the undo functionality on RotatingToy. The important thing to observe here is that we are using the same RemoteControl object to operate on both the receivers. So the next time we add a toy (receiver), we only need to create its ConcreteCommand classes and pass it to RemoteControl. The RemoteControl remains unchanged. The output of the code is this.

---

## T E S T S

Running guru.springframework.gof.command.invoker.RemoteControlTest
-----Testing onButtonPressed on RemoteControl for Car-----
CarMoveCommand.execute(): Invoking move() on Car
Car is moving
-----Testing offButtonPressed on RemoteControl for Car-----
CarStopCommand.execute(): Invoking stop() on Car
Car has stopped
-----Testing undoButtonPressed() on RemoteControl for Car-----
CarStopCommand.undo(): Undoing previous action-> Invoking move() on Car
Car is moving
-----Testing onButtonPressed on RemoteControl for RotatingTop-----
TopRotateCommand.execute(): Invoking startRotating() on RotatingTop
Top has start rotating
-----Testing offButtonPressed on RemoteControl for RotatingTop-----
TopStopRotateCommand.execute(): Invoking stopRotating() on RotatingTop
Top has stopped rotating
-----Testing undoButtonPressed on RemoteControl for RotatingTop-----
TopStopRotateCommand.undo(): Undoing previous action->Invoking startRotating() on RotatingTop
Top has start rotating
Tests run: 1, Failures: 0, Errors: 0, Skipped: 0, Time elapsed: 0.003 sec - in guru.springframework.gof.command.invoker.RemoteControlTest

## Summary

Beginning programmers may find the Command Pattern daunting. They may not have the vision nor experience to see how it applies to common programming tasks. The Command Pattern is commonly seen in GUI code, such as handling actions of buttons, menu items, action links and also Java progress bars and wizards. It is also seen in Runnable related code. But, this pattern is not limited to them. In developing web applications with Spring MVC, you often see the concepts of the Command Pattern applied through the use of Command Objects. Think about an e-commerce application where you are adding an item to your cart. That form post to add the item to your cart is effectively a “command”. Hence a core component of Spring MVC is the AbstractCommandController.

The Command Pattern isn’t just limited to Spring MVC. When developing enterprise applications using the Spring Framework, you will find plenty of other opportunities to apply the Command Pattern. Whenever you are writing code that requires some invoker to perform different actions on multiple receivers, consider using the Command Pattern – after all, it is one of the classic GoF design patterns.
