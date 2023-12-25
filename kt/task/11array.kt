package task

class ArrayTask {

    var datScanner = scanner.Scanner()
    var ansScanner = scanner.Scanner()

    fun solve(pathPrefix: String) {
        for (taskNo in 1..140) {
            var taskNoStr = "%03d".format(taskNo)
            var fileName: String
            for (testNo in 1..100) {
                var testNoStr = "%03d".format(testNo)
                fileName = "$pathPrefix/11array/Array$taskNoStr/$testNoStr/$testNoStr"
                if (!datScanner.changeReader("$fileName.dat")) {
                    println("Array$taskNoStr couldn't open datScanner source")
                    return
                }
                if (!ansScanner.changeReader("$fileName.ans")) {
                    println("Array$taskNoStr couldn't open ansScanner source")
                    return
                }
                val ok = when (taskNo) {
                    1 -> this.array1()
                    2 -> this.array2()
                    3 -> this.array3()
                    4 -> this.array4()
                    5 -> this.array5()
                    6 -> this.array6()
                    7 -> this.array7()
                    8 -> this.array8()
                    9 -> this.array9()
                    10 -> this.array10()
                    11 -> this.array11()
                    12 -> this.array12()
                    13 -> this.array13()
                    14 -> this.array14()
                    15 -> this.array15()
                    16 -> this.array16()
                    17 -> this.array17()
                    18 -> this.array18()
                    19 -> this.array19()
                    20 -> this.array20()
                    21 -> this.array21()
                    22 -> this.array22()
                    23 -> this.array23()
                    24 -> this.array24()
                    25 -> this.array25()
                    26 -> this.array26()
                    27 -> this.array27()
                    28 -> this.array28()
                    29 -> this.array29()
                    30 -> this.array30()
                    31 -> this.array31()
                    32 -> this.array32()
                    33 -> this.array33()
                    34 -> this.array34()
                    35 -> this.array35()
                    36 -> this.array36()
                    37 -> this.array37()
                    38 -> this.array38()
                    39 -> this.array39()
                    40 -> this.array40()
                    41 -> this.array41()
                    42 -> this.array42()
                    43 -> this.array43()
                    44 -> this.array44()
                    45 -> this.array45()
                    46 -> this.array46()
                    47 -> this.array47()
                    48 -> this.array48()
                    49 -> this.array49()
                    50 -> this.array50()
                    51 -> this.array51()
                    52 -> this.array52()
                    53 -> this.array53()
                    54 -> this.array54()
                    55 -> this.array55()
                    56 -> this.array56()
                    57 -> this.array57()
                    58 -> this.array58()
                    59 -> this.array59()
                    60 -> this.array60()
                    61 -> this.array61()
                    62 -> this.array62()
                    63 -> this.array63()
                    64 -> this.array64()
                    65 -> this.array65()
                    66 -> this.array66()
                    67 -> this.array67()
                    68 -> this.array68()
                    69 -> this.array69()
                    70 -> this.array70()
                    71 -> this.array71()
                    72 -> this.array72()
                    73 -> this.array73()
                    74 -> this.array74()
                    75 -> this.array75()
                    76 -> this.array76()
                    77 -> this.array77()
                    78 -> this.array78()
                    79 -> this.array79()
                    80 -> this.array80()
                    81 -> this.array81()
                    82 -> this.array82()
                    83 -> this.array83()
                    84 -> this.array84()
                    85 -> this.array85()
                    86 -> this.array86()
                    87 -> this.array87()
                    88 -> this.array88()
                    89 -> this.array89()
                    90 -> this.array90()
                    91 -> this.array91()
                    92 -> this.array92()
                    93 -> this.array93()
                    94 -> this.array94()
                    95 -> this.array95()
                    96 -> this.array96()
                    97 -> this.array97()
                    98 -> this.array98()
                    99 -> this.array99()
                    100 -> this.array100()
                    101 -> this.array101()
                    102 -> this.array102()
                    103 -> this.array103()
                    104 -> this.array104()
                    105 -> this.array105()
                    106 -> this.array106()
                    107 -> this.array107()
                    108 -> this.array108()
                    109 -> this.array109()
                    110 -> this.array110()
                    111 -> this.array111()
                    112 -> this.array112()
                    113 -> this.array113()
                    114 -> this.array114()
                    115 -> this.array115()
                    116 -> this.array116()
                    117 -> this.array117()
                    118 -> this.array118()
                    119 -> this.array119()
                    120 -> this.array120()
                    121 -> this.array121()
                    122 -> this.array122()
                    123 -> this.array123()
                    124 -> this.array124()
                    125 -> this.array125()
                    126 -> this.array126()
                    127 -> this.array127()
                    128 -> this.array128()
                    129 -> this.array129()
                    130 -> this.array130()
                    131 -> this.array131()
                    132 -> this.array132()
                    133 -> this.array133()
                    134 -> this.array134()
                    135 -> this.array135()
                    136 -> this.array136()
                    137 -> this.array137()
                    138 -> this.array138()
                    139 -> this.array139()
                    140 -> this.array140()
                    else -> false
                }
                scanner.removeFiles()
                if (!ok) {
                    println("Array$taskNoStr Test#$testNoStr has not been tested")
                    return
                }
                if (!datScanner.eof()) {
                    println("Array$taskNoStr Test#$testNoStr input is not enough")
                    return
                }
                if (!ansScanner.eof()) {
                    println("Array$taskNoStr Test#$testNoStr output is not enough")
                    return
                }
            }
            println("Array$taskNoStr has been tested!")
        }
    }

