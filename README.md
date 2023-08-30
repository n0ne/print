# Print

[![GoDoc](https://img.shields.io/badge/api-reference-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/n0ne/print) 

Print is a Go package that makes terminal console output a little prettier  (-:

Getting Started
===============
## Installing

To start using Print, install Go and run `go get`:

```sh
$ go get -u github.com/n0ne/print
```

This will retrieve the library.

## How to use

### Print struct

Create a struct:
```go
type User struct {
  Name  string
  Age   int
  Email string
}
```

then create a user:
```go
user := User{
  Name:  "Alex",
  Age:   25,
  Email: "test@test.com",
}
```

Now we can print our user:
```go
print.Struct(user)
```

The output will be:

<img width="206" alt="Screenshot 2023-08-30 at 17 54 34" src="https://github.com/n0ne/print/assets/783906/e4fd4592-a953-4d6b-8c12-30171b318a53">

Or you can call `Structc()` function, which will add some colors to the output:
```go
print.Structc(user)
```

The output in this case will be:

<img width="203" alt="Screenshot 2023-08-30 at 17 54 50" src="https://github.com/n0ne/print/assets/783906/d177d77f-c383-4c25-9b9f-030c33ce89fb">

### Print slice

Let's create a few slices:
```go
  sliceInt := []int{1, 2, 3, 4, 5, 6}
  sliceString := []string{"Some", "strings", "here"}
  sliceDouble := []float64{1.1, 2.2, 3.3}
```

and print them:
```go
  print.Slice(sliceInt)
  print.Slice(sliceString)
  print.Slice(sliceDouble)
```

The output for slices will be:

<img width="231" alt="Screenshot 2023-08-30 at 17 55 03" src="https://github.com/n0ne/print/assets/783906/17779377-94aa-4e55-9942-47a5fcfb8f8a">

### Print error message

Let's add some color to our error message:
```go
  print.Error("Error: something went wrong")
```

And this you will see in the console:

<img width="205" alt="Screenshot 2023-08-30 at 17 55 10" src="https://github.com/n0ne/print/assets/783906/51aa784b-56da-44a2-81c4-c7ecec5fa713">



    

## License

Print source code is available under the MIT [License](/LICENSE).
