package mat
 
import "testing"
import "fmt"
import "github.com/stretchr/testify/assert"
import "github.com/stretchr/testify/require"

func TestCekGanjilGenap(t *testing.T){
	result := CekGanjilGenap(1,2,3)
	if result != "ganjil, genap, ganjil, "{
		t.Fail()
	}
	fmt.Println("running unit test")

}

func TestAssert(t *testing.T){
	assert := assert.New(t)
	result := CekGanjilGenap(1,2,3)
	assert.Equal("ganjil, genap, ganjil, ", result, "they should be equal")
}
func TestRequire(t *testing.T){
	require := require.New(t)
	result := CekGanjilGenap(1,2,3)
	require.Equal("ganjil, genap, ganjil, ", result, "they should be equal")
}
func TestSubTest(t *testing.T){
	assert := assert.New(t)
	t.Run("01", func(t *testing.T){
		result := CekGanjilGenap(1,2,3)
		assert.Equal("ganjil, genap, ganjil, ", result, "they should be equal")
	}) 
	t.Run("02", func(t *testing.T){
		result := CekGanjilGenap(1)
		assert.Equal("ganjil, ", result, "they should be equal")
	})
	t.Run("03", func(t *testing.T){
		result := CekGanjilGenap(2)
		assert.Equal("genap, ", result, "they should be equal")
	})
}
func TestTableTest(t *testing.T){
	tests := []struct {
		name string
		expected string
		param []int
	}{
		{
			name: "01",
			expected : "ganjil, genap, ganjil, ",
			param: []int{1,2,3},
		},
		{
			name: "02",
			expected : "ganjil, ",
			param: []int{1},
		},
		{
			name: "01",
			expected : "genap, ",
			param: []int{2},
		},
	}
	assert := assert.New(t)
	for _, test := range tests{
		t.Run(test.name, func(t *testing.T){
			result := CekGanjilGenap(test.param...)
			assert.Equal(test.expected, result, "they should be equal")
		})
	}
}

func BenchmarkCekGanjilGenap(b *testing.B){
	for i :=0;i<b.N ;i++{
		CekGanjilGenap(1,2,3)
	}
}

func BenchmarkTable(b *testing.B){
	tests := []struct{
		name string
		param []int
	}{
		{
			name : "01",
			param : []int{1,2,3} ,
		},
		{
			name : "02",
			param : []int{1}  ,
		},
		{
			name : "03",
			param : []int{2}  ,
		},
	}
	for _,test := range tests{
		b.Run(test.name,func(b *testing.B){
			for i :=0;i<b.N;i++{
				CekGanjilGenap(test.param...)
			}
		})
	}
}