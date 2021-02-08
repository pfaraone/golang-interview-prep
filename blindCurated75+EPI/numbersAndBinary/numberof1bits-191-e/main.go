func hammingWeight(num uint32) int {
	res := 0
	bitmask := uint32(1)
	for i := 0; i < 32; i++ {
		if num&bitmask != uint32(0) {
			res += 1
		}
		bitmask <<= 1
		// if (num & uint32(1) != uint32(0)){
		//     res += 1
		// }

	}
	return res
}

func hammingWeight(num uint32) int {
	sum := 0
	for num != uint32(0) {
		sum++
		num &= (num - 1)
	}
	return sum
	//     res := 0
	//     bitmask := uint32(1)
	//     for i:=0;i<32;i++{
	//         if (num & bitmask != uint32(0)){
	//             res += 1
	//         }
	//         bitmask <<= 1
	//         // if (num & uint32(1) != uint32(0)){
	//         //     res += 1
	//         // }

	//     }
	//     return res
}