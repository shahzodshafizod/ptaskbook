package scanner

import utils.*
import utils.List
import java.io.File
import java.nio.ByteOrder
import java.util.Scanner
import kotlin.math.*

private var filesForRemove = arrayListOf<String>()

class Scanner {

    private var scanner: Scanner? = null
    private var scanPath: String = ""

    fun changeReader(filename: String): Boolean {
        val file = File(filename)
        if (!file.exists()) {
            return false
        }
        this.scanPath = file.parentFile.absolutePath
        this.scanner = Scanner(file.bufferedReader())
        return true
    }

    fun eof(): Boolean {
        return !this.scanner!!.hasNext()
    }

    fun skip() {
        this.scanner!!.skip(".")
    }

    fun nextWord(): String {
        if (this.eof()) {
            return ""
        }
        return this.scanner!!.next()
    }

    fun checkWord(vararg args: String): Boolean {
        var value: String
        for (arg in args) {
            if (this.eof()) {
                println("output EOF")
                return false
            }
            value = this.nextWord()
            if (value != arg) {
                println("expected: $value")
                println("have got: $arg")
                return false
            }
        }
        return true
    }

    fun nextLine(): String {
        if (!this.scanner!!.hasNextLine()) {
            return ""
        }
        return this.scanner!!.nextLine()
    }

    fun checkLine(vararg args: String): Boolean {
        var value: String
        for (arg in args) {
            if (!this.scanner!!.hasNextLine()) {
                println("output EOF")
                return false
            }
            value = this.nextLine()
            if (value != arg) {
                println("expected: $value")
                println("have got: $arg")
                return false
            }
        }
        return true
    }

    fun nextChar(): Char {
        var chr = ' '
        while (!this.eof()) {
            chr = this.scanner!!.next()[0]
            if (chr != '\n') {
                break
            }
        }
        return chr
    }

    fun checkChar(vararg args: Char): Boolean {
        var value: Char
        for (arg in args) {
            if (this.eof()) {
                println("output EOF")
                return false
            }
            value = this.nextChar()
            if (value != arg) {
                println("expected: $value")
                println("have got: $arg")
                return false
            }
        }
        return true
    }

    fun nextBoolean(): Boolean {
        if (!this.scanner!!.hasNextBoolean()) {
            return false
        }
        return this.scanner!!.nextBoolean()
    }

    fun checkBoolean(vararg args: Boolean): Boolean {
        var value: Boolean
        for (arg in args) {
            if (!this.scanner!!.hasNextBoolean()) {
                println("output EOF")
                return false
            }
            value = this.nextBoolean()
            if (value != arg) {
                println("expected: $value")
                println("have got: $arg")
                return false
            }
        }
        return true
    }

    fun nextInt(): Int {
        if (!this.scanner!!.hasNextInt()) {
            return 0
        }
        return this.scanner!!.nextInt()
    }

    fun checkInt(vararg args: Int): Boolean {
        var value: Int
        for (arg in args) {
            if (!this.scanner!!.hasNextInt()) {
                println("output EOF")
                return false
            }
            value = this.nextInt()
            if (value != arg) {
                println("expected: $value")
                println("have got: $arg")
                return false
            }
        }
        return true
    }

    fun nextFloat(): Float {
        if (!this.scanner!!.hasNextFloat()) {
            return 0F
        }
        return this.scanner!!.nextFloat()
    }

    fun checkFloat(precision: Int, vararg args: Float): Boolean {
        var value: Float
        for (arg in args) {
            if (!this.scanner!!.hasNextFloat()) {
                println("output EOF")
                return false
            }
            value = this.nextFloat().round(precision)
            if (value != arg.round(precision)) {
                println("expected: $value")
                println("have got: $arg")
                return false
            }
        }
        return true
    }

    fun nextDouble(): Double {
        if (!this.scanner!!.hasNextDouble()) {
            return 0.0
        }
        return this.scanner!!.nextDouble()
    }

    fun checkDouble(precision: Int, vararg args: Double): Boolean {
        var value: Double
        for (arg in args) {
            if (!this.scanner!!.hasNextDouble()) {
                println("output EOF")
                return false
            }
            value = this.nextDouble().round(precision)
            if (value != arg.round(precision)) {
                println("expected: $value")
                println("have got: $arg")
                return false
            }
        }
        return true
    }

