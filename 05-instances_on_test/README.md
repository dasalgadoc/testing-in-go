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