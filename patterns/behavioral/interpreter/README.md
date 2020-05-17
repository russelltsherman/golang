# Interpreter Pattern

Interpreter pattern is part of the Behavioral pattern family of the Gang of Four design patterns. Behavioral patterns address responsibilities of objects in an application and how they communicate between them. We have already learned about the Command, Chain of Responsibility, Iterator, Mediator, and Memento Behavioral patterns. In this post we’ll learn about the Interpreter Pattern in the Java programming language.

## Interpreter Pattern: Introduction

The best analogy of the Interpreter pattern are we, who at some times need to interpret gestures. We need to observe a gesture, and present an interpretation. Based on the culture, our interpretation may have different significance and it is our interpretation that will give the gesture different meaning.

Similarly, in the programming world, applications can receive user input from a wide variety of sources, such as browsers, GUI, command line tools, and mobile phones. The input, can be expressions in different formats, such as mathematical expressions following one of the Infix, Prefix, or Postfix notations. Also, when a new type of input format is introduced, we don’t want to change the client code sending the input. The solution is to use the Interpreter pattern that allows automated and flexible processing of expression provided as input by users through client code.

The purpose of using the interpreter pattern is to process user input expressions and build an Abstract Syntax Tree, which we will refer as AST. This AST is an instance of the Composite pattern which I wrote about earlier. You then need a parser to parse the AST and produce the output.

When we are talking about input expressions, you can relate them with mathematical expressions that follows one of the notation, such as Reverse Polish Notation (RPN) (Postfix) in which every operator follows all of its operands. An example of an expression with Postfix notation is 2 1 5 + *, which the Java compiler will interpret as 2*1+5 that will result into 12.

Expressions defined by the interpreter pattern can be:

Terminal: Are the leaf nodes of the tree. They don’t contain other expressions. Terminal expressions in our Postfix mathematical expression example are the operands, 2, 1, and 5.
Nonterminal: Are the non-leaf nodes of the tree. They contain other expressions. The +and \* operators are nonterminal expressions.
The AST for our Postfix expression is this.

The easiest way to understand how the Interpreter Pattern works is by looking at an example that will lead us to the participants of the pattern.

## Participants of the Interpreter Pattern

We will apply the Interpreter Pattern to interpret a string of a mathematical expression with postfix notation. To do so, we start by creating an Expression interface and a NumberExpression subclass to represent numbers (terminal expressions). For the nonterminal expressions, we create a class for each of the operators by implementing the Expression interface. Let’s name them AdditionExpression, SubtractionExpression, and MultiplicationExpression. To parse the abstract syntax tree, we will create an ExpressionParser class.

￼We will now summarize the participants of the Interpreter pattern.

AbstractExpression (Expression): Declares an interpret() operation that all nodes (terminal and nonterminal) in the AST overrides.
TerminalExpression (NumberExpression): Implements the interpret() operation for terminal expressions.
NonterminalExpression (AdditionExpression, SubtractionExpression, and MultiplicationExpression): Implementoperations the interpret() operation for all nonterminal expressions.
Context (String): Contains information that is global to the interpreter. It is this String expression with the Postfix notation that has to be interpreted and parsed.
Client (ExpressionParser): Builds (or is provided) the AST assembled from TerminalExpression and NonTerminalExpression. The Client invokes the interpret() operation.
The following figure illustrates the relationship between the participants of the Interpreter pattern.Relationship between the participants of the Interpreter Pattern

## Applying the Iterator Pattern

We will start applying the Interpreter pattern to parse Postfix mathematical expressions by writing the Expression interface (AbstractExpression), and the NumberExpression class (TerminalExpression).

Expression.java

```java
package guru.springframework.gof.interpreter;
public interface Expression {
    int interpret();
}
```

NumberExpression.java

```java
package guru.springframework.gof.interpreter;
public class NumberExpression implements Expression{
    private int number;
    public NumberExpression(int number){
        this.number=number;
    }
    public NumberExpression(String number){
        this.number=Integer.parseInt(number);
    }
    @Override
    public int interpret(){
        return this.number;
    }
}
```

The Expression interface we wrote has a single interpret() method. In the NumberExpression class, we wrote two constructor to initialize an integer instance variable and returned the variable from the overridden interpret() method.

We will next write the NonterminalExpression classes (AdditionExpression, SubtractionExpression, and MultiplicationExpression).

AdditionExpression.java

```java
package guru.springframework.gof.interpreter;
public class AdditionExpression implements Expression {
    private Expression firstExpression,secondExpression;
    public AdditionExpression(Expression firstExpression, Expression secondExpression){
        this.firstExpression=firstExpression;
        this.secondExpression=secondExpression;
    }
    @Override
    public int interpret(){
        return this.firstExpression.interpret()+this.secondExpression.interpret();
    }
    @Override
    public String toString(){
        return "+";
    }
}
```

SubtractionExpression.java

```java
package guru.springframework.gof.interpreter;
public class SubtractionExpression implements Expression{
    private Expression firstExpression,secondExpression;
    public SubtractionExpression(Expression firstExpression, Expression secondExpression){
        this.firstExpression=firstExpression;
        this.secondExpression=secondExpression;
    }
    @Override
    public int interpret(){
        return this.firstExpression.interpret()-this.secondExpression.interpret();
    }
    @Override
    public String toString(){
        return "-";
    }
}
```

