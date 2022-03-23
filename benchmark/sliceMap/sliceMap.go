package sliceMap

const OPTYE_ARRAY = 1
const OPTYE_SLICE_APPEND = 2
const OPTYE_SLICE_APPEND_PREDEFINED = 3
const OPTYE_SLICE_FIXED_SIZE = 4
const OPTYE_MAP = 5
const OPTYE_MAP_PREDEFINED = 6

func SliceOperations(opType int, dataSurce *[500]int) int {
	var sum int

	switch opType {
	case OPTYE_ARRAY:
		var localArray [len(dataSurce)]int

		for i, val := range *dataSurce {
			localArray[i] = val
		}

		for _, val := range localArray {
			sum += val
		}

	case OPTYE_SLICE_APPEND:
		localSlice := make([]int, 0)
		for _, val := range *dataSurce {
			localSlice = append(localSlice, val)
		}

		for _, val := range localSlice {
			sum += val
		}

	case OPTYE_SLICE_APPEND_PREDEFINED:
		localSlice := make([]int, 0, len(dataSurce))
		for _, val := range *dataSurce {
			localSlice = append(localSlice, val)
		}

		for _, val := range localSlice {
			sum += val
		}
	case OPTYE_SLICE_FIXED_SIZE:
		localSlice := make([]int, len(dataSurce))

		for i, val := range *dataSurce {
			localSlice[i] = val
		}

		for _, val := range localSlice {
			sum += val
		}

	case OPTYE_MAP:
		localmap := map[int]int{}
		for i, val := range *dataSurce {
			localmap[i] = val
		}

		for _, val := range localmap {
			sum += val
		}

	case OPTYE_MAP_PREDEFINED:
		localmap := make(map[int]int, len(dataSurce))
		for i, val := range *dataSurce {
			localmap[i] = val
		}

		for _, val := range localmap {
			sum += val
		}
	}

	return sum
}
