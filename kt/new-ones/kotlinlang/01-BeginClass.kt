package pack01begin

import java.lang.Math.pow
import kotlin.math.abs
import kotlin.math.sqrt

class BeginClass {

    var datScanner = scanner.Scanner()
    var ansScanner = scanner.Scanner()

    fun solve(taskNo: Int): Boolean {
        val path = "src/pack01begin/tests/Begin$taskNo/"
        var fileName: String
        for (index in 1..100) {
            fileName = path + "%03d".format(index)
            if (!datScanner.changeReader("$fileName.dat")) {
                println("Begin$taskNo couldn't open datScanner source")
                return false
            }
            if (!ansScanner.changeReader("$fileName.ans")) {
                println("Begin$taskNo couldn't open ansScanner source")
                return false
            }
            val ok = when (taskNo) {
                1 -> this.begin01()
                2 -> this.begin02()
                3 -> this.begin03()
                4 -> this.begin04()
                5 -> this.begin05()
                6 -> this.begin06()
                7 -> this.begin07()
                8 -> this.begin08()
                9 -> this.begin09()
                10 -> this.begin10()
                11 -> this.begin11()
                12 -> this.begin12()
                13 -> this.begin13()
                14 -> this.begin14()
                15 -> this.begin15()
                16 -> this.begin16()
                17 -> this.begin17()
                18 -> this.begin18()
                19 -> this.begin19()
                20 -> this.begin20()
                21 -> this.begin21()
                22 -> this.begin22()
                23 -> this.begin23()
                24 -> this.begin24()
                25 -> this.begin25()
                26 -> this.begin26()
                27 -> this.begin27()
                28 -> this.begin28()
                29 -> this.begin29()
                30 -> this.begin30()
                31 -> this.begin31()
                32 -> this.begin32()
                33 -> this.begin33()
                34 -> this.begin34()
                35 -> this.begin35()
                36 -> this.begin36()
                37 -> this.begin37()
                38 -> this.begin38()
                39 -> this.begin39()
                40 -> this.begin40()
                else -> false
            }
            scanner.removeFiles()
            if (!ok) {
                println("Begin$taskNo Test#$index has not been tested")
                return false
            }
            if (!datScanner.eof()) {
                println("Begin$taskNo Test#$index input is not enough")
                return false
            }
            if (!ansScanner.eof()) {
                println("Begin$taskNo Test#$index output is not enough")
                return false
            }
        }
        println("Begin$taskNo Tested!")
        return true
    }

    private fun begin01(): Boolean {
        val a = this.datScanner.nextDouble()
        val p = 4 * a
        return this.ansScanner.checkDouble(2, p)
    }

    private fun begin02(): Boolean {
        val a = this.datScanner.nextDouble()
        val s = a * a
        return this.ansScanner.checkDouble(2, s)
    }

    private fun begin03(): Boolean {
        val a = this.datScanner.nextDouble()
        val b = this.datScanner.nextDouble()
        val s = a * b
        val p = 2 * (a + b)
        return this.ansScanner.checkDouble(2, s, p)
    }

    private fun begin04(): Boolean {
        val d = this.datScanner.nextDouble()
        val pi = 3.14
        val l = pi * d
        return this.ansScanner.checkDouble(2, l)
    }

    private fun begin05(): Boolean {
        val a = this.datScanner.nextDouble()
        val v = a * a * a
        val s = 6 * a * a
        return this.ansScanner.checkDouble(3, v, s)
    }

    private fun begin06(): Boolean {
        val a = this.datScanner.nextDouble()
        val b = this.datScanner.nextDouble()
        val c = this.datScanner.nextDouble()
        val v = a * b * c
        val s = 2 * (a * b + b * c + c * a)
        return this.ansScanner.checkDouble(3, v, s)
    }

    private fun begin07(): Boolean {
        val r = this.datScanner.nextDouble()
        val pi = 3.14
        val l = 2 * pi * r
        val s = pi * r * r
        return this.ansScanner.checkDouble(3, l, s)
    }

    private fun begin08(): Boolean {
        val a = this.datScanner.nextDouble()
        val b = this.datScanner.nextDouble()
        val average = (a + b) / 2
        return this.ansScanner.checkDouble(2, average)
    }

    private fun begin09(): Boolean {
        val a = this.datScanner.nextDouble()
        val b = this.datScanner.nextDouble()
        val gMean = sqrt(a * b)
        return this.ansScanner.checkDouble(2, gMean)
    }

    private fun begin10(): Boolean {
        var a = this.datScanner.nextDouble()
        var b = this.datScanner.nextDouble()
        a *= a
        b *= b
        val sum = a + b
        val sub = a - b
        val mul = a * b
        val div = a / b
        return this.ansScanner.checkDouble(2, sum, sub, mul, div)
    }

