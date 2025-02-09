package utils

class Node (
    var data: Int,
    var prev: Node? = null,
    var next: Node? = null,
    var parent: Node? = null,
    var left: Node? = null,
    var right: Node? = null
)