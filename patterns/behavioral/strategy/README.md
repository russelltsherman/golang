# Strategy Pattern

The Behavioral pattern family of the Gang of Four design patterns address responsibilities of objects in an application and how they communicate between them at runtime. The Behavioral patterns that I already wrote in this series of the GoF patterns are the Command, Chain of Responsibility, Iterator, Mediator, Interpreter, Memento, Observer and State patterns. In this post, I will discuss the Strategy Pattern – one of the most fundamental design pattern that all programmers should possess in their design toolkit. The intent of the Strategy Pattern, as mentioned above, suggests that this pattern is applicable when you have multiple algorithms and you want to treat them as independent objects that can be interchanged dynamically at runtime to achieve high cohesion and loose coupling in your application.

## Strategy Pattern: Introduction

In enterprise applications, you will often have objects that use multiple algorithms to implement some business requirements. A common example is a number sorting class that supports multiple sorting algorithms, such as bubble sort, merge sort, and quick sort. Similarly, a file compression class can support different compression algorithm, such as ZIP, GZIP, LZ4, or even a custom compression algorithm. Another example can be a data encryption class that encrypts data using different encryption algorithms, such as AES, TripleDES, and Blowfish. Typically, programmers tend to bundle all the algorithm logic in the host class, resulting in a monolithic class with multiple switch case or conditional statements. The following example shows the structure of one such class that supports multiple algorithms to encrypt data.

```java
public class Encryptor {
    private String algorithmName;
    private String plainText;
    public Encryptor(String algorithmName){
       this.algorithmName=algorithmName;
   }
    public void encrypt(){
        if (algorithmName.equals("Aes")){
            System.out.println("Encrypting data using AES algorithm");
            /*Code to encrypt data using AES algorithm*/
        }
       else if (algorithmName.equals("Blowfish")){
            System.out.println("Encrypting data using Blowfish algorithm");
            /*Code to encrypt data using Blowfish algorithm */
        }
        /*More else if statements for other encryption algorithms*/
    }
   /*Getter and setter methods for plainText*/

}
```

The above Encryptor class has conditional statements for different encryption algorithms. At runtime, the code loops through the statements to perform encryption based on the client specified algorithm. The result is a tightly coupled and rigid software that is difficult-to-change. You can imagine the consequences if you try to implement a new encryption algorithm, say TripleDES or a custom encryption algorithm. You’d have to open and modify the Encryptor class. Also, if an existing algorithm needs to be changed, the Encryptor class will again require modification. As you can see, our Encryptor class is a clear violation of the Open Closed principle – one of the SOLID design principles. As per the principle, new functionality should be added by writing new code, rather than modifying existing code. The violation occurred because we did not follow the fundamental tenet of Object-Oriented (OO) programming practice that states “encapsulate what varies”.

Such pitfalls in enterprise applications result in rippling effects across the application making the application fragile, and you can avoid them by using the Strategy pattern. Using the Strategy pattern, we define a set of related algorithm and encapsulate them in classes separated from the host class (Encryptor). Clients can choose the algorithm to use at run time. By doing so, we can easily add a new algorithm or remove an existing one without modifying the other existing algorithms or the host class. Each of the algorithm classes adhere to the Single Responsibility principle, another SOLID principle as they will only be concerned with encrypting data with a specific algorithm, which is currently lacking in our Encryptor class. In addition, with smaller algorithm classes, unit testing becomes easier to focus on testing one particular situation.

## Participants of the Strategy Pattern

Using our example of data encryption, we will first implement the interface that will be used by all of the different encryption algorithm-specific classes. Let’s name the interface EncryptionStrategy and name the algorithm specific classes AesEncryptionStrategy and BlowfishEncryptionStrategy. Ultimately, these are our strategies.

We will next refactor the Encryptor class to remove all conditional statements and delegate any encryption request to an algorithm-specific class that the client specifies.

We can now summarize the participants of the strategy pattern as:

