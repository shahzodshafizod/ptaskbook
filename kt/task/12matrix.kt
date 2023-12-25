package task

class MatrixTask {

    var datScanner = scanner.Scanner()
    var ansScanner = scanner.Scanner()

    fun solve(pathPrefix: String) {
        for (taskNo in 1..100) {
            var taskNoStr = "%03d".format(taskNo)
            var fileName: String
            for (testNo in 1..100) {
                var testNoStr = "%03d".format(testNo)
                fileName = "$pathPrefix/12matrix/Matrix$taskNoStr/$testNoStr/$testNoStr"
                if (!datScanner.changeReader("$fileName.dat")) {
                    println("Matrix$taskNoStr couldn't open datScanner source")
                    return
                }
                if (!ansScanner.changeReader("$fileName.ans")) {
                    println("Matrix$taskNoStr couldn't open ansScanner source")
                    return
                }
                val ok = when (taskNo) {
                    1 -> this.matrix1()
                    2 -> this.matrix2()
                    3 -> this.matrix3()
                    4 -> this.matrix4()
                    5 -> this.matrix5()
                    6 -> this.matrix6()
                    7 -> this.matrix7()
                    8 -> this.matrix8()
                    9 -> this.matrix9()
                    10 -> this.matrix10()
                    11 -> this.matrix11()
                    12 -> this.matrix12()
                    13 -> this.matrix13()
                    14 -> this.matrix14()
                    15 -> this.matrix15()
                    16 -> this.matrix16()
                    17 -> this.matrix17()
                    18 -> this.matrix18()
                    19 -> this.matrix19()
                    20 -> this.matrix20()
                    21 -> this.matrix21()
                    22 -> this.matrix22()
                    23 -> this.matrix23()
                    24 -> this.matrix24()
                    25 -> this.matrix25()
                    26 -> this.matrix26()
                    27 -> this.matrix27()
                    28 -> this.matrix28()
                    29 -> this.matrix29()
                    30 -> this.matrix30()
                    31 -> this.matrix31()
                    32 -> this.matrix32()
                    33 -> this.matrix33()
                    34 -> this.matrix34()
                    35 -> this.matrix35()
                    36 -> this.matrix36()
                    37 -> this.matrix37()
                    38 -> this.matrix38()
                    39 -> this.matrix39()
                    40 -> this.matrix40()
                    41 -> this.matrix41()
                    42 -> this.matrix42()
                    43 -> this.matrix43()
                    44 -> this.matrix44()
                    45 -> this.matrix45()
                    46 -> this.matrix46()
                    47 -> this.matrix47()
                    48 -> this.matrix48()
                    49 -> this.matrix49()
                    50 -> this.matrix50()
                    51 -> this.matrix51()
                    52 -> this.matrix52()
                    53 -> this.matrix53()
                    54 -> this.matrix54()
                    55 -> this.matrix55()
                    56 -> this.matrix56()
                    57 -> this.matrix57()
                    58 -> this.matrix58()
                    59 -> this.matrix59()
                    60 -> this.matrix60()
                    61 -> this.matrix61()
                    62 -> this.matrix62()
                    63 -> this.matrix63()
                    64 -> this.matrix64()
                    65 -> this.matrix65()
                    66 -> this.matrix66()
                    67 -> this.matrix67()
                    68 -> this.matrix68()
                    69 -> this.matrix69()
                    70 -> this.matrix70()
                    71 -> this.matrix71()
                    72 -> this.matrix72()
                    73 -> this.matrix73()
                    74 -> this.matrix74()
                    75 -> this.matrix75()
                    76 -> this.matrix76()
                    77 -> this.matrix77()
                    78 -> this.matrix78()
                    79 -> this.matrix79()
                    80 -> this.matrix80()
                    81 -> this.matrix81()
                    82 -> this.matrix82()
                    83 -> this.matrix83()
                    84 -> this.matrix84()
                    85 -> this.matrix85()
                    86 -> this.matrix86()
                    87 -> this.matrix87()
                    88 -> this.matrix88()
                    89 -> this.matrix89()
                    90 -> this.matrix90()
                    91 -> this.matrix91()
                    92 -> this.matrix92()
                    93 -> this.matrix93()
                    94 -> this.matrix94()
                    95 -> this.matrix95()
                    96 -> this.matrix96()
                    97 -> this.matrix97()
                    98 -> this.matrix98()
                    99 -> this.matrix99()
                    100 -> this.matrix100()
                    else -> false
                }
                scanner.removeFiles()
                if (!ok) {
                    println("Matrix$taskNoStr Test#$testNoStr has not been tested")
                    return
                }
                if (!datScanner.eof()) {
                    println("Matrix$taskNoStr Test#$testNoStr input is not enough")
                    return
                }
                if (!ansScanner.eof()) {
                    println("Matrix$taskNoStr Test#$testNoStr output is not enough")
                    return
                }
            }
            println("Matrix$taskNoStr has been tested!")
        }
    }

    private fun matrix1(): Boolean {
        return false
    }

    private fun matrix2(): Boolean {
        return false
    }

    private fun matrix3(): Boolean {
        return false
    }

    private fun matrix4(): Boolean {
        return false
    }

    private fun matrix5(): Boolean {
        return false
    }

    private fun matrix6(): Boolean {
        return false
    }

    private fun matrix7(): Boolean {
        return false
    }

