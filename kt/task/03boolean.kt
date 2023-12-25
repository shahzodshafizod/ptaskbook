package task

import kotlin.math.abs

class BooleanTask {

    var datScanner = scanner.Scanner()
    var ansScanner = scanner.Scanner()

    fun solve(pathPrefix: String) {
        for (taskNo in 1..40) {
            var taskNoStr = "%03d".format(taskNo)
            var fileName: String
            for (testNo in 1..100) {
                var testNoStr = "%03d".format(testNo)
                fileName = "$pathPrefix/03boolean/Boolean$taskNoStr/$testNoStr/$testNoStr"
                if (!datScanner.changeReader("$fileName.dat")) {
                    println("Boolean$taskNoStr couldn't open datScanner source")
                    return
                }
                if (!ansScanner.changeReader("$fileName.ans")) {
                    println("Boolean$taskNoStr couldn't open ansScanner source")
                    return
                }
                val ok = when (taskNo) {
                    1 -> this.boolean1()
                    2 -> this.boolean2()
                    3 -> this.boolean3()
                    4 -> this.boolean4()
                    5 -> this.boolean5()
                    6 -> this.boolean6()
                    7 -> this.boolean7()
                    8 -> this.boolean8()
                    9 -> this.boolean9()
                    10 -> this.boolean10()
                    11 -> this.boolean11()
                    12 -> this.boolean12()
                    13 -> this.boolean13()
                    14 -> this.boolean14()
                    15 -> this.boolean15()
                    16 -> this.boolean16()
                    17 -> this.boolean17()
                    18 -> this.boolean18()
                    19 -> this.boolean19()
                    20 -> this.boolean20()
                    21 -> this.boolean21()
                    22 -> this.boolean22()
                    23 -> this.boolean23()
                    24 -> this.boolean24()
                    25 -> this.boolean25()
                    26 -> this.boolean26()
                    27 -> this.boolean27()
                    28 -> this.boolean28()
                    29 -> this.boolean29()
                    30 -> this.boolean30()
                    31 -> this.boolean31()
                    32 -> this.boolean32()
                    33 -> this.boolean33()
                    34 -> this.boolean34()
                    35 -> this.boolean35()
                    36 -> this.boolean36()
                    37 -> this.boolean37()
                    38 -> this.boolean38()
                    39 -> this.boolean39()
                    40 -> this.boolean40()
                    else -> false
                }
                scanner.removeFiles()
                if (!ok) {
                    println("Boolean$taskNoStr Test#$testNoStr has not been tested")
                    return
                }
                if (!datScanner.eof()) {
                    println("Boolean$taskNoStr Test#$testNoStr input is not enough")
                    return
                }
                if (!ansScanner.eof()) {
                    println("Boolean$taskNoStr Test#$testNoStr output is not enough")
                    return
                }
            }
            println("Boolean$taskNoStr has been tested!")
        }
    }

    private fun boolean1(): Boolean {
        val a = this.datScanner.nextInt()
        val isPositive = a > 0
        return this.ansScanner.checkBoolean(isPositive)
    }

    private fun boolean2(): Boolean {
        val a = this.datScanner.nextInt()
        val isOdd = a % 2 != 0
        return this.ansScanner.checkBoolean(isOdd)
    }

    private fun boolean3(): Boolean {
        val a = this.datScanner.nextInt()
        val isEven = a % 2 == 0
        return this.ansScanner.checkBoolean(isEven)
    }

    private fun boolean4(): Boolean {
        val a = this.datScanner.nextInt()
        val b = this.datScanner.nextInt()
        val result = a > 2 && b <= 3
        return this.ansScanner.checkBoolean(result)
    }

    private fun boolean5(): Boolean {
        val a = this.datScanner.nextInt()
        val b = this.datScanner.nextInt()
        val result = a >= 0 || b < -2
        return this.ansScanner.checkBoolean(result)
    }

    private fun boolean6(): Boolean {
        val a = this.datScanner.nextInt()
        val b = this.datScanner.nextInt()
        val c = this.datScanner.nextInt()
        val result = b in (a + 1) until c
        return this.ansScanner.checkBoolean(result)
    }

    private fun boolean7(): Boolean {
        val a = this.datScanner.nextInt()
        val b = this.datScanner.nextInt()
        val c = this.datScanner.nextInt()
        val result = b in (a + 1) until  c || b in (c + 1) until a
        return this.ansScanner.checkBoolean(result)
    }

    private fun boolean8(): Boolean {
        val a = this.datScanner.nextInt()
        val b = this.datScanner.nextInt()
        val eachOdd = a % 2 != 0 && b % 2 != 0
        return this.ansScanner.checkBoolean(eachOdd)
    }

    private fun boolean9(): Boolean {
        val a = this.datScanner.nextInt()
        val b = this.datScanner.nextInt()
        val anyOdd = a % 2 != 0 || b % 2 != 0
        return this.ansScanner.checkBoolean(anyOdd)
    }

