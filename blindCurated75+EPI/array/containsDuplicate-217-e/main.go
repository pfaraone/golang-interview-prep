import "fmt"

// "sort"

func myinsertionsort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		valueToSort := arr[i]
		// j := i-1
		for i > 0 && arr[i-1] > valueToSort {
			arr[i], arr[i-1] = arr[i-1], arr[i]
			i--
		}
	}
	return arr
}

func containsDuplicate(nums []int) bool {
	// sort.Ints(nums)
	my_nums := myinsertionsort(nums[:])
	counter := make(map[int]int)
	// fmt.Println(my_nums)
	for _, num := range my_nums {
		_, ok := counter[num]
		if ok {
			counter[num]++
			if counter[num] == 2 {
				return true
			}
		} else {
			counter[num] = 1
		}
		// fmt.Println(counter)
	}

	return false
}

func containsDuplicate2(nums []int) bool {
	// sort.Ints(nums)
	my_nums := myinsertionsort(nums[:])
	counter := make(map[int]int)
	fmt.Println(my_nums)
	for _, num := range my_nums {
		_, ok := counter[num]
		if ok {
			counter[num]++
			if counter[num] == 2 {
				return true
			}
		} else {
			counter[num] = 1
		}
		// fmt.Println(counter)
	}

	return false
}

func containsDuplicate3(nums []int) bool {
    sort.Ints(nums)
    for i:=0; i<=len(nums)-2;i++{
        if nums[i] == nums[i+1]{
            return true
        }
    }
    return false