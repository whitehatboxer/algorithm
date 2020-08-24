package main

import ("fmt")

func minimumTotal(triangle [][]int) int {
    return 0;
}


func main() {
    a := [][]int {
        {2},
        {3, 4},
        {6, 5, 7},
        {4, 1, 8, 3}
    }
    fmt.Println("res is: ", minimumTotal(a));
}
