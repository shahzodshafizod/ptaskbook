package pack12matrix

class MatrixClass {

    var datScanner = scanner.Scanner()
    var ansScanner = scanner.Scanner()

    fun solve(taskNo: Int): Boolean {
        val path = "src/pack12matrix/tests/Matrix$taskNo/"
        var fileName: String
        for (index in 1..100) {
            fileName = path + "%03d".format(index)
            if (!datScanner.changeReader("$fileName.dat")) {
                println("Matrix$taskNo couldn't open datScanner source")
                return false
            }
            if (!ansScanner.changeReader("$fileName.ans")) {
                println("Matrix$taskNo couldn't open ansScanner source")
                return false
            }
            val ok = when (taskNo) {
                1 -> this.matrix001()
                2 -> this.matrix002()
                3 -> this.matrix003()
                4 -> this.matrix004()
                5 -> this.matrix005()
                6 -> this.matrix006()
                7 -> this.matrix007()
                8 -> this.matrix008()
                9 -> this.matrix009()
                10 -> this.matrix010()
                11 -> this.matrix011()
                12 -> this.matrix012()
                13 -> this.matrix013()
                14 -> this.matrix014()
                15 -> this.matrix015()
                16 -> this.matrix016()
                17 -> this.matrix017()
                18 -> this.matrix018()
                19 -> this.matrix019()
                20 -> this.matrix020()
                21 -> this.matrix021()
                22 -> this.matrix022()
                23 -> this.matrix023()
                24 -> this.matrix024()
                25 -> this.matrix025()
                26 -> this.matrix026()
                27 -> this.matrix027()
                28 -> this.matrix028()
                29 -> this.matrix029()
                30 -> this.matrix030()
                31 -> this.matrix031()
                32 -> this.matrix032()
                33 -> this.matrix033()
                34 -> this.matrix034()
                35 -> this.matrix035()
                36 -> this.matrix036()
                37 -> this.matrix037()
                38 -> this.matrix038()
                39 -> this.matrix039()
                40 -> this.matrix040()
                41 -> this.matrix041()
                42 -> this.matrix042()
                43 -> this.matrix043()
                44 -> this.matrix044()
                45 -> this.matrix045()
                46 -> this.matrix046()
                47 -> this.matrix047()
                48 -> this.matrix048()
                49 -> this.matrix049()
                50 -> this.matrix050()
                51 -> this.matrix051()
                52 -> this.matrix052()
                53 -> this.matrix053()
                54 -> this.matrix054()
                55 -> this.matrix055()
                56 -> this.matrix056()
                57 -> this.matrix057()
                58 -> this.matrix058()
                59 -> this.matrix059()
                60 -> this.matrix060()
                61 -> this.matrix061()
                62 -> this.matrix062()
                63 -> this.matrix063()
                64 -> this.matrix064()
                65 -> this.matrix065()
                66 -> this.matrix066()
                67 -> this.matrix067()
                68 -> this.matrix068()
                69 -> this.matrix069()
                70 -> this.matrix070()
                71 -> this.matrix071()
                72 -> this.matrix072()
                73 -> this.matrix073()
                74 -> this.matrix074()
                75 -> this.matrix075()
                76 -> this.matrix076()
                77 -> this.matrix077()
                78 -> this.matrix078()
                79 -> this.matrix079()
                80 -> this.matrix080()
                81 -> this.matrix081()
                82 -> this.matrix082()
                83 -> this.matrix083()
                84 -> this.matrix084()
                85 -> this.matrix085()
                86 -> this.matrix086()
                87 -> this.matrix087()
                88 -> this.matrix088()
                89 -> this.matrix089()
                90 -> this.matrix090()
                91 -> this.matrix091()
                92 -> this.matrix092()
                93 -> this.matrix093()
                94 -> this.matrix094()
                95 -> this.matrix095()
                96 -> this.matrix096()
                97 -> this.matrix097()
                98 -> this.matrix098()
                99 -> this.matrix099()
                100 -> this.matrix100()
                else -> false
            }
            scanner.removeFiles()
            if (!ok) {
                println("Matrix$taskNo Test#$index has not been tested")
                return false
            }
            if (!datScanner.eof()) {
                println("Matrix$taskNo Test#$index input is not enough")
                return false
            }
            if (!ansScanner.eof()) {
                println("Matrix$taskNo Test#$index output is not enough")
                return false
            }
        }
        println("Matrix$taskNo Tested!")
        return true
    }

    private fun matrix001(): Boolean {
        return false
    }

    private fun matrix002(): Boolean {
        return false
    }

    private fun matrix003(): Boolean {
        return false
    }

    private fun matrix004(): Boolean {
        return false
    }

    private fun matrix005(): Boolean {
        return false
    }

    private fun matrix006(): Boolean {
        return false
    }

    private fun matrix007(): Boolean {
        return false
    }

    private fun matrix008(): Boolean {
        return false
    }

    private fun matrix009(): Boolean {
        return false
    }

    private fun matrix010(): Boolean {
        return false
    }

    private fun matrix011(): Boolean {
        return false
    }

    private fun matrix012(): Boolean {
        return false
    }

