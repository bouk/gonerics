type Function func(T) U
type Function2 func(U, T) U

func Map(f Function, input []T) []U {
    result := make([]U, len(input))

    for i, element := range input {
        result[i] = f(element)
    }

    return result
}

func Reduce(f Function2, initial U, input []T) U {
    for _, v := range input {
        initial = f(initial, v)
    }
    return initial
}
