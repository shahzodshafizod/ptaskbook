package pack11array

class ArrayClass {

    var datScanner = scanner.Scanner()
    var ansScanner = scanner.Scanner()

    fun solve(taskNo: Int): Boolean {
        val path = "src/pack11array/tests/Array$taskNo/"
        var fileName: String
        for (index in 1..100) {
            fileName = path + "%03d".format(index)
            if (!datScanner.changeReader("$fileName.dat")) {
                println("Array$taskNo couldn't open datScanner source")
                return false
            }
            if (!ansScanner.changeReader("$fileName.ans")) {
                println("Array$taskNo couldn't open ansScanner source")
                return false
            }
            val ok = when (taskNo) {
                1 -> this.array001()
                2 -> this.array002()
                3 -> this.array003()
                4 -> this.array004()
                5 -> this.array005()
                6 -> this.array006()
                7 -> this.array007()
                8 -> this.array008()
                9 -> this.array009()
                10 -> this.array010()
                11 -> this.array011()
                12 -> this.array012()
                13 -> this.array013()
                14 -> this.array014()
                15 -> this.array015()
                16 -> this.array016()
                17 -> this.array017()
                18 -> this.array018()
                19 -> this.array019()
                20 -> this.array020()
                21 -> this.array021()
                22 -> this.array022()
                23 -> this.array023()
                24 -> this.array024()
                25 -> this.array025()
                26 -> this.array026()
                27 -> this.array027()
                28 -> this.array028()
                29 -> this.array029()
                30 -> this.array030()
                31 -> this.array031()
                32 -> this.array032()
                33 -> this.array033()
                34 -> this.array034()
                35 -> this.array035()
                36 -> this.array036()
                37 -> this.array037()
                38 -> this.array038()
                39 -> this.array039()
                40 -> this.array040()
                41 -> this.array041()
                42 -> this.array042()
                43 -> this.array043()
                44 -> this.array044()
                45 -> this.array045()
                46 -> this.array046()
                47 -> this.array047()
                48 -> this.array048()
                49 -> this.array049()
                50 -> this.array050()
                51 -> this.array051()
                52 -> this.array052()
                53 -> this.array053()
                54 -> this.array054()
                55 -> this.array055()
                56 -> this.array056()
                57 -> this.array057()
                58 -> this.array058()
                59 -> this.array059()
                60 -> this.array060()
                61 -> this.array061()
                62 -> this.array062()
                63 -> this.array063()
                64 -> this.array064()
                65 -> this.array065()
                66 -> this.array066()
                67 -> this.array067()
                68 -> this.array068()
                69 -> this.array069()
                70 -> this.array070()
                71 -> this.array071()
                72 -> this.array072()
                73 -> this.array073()
                74 -> this.array074()
                75 -> this.array075()
                76 -> this.array076()
                77 -> this.array077()
                78 -> this.array078()
                79 -> this.array079()
                80 -> this.array080()
                81 -> this.array081()
                82 -> this.array082()
                83 -> this.array083()
                84 -> this.array084()
                85 -> this.array085()
                86 -> this.array086()
                87 -> this.array087()
                88 -> this.array088()
                89 -> this.array089()
                90 -> this.array090()
                91 -> this.array091()
                92 -> this.array092()
                93 -> this.array093()
                94 -> this.array094()
                95 -> this.array095()
                96 -> this.array096()
                97 -> this.array097()
                98 -> this.array098()
                99 -> this.array099()
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
                println("Array$taskNo Test#$index has not been tested")
                return false
            }
            if (!datScanner.eof()) {
                println("Array$taskNo Test#$index input is not enough")
                return false
            }
            if (!ansScanner.eof()) {
                println("Array$taskNo Test#$index output is not enough")
                return false
            }
        }
        println("Array$taskNo Tested!")
        return true
    }

    private fun array001(): Boolean {
        return false
    }

    private fun array002(): Boolean {
        return false
    }

    private fun array003(): Boolean {
        return false
    }

    private fun array004(): Boolean {
        return false
    }

    private fun array005(): Boolean {
        return false
    }

    private fun array006(): Boolean {
        return false
    }

    private fun array007(): Boolean {
        return false
    }

    private fun array008(): Boolean {
        return false
    }

    private fun array009(): Boolean {
        return false
    }

    private fun array010(): Boolean {
        return false
    }

    private fun array011(): Boolean {
        return false
    }

    private fun array012(): Boolean {
        return false
    }

    private fun array013(): Boolean {
        return false
    }

