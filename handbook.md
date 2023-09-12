<h1>Handbook on Programming in Go</h1>

**AP Computer Science Principles 2023-2024**

Camila Granja

<!-- This is a comment (which will not be displayed in the live file);
replace all "???" with your own text. -->




___





<h1>Table of Contents</h1>

- [1. Compiling and Running](#1-compiling-and-running)
- [2. Data Types](#2-data-types)
- [3. Console I/O](#3-console-io)
- [4. Arithmetic Operations](#4-arithmetic-operations)
- [5. Assignment Operations](#5-assignment-operations)
- [6. Comments](#6-comments)
- [7. Decision Structures](#7-decision-structures)
- [8. Conditional Operators](#8-conditional-operators)
- [9. Logic Operators](#9-logic-operators)
- [10. Advanced Decision Structures](#10-advanced-decision-structures)
- [11. String Methods](#11-string-methods)
- [12. Random Generation](#12-random-generation)
- [13. Looping Structures](#13-looping-structures)
- [14. Functions/Methods](#14-functionsmethods)
- [15. Elementary Data Structures](#15-elementary-data-structures)
  - [15.1 Arrays/Lists](#151-arrayslists)
  - [15.2 Matrices](#152-matrices)
- [References](#references)

<!-- 
- [16. Major Keywords](#16-major-keywords)
- [17. Error Handling](#17-error-handling)
- [18. Working with Files](#18-working-with-files)
- [19. Major Language Features](#19-major-language-features)
  - [19.1 Classes](#191-classes)
  - [19.2 Inheritance](#192-inheritance)
  - [19.3 Generic Typing (Templates)](#193-generic-typing-templates)
  - [19.4 Pointers](#194-pointers)
- [20. Importing Local Libraries](#20-importing-local-libraries)
- [21. Working with Time](#21-working-with-time)
- [22. Importing Libaries from Package managers](#22-importing-libaries-from-package-managers)
- [23. Bitwise Operators](#23-bitwise-operators)
- [24. Common Data Structures](#24-common-data-structures)
- [25. Advanced Language Features](#25-advanced-language-features)
-->




___





# 1. Compiling and Running
* "go run filename.go" to run





___





# 2. Data Types
|type|desc|ex|
|--|--|--|
|int|Whole numbers pos or neg, integers|5, -4, 3|
|float|Numbers with decimals|5.4, 3.2, 8.7455|
|string|Sequence of characters|"Hello", "HI!!"|
|bool|True or false|true, false|

include type of ints + floats
```LANGUAGE_HERE

```





___





# 3. Console I/O
### Input
* "fmt.Scanf("%type", &var)"

### Output
* "fmt.Printf("Words")"


___





# 4. Arithmetic Operations
|Operator|Desc|Ex.|
|--|--|--|
|+|Adds|5 + 2 = 7|
|-|Subtracts|5 - 2 = 3|
|*|Multiplies||
|/|Divides||
|%|Modulus, gives the remainder||
|++|Increment operator, adds 1||
|--|Increment operator, subtracts 1||





___





# 5. Assignment Operations

???





___





# 6. Comments

???





___





# 7. Decision Structures

???





___





# 8. Conditional Operators

???





___





# 9. Logic Operators

???





___





# 10. Advanced Decision Structures

???





___





# 11. String Methods

???





___





# 12. Random Generation

???





___





# 13. Looping Structures
* Go only has for loops

**for loop**
```go
for i := 0; i < 10; i++ {
		sum += i
}
```
* i := 0 initializes the loop control var
* i < 10 checks that i is less than 10, if it is the loop runs
* i++ adds one to i after the loop is done running

**while loop**
```go
i := 0
for i < 5 {
  fmt.Print(i)
  i++
}
```
* skip the init, just put the condition expression
* loop will run until condition expression is false

**for each loop**
```go
nums = []int{1, 2, 3, 4}
for i, n := range nums {
  fmt.Print(n)
}

```
___





# 14. Functions/Methods

???





___





# 15. Elementary Data Structures

???





## 15.1 Arrays/Lists

???






## 15.2 Matrices

???





___





<!-- 
EVERYTHING BELOW IS OPTIONAL; 
UNCOMMENT BY REMOVING THE ARROW TAGS SURROUNDING
(i.e., delete the "< !--" and "-- >" tags)

CHANGE THE SECTION NUMBERS AS DESIRED
-->

<!-- # 16. Major Keywords

???





___ -->





<!-- # 17. Error Handling

???





___ -->





<!-- # 18. Working with Files

???





___ -->





<!-- # 19. Major Language Features

???







## 19.1 Classes

???





## 19.2 Inheritance

???





## 19.3 Generic Typing (Templates)

???





## 19.4 Pointers

???





___ -->





<!-- # 20. Importing Local Libraries

???





___ -->





<!-- # 21. Working with Time

???





___ -->





<!-- # 22. Importing Libaries from Package managers

???





___ -->





<!-- # 23. Bitwise Operators

???





___ -->





<!-- # 24. Common Data Structures

???





___ -->





<!-- # 25. Advanced Language Features

???





___ -->





# References

* [Markdown Cheatsheet](https://gist.github.com/jonschlinkert/5854601)
* [description](http://example.com)