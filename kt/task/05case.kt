package task

import kotlin.math.sqrt

class CaseTask {

    var datScanner = scanner.Scanner()
    var ansScanner = scanner.Scanner()

    fun solve(pathPrefix: String) {
        for (taskNo in 1..20) {
            var taskNoStr = "%03d".format(taskNo)
            var fileName: String
            for (testNo in 1..100) {
                var testNoStr = "%03d".format(testNo)
                fileName = "$pathPrefix/05case/Case$taskNoStr/$testNoStr/$testNoStr"
                if (!datScanner.changeReader("$fileName.dat")) {
                    println("Case$taskNoStr couldn't open datScanner source")
                    return
                }
                if (!ansScanner.changeReader("$fileName.ans")) {
                    println("Case$taskNoStr couldn't open ansScanner source")
                    return
                }
                val ok = when (taskNo) {
                    1 -> this.case1()
                    2 -> this.case2()
                    3 -> this.case3()
                    4 -> this.case4()
                    5 -> this.case5()
                    6 -> this.case6()
                    7 -> this.case7()
                    8 -> this.case8()
                    9 -> this.case9()
                    10 -> this.case10()
                    11 -> this.case11()
                    12 -> this.case12()
                    13 -> this.case13()
                    14 -> this.case14()
                    15 -> this.case15()
                    16 -> this.case16()
                    17 -> this.case17()
                    18 -> this.case18()
                    19 -> this.case19()
                    20 -> this.case20()
                    else -> false
                }
                scanner.removeFiles()
                if (!ok) {
                    println("Case$taskNoStr Test#$testNoStr has not been tested")
                    return
                }
                if (!datScanner.eof()) {
                    println("Case$taskNoStr Test#$testNoStr input is not enough")
                    return
                }
                if (!ansScanner.eof()) {
                    println("Case$taskNoStr Test#$testNoStr output is not enough")
                    return
                }
            }
            println("Case$taskNoStr has been tested!")
        }
    }

    private fun case1(): Boolean {
        val number = this.datScanner.nextInt()
        val weekDay = when (number) {
            1 -> "Monday"
            2 -> "Tuesday"
            3 -> "Wednesday"
            4 -> "Thursday"
            5 -> "Friday"
            6 -> "Saturday"
            7 -> "Sunday"
            else -> ""
        }
        return this.ansScanner.checkWord(weekDay)
    }

    private fun case2(): Boolean {
        val k = this.datScanner.nextInt()
        val mark = when (k) {
            1 -> "bad"
            2 -> "unsatisfactory"
            3 -> "mediocre"
            4 -> "good"
            5 -> "excellent"
            else -> "error"
        }
        return this.ansScanner.checkWord(mark)
    }

    private fun case3(): Boolean {
        val monthNo = this.datScanner.nextInt()
        val season = when (monthNo) {
            1, 2, 12 -> "Winter"
            3, 4, 5 -> "Spring"
            6, 7, 8 -> "Summer"
            9, 10, 11 -> "Autumn"
            else -> ""
        }
        return this.ansScanner.checkWord(season)
    }

    private fun case4(): Boolean {
        val monthNo = this.datScanner.nextInt()
        val monthDays = when (monthNo) {
            1, 3, 5, 7, 8, 10, 12 -> 31
            4, 6, 9, 11 -> 30
            2 -> 28
            else -> 0
        }
        return this.ansScanner.checkInt(monthDays)
    }

    private fun case5(): Boolean {
        val n = this.datScanner.nextInt()
        val a = this.datScanner.nextDouble()
        val b = this.datScanner.nextDouble()
        val result = when (n) {
            1 -> a + b
            2 -> a - b
            3 -> a * b
            4 -> a / b
            else -> 0.0
        }
        return this.ansScanner.checkDouble(2, result)
    }

    private fun case6(): Boolean {
        val n = this.datScanner.nextInt()
        val l = this.datScanner.nextDouble()
        val meters = when (n) {
            1 -> l / 10
            2 -> l * 1000
            3 -> l
            4 -> l / 1000
            5 -> l / 100
            else -> 0.0
        }
        return this.ansScanner.checkDouble(2, meters)
    }

