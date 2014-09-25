type FilterFunction func(T) bool
type Function func(T) T
type Function2 func(T, T) T

func combine(a Function, b Function) Function {
    return func(arg T) T {
        return a(b(arg))
    }
}

func Compose(first Function, f ...Function) Function {
    for _, v := range f {
        first = combine(first, v)
    }
    return first
}

func Map(f Function, input []T) []T {
    result := make([]T, len(input))

    for i, element := range input {
        result[i] = f(element)
    }

    return result
}

func Reduce(f Function2, input []T) (accum T) {
    if len(input) == 0 {
        return accum
    }

    accum = input[0]
    for _, v := range input[1:] {
        accum = f(accum, v)
    }

    return
}

func Filter(f FilterFunction, input []T) (result []T) {
    for _, v := range input {
        if f(v) {
            result = append(result, v)
        }
    }
    return
}

func TakeWhile(f FilterFunction, input []T) (result []T) {
    for _, v := range input {
        if !f(v) {
            break
        }
        result = append(result, v)
    }
    return
}

func DropWhile(f FilterFunction, input []T) (result []T) {
    take := false
    for _, v := range input {
        if take = take || f(v); take {
            result = append(result, v)
        }
    }
    return
}

func ZipWith(f Function2, inputa []T, inputb []T) (result []T) {
    upto := len(inputa)
    if len(inputb) < upto {
        upto = len(inputb)
    }

    result = make([]T, upto)

    for i := 0; i < upto; i++ {
        result[i] = f(inputa[i], inputb[i])
    }

    return
}