    private fun boolean10(): Boolean {
        val a = this.datScanner.nextInt()
        val b = this.datScanner.nextInt()
        val oneOdd = (a + b) % 2 != 0
        return this.ansScanner.checkBoolean(oneOdd)
    }

    private fun boolean11(): Boolean {
        val a = this.datScanner.nextInt()
        val b = this.datScanner.nextInt()
        val equalParity = (a + b) % 2 == 0
        return this.ansScanner.checkBoolean(equalParity)
    }

    private fun boolean12(): Boolean {
        val a = this.datScanner.nextInt()
        val b = this.datScanner.nextInt()
        val c = this.datScanner.nextInt()
        val eachPositive = a > 0 && b > 0 && c > 0
        return this.ansScanner.checkBoolean(eachPositive)
    }

    private fun boolean13(): Boolean {
        val a = this.datScanner.nextInt()
        val b = this.datScanner.nextInt()
        val c = this.datScanner.nextInt()
        val anyPositive = a > 0 || b > 0 || c > 0
        return this.ansScanner.checkBoolean(anyPositive)
    }

    private fun boolean14(): Boolean {
        val a = this.datScanner.nextInt()
        val b = this.datScanner.nextInt()
        val c = this.datScanner.nextInt()
        val onePositive = (a > 0 && b <= 0 && c <= 0) ||
                (b > 0 && a <= 0 && c <= 0) ||
                (c > 0 && a <= 0 && b <= 0)
        return this.ansScanner.checkBoolean(onePositive)
    }

    private fun boolean15(): Boolean {
        val a = this.datScanner.nextInt()
        val b = this.datScanner.nextInt()
        val c = this.datScanner.nextInt()
        val twoPositive = (a > 0 && b > 0 && c <= 0) ||
                (b > 0 && c > 0 && a <= 0) ||
                (c > 0 && a > 0 && b <= 0)
        return this.ansScanner.checkBoolean(twoPositive)
    }

    private fun boolean16(): Boolean {
        val number = this.datScanner.nextInt()
        val isTwoDigitalEven = number in 10..99 && number % 2 == 0
        return this.ansScanner.checkBoolean(isTwoDigitalEven)
    }

    private fun boolean17(): Boolean {
        val number = this.datScanner.nextInt()
        val threeDigitalOdd = number in 100..999 && number % 2 != 0
        return this.ansScanner.checkBoolean(threeDigitalOdd)
    }

    private fun boolean18(): Boolean {
        val a = this.datScanner.nextInt()
        val b = this.datScanner.nextInt()
        val c = this.datScanner.nextInt()
        val hasPair = a == b || b == c || c == a
        return this.ansScanner.checkBoolean(hasPair)
    }

    private fun boolean19(): Boolean {
        val a = this.datScanner.nextInt()
        val b = this.datScanner.nextInt()
        val c = this.datScanner.nextInt()
        val hasOppositePair = a == -b || b == -c || c == -a
        return this.ansScanner.checkBoolean(hasOppositePair)
    }

    private fun boolean20(): Boolean {
        val number = this.datScanner.nextInt()
        val hundreds = number / 100
        val tens = number / 10 % 10
        val ones = number % 10
        val areDifferent = hundreds != tens && tens != ones && ones != hundreds
        return this.ansScanner.checkBoolean(areDifferent)
    }

    private fun boolean21(): Boolean {
        val number = this.datScanner.nextInt()
        val hundreds = number / 100
        val tens = number / 10 % 10
        val ones = number % 10
        val areAscending = tens in (hundreds + 1) until ones
        return this.ansScanner.checkBoolean(areAscending)
    }

    private fun boolean22(): Boolean {
        val number = this.datScanner.nextInt()
        val hundreds = number / 100
        val tens = number / 10 % 10
        val ones = number % 10
        val areOrdered = tens in (hundreds + 1) until ones ||
                tens in (ones + 1) until hundreds
        return this.ansScanner.checkBoolean(areOrdered)
    }

    private fun boolean23(): Boolean {
        val number = this.datScanner.nextInt()
        val thousands = number / 1000
        val hundreds = number / 100 % 10
        val isPalindrome = hundreds * 10 + thousands == number % 100
        return this.ansScanner.checkBoolean(isPalindrome)
    }

    private fun boolean24(): Boolean {
        val a = this.datScanner.nextDouble()
        val b = this.datScanner.nextDouble()
        val c = this.datScanner.nextDouble()
        val d = b * b - 4 * a * c
        val hasRoot = d >= 0
        return this.ansScanner.checkBoolean(hasRoot)
    }

    private fun boolean25(): Boolean {
        val x = this.datScanner.nextDouble()
        val y = this.datScanner.nextDouble()
        val inTheSecondQuarter = x < 0 && y > 0
        return this.ansScanner.checkBoolean(inTheSecondQuarter)
    }

