# Guess the Number Game

## **Description**

The **Guess the Number** game is a simple command-line program where players try to guess a randomly chosen number between 1 and 100. The game provides hints ("Too low!" or "Too high!") to guide the player toward the correct number. The program tracks the total time taken to find the correct number, adding an exciting time-based challenge.

## **How the Game Works**

1. The program randomly selects a number between 1 and 100.
2. The player is prompted to enter their guesses.
3. The program provides feedback after each guess:
   - **"Too low!"** if the guess is lower than the target number.
   - **"Too high!"** if the guess is higher than the target number.
4. When the player guesses correctly:
   - The program announces the win.
   - It displays the total time taken to find the correct number.

---

## **How to Run the Game**

### **Requirements**
- Go programming language installed on your system.
- Basic knowledge of how to run Go programs.

### **Steps to Run**
1. Open a terminal and clone the repository.
```bash
git clone https://github.com/OumarLAM/guess_the_number.git
```
2. Navigate to the directory containing the project.
```bash
cd guess_the_number
```
3. Run the game using the command:  
```bash
go run main.go
```
4. Follow the on-screen instructions to play the game.

## **Development Steps**

1. ### **Generate a Random Number**
- Use Go's math/rand package to generate a random number.
- Seed the random number generator using the current timestamp for true randomness:
```go
rand.New(rand.NewSource(time.Now().UnixNano()))
```
- Generate a number between 1 and 100:
```go
correctNumber := rand.Intn(101)
```

2. ### **Prompt for User Input**
- Use fmt.Scanln() to read user input.
- Trim spaces and convert the input string to an integer: 
```go
input = strings.TrimSpace(input)
guess, err := strconv.Atoi(input)
```

3. ### **Validate the Input**
- If the input cannot be converted to an integer, display an error message:
```go
if err != nil {
    fmt.Println("Invalid input. Please enter a number between 1 and 100.")
    continue
}
```

4. ### **Check the Guess**
- Compare the guess to the random number:
    - Print "Too low!" if the guess is smaller.
    - Print "Too high!" if the guess is larger.
    - End the game if the guess is correct.

5. ### **Track the Time**
- Record the start time at the beginning of the game: 
```go
startTime := time.Now()
```
- Calculate the elapsed time when the player guesses correctly:
```go
elapsedTime := time.Since(startTime)
```

6. ### **Display Results**
- Congratulate the player and display the correct number:
```go
fmt.Printf("🎉 Congratulations! You guessed the correct number: %d\n", correctNumber)
```
- Show the total time taken:
```go
fmt.Printf("It took you %.2f seconds.\n", elapsedTime.Seconds())
```

## **Author**
This project was developed by **Oumar LAM** as an introductory programming activity for high school students from Mariama BA. The game is designed to be fun, intuitive, and challenging while reinforcing programming fundamentals.