- Strategy (EncryptionStrategy): Is an interface common to all supported algorithm-specific classes.
- ConcreteStrategy (AesEncryptionStrategy and BlowfishEncryptionStrategy): Implements the algorithm using the Strategy interface.
- Context (Encryptor): Provides the interface to client for encrypting data. The Context maintains a reference to a Strategy object and is instantiated and initialized by clients with a ConcreteStrategy object.

## Applying the Strategy Pattern

Let’s now apply the Strategy pattern to implement the same requirement done for the Encryptor class we wrote earlier. We will start with the Strategy interface and then move to the ConcreteStrategy classes.

EncryptionStrategy.java

```java
package guru.springframework.gof.strategy.strategies;
public interface EncryptionStrategy {
    void encryptData(String plainText);
}
```

AesEncryptionStrategy.java

```java
package guru.springframework.gof.strategy.strategies;
import javax.crypto.Cipher;
import javax.crypto.KeyGenerator;
import javax.crypto.SecretKey;
public class AesEncryptionStrategy implements EncryptionStrategy{
   @Override
    public void encryptData(String plaintext) {
       System.out.println("-------Encrypting data using AES algorithm-------");
       try {
           KeyGenerator keyGenerator = KeyGenerator.getInstance("AES");
           keyGenerator.init(128);
           SecretKey secretKey = keyGenerator.generateKey();
           byte[] plaintTextByteArray = plaintext.getBytes("UTF8");
           Cipher cipher = Cipher.getInstance("AES");
           cipher.init(Cipher.ENCRYPT_MODE, secretKey);
           byte[] cipherText = cipher.doFinal(plaintTextByteArray);
           System.out.println("Original data: " + plaintext);
           System.out.println("Encrypted data:");
           for (int i = 0; i < cipherText.length; i++) {
               System.out.print(cipherText[i] + " ");
           }
       }
           catch(Exception ex){
               ex.printStackTrace();
           }
       }
   }
```

BlowfishEncryptionStrategy.java

```java
package guru.springframework.gof.strategy.strategies;
import javax.crypto.Cipher;
import javax.crypto.KeyGenerator;
import javax.crypto.SecretKey;
public class BlowfishEncryptionStrategy implements EncryptionStrategy{
    @Override
    public void encryptData(String plaintext) {
        System.out.println("\n-------Encrypting data using Blowfish algorithm-------");
        try {
            KeyGenerator keyGenerator = KeyGenerator.getInstance("Blowfish");
            keyGenerator.init(128);
            SecretKey secretKey = keyGenerator.generateKey();
            byte[] plaintTextByteArray = plaintext.getBytes("UTF8");
            Cipher cipher = Cipher.getInstance("Blowfish");
            cipher.init(Cipher.ENCRYPT_MODE, secretKey);
            byte[] cipherText = cipher.doFinal(plaintTextByteArray);
            System.out.println("Original data: " + plaintext);
            System.out.println("Encrypted data:");
            for (int i = 0; i < cipherText.length; i++) {
                System.out.print(cipherText[i] + " ");
            }
        }
        catch(Exception ex){
            ex.printStackTrace();
        }
    }
}
```

In the EncryptionStrategy interface we wrote above, we declared a single encryptData() method that both the AesEncryptionStrategy and BlowfishEncryptionStrategy implements. In the AesEncryptionStrategy class, we used the AES encryption algorithm to symmetrically encrypt the string passed to encryptData(). We performed the encryption using the cryptography classes of the javax.crypto package. After encryption, we printed out the encrypted string. We performed the same function in the BlowfishEncryptionStrategy class, but this time using the Blowfish encryption algorithm.

Next, we will write the Context – The Encryptor class.

Encryptor.java

```java
package guru.springframework.gof.strategy.context;
import guru.springframework.gof.strategy.strategies.EncryptionStrategy;
public class Encryptor {
    private EncryptionStrategy strategy;
    private String plainText;
    public Encryptor(EncryptionStrategy strategy){
        this.strategy=strategy;
    }
    public void encrypt(){
        strategy.encryptData(plainText);
    }
    public String getPlainText() {
        return plainText;
    }
    public void setPlainText(String plainText) {
        this.plainText = plainText;
    }
}
```

