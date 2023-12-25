package task

class TreeTask {

    var datScanner = scanner.Scanner()
    var ansScanner = scanner.Scanner()

    fun solve(pathPrefix: String) {
        for (taskNo in 1..100) {
            var taskNoStr = "%03d".format(taskNo)
            var fileName: String
            for (testNo in 1..100) {
                var testNoStr = "%03d".format(testNo)
                fileName = "$pathPrefix/19tree/Tree$taskNoStr/$testNoStr/$testNoStr"
                if (!datScanner.changeReader("$fileName.dat")) {
                    println("Tree$taskNoStr couldn't open datScanner source")
                    return
                }
                if (!ansScanner.changeReader("$fileName.ans")) {
                    println("Tree$taskNoStr couldn't open ansScanner source")
                    return
                }
                val ok = when (taskNo) {
                    1 -> this.tree1()
                    2 -> this.tree2()
                    3 -> this.tree3()
                    4 -> this.tree4()
                    5 -> this.tree5()
                    6 -> this.tree6()
                    7 -> this.tree7()
                    8 -> this.tree8()
                    9 -> this.tree9()
                    10 -> this.tree10()
                    11 -> this.tree11()
                    12 -> this.tree12()
                    13 -> this.tree13()
                    14 -> this.tree14()
                    15 -> this.tree15()
                    16 -> this.tree16()
                    17 -> this.tree17()
                    18 -> this.tree18()
                    19 -> this.tree19()
                    20 -> this.tree20()
                    21 -> this.tree21()
                    22 -> this.tree22()
                    23 -> this.tree23()
                    24 -> this.tree24()
                    25 -> this.tree25()
                    26 -> this.tree26()
                    27 -> this.tree27()
                    28 -> this.tree28()
                    29 -> this.tree29()
                    30 -> this.tree30()
                    31 -> this.tree31()
                    32 -> this.tree32()
                    33 -> this.tree33()
                    34 -> this.tree34()
                    35 -> this.tree35()
                    36 -> this.tree36()
                    37 -> this.tree37()
                    38 -> this.tree38()
                    39 -> this.tree39()
                    40 -> this.tree40()
                    41 -> this.tree41()
                    42 -> this.tree42()
                    43 -> this.tree43()
                    44 -> this.tree44()
                    45 -> this.tree45()
                    46 -> this.tree46()
                    47 -> this.tree47()
                    48 -> this.tree48()
                    49 -> this.tree49()
                    50 -> this.tree50()
                    51 -> this.tree51()
                    52 -> this.tree52()
                    53 -> this.tree53()
                    54 -> this.tree54()
                    55 -> this.tree55()
                    56 -> this.tree56()
                    57 -> this.tree57()
                    58 -> this.tree58()
                    59 -> this.tree59()
                    60 -> this.tree60()
                    61 -> this.tree61()
                    62 -> this.tree62()
                    63 -> this.tree63()
                    64 -> this.tree64()
                    65 -> this.tree65()
                    66 -> this.tree66()
                    67 -> this.tree67()
                    68 -> this.tree68()
                    69 -> this.tree69()
                    70 -> this.tree70()
                    71 -> this.tree71()
                    72 -> this.tree72()
                    73 -> this.tree73()
                    74 -> this.tree74()
                    75 -> this.tree75()
                    76 -> this.tree76()
                    77 -> this.tree77()
                    78 -> this.tree78()
                    79 -> this.tree79()
                    80 -> this.tree80()
                    81 -> this.tree81()
                    82 -> this.tree82()
                    83 -> this.tree83()
                    84 -> this.tree84()
                    85 -> this.tree85()
                    86 -> this.tree86()
                    87 -> this.tree87()
                    88 -> this.tree88()
                    89 -> this.tree89()
                    90 -> this.tree90()
                    91 -> this.tree91()
                    92 -> this.tree92()
                    93 -> this.tree93()
                    94 -> this.tree94()
                    95 -> this.tree95()
                    96 -> this.tree96()
                    97 -> this.tree97()
                    98 -> this.tree98()
                    99 -> this.tree99()
                    100 -> this.tree100()
                    else -> false
                }
                scanner.removeFiles()
                if (!ok) {
                    println("Tree$taskNoStr Test#$testNoStr has not been tested")
                    return
                }
                if (!datScanner.eof()) {
                    println("Tree$taskNoStr Test#$testNoStr input is not enough")
                    return
                }
                if (!ansScanner.eof()) {
                    println("Tree$taskNoStr Test#$testNoStr output is not enough")
                    return
                }
            }
            println("Tree$taskNoStr has been tested!")
        }
    }

    private fun tree1(): Boolean {
        return false
    }

    private fun tree2(): Boolean {
        return false
    }

    private fun tree3(): Boolean {
        return false
    }

    private fun tree4(): Boolean {
        return false
    }

    private fun tree5(): Boolean {
        return false
    }

    private fun tree6(): Boolean {
        return false
    }

    private fun tree7(): Boolean {
        return false
    }

    private fun tree8(): Boolean {
        return false
    }