    fun nextBinaryFileName(endian: ByteOrder): String? {
        var fileName = this.nextLine()
        if (endian == ByteOrder.BIG_ENDIAN) {
            fileName += ".big"
        }
        val fromFileName = this.scanPath + "/" + fileName
        val fromFile = File(fromFileName)
        if (!fromFile.exists()) {
            println("doesn't exist the file: $fromFileName")
            return null
        }
        val toFile = File(fileName)
        if (!toFile.createNewFile()) {
            println("couldn't create the file: $fileName")
            return null
        }
        fromFile.copyTo(toFile, true)
        addToRemoveQueue(fileName)
        return fileName
    }

    fun checkBinaryFile(endian: ByteOrder, vararg fileNames: String): Boolean {
        for (fileName in fileNames) {
            if (!File(fileName).exists()) {
                println("File $fileName doesn't exists")
                continue
            }
            addToRemoveQueue(fileName)
        }
        for (fileName in fileNames) {
            if (!File(fileName).exists()) {
                continue
            }
            if (this.eof()) {
                if (fileName.isEmpty()) {
                    continue
                }
                println("output EOF")
                return false
            }
            var correctFileName = this.scanPath + "/" + this.nextLine()
            if (endian == ByteOrder.BIG_ENDIAN) {
                correctFileName += ".big"
            }
            val correctChs = utils.getCheckSum(correctFileName)
            val checkChs2 = utils.getCheckSum(fileName)
            // check with hashed checkSums;
            if (checkChs2 != correctChs) {
                println("files $fileName and $correctFileName have different contents")
                return false
            }
        }
        return true
    }

    fun nextTextFileName(): String? {
        val fileName = this.nextLine()
        val fromFile = File(scanPath + "/" + fileName)
        if (!fromFile.exists()) {
            return null
        }
        val toFile = File(fileName)
        if (!toFile.createNewFile()) {
            return null
        }
        fromFile.copyTo(toFile, true)
        addToRemoveQueue(fileName)
        return fileName
    }

    fun checkTextFile(vararg fileNames: String): Boolean {
        for (fileName in fileNames) {
            if (!File(fileName).exists()) {
                println("File $fileName doesn't exists")
                continue
            }
            addToRemoveQueue(fileName)
        }
        for (fileName in fileNames) {
            if (this.eof()) {
                if (fileName.isEmpty()) {
                    continue
                }
                println("output EOF")
                return false
            }
            val correctScanner = Scanner(File(this.scanPath + "/" + this.nextLine()))
            val checkScanner = Scanner(File(fileName))
            while (true) {
                if (!correctScanner.hasNextLine() || !checkScanner.hasNextLine()) {
                    break
                }
                val correctLine = correctScanner.nextLine()
                val checkLine = checkScanner.nextLine()
                if (correctLine != checkLine) {
                    println("expected: '$correctLine'")
                    println("have got: '$checkLine'")
                    return false
                }
            }
            if (correctScanner.hasNextLine()) {
                println("Correct file is not EOF")
                return false
            }
            if (checkScanner.hasNextLine()) {
                println("Check file is not EOF")
                return false
            }
        }
        return true
    }

    fun nextStack(): Stack {
        val count =  this.nextInt()
        val stack = Stack()
        for (index in 1..count) {
            stack.top = Node(data = this.nextInt(), next = stack.top)
        }
        return stack
    }

    fun checkStack(vararg stacks: Stack): Boolean {
        var correctStack: Stack
        var checkStack: Stack
        for (stack in stacks) {
            if (this.eof()) {
                println("output EOF")
                return false
            }
            correctStack = this.nextStack()
            checkStack = stack
            while (checkStack.top != null && correctStack.top != null) {
                if (checkStack.top!!.data != correctStack.top!!.data) {
                    println("expected: ${correctStack.top!!.data}")
                    println("have got: ${checkStack.top!!.data}")
                    return false
                }
                checkStack.top = checkStack.top!!.next
                correctStack.top = correctStack.top!!.next
            }
            if (checkStack.top != null) {
                println("Check stack is not empty")
                return false
            }
            if (correctStack.top != null) {
                println("Correct stack is not empty")
                return false
            }
        }
        return true
    }

    fun nextQueue(): Queue {
        val count = this.nextInt()
        val queue = Queue()
        var newNode: Node?
        for (index in 1..count) {
            newNode = Node(this.nextInt())
            if (queue.tail != null) {
                queue.tail!!.next = newNode
            }
            queue.tail = newNode
            if (queue.head == null) {
                queue.head = newNode
            }
        }
        return queue
    }

