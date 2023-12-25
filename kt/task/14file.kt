package task

class FileTask {

    var datScanner = scanner.Scanner()
    var ansScanner = scanner.Scanner()

    fun solve(pathPrefix: String) {
        for (taskNo in 1..90) {
            var taskNoStr = "%03d".format(taskNo)
            var fileName: String
            for (testNo in 1..100) {
                var testNoStr = "%03d".format(testNo)
                fileName = "$pathPrefix/14file/File$taskNoStr/$testNoStr/$testNoStr"
                if (!datScanner.changeReader("$fileName.dat")) {
                    println("File$taskNoStr couldn't open datScanner source")
                    return
                }
                if (!ansScanner.changeReader("$fileName.ans")) {
                    println("File$taskNoStr couldn't open ansScanner source")
                    return
                }
                val ok = when (taskNo) {
                    1 -> this.file1()
                    2 -> this.file2()
                    3 -> this.file3()
                    4 -> this.file4()
                    5 -> this.file5()
                    6 -> this.file6()
                    7 -> this.file7()
                    8 -> this.file8()
                    9 -> this.file9()
                    10 -> this.file10()
                    11 -> this.file11()
                    12 -> this.file12()
                    13 -> this.file13()
                    14 -> this.file14()
                    15 -> this.file15()
                    16 -> this.file16()
                    17 -> this.file17()
                    18 -> this.file18()
                    19 -> this.file19()
                    20 -> this.file20()
                    21 -> this.file21()
                    22 -> this.file22()
                    23 -> this.file23()
                    24 -> this.file24()
                    25 -> this.file25()
                    26 -> this.file26()
                    27 -> this.file27()
                    28 -> this.file28()
                    29 -> this.file29()
                    30 -> this.file30()
                    31 -> this.file31()
                    32 -> this.file32()
                    33 -> this.file33()
                    34 -> this.file34()
                    35 -> this.file35()
                    36 -> this.file36()
                    37 -> this.file37()
                    38 -> this.file38()
                    39 -> this.file39()
                    40 -> this.file40()
                    41 -> this.file41()
                    42 -> this.file42()
                    43 -> this.file43()
                    44 -> this.file44()
                    45 -> this.file45()
                    46 -> this.file46()
                    47 -> this.file47()
                    48 -> this.file48()
                    49 -> this.file49()
                    50 -> this.file50()
                    51 -> this.file51()
                    52 -> this.file52()
                    53 -> this.file53()
                    54 -> this.file54()
                    55 -> this.file55()
                    56 -> this.file56()
                    57 -> this.file57()
                    58 -> this.file58()
                    59 -> this.file59()
                    60 -> this.file60()
                    61 -> this.file61()
                    62 -> this.file62()
                    63 -> this.file63()
                    64 -> this.file64()
                    65 -> this.file65()
                    66 -> this.file66()
                    67 -> this.file67()
                    68 -> this.file68()
                    69 -> this.file69()
                    70 -> this.file70()
                    71 -> this.file71()
                    72 -> this.file72()
                    73 -> this.file73()
                    74 -> this.file74()
                    75 -> this.file75()
                    76 -> this.file76()
                    77 -> this.file77()
                    78 -> this.file78()
                    79 -> this.file79()
                    80 -> this.file80()
                    81 -> this.file81()
                    82 -> this.file82()
                    83 -> this.file83()
                    84 -> this.file84()
                    85 -> this.file85()
                    86 -> this.file86()
                    87 -> this.file87()
                    88 -> this.file88()
                    89 -> this.file89()
                    90 -> this.file90()
                    else -> false
                }
                scanner.removeFiles()
                if (!ok) {
                    println("File$taskNoStr Test#$testNoStr has not been tested")
                    return
                }
                if (!datScanner.eof()) {
                    println("File$taskNoStr Test#$testNoStr input is not enough")
                    return
                }
                if (!ansScanner.eof()) {
                    println("File$taskNoStr Test#$testNoStr output is not enough")
                    return
                }
            }
            println("File$taskNoStr has been tested!")
        }
    }

    private fun file1(): Boolean {
        return false
    }

    private fun file2(): Boolean {
        return false
    }

    private fun file3(): Boolean {
        return false
    }

    private fun file4(): Boolean {
        return false
    }

    private fun file5(): Boolean {
        return false
    }

    private fun file6(): Boolean {
        return false
    }

    private fun file7(): Boolean {
        return false
    }

