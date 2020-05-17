# Memento Pattern

In the series of the Gang of Four design patterns, we have learned about the Command, Chain of Responsibility, Iterator, and Mediator patterns These patterns are part of the Behavioral pattern family that address responsibilities of objects in an application and how they communicate between them. In this post, I’ll discuss the Memento Pattern which you can use to capture an object’s internal state and save it externally so you can restore it later.

## Memento Pattern: Introduction

Consider a computer game that goes on for a long time. In a well-designed game, it is necessary to keep saving the state of the game periodically. This will allow a player to undo a wrong move or revert back to an earlier level and resume playing. You may be thinking that this is possible without applying any design pattern by implementing the logic to undo and revert back in the core applications itself. But, as I mentioned earlier, we are working on a well-designed game where the object’s internal states are not exposed because by doing so would violate encapsulation – a key OOP concept.

Besides gaming applications, such requirements and challenges are also part of other types of applications. For example, a desktop word processing application should allow users undo their operation while an enterprise e-commerce application might require allowing a user to revert back a checkout process comprising of multiple steps to an earlier step. The Memento Pattern is there to address such requirements without violating encapsulation.

When using the Memento Pattern, you have an object, called the Originator whose state (or even its partial state) needs to be saved. You then create another object, called the Memento which will hold different states of the Originator. Therefore the Memento class needs to have the same properties as the Originator in order to save the state. But if the Originator has its properties as private fields, then they won’t be accessible outside the Originator. This means the Memento object cannot access the private fields? The solution to this problem is the is core to the Memento Pattern. By applying the Memento Pattern, the Originator object will be able to:

Create a Memento object with the current state of the Originator object. The Originator will update the Memento object whenever it’s state changes and it deems it necessary to save the changed state.
Restore its previous state from a Memento object. By separating the logic of saving an object state from the object (Originator) itself, the Memento pattern adheres to the Single Responsibility principle, one of the SOLID design principles.
The Memento itself is a POJO that remains “opaque” to other objects. It is only the Originator that can store and retrieve state information from the Memento.

We have an Originator object that will create different Memento objects to hold it’s current state as it’s state changes. But how will you manage the Memento objects? This is where the Caretaker object comes into play. The purpose of the Caretaker object is the safekeeping of Memento objects. It is equally important that the Caretaker object never modifies the state of the Memento object. This modification would ripple back to the Originator object, and would be a violation of encapsulation.

Let’s look closer at the participants of the Memento Pattern.

## Participants of the Memento Pattern

Consider an employee management application which a user of the HR department uses to enter employee details when a new employee joins. The process is long and divided into multiple steps. It starts with storing basic information of the employee, and continues to recording previous working history of the employee. Imagine, if the user makes an error during the process and finds it in the last step of the process. It would be a poor user experience to force the user to start over. We need to allow the user to undo the steps and work backward to correct the error.

By applying the Memento pattern, we will create a EmpOriginator class, which is object we wish to save the state of. We will then create an EmpMemento class which will represent the different states of the EmpOriginator object. Finally, we will create the EmpCaretaker class which will manage the EmpMemento objects.

In summary, there are the objects we will be creating to implement the Memento Pattern:

- Originator (EmpOriginator): Is a class the state of whose object needs to be saved. It creates a Memento containing a snapshot of its current state. Originator uses the Memento to restore back its state.
- Memento (EmpMemento): Is a class whose objects stores states of the Originator. Memento denies other objects access to itself except the Originator.
- Caretaker (EmpCaretaker): Manages and safeguards Memento.

## Applying the Memento Pattern

Moving ahead with our example of employee application, let’s apply the Memento pattern to solve the undo problem we discussed. We will start with the EmpOriginator class – the Originator.

EmpOriginator.java

