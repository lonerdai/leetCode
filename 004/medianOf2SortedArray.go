func findKth(nums1 []int,int m,nums2[]int,int n,int k)float64{
    if m>n {
        return findKth(nums2,n,num1,m,k)
    }
    if m==0 {
        return nums[k-1]
    }
    if k==1 {
        return min(nums1[0],nums2[0])
    }
    pa,pb := min(k/2,m),k-pa
    if *(nums1+pa-1)<*(nums2+pb-1) {
        return findKth(nums1+pa,m-pa,b,n,k-pa)
    }else if *(nums1+pa-1)>*(nums2+pb-1) {
        return findKth(nums1,m,nums2+pb,n-pb,k-pb)
    }else {
        return *(nums1+pa-1)
    }
    
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    
}