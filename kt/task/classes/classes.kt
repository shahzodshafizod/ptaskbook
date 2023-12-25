package classes

class IntClass(var data: Int)

class IntArrayClass(var data: IntArray)

class DoubleArrayClass(var data: DoubleArray)

class IntMatrixClass(var data: Array<IntArray>)

//////////////////////////////////////////////////////

class Node (
    var data: Int,
    var prev: Node? = null,
    var next: Node? = null,
    var parent: Node? = null,
    var left: Node? = null,
    var right: Node? = null
)

class Stack(var top: Node? = null)

class Queue(
    var head: Node? = null,
    var tail: Node? = null
)

class List(
    var first: Node? = null,
    var last: Node? = null,
    var current: Node? = null
)

class Barrier(
    var barrier: Node? = null,
    var current: Node? = null
)