    private fun array014(): Boolean {
        return false
    }

    private fun array015(): Boolean {
        return false
    }

    private fun array016(): Boolean {
        return false
    }

    private fun array017(): Boolean {
        return false
    }

    private fun array018(): Boolean {
        return false
    }

    private fun array019(): Boolean {
        return false
    }

    private fun array020(): Boolean {
        return false
    }

    private fun array021(): Boolean {
        return false
    }

    private fun array022(): Boolean {
        return false
    }

    private fun array023(): Boolean {
        return false
    }

    private fun array024(): Boolean {
        return false
    }

    private fun array025(): Boolean {
        return false
    }

    private fun array026(): Boolean {
        return false
    }

    private fun array027(): Boolean {
        return false
    }

    private fun array028(): Boolean {
        return false
    }

    private fun array029(): Boolean {
        return false
    }

    private fun array030(): Boolean {
        return false
    }

    private fun array031(): Boolean {
        return false
    }

    private fun array032(): Boolean {
        return false
    }

    private fun array033(): Boolean {
        return false
    }

    private fun array034(): Boolean {
        return false
    }

    private fun array035(): Boolean {
        return false
    }

    private fun array036(): Boolean {
        return false
    }

    private fun array037(): Boolean {
        return false
    }

    private fun array038(): Boolean {
        return false
    }

    private fun array039(): Boolean {
        return false
    }

    private fun array040(): Boolean {
        return false
    }

    private fun array041(): Boolean {
        return false
    }

    private fun array042(): Boolean {
        return false
    }

    private fun array043(): Boolean {
        return false
    }

    private fun array044(): Boolean {
        return false
    }

    private fun array045(): Boolean {
        return false
    }

    private fun array046(): Boolean {
        return false
    }

    private fun array047(): Boolean {
        return false
    }

    private fun array048(): Boolean {
        return false
    }

    private fun array049(): Boolean {
        return false
    }

    private fun array050(): Boolean {
        return false
    }

    private fun array051(): Boolean {
        return false
    }

    private fun array052(): Boolean {
        return false
    }

    private fun array053(): Boolean {
        return false
    }

    private fun array054(): Boolean {
        return false
    }

    private fun array055(): Boolean {
        return false
    }

    private fun array056(): Boolean {
        return false
    }

    private fun array057(): Boolean {
        return false
    }

    private fun array058(): Boolean {
        return false
    }

    private fun array059(): Boolean {
        return false
    }

    private fun array060(): Boolean {
        return false
    }

    private fun array061(): Boolean {
        return false
    }

    private fun array062(): Boolean {
        return false
    }

    private fun array063(): Boolean {
        return false
    }

    private fun array064(): Boolean {
        return false
    }

    private fun array065(): Boolean {
        return false
    }

    private fun array066(): Boolean {
        return false
    }

    private fun array067(): Boolean {
        return false
    }

    private fun array068(): Boolean {
        return false
    }

    private fun array069(): Boolean {
        return false
    }

    private fun array070(): Boolean {
        return false
    }

    private fun array071(): Boolean {
        return false
    }

    private fun array072(): Boolean {
        return false
    }

    private fun array073(): Boolean {
        return false
    }

    private fun array074(): Boolean {
        return false
    }

    private fun array075(): Boolean {
        return false
    }

    private fun array076(): Boolean {
        return false
    }

    private fun array077(): Boolean {
        return false
    }

    private fun array078(): Boolean {
        return false
    }

    private fun array079(): Boolean {
        return false
    }

    private fun array080(): Boolean {
        return false
    }

    private fun array081(): Boolean {
        return false
    }

    private fun array082(): Boolean {
        return false
    }

    private fun array083(): Boolean {
        return false
    }

    private fun array084(): Boolean {
        return false
    }

    private fun array085(): Boolean {
        return false
    }

    private fun array086(): Boolean {
        return false
    }

    private fun array087(): Boolean {
        return false
    }

    private fun array088(): Boolean {
        return false
    }

    private fun array089(): Boolean {
        return false
    }

    private fun array090(): Boolean {
        return false
    }

    private fun array091(): Boolean {
        return false
    }

    private fun array092(): Boolean {
        return false
    }

    private fun array093(): Boolean {
        return false
    }

    private fun array094(): Boolean {
        return false
    }

    private fun array095(): Boolean {
        return false
    }

    private fun array096(): Boolean {
        return false
    }

    private fun array097(): Boolean {
        return false
    }

    private fun array098(): Boolean {
        return false
    }

    private fun array099(): Boolean {
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