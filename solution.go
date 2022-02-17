package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type Num int

const SidesCircle = 0
const SidesTriangle = 3
const SidesSquare = 4

func CalcSquare(sideLen float64, sidesNum Num) float64 {

	var area float64

	if sidesNum == SidesCircle {
		area = math.Pi * math.Pow(sideLen, 2.0)
	} else if sidesNum == SidesTriangle {
		area = math.Sqrt(3) * math.Pow(sideLen, 2.0) / 4
	} else if sidesNum == SidesSquare {
		area = math.Pow(sideLen, 2)
	} else {
		area = 0
	}
	return area
}
