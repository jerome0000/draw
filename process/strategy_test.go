package process

import (
	"fmt"
	"github.com/jerome0000/draw/conf"
	"sort"
	"testing"
	"time"
)

func TestStrategyHandler(t *testing.T) {

}

func Test_checkStrategyStatus(t *testing.T) {
	b := checkStrategyStatus(conf.Strategy{
		ID:           0,
		Rules:        []int64{1, 2, 3},
		StartTime:    time.Now().AddDate(0, 0, -1),
		EndTime:      time.Now().AddDate(0, 0, 1),
		StartTimeDay: "11:00:00",
		EndTimeDay:   "24:00:00",
		Weights:      1,
	}, time.Now(), map[string]interface{}{})
	fmt.Println(b)
}

func TestSort(t *testing.T) {
	arr := make([]conf.Strategy, 0)
	arr = append(arr, conf.Strategy{
		ID:      0,
		Weights: 0,
	})
	arr = append(arr, conf.Strategy{
		ID:      1,
		Weights: 1,
	})

	sort.Sort(WeightsSort(arr))

	fmt.Println(arr)
}