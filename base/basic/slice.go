package basic

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func TestSilce() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:] =", arr[:])

	s1 := arr[2:] //视图索引
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)
	arr[2] = 2
	fmt.Println("Extending slice")
	s1 = arr[2:6]
	s2 := s1[3:5] //slice可以向后拓展，不可以向前拓展
	// s[i]不可以超越len(s)，向后拓展不能超越cap(s)
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println("s3, s4, s5 =", s3, s4, s5)
	fmt.Println("arr =", arr)
}
