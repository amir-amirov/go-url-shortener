package shorten

const alphabet = "abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ123456789"

var alphabetLen = uint32(len(alphabet))

func Shorten(id uint32) string {
	var nums []uint32 = []uint32{}

	for id > 0 {
		n := id % alphabetLen
		nums = append(nums, n)
		id = id / alphabetLen
	}

	reverse(nums)

	var chars []byte = []byte{}

	for _, n := range nums {
		if n >= alphabetLen {
			panic("n is out of bounds")
		}
		chars = append(chars, alphabet[n])
	}

	return string(chars)

}

func reverse(nums []uint32) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
