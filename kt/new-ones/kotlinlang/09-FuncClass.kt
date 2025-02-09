package pack09proc

class FuncClass {

    var datScanner = scanner.Scanner()
    var ansScanner = scanner.Scanner()

    fun solve(taskNo: Int): Boolean {
        val path = "src/pack09proc/tests/Func$taskNo/"
        var fileName: String
        for (index in 1..100) {
            fileName = path + "%03d".format(index)
            if (!datScanner.changeReader("$fileName.dat")) {
                println("Func$taskNo couldn't open datScanner source")
                return false
            }
            if (!ansScanner.changeReader("$fileName.ans")) {
                println("Func$taskNo couldn't open ansScanner source")
                return false
            }
            val ok = when (taskNo) {
                1 -> this.func01()
                2 -> this.func02()
                3 -> this.func03()
                4 -> this.func04()
                5 -> this.func05()
                6 -> this.func06()
                7 -> this.func07()
                8 -> this.func08()
                9 -> this.func09()
                10 -> this.func10()
                11 -> this.func11()
                12 -> this.func12()
                13 -> this.func13()
                14 -> this.func14()
                15 -> this.func15()
                16 -> this.func16()
                17 -> this.func17()
                18 -> this.func18()
                19 -> this.func19()
                20 -> this.func20()
                21 -> this.func21()
                22 -> this.func22()
                23 -> this.func23()
                24 -> this.func24()
                25 -> this.func25()
                26 -> this.func26()
                27 -> this.func27()
                28 -> this.func28()
                29 -> this.func29()
                30 -> this.func30()
                31 -> this.func31()
                32 -> this.func32()
                33 -> this.func33()
                34 -> this.func34()
                35 -> this.func35()
                36 -> this.func36()
                37 -> this.func37()
                38 -> this.func38()
                39 -> this.func39()
                40 -> this.func40()
                41 -> this.func41()
                42 -> this.func42()
                43 -> this.func43()
                44 -> this.func44()
                45 -> this.func45()
                46 -> this.func46()
                47 -> this.func47()
                48 -> this.func48()
                49 -> this.func49()
                50 -> this.func50()
                51 -> this.func51()
                52 -> this.func52()
                53 -> this.func53()
                54 -> this.func54()
                55 -> this.func55()
                56 -> this.func56()
                57 -> this.func57()
                58 -> this.func58()
                59 -> this.func59()
                60 -> this.func60()
                else -> false
            }
            scanner.removeFiles()
            if (!ok) {
                println("Func$taskNo Test#$index has not been tested")
                return false
            }
            if (!datScanner.eof()) {
                println("Func$taskNo Test#$index input is not enough")
                return false
            }
            if (!ansScanner.eof()) {
                println("Func$taskNo Test#$index output is not enough")
                return false
            }
        }
        println("Func$taskNo Tested!")
        return true
    }

    private fun func01(): Boolean {
        return false
    }

    private fun func02(): Boolean {
        return false
    }

    private fun func03(): Boolean {
        return false
    }

    private fun func04(): Boolean {
        return false
    }

    private fun func05(): Boolean {
        return false
    }

    private fun func06(): Boolean {
        return false
    }

    private fun func07(): Boolean {
        return false
    }

    private fun func08(): Boolean {
        return false
    }

    private fun func09(): Boolean {
        return false
    }

    private fun func10(): Boolean {
        return false
    }

    private fun func11(): Boolean {
        return false
    }

    private fun func12(): Boolean {
        return false
    }

    private fun func13(): Boolean {
        return false
    }

    private fun func14(): Boolean {
        return false
    }

    private fun func15(): Boolean {
        return false
    }

    private fun func16(): Boolean {
        return false
    }

    private fun func17(): Boolean {
        return false
    }

    private fun func18(): Boolean {
        return false
    }

    private fun func19(): Boolean {
        return false
    }

    private fun func20(): Boolean {
        return false
    }

    private fun func21(): Boolean {
        return false
    }

    private fun func22(): Boolean {
        return false
    }

    private fun func23(): Boolean {
        return false
    }

    private fun func24(): Boolean {
        return false
    }

    private fun func25(): Boolean {
        return false
    }

    private fun func26(): Boolean {
        return false
    }

    private fun func27(): Boolean {
        return false
    }

    private fun func28(): Boolean {
        return false
    }

    private fun func29(): Boolean {
        return false
    }

    private fun func30(): Boolean {
        return false
    }

    private fun func31(): Boolean {
        return false
    }

    private fun func32(): Boolean {
        return false
    }

    private fun func33(): Boolean {
        return false
    }

    private fun func34(): Boolean {
        return false
    }

    private fun func35(): Boolean {
        return false
    }

    private fun func36(): Boolean {
        return false
    }

    private fun func37(): Boolean {
        return false
    }

    private fun func38(): Boolean {
        return false
    }

    private fun func39(): Boolean {
        return false
    }

    private fun func40(): Boolean {
        return false
    }

    private fun func41(): Boolean {
        return false
    }

    private fun func42(): Boolean {
        return false
    }

    private fun func43(): Boolean {
        return false
    }

    private fun func44(): Boolean {
        return false
    }

    private fun func45(): Boolean {
        return false
    }

    private fun func46(): Boolean {
        return false
    }

    private fun func47(): Boolean {
        return false
    }

    private fun func48(): Boolean {
        return false
    }

    private fun func49(): Boolean {
        return false
    }

    private fun func50(): Boolean {
        return false
    }

    private fun func51(): Boolean {
        return false
    }

    private fun func52(): Boolean {
        return false
    }

    private fun func53(): Boolean {
        return false
    }

    private fun func54(): Boolean {
        return false
    }

    private fun func55(): Boolean {
        return false
    }

    private fun func56(): Boolean {
        return false
    }

    private fun func57(): Boolean {
        return false
    }

    private fun func58(): Boolean {
        return false
    }

    private fun func59(): Boolean {
        return false
    }

    private fun func60(): Boolean {
        return false
    }
}