package main

import (
    "fmt"
    "time"
)

type Person struct {
    firstName string
    lastName string
    birthday time.Time
}

func main() {
    birthday, _ := time.Parse("2006-Jan-02", "1994-Aug-10")

    person := Person{
        firstName: "Alex",
        lastName: "Kotov",
        birthday: birthday,
    }

    printPersonByValue(person)
    printPersonByRefSafe(&person)
    printPersonByRefUnsafe(&person)
    printPersonByRefSafe(nil)
    printPersonByRefUnsafe(nil)
}

func printPersonByValue(person Person) {
    fmt.Println("=== by value =====================================")
    fmt.Println("First name:", person.firstName)
    fmt.Println("Last name: ", person.lastName)
    fmt.Println("Birthday:  ", person.birthday)
}

func printPersonByRefSafe(person *Person) {
    fmt.Println("=== by reference (safe) ==========================")

    if person == nil {
        fmt.Println("nil pointer, skip...")
        return
    }

    fmt.Println("First name:", person.firstName)
    fmt.Println("Last name: ", person.lastName)
    fmt.Println("Birthday:  ", person.birthday)
}

func printPersonByRefUnsafe(person *Person) {
    fmt.Println("=== by reference (unsafe) ========================")
    fmt.Println("First name:", person.firstName)
    fmt.Println("Last name: ", person.lastName)
    fmt.Println("Birthday:  ", person.birthday)
}
