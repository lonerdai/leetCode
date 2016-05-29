func abs(x int)int{
    if x>0{
        return x
    }else{
        return -x
    }
}
func reverse(x int) int {
    res := 0
    for x!=0{
        if(abs(res)>214748364){
            return 0
        }
        res = res*10+x%10
        x/=10
    }
    return res
}