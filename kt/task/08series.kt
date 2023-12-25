package task

class SeriesTask {

    var datScanner = scanner.Scanner()
    var ansScanner = scanner.Scanner()

    fun solve(pathPrefix: String) {
        for (taskNo in 1..40) {
            var taskNoStr = "%03d".format(taskNo)
            var fileName: String
            for (testNo in 1..100) {
                var testNoStr = "%03d".format(testNo)
                fileName = "$pathPrefix/08series/Series$taskNoStr/$testNoStr/$testNoStr"
                if (!datScanner.changeReader("$fileName.dat")) {
                    println("Series$taskNoStr couldn't open datScanner source")
                    return
                }
                if (!ansScanner.changeReader("$fileName.ans")) {
                    println("Series$taskNoStr couldn't open ansScanner source")
                    return
                }
                val ok = when (taskNo) {
                    1 -> this.series1()
                    2 -> this.series2()
                    3 -> this.series3()
                    4 -> this.series4()
                    5 -> this.series5()
                    6 -> this.series6()
                    7 -> this.series7()
                    8 -> this.series8()
                    9 -> this.series9()
                    10 -> this.series10()
                    11 -> this.series11()
                    12 -> this.series12()
                    13 -> this.series13()
                    14 -> this.series14()
                    15 -> this.series15()
                    16 -> this.series16()
                    17 -> this.series17()
                    18 -> this.series18()
                    19 -> this.series19()
                    20 -> this.series20()
                    21 -> this.series21()
                    22 -> this.series22()
                    23 -> this.series23()
                    24 -> this.series24()
                    25 -> this.series25()
                    26 -> this.series26()
                    27 -> this.series27()
                    28 -> this.series28()
                    29 -> this.series29()
                    30 -> this.series30()
                    31 -> this.series31()
                    32 -> this.series32()
                    33 -> this.series33()
                    34 -> this.series34()
                    35 -> this.series35()
                    36 -> this.series36()
                    37 -> this.series37()
                    38 -> this.series38()
                    39 -> this.series39()
                    40 -> this.series40()
                    else -> false
                }
                scanner.removeFiles()
                if (!ok) {
                    println("Series$taskNoStr Test#$testNoStr has not been tested")
                    return
                }
                if (!datScanner.eof()) {
                    println("Series$taskNoStr Test#$testNoStr input is not enough")
                    return
                }
                if (!ansScanner.eof()) {
                    println("Series$taskNoStr Test#$testNoStr output is not enough")
                    return
                }
            }
            println("Series$taskNoStr has been tested!")
        }
    }

    private fun series1(): Boolean {
        return false
    }

    private fun series2(): Boolean {
        return false
    }

    private fun series3(): Boolean {
        return false
    }

    private fun series4(): Boolean {
        return false
    }

    private fun series5(): Boolean {
        return false
    }

    private fun series6(): Boolean {
        return false
    }

    private fun series7(): Boolean {
        return false
    }

    private fun series8(): Boolean {
        return false
    }

    private fun series9(): Boolean {
        return false
    }

    private fun series10(): Boolean {
        return false
    }

    private fun series11(): Boolean {
        return false
    }

    private fun series12(): Boolean {
        return false
    }

    private fun series13(): Boolean {
        return false
    }

    private fun series14(): Boolean {
        return false
    }

    private fun series15(): Boolean {
        return false
    }

    private fun series16(): Boolean {
        return false
    }

    private fun series17(): Boolean {
        return false
    }

    private fun series18(): Boolean {
        return false
    }

    private fun series19(): Boolean {
        return false
    }

    private fun series20(): Boolean {
        return false
    }

    private fun series21(): Boolean {
        return false
    }

    private fun series22(): Boolean {
        return false
    }

    private fun series23(): Boolean {
        return false
    }

    private fun series24(): Boolean {
        return false
    }

    private fun series25(): Boolean {
        return false
    }

    private fun series26(): Boolean {
        return false
    }

    private fun series27(): Boolean {
        return false
    }

    private fun series28(): Boolean {
        return false
    }

    private fun series29(): Boolean {
        return false
    }

    private fun series30(): Boolean {
        return false
    }

    private fun series31(): Boolean {
        return false
    }

    private fun series32(): Boolean {
        return false
    }

    private fun series33(): Boolean {
        return false
    }

    private fun series34(): Boolean {
        return false
    }

    private fun series35(): Boolean {
        return false
    }

    private fun series36(): Boolean {
        return false
    }

    private fun series37(): Boolean {
        return false
    }

    private fun series38(): Boolean {
        return false
    }

    private fun series39(): Boolean {
        return false
    }

    private fun series40(): Boolean {
        return false
    }
}