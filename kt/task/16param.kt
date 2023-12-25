package task

class ParamTask {

    var datScanner = scanner.Scanner()
    var ansScanner = scanner.Scanner()

    fun solve(pathPrefix: String) {
        for (taskNo in 1..70) {
            var taskNoStr = "%03d".format(taskNo)
            var fileName: String
            for (testNo in 1..100) {
                var testNoStr = "%03d".format(testNo)
                fileName = "$pathPrefix/16param/Param$taskNoStr/$testNoStr/$testNoStr"
                if (!datScanner.changeReader("$fileName.dat")) {
                    println("Param$taskNoStr couldn't open datScanner source")
                    return
                }
                if (!ansScanner.changeReader("$fileName.ans")) {
                    println("Param$taskNoStr couldn't open ansScanner source")
                    return
                }
                val ok = when (taskNo) {
                    1 -> this.param1()
                    2 -> this.param2()
                    3 -> this.param3()
                    4 -> this.param4()
                    5 -> this.param5()
                    6 -> this.param6()
                    7 -> this.param7()
                    8 -> this.param8()
                    9 -> this.param9()
                    10 -> this.param10()
                    11 -> this.param11()
                    12 -> this.param12()
                    13 -> this.param13()
                    14 -> this.param14()
                    15 -> this.param15()
                    16 -> this.param16()
                    17 -> this.param17()
                    18 -> this.param18()
                    19 -> this.param19()
                    20 -> this.param20()
                    21 -> this.param21()
                    22 -> this.param22()
                    23 -> this.param23()
                    24 -> this.param24()
                    25 -> this.param25()
                    26 -> this.param26()
                    27 -> this.param27()
                    28 -> this.param28()
                    29 -> this.param29()
                    30 -> this.param30()
                    31 -> this.param31()
                    32 -> this.param32()
                    33 -> this.param33()
                    34 -> this.param34()
                    35 -> this.param35()
                    36 -> this.param36()
                    37 -> this.param37()
                    38 -> this.param38()
                    39 -> this.param39()
                    40 -> this.param40()
                    41 -> this.param41()
                    42 -> this.param42()
                    43 -> this.param43()
                    44 -> this.param44()
                    45 -> this.param45()
                    46 -> this.param46()
                    47 -> this.param47()
                    48 -> this.param48()
                    49 -> this.param49()
                    50 -> this.param50()
                    51 -> this.param51()
                    52 -> this.param52()
                    53 -> this.param53()
                    54 -> this.param54()
                    55 -> this.param55()
                    56 -> this.param56()
                    57 -> this.param57()
                    58 -> this.param58()
                    59 -> this.param59()
                    60 -> this.param60()
                    61 -> this.param61()
                    62 -> this.param62()
                    63 -> this.param63()
                    64 -> this.param64()
                    65 -> this.param65()
                    66 -> this.param66()
                    67 -> this.param67()
                    68 -> this.param68()
                    69 -> this.param69()
                    70 -> this.param70()
                    else -> false
                }
                scanner.removeFiles()
                if (!ok) {
                    println("Param$taskNoStr Test#$testNoStr has not been tested")
                    return
                }
                if (!datScanner.eof()) {
                    println("Param$taskNoStr Test#$testNoStr input is not enough")
                    return
                }
                if (!ansScanner.eof()) {
                    println("Param$taskNoStr Test#$testNoStr output is not enough")
                    return
                }
            }
            println("Param$taskNoStr has been tested!")
        }
    }

    private fun param1(): Boolean {
        return false
    }

    private fun param2(): Boolean {
        return false
    }

    private fun param3(): Boolean {
        return false
    }

    private fun param4(): Boolean {
        return false
    }

    private fun param5(): Boolean {
        return false
    }

    private fun param6(): Boolean {
        return false
    }

    private fun param7(): Boolean {
        return false
    }

    private fun param8(): Boolean {
        return false
    }

    private fun param9(): Boolean {
        return false
    }

    private fun param10(): Boolean {
        return false
    }

    private fun param11(): Boolean {
        return false
    }

    private fun param12(): Boolean {
        return false
    }

    private fun param13(): Boolean {
        return false
    }

    private fun param14(): Boolean {
        return false
    }

    private fun param15(): Boolean {
        return false
    }

    private fun param16(): Boolean {
        return false
    }

    private fun param17(): Boolean {
        return false
    }

    private fun param18(): Boolean {
        return false
    }

    private fun param19(): Boolean {
        return false
    }

    private fun param20(): Boolean {
        return false
    }

    private fun param21(): Boolean {
        return false
    }

    private fun param22(): Boolean {
        return false
    }

    private fun param23(): Boolean {
        return false
    }

    private fun param24(): Boolean {
        return false
    }

    private fun param25(): Boolean {
        return false
    }

    private fun param26(): Boolean {
        return false
    }

    private fun param27(): Boolean {
        return false
    }

    private fun param28(): Boolean {
        return false
    }

    private fun param29(): Boolean {
        return false
    }

    private fun param30(): Boolean {
        return false
    }

    private fun param31(): Boolean {
        return false
    }

    private fun param32(): Boolean {
        return false
    }

    private fun param33(): Boolean {
        return false
    }

    private fun param34(): Boolean {
        return false
    }

    private fun param35(): Boolean {
        return false
    }

    private fun param36(): Boolean {
        return false
    }

    private fun param37(): Boolean {
        return false
    }

    private fun param38(): Boolean {
        return false
    }

    private fun param39(): Boolean {
        return false
    }

    private fun param40(): Boolean {
        return false
    }

    private fun param41(): Boolean {
        return false
    }

    private fun param42(): Boolean {
        return false
    }

    private fun param43(): Boolean {
        return false
    }

    private fun param44(): Boolean {
        return false
    }

    private fun param45(): Boolean {
        return false
    }

    private fun param46(): Boolean {
        return false
    }

    private fun param47(): Boolean {
        return false
    }

    private fun param48(): Boolean {
        return false
    }

    private fun param49(): Boolean {
        return false
    }

    private fun param50(): Boolean {
        return false
    }

    private fun param51(): Boolean {
        return false
    }

    private fun param52(): Boolean {
        return false
    }

    private fun param53(): Boolean {
        return false
    }

    private fun param54(): Boolean {
        return false
    }

    private fun param55(): Boolean {
        return false
    }

    private fun param56(): Boolean {
        return false
    }

    private fun param57(): Boolean {
        return false
    }

    private fun param58(): Boolean {
        return false
    }

    private fun param59(): Boolean {
        return false
    }

    private fun param60(): Boolean {
        return false
    }

    private fun param61(): Boolean {
        return false
    }

    private fun param62(): Boolean {
        return false
    }

    private fun param63(): Boolean {
        return false
    }

    private fun param64(): Boolean {
        return false
    }

    private fun param65(): Boolean {
        return false
    }

    private fun param66(): Boolean {
        return false
    }

    private fun param67(): Boolean {
        return false
    }

    private fun param68(): Boolean {
        return false
    }

    private fun param69(): Boolean {
        return false
    }

    private fun param70(): Boolean {
        return false
    }
}