package randomutils

import (
	"math/rand"
	"time"
	"math"
)

type awards [][]int
type awardlist [][][]int
type awardlists [][][][]int


// 根据范围（闭区间）随机int
func RandomInt(min, max int) int{
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}

func GetAwardByWeight(awards awards) []int {
	var weights = 0
	for _, award := range awards {
		weights += award[len(award)-1]
	}
	rand.Seed(time.Now().Unix())
	randomNum := rand.Intn(weights)
	sumNum := 0
	for _, award := range awards {
		sumNum += award[len(award)-1]
		if sumNum > randomNum {
			return award
		}
	}
	return nil
}

//输入格式{{2},{10},{100}} 或 {{2},{5,10},{100}}
func GetAwardByPercentage(awards [][]int, fullPercent int) []int {
	var result []int
	rand.Seed(time.Now().Unix())
	var randomNum = rand.Intn(fullPercent)
	var rateSum = 0
	for i := 0; i < len(awards); i++ {
		rateSum = awards[2][0]
		if rateSum >= randomNum {
			if len(awards[1]) > 1 {
				// 取随机数量
				var maxi = math.Max(float64(awards[1][0]), float64(awards[1][1]))
				var mini = math.Min(float64(awards[1][0]), float64(awards[1][1]))
				var num = rand.Intn(int(maxi)-int(mini)) + int(mini)
				slice1 := make([]int, 2)
				slice1[0] = awards[0][0]
				slice1[1] = num
				return slice1
			} else {
				slice1 := make([]int, 2)
				slice1[0] = awards[0][0]
				slice1[1] = awards[1][0]
				return slice1
			}
		}
	}
	return result
}

//输入格式 {{{6},{10},{30}},{{7},{1,5},{20}}}
func GetAwardByPercentageAlternative(awards awardlist, fullPercent int) []int {
	var result []int
	rand.Seed(time.Now().Unix())
	var randomNum = rand.Intn(fullPercent)
	var rateSum = 0
	for i := 0; i < len(awards); i++ {
		rateSum += awards[i][2][0]
		if rateSum >= randomNum {
			if len(awards[i][1]) > 1 {
				// 取随机数量
				var maxi = math.Max(float64(awards[i][1][0]), float64(awards[i][1][1]))
				var mini = math.Min(float64(awards[i][1][0]), float64(awards[i][1][1]))
				var num = rand.Intn(int(maxi)-int(mini)) + int(mini)
				slice1 := make([]int, 2)
				slice1[0] = awards[i][0][0]
				slice1[1] = num
				return slice1
			} else {
				slice1 := make([]int, 2)
				slice1[0] = awards[i][0][0]
				slice1[1] = awards[i][1][0]
				return slice1
			}
		}
	}
	return result
}

func GetAwardsByPercentage(awards awardlists, fullPercent int) [][]int {
	//百分比奖励获取
	//    例子：get_awards_by_percentage([[[2, [6, 10], 100]], [[5, 5, 100]], [[6, 10, 30], [7, [1,10], 20]]], 100)
	//    100为百分比的基数  100则最小单位为1% 1000为0.1% 10000为0.01% 以此类推
	//    该例子表示产出 2-3种道具
	//    [[2, [6, 10], 100]] 表示第一个位置上100%出产 6-10个道具id为2的道具
	//    [[5, 5, 100]] 表示第二个位置上100%出产 5个道具id为5的道具
	//    如上[[6, 10, 30], [7, [1,10], 20]] 表示第三个出产物品的位置有三种可能性
	//        1:什么都不出  2.30%的几率出10个道具id为6的道具 3.20%的几率出 1-10个道具id为7的道具
	var awardsItem = [][]int{}
	var item []int
	for _, value := range awards {
		if len(value) > 1 {
			item = GetAwardByPercentageAlternative(value, fullPercent)
			if len(item) > 0 {
				awardsItem = append(awardsItem, item)
			}
		} else {
			for _, v := range value {
				item = GetAwardByPercentage(v, fullPercent)
				if len(item) > 0 {
					awardsItem = append(awardsItem, item)
				}
			}
		}
	}
	return awardsItem
}

// 获取count个互不相等的随机数
func GetDifferentRandomNum(minNum, maxNum, count int) []int {
	var nums []int
	rand.Seed(time.Now().Unix())
	if maxNum < minNum || (maxNum-minNum) < count {
		return nil
	}
	for len(nums) < count {
		num := rand.Intn(maxNum-minNum) + minNum
		exist := false
		for _, value := range nums {
			if value == num {
				exist = true
				break
			}
		}
		if !exist {
			nums = append(nums, num)
		}
	}
	return nums
}