```java
package guru.springframework.gof.memento;
public class EmpOriginator {
    private int empId;
    private String empName;
    private String empPhoneNo;
    private String empDesignation;
    public EmpOriginator(int empId, String empName, String empPhoneNo,String empDesignation)
    {
        this.empId=empId;
        this.empName=empName;
        this.empPhoneNo=empPhoneNo;
        this.empDesignation=empDesignation;
    }
    public int getEmpId() {
        return empId;
    }
    public void setEmpId(int empId) {
        this.empId = empId;
    }
    public String getEmpName() {
        return empName;
    }
    public void setEmpName(String empName) {
        this.empName = empName;
    }
    public String getEmpPhoneNo() {
        return empPhoneNo;
    }
    public void setEmpPhoneNo(String empPhoneNo) {
        this.empPhoneNo = empPhoneNo;
    }
    public String getEmpDesignation() {
        return empDesignation;
    }
    public void setEmpDesignation(String empDesignation) {
        this.empDesignation = empDesignation;
    }
    public EmpMemento saveToMemento() {
        EmpMemento empMemento=new EmpMemento(this.empId, this.empName, this.empPhoneNo, this.empDesignation );
        return empMemento;
    }
    public  void undoFromMemento(EmpMemento memento)
    {
        this.empId = memento.getEmpId();
        this.empName = memento.getEmpName();
        this.empPhoneNo = memento.getEmpPhoneNo();
        this.empDesignation = memento.getEmpDesignation();
    }
    public void printInfo()
    {
        System.out.println("ID: "+ this.empId);
        System.out.println("Name: "+ this.empName);
        System.out.println("Phone Number: "+ this.empPhoneNo);
        System.out.println("Designation: "+ this.empDesignation);
    }
}
```

In the initial part of the EmpOriginator class above, we wrote fields to hold the state of EmpOriginator that are initialized through the constructor. For each property, we wrote the corresponding getter and setter method. It is from Line 50 – Line 54 , where we wrote the saveToMemento() method, we start using the Memento pattern. In the saveToMemento() method, we created and returned back an EmpMemento object initialized with the current state of EmpOriginator. From Line 55 – Line 62 we wrote the undoFromMemento() method that accepts a EmpMemento object. In the undoFromMemento() method, we assigned the current state of EmpOriginator with the state of the EmpMemento object. The printInfo() method that we wrote from Line 64 – Line 70 is a utility method to print out the current state of EmpOriginator.

EmpMemento.java

```java
package guru.springframework.gof.memento;
public class EmpMemento {
    private int empId;
    private String empName;
    private String empPhoneNo;
    private String empDesignation;
    public EmpMemento(int empId,String empName,String empPhoneNo,String empDesignation) {
        this.empId = empId;
        this.empName = empName;
        this.empPhoneNo = empPhoneNo;
        this.empDesignation = empDesignation;
    }
   public int getEmpId() {
        return empId;
    }
    public String getEmpName() {
        return empName;
    }
    public String getEmpDesignation() {
        return empDesignation;
    }
    public String getEmpPhoneNo() {
        return empPhoneNo;
    }
    @Override
    public String toString(){
        String str="Current Memento State" + this.empId +" , "+this.empName +" , "+this.getEmpPhoneNo()+" , "+this.getEmpDesignation();
        return str;
    }
    }
```

As you can observe, the EmpMemento class we wrote above have the same fields as the EmpOriginator class. We declared the properties as private because we don’t want other objects (except EmpOriginator) to modify the properties of our Memento. So, we also haven’t wrote any setter methods for the properties.

Next we need to implement Caretaker object.

EmpCaretaker.java

```java
package guru.springframework.gof.memento;
import java.util.ArrayDeque;
import java.util.Deque;
import java.util.Stack;
public class EmpCaretaker {
    final Deque<EmpMemento> mementos = new ArrayDeque<>();
    public EmpMemento getMemento()
    {
        EmpMemento empMemento= mementos.pop();
        return empMemento;
    }
    public void addMemento(EmpMemento memento)
    {
        mementos.push(memento);
    }
}
```

In the EmpCaretaker class above, we used the ArrayDeque class which is a linear collection that supports element insertion and removal at both ends. From Line 11 – Line 16 we wrote the getMemento() method, where we called the pop() method of ArrayDeque that removes and returns the first EmpMemento of this deque. From Line 18 – Line 22 we wrote the addMemento() method that accepts a EmpMemento object. In this method, we called the push() method of ArrayDeque passing the EmpMemento object to be pushed at the head of this deque. Let’s now write some test code to see the Memento Pattern in action.

