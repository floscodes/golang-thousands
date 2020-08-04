thousands is a small package for Gophers to add thousands seperators to numbers.
Before use, a number has to be converted to a string. Then you can use the Seperate-function
to add thousand seperators to your number. The Seperate-function takes two arguments: the number itself and the language-code of the language you want to use. It will return a string containing the number with thousands seperators in the language you set with the second argument. For now the package supports English and German. More languages will be added in the future.

Examples:

n := 5000.61

N:=strconv.Itoa(n)

fmt.Println("The number is: " + thousands.Seperate(N, "en"))

// The number is: 5,000.61

fmt.Println("Die Nummer ist: " + thousands.Seperate(N, "de") + "(German)")

//Die Nummer ist: 5.000,61 (German)



