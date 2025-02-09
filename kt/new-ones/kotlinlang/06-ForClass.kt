package pack06for

import kotlin.math.sin

class ForClass {

    var datScanner = scanner.Scanner()
    var ansScanner = scanner.Scanner()

    fun solve(taskNo: Int): Boolean {
        val path = "src/pack06for/tests/For$taskNo/"
        var fileName: String
        for (index in 1..100) {
            fileName = path + "%03d".format(index)
            if (!datScanner.changeReader("$fileName.dat")) {
                println("For$taskNo couldn't open datScanner source")
                return false
            }
            if (!ansScanner.changeReader("$fileName.ans")) {
                println("For$taskNo couldn't open ansScanner source")
                return false
            }
            val ok = when (taskNo) {
                1 -> this.for01()
                2 -> this.for02()
                3 -> this.for03()
                4 -> this.for04()
                5 -> this.for05()
                6 -> this.for06()
                7 -> this.for07()
                8 -> this.for08()
                9 -> this.for09()
                10 -> this.for10()
                11 -> this.for11()
                12 -> this.for12()
                13 -> this.for13()
                14 -> this.for14()
                15 -> this.for15()
                16 -> this.for16()
                17 -> this.for17()
                18 -> this.for18()
                19 -> this.for19()
                20 -> this.for20()
                21 -> this.for21()
                22 -> this.for22()
                23 -> this.for23()
                24 -> this.for24()
                25 -> this.for25()
                26 -> this.for26()
                27 -> this.for27()
                28 -> this.for28()
                29 -> this.for29()
                30 -> this.for30()
                31 -> this.for31()
                32 -> this.for32()
                33 -> this.for33()
                34 -> this.for34()
                35 -> this.for35()
                36 -> this.for36()
                37 -> this.for37()
                38 -> this.for38()
                39 -> this.for39()
                40 -> this.for40()
                else -> false
            }
            scanner.removeFiles()
            if (!ok) {
                println("For$taskNo Test#$index has not been tested")
                return false
            }
            if (!datScanner.eof()) {
                println("For$taskNo Test#$index input is not enough")
                return false
            }
            if (!ansScanner.eof()) {
                println("For$taskNo Test#$index output is not enough")
                return false
            }
        }
        println("For$taskNo Tested!")
        return true
    }

    private fun for01(): Boolean {
        val k = this.datScanner.nextInt()
        val n = this.datScanner.nextInt()
        for (index in 1..n) {
            if (!this.ansScanner.checkInt(k)) {
                return false
            }
        }
        return true
    }

    private fun for02(): Boolean {
        val a = this.datScanner.nextInt()
        val b = this.datScanner.nextInt()
        var n = 0
        for (index in a..b) {
            n++
            if (!this.ansScanner.checkInt(index)) {
                return false
            }
        }
        return this.ansScanner.checkInt(n)
    }

    private fun for03(): Boolean {
        val a = this.datScanner.nextInt()
        val b = this.datScanner.nextInt()
        var n = 0
        for (index in b-1 downTo a+1) {
            n++
            if (!this.ansScanner.checkInt(index)) {
                return false
            }
        }
        return this.ansScanner.checkInt(n)
    }

    private fun for04(): Boolean {
        val price = this.datScanner.nextDouble()
        for (index in 1..10) {
            if (!this.ansScanner.checkDouble(2, price * index)) {
                return false
            }
        }
        return true
    }

    private fun for05(): Boolean {
        var price = this.datScanner.nextDouble()
        price /= 10
        for (index in 1..10) {
            if (!this.ansScanner.checkDouble(2, price * index)) {
                return false
            }
        }
        return true
    }

    private fun for06(): Boolean {
        var price = this.datScanner.nextDouble()
        price /= 10
        for (index in 12..20 step 2) {
            if (!this.ansScanner.checkDouble(2, price * index)) {
                return false
            }
        }
        return true
    }

