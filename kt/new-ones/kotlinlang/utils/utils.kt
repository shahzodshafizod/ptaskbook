package utils

import java.io.File
import java.math.BigInteger
import java.security.MessageDigest

fun getCheckSum(fileName: String): String {
    val file = File(fileName)
    if (!file.exists()) {
        println("getCheckSum $fileName does not exists")
        return ""
    }
    val data = file.readBytes()
    val hashed = MessageDigest.getInstance("SHA-256").digest(data)
    return String.format("%032x", BigInteger(1, hashed))
}