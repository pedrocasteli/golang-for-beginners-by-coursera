## Commands and events

The first command was `go mod init silverkarma/faremenu`. It creates a module file, specifying the path of the module and the version of **Go** we are using. This file will be used to build an run the project in the future.

The `faremenu.go` file was created normally by us.

The nest thing we did was run the command `go run .`, and the strings were showed in the console.

The `&` used to read the value typed by the user, is a reference to the memory address of tha variable

In **Go**, we don't have to put everything in one single file. That's why we created the `data.go` file. We especify the same package as `faremenu.go` so they can see each other. We created a custom data type called `City`: A struct that holds some values about cities. Then we created a variable of the type `array` of `City` elements called `cities`. We can then access the `cities` variable in the `faremenu.go` file.

We also created the custom type `CabinClass`, and an `array` of `CabinClass` called `cabinClasses`.

In **Practice-Task-1**, we had to fix the code. In `greeter.go`, I changed the **package** to "main". The main function's name was changed to "main" with lowercase "m". And in the `fmt.Scanln` function I added "&" before the variable name.

In the function `getCityFromCode`, if there is no error, the `error` will be `nil`.
The `:=` characters means you are declaring the variable and assigning a value to it at the same time. That way we don't need the `var` keyword. We also create the `validOriginEnterned` variable, that starts as `false`. As long as it stays `false`, we continue inside the `FOR` loop. `originCity` and `originError` receive the returning values of the `getCityFromCode` function. If `originError` in `nil`, then we know that the function found a corresponding origin code, and `validOriginEnterned` becomes `true`.

We applied the same validations for the destination and for the cabin class.