    private fun array1(): Boolean {
        return false
    }

    private fun array2(): Boolean {
        return false
    }

    private fun array3(): Boolean {
        return false
    }

    private fun array4(): Boolean {
        return false
    }

    private fun array5(): Boolean {
        return false
    }

    private fun array6(): Boolean {
        return false
    }

    private fun array7(): Boolean {
        return false
    }

    private fun array8(): Boolean {
        return false
    }

    private fun array9(): Boolean {
        return false
    }

    private fun array10(): Boolean {
        return false
    }

    private fun array11(): Boolean {
        return false
    }

    private fun array12(): Boolean {
        return false
    }

    private fun array13(): Boolean {
        return false
    }

    private fun array14(): Boolean {
        return false
    }

    private fun array15(): Boolean {
        return false
    }

    private fun array16(): Boolean {
        return false
    }

    private fun array17(): Boolean {
        return false
    }

    private fun array18(): Boolean {
        return false
    }

    private fun array19(): Boolean {
        return false
    }

    private fun array20(): Boolean {
        return false
    }

    private fun array21(): Boolean {
        return false
    }

    private fun array22(): Boolean {
        return false
    }

    private fun array23(): Boolean {
        return false
    }

    private fun array24(): Boolean {
        return false
    }

    private fun array25(): Boolean {
        return false
    }

    private fun array26(): Boolean {
        return false
    }

    private fun array27(): Boolean {
        return false
    }

    private fun array28(): Boolean {
        return false
    }

    private fun array29(): Boolean {
        return false
    }

    private fun array30(): Boolean {
        return false
    }

    private fun array31(): Boolean {
        return false
    }

    private fun array32(): Boolean {
        return false
    }

    private fun array33(): Boolean {
        return false
    }

    private fun array34(): Boolean {
        return false
    }

    private fun array35(): Boolean {
        return false
    }

    private fun array36(): Boolean {
        return false
    }

    private fun array37(): Boolean {
        return false
    }

    private fun array38(): Boolean {
        return false
    }

    private fun array39(): Boolean {
        return false
    }

    private fun array40(): Boolean {
        return false
    }

    private fun array41(): Boolean {
        return false
    }

    private fun array42(): Boolean {
        return false
    }

    private fun array43(): Boolean {
        return false
    }

    private fun array44(): Boolean {
        return false
    }

    private fun array45(): Boolean {
        return false
    }

    private fun array46(): Boolean {
        return false
    }

    private fun array47(): Boolean {
        return false
    }

    private fun array48(): Boolean {
        return false
    }

    private fun array49(): Boolean {
        return false
    }

    private fun array50(): Boolean {
        return false
    }

    private fun array51(): Boolean {
        return false
    }

    private fun array52(): Boolean {
        return false
    }

    private fun array53(): Boolean {
        return false
    }

    private fun array54(): Boolean {
        return false
    }

    private fun array55(): Boolean {
        return false
    }

    private fun array56(): Boolean {
        return false
    }

    private fun array57(): Boolean {
        return false
    }

    private fun array58(): Boolean {
        return false
    }