    private fun boolean26(): Boolean {
        val x = this.datScanner.nextDouble()
        val y = this.datScanner.nextDouble()
        val inTheFourthQuarter = x > 0 && y < 0
        return this.ansScanner.checkBoolean(inTheFourthQuarter)
    }

    private fun boolean27(): Boolean {
        val x = this.datScanner.nextDouble()
        val y = this.datScanner.nextDouble()
        val result = x < 0 && y != 0.0
        return this.ansScanner.checkBoolean(result)
    }

    private fun boolean28(): Boolean {
        val x = this.datScanner.nextDouble()
        val y = this.datScanner.nextDouble()
        val result = x * y > 0
        return this.ansScanner.checkBoolean(result)
    }

    private fun boolean29(): Boolean {
        val x = this.datScanner.nextDouble()
        val y = this.datScanner.nextDouble()
        val x1 = this.datScanner.nextDouble()
        val y1 = this.datScanner.nextDouble()
        val x2 = this.datScanner.nextDouble()
        val y2 = this.datScanner.nextDouble()
        val isInside = x > x1 && x < x2 && y > y2 && y < y1
        return this.ansScanner.checkBoolean(isInside)
    }

    private fun boolean30(): Boolean {
        val a = this.datScanner.nextInt()
        val b = this.datScanner.nextInt()
        val c = this.datScanner.nextInt()
        val isEquilateral = a == b && b == c && c == a
        return this.ansScanner.checkBoolean(isEquilateral)
    }

    private fun boolean31(): Boolean {
        val a = this.datScanner.nextInt()
        val b = this.datScanner.nextInt()
        val c = this.datScanner.nextInt()
        val isIsosceles = a == b || b == c || c == a
        return this.ansScanner.checkBoolean(isIsosceles)
    }

    private fun boolean32(): Boolean {
        val a = this.datScanner.nextInt()
        val b = this.datScanner.nextInt()
        val c = this.datScanner.nextInt()
        val isRightTriangle = a*a == b*b + c*c ||
                b*b == a*a + c*c || c*c == a*a + b*b
        return this.ansScanner.checkBoolean(isRightTriangle)
    }

    private fun boolean33(): Boolean {
        val a = this.datScanner.nextInt()
        val b = this.datScanner.nextInt()
        val c = this.datScanner.nextInt()
        val existsTriangle = a + b > c && b + c > a && c + a > b
        return this.ansScanner.checkBoolean(existsTriangle)
    }

    private fun boolean34(): Boolean {
        val x = this.datScanner.nextInt()
        val y = this.datScanner.nextInt()
        val isWhite = (x + y) % 2 != 0
        return this.ansScanner.checkBoolean(isWhite)
    }

    private fun boolean35(): Boolean {
        val x1 = this.datScanner.nextInt()
        val y1 = this.datScanner.nextInt()
        val x2 = this.datScanner.nextInt()
        val y2 = this.datScanner.nextInt()
        val sameColor = (x1 + y1) % 2 == (x2 + y2) % 2
        return this.ansScanner.checkBoolean(sameColor)
    }

    private fun boolean36(): Boolean {
        val x1 = this.datScanner.nextInt()
        val y1 = this.datScanner.nextInt()
        val x2 = this.datScanner.nextInt()
        val y2 = this.datScanner.nextInt()
        val canMoveRook = x1 == x2 || y1 == y2
        return this.ansScanner.checkBoolean(canMoveRook)
    }

    private fun boolean37(): Boolean {
        val x1 = this.datScanner.nextInt()
        val y1 = this.datScanner.nextInt()
        val x2 = this.datScanner.nextInt()
        val y2 = this.datScanner.nextInt()
        val canMoveKing = abs(x2-x1) < 2 && abs(y2-y1) < 2
        return this.ansScanner.checkBoolean(canMoveKing)
    }

    private fun boolean38(): Boolean {
        val x1 = this.datScanner.nextInt()
        val y1 = this.datScanner.nextInt()
        val x2 = this.datScanner.nextInt()
        val y2 = this.datScanner.nextInt()
        val canMoveBishop = abs(x2-x1) == abs(y2-y1)
        return this.ansScanner.checkBoolean(canMoveBishop)
    }

    private fun boolean39(): Boolean {
        val x1 = this.datScanner.nextInt()
        val y1 = this.datScanner.nextInt()
        val x2 = this.datScanner.nextInt()
        val y2 = this.datScanner.nextInt()
        val canMoveQueen = x1 == x2 || y1 == y2 || abs(x2-x1) == abs(y2-y1)
        return this.ansScanner.checkBoolean(canMoveQueen)
    }

    private fun boolean40(): Boolean {
        val x1 = this.datScanner.nextInt()
        val y1 = this.datScanner.nextInt()
        val x2 = this.datScanner.nextInt()
        val y2 = this.datScanner.nextInt()
        val xs = abs(x2 - x1)
        val ys = abs(y2 - y1)
        val canMoveKnight = xs*ys == 2 && xs != ys
        return this.ansScanner.checkBoolean(canMoveKnight)
    }
}