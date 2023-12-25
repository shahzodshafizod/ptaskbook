package task

class ProcTask {

    var datScanner = scanner.Scanner()
    var ansScanner = scanner.Scanner()

    fun solve(pathPrefix: String) {
        for (taskNo in 1..60) {
            var taskNoStr = "%03d".format(taskNo)
            var fileName: String
            for (testNo in 1..100) {
                var testNoStr = "%03d".format(testNo)
                fileName = "$pathPrefix/09proc/Proc$taskNoStr/$testNoStr/$testNoStr"
                if (!datScanner.changeReader("$fileName.dat")) {
                    println("Proc$taskNoStr couldn't open datScanner source")
                    return
                }
                if (!ansScanner.changeReader("$fileName.ans")) {
                    println("Proc$taskNoStr couldn't open ansScanner source")
                    return
                }
                val ok = when (taskNo) {
                    1 -> this.proc1()
                    2 -> this.proc2()
                    3 -> this.proc3()
                    4 -> this.proc4()
                    5 -> this.proc5()
                    6 -> this.proc6()
                    7 -> this.proc7()
                    8 -> this.proc8()
                    9 -> this.proc9()
                    10 -> this.proc10()
                    11 -> this.proc11()
                    12 -> this.proc12()
                    13 -> this.proc13()
                    14 -> this.proc14()
                    15 -> this.proc15()
                    16 -> this.proc16()
                    17 -> this.proc17()
                    18 -> this.proc18()
                    19 -> this.proc19()
                    20 -> this.proc20()
                    21 -> this.proc21()
                    22 -> this.proc22()
                    23 -> this.proc23()
                    24 -> this.proc24()
                    25 -> this.proc25()
                    26 -> this.proc26()
                    27 -> this.proc27()
                    28 -> this.proc28()
                    29 -> this.proc29()
                    30 -> this.proc30()
                    31 -> this.proc31()
                    32 -> this.proc32()
                    33 -> this.proc33()
                    34 -> this.proc34()
                    35 -> this.proc35()
                    36 -> this.proc36()
                    37 -> this.proc37()
                    38 -> this.proc38()
                    39 -> this.proc39()
                    40 -> this.proc40()
                    41 -> this.proc41()
                    42 -> this.proc42()
                    43 -> this.proc43()
                    44 -> this.proc44()
                    45 -> this.proc45()
                    46 -> this.proc46()
                    47 -> this.proc47()
                    48 -> this.proc48()
                    49 -> this.proc49()
                    50 -> this.proc50()
                    51 -> this.proc51()
                    52 -> this.proc52()
                    53 -> this.proc53()
                    54 -> this.proc54()
                    55 -> this.proc55()
                    56 -> this.proc56()
                    57 -> this.proc57()
                    58 -> this.proc58()
                    59 -> this.proc59()
                    60 -> this.proc60()
                    else -> false
                }
                scanner.removeFiles()
                if (!ok) {
                    println("Proc$taskNoStr Test#$testNoStr has not been tested")
                    return
                }
                if (!datScanner.eof()) {
                    println("Proc$taskNoStr Test#$testNoStr input is not enough")
                    return
                }
                if (!ansScanner.eof()) {
                    println("Proc$taskNoStr Test#$testNoStr output is not enough")
                    return
                }
            }
            println("Proc$taskNoStr has been tested!")
        }
    }

    private fun proc1(): Boolean {
        return false
    }

    private fun proc2(): Boolean {
        return false
    }

    private fun proc3(): Boolean {
        return false
    }

    private fun proc4(): Boolean {
        return false
    }

    private fun proc5(): Boolean {
        return false
    }

    private fun proc6(): Boolean {
        return false
    }

    private fun proc7(): Boolean {
        return false
    }

    private fun proc8(): Boolean {
        return false
    }

    private fun proc9(): Boolean {
        return false
    }

    private fun proc10(): Boolean {
        return false
    }

    private fun proc11(): Boolean {
        return false
    }

    private fun proc12(): Boolean {
        return false
    }

    private fun proc13(): Boolean {
        return false
    }

    private fun proc14(): Boolean {
        return false
    }

    private fun proc15(): Boolean {
        return false
    }

    private fun proc16(): Boolean {
        return false
    }

    private fun proc17(): Boolean {
        return false
    }

    private fun proc18(): Boolean {
        return false
    }

    private fun proc19(): Boolean {
        return false
    }

    private fun proc20(): Boolean {
        return false
    }

    private fun proc21(): Boolean {
        return false
    }

    private fun proc22(): Boolean {
        return false
    }

    private fun proc23(): Boolean {
        return false
    }

    private fun proc24(): Boolean {
        return false
    }

    private fun proc25(): Boolean {
        return false
    }

    private fun proc26(): Boolean {
        return false
    }

    private fun proc27(): Boolean {
        return false
    }

    private fun proc28(): Boolean {
        return false
    }

    private fun proc29(): Boolean {
        return false
    }

    private fun proc30(): Boolean {
        return false
    }

    private fun proc31(): Boolean {
        return false
    }

    private fun proc32(): Boolean {
        return false
    }

    private fun proc33(): Boolean {
        return false
    }

    private fun proc34(): Boolean {
        return false
    }

    private fun proc35(): Boolean {
        return false
    }

    private fun proc36(): Boolean {
        return false
    }

    private fun proc37(): Boolean {
        return false
    }

    private fun proc38(): Boolean {
        return false
    }

    private fun proc39(): Boolean {
        return false
    }

    private fun proc40(): Boolean {
        return false
    }

    private fun proc41(): Boolean {
        return false
    }

    private fun proc42(): Boolean {
        return false
    }

    private fun proc43(): Boolean {
        return false
    }

    private fun proc44(): Boolean {
        return false
    }

    private fun proc45(): Boolean {
        return false
    }

    private fun proc46(): Boolean {
        return false
    }

    private fun proc47(): Boolean {
        return false
    }

    private fun proc48(): Boolean {
        return false
    }

    private fun proc49(): Boolean {
        return false
    }

    private fun proc50(): Boolean {
        return false
    }

    private fun proc51(): Boolean {
        return false
    }

    private fun proc52(): Boolean {
        return false
    }

    private fun proc53(): Boolean {
        return false
    }

    private fun proc54(): Boolean {
        return false
    }

    private fun proc55(): Boolean {
        return false
    }

    private fun proc56(): Boolean {
        return false
    }

    private fun proc57(): Boolean {
        return false
    }

    private fun proc58(): Boolean {
        return false
    }

    private fun proc59(): Boolean {
        return false
    }

    private fun proc60(): Boolean {
        return false
    }
}