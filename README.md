# cbreakable
Simple library to create text based interfaces

Heavy-duty golang libraries for this purpose exist, but the goal of this one is to give light functionality for purposes where using something heavier would be overkill.

# Contents
(Usage)[#Usage]
  (Char datatype)[###Char]
  (Writer datatype)[###Writer]
  (Entering Cbreak mode and the Reader datatype)[###Reader]
(Licensing)[#License]

# Usage
This library is simple enough to be explained in the README!

### Char
This is an interface that represents a cbreakable Char. It could be a "literal char", like 'a' or '-' or '!', or it could be an escape char, like "\x1b[A" (the up arrow key). There are only 3 methods on cbrk.Char:

##### Char.String
This is the string representation of the char. Literal chars return a string version of their literal value, and escape chars return their escape sequence

##### Char.Equals
To test for equality between either another cbrk.Char, or any other value that implements interface{} (which is all of them). 

##### Char.Repr
This is another string representation of a char, which you will probably use less. Literal chars get printed out as their ascii value, and escape chars get printed as a series of space seperated ascii values.

##### CharString
This is a function that takes in a []cbrk.Char and returns a string.

### Writer
This relatively simple type allows you to print out text and easily delete it. 2 instances are available to you: cbrk.Stdout and cbrk.Stderr, which print to os.Stdout and os.Stderr respectively. 

##### Print
Print out a given value. Really thats it.

```go
cbrk.Stdout.Print("hello, world!\n")
```

##### Println
Print out a given value with a newline ('\n' char) at the end.

```go
cbrk.Stdout.Println("hello, world")
```

##### Clear
Clear every line printed out by the writer. Make sure that you call Writer.Println, or end your last Writer.Print call with a "\n" before you call this.

```go
cbrk.Stdout.Println("hello\nworld!")
time.Sleep(time.Second)
cbrk.Stdout.Clear()
```

##### Clearln
Clear the last line printed out by the writer. The same warning to end your last message with a newline (or use Writer.Println) applies as in Writer.Clear.

```go
cbrk.Stdout.Print("hello\nworld\n")
time.Sleep(time.Second)
cbrk.Stdout.Clearln()
cbrk.Stdout.Clearln()
```

##### ClearPage
Clear the entire terminal, dont worry about calling Writer.Println before you call this.

```go
cbrk.Stdout.Print("wooo, no trailing newline")
cbrk.Stdout.ClearPage()
```

### Input
A big part of this library is dealing with raw input. Normally you need to hit enter before your code can read text, and any text you do type gets printed out as you type it. Cbreakable provides an easy way to read input 1 "raw" char at a time by putting the terminal in "cbreak" mode (hence the name).

##### Cbreak
This function is used to enable and disable cbreak mode in the terminal. Call `Cbreak(true)` to enable cbreak, and `Cbreak(false)` to disable it. When writing applications in cbreak it is common practice to use the following to make sure that you exit cbreak before your code exits:


```go
cbrk.Cbreak(true)
defer cbrk.Cbreak(false)
```

If your code exits without calling cbreak(false), (like during a panic), simply type "stty sane" into your terminal and hit enter. This will revert it back to normal.

##### Reader
This type is used to read chars from a buffer. You have access to 1 instance: cbrk.Stdin, which reads from os.Stdin. You should call `Cbreak(true)` before using it.

##### Get
This function will block until a char appears in os.Stdin, then return it. Note that chars like `'\x1b'` will continue to block, as Reader.Get blocks until it can construct a valid cbrk.Char.

```go
cbrk.Cbreak(true)
defer cbrk.Cbreak(false)
cbrk.Stdout.Print("Allow this program to continue? (y/n)")
char := cbrk.Stdin.Get()
cbrk.Stdout.Println(char)
if char.Equals("n") {
  return
}
cbrk.Stdout.Println("guess what, it continued!")
```

##### Getln
This function will block until it can read an entire line from os.Stdin, then return a []cbrk.Char representing it.

```go
cbrk.Cbreak(true)
defer cbrk.Cbreak(false)
cbrk.Stdout.Println("What is your password?")
if cbrk.CharString(cbrk.Stdin.Getln()) == "incorrect\n" {
  cbrk.Stdout.Println("thats right!")
}else{
  cbrk.Stdout.Println("thats wrong :(")
}
```

# License
This project is licensed under MIT, see /LICENSE for more information.
