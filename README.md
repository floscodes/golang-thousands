thousands is a small package for Gophers to add thousands separators to numbers.
Before use, a number has to be converted to a string. Then you can use the separate-function
to add thousand separators to your number. The separate-function takes two arguments: the number itself and the language-code of the language you want to use. It will return a string containing the number with thousands separators in the language you set with the second argument. For now the package supports English and German. More languages will be added in the future.


func Separate(n string, lang string) string


Examples:

n := 5000.61

N:=strconv.Itoa(n)

fmt.Println("The number is: " + thousands.Separate(N, "en"))

// The number is: 5,000.61

fmt.Println("Die Nummer ist: " + thousands.Separate(N, "de") + "(German)")

//Die Nummer ist: 5.000,61 (German)



