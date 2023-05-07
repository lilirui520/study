package pkg

import (
	"fmt"
	"math"
)

type Vector struct {
	X float32
	Y float32
}

// 坐标点相加
func (vector1 Vector) add(vector2 Vector) Vector {
	return Vector{vector1.X + vector2.X, vector1.Y + vector2.Y}
}

// 坐标点相减
func (vector1 Vector) sub(vector2 Vector) Vector {
	return Vector{vector1.X - vector2.X, vector1.Y - vector2.Y}
}

//坐标点相乘

func (vector1 Vector) multi(speed float32) Vector {
	return Vector{vector1.X * speed, vector1.Y * speed}
}

//计算距离

func (vector1 Vector) distanceTo(vector2 Vector) float32 {
	dX := vector1.X * vector2.X
	dY := vector1.Y * vector2.Y
	//根号
	distance := math.Sqrt(float64(dX*dX + dY*dY))
	fmt.Printf("%v,%v,%v \n", dX, dY, float32(distance))
	return float32(distance)
}

// 适量单位化
// 标记
func (vector1 Vector) normalize() Vector {
	mag := vector1.X*vector1.X + vector1.Y*vector1.Y
	if mag > 0 {
		oneOverMag := 1 / float32(math.Sqrt(float64(mag)))
		return Vector{vector1.X * oneOverMag, vector1.Y * oneOverMag}
	} else {
		return Vector{0, 0}
	}

}
