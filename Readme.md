# Simple Bank Application

A command-line banking application written in Go that allows users to check balance, deposit money, and withdraw money.

## Features

- **Check Balance**: View current account balance
- **Deposit Money**: Add funds to your account
- **Withdraw Money**: Remove funds from your account
- **Exit**: Close the application

## Prerequisites

- Go 1.x or higher installed on your system

## Installation

1. Clone the repository or download the source code:
```bash
git clone [your-repository-url]
```

2. Navigate to the project directory:
```bash
cd bank-application
```

3. Build the application:
```bash
go build -o bankapp
```

## Usage

Run the application:
```bash
go run main.go
```
Or if you built it:
```bash
./bankapp
```

### Menu Options

1. **Check Balance** - Displays your current account balance
2. **Deposit Money** - Enter an amount to add to your account
3. **Withdraw Money** - Enter an amount to withdraw from your account
4. **Exit** - Close the application

## Example

```
welcome to the bank:what do you want to do:
1.check balance:
2.deposit money:
3.withdraw money:
4.exit:
your chioce:1
Available amount:bank amount: 1000000
```

## Code Structure

- `main()` - Entry point of the application
- `userPrompt()` - Displays menu and gets user input
- `match()` - Handles user menu selection
- `checkAmount()` - Displays current balance
- `depositAmount()` - Processes deposits
- `withdrawAmount()` - Processes withdrawals

## Initial Balance

The application starts with an initial balance of $1,000,000.

## Notes

- The application currently runs only once and exits after performing one operation
- Input validation is implemented for menu choices
- All monetary values are handled as float64



