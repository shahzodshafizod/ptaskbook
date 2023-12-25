package task

import kotlin.math.pow

class WhileTask {

    var datScanner = scanner.Scanner()
    var ansScanner = scanner.Scanner()

    fun solve(pathPrefix: String) {
        for (taskNo in 1..30) {
            var taskNoStr = "%03d".format(taskNo)
            var fileName: String
            for (testNo in 1..100) {
                var testNoStr = "%03d".format(testNo)
                fileName = "$pathPrefix/07while/While$taskNoStr/$testNoStr/$testNoStr"
                if (!datScanner.changeReader("$fileName.dat")) {
                    println("While$taskNoStr couldn't open datScanner source")
                    return
                }
                if (!ansScanner.changeReader("$fileName.ans")) {
                    println("While$taskNoStr couldn't open ansScanner source")
                    return
                }
                val ok = when (taskNo) {
                    1 -> this.while1()
                    2 -> this.while2()
                    3 -> this.while3()
                    4 -> this.while4()
                    5 -> this.while5()
                    6 -> this.while6()
                    7 -> this.while7()
                    8 -> this.while8()
                    9 -> this.while9()
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
                    println("While$taskNoStr Test#$testNoStr has not been tested")
                    return
                }
                if (!datScanner.eof()) {
                    println("While$taskNoStr Test#$testNoStr input is not enough")
                    return
                }
                if (!ansScanner.eof()) {
                    println("While$taskNoStr Test#$testNoStr output is not enough")
                    return
                }
            }
            println("While$taskNoStr has been tested!")
        }
    }

    private fun while1(): Boolean {
        var a = this.datScanner.nextDouble()
        val b = this.datScanner.nextDouble()
        while (a >= b) a -= b
        return this.ansScanner.checkDouble(2, a)
    }

    private fun while2(): Boolean {
        var a = this.datScanner.nextDouble()
        val b = this.datScanner.nextDouble()
        var segmentsCount = 0
        while (a >= b) {
            a -= b
            segmentsCount++
        }
        return this.ansScanner.checkInt(segmentsCount)
    }

    private fun while3(): Boolean {
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

    private fun while4(): Boolean {
        val n = this.datScanner.nextInt()
        var degree = 1
        while (degree < n) degree *= 3
        return this.ansScanner.checkBoolean(degree == n)
    }

    private fun while5(): Boolean {
        val n = this.datScanner.nextInt()
        var k = 0
        while (2.0.pow(k).toInt() < n) k++
        return this.ansScanner.checkInt(k)
    }

    private fun while6(): Boolean {
        val n = this.datScanner.nextInt()
        var doubleFact = 1.0
        for (index in n downTo 1 step 2) {
            doubleFact *= index
        }
        return this.ansScanner.checkDouble(2, doubleFact)
    }

    private fun while7(): Boolean {
        val n = this.datScanner.nextInt()
        var k = 1
        while (k*k <= n) {
            k++
        }
        return this.ansScanner.checkInt(k)
    }

    private fun while8(): Boolean {
        val n = this.datScanner.nextInt()
        var k = 0
        while ((k+1)*(k+1) <= n) {
            k++
        }
        return this.ansScanner.checkInt(k)
    }

    private fun while9(): Boolean {
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
        return false
    }

    private fun while16(): Boolean {
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
        return false
    }

    private fun while29(): Boolean {
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