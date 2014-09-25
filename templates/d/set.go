type Set map[T]bool

func New() Set {
    return make(Set)
}

func (set Set) Add(key T) {
    set[key] = true
}

func (set Set) Contains(key T) bool {
    return set[key]
}

func (set Set) Remove(key T) {
    delete(set, key)
}
