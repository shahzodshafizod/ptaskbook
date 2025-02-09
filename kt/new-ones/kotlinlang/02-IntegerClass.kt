package pack02integer

class IntegerClass {

    var datScanner = scanner.Scanner()
    var ansScanner = scanner.Scanner()

    fun solve(taskNo: Int): Boolean {
        val path = "src/pack02integer/tests/Integer$taskNo/"
        var fileName: String
        for (index in 1..100) {
            fileName = path + "%03d".format(index)
            if (!datScanner.changeReader("$fileName.dat")) {
                println("Integer$taskNo couldn't open datScanner source")
                return false
            }
            if (!ansScanner.changeReader("$fileName.ans")) {
                println("Integer$taskNo couldn't open ansScanner source")
                return false
            }
            val ok = when (taskNo) {
                1 -> this.integer01()
                2 -> this.integer02()
                3 -> this.integer03()
                4 -> this.integer04()
                5 -> this.integer05()
                6 -> this.integer06()
                7 -> this.integer07()
                8 -> this.integer08()
                9 -> this.integer09()
                10 -> this.integer10()
                11 -> this.integer11()
                12 -> this.integer12()
                13 -> this.integer13()
                14 -> this.integer14()
                15 -> this.integer15()
                16 -> this.integer16()
                17 -> this.integer17()
                18 -> this.integer18()
                19 -> this.integer19()
                20 -> this.integer20()
                21 -> this.integer21()
                22 -> this.integer22()
                23 -> this.integer23()
                24 -> this.integer24()
                25 -> this.integer25()
                26 -> this.integer26()
                27 -> this.integer27()
                28 -> this.integer28()
                29 -> this.integer29()
                30 -> this.integer30()
                else -> false
            }
            scanner.removeFiles()
            if (!ok) {
                println("Integer$taskNo Test#$index has not been tested")
                return false
            }
            if (!datScanner.eof()) {
                println("Integer$taskNo Test#$index input is not enough")
                return false
            }
            if (!ansScanner.eof()) {
                println("Integer$taskNo Test#$index output is not enough")
                return false
            }
        }
        println("Integer$taskNo Tested!")
        return true
    }

    private fun integer01(): Boolean {
        val l = this.datScanner.nextInt()
        val meters = l / 100
        return this.ansScanner.checkInt(meters)
    }

    private fun integer02(): Boolean {
        val m = this.datScanner.nextInt()
        val kg = m / 1000
        return this.ansScanner.checkInt(kg)
    }

    private fun integer03(): Boolean {
        val b = this.datScanner.nextInt()
        val kBytes = b / 1024
        return this.ansScanner.checkInt(kBytes)
    }

    private fun integer04(): Boolean {
        val a = this.datScanner.nextInt()
        val b = this.datScanner.nextInt()
        val segments = a / b
        return this.ansScanner.checkInt(segments)
    }

    private fun integer05(): Boolean {
        val a = this.datScanner.nextInt()
        val b = this.datScanner.nextInt()
        val unusedPart = a % b
        return this.ansScanner.checkInt(unusedPart)
    }

    private fun integer06(): Boolean {
        val number = this.datScanner.nextInt()
        val tens = number / 10
        val ones = number % 10
        return this.ansScanner.checkInt(tens, ones)
    }

    private fun integer07(): Boolean {
        val number = this.datScanner.nextInt()
        val tens = number / 10
        val ones = number % 10
        val sum = tens + ones
        val mul = tens * ones
        return this.ansScanner.checkInt(sum, mul)
    }

    private fun integer08(): Boolean {
        val number = this.datScanner.nextInt()
        val tens = number / 10
        val ones = number % 10
        val obtainedNumber = ones * 10 + tens
        return this.ansScanner.checkInt(obtainedNumber)
    }

    private fun integer09(): Boolean {
        val number = this.datScanner.nextInt()
        val hundreds = number / 100
        return this.ansScanner.checkInt(hundreds)
    }

    private fun integer10(): Boolean {
        val number = this.datScanner.nextInt()
        val ones = number % 10
        val tens = number / 10 % 10
        return this.ansScanner.checkInt(ones, tens)
    }

