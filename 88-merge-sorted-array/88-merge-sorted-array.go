func merge(nums1 []int, m int, nums2 []int, n int)  {
    res := append(nums1[:m], nums2...)
	sort.Ints(res)
}