package task

class MinmaxTask {

    var datScanner = scanner.Scanner()
    var ansScanner = scanner.Scanner()

    fun solve(pathPrefix: String) {
        for (taskNo in 1..30) {
            var taskNoStr = "%03d".format(taskNo)
            var fileName: String
            for (testNo in 1..100) {
                var testNoStr = "%03d".format(testNo)
                fileName = "$pathPrefix/10minmax/Minmax$taskNoStr/$testNoStr/$testNoStr"
                if (!datScanner.changeReader("$fileName.dat")) {
                    println("MinMax$taskNoStr couldn't open datScanner source")
                    return
                }
                if (!ansScanner.changeReader("$fileName.ans")) {
                    println("MinMax$taskNoStr couldn't open ansScanner source")
                    return
                }
                val ok = when (taskNo) {
                    1 -> this.minmax1()
                    2 -> this.minmax2()
                    3 -> this.minmax3()
                    4 -> this.minmax4()
                    5 -> this.minmax5()
                    6 -> this.minmax6()
                    7 -> this.minmax7()
                    8 -> this.minmax8()
                    9 -> this.minmax9()
                    10 -> this.minmax10()
                    11 -> this.minmax11()
                    12 -> this.minmax12()
                    13 -> this.minmax13()
                    14 -> this.minmax14()
                    15 -> this.minmax15()
                    16 -> this.minmax16()
                    17 -> this.minmax17()
                    18 -> this.minmax18()
                    19 -> this.minmax19()
                    20 -> this.minmax20()
                    21 -> this.minmax21()
                    22 -> this.minmax22()
                    23 -> this.minmax23()
                    24 -> this.minmax24()
                    25 -> this.minmax25()
                    26 -> this.minmax26()
                    27 -> this.minmax27()
                    28 -> this.minmax28()
                    29 -> this.minmax29()
                    30 -> this.minmax30()
                    else -> false
                }
                scanner.removeFiles()
                if (!ok) {
                    println("Minmax$taskNoStr Test#$testNoStr has not been tested")
                    return
                }
                if (!datScanner.eof()) {
                    println("Minmax$taskNoStr Test#$testNoStr input is not enough")
                    return
                }
                if (!ansScanner.eof()) {
                    println("Minmax$taskNoStr Test#$testNoStr output is not enough")
                    return
                }
            }
            println("Minmax$taskNoStr has been tested!")
        }
    }

    private fun minmax1(): Boolean {
        return false
    }

    private fun minmax2(): Boolean {
        return false
    }

    private fun minmax3(): Boolean {
        return false
    }

    private fun minmax4(): Boolean {
        return false
    }

    private fun minmax5(): Boolean {
        return false
    }

    private fun minmax6(): Boolean {
        return false
    }

    private fun minmax7(): Boolean {
        return false
    }

    private fun minmax8(): Boolean {
        return false
    }

    private fun minmax9(): Boolean {
        return false
    }

    private fun minmax10(): Boolean {
        return false
    }

    private fun minmax11(): Boolean {
        return false
    }

    private fun minmax12(): Boolean {
        return false
    }

    private fun minmax13(): Boolean {
        return false
    }

    private fun minmax14(): Boolean {
        return false
    }

    private fun minmax15(): Boolean {
        return false
    }

    private fun minmax16(): Boolean {
        return false
    }

    private fun minmax17(): Boolean {
        return false
    }

    private fun minmax18(): Boolean {
        return false
    }

    private fun minmax19(): Boolean {
        return false
    }

    private fun minmax20(): Boolean {
        return false
    }

    private fun minmax21(): Boolean {
        return false
    }

    private fun minmax22(): Boolean {
        return false
    }

    private fun minmax23(): Boolean {
        return false
    }

    private fun minmax24(): Boolean {
        return false
    }

    private fun minmax25(): Boolean {
        return false
    }

    private fun minmax26(): Boolean {
        return false
    }

    private fun minmax27(): Boolean {
        return false
    }

    private fun minmax28(): Boolean {
        return false
    }

    private fun minmax29(): Boolean {
        return false
    }

    private fun minmax30(): Boolean {
        return false
    }
}