EmpOriginatorTest.java

```java
package guru.springframework.gof.memento;
import org.junit.Test;
import static org.junit.Assert.*;
public class EmpOriginatorTest {
    @Test
    public void testMemento() throws Exception {
        EmpOriginator empOriginator= new EmpOriginator(306,"Mark Ferguson", "131011789610","Sales Manager");
        EmpMemento empMemento=empOriginator.saveToMemento();
        EmpCaretaker empCaretaker=new EmpCaretaker();
        empCaretaker.addMemento(empMemento);
        System.out.println("\n Original EmpOriginator");
        empOriginator.printInfo();
        System.out.println("\n EmpOriginator after updating phone number");
        empOriginator.setEmpPhoneNo("131011888886");
        empMemento=empOriginator.saveToMemento();
        empCaretaker.addMemento(empMemento);
        empOriginator.printInfo();
        System.out.println("\n EmpOriginator after updating designation");
        empOriginator.setEmpDesignation("Senior Sales Manager");
        empMemento=empOriginator.saveToMemento();
        empCaretaker.addMemento(empMemento);
       empOriginator.printInfo();
        System.out.println("\n EmpOriginator after undoing designation update");
        empMemento=empCaretaker.getMemento();
        empOriginator.undoFromMemento(empMemento);
        empMemento=empCaretaker.getMemento();
        empOriginator.undoFromMemento(empMemento);
        empOriginator.printInfo();
        System.out.println("\n Original EmpOriginator after undoing phone number update");
        empMemento=empCaretaker.getMemento();
        empOriginator.undoFromMemento(empMemento);
        empOriginator.printInfo();
    }
}
```

In the test class above, from line 12 to line 16 we created an EmpOriginator object and called its saveToMemento() method that returns back a EmpMemento object containing the snapshot of the EmpOriginator object’s state. We then created an EmpCaretaker object and called its addMemento() method passing the EmpMemento to be stored. In line 22, we updated the empPhoneNo field of EmpOriginator by calling the setEmpPhoneNo() method.

We then called the saveToMemento() mthod to obtain a new EmpMemento object with the updated state of EmpOriginator and passed the new EmpMemento object to EmpCaretaker. From line 28 to line 30, we performed the same steps after updating the empDesignation property of EmpOriginator. For the undo operations, we called the getMemento() method of EmpCaretaker. We then called the undoFromMemento() method of EmpOriginator passing the EmpMemento object that the getMemento() method returns. We wrote the same code from line 43 to line 44 to revert the EmpOriginator object back to its original state.

Here is the output of the test:

---

## T E S T S

Running guru.springframework.gof.memento.EmpOriginatorTest
Original EmpOriginator
ID: 306
Name: Mark Ferguson
Phone Number: 131011789610
Designation: Sales Manager
EmpOriginator after updating phone number
ID: 306
Name: Mark Ferguson
Phone Number: 131011888886
Designation: Sales Manager
EmpOriginator after updating designation
ID: 306
Name: Mark Ferguson
Phone Number: 131011888886
Designation: Senior Sales Manager
EmpOriginator after undoing designation update
ID: 306
Name: Mark Ferguson
Phone Number: 131011888886
Designation: Sales Manager
Original EmpOriginator after undoing phone number update
ID: 306
Name: Mark Ferguson
Phone Number: 131011789610
Designation: Sales Manager
Tests run: 1, Failures: 0, Errors: 0, Skipped: 0, Time elapsed: 0.008 sec - in guru.springframework.gof.memento.EmpOriginatorTest

## Summary

The Memento Pattern is a powerful design pattern, which should be in your programming toolbox. A common usage of this pattern is an implementation of ‘Ok’ and ‘Cancel’ dialogs. When the dialog loads, its state is stored and you work on the dialog. If you press Cancel, the initial state of the dialog is restored.
