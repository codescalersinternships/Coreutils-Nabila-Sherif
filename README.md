# Unix-Based Commands Implemented in Go
This project provides an implementation of several Unix-like Coreutils commands using the Go programming language
## Implemented commands:
### 1. head
reads a file where the file path is given as a command line argument 
the default if to print the 1st 10 lines but if -n flag is present then print the first n lines 
### 2. tail
same as head but the last 10 or n lines
### 3. cat
reads a file and prints its content
### 4. echo
given a string passed as command line arguments , the string is printed
If the -n flag is present, it add a newline at the end.
### 5. true
do nothing, successfuly
### 6. false
do nothing, unsuccessfully
### 7. env
display the current environment
### 8. wc (word count)
the func counts the number of lines, words and characters and prints the count depending on whether flags -l, -w, and -c are present respectively 
### 9. tree
reading a directory and depending on the depth passed(1 or 2) print the up to the depth indicated
