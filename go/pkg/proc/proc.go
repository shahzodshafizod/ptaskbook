package proc

import "math"

func PowerA3(a float64, b *float64) {
	*b = a * a * a
}

func PowerA234(a float64, b, c, d *float64) {
	*b = a * a
	*c = *b * a
	*d = *c * a
}

func Mean(x, y float64, aMean, gMean *float64) {
	*aMean = (x + y) / 2
	*gMean = math.Sqrt(x * y)
}

func TrianglePS(a float64, p, s *float64) {
	*p = 3 * a
	*s = a * a / 4 * math.Sqrt(3)
}

func RectPS(x1, y1, x2, y2 float64, p, s *float64) {
	var a = math.Abs(x2 - x1)
	var b = math.Abs(y2 - y1)
	*p = 2 * (a + b)
	*s = a * b
}

func DigitCountSum(k int, c, s *int) {
	*c, *s = 0, 0
	var digit int
	for k > 0 {
		digit = k % 10
		*c++
		*s += digit
		k /= 10
	}
}

func InvDigits(k *int) {
	var invertedNumber = 0
	for *k > 0 {
		invertedNumber = invertedNumber*10 + *k%10
		*k /= 10
	}
	*k = invertedNumber
}

func AddRightDigit(d int, k *int) {
	*k = *k*10 + d
}

func AddLeftDigit(d int, k *int) {
	var factor = 1
	var tempK = *k
	for tempK > 0 {
		factor *= 10
		tempK /= 10
	}
	*k = d*factor + *k
}

func Swap(x, y *float64) {
	*x, *y = *y, *x
}

func Minmax(x, y *float64) {
	if *x > *y {
		Swap(x, y)
	}
}

func SortInc(x, y *int) {
	if *x > *y {
		*x, *y = *y, *x
	}
}

func SortInc3(a, b, c *float64) {
	Minmax(a, b)
	Minmax(a, c)
	Minmax(b, c)
}

func SortDec3(a, b, c *float64) {
	Minmax(c, b)
	Minmax(c, a)
	Minmax(b, a)
}

func ShiftRight3(a, b, c *float64) {
	*a, *b, *c = *c, *a, *b
}

func ShiftLeft3(a, b, c *float64) {
	*a, *b, *c = *b, *c, *a
}

func Sign(x float64) int {
	if x < 0 {
		return -1
	} else if x == 0 {
		return 0
	} else {
		return 1
	}
}

func RootCount(a, b, c float64) int {
	d := b*b - 4*a*c
	if d > 0 {
		return 2
	} else if d == 0 {
		return 1
	}
	return 0
}

func CircleS(r float64) float64 {
	const pi = 3.14
	return pi * r * r
}

func RingS(r1, r2 float64) float64 {
	const pi = 3.14
	var s1 = pi * r1 * r1
	var s2 = pi * r2 * r2
	return math.Abs(s1 - s2)
}

func TriangleP(a, h float64) float64 {
	var b = math.Sqrt(math.Pow(a/2, 2) + h*h)
	var c = b
	return a + b + c
}

func SumRange(a, b int) int {
	var sum = 0
	for element := a; element <= b; element++ {
		sum += element
	}
	return sum
}

func Calc(a, b float64, op int) float64 {
	switch op {
	case 1:
		return a - b
	case 2:
		return a * b
	case 3:
		return a / b
	default:
		return a + b
	}
}

func Quarter(x, y float64) int {
	if x > 0 {
		if y > 0 {
			return 1
		} else {
			return 4
		}
	} else if y > 0 {
		return 2
	}
	return 3
}

func Even(k int) bool {
	return k%2 == 0
}

func IsSquare(k int) bool {
	var number = 1
	for number*number < k {
		number++
	}
	return number*number == k
}

func IsPower5(k int) bool {
	return IsPowerN(k, 5)
}

func IsPowerN(k, n int) bool {
	var number = 1
	for number < k {
		number *= n
	}
	return number == k
}

func IsPrime(n int) bool {
	for index := int(math.Sqrt(float64(n))); index > 1; index-- {
		if n%index == 0 {
			return false
		}
	}
	return true
}

func DigitCount(k int) int {
	var count = 0
	for k > 0 {
		count++
		k /= 10
	}
	return count
}

func DigitN(k, n int) int {
	var digit = -1
	var index = 0
	for k > 0 {
		index++
		if index == n {
			digit = k % 10
			break
		}
		k /= 10
	}
	return digit
}

func IsPalindrome(k int) bool {
	invertedK := k
	InvDigits(&invertedK)
	return k == invertedK
}

func DegToRad(degree float64) float64 {
	const pi = 3.14
	return degree / 180 * pi
}

func RadToDeg(radian float64) float64 {
	const pi = 3.14
	return 180 / pi * radian
}

func Fact(n int) float64 {
	var factorial float64 = 1
	for index := 2; index <= n; index++ {
		factorial *= float64(index)
	}
	return factorial
}

func Fact2(n int) float64 {
	var factorial float64 = 1
	for index := n; index > 0; index -= 2 {
		factorial *= float64(index)
	}
	return factorial
}

func Fib(n int) int {
	var prevF int
	var currF = 1
	for index := 2; index <= n; index++ {
		currF += prevF
		prevF = currF - prevF
	}
	return currF
}

func Power1(a, b float64) float64 {
	if a <= 0 {
		return 0
	}
	return math.Exp(b * math.Log(a))
}

func Power2(a float64, n int) float64 {
	if n == 0 {
		return 1
	}
	var isNegative = n < 0
	if isNegative {
		n = -n
	}
	var degree = 1.0
	for index := 0; index < n; index++ {
		degree *= a
	}
	if !isNegative {
		return degree
	}
	return 1.0 / degree
}