    fun checkQueue(vararg queues: Queue): Boolean {
        var correctQueue: Queue
        var checkQueue: Queue
        for (queue in queues) {
            if (this.eof()) {
                println("output EOF")
                return false
            }
            correctQueue = this.nextQueue()
            checkQueue = queue
            while (checkQueue.head != null && correctQueue.head != null) {
                if (checkQueue.head!!.data != correctQueue.head!!.data) {
                    println("expected: ${correctQueue.head!!.data}")
                    println("have got: ${checkQueue.head!!.data}")
                    return false
                }
                checkQueue.head = checkQueue.head!!.next
                correctQueue.head = correctQueue.head!!.next
            }
            if (checkQueue.head != null) {
                println("Check queue is not empty")
                return false
            }
            if (correctQueue.head != null) {
                println("Correct queue is not empty")
                return false
            }
        }
        return true
    }

    fun nextList(): List {
        val count = this.nextInt()
        val currentOrderNumber = this.nextInt()
        val list = List()
        var newNode: Node
        for (index in 1..count) {
            newNode = Node(data = this.nextInt(), prev = list.last)
            if (list.last != null) {
                list.last!!.next = newNode
            }
            list.last = newNode
            if (currentOrderNumber == index) {
                list.current = newNode
            }
            if (list.first == null) {
                list.first = newNode
            }
        }
        return list
    }

    fun checkList(vararg lists: List): Boolean {
        var checkList: List
        var correctList: List
        for (list in  lists) {
            if (this.eof()) {
                println("output EOF")
                return false
            }
            correctList = this.nextList()
            checkList = list
            if (!checkListAsc(checkList, correctList)) {
                return false
            }
            if (!checkListDesc(checkList, correctList)) {
                return false
            }
            if (correctList.current != null && checkList.current != null &&
                    checkList.current!!.data != correctList.current!!.data) {
                println("CheckList Current expected: ${correctList.current!!.data}")
                println("CheckList Current have got: ${checkList.current!!.data}")
                return false
            }
        }
        return true
    }

    fun checkListAsc(checkList: List, correctList: List): Boolean {
        if (checkList.first == null) {
            if (checkList.current != null) {
                checkList.first = checkList.current
            } else {
                checkList.first = checkList.last
            }
            while (checkList.first != null && checkList.first!!.prev != null &&
                    checkList.first!!.prev != checkList.last) {
                checkList.first = checkList.first!!.prev
            }
        }
        while (checkList.first != null && correctList.first != null) {
            if (correctList.first!!.data != checkList.first!!.data) {
                println("expected: ${correctList.first!!.data}")
                println("have got: ${checkList.first!!.data}")
                return false
            }
            checkList.first = checkList.first!!.next
            correctList.first = correctList.first!!.next
        }
        if (checkList.first != null) {
            println("checkListAsc check list is not empty")
            return false
        }
        if (correctList.first != null) {
            println("checkListAsc correct list is not empty")
            return false
        }
        return true
    }

    fun checkListDesc(checkList: List, correctList: List): Boolean {
        if (checkList.last == null) {
            if (checkList.current != null) {
                checkList.last = checkList.current
            } else {
                checkList.last = checkList.first
            }
            while (checkList.last != null && checkList.last!!.next != null &&
                    checkList.last!!.next != checkList.first) {
                checkList.last = checkList.last!!.next
            }
        }
        while (checkList.last != null && correctList.last != null) {
            if (checkList.last!!.data != correctList.last!!.data) {
                println("expected: ${correctList.last!!.data}")
                println("have got: ${checkList.last!!.data}")
                return false
            }
            checkList.last = checkList.last!!.prev
            correctList.last = correctList.last!!.prev
        }
        if (checkList.last != null) {
            println("checkListDesc check list is not empty")
            return false
        }
        if (correctList.last != null) {
            println("checkListDesc correct list is not empty")
            return false
        }
        return true
    }

    fun nextBarrierList(): Barrier {
        val list = this.nextList()
        val barrier = Barrier(barrier = Node(data = 0, next = list.first, prev = list.last))
        if (list.first == null || list.last == null) {
            barrier.barrier!!.next = barrier.barrier
            barrier.barrier!!.prev = barrier.barrier
            barrier.current = barrier.barrier
        } else {
            list.first!!.prev = barrier.barrier
            list.last!!.next = barrier.barrier
            barrier.current = if (list.current == null) barrier.barrier else list.current
        }
        return barrier
    }