    private fun for07(): Boolean {
        val a = this.datScanner.nextInt()
        val b = this.datScanner.nextInt()
        var sum = 0
        for (index in a..b) {
            sum += index
        }
        return this.ansScanner.checkInt(sum)
    }

    private fun for08(): Boolean {
        val a = this.datScanner.nextInt()
        val b = this.datScanner.nextInt()
        var mul = 1
        for (index in a..b) {
            mul *= index
        }
        return this.ansScanner.checkInt(mul)
    }

    private fun for09(): Boolean {
        val a = this.datScanner.nextInt()
        val b = this.datScanner.nextInt()
        var sum = 0
        for (index in a..b) {
            sum += index * index
        }
        return this.ansScanner.checkInt(sum)
    }

    private fun for10(): Boolean {
        val n = this.datScanner.nextInt()
        var sum = 0.0
        for (index in 1..n) {
            sum += 1.0 / index.toDouble()
        }
        return this.ansScanner.checkDouble(6, sum)
    }

    private fun for11(): Boolean {
        val n = this.datScanner.nextInt()
        var sum = 0
        for (index in 0..n) {
            sum += n * n + 2 * n * index + index * index
        }
        return this.ansScanner.checkInt(sum)
    }

    private fun for12(): Boolean {
        val n = this.datScanner.nextInt()
        var factor = 1.1
        var pro = 1.0
        for (index in 1..n) {
            pro *= factor
            factor += 0.1
        }
        return this.ansScanner.checkDouble(6, pro)
    }

    private fun for13(): Boolean {
        val n = this.datScanner.nextInt()
        var number = 1.1
        var sign = 1
        var result = 0.0
        for (index in 1..n) {
            result += sign * number
            sign *= -1
            number += 0.1
        }
        return this.ansScanner.checkDouble(2, result)
    }

    private fun for14(): Boolean {
        val n = this.datScanner.nextInt()
        var sum = 0
        for (index in 1..n) {
            sum += 2 * index - 1
            if (!this.ansScanner.checkInt(sum)) {
                return false
            }
        }
        return true
    }

    private fun for15(): Boolean {
        val a = this.datScanner.nextDouble()
        val n = this.datScanner.nextInt()
        var power = 1.0
        for (index in 1..n) {
            power *= a
        }
        return this.ansScanner.checkDouble(2, power)
    }

    private fun for16(): Boolean {
        val a = this.datScanner.nextDouble()
        val n = this.datScanner.nextInt()
        var power = 1.0
        for (k in 1..n) {
            power *= a
            if (!this.ansScanner.checkDouble(2, power)) {
                return false
            }
        }
        return true
    }

    private fun for17(): Boolean {
        val a = this.datScanner.nextDouble()
        val n = this.datScanner.nextInt()
        var power = 1.0
        var sum = 0.0
        for (index in 0..n) {
            sum += power
            power *= a
        }
        return this.ansScanner.checkDouble(2, sum)
    }

    private fun for18(): Boolean {
        val a = this.datScanner.nextDouble()
        val n = this.datScanner.nextInt()
        var power = 1.0
        var sign = 1
        var result = 0.0
        for (index in 0..n) {
            result += sign * power
            sign *= -1
            power *= a
        }
        return this.ansScanner.checkDouble(2, result)
    }

    private fun for19(): Boolean {
        val n = this.datScanner.nextInt()
        var factorial = 1.0
        for (index in 1..n) {
            factorial *= index
        }
        return this.ansScanner.checkDouble(2, factorial)
    }

    private fun for20(): Boolean {
        val n = this.datScanner.nextInt()
        var factorial = 1.0
        var sum = 0.0
        for (index in 1..n) {
            factorial *= index
            sum += factorial
        }
        return this.ansScanner.checkDouble(2, sum)
    }

    private fun for21(): Boolean {
        val n = this.datScanner.nextInt()
        var factorial = 1.0
        var sum = 0.0
        for (index in 0..n) {
            sum += 1 / factorial
            factorial *= index + 1
        }
        return this.ansScanner.checkDouble(8, sum)
    }