    private fun integer11(): Boolean {
        val number = this.datScanner.nextInt()
        val hundreds = number / 100
        val tens = number / 10 % 10
        val ones = number % 10
        val sum = hundreds + tens + ones
        val mul = hundreds * tens * ones
        return this.ansScanner.checkInt(sum, mul)
    }

    private fun integer12(): Boolean {
        val number = this.datScanner.nextInt()
        val hundreds = number / 100
        val tens = number / 10 % 10
        val ones = number % 10
        val obtained = ones * 100 + tens * 10 + hundreds
        return this.ansScanner.checkInt(obtained)
    }

    private fun integer13(): Boolean {
        val number = this.datScanner.nextInt()
        val hundreds = number / 100
        val obtained = number % 100 * 10 + hundreds
        return this.ansScanner.checkInt(obtained)
    }

    private fun integer14(): Boolean {
        val number = this.datScanner.nextInt()
        val ones = number % 10
        val obtained = ones * 100 + number / 10
        return this.ansScanner.checkInt(obtained)
    }

    private fun integer15(): Boolean {
        val number = this.datScanner.nextInt()
        val hundreds = number / 100
        val tens = number / 10 % 10
        val ones = number % 10
        val obtained = tens * 100 + hundreds * 10 + ones
        return this.ansScanner.checkInt(obtained)
    }

    private fun integer16(): Boolean {
        val number = this.datScanner.nextInt()
        val hundreds = number / 100
        val tens = number / 10 % 10
        val ones = number % 10
        val obtained = hundreds * 100 + ones * 10 + tens
        return this.ansScanner.checkInt(obtained)
    }

    private fun integer17(): Boolean {
        val number = this.datScanner.nextInt()
        val hundreds = number / 100 % 10
        return this.ansScanner.checkInt(hundreds)
    }

    private fun integer18(): Boolean {
        val number = this.datScanner.nextInt()
        val thousands = number / 1000 % 10
        return this.ansScanner.checkInt(thousands)
    }

    private fun integer19(): Boolean {
        val n = this.datScanner.nextInt()
        val minutes = n / 60
        return this.ansScanner.checkInt(minutes)
    }

    private fun integer20(): Boolean {
        val n = this.datScanner.nextInt()
        val hours = n / 3600
        return this.ansScanner.checkInt(hours)
    }

    private fun integer21(): Boolean {
        val n = this.datScanner.nextInt()
        val seconds = n % 60
        return this.ansScanner.checkInt(seconds)
    }

    private fun integer22(): Boolean {
        val n = this.datScanner.nextInt()
        val seconds = n % 3600
        return this.ansScanner.checkInt(seconds)
    }

    private fun integer23(): Boolean {
        val n = this.datScanner.nextInt()
        val minutes = n % 3600 / 60
        return this.ansScanner.checkInt(minutes)
    }

    private fun integer24(): Boolean {
        val k = this.datScanner.nextInt()
        val weedDay = k % 7
        return this.ansScanner.checkInt(weedDay)
    }

    private fun integer25(): Boolean {
        val k = this.datScanner.nextInt()
        val weekDay = (k + 3) % 7
        return this.ansScanner.checkInt(weekDay)
    }

    private fun integer26(): Boolean {
        val k = this.datScanner.nextInt()
        val weekDay = k % 7 + 1
        return this.ansScanner.checkInt(weekDay)
    }

    private fun integer27(): Boolean {
        val k = this.datScanner.nextInt()
        val weekDay = (k + 4) % 7 + 1
        return this.ansScanner.checkInt(weekDay)
    }

    private fun integer28(): Boolean {
        val k = this.datScanner.nextInt()
        val n = this.datScanner.nextInt()
        val weekDay = (k + n - 2) % 7 + 1
        return this.ansScanner.checkInt(weekDay)
    }

    private fun integer29(): Boolean {
        val a = this.datScanner.nextInt()
        val b = this.datScanner.nextInt()
        val c = this.datScanner.nextInt()
        val kvads = a / c * (b / c)
        val freeSpace = a * b - c * c * kvads
        return this.ansScanner.checkInt(kvads, freeSpace)
    }

    private fun integer30(): Boolean {
        val year = this.datScanner.nextInt()
        val century = (year - 1) / 100 + 1
        return this.ansScanner.checkInt(century)
    }
}