MultiplicationExpression.java

```java
package guru.springframework.gof.interpreter;
public class MultiplicationExpression implements Expression{
    private Expression firstExpression,secondExpression;
    public MultiplicationExpression(Expression firstExpression, Expression secondExpression){
        this.firstExpression=firstExpression;
        this.secondExpression=secondExpression;
    }
    @Override
    public int interpret(){
        return this.firstExpression.interpret()*this.secondExpression.interpret();
    }
    @Override
    public String toString(){
        return "*";
    }
}
```

In all of the classes above, we declared two variables of type Expression and initialized them through the constructor. It is in the interpret() method of each class that differs. In each class, we called the interpret() method of the Expression variables, performed the corresponding mathematical operation, and returned back the result. Let’s now write our expression parser. We will start with a utility class that our expression parser will use.

ParserUtil.java

```java
package guru.springframework.gof.interpreter;
public class ParserUtil {
    public static boolean isOperator(String symbol) {
        return (symbol.equals("+") || symbol.equals("-") || symbol.equals("*"));
    }
    public static Expression getExpressionObject(Expression firstExpression,Expression secondExpression,String symbol){
        if(symbol.equals("+"))
            return new AdditionExpression(firstExpression,secondExpression);
        else if(symbol.equals("-"))
            return new SubtractionExpression(firstExpression,secondExpression);
        else
            return new MultiplicationExpression(firstExpression,secondExpression);
    }
}
```

The ParserUtil class we wrote above have two methods. The isOperator() method returns a true Boolean value if the String passed to it is an operator. The getExpression() method returns an Expression instance based on the two Expression objects and the symbol passed to it.

We will now write the expression parser.

ExpressionParser.java

```java
package guru.springframework.gof.interpreter;
import java.util.Stack;
public class ExpressionParser {
    Stack stack=new Stack<>();
    public int parse(String str){
        String[] tokenList = str.split(" ");
        for (String symbol : tokenList) {
            if (!ParserUtil.isOperator(symbol)) {
                Expression numberExpression = new NumberExpression(symbol);
                stack.push(numberExpression);
                System.out.println(String.format("Pushed to stack: %d", numberExpression.interpret()));
            } else  if (ParserUtil.isOperator(symbol)) {
                Expression firstExpression = stack.pop();
                Expression secondExpression = stack.pop();
                System.out.println(String.format("Popped operands %d and %d",
                        firstExpression.interpret(), secondExpression.interpret()));
                Expression operator = ParserUtil.getExpressionObject(firstExpression, secondExpression, symbol);
                System.out.println(String.format("Applying Operator: %s", operator));
                int result = operator.interpret();
                NumberExpression resultExpression = new NumberExpression(result);
                stack.push(resultExpression);
                System.out.println(String.format("Pushed result to stack: %d", resultExpression.interpret()));
            }
        }
       int result= stack.pop().interpret();
        return result;
    }
}
```

In the ExpressionParser class above, we wrote a parse() method that generates a token of symbol from the expression passed to it. It starts from the beginning of the expression and loops through each token one by one till the end of the expression. For each symbol, it checks with the ParserUtil class, whether it is an operator or not. If the symbol is not an operator it pushes the NumberExpression that represents the numeric symbol to a stack. If it is an operator, the code pop two elements from the stack, applies the operator, and push back the result to the stack. Things will get even more clear once we write a test and observe the output.

ExpressionParserTest.java

```java
package guru.springframework.gof.interpreter;
import org.junit.Test;
import static org.junit.Assert.*;
public class ExpressionParserTest {
    @Test
    public void testParse() throws Exception {
      String input="2 1 5 + *";
       ExpressionParser expressionParser=new ExpressionParser();
       int result=expressionParser.parse(input);
       System.out.println("Final result: "+result);
    }
}
```

The output of the test is this.

---

## T E S T S

Running guru.springframework.gof.interpreter.ExpressionParserTest
Pushed to stack: 2
Pushed to stack: 1
Pushed to stack: 5
Popped operands 5 and 1
Applying Operator: +
Pushed result to stack: 6
Popped operands 6 and 2
Applying Operator: \*
Pushed result to stack: 12
Final result: 12
Tests run: 1, Failures: 0, Errors: 0, Skipped: 0, Time elapsed: 0.001 sec - in guru.springframework.gof.interpreter.ExpressionParserTest

## Summary

It may seem that the Interpreter Pattern is limited to situations in which it can be applied. Thus many developers tend to overlook it. As you grow as a developer, you will realize the importance of this design pattern. Consider what the GoF said about the Interpreter Pattern: “If a particular kind of problem occurs often enough, then it might be worthwhile to express instances of the problem as sentences in a simple language. Then you can build an interpreter that solves the problem by interpreting these sentences. For example, searching for strings that match a pattern is a common problem. Regular expressions are a standard language for specifying patterns of strings. Rather than building custom algorithms to match each pattern against strings, search algorithms could interpret a regular expression that specifies a set of strings to match”.

In a nutshell, the Interpreter Pattern allows you to create almost a mini language to implement program logic. For example, scenarios like regular expressions and interpreting mathematical expression, which I already explained. In Java SE, java.util.Pattern and subclasses of java.text.Format are some of the examples of interpreter pattern being used.
