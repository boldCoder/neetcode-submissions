func hasDuplicate(nums []int) bool {
    hash:= make(map[int]int)

    for _, val := range nums {
        hash[val] += 1
        if hash[val] > 1 {
            return true
        }
    }

    return false
}