    private fun tree9(): Boolean {
        return false
    }

    private fun tree10(): Boolean {
        return false
    }

    private fun tree11(): Boolean {
        return false
    }

    private fun tree12(): Boolean {
        return false
    }

    private fun tree13(): Boolean {
        return false
    }

    private fun tree14(): Boolean {
        return false
    }

    private fun tree15(): Boolean {
        return false
    }

    private fun tree16(): Boolean {
        return false
    }

    private fun tree17(): Boolean {
        return false
    }

    private fun tree18(): Boolean {
        return false
    }

    private fun tree19(): Boolean {
        return false
    }

    private fun tree20(): Boolean {
        return false
    }

    private fun tree21(): Boolean {
        return false
    }

    private fun tree22(): Boolean {
        return false
    }

    private fun tree23(): Boolean {
        return false
    }

    private fun tree24(): Boolean {
        return false
    }

    private fun tree25(): Boolean {
        return false
    }

    private fun tree26(): Boolean {
        return false
    }

    private fun tree27(): Boolean {
        return false
    }

    private fun tree28(): Boolean {
        return false
    }

    private fun tree29(): Boolean {
        return false
    }

    private fun tree30(): Boolean {
        return false
    }

    private fun tree31(): Boolean {
        return false
    }

    private fun tree32(): Boolean {
        return false
    }

    private fun tree33(): Boolean {
        return false
    }

    private fun tree34(): Boolean {
        return false
    }

    private fun tree35(): Boolean {
        return false
    }

    private fun tree36(): Boolean {
        return false
    }

    private fun tree37(): Boolean {
        return false
    }

    private fun tree38(): Boolean {
        return false
    }

    private fun tree39(): Boolean {
        return false
    }

    private fun tree40(): Boolean {
        return false
    }

    private fun tree41(): Boolean {
        return false
    }

    private fun tree42(): Boolean {
        return false
    }

    private fun tree43(): Boolean {
        return false
    }

    private fun tree44(): Boolean {
        return false
    }

    private fun tree45(): Boolean {
        return false
    }

    private fun tree46(): Boolean {
        return false
    }

    private fun tree47(): Boolean {
        return false
    }

    private fun tree48(): Boolean {
        return false
    }

    private fun tree49(): Boolean {
        return false
    }

    private fun tree50(): Boolean {
        return false
    }

    private fun tree51(): Boolean {
        return false
    }

    private fun tree52(): Boolean {
        return false
    }

    private fun tree53(): Boolean {
        return false
    }

    private fun tree54(): Boolean {
        return false
    }

    private fun tree55(): Boolean {
        return false
    }

    private fun tree56(): Boolean {
        return false
    }

    private fun tree57(): Boolean {
        return false
    }

    private fun tree58(): Boolean {
        return false
    }

    private fun tree59(): Boolean {
        return false
    }

    private fun tree60(): Boolean {
        return false
    }

    private fun tree61(): Boolean {
        return false
    }

    private fun tree62(): Boolean {
        return false
    }

    private fun tree63(): Boolean {
        return false
    }

    private fun tree64(): Boolean {
        return false
    }

    private fun tree65(): Boolean {
        return false
    }

    private fun tree66(): Boolean {
        return false
    }

    private fun tree67(): Boolean {
        return false
    }

    private fun tree68(): Boolean {
        return false
    }

    private fun tree69(): Boolean {
        return false
    }

    private fun tree70(): Boolean {
        return false
    }

    private fun tree71(): Boolean {
        return false
    }

    private fun tree72(): Boolean {
        return false
    }

    private fun tree73(): Boolean {
        return false
    }

    private fun tree74(): Boolean {
        return false
    }

    private fun tree75(): Boolean {
        return false
    }

    private fun tree76(): Boolean {
        return false
    }

    private fun tree77(): Boolean {
        return false
    }

    private fun tree78(): Boolean {
        return false
    }

    private fun tree79(): Boolean {
        return false
    }

    private fun tree80(): Boolean {
        return false
    }

    private fun tree81(): Boolean {
        return false
    }

    private fun tree82(): Boolean {
        return false
    }

    private fun tree83(): Boolean {
        return false
    }

    private fun tree84(): Boolean {
        return false
    }

    private fun tree85(): Boolean {
        return false
    }

    private fun tree86(): Boolean {
        return false
    }

    private fun tree87(): Boolean {
        return false
    }

    private fun tree88(): Boolean {
        return false
    }

    private fun tree89(): Boolean {
        return false
    }

    private fun tree90(): Boolean {
        return false
    }

    private fun tree91(): Boolean {
        return false
    }

    private fun tree92(): Boolean {
        return false
    }

    private fun tree93(): Boolean {
        return false
    }

    private fun tree94(): Boolean {
        return false
    }

    private fun tree95(): Boolean {
        return false
    }

    private fun tree96(): Boolean {
        return false
    }

    private fun tree97(): Boolean {
        return false
    }

    private fun tree98(): Boolean {
        return false
    }

    private fun tree99(): Boolean {
        return false
    }

    private fun tree100(): Boolean {
        return false
    }
}