    private fun file8(): Boolean {
        return false
    }

    private fun file9(): Boolean {
        return false
    }

    private fun file10(): Boolean {
        return false
    }

    private fun file11(): Boolean {
        return false
    }

    private fun file12(): Boolean {
        return false
    }

    private fun file13(): Boolean {
        return false
    }

    private fun file14(): Boolean {
        return false
    }

    private fun file15(): Boolean {
        return false
    }

    private fun file16(): Boolean {
        return false
    }

    private fun file17(): Boolean {
        return false
    }

    private fun file18(): Boolean {
        return false
    }

    private fun file19(): Boolean {
        return false
    }

    private fun file20(): Boolean {
        return false
    }

    private fun file21(): Boolean {
        return false
    }

    private fun file22(): Boolean {
        return false
    }

    private fun file23(): Boolean {
        return false
    }

    private fun file24(): Boolean {
        return false
    }

    private fun file25(): Boolean {
        return false
    }

    private fun file26(): Boolean {
        return false
    }

    private fun file27(): Boolean {
        return false
    }

    private fun file28(): Boolean {
        return false
    }

    private fun file29(): Boolean {
        return false
    }

    private fun file30(): Boolean {
        return false
    }

    private fun file31(): Boolean {
        return false
    }

    private fun file32(): Boolean {
        return false
    }

    private fun file33(): Boolean {
        return false
    }

    private fun file34(): Boolean {
        return false
    }

    private fun file35(): Boolean {
        return false
    }

    private fun file36(): Boolean {
        return false
    }

    private fun file37(): Boolean {
        return false
    }

    private fun file38(): Boolean {
        return false
    }

    private fun file39(): Boolean {
        return false
    }

    private fun file40(): Boolean {
        return false
    }

    private fun file41(): Boolean {
        return false
    }

    private fun file42(): Boolean {
        return false
    }

    private fun file43(): Boolean {
        return false
    }

    private fun file44(): Boolean {
        return false
    }

    private fun file45(): Boolean {
        return false
    }

    private fun file46(): Boolean {
        return false
    }

    private fun file47(): Boolean {
        return false
    }

    private fun file48(): Boolean {
        return false
    }

    private fun file49(): Boolean {
        return false
    }

    private fun file50(): Boolean {
        return false
    }

    private fun file51(): Boolean {
        return false
    }

    private fun file52(): Boolean {
        return false
    }

    private fun file53(): Boolean {
        return false
    }

    private fun file54(): Boolean {
        return false
    }

    private fun file55(): Boolean {
        return false
    }

    private fun file56(): Boolean {
        return false
    }

    private fun file57(): Boolean {
        return false
    }

    private fun file58(): Boolean {
        return false
    }

    private fun file59(): Boolean {
        return false
    }

    private fun file60(): Boolean {
        return false
    }

    private fun file61(): Boolean {
        return false
    }

    private fun file62(): Boolean {
        return false
    }

    private fun file63(): Boolean {
        return false
    }

    private fun file64(): Boolean {
        return false
    }

    private fun file65(): Boolean {
        return false
    }

    private fun file66(): Boolean {
        return false
    }

    private fun file67(): Boolean {
        return false
    }

    private fun file68(): Boolean {
        return false
    }

    private fun file69(): Boolean {
        return false
    }

    private fun file70(): Boolean {
        return false
    }

    private fun file71(): Boolean {
        return false
    }

    private fun file72(): Boolean {
        return false
    }

    private fun file73(): Boolean {
        return false
    }

    private fun file74(): Boolean {
        return false
    }

    private fun file75(): Boolean {
        return false
    }

    private fun file76(): Boolean {
        return false
    }

    private fun file77(): Boolean {
        return false
    }

    private fun file78(): Boolean {
        return false
    }

    private fun file79(): Boolean {
        return false
    }

    private fun file80(): Boolean {
        return false
    }

    private fun file81(): Boolean {
        return false
    }

    private fun file82(): Boolean {
        return false
    }

    private fun file83(): Boolean {
        return false
    }

    private fun file84(): Boolean {
        return false
    }

    private fun file85(): Boolean {
        return false
    }

    private fun file86(): Boolean {
        return false
    }

    private fun file87(): Boolean {
        return false
    }

    private fun file88(): Boolean {
        return false
    }

    private fun file89(): Boolean {
        return false
    }

    private fun file90(): Boolean {
        return false
    }
}