    private fun for22(): Boolean {
        val x = this.datScanner.nextDouble()
        val n = this.datScanner.nextInt()
        var factorial = 1.0
        var degree = 1.0
        var result = 0.0
        for (index in 0..n) {
            result += degree / factorial
            degree *= x
            factorial *= index + 1
        }
        return this.ansScanner.checkDouble(8, result)
    }

    private fun for23(): Boolean {
        val x = this.datScanner.nextDouble()
        val n = this.datScanner.nextInt()
        var degree = 1.0
        var factorial = 1.0
        var sign = 1
        var result = 0.0
        for (index in 0..n) {
            degree *= x
            if (index != 0) {
                factorial *= (2*index) * (2*index+1)
            }
            result += sign * degree / factorial
            sign *= -1
            degree *= x
        }
        return this.ansScanner.checkDouble(8, result)
    }

    private fun for24(): Boolean {
        val x = this.datScanner.nextDouble()
        val n = this.datScanner.nextInt()
        var result = 0.0
        var degree = 1.0
        var factorial = 1.0
        var sign = 1
        for (index in 0..n) {
            if (index != 0) {
                factorial *= (2*index-1) * (2*index)
            }
            result += sign * degree / factorial
            sign *= -1
            degree *= x * x
        }
        return this.ansScanner.checkDouble(8, result)
    }

    private fun for25(): Boolean {
        val x = this.datScanner.nextDouble()
        val n = this.datScanner.nextInt()
        var result = 0.0
        var sign = 1
        var degree = 1.0
        for (index in 1..n) {
            degree *= x
            result += sign * degree / index
            sign *= -1
        }
        return this.ansScanner.checkDouble(8, result)
    }

    private fun for26(): Boolean {
        val x = this.datScanner.nextDouble()
        val n = this.datScanner.nextInt()
        var result = 0.0
        var degree = 1.0
        var sign = 1
        for (index in 0..n) {
            degree *= x
            result += sign * degree / (2*index+1)
            sign *= -1
            degree *= x
        }
        return this.ansScanner.checkDouble(8, result)
    }

    private fun for27(): Boolean {
        val x = this.datScanner.nextDouble()
        val n = this.datScanner.nextInt()
        var sum = 0.0
        var oddsPro = 1.0
        var evensPro = 1.0
        var degree = 1.0
        for (index in 0..n) {
            degree *= x
            if (index != 0) {
                evensPro *= 2 * index
                oddsPro *= 2 * index - 1
            }
            sum += oddsPro * degree / (evensPro * (2*index+1))
            degree *= x
        }
        return this.ansScanner.checkDouble(8, sum)
    }

    private fun for28(): Boolean {
        val x = this.datScanner.nextDouble()
        val n = this.datScanner.nextInt()
        var oddsPro = 1.0
        var evensPro = 1.0
        var degree = 1.0
        var sign = 1
        var result = 1.0
        for (index in 1..n) {
            degree *= x
            evensPro *= 2 * index
            if (index > 1) {
                oddsPro *= 2 * index - 3
            }
            result += sign * oddsPro * degree / evensPro
            sign *= -1
        }
        return this.ansScanner.checkDouble(8, result)
    }

    private fun for29(): Boolean {
        val n = this.datScanner.nextInt()
        val a = this.datScanner.nextDouble()
        val b = this.datScanner.nextDouble()
        val h = (b - a) / n
        if (!this.ansScanner.checkDouble(5, h)) {
            return false
        }
        var sequence = a
        for (index in 0..n) {
            if (!this.ansScanner.checkDouble(5, sequence)) {
                return false
            }
            sequence += h
        }
        return true
    }

    private fun for30(): Boolean {
        val n = this.datScanner.nextInt()
        val a = this.datScanner.nextDouble()
        val b = this.datScanner.nextDouble()
        val h = (b - a) / n
        if (!this.ansScanner.checkDouble(5, h)) {
            return false
        }
        var fx: Double
        var sequence = a
        for (index in 0..n) {
            fx = 1 - sin(sequence)
            if (!this.ansScanner.checkDouble(5, fx)) {
                return false
            }
            sequence += h
        }
        return true
    }

