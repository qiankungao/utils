package rand

import (
	"math/rand"
	"time"
)

type Draw struct {
	Id     int32
	Weight int32
	Index  int //这是个SB字段
}

//加权随机
func RandomDraw(draw map[int32]*Draw) *Draw {
	//权重累加求和
	var weightSum int32
	for _, v := range draw {
		weightSum += v.Weight
	}

	//生成一个权重随机数，介于0-weightSum之间
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Int31n(weightSum)

	total := int32(0)
	tar := int32(0)
	for index, p := range draw {
		total += p.Weight
		if total > randomNum {
			tar = index
			break
		}
	}
	return draw[tar]
}

func RandInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}

//在数组中随机N个不同的数
func RandElement(tar []int32, num int) []int32 {
	var result []int32
	l := len(tar)
	if num > l {
		return nil
	}
	for i := 0; i < num; i++ {
		index := RandInt(0, l)
		result = append(result, tar[index])
		tar[index], tar[l-1] = tar[l-1], tar[index]
		l -= 1
	}
	return result
}
