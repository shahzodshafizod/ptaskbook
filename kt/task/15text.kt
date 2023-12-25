package task

class TextTask {

    var datScanner = scanner.Scanner()
    var ansScanner = scanner.Scanner()

    fun solve(pathPrefix: String) {
        for (taskNo in 1..60) {
            var taskNoStr = "%03d".format(taskNo)
            var fileName: String
            for (testNo in 1..100) {
                var testNoStr = "%03d".format(testNo)
                fileName = "$pathPrefix/15text/Text$taskNoStr/$testNoStr/$testNoStr"
                if (!datScanner.changeReader("$fileName.dat")) {
                    println("Text$taskNoStr couldn't open datScanner source")
                    return
                }
                if (!ansScanner.changeReader("$fileName.ans")) {
                    println("Text$taskNoStr couldn't open ansScanner source")
                    return
                }
                val ok = when (taskNo) {
                    1 -> this.text1()
                    2 -> this.text2()
                    3 -> this.text3()
                    4 -> this.text4()
                    5 -> this.text5()
                    6 -> this.text6()
                    7 -> this.text7()
                    8 -> this.text8()
                    9 -> this.text9()
                    10 -> this.text10()
                    11 -> this.text11()
                    12 -> this.text12()
                    13 -> this.text13()
                    14 -> this.text14()
                    15 -> this.text15()
                    16 -> this.text16()
                    17 -> this.text17()
                    18 -> this.text18()
                    19 -> this.text19()
                    20 -> this.text20()
                    21 -> this.text21()
                    22 -> this.text22()
                    23 -> this.text23()
                    24 -> this.text24()
                    25 -> this.text25()
                    26 -> this.text26()
                    27 -> this.text27()
                    28 -> this.text28()
                    29 -> this.text29()
                    30 -> this.text30()
                    31 -> this.text31()
                    32 -> this.text32()
                    33 -> this.text33()
                    34 -> this.text34()
                    35 -> this.text35()
                    36 -> this.text36()
                    37 -> this.text37()
                    38 -> this.text38()
                    39 -> this.text39()
                    40 -> this.text40()
                    41 -> this.text41()
                    42 -> this.text42()
                    43 -> this.text43()
                    44 -> this.text44()
                    45 -> this.text45()
                    46 -> this.text46()
                    47 -> this.text47()
                    48 -> this.text48()
                    49 -> this.text49()
                    50 -> this.text50()
                    51 -> this.text51()
                    52 -> this.text52()
                    53 -> this.text53()
                    54 -> this.text54()
                    55 -> this.text55()
                    56 -> this.text56()
                    57 -> this.text57()
                    58 -> this.text58()
                    59 -> this.text59()
                    60 -> this.text60()
                    else -> false
                }
                scanner.removeFiles()
                if (!ok) {
                    println("Text$taskNoStr Test#$testNoStr has not been tested")
                    return
                }
                if (!datScanner.eof()) {
                    println("Text$taskNoStr Test#$testNoStr input is not enough")
                    return
                }
                if (!ansScanner.eof()) {
                    println("Text$taskNoStr Test#$testNoStr output is not enough")
                    return
                }
            }
            println("Text$taskNoStr has been tested!")
        }
    }

    private fun text1(): Boolean {
        return false
    }

    private fun text2(): Boolean {
        return false
    }

    private fun text3(): Boolean {
        return false
    }

    private fun text4(): Boolean {
        return false
    }

    private fun text5(): Boolean {
        return false
    }

    private fun text6(): Boolean {
        return false
    }

    private fun text7(): Boolean {
        return false
    }

    private fun text8(): Boolean {
        return false
    }

    private fun text9(): Boolean {
        return false
    }

    private fun text10(): Boolean {
        return false
    }

    private fun text11(): Boolean {
        return false
    }

    private fun text12(): Boolean {
        return false
    }

    private fun text13(): Boolean {
        return false
    }

    private fun text14(): Boolean {
        return false
    }

    private fun text15(): Boolean {
        return false
    }

    private fun text16(): Boolean {
        return false
    }

    private fun text17(): Boolean {
        return false
    }

    private fun text18(): Boolean {
        return false
    }

    private fun text19(): Boolean {
        return false
    }

    private fun text20(): Boolean {
        return false
    }

    private fun text21(): Boolean {
        return false
    }

    private fun text22(): Boolean {
        return false
    }

    private fun text23(): Boolean {
        return false
    }

    private fun text24(): Boolean {
        return false
    }

    private fun text25(): Boolean {
        return false
    }

    private fun text26(): Boolean {
        return false
    }

    private fun text27(): Boolean {
        return false
    }

    private fun text28(): Boolean {
        return false
    }

    private fun text29(): Boolean {
        return false
    }

    private fun text30(): Boolean {
        return false
    }

    private fun text31(): Boolean {
        return false
    }

    private fun text32(): Boolean {
        return false
    }

    private fun text33(): Boolean {
        return false
    }

    private fun text34(): Boolean {
        return false
    }

    private fun text35(): Boolean {
        return false
    }

    private fun text36(): Boolean {
        return false
    }

    private fun text37(): Boolean {
        return false
    }

    private fun text38(): Boolean {
        return false
    }

    private fun text39(): Boolean {
        return false
    }

    private fun text40(): Boolean {
        return false
    }

    private fun text41(): Boolean {
        return false
    }

    private fun text42(): Boolean {
        return false
    }

    private fun text43(): Boolean {
        return false
    }

    private fun text44(): Boolean {
        return false
    }

    private fun text45(): Boolean {
        return false
    }

    private fun text46(): Boolean {
        return false
    }

    private fun text47(): Boolean {
        return false
    }

    private fun text48(): Boolean {
        return false
    }

    private fun text49(): Boolean {
        return false
    }

    private fun text50(): Boolean {
        return false
    }

    private fun text51(): Boolean {
        return false
    }

    private fun text52(): Boolean {
        return false
    }

    private fun text53(): Boolean {
        return false
    }

    private fun text54(): Boolean {
        return false
    }

    private fun text55(): Boolean {
        return false
    }

    private fun text56(): Boolean {
        return false
    }

    private fun text57(): Boolean {
        return false
    }

    private fun text58(): Boolean {
        return false
    }

    private fun text59(): Boolean {
        return false
    }

    private fun text60(): Boolean {
        return false
    }
}