package task

class RecurTask {

    var datScanner = scanner.Scanner()
    var ansScanner = scanner.Scanner()

    fun solve(pathPrefix: String) {
        for (taskNo in 1..30) {
            var taskNoStr = "%03d".format(taskNo)
            var fileName: String
            for (testNo in 1..100) {
                var testNoStr = "%03d".format(testNo)
                fileName = "$pathPrefix/17recur/Recur$taskNoStr/$testNoStr/$testNoStr"
                if (!datScanner.changeReader("$fileName.dat")) {
                    println("Recur$taskNoStr couldn't open datScanner source")
                    return
                }
                if (!ansScanner.changeReader("$fileName.ans")) {
                    println("Recur$taskNoStr couldn't open ansScanner source")
                    return
                }
                val ok = when (taskNo) {
                    1 -> this.recur1()
                    2 -> this.recur2()
                    3 -> this.recur3()
                    4 -> this.recur4()
                    5 -> this.recur5()
                    6 -> this.recur6()
                    7 -> this.recur7()
                    8 -> this.recur8()
                    9 -> this.recur9()
                    10 -> this.recur10()
                    11 -> this.recur11()
                    12 -> this.recur12()
                    13 -> this.recur13()
                    14 -> this.recur14()
                    15 -> this.recur15()
                    16 -> this.recur16()
                    17 -> this.recur17()
                    18 -> this.recur18()
                    19 -> this.recur19()
                    20 -> this.recur20()
                    21 -> this.recur21()
                    22 -> this.recur22()
                    23 -> this.recur23()
                    24 -> this.recur24()
                    25 -> this.recur25()
                    26 -> this.recur26()
                    27 -> this.recur27()
                    28 -> this.recur28()
                    29 -> this.recur29()
                    30 -> this.recur30()
                    else -> false
                }
                scanner.removeFiles()
                if (!ok) {
                    println("Recur$taskNoStr Test#$testNoStr has not been tested")
                    return
                }
                if (!datScanner.eof()) {
                    println("Recur$taskNoStr Test#$testNoStr input is not enough")
                    return
                }
                if (!ansScanner.eof()) {
                    println("Recur$taskNoStr Test#$testNoStr output is not enough")
                    return
                }
            }
            println("Recur$taskNoStr has been tested!")
        }
    }

    private fun recur1(): Boolean {
        return false
    }

    private fun recur2(): Boolean {
        return false
    }

    private fun recur3(): Boolean {
        return false
    }

    private fun recur4(): Boolean {
        return false
    }

    private fun recur5(): Boolean {
        return false
    }

    private fun recur6(): Boolean {
        return false
    }

    private fun recur7(): Boolean {
        return false
    }

    private fun recur8(): Boolean {
        return false
    }

    private fun recur9(): Boolean {
        return false
    }

    private fun recur10(): Boolean {
        return false
    }

    private fun recur11(): Boolean {
        return false
    }

    private fun recur12(): Boolean {
        return false
    }

    private fun recur13(): Boolean {
        return false
    }

    private fun recur14(): Boolean {
        return false
    }

    private fun recur15(): Boolean {
        return false
    }

    private fun recur16(): Boolean {
        return false
    }

    private fun recur17(): Boolean {
        return false
    }

    private fun recur18(): Boolean {
        return false
    }

    private fun recur19(): Boolean {
        return false
    }

    private fun recur20(): Boolean {
        return false
    }

    private fun recur21(): Boolean {
        return false
    }

    private fun recur22(): Boolean {
        return false
    }

    private fun recur23(): Boolean {
        return false
    }

    private fun recur24(): Boolean {
        return false
    }

    private fun recur25(): Boolean {
        return false
    }

    private fun recur26(): Boolean {
        return false
    }

    private fun recur27(): Boolean {
        return false
    }

    private fun recur28(): Boolean {
        return false
    }

    private fun recur29(): Boolean {
        return false
    }

    private fun recur30(): Boolean {
        return false
    }
}