    private fun case7(): Boolean {
        val n = this.datScanner.nextInt()
        val m = this.datScanner.nextDouble()
        val kg = when (n) {
            1 -> m
            2 -> m / 100000
            3 -> m / 1000
            4 -> m * 1000
            5 -> m * 100
            else -> 0.0
        }
        return this.ansScanner.checkDouble(2, kg)
    }

    private fun case8(): Boolean {
        var d = this.datScanner.nextInt()
        var m = this.datScanner.nextInt()
        when (d) {
            1 -> {
                d = when (m) {
                    1 -> {
                        m = 13; 31
                    }
                    2, 4, 6, 8, 9, 11 -> 31
                    3 -> 28
                    5, 7, 10, 12 -> 30
                    else -> 0
                }
                m--
            }
            else -> d--
        }
        return this.ansScanner.checkInt(d, m)
    }

    private fun case9(): Boolean {
        var d = this.datScanner.nextInt()
        var m = this.datScanner.nextInt()
        when (d) {
            31 -> {
                d = 0
                m = when (m) {
                    12 -> 0
                    else -> m
                }
            }
            30 -> d = when (m) {
                4, 6, 9, 11 -> 0
                else -> d
            }
            28 -> d = when (m) {
                2 -> 0
                else -> d
            }
        }
        if (d == 0) {
            m++
        }
        d++
        return this.ansScanner.checkInt(d, m)
    }

    private fun case10(): Boolean {
        var c = this.datScanner.nextChar()
        val n = this.datScanner.nextInt()
        c = when (c) {
            'N' -> when (n) {
                1 -> 'W'
                -1 -> 'E'
                else -> c
            }
            'W' -> when (n) {
                1 -> 'S'
                -1 -> 'N'
                else -> c
            }
            'S' -> when (n) {
                1 -> 'E'
                -1 -> 'W'
                else -> c
            }
            'E' -> when (n) {
                1 -> 'N'
                -1 -> 'S'
                else -> c
            }
            else -> c
        }
        return this.ansScanner.checkChar(c)
    }

    private fun case11(): Boolean {
        var c = this.datScanner.nextChar()
        val n1 = this.datScanner.nextInt()
        val n2 = this.datScanner.nextInt()
        c = when (n1 + n2) {
            1 -> when (c) { // left;
                'N' -> 'W'
                'E' -> 'N'
                'S' -> 'E'
                'W' -> 'S'
                else -> c
            }
            3 -> when (c) { // right;
                'N' -> 'E'
                'E' -> 'S'
                'S' -> 'W'
                'W' -> 'N'
                else -> c
            }
            2, -2 -> when (c) { // 180;
                'N' -> 'S'
                'E' -> 'W'
                'S' -> 'N'
                'W' -> 'E'
                else -> c
            }
            else -> c
        }
        return this.ansScanner.checkChar(c)
    }

    private fun case12(): Boolean {
        val index = this.datScanner.nextInt()
        val value = this.datScanner.nextDouble()
        var r = 0.0
        var d = 0.0
        var l = 0.0
        var s = 0.0
        val pi = 3.14
        when (index) {
            1 -> {
                r = value
                d = 2 * r
                l = 2 * pi * r
                s = pi * r * r
            }
            2 -> {
                d = value
                r = d / 2
                l = 2 * pi * r
                s = pi * r * r
            }
            3 -> {
                l = value
                r = l / (2 * pi)
                d = 2 * r
                s = pi * r * r
            }
            4 -> {
                s = value
                r = sqrt(s / pi)
                d = 2 * r
                l = 2 * pi * r
            }
        }
        if (index != 1 && !this.ansScanner.checkDouble(2, r)) {
            return false
        }
        if (index != 2 && !this.ansScanner.checkDouble(2, d)) {
            return false
        }
        if (index != 3 && !this.ansScanner.checkDouble(2, l)) {
            return false
        }
        if (index != 4 && !this.ansScanner.checkDouble(2, s)) {
            return false
        }
        return true
    }

