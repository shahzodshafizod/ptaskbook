package pack19tree

class TreeClass {

    var datScanner = scanner.Scanner()
    var ansScanner = scanner.Scanner()

    fun solve(taskNo: Int): Boolean {
        val path = "src/pack19tree/tests/Tree$taskNo/"
        var fileName: String
        for (index in 1..100) {
            fileName = path + "%03d".format(index)
            if (!datScanner.changeReader("$fileName.dat")) {
                println("Tree$taskNo couldn't open datScanner source")
                return false
            }
            if (!ansScanner.changeReader("$fileName.ans")) {
                println("Tree$taskNo couldn't open ansScanner source")
                return false
            }
            val ok = when (taskNo) {
                1 -> this.tree001()
                2 -> this.tree002()
                3 -> this.tree003()
                4 -> this.tree004()
                5 -> this.tree005()
                6 -> this.tree006()
                7 -> this.tree007()
                8 -> this.tree008()
                9 -> this.tree009()
                10 -> this.tree010()
                11 -> this.tree011()
                12 -> this.tree012()
                13 -> this.tree013()
                14 -> this.tree014()
                15 -> this.tree015()
                16 -> this.tree016()
                17 -> this.tree017()
                18 -> this.tree018()
                19 -> this.tree019()
                20 -> this.tree020()
                21 -> this.tree021()
                22 -> this.tree022()
                23 -> this.tree023()
                24 -> this.tree024()
                25 -> this.tree025()
                26 -> this.tree026()
                27 -> this.tree027()
                28 -> this.tree028()
                29 -> this.tree029()
                30 -> this.tree030()
                31 -> this.tree031()
                32 -> this.tree032()
                33 -> this.tree033()
                34 -> this.tree034()
                35 -> this.tree035()
                36 -> this.tree036()
                37 -> this.tree037()
                38 -> this.tree038()
                39 -> this.tree039()
                40 -> this.tree040()
                41 -> this.tree041()
                42 -> this.tree042()
                43 -> this.tree043()
                44 -> this.tree044()
                45 -> this.tree045()
                46 -> this.tree046()
                47 -> this.tree047()
                48 -> this.tree048()
                49 -> this.tree049()
                50 -> this.tree050()
                51 -> this.tree051()
                52 -> this.tree052()
                53 -> this.tree053()
                54 -> this.tree054()
                55 -> this.tree055()
                56 -> this.tree056()
                57 -> this.tree057()
                58 -> this.tree058()
                59 -> this.tree059()
                60 -> this.tree060()
                61 -> this.tree061()
                62 -> this.tree062()
                63 -> this.tree063()
                64 -> this.tree064()
                65 -> this.tree065()
                66 -> this.tree066()
                67 -> this.tree067()
                68 -> this.tree068()
                69 -> this.tree069()
                70 -> this.tree070()
                71 -> this.tree071()
                72 -> this.tree072()
                73 -> this.tree073()
                74 -> this.tree074()
                75 -> this.tree075()
                76 -> this.tree076()
                77 -> this.tree077()
                78 -> this.tree078()
                79 -> this.tree079()
                80 -> this.tree080()
                81 -> this.tree081()
                82 -> this.tree082()
                83 -> this.tree083()
                84 -> this.tree084()
                85 -> this.tree085()
                86 -> this.tree086()
                87 -> this.tree087()
                88 -> this.tree088()
                89 -> this.tree089()
                90 -> this.tree090()
                91 -> this.tree091()
                92 -> this.tree092()
                93 -> this.tree093()
                94 -> this.tree094()
                95 -> this.tree095()
                96 -> this.tree096()
                97 -> this.tree097()
                98 -> this.tree098()
                99 -> this.tree099()
                100 -> this.tree100()
                else -> false
            }
            scanner.removeFiles()
            if (!ok) {
                println("Tree$taskNo Test#$index has not been tested")
                return false
            }
            if (!datScanner.eof()) {
                println("Tree$taskNo Test#$index input is not enough")
                return false
            }
            if (!ansScanner.eof()) {
                println("Tree$taskNo Test#$index output is not enough")
                return false
            }
        }
        println("Tree$taskNo Tested!")
        return true
    }

    private fun tree001(): Boolean {
        return false
    }

    private fun tree002(): Boolean {
        return false
    }

    private fun tree003(): Boolean {
        return false
    }

    private fun tree004(): Boolean {
        return false
    }

    private fun tree005(): Boolean {
        return false
    }

    private fun tree006(): Boolean {
        return false
    }

    private fun tree007(): Boolean {
        return false
    }

    private fun tree008(): Boolean {
        return false
    }

    private fun tree009(): Boolean {
        return false
    }

    private fun tree010(): Boolean {
        return false
    }

    private fun tree011(): Boolean {
        return false
    }

    private fun tree012(): Boolean {
        return false
    }

