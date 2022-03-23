package sliceMap

import (
	"math/rand"
	"testing"
	"time"
)

func prepareData() [500]int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	var data [500]int
	for i := range data {
		data[i] = r1.Intn(100)
	}

	return data
}

func BenchmarkSliceOperations_Array(b *testing.B) {
	data := prepareData()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		SliceOperations(OPTYE_ARRAY, &data)
	}
}

func BenchmarkSliceOperations_SliceAppend(b *testing.B) {
	data := prepareData()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		SliceOperations(OPTYE_SLICE_APPEND, &data)
	}
}

func BenchmarkSliceOperations_SliceAppendPredefinedLimit(b *testing.B) {
	data := prepareData()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		SliceOperations(OPTYE_SLICE_APPEND_PREDEFINED, &data)
	}
}

func BenchmarkSliceOperations_SliceFixedSize(b *testing.B) {
	data := prepareData()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		SliceOperations(OPTYE_SLICE_FIXED_SIZE, &data)
	}
}

func BenchmarkSliceOperations_Map(b *testing.B) {
	data := prepareData()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		SliceOperations(OPTYE_MAP, &data)
	}
}

func BenchmarkSliceOperations_MapPedefined(b *testing.B) {
	data := prepareData()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		SliceOperations(OPTYE_MAP_PREDEFINED, &data)
	}
}
