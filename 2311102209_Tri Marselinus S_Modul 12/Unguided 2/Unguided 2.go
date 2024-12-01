package main

import "fmt"

func insertionSort(nums []int) {
    for i := 1; i < len(nums); i++ {
        key := nums[i]
        j := i - 1
        for j >= 0 && nums[j] > key {
            nums[j+1] = nums[j]
            j = j - 1
        }
        nums[j+1] = key
    }
}

func median(nums []int) float64 {
    insertionSort(nums)
    n := len(nums)
    if n%2 == 0 {
        return float64(nums[n/2-1]+nums[n/2]) / 2
    }
    return float64(nums[n/2])
}

func main() {
    var nums []int
    var num int
    for {
        fmt.Scan(&num)
        if num == 0 {
            if len(nums) > 0 {
                fmt.Println("Median:", median(nums))
            }
        } else if num == -5313 {
            break 
        } else {
            nums = append(nums, num) 
        }
    }
}