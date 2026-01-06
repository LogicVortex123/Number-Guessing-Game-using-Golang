# Description

This is a command-line number guessing game developed in Go.
The program generates a random number between 1 and 100, and the user has to guess it.
After each guess, the program tells the user whether the guess is too low, too high, or correct.
It also keeps track of the number of attempts.

This project demonstrates basic Go programming concepts such as loops, conditionals, user input, random number generation, and error handling.

# Features

1. Generates a random number between 1 and 100
2. Prompts the user to guess the number
3. Gives hints (too high / too low) for each guess
4. Counts the number of attempts
5. Prints a congratulatory message when the user guesses correctly

# Prerequisites

1. Go installed on your machine (version 1.18+ recommended)

2. Terminal / Command Prompt to run the program

# How to Run

1. Clone the repository:

git clone https://github.com/LogicVortex123/Number-Guessing-Game-using-Golang


2. Navigate into the project folder:

cd Golang-Number-Guessing-Game


3. Run the game using Go:

go run main.go


4. Follow the on-screen instructions to guess the number!

Example Output
🎯 Welcome to the Number Guessing Game!
I have selected a number between 1 and 100.
Enter your guess: 50
Too low! Try again.
Enter your guess: 75
Too high! Try again.
Enter your guess: 62
✅ Correct! You guessed the number in 3 attempts.

# Concepts Learned

1. Random number generation (math/rand)
2. Loops (for)
3. Conditional statements (if-else)
4. User input (fmt.Scan)
5. Basic CLI interaction

# Author

Anushka Dudhe | Keploy GSoC 2026 Participant
anushkadudhe9@gmail.com
