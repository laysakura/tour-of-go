package main

import (
    "code.google.com/p/go-tour/tree"
    "fmt"
    "reflect"
)

func walk(root bool, t *tree.Tree, ch chan int) {
    if t.Left  != nil { walk(false, t.Left, ch) }
	ch <- t.Value
    if t.Right != nil { walk(false, t.Right, ch) }

    if root { close(ch) }
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
    walk(true, t, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
    ch1 := make(chan int)
    vals1 := []int {}
    ch2 := make(chan int)
    vals2 := []int {}

    go Walk(t1, ch1)
    go Walk(t2, ch2)

    for v := range ch1 {
        vals1 = append(vals1, v)
    }
    for v := range ch2 {
        vals2 = append(vals2, v)
    }

    return reflect.DeepEqual(vals1, vals2)
}

func main() {
    ch := make(chan int)
    go Walk(tree.New(1), ch)

    for v := range ch {
        fmt.Println(v)
    }


    fmt.Println(Same(tree.New(1), tree.New(1)))
    fmt.Println(Same(tree.New(1), tree.New(2)))
}
