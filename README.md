thousands is a small package for Gophers to add thousands separators to numbers.
The Separate-function takes two arguments: the number itself and the language-code of the language you want to use. It will return a string containing the number with thousands separators in the language you set with the second argument. For now the package supports English and German. More languages will be added in the future.

```
func Separate(N interface{}, lang string) string
```

Examples:
```
n := 5000.61

fmt.Println("The number is: " + thousands.Separate(n, "en"))

// The number is: 5,000.61

fmt.Println("Die Nummer ist: " + thousands.Separate(n, "de") + "(German)")

//Die Nummer ist: 5.000,61 (German)

```
