## Commands and events

The first command was `go mod init silverkarma/faremenu`. It creates a module file, specifying the path of the module and the version of **Go** we are using. This file will be used to build an run the project in the future.

The `faremenu.go` file was created normally by us.

The nest thing we did was run the command `go run .` inside the `project-01/` folder, and the strings were showed in the console.

The `&` used to read the value typed by the user, is a reference to the memory address of tha variable

In **Go**, we don't have to put everything in one single file. That's why we created the `data.go` file. We especify the same package as `faremenu.go` so they can see each other. We created a custom data type called `City`: A struct that holds some values about cities. Then we created a variable of the type `array` of `City` elements called `cities`. We can then access the `cities` variable in the `faremenu.go` file.

We also created the custom type `CabinClass`, and an `array` of `CabinClass` called `cabinClasses`.

In **Practice-Task-1**, we had to fix the code. In `greeter.go`, I changed the **package** to "main". The main function's name was changed to "main" with lowercase "m". And in the `fmt.Scanln` function I added "&" before the variable name.

In the function `getCityFromCode`, if there is no error, the `error` will be `nil`.
The `:=` characters means you are declaring the variable and assigning a value to it at the same time. That way we don't need the `var` keyword. We also create the `validOriginEnterned` variable, that starts as `false`. As long as it stays `false`, we continue inside the `FOR` loop. `originCity` and `originError` receive the returning values of the `getCityFromCode` function. If `originError` in `nil`, then we know that the function found a corresponding origin code, and `validOriginEnterned` becomes `true`.

We applied the same validations for the destination and for the cabin class.

To show the distance, we use `%.1f` to make it have just one decimal point. However, it also rounds the number. Just using `%f` gives the result **14327.960709**, while using `%.1f` gives the result **14328.0**.

To calculate the fare, I'm going to use the **github.com/shopspring/decimal** library.
For that, I run the following command in the terminal:

```
go get github.com/shopspring/decimal
```

This changes the `go.mod` file, by adding the `require` statement. It also creates the `go.sum` file.
We then import the package into our files.

At the IDE's suggestion, I executed the following command in the terminal:

```
go mod tidy
```

To create an `.exe` file of the application (a production version), I ran:

```
go build .
```

I imported the `strings` package to use in the validations of city and cabin class. I used the `strings.toUpper()` function so that the user doesn't always have to type in uppercase to inform his choice.
