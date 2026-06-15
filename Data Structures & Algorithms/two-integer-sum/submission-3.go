func twoSum(nums []int, target int) []int {

seen := make(map[int]int)

for i, num := range nums {
    delta := target - num
    if idx, found := seen[delta]; found {
        return []int{idx,i}
    }

    seen[num] = i
}

return nil
}