    private fun matrix8(): Boolean {
        return false
    }

    private fun matrix9(): Boolean {
        return false
    }

    private fun matrix10(): Boolean {
        return false
    }

    private fun matrix11(): Boolean {
        return false
    }

    private fun matrix12(): Boolean {
        return false
    }

    private fun matrix13(): Boolean {
        return false
    }

    private fun matrix14(): Boolean {
        return false
    }

    private fun matrix15(): Boolean {
        return false
    }

    private fun matrix16(): Boolean {
        return false
    }

    private fun matrix17(): Boolean {
        return false
    }

    private fun matrix18(): Boolean {
        return false
    }

    private fun matrix19(): Boolean {
        return false
    }

    private fun matrix20(): Boolean {
        return false
    }

    private fun matrix21(): Boolean {
        return false
    }

    private fun matrix22(): Boolean {
        return false
    }

    private fun matrix23(): Boolean {
        return false
    }

    private fun matrix24(): Boolean {
        return false
    }

    private fun matrix25(): Boolean {
        return false
    }

    private fun matrix26(): Boolean {
        return false
    }

    private fun matrix27(): Boolean {
        return false
    }

    private fun matrix28(): Boolean {
        return false
    }

    private fun matrix29(): Boolean {
        return false
    }

    private fun matrix30(): Boolean {
        return false
    }

    private fun matrix31(): Boolean {
        return false
    }

    private fun matrix32(): Boolean {
        return false
    }

    private fun matrix33(): Boolean {
        return false
    }

    private fun matrix34(): Boolean {
        return false
    }

    private fun matrix35(): Boolean {
        return false
    }

    private fun matrix36(): Boolean {
        return false
    }

    private fun matrix37(): Boolean {
        return false
    }

    private fun matrix38(): Boolean {
        return false
    }

    private fun matrix39(): Boolean {
        return false
    }

    private fun matrix40(): Boolean {
        return false
    }

    private fun matrix41(): Boolean {
        return false
    }

    private fun matrix42(): Boolean {
        return false
    }

    private fun matrix43(): Boolean {
        return false
    }

    private fun matrix44(): Boolean {
        return false
    }

    private fun matrix45(): Boolean {
        return false
    }

    private fun matrix46(): Boolean {
        return false
    }

    private fun matrix47(): Boolean {
        return false
    }

    private fun matrix48(): Boolean {
        return false
    }

    private fun matrix49(): Boolean {
        return false
    }

    private fun matrix50(): Boolean {
        return false
    }

    private fun matrix51(): Boolean {
        return false
    }

    private fun matrix52(): Boolean {
        return false
    }

    private fun matrix53(): Boolean {
        return false
    }

    private fun matrix54(): Boolean {
        return false
    }

    private fun matrix55(): Boolean {
        return false
    }

    private fun matrix56(): Boolean {
        return false
    }

    private fun matrix57(): Boolean {
        return false
    }

    private fun matrix58(): Boolean {
        return false
    }

    private fun matrix59(): Boolean {
        return false
    }

    private fun matrix60(): Boolean {
        return false
    }

    private fun matrix61(): Boolean {
        return false
    }

    private fun matrix62(): Boolean {
        return false
    }

    private fun matrix63(): Boolean {
        return false
    }

    private fun matrix64(): Boolean {
        return false
    }

    private fun matrix65(): Boolean {
        return false
    }

    private fun matrix66(): Boolean {
        return false
    }

    private fun matrix67(): Boolean {
        return false
    }

    private fun matrix68(): Boolean {
        return false
    }

    private fun matrix69(): Boolean {
        return false
    }

    private fun matrix70(): Boolean {
        return false
    }

    private fun matrix71(): Boolean {
        return false
    }

    private fun matrix72(): Boolean {
        return false
    }

    private fun matrix73(): Boolean {
        return false
    }

    private fun matrix74(): Boolean {
        return false
    }

    private fun matrix75(): Boolean {
        return false
    }

    private fun matrix76(): Boolean {
        return false
    }

    private fun matrix77(): Boolean {
        return false
    }

    private fun matrix78(): Boolean {
        return false
    }

    private fun matrix79(): Boolean {
        return false
    }

    private fun matrix80(): Boolean {
        return false
    }

    private fun matrix81(): Boolean {
        return false
    }

    private fun matrix82(): Boolean {
        return false
    }

    private fun matrix83(): Boolean {
        return false
    }

    private fun matrix84(): Boolean {
        return false
    }

    private fun matrix85(): Boolean {
        return false
    }

    private fun matrix86(): Boolean {
        return false
    }

    private fun matrix87(): Boolean {
        return false
    }

    private fun matrix88(): Boolean {
        return false
    }

    private fun matrix89(): Boolean {
        return false
    }

    private fun matrix90(): Boolean {
        return false
    }

    private fun matrix91(): Boolean {
        return false
    }

    private fun matrix92(): Boolean {
        return false
    }

    private fun matrix93(): Boolean {
        return false
    }

    private fun matrix94(): Boolean {
        return false
    }

    private fun matrix95(): Boolean {
        return false
    }

    private fun matrix96(): Boolean {
        return false
    }

    private fun matrix97(): Boolean {
        return false
    }

    private fun matrix98(): Boolean {
        return false
    }

    private fun matrix99(): Boolean {
        return false
    }

    private fun matrix100(): Boolean {
        return false
    }
}