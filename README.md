thousands is a small package for Gophers to add thousands seperators to numbers.
Before use, a number has to be converted to a string. Then you can use the Seperate-function
to add thousand seperators to your number. The Seperate-function will return a string containing
the number with thousands seperators. The decimal point will be replaced by a comma.

Example:

n := 5000

N:=strconv.Itoa(n)

fmt.Println("The number is: " + thousands.Seperate(N))

// The number is: 5.000



