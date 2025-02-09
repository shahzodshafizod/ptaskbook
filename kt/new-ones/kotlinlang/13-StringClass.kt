package pack13string

class StringClass {

    var datScanner = scanner.Scanner()
    var ansScanner = scanner.Scanner()

    fun solve(taskNo: Int): Boolean {
        val path = "src/pack13string/tests/String$taskNo/"
        var fileName: String
        for (index in 1..100) {
            fileName = path + "%03d".format(index)
            if (!datScanner.changeReader("$fileName.dat")) {
                println("String$taskNo couldn't open datScanner source")
                return false
            }
            if (!ansScanner.changeReader("$fileName.ans")) {
                println("String$taskNo couldn't open ansScanner source")
                return false
            }
            val ok = when (taskNo) {
                1 -> this.string01()
                2 -> this.string02()
                3 -> this.string03()
                4 -> this.string04()
                5 -> this.string05()
                6 -> this.string06()
                7 -> this.string07()
                8 -> this.string08()
                9 -> this.string09()
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
                println("String$taskNo Test#$index has not been tested")
                return false
            }
            if (!datScanner.eof()) {
                println("String$taskNo Test#$index input is not enough")
                return false
            }
            if (!ansScanner.eof()) {
                println("String$taskNo Test#$index output is not enough")
                return false
            }
        }
        println("String$taskNo Tested!")
        return true
    }

    private fun string01(): Boolean {
        return false
    }

    private fun string02(): Boolean {
        return false
    }

    private fun string03(): Boolean {
        return false
    }

    private fun string04(): Boolean {
        return false
    }

    private fun string05(): Boolean {
        return false
    }

    private fun string06(): Boolean {
        return false
    }

    private fun string07(): Boolean {
        return false
    }

    private fun string08(): Boolean {
        return false
    }

    private fun string09(): Boolean {
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