    fun checkBarrierList(vararg barriers: Barrier): Boolean {
        for (barrier in barriers) {
            if (this.eof()) {
                println("output EOF")
                return false
            }
            val list = List(
                first = barrier.barrier!!.next,
                last = barrier.barrier!!.prev,
                current = barrier.current
            )
            if (list.first == list.last || list.first == barrier.barrier) {
                list.first = null
                list.last = null
                list.current = null
            }
            if (!this.checkList(list)) {
                return false
            }
        }
        return true
    }

    fun nextTree(): Node? {
        // read level of tree;
        val level = this.nextInt()
        this.skip()
        // read source tree into words matrix;
        val wordsArray = arrayListOf<kotlin.collections.List<String>>()
        for (index in 1..level) {
            wordsArray.add(this.nextLine().split(' '))
        }
        // make an array for indexes of every row array;
        val colIndexes = IntArray(level){0}
        // parse tree;
        val root = nextTree(null, wordsArray, 0, level, colIndexes)
        return root
    }

    fun nextTree(parent: Node?,
                 wordsArray: ArrayList<kotlin.collections.List<String>>,
                 rowIndex: Int,
                 rowsCount: Int,
                 colIndexes: IntArray): Node? {

        if (rowIndex >= rowsCount || colIndexes[rowIndex] >= wordsArray[rowIndex].size) {
            return null
        }
        val word = wordsArray[rowIndex][colIndexes[rowIndex]]
        colIndexes[rowIndex]++
        var dataStr = ""
        var firstChar = ' '
        var lastChar = ' '
        word.forEachIndexed { index, chr ->
            run {
                if (chr.isDigit()) dataStr += chr.toString()
                if (index == 0) {
                    firstChar = chr
                }
                lastChar = chr
            }
        }
        val data = dataStr.toInt()
        val node = Node(data = data, parent = parent)
        if (firstChar == '┌') {
            node.left = nextTree(node, wordsArray, rowIndex+1, rowsCount, colIndexes)
        }
        if (lastChar == '┐') {
            node.right = nextTree(node, wordsArray, rowIndex+1, rowsCount, colIndexes)
        }
        return node
    }

    fun checkTree(checkRoot: Node?, vararg currents: Node?): Boolean {
        if (checkRoot == null) {
            println("Root is null")
            return false
        }
        val correctRoot = this.nextTree()
        if (!checkTree(checkRoot, correctRoot)) {
            return false
        }
        for (current in currents) {
            if (current == null) {
                println("Check current node doesn't exist")
                return false
            }
            var correctCurrent = correctRoot
            var exit = false
            while (!exit && correctCurrent != null) {
                val direction = this.nextInt()
                correctCurrent = when (direction) {
                    1 -> correctCurrent.left
                    2 -> correctCurrent.right
                    else -> { exit = true; correctCurrent }
                }
            }
            if (correctCurrent == null) {
                println("Correct current node doesn't exist")
                return false
            }
            if (current.data != correctCurrent.data) {
                println("expected: ${correctCurrent.data}")
                println("have got: ${current.data}")
                return false
            }
        }
        return true
    }

    fun checkTree(checkNode: Node?, correctNode: Node?): Boolean {
        if (checkNode == null && correctNode == null) {
            return true
        }
        if (checkNode == null) {
            println("Check Tree is not completed")
            return false
        }
        if (correctNode == null) {
            println("Correct Tree is not completed")
            return false
        }
        return checkNode.data == correctNode.data &&
                checkTree(checkNode.left, correctNode.left) &&
                checkTree(checkNode.right, correctNode.right)
    }
}

private fun Float.round(precision: Int): Float {
    var multiplier = 1F
    repeat(precision) { multiplier *= 10F }
    return (this * multiplier).roundToInt() / multiplier
}

private fun Double.round(precision: Int): Double {
    var multiplier = 1.0
    repeat(precision) { multiplier *= 10.0 }
    return (this * multiplier).roundToLong() / multiplier
}

fun addToRemoveQueue(fileName: String) {
    filesForRemove.add(fileName)
}

fun removeFiles() {
    filesForRemove.forEach { fileName ->
        run {
            val file = File(fileName)
            if (file.exists()) {
                file.delete()
            }
        }
    }
}