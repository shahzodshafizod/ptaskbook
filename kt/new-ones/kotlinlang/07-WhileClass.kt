package pack07while

import kotlin.math.pow

class WhileClass {

    var datScanner = scanner.Scanner()
    var ansScanner = scanner.Scanner()

    fun solve(taskNo: Int): Boolean {
        val path = "src/pack07while/tests/While$taskNo/"
        var fileName: String
        for (index in 1..100) {
            fileName = path + "%03d".format(index)
            if (!datScanner.changeReader("$fileName.dat")) {
                println("While$taskNo couldn't open datScanner source")
                return false
            }
            if (!ansScanner.changeReader("$fileName.ans")) {
                println("While$taskNo couldn't open ansScanner source")
                return false
            }
            val ok = when (taskNo) {
                1 -> this.while01()
                2 -> this.while02()
                3 -> this.while03()
                4 -> this.while04()
                5 -> this.while05()
                6 -> this.while06()
                7 -> this.while07()
                8 -> this.while08()
                9 -> this.while09()
                10 -> this.while10()
                11 -> this.while11()
                12 -> this.while12()
                13 -> this.while13()
                14 -> this.while14()
                15 -> this.while15()
                16 -> this.while16()
                17 -> this.while17()
                18 -> this.while18()
                19 -> this.while19()
                20 -> this.while20()
                21 -> this.while21()
                22 -> this.while22()
                23 -> this.while23()
                24 -> this.while24()
                25 -> this.while25()
                26 -> this.while26()
                27 -> this.while27()
                28 -> this.while28()
                29 -> this.while29()
                30 -> this.while30()
                else -> false
            }
            scanner.removeFiles()
            if (!ok) {
                println("While$taskNo Test#$index has not been tested")
                return false
            }
            if (!datScanner.eof()) {
                println("While$taskNo Test#$index input is not enough")
                return false
            }
            if (!ansScanner.eof()) {
                println("While$taskNo Test#$index output is not enough")
                return false
            }
        }
        println("While$taskNo Tested!")
        return true
    }

    private fun while01(): Boolean {
        var a = this.datScanner.nextDouble()
        val b = this.datScanner.nextDouble()
        while (a >= b) a -= b
        return this.ansScanner.checkDouble(2, a)
    }

    private fun while02(): Boolean {
        var a = this.datScanner.nextDouble()
        val b = this.datScanner.nextDouble()
        var segmentsCount = 0
        while (a >= b) {
            a -= b
            segmentsCount++
        }
        return this.ansScanner.checkInt(segmentsCount)
    }

    private fun while03(): Boolean {
        var n = this.datScanner.nextInt()
        val k = this.datScanner.nextInt()
        var div = 0
        while (n >= k) {
            n -= k
            div++
        }
        val rem = n
        return this.ansScanner.checkInt(div, rem)
    }

    private fun while04(): Boolean {
        val n = this.datScanner.nextInt()
        var degree = 1
        while (degree < n) degree *= 3
        return this.ansScanner.checkBoolean(degree == n)
    }

    private fun while05(): Boolean {
        val n = this.datScanner.nextInt()
        var k = 0
        while (2.0.pow(k).toInt() < n) k++
        return this.ansScanner.checkInt(k)
    }

    private fun while06(): Boolean {
        val n = this.datScanner.nextInt()
        var doubleFact = 1.0
        for (index in n downTo 1 step 2) {
            doubleFact *= index
        }
        return this.ansScanner.checkDouble(2, doubleFact)
    }

    private fun while07(): Boolean {
        val n = this.datScanner.nextInt()
        var k = 1
        while (k*k <= n) {
            k++
        }
        return this.ansScanner.checkInt(k)
    }

    private fun while08(): Boolean {
        val n = this.datScanner.nextInt()
        var k = 0
        while ((k+1)*(k+1) <= n) {
            k++
        }
        return this.ansScanner.checkInt(k)
    }

    private fun while09(): Boolean {
        val n = this.datScanner.nextInt()
        var k = 0
        var degree3 = 1
        while (degree3 <= n) {
            degree3 *= 3
            k++
        }
        return this.ansScanner.checkInt(k)
    }

    private fun while10(): Boolean {
        val n = this.datScanner.nextInt()
        var k = -1
        var degree3 = 1
        do {
            degree3 *= 3
            k++
        } while (degree3 < n)
        return this.ansScanner.checkInt(k)
    }

    private fun while11(): Boolean {
        val n = this.datScanner.nextInt()
        var sum = 0
        var k = 0
        while (sum < n) {
            k++
            sum += k
        }
        return this.ansScanner.checkInt(k, sum)
    }

    private fun while12(): Boolean {
        val n = this.datScanner.nextInt()
        var sum = 0
        var k = 0
        while (sum + k + 1 <= n) {
            k++
            sum += k
        }
        return this.ansScanner.checkInt(k, sum)
    }

    private fun while13(): Boolean {
        val a = this.datScanner.nextDouble()
        var k = 0
        var sum = 0.0
        while (sum <= a) {
            k++
            sum += 1.0/k
        }
        return this.ansScanner.checkInt(k) && this.ansScanner.checkDouble(5, sum)
    }

    private fun while14(): Boolean {
        val a = this.datScanner.nextDouble()
        var k = 0
        var sum = 0.0
        while (sum < a) {
            k++
            sum += 1.0 / k
        }
        sum -= k
        k--
        return this.ansScanner.checkInt(k) && this.ansScanner.checkDouble(5, sum)
    }

    private fun while15(): Boolean {
//        While15: 2
        return false
    }

    private fun while16(): Boolean {
//        While16: 3
        return false
    }

    private fun while17(): Boolean {
        return false
    }

    private fun while18(): Boolean {
        return false
    }

    private fun while19(): Boolean {
        return false
    }

    private fun while20(): Boolean {
        return false
    }

    private fun while21(): Boolean {
        return false
    }

    private fun while22(): Boolean {
        return false
    }

    private fun while23(): Boolean {
        return false
    }

    private fun while24(): Boolean {
        return false
    }

    private fun while25(): Boolean {
        return false
    }

    private fun while26(): Boolean {
        return false
    }

    private fun while27(): Boolean {
        return false
    }

    private fun while28(): Boolean {
//        While28: 6
        return false
    }

    private fun while29(): Boolean {
//        While29: 6
        return false
    }

    private fun while30(): Boolean {
        return false
    }
}

/*
While1: 2
While6: 2
While13: 5
While14: 5
While15: 2
While16: 3
While28: 6
While29: 6
 */