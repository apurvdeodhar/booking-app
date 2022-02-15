## Notes
- Multithreading
  - Run multiple processes at the same time as watch video on YouTube and comment at the same time.
  - Multiple users writing to same google Doc
- Go was built to run multiple cores and **built to support concurrency**
  - Concurrency is cheap and easy.
- Characteristics of Go
  - simple syntax like python
  - efficiency and safety like c++
  - fast build time, start up and run
  - requires fewer resources
  - compiled lang like python
- It is used on backend oder server side.


- Go always needs to know where to start the program? OR what's the entrypoint?
  - meaning it needs `main()`
  ```go
  package main
  
  //main package a bunch of functions 
  //one of them is "fmt"
  
  import "fmt"
  ```
- To find which function belongs to which package [see](https://pkg.go.dev/fmt)
- Variable names are in `camelCase`
  - use `const conferenceTickets = 314` for keeping the value constant throughout the app.
  - for `%v`, `%s` in `fmt.printf("This is my name: %v", name)` see [fmt package](https://pkg.go.dev/fmt) 


- Data types
  - ```go 
    "this is a string"
    // declare and assing varible at the same time
    // so instead var name string = "Jim" this is called syntactic suger
    name := "Jim"
    ```
- Difference between scan and scanf [ref](https://fernandocorreia.dev/posts/learning-golang-part-18/)

- Arrays
  - Have fixed size see `var booking` in the code.
  - Can have only single data type mixed does not work in go.
  - ```go
      var bookings = [50]string{}
      // define an empty array with size and datatype
      var bookings [50]string
      // can also be written like this it is called array type in go
    ```
    
- Slices
  - Abstraction of array.
  - Arrays without size like `var bookings []string`

- Loops
  - Go just has a **for** loop for all use cases.
  - If statement
    - ```go
        if remainingTickets == 0 {
            // do something
      } 
        else {
            // do or skip something
      }
      ```
- Functions
  - If you want to return values use this or see getFirstNames
  - ```go
    func getFirstNames (booking []string) [] string{
        // bla bla bla
        return something
    }
    ```
  - Interesting thing in Go inside `()` of function takes arguments but outside `()` it needs to know what type it will return
  - Best thing until now 💪🏽, **You can return any no of values from a func unlike other programming languages.** 
    See `userInputValidate` specially `(bool,bool,bool)` and line `36` where it is called.

- Package level variables
  - Just define them on package level
  - These variables can not be created shorthanded using `:=` rather define them with `var foo string = "foobar"` 

- Packages
  - If you want to use a func from imported package let's say `booking-app/helper` has been imported and now func `userInputValidate` needs to be used in `main.go`
    - **This func `userInputValidate` needs to be exported from `booking-app/helper`** by just making first letter capital `UserInputValidate` 
      (same applies for var names) just like `fmt.Println`
