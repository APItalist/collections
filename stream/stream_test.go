package stream_test

import (
    "fmt"

    "github.com/apitalist/collections/stream"
)

func ExampleStream_Filter() {
    r, err := stream.
        Of(1, 2, 3, 4, 5, 6).
        Filter(
            func(e int) bool {
                return e%2 == 0
            },
        ).
        ToSlice()
    if err != nil {
        panic(err)
    }
    fmt.Println(r)

    // Output: [2 4 6]
}

func ExampleStream_AllMatch() {
    allEven, err := stream.
        Of(2, 4, 6).
        AllMatch(
            func(e int) bool {
                return e%2 == 0
            },
        )
    if err != nil {
        panic(err)
    }
    if allEven {
        fmt.Println("All numbers are even.")
    } else {
        fmt.Printf("Not all numbers are even.")
    }

    // Output: All numbers are even.
}

func ExampleStream_AnyMatch() {
    allEven, err := stream.
        Of(1, 2, 3, 4, 5, 6).
        AnyMatch(
            func(e int) bool {
                return e%2 == 0
            },
        )
    if err != nil {
        panic(err)
    }
    if allEven {
        fmt.Println("There are even numbers in the stream.")
    } else {
        fmt.Printf("There are no even numbers in the stream.")
    }

    // Output: There are even numbers in the stream.
}

func ExampleMap() {
    s, err := stream.Map(
        stream.
            Of(1, 2, 3, 4, 5, 6, 7, 8),
        func(input int) (string, error) {
            return fmt.Sprintf("%d", input), nil
        },
    ).ToSlice()
    if err != nil {
        panic(err)
    }
    fmt.Println(s)
    // Output: [1 2 3 4 5 6 7 8]
}
