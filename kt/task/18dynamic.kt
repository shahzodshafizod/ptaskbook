package task

class DynamicTask {

    var datScanner = scanner.Scanner()
    var ansScanner = scanner.Scanner()

    fun solve(pathPrefix: String) {
        for (taskNo in 1..80) {
            var taskNoStr = "%03d".format(taskNo)
            var fileName: String
            for (testNo in 1..100) {
                var testNoStr = "%03d".format(testNo)
                fileName = "$pathPrefix/18dynamic/Dynamic$taskNoStr/$testNoStr/$testNoStr"
                if (!datScanner.changeReader("$fileName.dat")) {
                    println("Dynamic$taskNoStr couldn't open datScanner source")
                    return
                }
                if (!ansScanner.changeReader("$fileName.ans")) {
                    println("Dynamic$taskNoStr couldn't open ansScanner source")
                    return
                }
                val ok = when (taskNo) {
                    1 -> this.dynamic1()
                    2 -> this.dynamic2()
                    3 -> this.dynamic3()
                    4 -> this.dynamic4()
                    5 -> this.dynamic5()
                    6 -> this.dynamic6()
                    7 -> this.dynamic7()
                    8 -> this.dynamic8()
                    9 -> this.dynamic9()
                    10 -> this.dynamic10()
                    11 -> this.dynamic11()
                    12 -> this.dynamic12()
                    13 -> this.dynamic13()
                    14 -> this.dynamic14()
                    15 -> this.dynamic15()
                    16 -> this.dynamic16()
                    17 -> this.dynamic17()
                    18 -> this.dynamic18()
                    19 -> this.dynamic19()
                    20 -> this.dynamic20()
                    21 -> this.dynamic21()
                    22 -> this.dynamic22()
                    23 -> this.dynamic23()
                    24 -> this.dynamic24()
                    25 -> this.dynamic25()
                    26 -> this.dynamic26()
                    27 -> this.dynamic27()
                    28 -> this.dynamic28()
                    29 -> this.dynamic29()
                    30 -> this.dynamic30()
                    31 -> this.dynamic31()
                    32 -> this.dynamic32()
                    33 -> this.dynamic33()
                    34 -> this.dynamic34()
                    35 -> this.dynamic35()
                    36 -> this.dynamic36()
                    37 -> this.dynamic37()
                    38 -> this.dynamic38()
                    39 -> this.dynamic39()
                    40 -> this.dynamic40()
                    41 -> this.dynamic41()
                    42 -> this.dynamic42()
                    43 -> this.dynamic43()
                    44 -> this.dynamic44()
                    45 -> this.dynamic45()
                    46 -> this.dynamic46()
                    47 -> this.dynamic47()
                    48 -> this.dynamic48()
                    49 -> this.dynamic49()
                    50 -> this.dynamic50()
                    51 -> this.dynamic51()
                    52 -> this.dynamic52()
                    53 -> this.dynamic53()
                    54 -> this.dynamic54()
                    55 -> this.dynamic55()
                    56 -> this.dynamic56()
                    57 -> this.dynamic57()
                    58 -> this.dynamic58()
                    59 -> this.dynamic59()
                    60 -> this.dynamic60()
                    61 -> this.dynamic61()
                    62 -> this.dynamic62()
                    63 -> this.dynamic63()
                    64 -> this.dynamic64()
                    65 -> this.dynamic65()
                    66 -> this.dynamic66()
                    67 -> this.dynamic67()
                    68 -> this.dynamic68()
                    69 -> this.dynamic69()
                    70 -> this.dynamic70()
                    71 -> this.dynamic71()
                    72 -> this.dynamic72()
                    73 -> this.dynamic73()
                    74 -> this.dynamic74()
                    75 -> this.dynamic75()
                    76 -> this.dynamic76()
                    77 -> this.dynamic77()
                    78 -> this.dynamic78()
                    79 -> this.dynamic79()
                    80 -> this.dynamic80()
                    else -> false
                }
                scanner.removeFiles()
                if (!ok) {
                    println("Dynamic$taskNoStr Test#$testNoStr has not been tested")
                    return
                }
                if (!datScanner.eof()) {
                    println("Dynamic$taskNoStr Test#$testNoStr input is not enough")
                    return
                }
                if (!ansScanner.eof()) {
                    println("Dynamic$taskNoStr Test#$testNoStr output is not enough")
                    return
                }
            }
            println("Dynamic$taskNoStr has been tested!")
        }
    }

    private fun dynamic1(): Boolean {
        return false
    }

    private fun dynamic2(): Boolean {
        return false
    }

    private fun dynamic3(): Boolean {
        return false
    }

