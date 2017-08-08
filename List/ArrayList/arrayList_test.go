package ArrayList

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testData = map[string]interface{}{
		"字符串": []string{"1", "2", "3", "4", "5"},
	}
)

func TestArrayList_Append(t *testing.T) {
	for title, data := range testData {
		t.Run(title, func(t *testing.T) {
			a := New()

			slicedData := make([]interface{}, 0, 10)

			switch title {
			case "字符串":
				for _, d := range data.([]string) {
					slicedData = append(slicedData, d)
				}
			}
			a.Append(slicedData...)

			iter := a.Iterator()

			for i:=0;i<len(slicedData);i++{

			}
			element, ok := iter.Head()
			assert.Equal(t, ok, true, "第一个元素应该存在")
			assert.Equal(t, element, "1")
		})

	}

}