Notice that our Encryptor class now doesn’t have any conditional statements. It maintains a reference to Strategy, which is initialized through the constructor. Notice, we don’t have any reference to any ConcreteStrategy (AesEncryptionStrategy and BlowfishEncryptionStrategy) classes. We are adhering to the OOP programming practice of “Program to an interface, not an implementation”. So, if we later want to accommodate a new encryption algorithm, say Triple DES, all we need to do is create a new ConcreteStrategy class, say TripleDesEncryptionStrategy. This class will implement our EncryptionStrategy interface and we are good to go without modifying our Encryptor class. Our Encryptor class is now open for extension and closed for modification – It’s now following the Open Close principle. The Encryptor class also contains a encrypt() method that clients will call to perform encryption. Instead of putting in any encryption logic in this method, we delegated any call to this method to the associated algorithm-specific object.

Let’s write a test class for our example.

EncryptorTest.java

```java
package guru.springframework.gof.strategy.context;
import guru.springframework.gof.strategy.strategies.AesEncryptionStrategy;
import guru.springframework.gof.strategy.strategies.BlowfishEncryptionStrategy;
import guru.springframework.gof.strategy.strategies.EncryptionStrategy;
import org.junit.Test;
import static org.junit.Assert.*;
public class EncryptorTest {
    @Test
    public void testEncrypt() throws Exception {
      EncryptionStrategy aesStrategy=new AesEncryptionStrategy();
      Encryptor aesEncryptor=new Encryptor(aesStrategy);
      aesEncryptor.setPlainText("This is plain text");
      aesEncryptor.encrypt();
        EncryptionStrategy blowfishStrategy=new BlowfishEncryptionStrategy();
        Encryptor blowfishEncryptor=new Encryptor(blowfishStrategy);
        blowfishEncryptor.setPlainText("This is plain text");
        blowfishEncryptor.encrypt();
    }
}
```

In the EncryptorTest class above, from Line 15 – 18 we created an AesEncryptionStrategy object and passed it to the constructor while instantiating the Encryptor class. We then called the setPlainText() method of Encryptor to set the plain text to encrypt and then called the encrypt() method to perform the encryption using AES. From Line 20 – Line 23, we performed the same steps but this time, we switched the encryption algorithm to Blowfish by using the BlowfishEncryptionStrategy class.

The output of the test is this.

---

## T E S T S

Running guru.springframework.gof.strategy.context.EncryptorTest
-------Encrypting data using AES algorithm-------
Original data: This is plain text
Encrypted data:
97 11 54 73 -21 92 109 -124 -120 110 -43 20 123 99 87 120 120 95 70 -63 -8 70 41 44 -73 -94 48 127 61 43 -96 110
-------Encrypting data using Blowfish algorithm-------
Original data: This is plain text
Encrypted data:
3 -34 -23 -74 -61 -55 99 -114 71 -113 124 -57 -65 -45 -128 37 -123 -83 118 107 42 -123 84 14
Tests run: 1, Failures: 0, Errors: 0, Skipped: 0, Time elapsed: 7.72 sec - in guru.springframework.gof.strategy.context.EncryptorTest

## Summary

To many, the Strategy and State patterns appear similar. It’s true that the structure of both the patterns are similar. It’s the intent that differs – that is, they solve different problems. The State pattern aims to facilitate state transition while the aim of the Strategy pattern is to change the behavior of a class by changing internal algorithm at runtime without modifying the class itself.

There is a lot of debate around the use of the Strategy Pattern with Spring. Often you’ll see the Strategy Pattern used in conjunction with Dependency Injection, where Springs IoC container is making the choice of which strategy to use. Different data sources as a great example. Using a H2 data source for local development is one strategy. Using MySQL for production is another strategy. Which one is to use at runtime is up to the Spring IoC container.

Spring 3, introduced a type conversion factory. In this, you provide a type converter, which implements the Converter interface. At run time, your code can ask the converter factory for the proper converter. This sure sounds like the Strategy Pattern, doesn’t it?