    private fun array59(): Boolean {
        return false
    }

    private fun array60(): Boolean {
        return false
    }

    private fun array61(): Boolean {
        return false
    }

    private fun array62(): Boolean {
        return false
    }

    private fun array63(): Boolean {
        return false
    }

    private fun array64(): Boolean {
        return false
    }

    private fun array65(): Boolean {
        return false
    }

    private fun array66(): Boolean {
        return false
    }

    private fun array67(): Boolean {
        return false
    }

    private fun array68(): Boolean {
        return false
    }

    private fun array69(): Boolean {
        return false
    }

    private fun array70(): Boolean {
        return false
    }

    private fun array71(): Boolean {
        return false
    }

    private fun array72(): Boolean {
        return false
    }

    private fun array73(): Boolean {
        return false
    }

    private fun array74(): Boolean {
        return false
    }

    private fun array75(): Boolean {
        return false
    }

    private fun array76(): Boolean {
        return false
    }

    private fun array77(): Boolean {
        return false
    }

    private fun array78(): Boolean {
        return false
    }

    private fun array79(): Boolean {
        return false
    }

    private fun array80(): Boolean {
        return false
    }

    private fun array81(): Boolean {
        return false
    }

    private fun array82(): Boolean {
        return false
    }

    private fun array83(): Boolean {
        return false
    }

    private fun array84(): Boolean {
        return false
    }

    private fun array85(): Boolean {
        return false
    }

    private fun array86(): Boolean {
        return false
    }

    private fun array87(): Boolean {
        return false
    }

    private fun array88(): Boolean {
        return false
    }

    private fun array89(): Boolean {
        return false
    }

    private fun array90(): Boolean {
        return false
    }

    private fun array91(): Boolean {
        return false
    }

    private fun array92(): Boolean {
        return false
    }

    private fun array93(): Boolean {
        return false
    }

    private fun array94(): Boolean {
        return false
    }

    private fun array95(): Boolean {
        return false
    }

    private fun array96(): Boolean {
        return false
    }

    private fun array97(): Boolean {
        return false
    }

    private fun array98(): Boolean {
        return false
    }

    private fun array99(): Boolean {
        return false
    }

    private fun array100(): Boolean {
        return false
    }

    private fun array101(): Boolean {
        return false
    }

    private fun array102(): Boolean {
        return false
    }

    private fun array103(): Boolean {
        return false
    }

    private fun array104(): Boolean {
        return false
    }

    private fun array105(): Boolean {
        return false
    }

    private fun array106(): Boolean {
        return false
    }

    private fun array107(): Boolean {
        return false
    }

    private fun array108(): Boolean {
        return false
    }

    private fun array109(): Boolean {
        return false
    }

    private fun array110(): Boolean {
        return false
    }

    private fun array111(): Boolean {
        return false
    }

    private fun array112(): Boolean {
        return false
    }

    private fun array113(): Boolean {
        return false
    }

    private fun array114(): Boolean {
        return false
    }

    private fun array115(): Boolean {
        return false
    }

    private fun array116(): Boolean {
        return false
    }

    private fun array117(): Boolean {
        return false
    }

    private fun array118(): Boolean {
        return false
    }

    private fun array119(): Boolean {
        return false
    }

    private fun array120(): Boolean {
        return false
    }

    private fun array121(): Boolean {
        return false
    }

    private fun array122(): Boolean {
        return false
    }

    private fun array123(): Boolean {
        return false
    }

    private fun array124(): Boolean {
        return false
    }

    private fun array125(): Boolean {
        return false
    }

    private fun array126(): Boolean {
        return false
    }

    private fun array127(): Boolean {
        return false
    }

    private fun array128(): Boolean {
        return false
    }

    private fun array129(): Boolean {
        return false
    }

    private fun array130(): Boolean {
        return false
    }

    private fun array131(): Boolean {
        return false
    }

    private fun array132(): Boolean {
        return false
    }

    private fun array133(): Boolean {
        return false
    }

    private fun array134(): Boolean {
        return false
    }

    private fun array135(): Boolean {
        return false
    }

    private fun array136(): Boolean {
        return false
    }

    private fun array137(): Boolean {
        return false
    }

    private fun array138(): Boolean {
        return false
    }

    private fun array139(): Boolean {
        return false
    }

    private fun array140(): Boolean {
        return false
    }
}