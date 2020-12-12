package main


func TestMerge(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	println(len(nums1), cap(nums1))

	// nums1 := []int{2, 0}
	// nums2 := []int{1}

	// nums1 := []int{1, 2, 3, 4, 5, 8}
	// nums2 := []int{2, 5, 6}

	merge(nums1, len(nums1), nums2, len(nums2))
	// fmt.Printf("TestMerge%#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&nums1)))
	fmt.Println(nums1)
}