    private fun begin11(): Boolean {
        var a = this.datScanner.nextDouble()
        var b = this.datScanner.nextDouble()
        a = abs(a)
        b = abs(b)
        val sum = a + b
        val sub = a - b
        val mul = a * b
        val div = a / b
        return this.ansScanner.checkDouble(2, sum, sub, mul, div)
    }

    private fun begin12(): Boolean {
        val a = this.datScanner.nextDouble()
        val b = this.datScanner.nextDouble()
        val c = sqrt(a * a + b * b)
        val p = a + b + c
        return this.ansScanner.checkDouble(2, c, p)
    }

    private fun begin13(): Boolean {
        val r1 = this.datScanner.nextDouble()
        val r2 = this.datScanner.nextDouble()
        val pi = 3.14
        val s1 = pi * r1 * r1
        val s2 = pi * r2 * r2
        val s3 = abs(s1 - s2)
        return this.ansScanner.checkDouble(2, s1, s2, s3)
    }

    private fun begin14(): Boolean {
        val l = this.datScanner.nextDouble()
        val pi = 3.14
        val r = l / (2 * pi)
        val s = pi * r * r
        return this.ansScanner.checkDouble(2, r, s)
    }

    private fun begin15(): Boolean {
        val s = this.datScanner.nextDouble()
        val pi = 3.14
        val d = sqrt(4 * s / pi)
        val l = pi * d
        return this.ansScanner.checkDouble(2, d, l)
    }

    private fun begin16(): Boolean {
        val x1 = this.datScanner.nextDouble()
        val x2 = this.datScanner.nextDouble()
        val distance = abs(x2 - x1)
        return this.ansScanner.checkDouble(2, distance)
    }

    private fun begin17(): Boolean {
        val a = this.datScanner.nextDouble()
        val b = this.datScanner.nextDouble()
        val c = this.datScanner.nextDouble()
        val ac = abs(c - a)
        val bc = abs(c - b)
        val sum = ac + bc
        return this.ansScanner.checkDouble(2, ac, bc, sum)
    }

    private fun begin18(): Boolean {
        val a = this.datScanner.nextDouble()
        val b = this.datScanner.nextDouble()
        val c = this.datScanner.nextDouble()
        val ac = abs(c - a)
        val bc = abs(c - b)
        val mul = ac * bc
        return this.ansScanner.checkDouble(2, mul)
    }

    private fun begin19(): Boolean {
        val x1 = this.datScanner.nextDouble()
        val y1 = this.datScanner.nextDouble()
        val x2 = this.datScanner.nextDouble()
        val y2 = this.datScanner.nextDouble()
        val a = abs(x2 - x1)
        val b = abs(y2 - y1)
        val p = 2 * (a + b)
        val s = a * b
        return this.ansScanner.checkDouble(2, p, s)
    }

    private fun begin20(): Boolean {
        val x1 = this.datScanner.nextDouble()
        val y1 = this.datScanner.nextDouble()
        val x2 = this.datScanner.nextDouble()
        val y2 = this.datScanner.nextDouble()
        val distance = sqrt(pow(x2 - x1, 2.0) + pow(y2 - y1, 2.0))
        return this.ansScanner.checkDouble(2, distance)
    }

    private fun begin21(): Boolean {
        val x1 = this.datScanner.nextDouble()
        val y1 = this.datScanner.nextDouble()
        val x2 = this.datScanner.nextDouble()
        val y2 = this.datScanner.nextDouble()
        val x3 = this.datScanner.nextDouble()
        val y3 = this.datScanner.nextDouble()
        val a = sqrt(pow(x2 - x1, 2.0) + pow(y2 - y1, 2.0))
        val b = sqrt(pow(x3 - x2, 2.0) + pow(y3 - y2, 2.0))
        val c = sqrt(pow(x3 - x1, 2.0) + pow(y3 - y1, 2.0))
        val p = a + b + c
        val halfP = p / 2
        val s = sqrt(halfP * (halfP - a) * (halfP - b) * (halfP - c))
        return this.ansScanner.checkDouble(2, p, s)
    }

    private fun begin22(): Boolean {
        var a = this.datScanner.nextDouble()
        var b = this.datScanner.nextDouble()
        a += b
        b = a - b
        a -= b
        return this.ansScanner.checkDouble(2, a, b)
    }

    private fun begin23(): Boolean {
        var a = this.datScanner.nextDouble()
        var b = this.datScanner.nextDouble()
        var c = this.datScanner.nextDouble()
        val d = a
        a = c
        c = b
        b = d
        return this.ansScanner.checkDouble(2, a, b, c)
    }

    private fun begin24(): Boolean {
        var a = this.datScanner.nextDouble()
        var b = this.datScanner.nextDouble()
        var c = this.datScanner.nextDouble()
        val d = a
        a = b
        b = c
        c = d
        return this.ansScanner.checkDouble(2, a, b, c)
    }

    private fun begin25(): Boolean {
        val x = this.datScanner.nextDouble()
        val y = 3 * pow(x, 6.0) - 6 * x * x - 7
        return this.ansScanner.checkDouble(2, y)
    }

