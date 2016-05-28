func convert(s string, numRows int) string {
    if numRows == 1 {
        return s
    }
    twoD := make([][]rune, numRows)
    for i := 0; i < numRows; i++ {
        twoD[i] = make([]rune, 0)
    }
    down ,at:= true,0
    for i := 0; i < len(s); i++ {
        twoD[at] = append(twoD[at],rune(s[i]))
        if down {
            if at == numRows-1 {
                down = false
                at = numRows-2
            }else {
                at++
            }
        }else {
            if at==0 {
                down = true
                at = 1
            }else {
                at--
            }
        }
    }
    res := ""
    for i := 0; i < numRows; i++ {
        for j := 0; j < len(twoD[i]); j++ {
            res+=string(twoD[i][j])
        }
    }
    return res
}