    private fun case13(): Boolean {
        val index = this.datScanner.nextInt()
        val value = this.datScanner.nextDouble()
        var a = 0.0
        var c = 0.0
        var h = 0.0
        var s = 0.0
        when (index) {
            1 -> {
                a = value
                c = a * sqrt(2.0)
                h = c / 2
                s = c * h / 2
            }
            2 -> {
                c = value
                a = c / sqrt(2.0)
                h = c / 2
                s = c * h / 2
            }
            3 -> {
                h = value
                c = 2 * h
                a = c / sqrt(2.0)
                s = c * h / 2
            }
            4 -> {
                s = value
                c = sqrt(4 * s)
                a = c / sqrt(2.0)
                h = c / 2
            }
        }
        if (index != 1 && !this.ansScanner.checkDouble(2, a)) {
            return false
        }
        if (index != 2 && !this.ansScanner.checkDouble(2, c)) {
            return false
        }
        if (index != 3 && !this.ansScanner.checkDouble(2, h)) {
            return false
        }
        if (index != 4 && !this.ansScanner.checkDouble(2, s)) {
            return false
        }
        return true
    }

    private fun case14(): Boolean {
        val index = this.datScanner.nextInt()
        val value = this.datScanner.nextDouble()
        var a = 0.0
        var r1 = 0.0
        var r2 = 0.0
        var s = 0.0
        when (index) {
            1 -> {
                a = value
                r1 = a * sqrt(3.0) / 6
                r2 = 2 * r1
                s = a * a * sqrt(3.0) / 4
            }
            2 -> {
                r1 = value
                a = 6 / sqrt(3.0) * r1
                r2 = 2 * r1
                s = a * a * sqrt(3.0) / 4
            }
            3 -> {
                r2 = value
                r1 = r2 / 2
                a = 6 / sqrt(3.0) * r1
                s = a * a * sqrt(3.0) / 4
            }
            4 -> {
                s = value
                a = sqrt(4 / sqrt(3.0) * s)
                r1 = a * sqrt(3.0) / 6
                r2 = 2 * r1
            }
        }
        if (index != 1 && !this.ansScanner.checkDouble(2, a)) {
            return false
        }
        if (index != 2 && !this.ansScanner.checkDouble(2, r1)) {
            return false
        }
        if (index != 3 && !this.ansScanner.checkDouble(2, r2)) {
            return false
        }
        if (index != 4 && !this.ansScanner.checkDouble(2, s)) {
            return false
        }
        return true
    }

    private fun case15(): Boolean {
        val n = this.datScanner.nextInt()
        val m = this.datScanner.nextInt()
        val description = when(n) {
            6 -> "six"
            7 -> "seven"
            8 -> "eight"
            9 -> "nine"
            10 -> "ten"
            11 -> "jack"
            12 -> "queen"
            13 -> "king"
            14 -> "ace"
            else -> ""
        } + " of " + when (m) {
            1 -> "spades"
            2 -> "clubs"
            3 -> "diamonds"
            4 -> "hearts"
            else -> ""
        }
        return this.ansScanner.checkLine(description)
    }

    private fun case16(): Boolean {
        val age = this.datScanner.nextInt()
        val description = when(age / 10) {
            2 -> "twenty"
            3 -> "thirty"
            4 -> "fourty"
            5 -> "fifty"
            6 -> "sixty"
            else -> ""
        } + when (age % 10) {
            1 -> "-one"
            2 -> "-two"
            3 -> "-three"
            4 -> "-four"
            5 -> "-five"
            6 -> "-six"
            7 -> "-seven"
            8 -> "-eight"
            9 -> "-nine"
            else -> ""
        } + " years"
        return this.ansScanner.checkLine(description)
    }

    private fun case17(): Boolean {
        val orderNumber = this.datScanner.nextInt()
        val tens = orderNumber / 10
        val ones = orderNumber % 10
        var alphabetic = "the " + when(tens) {
            1 -> when (ones) {
                0 -> "tenth"
                1 -> "eleventh"
                2 -> "twelfth"
                3 -> "thirteenth"
                4 -> "fourteenth"
                5 -> "fifteenth"
                6 -> "sixteenth"
                7 -> "seventeenth"
                8 -> "eighteenth"
                9 -> "nineteenth"
                else -> ""
            }
            2 -> "twent"
            3 -> "thirt"
            4 -> "fort"
            else -> ""
        }
        if (tens != 1) {
            alphabetic += when (ones) {
                1 -> "y-first"
                2 -> "y-second"
                3 -> "y-third"
                4 -> "y-fourth"
                5 -> "y-fifth"
                6 -> "y-sixth"
                7 -> "y-seventh"
                8 -> "y-eighth"
                9 -> "y-nineth"
                0 -> "ieth"
                else -> ""
            }
        }
        alphabetic += " task"
        return this.ansScanner.checkLine(alphabetic)
    }

