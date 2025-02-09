package pack08series

class SeriesClass {

    var datScanner = scanner.Scanner()
    var ansScanner = scanner.Scanner()

    fun solve(taskNo: Int): Boolean {
        val path = "src/pack08series/tests/Series$taskNo/"
        var fileName: String
        for (index in 1..100) {
            fileName = path + "%03d".format(index)
            if (!datScanner.changeReader("$fileName.dat")) {
                println("Series$taskNo couldn't open datScanner source")
                return false
            }
            if (!ansScanner.changeReader("$fileName.ans")) {
                println("Series$taskNo couldn't open ansScanner source")
                return false
            }
            val ok = when (taskNo) {
                1 -> this.series01()
                2 -> this.series02()
                3 -> this.series03()
                4 -> this.series04()
                5 -> this.series05()
                6 -> this.series06()
                7 -> this.series07()
                8 -> this.series08()
                9 -> this.series09()
                10 -> this.series10()
                11 -> this.series11()
                12 -> this.series12()
                13 -> this.series13()
                14 -> this.series14()
                15 -> this.series15()
                16 -> this.series16()
                17 -> this.series17()
                18 -> this.series18()
                19 -> this.series19()
                20 -> this.series20()
                21 -> this.series21()
                22 -> this.series22()
                23 -> this.series23()
                24 -> this.series24()
                25 -> this.series25()
                26 -> this.series26()
                27 -> this.series27()
                28 -> this.series28()
                29 -> this.series29()
                30 -> this.series30()
                31 -> this.series31()
                32 -> this.series32()
                33 -> this.series33()
                34 -> this.series34()
                35 -> this.series35()
                36 -> this.series36()
                37 -> this.series37()
                38 -> this.series38()
                39 -> this.series39()
                40 -> this.series40()
                else -> false
            }
            scanner.removeFiles()
            if (!ok) {
                println("Series$taskNo Test#$index has not been tested")
                return false
            }
            if (!datScanner.eof()) {
                println("Series$taskNo Test#$index input is not enough")
                return false
            }
            if (!ansScanner.eof()) {
                println("Series$taskNo Test#$index output is not enough")
                return false
            }
        }
        println("Series$taskNo Tested!")
        return true
    }

    private fun series01(): Boolean {
        return false
    }

    private fun series02(): Boolean {
        return false
    }

    private fun series03(): Boolean {
        return false
    }

    private fun series04(): Boolean {
        return false
    }

    private fun series05(): Boolean {
        return false
    }

    private fun series06(): Boolean {
        return false
    }

    private fun series07(): Boolean {
        return false
    }

    private fun series08(): Boolean {
        return false
    }

    private fun series09(): Boolean {
        return false
    }

    private fun series10(): Boolean {
        return false
    }

    private fun series11(): Boolean {
        return false
    }

    private fun series12(): Boolean {
        return false
    }

    private fun series13(): Boolean {
        return false
    }

    private fun series14(): Boolean {
        return false
    }

    private fun series15(): Boolean {
        return false
    }

    private fun series16(): Boolean {
        return false
    }

    private fun series17(): Boolean {
        return false
    }

    private fun series18(): Boolean {
        return false
    }

    private fun series19(): Boolean {
        return false
    }

    private fun series20(): Boolean {
        return false
    }

    private fun series21(): Boolean {
        return false
    }

    private fun series22(): Boolean {
        return false
    }

    private fun series23(): Boolean {
        return false
    }

    private fun series24(): Boolean {
        return false
    }

    private fun series25(): Boolean {
        return false
    }

    private fun series26(): Boolean {
        return false
    }

    private fun series27(): Boolean {
        return false
    }

    private fun series28(): Boolean {
        return false
    }

    private fun series29(): Boolean {
        return false
    }

    private fun series30(): Boolean {
        return false
    }

    private fun series31(): Boolean {
        return false
    }

    private fun series32(): Boolean {
        return false
    }

    private fun series33(): Boolean {
        return false
    }

    private fun series34(): Boolean {
        return false
    }

    private fun series35(): Boolean {
        return false
    }

    private fun series36(): Boolean {
        return false
    }

    private fun series37(): Boolean {
        return false
    }

    private fun series38(): Boolean {
        return false
    }

    private fun series39(): Boolean {
        return false
    }

    private fun series40(): Boolean {
        return false
    }
}