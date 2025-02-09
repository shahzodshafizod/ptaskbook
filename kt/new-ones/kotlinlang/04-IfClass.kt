package pack04if

import kotlin.math.abs
import kotlin.math.sin

class IfClass {

    var datScanner = scanner.Scanner()
    var ansScanner = scanner.Scanner()

    fun solve(taskNo: Int): Boolean {
        val path = "src/pack04if/tests/If$taskNo/"
        var fileName: String
        for (index in 1..100) {
            fileName = path + "%03d".format(index)
            if (!datScanner.changeReader("$fileName.dat")) {
                println("If$taskNo couldn't open datScanner source")
                return false
            }
            if (!ansScanner.changeReader("$fileName.ans")) {
                println("If$taskNo couldn't open ansScanner source")
                return false
            }
            val ok = when (taskNo) {
                1 -> this.if01()
                2 -> this.if02()
                3 -> this.if03()
                4 -> this.if04()
                5 -> this.if05()
                6 -> this.if06()
                7 -> this.if07()
                8 -> this.if08()
                9 -> this.if09()
                10 -> this.if10()
                11 -> this.if11()
                12 -> this.if12()
                13 -> this.if13()
                14 -> this.if14()
                15 -> this.if15()
                16 -> this.if16()
                17 -> this.if17()
                18 -> this.if18()
                19 -> this.if19()
                20 -> this.if20()
                21 -> this.if21()
                22 -> this.if22()
                23 -> this.if23()
                24 -> this.if24()
                25 -> this.if25()
                26 -> this.if26()
                27 -> this.if27()
                28 -> this.if28()
                29 -> this.if29()
                30 -> this.if30()
                else -> false
            }
            scanner.removeFiles()
            if (!ok) {
                println("If$taskNo Test#$index has not been tested")
                return false
            }
            if (!datScanner.eof()) {
                println("If$taskNo Test#$index input is not enough")
                return false
            }
            if (!ansScanner.eof()) {
                println("If$taskNo Test#$index output is not enough")
                return false
            }
        }
        println("If$taskNo Tested!")
        return true
    }

    private fun if01(): Boolean {
        var number = this.datScanner.nextInt()
        if (number > 0) {
            number -= 8
        }
        return this.ansScanner.checkInt(number)
    }

    private fun if02(): Boolean {
        var number = this.datScanner.nextInt()
        number += if (number > 0) -8 else 6
        return this.ansScanner.checkInt(number)
    }

    private fun if03(): Boolean {
        var number = this.datScanner.nextInt()
        number += if (number > 0) -8 else if (number < 0) 6 else 10 - number
        return this.ansScanner.checkInt(number)
    }

    private fun if04(): Boolean {
        val a = this.datScanner.nextInt()
        val b = this.datScanner.nextInt()
        val c = this.datScanner.nextInt()
        var positivesCount = 0
        if (a > 0) positivesCount++
        if (b > 0) positivesCount++
        if (c > 0) positivesCount++
        return this.ansScanner.checkInt(positivesCount)
    }

    private fun if05(): Boolean {
        val a = this.datScanner.nextInt()
        val b = this.datScanner.nextInt()
        val c = this.datScanner.nextInt()
        var positivesCount = 0
        var negativesCount = 0
        if (a > 0) positivesCount++
        if (b > 0) positivesCount++
        if (c > 0) positivesCount++
        if (a < 0) negativesCount++
        if (b < 0) negativesCount++
        if (c < 0) negativesCount++
        return this.ansScanner.checkInt(positivesCount, negativesCount)
    }

    private fun if06(): Boolean {
        val a = this.datScanner.nextDouble()
        val b = this.datScanner.nextDouble()
        val larger = if (a > b) a else b
        return this.ansScanner.checkDouble(2, larger)
    }

    private fun if07(): Boolean {
        val a = this.datScanner.nextDouble()
        val b = this.datScanner.nextDouble()
        val smallerIndex = if (a < b) 1 else 2
        return this.ansScanner.checkInt(smallerIndex)
    }

    private fun if08(): Boolean {
        val a = this.datScanner.nextDouble()
        val b = this.datScanner.nextDouble()
        val larger = if (a > b) a else b
        val smaller = a + b - larger
        return this.ansScanner.checkDouble(2, larger, smaller)
    }

    private fun if09(): Boolean {
        var a = this.datScanner.nextDouble()
        var b = this.datScanner.nextDouble()
        if (a > b) {
            a += b
            b = a - b
            a -= b
        }
        return this.ansScanner.checkDouble(2, a, b)
    }

    private fun if10(): Boolean {
        var a = this.datScanner.nextInt()
        var b = this.datScanner.nextInt()
        a = if (a != b) a + b else 0
        b = a
        return this.ansScanner.checkInt(a, b)
    }

    private fun if11(): Boolean {
        var a = this.datScanner.nextInt()
        var b = this.datScanner.nextInt()
        a = if (a != b) if (a > b) a else b else 0
        b = a
        return this.ansScanner.checkInt(a, b)
    }

    private fun if12(): Boolean {
        val a = this.datScanner.nextDouble()
        val b = this.datScanner.nextDouble()
        val c = this.datScanner.nextDouble()
        val minimal = if (a < b && a < c) a else if (b < c) b else c
        return this.ansScanner.checkDouble(2, minimal)
    }

    private fun if13(): Boolean {
        val a = this.datScanner.nextDouble()
        val b = this.datScanner.nextDouble()
        val c = this.datScanner.nextDouble()
        val between = if (a < b && b < c || a > b && b > c) b
        else if (b < a && a < c || b > a && a > c) a else c
        return this.ansScanner.checkDouble(2, between)
    }

