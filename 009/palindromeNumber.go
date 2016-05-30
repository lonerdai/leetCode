func isPalindrome(x int) bool {
    if x<0{
        return false
    }
    div:=1
    for div*10<=x{
        div*=10
    }
    for div>1{
        if x/div != x%10 {
            return false
        }
        x%=div
        div/=100
        x/=10
    }
    return true
}