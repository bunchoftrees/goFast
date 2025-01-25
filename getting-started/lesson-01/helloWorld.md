# Writing Your First Code in Go

Welcome to your first Go programming lesson! In this guide, we will walk you through writing and running your first Go program.

## Prerequisites

Before you begin, ensure you have the following installed on your system:
- Go (https://golang.org/dl/)
- A text editor or IDE of your choice
    * You can get **VS Code** here: (https://code.visualstudio.com/Download)  

## Steps

1. **Create a New Directory**

    Open your terminal and create a new directory for your project:
    ```sh
    mkdir helloWorld
    cd helloWorld
    ```

2. **Create Your First Go File**

    Create a new file named `main.go` in the `helloWorld` directory and open it in your text editor. 
    ```sh
    touch helloWorld.go
    ```

    
    Add the following code:
    ```go
    package main

    import "fmt"

    func main() {
         fmt.Println("Hello, World!")
    }
    ```

    ### Explanation

    - **Package Declaration (`package main`)**: In Go, every file belongs to a package. The `main` package is a special package that tells the Go compiler that the program should be compiled as an executable program. The `main` package must have a `main` function, which is the entry point of the program.
    
    - **Import Statement (`import "fmt"`)**: The `import` statement is used to include the code from other packages. The `fmt` package is part of the Go standard library and provides functions for formatted I/O, including printing to the console. Here, we use it to print "Hello, World!" to the console.
    
    - **Main Function (`func main()`)**: The `main` function is the entry point of the program. When the program is executed, the code inside the `main` function is run. In this case, it calls `fmt.Println("Hello, World!")` to print "Hello, World!" to the console.

4. **Run Your Go Program**

    Save the file and return to your terminal. Run your Go program with the following command:
    ```sh
    go run main.go
    ```

    You should see the output:
    ```
    Hello, World!
    ```

## Conclusion

Congratulations! You've written and executed your first Go program. Continue exploring Go by experimenting with more code and building new projects.

Happy coding!
