package task

class StringTask {

    var datScanner = scanner.Scanner()
    var ansScanner = scanner.Scanner()

    fun solve(pathPrefix: String) {
        for (taskNo in 1..70) {
            var taskNoStr = "%03d".format(taskNo)
            var fileName: String
            for (testNo in 1..100) {
                var testNoStr = "%03d".format(testNo)
                fileName = "$pathPrefix/13string/String$taskNoStr/$testNoStr/$testNoStr"
                if (!datScanner.changeReader("$fileName.dat")) {
                    println("String$taskNoStr couldn't open datScanner source")
                    return
                }
                if (!ansScanner.changeReader("$fileName.ans")) {
                    println("String$taskNoStr couldn't open ansScanner source")
                    return
                }
                val ok = when (taskNo) {
                    1 -> this.string1()
                    2 -> this.string2()
                    3 -> this.string3()
                    4 -> this.string4()
                    5 -> this.string5()
                    6 -> this.string6()
                    7 -> this.string7()
                    8 -> this.string8()
                    9 -> this.string9()
                    10 -> this.string10()
                    11 -> this.string11()
                    12 -> this.string12()
                    13 -> this.string13()
                    14 -> this.string14()
                    15 -> this.string15()
                    16 -> this.string16()
                    17 -> this.string17()
                    18 -> this.string18()
                    19 -> this.string19()
                    20 -> this.string20()
                    21 -> this.string21()
                    22 -> this.string22()
                    23 -> this.string23()
                    24 -> this.string24()
                    25 -> this.string25()
                    26 -> this.string26()
                    27 -> this.string27()
                    28 -> this.string28()
                    29 -> this.string29()
                    30 -> this.string30()
                    31 -> this.string31()
                    32 -> this.string32()
                    33 -> this.string33()
                    34 -> this.string34()
                    35 -> this.string35()
                    36 -> this.string36()
                    37 -> this.string37()
                    38 -> this.string38()
                    39 -> this.string39()
                    40 -> this.string40()
                    41 -> this.string41()
                    42 -> this.string42()
                    43 -> this.string43()
                    44 -> this.string44()
                    45 -> this.string45()
                    46 -> this.string46()
                    47 -> this.string47()
                    48 -> this.string48()
                    49 -> this.string49()
                    50 -> this.string50()
                    51 -> this.string51()
                    52 -> this.string52()
                    53 -> this.string53()
                    54 -> this.string54()
                    55 -> this.string55()
                    56 -> this.string56()
                    57 -> this.string57()
                    58 -> this.string58()
                    59 -> this.string59()
                    60 -> this.string60()
                    61 -> this.string61()
                    62 -> this.string62()
                    63 -> this.string63()
                    64 -> this.string64()
                    65 -> this.string65()
                    66 -> this.string66()
                    67 -> this.string67()
                    68 -> this.string68()
                    69 -> this.string69()
                    70 -> this.string70()
                    else -> false
                }
                scanner.removeFiles()
                if (!ok) {
                    println("String$taskNoStr Test#$testNoStr has not been tested")
                    return
                }
                if (!datScanner.eof()) {
                    println("String$taskNoStr Test#$testNoStr input is not enough")
                    return
                }
                if (!ansScanner.eof()) {
                    println("String$taskNoStr Test#$testNoStr output is not enough")
                    return
                }
            }
            println("String$taskNoStr has been tested!")
        }
    }

    private fun string1(): Boolean {
        return false
    }

    private fun string2(): Boolean {
        return false
    }

    private fun string3(): Boolean {
        return false
    }

    private fun string4(): Boolean {
        return false
    }

    private fun string5(): Boolean {
        return false
    }

    private fun string6(): Boolean {
        return false
    }

    private fun string7(): Boolean {
        return false
    }

    private fun string8(): Boolean {
        return false
    }

    private fun string9(): Boolean {
        return false
    }

    private fun string10(): Boolean {
        return false
    }

    private fun string11(): Boolean {
        return false
    }

    private fun string12(): Boolean {
        return false
    }

    private fun string13(): Boolean {
        return false
    }

    private fun string14(): Boolean {
        return false
    }

    private fun string15(): Boolean {
        return false
    }

    private fun string16(): Boolean {
        return false
    }

    private fun string17(): Boolean {
        return false
    }

    private fun string18(): Boolean {
        return false
    }

    private fun string19(): Boolean {
        return false
    }

    private fun string20(): Boolean {
        return false
    }

    private fun string21(): Boolean {
        return false
    }

    private fun string22(): Boolean {
        return false
    }

    private fun string23(): Boolean {
        return false
    }

    private fun string24(): Boolean {
        return false
    }

    private fun string25(): Boolean {
        return false
    }

    private fun string26(): Boolean {
        return false
    }

    private fun string27(): Boolean {
        return false
    }

    private fun string28(): Boolean {
        return false
    }

    private fun string29(): Boolean {
        return false
    }

    private fun string30(): Boolean {
        return false
    }

    private fun string31(): Boolean {
        return false
    }

    private fun string32(): Boolean {
        return false
    }

    private fun string33(): Boolean {
        return false
    }

    private fun string34(): Boolean {
        return false
    }

    private fun string35(): Boolean {
        return false
    }

    private fun string36(): Boolean {
        return false
    }

    private fun string37(): Boolean {
        return false
    }

    private fun string38(): Boolean {
        return false
    }

    private fun string39(): Boolean {
        return false
    }

    private fun string40(): Boolean {
        return false
    }

    private fun string41(): Boolean {
        return false
    }

    private fun string42(): Boolean {
        return false
    }

    private fun string43(): Boolean {
        return false
    }

    private fun string44(): Boolean {
        return false
    }

    private fun string45(): Boolean {
        return false
    }

    private fun string46(): Boolean {
        return false
    }

    private fun string47(): Boolean {
        return false
    }

    private fun string48(): Boolean {
        return false
    }

    private fun string49(): Boolean {
        return false
    }

    private fun string50(): Boolean {
        return false
    }

    private fun string51(): Boolean {
        return false
    }

    private fun string52(): Boolean {
        return false
    }

    private fun string53(): Boolean {
        return false
    }

    private fun string54(): Boolean {
        return false
    }

    private fun string55(): Boolean {
        return false
    }

    private fun string56(): Boolean {
        return false
    }

    private fun string57(): Boolean {
        return false
    }

    private fun string58(): Boolean {
        return false
    }

    private fun string59(): Boolean {
        return false
    }

    private fun string60(): Boolean {
        return false
    }

    private fun string61(): Boolean {
        return false
    }

    private fun string62(): Boolean {
        return false
    }

    private fun string63(): Boolean {
        return false
    }

    private fun string64(): Boolean {
        return false
    }

    private fun string65(): Boolean {
        return false
    }

    private fun string66(): Boolean {
        return false
    }

    private fun string67(): Boolean {
        return false
    }

    private fun string68(): Boolean {
        return false
    }

    private fun string69(): Boolean {
        return false
    }

    private fun string70(): Boolean {
        return false
    }
}