    private fun matrix013(): Boolean {
        return false
    }

    private fun matrix014(): Boolean {
        return false
    }

    private fun matrix015(): Boolean {
        return false
    }

    private fun matrix016(): Boolean {
        return false
    }

    private fun matrix017(): Boolean {
        return false
    }

    private fun matrix018(): Boolean {
        return false
    }

    private fun matrix019(): Boolean {
        return false
    }

    private fun matrix020(): Boolean {
        return false
    }

    private fun matrix021(): Boolean {
        return false
    }

    private fun matrix022(): Boolean {
        return false
    }

    private fun matrix023(): Boolean {
        return false
    }

    private fun matrix024(): Boolean {
        return false
    }

    private fun matrix025(): Boolean {
        return false
    }

    private fun matrix026(): Boolean {
        return false
    }

    private fun matrix027(): Boolean {
        return false
    }

    private fun matrix028(): Boolean {
        return false
    }

    private fun matrix029(): Boolean {
        return false
    }

    private fun matrix030(): Boolean {
        return false
    }

    private fun matrix031(): Boolean {
        return false
    }

    private fun matrix032(): Boolean {
        return false
    }

    private fun matrix033(): Boolean {
        return false
    }

    private fun matrix034(): Boolean {
        return false
    }

    private fun matrix035(): Boolean {
        return false
    }

    private fun matrix036(): Boolean {
        return false
    }

    private fun matrix037(): Boolean {
        return false
    }

    private fun matrix038(): Boolean {
        return false
    }

    private fun matrix039(): Boolean {
        return false
    }

    private fun matrix040(): Boolean {
        return false
    }

    private fun matrix041(): Boolean {
        return false
    }

    private fun matrix042(): Boolean {
        return false
    }

    private fun matrix043(): Boolean {
        return false
    }

    private fun matrix044(): Boolean {
        return false
    }

    private fun matrix045(): Boolean {
        return false
    }

    private fun matrix046(): Boolean {
        return false
    }

    private fun matrix047(): Boolean {
        return false
    }

    private fun matrix048(): Boolean {
        return false
    }

    private fun matrix049(): Boolean {
        return false
    }

    private fun matrix050(): Boolean {
        return false
    }

    private fun matrix051(): Boolean {
        return false
    }

    private fun matrix052(): Boolean {
        return false
    }

    private fun matrix053(): Boolean {
        return false
    }

    private fun matrix054(): Boolean {
        return false
    }

    private fun matrix055(): Boolean {
        return false
    }

    private fun matrix056(): Boolean {
        return false
    }

    private fun matrix057(): Boolean {
        return false
    }

    private fun matrix058(): Boolean {
        return false
    }

    private fun matrix059(): Boolean {
        return false
    }

    private fun matrix060(): Boolean {
        return false
    }

    private fun matrix061(): Boolean {
        return false
    }

    private fun matrix062(): Boolean {
        return false
    }

    private fun matrix063(): Boolean {
        return false
    }

    private fun matrix064(): Boolean {
        return false
    }

    private fun matrix065(): Boolean {
        return false
    }

    private fun matrix066(): Boolean {
        return false
    }

    private fun matrix067(): Boolean {
        return false
    }

    private fun matrix068(): Boolean {
        return false
    }

    private fun matrix069(): Boolean {
        return false
    }

    private fun matrix070(): Boolean {
        return false
    }

    private fun matrix071(): Boolean {
        return false
    }

    private fun matrix072(): Boolean {
        return false
    }

    private fun matrix073(): Boolean {
        return false
    }

    private fun matrix074(): Boolean {
        return false
    }

    private fun matrix075(): Boolean {
        return false
    }

    private fun matrix076(): Boolean {
        return false
    }

    private fun matrix077(): Boolean {
        return false
    }

    private fun matrix078(): Boolean {
        return false
    }

    private fun matrix079(): Boolean {
        return false
    }

    private fun matrix080(): Boolean {
        return false
    }

    private fun matrix081(): Boolean {
        return false
    }

    private fun matrix082(): Boolean {
        return false
    }

    private fun matrix083(): Boolean {
        return false
    }

    private fun matrix084(): Boolean {
        return false
    }

    private fun matrix085(): Boolean {
        return false
    }

    private fun matrix086(): Boolean {
        return false
    }

    private fun matrix087(): Boolean {
        return false
    }

    private fun matrix088(): Boolean {
        return false
    }

    private fun matrix089(): Boolean {
        return false
    }

    private fun matrix090(): Boolean {
        return false
    }

    private fun matrix091(): Boolean {
        return false
    }

    private fun matrix092(): Boolean {
        return false
    }

    private fun matrix093(): Boolean {
        return false
    }

    private fun matrix094(): Boolean {
        return false
    }

    private fun matrix095(): Boolean {
        return false
    }

    private fun matrix096(): Boolean {
        return false
    }

    private fun matrix097(): Boolean {
        return false
    }

    private fun matrix098(): Boolean {
        return false
    }

    private fun matrix099(): Boolean {
        return false
    }

    private fun matrix100(): Boolean {
        return false
    }
}