    private fun for31(): Boolean {
        val n = this.datScanner.nextInt()
        var a = 2.0
        for (index in 0 until  n) {
            a = 2 + 1/a
            if (!this.ansScanner.checkDouble(5, a)) {
                return false
            }
        }
        return true
    }

    private fun for32(): Boolean {
        val n = this.datScanner.nextInt()
        var a = 1.0
        for (index in 1..n) {
            a = (a + 1) / index
            if (!this.ansScanner.checkDouble(5, a)) {
                return false
            }
        }
        return true
    }

    private fun for33(): Boolean {
        val n = this.datScanner.nextInt()
        var prev = 1
        if (!this.ansScanner.checkInt(prev)) {
            return false
        }
        var curr = 1
        for (index in 2..n) {
            if (!this.ansScanner.checkInt(curr)) {
                return false
            }
            curr += prev
            prev = curr - prev
        }
        return true
    }

    private fun for34(): Boolean {
        val n = this.datScanner.nextInt()
        var prevA = 1.0
        var nextA = 2.0
        if (!this.ansScanner.checkDouble(6, prevA)) {
            return false
        }
        var temp: Double
        for (index in 2..n) {
            if (!this.ansScanner.checkDouble(6, nextA)) {
                return false
            }
            temp = nextA
            nextA = (prevA+2*nextA)/3
            prevA = temp
        }
        return true
    }

    private fun for35(): Boolean {
        val n = this.datScanner.nextInt()
        var prevA = 1
        var currA = 2
        var nextA = 3
        if (!this.ansScanner.checkInt(prevA)) {
            return false
        }
        var temp: Int
        for (index in 2..n) {
            if (!this.ansScanner.checkInt(currA)) {
                return false
            }
            temp = nextA
            nextA = nextA+currA-2*prevA
            prevA = currA
            currA = temp
        }
        return true
    }

    private fun for36(): Boolean {
        val n = this.datScanner.nextInt()
        val k = this.datScanner.nextInt()
        var sum = 0.0
        var degree: Double
        for (number in 1..n) {
            degree = 1.0
            for (index in 1..k) {
                degree *= number
            }
            sum += degree
        }
        return this.ansScanner.checkDouble(2, sum)
    }

    private fun for37(): Boolean {
        val n = this.datScanner.nextInt()
        var sum = 0.0
        var degree: Double
        for (number in 1..n) {
            degree = 1.0
            for (index in 1..number) {
                degree *= number
            }
            sum += degree
        }
        return this.ansScanner.checkDouble(2, sum)
    }

    private fun for38(): Boolean {
        val n = this.datScanner.nextInt()
        var sum = 0.0
        var degree: Double
        for (number in 1..n) {
            degree = 1.0
            for (index in 1..n-number+1) {
                degree *= number
            }
            sum += degree
        }
        return this.ansScanner.checkDouble(2, sum)
    }

    private fun for39(): Boolean {
        val a = this.datScanner.nextInt()
        val b = this.datScanner.nextInt()
        for (k in a..b) {
            for (times in 1..k) {
                if (!this.ansScanner.checkInt(k)) {
                    return false
                }
            }
        }
        return true
    }

    private fun for40(): Boolean {
        val a = this.datScanner.nextInt()
        val b = this.datScanner.nextInt()
        for (number in a..b) {
            for (times in 1..number-a+1) {
                if (!this.ansScanner.checkInt(number)) {
                    return false
                }
            }
        }
        return true
    }
}

/*
for04:2
for05:2
for06:2
for10:6
for12:6
for13:2
for15:2
for16:2
for17:2
for18:2
for19:2
for20:2
for21:8
for22:8
for23:8
for24:8
for25:8
for26:8
for27:8
for28:8
for29:5
for30:5
for31:5
for32:5
for34:6
for36:2
for37:2
for38:2
*/
