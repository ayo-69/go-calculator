# Calculator

This is a simple calculator application written in Go. It performs basic arithmetic operations such as addition, subtraction, multiplication, and division.

## Getting Started

### Prerequisites

- Go 1.23.4 or later

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/ayo-69/calculator.git
    ```
2. Navigate to the project directory:
    ```sh
    cd calculator
    ```

### Usage

1. Run the application:
    ```sh
    go run main.go
    ```
2. Use the following commands to perform operations:
    - Addition: `calc add <num1> <num2>`
    - Subtraction: `calc sub <num1> <num2>`
    - Multiplication: `calc mul <num1> <num2>`
    - Division: `calc div <num1> <num2>`

### Example

```sh
go run main.go add 5 3
# Output: A + B = 8

go run main.go sub 5 3
# Output: A - B = 2

go run main.go mul 5 3
# Output: A x B = 15

go run main.go div 6 3
# Output: A / B = 2
```

## Project Structure

- `main.go`: The main file containing the calculator logic.
- `go.mod`: The Go module file.

## Functions

- `add(x, y float64)`: Adds two numbers.
- `sub(x, y float64)`: Subtracts the second number from the first.
- `mul(x, y float64)`: Multiplies two numbers.
- `div(x, y float64)`: Divides the first number by the second.
- `printAnwser(value float64, operator string)`: Prints the result of the operation.
- `parseArgs(args []string) (float64, float64)`: Parses the command line arguments into float64 numbers.

## License

This project is licensed under the MIT License.
