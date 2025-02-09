package pack17recur

class RecurClass {

    var datScanner = scanner.Scanner()
    var ansScanner = scanner.Scanner()

    fun solve(taskNo: Int): Boolean {
        val path = "src/pack17recur/tests/Recur$taskNo/"
        var fileName: String
        for (index in 1..100) {
            fileName = path + "%03d".format(index)
            if (!datScanner.changeReader("$fileName.dat")) {
                println("Recur$taskNo couldn't open datScanner source")
                return false
            }
            if (!ansScanner.changeReader("$fileName.ans")) {
                println("Recur$taskNo couldn't open ansScanner source")
                return false
            }
            val ok = when (taskNo) {
                1 -> this.recur01()
                2 -> this.recur02()
                3 -> this.recur03()
                4 -> this.recur04()
                5 -> this.recur05()
                6 -> this.recur06()
                7 -> this.recur07()
                8 -> this.recur08()
                9 -> this.recur09()
                10 -> this.recur10()
                11 -> this.recur11()
                12 -> this.recur12()
                13 -> this.recur13()
                14 -> this.recur14()
                15 -> this.recur15()
                16 -> this.recur16()
                17 -> this.recur17()
                18 -> this.recur18()
                19 -> this.recur19()
                20 -> this.recur20()
                21 -> this.recur21()
                22 -> this.recur22()
                23 -> this.recur23()
                24 -> this.recur24()
                25 -> this.recur25()
                26 -> this.recur26()
                27 -> this.recur27()
                28 -> this.recur28()
                29 -> this.recur29()
                30 -> this.recur30()
                else -> false
            }
            scanner.removeFiles()
            if (!ok) {
                println("Recur$taskNo Test#$index has not been tested")
                return false
            }
            if (!datScanner.eof()) {
                println("Recur$taskNo Test#$index input is not enough")
                return false
            }
            if (!ansScanner.eof()) {
                println("Recur$taskNo Test#$index output is not enough")
                return false
            }
        }
        println("Recur$taskNo Tested!")
        return true
    }

    private fun recur01(): Boolean {
        return false
    }

    private fun recur02(): Boolean {
        return false
    }

    private fun recur03(): Boolean {
        return false
    }

    private fun recur04(): Boolean {
        return false
    }

    private fun recur05(): Boolean {
        return false
    }

    private fun recur06(): Boolean {
        return false
    }

    private fun recur07(): Boolean {
        return false
    }

    private fun recur08(): Boolean {
        return false
    }

    private fun recur09(): Boolean {
        return false
    }

    private fun recur10(): Boolean {
        return false
    }

    private fun recur11(): Boolean {
        return false
    }

    private fun recur12(): Boolean {
        return false
    }

    private fun recur13(): Boolean {
        return false
    }

    private fun recur14(): Boolean {
        return false
    }

    private fun recur15(): Boolean {
        return false
    }

    private fun recur16(): Boolean {
        return false
    }

    private fun recur17(): Boolean {
        return false
    }

    private fun recur18(): Boolean {
        return false
    }

    private fun recur19(): Boolean {
        return false
    }

    private fun recur20(): Boolean {
        return false
    }

    private fun recur21(): Boolean {
        return false
    }

    private fun recur22(): Boolean {
        return false
    }

    private fun recur23(): Boolean {
        return false
    }

    private fun recur24(): Boolean {
        return false
    }

    private fun recur25(): Boolean {
        return false
    }

    private fun recur26(): Boolean {
        return false
    }

    private fun recur27(): Boolean {
        return false
    }

    private fun recur28(): Boolean {
        return false
    }

    private fun recur29(): Boolean {
        return false
    }

    private fun recur30(): Boolean {
        return false
    }
}