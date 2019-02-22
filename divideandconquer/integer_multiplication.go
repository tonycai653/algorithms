package divideandconquer

import (
	"math"
	"math/big"
	"strconv"
)

/*
Karatsuba 解决两个数x和y相乘的问题, 使用分值法:

假设x,y的位数相等，那么x,y可以分别写成这种形式:

  x = x1 * (B ^ m) + x0
  y = y1 * (B ^ m) + y0

也就是:
  xy = z2 * (B ^ 2m) + z1 * (B ^ m) + z0

其中z0, z1, z2的值分别如下:
  z2 = x1y1, z1 = x1y0 + x0y1, z0 = x0y0

z1 可以简化成如下的形式:
  z1 = (x1 + x0)(y1 + y0) - x1y1 - x0y0 = (x1 + x0)(y1 + y0) - z2 - z0
*/
func Karatsuba(num1, num2 *big.Int) (product *big.Int) {
	if num1 == nil || num2 == nil {
		panic("numbers pointer should not be nil")
	}
	big10 := big.NewInt(10)
	if num1.Cmp(big10) < 0 || num2.Cmp(big10) < 0 {
		return big.NewInt(num1 * num2)
	}
	m := min(num1, num2)

	m2 := lengthOf(m).Div(m, big.NewInt(2))

	high1, low1 := SplitAt(num1, m2)
	high2, low2 := SplitAt(num2, m2)

	z0 := Karatsuba(low1, low2)
	z2 := Karatsuba(high1, high2)
	z1 := Karatsuba(low1+high1, low2+high2)

	t := big.NewInt(z1*int64(math.Pow(10, float64(m2))) + z0)

	if product == nil {
		product = new(big.Int)
	}
	product.Mul(big.NewInt(z2), big.NewInt(int64(math.Pow(10, float64(2*m2)))))
	product.Add(product, t)

	return
}

func min(a, b *big.Int) *big.Int {
	if a.Cmp(b) < 0 {
		return a
	}
	return b
}

func lengthOf(num *big.Int) int {

}

func SplitAt(num *big.Int, spos *big.Int) (high, low *big.Int) {

	p := new(big.Int)
	p.Exp(big.NewInt(10), spos, nil)

	return p.DivMod(num, p, low)
}