    private fun dynamic4(): Boolean {
        return false
    }

    private fun dynamic5(): Boolean {
        return false
    }

    private fun dynamic6(): Boolean {
        return false
    }

    private fun dynamic7(): Boolean {
        return false
    }

    private fun dynamic8(): Boolean {
        return false
    }

    private fun dynamic9(): Boolean {
        return false
    }

    private fun dynamic10(): Boolean {
        return false
    }

    private fun dynamic11(): Boolean {
        return false
    }

    private fun dynamic12(): Boolean {
        return false
    }

    private fun dynamic13(): Boolean {
        return false
    }

    private fun dynamic14(): Boolean {
        return false
    }

    private fun dynamic15(): Boolean {
        return false
    }

    private fun dynamic16(): Boolean {
        return false
    }

    private fun dynamic17(): Boolean {
        return false
    }

    private fun dynamic18(): Boolean {
        return false
    }

    private fun dynamic19(): Boolean {
        return false
    }

    private fun dynamic20(): Boolean {
        return false
    }

    private fun dynamic21(): Boolean {
        return false
    }

    private fun dynamic22(): Boolean {
        return false
    }

    private fun dynamic23(): Boolean {
        return false
    }

    private fun dynamic24(): Boolean {
        return false
    }

    private fun dynamic25(): Boolean {
        return false
    }

    private fun dynamic26(): Boolean {
        return false
    }

    private fun dynamic27(): Boolean {
        return false
    }

    private fun dynamic28(): Boolean {
        return false
    }

    private fun dynamic29(): Boolean {
        return false
    }

    private fun dynamic30(): Boolean {
        return false
    }

    private fun dynamic31(): Boolean {
        return false
    }

    private fun dynamic32(): Boolean {
        return false
    }

    private fun dynamic33(): Boolean {
        return false
    }

    private fun dynamic34(): Boolean {
        return false
    }

    private fun dynamic35(): Boolean {
        return false
    }

    private fun dynamic36(): Boolean {
        return false
    }

    private fun dynamic37(): Boolean {
        return false
    }

    private fun dynamic38(): Boolean {
        return false
    }

    private fun dynamic39(): Boolean {
        return false
    }

    private fun dynamic40(): Boolean {
        return false
    }

    private fun dynamic41(): Boolean {
        return false
    }

    private fun dynamic42(): Boolean {
        return false
    }

    private fun dynamic43(): Boolean {
        return false
    }

    private fun dynamic44(): Boolean {
        return false
    }

    private fun dynamic45(): Boolean {
        return false
    }

    private fun dynamic46(): Boolean {
        return false
    }

    private fun dynamic47(): Boolean {
        return false
    }

    private fun dynamic48(): Boolean {
        return false
    }

    private fun dynamic49(): Boolean {
        return false
    }

    private fun dynamic50(): Boolean {
        return false
    }

    private fun dynamic51(): Boolean {
        return false
    }

    private fun dynamic52(): Boolean {
        return false
    }

    private fun dynamic53(): Boolean {
        return false
    }

    private fun dynamic54(): Boolean {
        return false
    }

    private fun dynamic55(): Boolean {
        return false
    }

    private fun dynamic56(): Boolean {
        return false
    }

    private fun dynamic57(): Boolean {
        return false
    }

    private fun dynamic58(): Boolean {
        return false
    }

    private fun dynamic59(): Boolean {
        return false
    }

    private fun dynamic60(): Boolean {
        return false
    }

    private fun dynamic61(): Boolean {
        return false
    }

    private fun dynamic62(): Boolean {
        return false
    }

    private fun dynamic63(): Boolean {
        return false
    }

    private fun dynamic64(): Boolean {
        return false
    }

    private fun dynamic65(): Boolean {
        return false
    }

    private fun dynamic66(): Boolean {
        return false
    }

    private fun dynamic67(): Boolean {
        return false
    }

    private fun dynamic68(): Boolean {
        return false
    }

    private fun dynamic69(): Boolean {
        return false
    }

    private fun dynamic70(): Boolean {
        return false
    }

    private fun dynamic71(): Boolean {
        return false
    }

    private fun dynamic72(): Boolean {
        return false
    }

    private fun dynamic73(): Boolean {
        return false
    }

    private fun dynamic74(): Boolean {
        return false
    }

    private fun dynamic75(): Boolean {
        return false
    }

    private fun dynamic76(): Boolean {
        return false
    }

    private fun dynamic77(): Boolean {
        return false
    }

    private fun dynamic78(): Boolean {
        return false
    }

    private fun dynamic79(): Boolean {
        return false
    }

    private fun dynamic80(): Boolean {
        return false
    }
}