    private fun if14(): Boolean {
        val a = this.datScanner.nextDouble()
        val b = this.datScanner.nextDouble()
        val c = this.datScanner.nextDouble()
        val minimal = if (a < b && a < c) a else if (b < c) b else c
        val maximal = if (a > b && a > c) a else if (b > c) b else c
        return this.ansScanner.checkDouble(2, minimal, maximal)
    }

    private fun if15(): Boolean {
        val a = this.datScanner.nextDouble()
        val b = this.datScanner.nextDouble()
        val c = this.datScanner.nextDouble()
        val minimal = if (a < b && a < c) a else if (b < c) b else c
        val sum = a + b + c - minimal
        return this.ansScanner.checkDouble(2, sum)
    }

    private fun if16(): Boolean {
        var a = this.datScanner.nextDouble()
        var b = this.datScanner.nextDouble()
        var c = this.datScanner.nextDouble()
        val multiplier = if (a < b && b < c) 2 else -1
        a *= multiplier
        b *= multiplier
        c *= multiplier
        return this.ansScanner.checkDouble(2, a, b, c)
    }

    private fun if17(): Boolean {
        var a = this.datScanner.nextDouble()
        var b = this.datScanner.nextDouble()
        var c = this.datScanner.nextDouble()
        val multiplier = if (a < b && b < c || a > b && b > c) 2 else -1
        a *= multiplier
        b *= multiplier
        c *= multiplier
        return this.ansScanner.checkDouble(2, a, b, c)
    }

    private fun if18(): Boolean {
        val a = this.datScanner.nextInt()
        val b = this.datScanner.nextInt()
        val c = this.datScanner.nextInt()
        val differIndex = if (b == c) 1 else if (a == c) 2 else 3
        return this.ansScanner.checkInt(differIndex)
    }

    private fun if19(): Boolean {
        val a = this.datScanner.nextInt()
        val b = this.datScanner.nextInt()
        val c = this.datScanner.nextInt()
        val d = this.datScanner.nextInt()
        val differIndex = if (b == c && c == d) 1
        else if (a == c && c == d) 2
        else if (a == b && b == d) 3 else 4
        return this.ansScanner.checkInt(differIndex)
    }

    private fun if20(): Boolean {
        val a = this.datScanner.nextDouble()
        val b = this.datScanner.nextDouble()
        val c = this.datScanner.nextDouble()
        val ab = abs(b - a)
        val ac = abs(c - a)
        val minDistance: Double
        val closerPoint: Double
        if (ab < ac) {
            minDistance = ab
            closerPoint = b
        } else {
            minDistance = ac
            closerPoint = c
        }
        return this.ansScanner.checkDouble(2, closerPoint, minDistance)
    }

    private fun if21(): Boolean {
        val x = this.datScanner.nextInt()
        val y = this.datScanner.nextInt()
        val output = if (x == 0 && y == 0) 0
        else if (y == 0) 1 else if (x == 0) 2 else 3
        return this.ansScanner.checkInt(output)
    }

    private fun if22(): Boolean {
        val x = this.datScanner.nextDouble()
        val y = this.datScanner.nextDouble()
        val quarterNo = if (x > 0) if (y > 0) 1 else 4
        else if (y > 0) 2 else 3
        return this.ansScanner.checkInt(quarterNo)
    }

    private fun if23(): Boolean {
        val x1 = this.datScanner.nextInt()
        val y1 = this.datScanner.nextInt()
        val x2 = this.datScanner.nextInt()
        val y2 = this.datScanner.nextInt()
        val x3 = this.datScanner.nextInt()
        val y3 = this.datScanner.nextInt()
        val x4 = if (x2 == x3) x1 else if (x1 == x3) x2 else x3
        val y4 = if (y2 == y3) y1 else if (y1 == y3) y2 else y3
        return this.ansScanner.checkInt(x4, y4)
    }

    private fun if24(): Boolean {
        val x = this.datScanner.nextDouble()
        val fx = if (x > 0) 2 * sin(x) else 6 - x
        return this.ansScanner.checkDouble(2, fx)
    }

    private fun if25(): Boolean {
        val x = this.datScanner.nextInt()
        val fx = if (x < -2 || x > 2) 2 * x else -3 * x
        return this.ansScanner.checkInt(fx)
    }

    private fun if26(): Boolean {
        val x = this.datScanner.nextDouble()
        val fx = if (x <= 0) -x else if (x >= 2) 4.0 else x * x
        return this.ansScanner.checkDouble(2, fx)
    }

    private fun if27(): Boolean {
        val x = this.datScanner.nextDouble()
        val fx = if (x < 0) 0 else if (x.toInt() % 2 == 0) 1 else -1
        return this.ansScanner.checkInt(fx)
    }

    private fun if28(): Boolean {
        val year = this.datScanner.nextInt()
        val days = if (year % 400 == 0) 366
        else if (year % 100 == 0) 365
        else if (year % 4 == 0) 366 else 365
        return this.ansScanner.checkInt(days)
    }

    private fun if29(): Boolean {
        val number = this.datScanner.nextInt()
        var description = if (number < 0) "negative"
        else if (number == 0) "zero" else "positive"
        description += if (number == 0) ""
        else if (number % 2 == 0) " even" else " odd"
        description += " number"
        return this.ansScanner.checkLine(description)
    }

    private fun if30(): Boolean {
        val number = this.datScanner.nextInt()
        var description = if (number in 1..9) "one"
        else if (number in 10..99) "two"
        else if (number in 100..999) "three"
        else "x"
        description += "-digit " + if (number % 2 == 0) "even" else "odd"
        description += " number"
        return this.ansScanner.checkLine(description)
    }
}