    private fun case18(): Boolean {
        val number = this.datScanner.nextInt()
        val hundreds = number / 100
        val tens = number / 10 % 10
        val ones = number % 10
        var alphabetic = when (hundreds) {
            1 -> "one"
            2 -> "two"
            3 -> "three"
            4 -> "four"
            5 -> "five"
            6 -> "six"
            7 -> "seven"
            8 -> "eight"
            9 -> "nine"
            else -> ""
        }
        if (hundreds != 0) {
            alphabetic += " hundred"
        }
        if (tens != 0 || ones != 0) {
            alphabetic += " and "
        }
        alphabetic += when (tens) {
            1 -> when (ones) {
                0 -> "ten"
                1 -> "eleven"
                2 -> "twelve"
                3 -> "thirteen"
                4 -> "fourteen"
                5 -> "fifteen"
                6 -> "sixteen"
                7 -> "seventeen"
                8 -> "eighteen"
                9 -> "nineteen"
                else -> ""
            }
            2 -> "twenty"
            3 -> "thirty"
            4 -> "forty"
            5 -> "fifty"
            6 -> "sixty"
            7 -> "seventy"
            8 -> "eighty"
            9 -> "ninety"
            else -> ""
        }
        if (tens > 1 && ones != 0) {
            alphabetic += "-"
        }
        if (tens != 1) {
            alphabetic += when (ones) {
                1 -> "one"
                2 -> "two"
                3 -> "three"
                4 -> "four"
                5 -> "five"
                6 -> "six"
                7 -> "seven"
                8 -> "eight"
                9 -> "nine"
                else -> ""
            }
        }
        return this.ansScanner.checkLine(alphabetic)
    }

    private fun case19(): Boolean {
        var year = this.datScanner.nextInt()
        year -= 4
        if (year < 0) {
            year = -year
        }
        year %= 60
        // colors;
        val yearName = "The " + when {
            year < 12 -> "Green"
            year < 24 -> "Red"
            year < 36 -> "Yellow"
            year < 48 -> "White"
            else -> "Black"
        } + " " + when (year % 60 % 12) {
            0 -> "Rat"
            1 -> "Cow"
            2 -> "Tiger"
            3 -> "Hare"
            4 -> "Dragon"
            5 -> "Snake"
            6 -> "Horse"
            7 -> "Sheep"
            8 -> "Monkey"
            9 -> "Hen"
            10 -> "Dog"
            11 -> "Pig"
            else -> ""
        } + "'s year"
        return this.ansScanner.checkLine(yearName)
    }

    private fun case20(): Boolean {
        val d = this.datScanner.nextInt()
        val m = this.datScanner.nextInt()
        val zodiacalName = when (m) {
            1 -> when {
                d < 20 -> "Capricorn"
                else -> "Aquarius"
            }
            2 -> when {
                d < 19 -> "Aquarius"
                else -> "Pisces"
            }
            3 -> when {
                d < 21 -> "Pisces"
                else ->"Aries"
            }
            4 -> when {
                d < 20 -> "Aries"
                else -> "Taurus"
            }
            5 -> when {
                d < 21 -> "Taurus"
                else -> "Gemini"
            }
            6 -> when {
                d < 22 -> "Gemini"
                else -> "Cancer"
            }
            7 -> when {
                d < 23 -> "Cancer"
                else -> "Leo"
            }
            8 -> when {
                d < 23 -> "Leo"
                else -> "Virgo"
            }
            9 -> when {
                d < 23 -> "Virgo"
                else -> "Libra"
            }
            10 -> when {
                d < 23 -> "Libra"
                else -> "Scorpio"
            }
            11 -> when {
                d < 23 -> "Scorpio"
                else -> "Sagittarius"
            }
            12 -> when {
                d < 22 -> "Sagittarius"
                else -> "Capricorn"
            }
            else -> ""
        }
        return this.ansScanner.checkWord(zodiacalName)
    }
}