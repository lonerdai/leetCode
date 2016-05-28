func lengthOfLongestSubstring(s string) int {
    m ,begin,res:= make(map[byte]int),0,0
    if s == "" {
        return 0
    }
    m[s[0]] = 0 
    for i := 1; i < len(s); i++ {
        _,exist := m[s[i]]
        if exist {
            if(i-begin > res){
                res = i-begin
            }
            for ; begin < i;begin++ {
                if s[begin] == s[i] {
                    delete(m,s[i])
                    break
                }else{
                    delete(m,s[begin])
                }
            }
        }else {
            m[s[i]]=0
        }
    }
    return res
}