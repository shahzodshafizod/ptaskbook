package pack10minmax

class MinMaxClass {

    var datScanner = scanner.Scanner()
    var ansScanner = scanner.Scanner()

    fun solve(taskNo: Int): Boolean {
        val path = "src/pack10minmax/tests/Minmax$taskNo/"
        var fileName: String
        for (index in 1..100) {
            fileName = path + "%03d".format(index)
            if (!datScanner.changeReader("$fileName.dat")) {
                println("MinMax$taskNo couldn't open datScanner source")
                return false
            }
            if (!ansScanner.changeReader("$fileName.ans")) {
                println("MinMax$taskNo couldn't open ansScanner source")
                return false
            }
            val ok = when (taskNo) {
                1 -> this.minMax01()
                2 -> this.minMax02()
                3 -> this.minMax03()
                4 -> this.minMax04()
                5 -> this.minMax05()
                6 -> this.minMax06()
                7 -> this.minMax07()
                8 -> this.minMax08()
                9 -> this.minMax09()
                10 -> this.minMax10()
                11 -> this.minMax11()
                12 -> this.minMax12()
                13 -> this.minMax13()
                14 -> this.minMax14()
                15 -> this.minMax15()
                16 -> this.minMax16()
                17 -> this.minMax17()
                18 -> this.minMax18()
                19 -> this.minMax19()
                20 -> this.minMax20()
                21 -> this.minMax21()
                22 -> this.minMax22()
                23 -> this.minMax23()
                24 -> this.minMax24()
                25 -> this.minMax25()
                26 -> this.minMax26()
                27 -> this.minMax27()
                28 -> this.minMax28()
                29 -> this.minMax29()
                30 -> this.minMax30()
                else -> false
            }
            scanner.removeFiles()
            if (!ok) {
                println("MinMax$taskNo Test#$index has not been tested")
                return false
            }
            if (!datScanner.eof()) {
                println("MinMax$taskNo Test#$index input is not enough")
                return false
            }
            if (!ansScanner.eof()) {
                println("MinMax$taskNo Test#$index output is not enough")
                return false
            }
        }
        println("MinMax$taskNo Tested!")
        return true
    }

    private fun minMax01(): Boolean {
        return false
    }

    private fun minMax02(): Boolean {
        return false
    }

    private fun minMax03(): Boolean {
        return false
    }

    private fun minMax04(): Boolean {
        return false
    }

    private fun minMax05(): Boolean {
        return false
    }

    private fun minMax06(): Boolean {
        return false
    }

    private fun minMax07(): Boolean {
        return false
    }

    private fun minMax08(): Boolean {
        return false
    }

    private fun minMax09(): Boolean {
        return false
    }

    private fun minMax10(): Boolean {
        return false
    }

    private fun minMax11(): Boolean {
        return false
    }

    private fun minMax12(): Boolean {
        return false
    }

    private fun minMax13(): Boolean {
        return false
    }

    private fun minMax14(): Boolean {
        return false
    }

    private fun minMax15(): Boolean {
        return false
    }

    private fun minMax16(): Boolean {
        return false
    }

    private fun minMax17(): Boolean {
        return false
    }

    private fun minMax18(): Boolean {
        return false
    }

    private fun minMax19(): Boolean {
        return false
    }

    private fun minMax20(): Boolean {
        return false
    }

    private fun minMax21(): Boolean {
        return false
    }

    private fun minMax22(): Boolean {
        return false
    }

    private fun minMax23(): Boolean {
        return false
    }

    private fun minMax24(): Boolean {
        return false
    }

    private fun minMax25(): Boolean {
        return false
    }

    private fun minMax26(): Boolean {
        return false
    }

    private fun minMax27(): Boolean {
        return false
    }

    private fun minMax28(): Boolean {
        return false
    }

    private fun minMax29(): Boolean {
        return false
    }

    private fun minMax30(): Boolean {
        return false
    }
}