    private fun tree013(): Boolean {
        return false
    }

    private fun tree014(): Boolean {
        return false
    }

    private fun tree015(): Boolean {
        return false
    }

    private fun tree016(): Boolean {
        return false
    }

    private fun tree017(): Boolean {
        return false
    }

    private fun tree018(): Boolean {
        return false
    }

    private fun tree019(): Boolean {
        return false
    }

    private fun tree020(): Boolean {
        return false
    }

    private fun tree021(): Boolean {
        return false
    }

    private fun tree022(): Boolean {
        return false
    }

    private fun tree023(): Boolean {
        return false
    }

    private fun tree024(): Boolean {
        return false
    }

    private fun tree025(): Boolean {
        return false
    }

    private fun tree026(): Boolean {
        return false
    }

    private fun tree027(): Boolean {
        return false
    }

    private fun tree028(): Boolean {
        return false
    }

    private fun tree029(): Boolean {
        return false
    }

    private fun tree030(): Boolean {
        return false
    }

    private fun tree031(): Boolean {
        return false
    }

    private fun tree032(): Boolean {
        return false
    }

    private fun tree033(): Boolean {
        return false
    }

    private fun tree034(): Boolean {
        return false
    }

    private fun tree035(): Boolean {
        return false
    }

    private fun tree036(): Boolean {
        return false
    }

    private fun tree037(): Boolean {
        return false
    }

    private fun tree038(): Boolean {
        return false
    }

    private fun tree039(): Boolean {
        return false
    }

    private fun tree040(): Boolean {
        return false
    }

    private fun tree041(): Boolean {
        return false
    }

    private fun tree042(): Boolean {
        return false
    }

    private fun tree043(): Boolean {
        return false
    }

    private fun tree044(): Boolean {
        return false
    }

    private fun tree045(): Boolean {
        return false
    }

    private fun tree046(): Boolean {
        return false
    }

    private fun tree047(): Boolean {
        return false
    }

    private fun tree048(): Boolean {
        return false
    }

    private fun tree049(): Boolean {
        return false
    }

    private fun tree050(): Boolean {
        return false
    }

    private fun tree051(): Boolean {
        return false
    }

    private fun tree052(): Boolean {
        return false
    }

    private fun tree053(): Boolean {
        return false
    }

    private fun tree054(): Boolean {
        return false
    }

    private fun tree055(): Boolean {
        return false
    }

    private fun tree056(): Boolean {
        return false
    }

    private fun tree057(): Boolean {
        return false
    }

    private fun tree058(): Boolean {
        return false
    }

    private fun tree059(): Boolean {
        return false
    }

    private fun tree060(): Boolean {
        return false
    }

    private fun tree061(): Boolean {
        return false
    }

    private fun tree062(): Boolean {
        return false
    }

    private fun tree063(): Boolean {
        return false
    }

    private fun tree064(): Boolean {
        return false
    }

    private fun tree065(): Boolean {
        return false
    }

    private fun tree066(): Boolean {
        return false
    }

    private fun tree067(): Boolean {
        return false
    }

    private fun tree068(): Boolean {
        return false
    }

    private fun tree069(): Boolean {
        return false
    }

    private fun tree070(): Boolean {
        return false
    }

    private fun tree071(): Boolean {
        return false
    }

    private fun tree072(): Boolean {
        return false
    }

    private fun tree073(): Boolean {
        return false
    }

    private fun tree074(): Boolean {
        return false
    }

    private fun tree075(): Boolean {
        return false
    }

    private fun tree076(): Boolean {
        return false
    }

    private fun tree077(): Boolean {
        return false
    }

    private fun tree078(): Boolean {
        return false
    }

    private fun tree079(): Boolean {
        return false
    }

    private fun tree080(): Boolean {
        return false
    }

    private fun tree081(): Boolean {
        return false
    }

    private fun tree082(): Boolean {
        return false
    }

    private fun tree083(): Boolean {
        return false
    }

    private fun tree084(): Boolean {
        return false
    }

    private fun tree085(): Boolean {
        return false
    }

    private fun tree086(): Boolean {
        return false
    }

    private fun tree087(): Boolean {
        return false
    }

    private fun tree088(): Boolean {
        return false
    }

    private fun tree089(): Boolean {
        return false
    }

    private fun tree090(): Boolean {
        return false
    }

    private fun tree091(): Boolean {
        return false
    }

    private fun tree092(): Boolean {
        return false
    }

    private fun tree093(): Boolean {
        return false
    }

    private fun tree094(): Boolean {
        return false
    }

    private fun tree095(): Boolean {
        return false
    }

    private fun tree096(): Boolean {
        return false
    }

    private fun tree097(): Boolean {
        return false
    }

    private fun tree098(): Boolean {
        return false
    }

    private fun tree099(): Boolean {
        return false
    }

    private fun tree100(): Boolean {
        return false
    }
}