    private fun begin26(): Boolean {
        val x = this.datScanner.nextDouble()
        val y = 4 * pow(x - 3, 6.0) - 7 * pow(x - 3, 3.0) + 2
        return this.ansScanner.checkDouble(2, y)
    }

    private fun begin27(): Boolean {
        val a = this.datScanner.nextDouble()
        val degA2 = a * a
        val degA4 = degA2 * degA2
        val degA8 = degA4 * degA4
        return this.ansScanner.checkDouble(2, degA2, degA4, degA8)
    }

    private fun begin28(): Boolean {
        val a = this.datScanner.nextDouble()
        val degA2 = a * a
        val degA3 = degA2 * a
        val degA5 = degA3 * degA2
        val degA10 = degA5 * degA5
        val degA15 = degA10 * degA5
        return this.ansScanner.checkDouble(2, degA2, degA3, degA5, degA10, degA15)
    }

    private fun begin29(): Boolean {
        val degrees = this.datScanner.nextDouble()
        val pi = 3.14
        val radians = degrees * pi / 180
        return this.ansScanner.checkDouble(2, radians)
    }

    private fun begin30(): Boolean {
        val radians = this.datScanner.nextDouble()
        val pi = 3.14
        val degrees = 180 * radians / pi
        return this.ansScanner.checkDouble(2, degrees)
    }

    private fun begin31(): Boolean {
        val tFahrenheit = this.datScanner.nextDouble()
        val tCentigrade = (tFahrenheit - 32) / 9 * 5
        return this.ansScanner.checkDouble(2, tCentigrade)
    }

    private fun begin32(): Boolean {
        val tCentigrade = this.datScanner.nextDouble()
        val tFahrenheit = tCentigrade / 5 * 9 + 32
        return this.ansScanner.checkDouble(2, tFahrenheit)
    }

    private fun begin33(): Boolean {
        val x = this.datScanner.nextDouble()
        val a = this.datScanner.nextDouble()
        val y = this.datScanner.nextDouble()
        val oneKg = a / x
        val yKgPrice = oneKg * y
        return this.ansScanner.checkDouble(2, oneKg, yKgPrice)
    }

    private fun begin34(): Boolean {
        val x = this.datScanner.nextDouble()
        val a = this.datScanner.nextDouble()
        val y = this.datScanner.nextDouble()
        val b = this.datScanner.nextDouble()
        val oneKgChocoPrice = a / x
        val oneKgCandyPrice = b / y
        val expensive = oneKgChocoPrice / oneKgCandyPrice
        return this.ansScanner.checkDouble(2, oneKgChocoPrice, oneKgCandyPrice, expensive)
    }

    private fun begin35(): Boolean {
        val v = this.datScanner.nextDouble()
        val u = this.datScanner.nextDouble()
        val t1 = this.datScanner.nextDouble()
        val t2 = this.datScanner.nextDouble()
        val s = v * t1 + (v - u) * t2
        return this.ansScanner.checkDouble(2, s)
    }

    private fun begin36(): Boolean {
        val v1 = this.datScanner.nextDouble()
        val v2 = this.datScanner.nextDouble()
        var s = this.datScanner.nextDouble()
        val t = this.datScanner.nextDouble()
        s += t * (v1 + v2)
        return this.ansScanner.checkDouble(2, s)
    }

    private fun begin37(): Boolean {
        val v1 = this.datScanner.nextDouble()
        val v2 = this.datScanner.nextDouble()
        var s = this.datScanner.nextDouble()
        val t = this.datScanner.nextDouble()
        s = abs(s - t * (v1 + v2))
        return this.ansScanner.checkDouble(2, s)
    }

    private fun begin38(): Boolean {
        val a = this.datScanner.nextDouble()
        val b = this.datScanner.nextDouble()
        val x = -b / a
        return this.ansScanner.checkDouble(2, x)
    }

    private fun begin39(): Boolean {
        val a = this.datScanner.nextDouble()
        val b = this.datScanner.nextDouble()
        val c = this.datScanner.nextDouble()
        val d = b * b - 4 * a * c
        val x1 = (-b - sqrt(d)) / 2 / a
        val x2 = (-b + sqrt(d)) / 2 / a
        return this.ansScanner.checkDouble(2, x1, x2)
    }

    private fun begin40(): Boolean {
        val a1 = this.datScanner.nextDouble()
        val b1 = this.datScanner.nextDouble()
        val c1 = this.datScanner.nextDouble()
        val a2 = this.datScanner.nextDouble()
        val b2 = this.datScanner.nextDouble()
        val c2 = this.datScanner.nextDouble()
        val d = a1 * b2 - a2 * b1
        val x = (c1 * b2 - c2 * b1) / d
        val y = (a1 * c2 - a2 * c1) / d
        return this.ansScanner.checkDouble(2, x, y)
    }
}