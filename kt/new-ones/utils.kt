import java.io.*
import java.math.BigInteger
import java.security.MessageDigest
import java.util.*

/*
------------FIRST FILE--------------
33554432
67108864
100663296
134217728
------------SECOND FILE--------------
2
4
6
8
 */

/*
------------FIRST FILE--------------
-2.4393123941940724E19
3.0181448022870605E-204
1.21796063925348E43
7.688168999772515E284
-9.603721405543067E-86
2.9549381756109897E237
1.2896809572350312E-251
-2.342644233522792E-128
5.1197696022784334E-5
-6.424945546342464E123
------------SECOND FILE--------------
5.94
8.290000000000001
10.64
12.99
15.34
17.69
20.040000000000003
22.390000000000004
24.740000000000006
27.090000000000007
 */

fun main() {
//    readDouble()
//    readInt()
//    file2()
//    checkFile()
    val str = "┌───521───────┐"
    var dig = ""
    str.forEach { chr -> if (chr.isDigit()) { dig += chr.toString() }}
    println(dig.toInt())
}

fun checkFile() {
    val filename = "s4dlzpa4.tst"
    println(getCheckSum(File(filename).readBytes()) == getCheckSum(File("$filename.shw").readBytes()))
}

fun file2() {
    val filename = "s4dlzpa4.tst.shw"
    val scanner = Scanner(File("dat.txt").bufferedReader())
    val ra = RandomAccessFile(filename, "rw")
    scanner.nextLine()
    val limit: Int = scanner.nextInt()
    var number = 2
    for (index in 1..limit) {
        ra.writeInt(number)
        number += 2
    }
    ra.close()
}

fun readInt() {
    println("------------FIRST FILE--------------")
    var filename = "s4dlzpa4.tst"
    readIntFile(filename)
    return

    println("------------SECOND FILE--------------")
    filename = "int/7639ie1x.tst.shw"
    readIntFile(filename)
}

fun readDouble() {
    println("------------FIRST FILE--------------")
    var filename = "double/vcloj0x4.tst"
    readDoubleFile(filename)

    println("------------SECOND FILE--------------")
    filename = "double/vcloj0x4.tst.shw"
    readDoubleFile(filename)
}

fun readIntFile(filename: String) {
    val ra = RandomAccessFile(filename, "r")
    var number: Int
    do {
        number = ra.readInt()
        println(number)
    } while (number != -1)
    ra.close()
}

fun readDoubleFile(filename: String) {
    val dis = DataInputStream(FileInputStream(File(filename)))
    for (index in 1..10) {
        println(dis.readDouble())
    }
    dis.close()
}

fun nextFilename(fname: String): String {
    val path = File(fname).absolutePath.substringBeforeLast('/') + "/"
    var filename = "$path$fname"
    val rd = File(filename).bufferedReader()
    val wr = File("$filename.shw").bufferedWriter()
    wr.write(rd.readText())
    return "$filename.shw"
}

fun getCheckSum(data: ByteArray): String {
    val hashed = MessageDigest.getInstance("SHA-256").digest(data)
    return String.format("%032x", BigInteger(1, hashed))
}
