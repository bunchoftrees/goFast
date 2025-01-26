# Defining Variables in Go

In Go, variables can be defined in several ways. 

## Variable Initialization and Assignment
In Go, all variables are initialized with a default value, even when a value is unassigned. You can assign a value after initializing a variable too.

```go
var message string
message = "Hello, World!"
fmt.Println(message)
```

In this example, we first initialize a `string` variable `message`, and on the next line assign a value and print the value to the terminal.

## Type Inference
Go can infer the type of a variable based on the initial value assigned to it. You don't need to explicitly specify the type.

```go
var name = "John Bear"
fmt.Println(name)
```
Here, `name` is inferred to be of type `string`.

## Explicit Type Declaration
You can also explicitly declare the type of a variable when you define it.

```go
var age int = 38
fmt.Println(age)
```

In this example, `age` is explicitly declared as an `int`.

## Multiple Variable Declaration
You can declare and initialize multiple variables in a single line.

```go
var city, state, year = "Chicago", "Illinois", 2025
fmt.Println(city, state, year)
```

Here, `city` and `state` are inferred to be of type `string`, and `year` is inferred to be of type `int`.

## Multi-line Variable Declaration
Sometimes it makes sense to define many variables. 

```go 
var (
		planet = "Earth"
		distance = 92.96
		unit = "million miles"
		life = true
	)
	fmt.Println(planet, distance, unit, life)
```

You can also explicitly define type in this longer form definition.

```go
var (
		planet string = "Earth"
		distance float32 = 92.96
		unit string = "million miles"
		life bool = true
	)
	fmt.Println(planet, distance, unit, life)
```

In this second example, we see the _types_ `string`, `float32`, and `bool`.

## Short Variable Declaration
Inside functions, you can use the shorthand `:=` to declare and initialize variables.

```go
animal := "Bear"
fmt.Println(animal)
```

In this case, `animal` is inferred to be of type `string`.

## Short Variable Declaration with Multiple Variables
You can also declare and initialize multiple variables using the shorthand `:=`. 

It's important to note that this special shorthand, `:=` can only be used in _functions_, which will be discussed in a later lesson.

```go
food, drink, cute := "Honey", "Water", true
fmt.Println(food, drink, cute)
```

Here, `food` and `drink` are inferred to be of type `string`, and cute is inferred to be of type `bool`.

It's easy to declare variables in Go!
