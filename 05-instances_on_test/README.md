# 🧶 Generating complex instances on test set

Generating data for testing is a common task. There are several approaches to do it. In this section we will see some of them.

## 👨🏻‍🦳 Traditional approach

The traditional approach is to create a function that returns a new instance of the type we want to test. This function will be used in the test set up.

```go
func traditionalInstance(name string, accessLevel int) User {
    userName, err := NewUserName(name)
    if err != nil {
    	log.Panic(err)
    	panic(err)
    }
    userAccessLevel, err := NewUserAccessLevelFromInt(accessLevel)
    if err != nil {
    	log.Panic(err)
    	panic(err)
    }
    user, err := NewUser(*userName, *userAccessLevel)
    if err != nil {
    	log.Panic(err)
    	panic(err)
    }
    
    return *user
}
```

## 🏢 Builder pattern

Builder pattern is a creational design pattern that lets you construct complex objects step by step. The pattern allows you to produce different types and representations of an object using the same construction code.

In test, we can use this pattern to create complex instances of the type we want to test with the exact data we want to test.

This method gain a lot in semantics and readability, but can be verbose has some integrity problems.

```go
user := NewUserBuilder().
   	WithAccessLevel(2).
   	Build()
```

## 👩‍👧 Object mother pattern with random data

The object mother pattern is a creational design pattern that allows you to create instances of a type with random data. This pattern is useful when you want to test the behavior of your code with different data.

In test, this pattern takes the form of a small piece of code that returns a new instance of the type we want to test with random data.

You can combine builder with the random data generation to create complex instances with half define and half random data.

- Object mother is a set of static functions.

## ❗️ Object mother with named parameters

Named parameters are a language feature that allows you to pass parameters to a function in any order. This feature is not available in Go, but we can emulate it with a function that receives a map of parameters.

In test, we can use this feature to create complex instances of the type we want to test with the exact data we want to test and the rest of pieces of information can be random.

```go
// named parameters emulation
testData := map[string]any{
    "name":        "John Doe",
    "accessLevel": 2,
}

func GenerateRandomInstance(data map[string]any) *SomeStruct {
	// ... generate random data or use values defined in data map
}
```

### 🧪 Object mother as solution to test Value Objects and Aggregates.

Value Objects are objects that represent a value and they are immutable, two value objects with the same value are equal.
Aggregate root can be a set of value objects and entities that are related to each other.

Object mothers are a good solution to test Value Objects and Aggregates for the next reasons:

- The initialization of values objects is isolated by object mothers, thus if the domain modeling changes, we simply need to alter the object mother.
- Random data makes it possible to establish different values for each test, which is helpful for testing the behavior of the code with various data while concentrating just on supply the test determinant data.
- Data can be generated randomly through libraries like [gofakeit](https://github.com/brianvoe/gofakeit) or [faker](https://github.com/icrowley/fake). Remember generate deterministic test!, For instance random age on an overage test.
- Random data allows you to generate some unexpected cases from exotic data.