func Power3(a, b float64) float64 {
	var intPart = int(b)
	if b-float64(intPart) == 0 {
		return Power2(a, intPart)
	}
	return Power1(a, b)
}

func Exp1(x, e float64) float64 {
	var (
		sum       = 0.0
		degree    = 1.0
		factorial = 1.0
		index     = 0
		term      float64
	)
	for {
		term = degree / factorial
		if term < e {
			break
		}
		sum += term

		degree *= x
		index++
		factorial *= float64(index)
	}
	return sum
}

func Sin1(x, e float64) float64 {
	var (
		degree    = 1.0
		factorial = 1.0
		sum       = 0.0
		sign      = 1.0
		index     = 0
		term      float64
	)
	for {
		degree *= x
		index++
		factorial *= float64(index)

		term = degree / factorial
		if math.Abs(term) < e {
			break
		}
		sum += sign * term
		sign *= -1

		degree *= x
		index++
		factorial *= float64(index)
	}
	return sum
}

func Cos1(x, e float64) float64 {
	var (
		degree    = 1.0
		factorial = 1.0
		sign      = 1.0
		term      float64
		sum       = 0.0
		index     = 0
	)
	for {
		term = degree / factorial
		if math.Abs(term) < e {
			break
		}
		sum += sign * term

		sign *= -1
		degree *= x * x
		index += 2
		factorial *= float64((index - 1) * index)
	}
	return sum
}

func Ln1(x, e float64) float64 {
	var (
		degree = 1.0
		index  = 1
		sign   = 1.0
		sum    = 0.0
		term   float64
	)
	for {
		degree *= x
		term = degree / float64(index)
		if math.Abs(term) < e {
			break
		}
		sum += sign * term
		sign *= -1
		index++
	}
	return sum
}

func Atan1(x, e float64) float64 {
	var (
		sum    = 0.0
		degree = 1.0
		sign   = 1.0
		term   float64
		index  = 1
	)
	for {
		degree *= x
		term = degree / float64(index)
		if math.Abs(term) < e {
			break
		}
		sum += sign * term
		sign *= -1
		degree *= x
		index += 2
	}
	return sum
}

func Power4(x, a, e float64) float64 {
	var (
		sum       = 0.0
		mult      = 1.0
		degree    = 1.0
		factorial = 1.0
		term      float64
		index     = 0.0
	)
	for {
		term = mult / factorial * degree
		if math.Abs(term) < e {
			break
		}
		sum += term
		mult *= (a - index)
		degree *= x
		index++
		factorial *= index
	}
	return sum
}

func Gcd2(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Frac1(a, b int, p, q *int) {
	greatestCommonDivisor := Gcd2(a, b)
	*p = a / greatestCommonDivisor
	*q = b / greatestCommonDivisor
	if *q < 0 {
		*p *= -1
		*q *= -1
	}
}

func Lcm2(a, b int) int {
	return b / Gcd2(a, b) * a
}

func Gcd3(a, b, c int) int {
	return Gcd2(Gcd2(a, b), c)
}

func TimeToHMS(t int, h, m, s *int) {
	*h = t / 3600
	t %= 3600
	*m, *s = t/60, t%60
}

func IncTime(h, m, s *int, t int) {
	*s += t % 60

	*m += t%3600/60 + *s/60
	*s %= 60

	*h += t/3600 + *m/60
	*m %= 60
}

func IsLeapYear(y int) bool {
	if y%400 == 0 || (y%100 != 0 && y%4 == 0) {
		return true
	}
	return false
}

func MonthDays(m, y int) int {
	switch m {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	case 2:
		if IsLeapYear(y) {
			return 29
		} else {
			return 28
		}
	default:
		return 0
	}
}

func PrevDate(d, m, y *int) {
	switch *d {
	case 1:
		*m--
		if *m < 1 {
			*m = 12
			*y--
		}
		*d = MonthDays(*m, *y)
	default:
		*d--
	}
}

func NextDate(d, m, y *int) {
	switch *d {
	case MonthDays(*m, *y):
		*m++
		if *m == 13 {
			*m = 1
			*y++
		}
		*d = 1
	default:
		*d++
	}
}

func Leng(xA, yA, xB, yB float64) float64 {
	return math.Sqrt(math.Pow(xB-xA, 2) + math.Pow(yB-yA, 2))
}

func Perim(xA, yA, xB, yB, xC, yC float64) float64 {
	var ab = Leng(xA, yA, xB, yB)
	var ac = Leng(xA, yA, xC, yC)
	var bc = Leng(xB, yB, xC, yC)
	return ab + ac + bc
}

func Area(xA, yA, xB, yB, xC, yC float64) float64 {
	halfP := Perim(xA, yA, xB, yB, xC, yC) / 2
	ab := Leng(xA, yA, xB, yB)
	ac := Leng(xA, yA, xC, yC)
	bc := Leng(xB, yB, xC, yC)
	return math.Sqrt(halfP * (halfP - ab) * (halfP - ac) * (halfP - bc))
}

func Dist(xP, yP, xA, yA, xB, yB float64) float64 {
	return Area(xP, yP, xA, yA, xB, yB) / Leng(xA, yA, xB, yB) * 2
}

func Alts(xA, yA, xB, yB, xC, yC float64, hA, hB, hC *float64) {
	*hA = Dist(xA, yA, xB, yB, xC, yC)
	*hB = Dist(xB, yB, xA, yA, xC, yC)
	*hC = Dist(xC, yC, xA, yA, xB, yB)
}
