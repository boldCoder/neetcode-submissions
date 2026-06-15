func isAnagram(s string, t string) bool {
if len(s) != len(t) {
    return false
}

hashComp1 := make(map[string]int)

for _, v := range s {
    hashComp1[string(v)] += 1
}

hashComp2 := make(map[string]int)
for _, v := range t {
    hashComp2[string(v)] += 1
}

for _, v := range s {
    if hashComp1[string(v)] != hashComp2[string(v)] {
        return false
    }
}

return true

}
