# READ ME

## Github link: https://github.com/jssuzuki1/command_line_application

## How to execute the command:
The main executable is assignment3.exe

To run it, change your directory to this folder, then change the parameters (denoted in brackets of the following code:

./assignment3.exe -op [min, max, or mean] -col [column number] housesInput.csv

  - ./assignment3.exe is the build file
	- -op [min, max, or mean] denotes the operation, which is followed up by "mean" in this example. Here, the operation can be mean, min, and max.
	- -col [column number] is the column number of your choosing
  - houseInput.csv is the file we are performing operations on.

For example, here is a valid command: 
`./assignment3.exe -op mean -col 2 housesInput.csv `

## What this file actually does
It should perform the computation and return it on the command prompt. It should also output a text file called "output.txt" that has the result generated.

## Testing
main_test.go loops the valid command example 100 times and checks to see if the value is correct against the true value. It also checks the computation time of running the function 100 times.

On my computer, it runs the loop in 0.882s
