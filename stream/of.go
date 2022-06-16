package stream

import "github.com/apitalist/collections"

func Of[E any](elements ...E) collections.Stream[E] {
    input := make(chan E)
    complete := make(chan struct{})
    errorInput := make(chan error)
    s := &stream[E]{
        input:      input,
        complete:   complete,
        errorInput: errorInput,
    }
    go func() {
        defer func() {
            close(input)
            close(errorInput)
        }()
        for _, e := range elements {
            select {
            case input <- e:
            case <-complete:
                return
            }
        }
    }()
    return s
}
