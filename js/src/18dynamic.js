import {
	initScanner
	, initChecker
	, inputsEOF
	, outputsEOF
} from './io.js'

(async () => {
	for (let taskNo = 1; taskNo <= 80; taskNo++) {
		const taskNoStr = String(taskNo).padStart(3, 0)
		for (let testNo = 1; testNo <= 100; testNo++) {
			const testNoStr = String(testNo).padStart(3, 0)
			const filePath = `../data/kt-gov2/18dynamic/Dynamic${taskNoStr}/${testNoStr}/${testNoStr}`

			await initScanner(filePath + '.dat')
			await initChecker(filePath + '.ans')

			let ok = true
			switch (taskNo) {
				case 1: ok = dynamic1(); break
                case 2: ok = dynamic2(); break
                case 3: ok = dynamic3(); break
                case 4: ok = dynamic4(); break
                case 5: ok = dynamic5(); break
                case 6: ok = dynamic6(); break
                case 7: ok = dynamic7(); break
                case 8: ok = dynamic8(); break
                case 9: ok = dynamic9(); break
                case 10: ok = dynamic10(); break
                case 11: ok = dynamic11(); break
                case 12: ok = dynamic12(); break
                case 13: ok = dynamic13(); break
                case 14: ok = dynamic14(); break
                case 15: ok = dynamic15(); break
                case 16: ok = dynamic16(); break
                case 17: ok = dynamic17(); break
                case 18: ok = dynamic18(); break
                case 19: ok = dynamic19(); break
                case 20: ok = dynamic20(); break
                case 21: ok = dynamic21(); break
                case 22: ok = dynamic22(); break
                case 23: ok = dynamic23(); break
                case 24: ok = dynamic24(); break
                case 25: ok = dynamic25(); break
                case 26: ok = dynamic26(); break
                case 27: ok = dynamic27(); break
                case 28: ok = dynamic28(); break
                case 29: ok = dynamic29(); break
                case 30: ok = dynamic30(); break
                case 31: ok = dynamic31(); break
                case 32: ok = dynamic32(); break
                case 33: ok = dynamic33(); break
                case 34: ok = dynamic34(); break
                case 35: ok = dynamic35(); break
                case 36: ok = dynamic36(); break
                case 37: ok = dynamic37(); break
                case 38: ok = dynamic38(); break
                case 39: ok = dynamic39(); break
                case 40: ok = dynamic40(); break
                case 41: ok = dynamic41(); break
                case 42: ok = dynamic42(); break
                case 43: ok = dynamic43(); break
                case 44: ok = dynamic44(); break
                case 45: ok = dynamic45(); break
                case 46: ok = dynamic46(); break
                case 47: ok = dynamic47(); break
                case 48: ok = dynamic48(); break
                case 49: ok = dynamic49(); break
                case 50: ok = dynamic50(); break
                case 51: ok = dynamic51(); break
                case 52: ok = dynamic52(); break
                case 53: ok = dynamic53(); break
                case 54: ok = dynamic54(); break
                case 55: ok = dynamic55(); break
                case 56: ok = dynamic56(); break
                case 57: ok = dynamic57(); break
                case 58: ok = dynamic58(); break
                case 59: ok = dynamic59(); break
                case 60: ok = dynamic60(); break
                case 61: ok = dynamic61(); break
                case 62: ok = dynamic62(); break
                case 63: ok = dynamic63(); break
                case 64: ok = dynamic64(); break
                case 65: ok = dynamic65(); break
                case 66: ok = dynamic66(); break
                case 67: ok = dynamic67(); break
                case 68: ok = dynamic68(); break
                case 69: ok = dynamic69(); break
                case 70: ok = dynamic70(); break
                case 71: ok = dynamic71(); break
                case 72: ok = dynamic72(); break
                case 73: ok = dynamic73(); break
                case 74: ok = dynamic74(); break
                case 75: ok = dynamic75(); break
                case 76: ok = dynamic76(); break
                case 77: ok = dynamic77(); break
                case 78: ok = dynamic78(); break
                case 79: ok = dynamic79(); break
                case 80: ok = dynamic80(); break
			}

			if (!ok) {
				console.log("Dynamic" + taskNoStr + " #" + testNoStr + " has NOT been tested!")
				return
			}

			if (!inputsEOF()) {
				console.log("Dynamic" + taskNoStr + " #" + testNoStr + " input NOT EOF!")
				return
			}

			if (!outputsEOF()) {
				console.log("Dynamic" + taskNoStr + " #" + testNoStr + " output NOT EOF!")
				return
			}
		}
		console.log("Dynamic" + taskNoStr + " has been tested!")
	}
	console.log("Dynamic has been tested!")
})()

const dynamic1 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
};

var P1 = makeStack();
document.write("first node data: " + P1.Data + "<br />")
document.write("second node data: " + P1.Next.Data + "<br />");

function makeStack() {
	var i, N = parseInt(prompt("N = "));
	var tempNode, top = null;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		tempNode.Next = top;
		top = tempNode;
		tempNode = null;
	}
	return top;
}
*/
    return false
}

const dynamic2 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
};

var P1 = makeStack();
var i, N = 0;
var lastNode = P1;
for (i = P1; i != null; i = i.Next) {
	document.write(i.Data + " ");
	++N;
	lastNode = i;
}
document.write("<br />N = " + N + "<br />");
document.write("lastNode.Data = " + lastNode.Data + "<br />");

function makeStack() {
	var i, N = parseInt(prompt("N = "));
	var tempNode, top = null;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		tempNode.Next = top;
		top = tempNode;
		tempNode = null;
	}
	return top;
}
*/
    return false
}

const dynamic3 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
};

var D = parseInt(prompt("D = "));
var P1 = makeStack();
var newTop = new TNode();
newTop.Data = D;
newTop.Next = P1;
P1 = newTop;
newTop = null;
var i;
for (i = P1; i != null; i = i.Next) {
	document.write(i.Data + " ");
}

function makeStack() {
	var i, N = parseInt(prompt("N = "));
	var tempNode, top = null;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		tempNode.Next = top;
		top = tempNode;
		tempNode = null;
	}
	return top;
}
*/
    return false
}

const dynamic4 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
};

var N = parseInt(prompt("N = "));
var i, number;
var tempNode, first = null;
for (i = 1; i <= N; ++i) {
	number = parseInt(prompt("number" + i + " = "));
	tempNode = new TNode();
	tempNode.Data = number;
	tempNode.Next = first;
	first = tempNode;
	tempNode = null;
}
printStack(first);

function printStack(node) {
	if (node == null) {
		return;
	}
	document.write(node.Data + " ");
	printStack(node.Next);
}
*/
    return false
}

const dynamic5 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
};

var P1 = makeStack();
var P2 = P1.Next;
document.write("D = " + P1.Data + "<br />");
P1 = null;
printStack(P2);

function makeStack() {
	var i, N = parseInt(prompt("N = "));
	var tempNode, top = null;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		tempNode.Next = top;
		top = tempNode;
		tempNode = null;
	}
	return top;
}

function printStack(node) {
	if (node == null) {
		return;
	}
	document.write(node.Data + " ");
	printStack(node.Next);
}
*/
    return false
}

const dynamic6 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
};

var P1 = makeStack();
var tempNode = null;
var i;
for (i = 0; i < 9; i++) {
	tempNode = P1;
	P1 = P1.Next;
	document.write(tempNode.Data + " ");
	tempNode = null;
}
document.write("<br />");
printStack(P1);

function makeStack() {
	var i, N = parseInt(prompt("N = "));
	var tempNode, top = null;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		tempNode.Next = top;
		top = tempNode;
		tempNode = null;
	}
	return top;
}

function printStack(node) {
	if (node == null) {
		return;
	}
	document.write(node.Data + " ");
	printStack(node.Next);
}
*/
    return false
}

const dynamic7 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
};

var P1 = makeStack();
var tempNode = null;
var N = 0;
while (P1 != null) {
	tempNode = P1;
	P1 = P1.Next;
	document.write(tempNode.Data + " ");
	tempNode = null;
	++N;
}
document.write("<br />N = " + N);

function makeStack() {
	var i, N = parseInt(prompt("N = "));
	var tempNode, top = null;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data:" + i + ':'));
		tempNode.Next = top;
		top = tempNode;
		tempNode = null;
	}
	return top;
}

function printStack(node) {
	if (node == null) {
		return;
	}
	document.write(node.Data + " ");
	printStack(node.Next);
}
*/
    return false
}

const dynamic8 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
};

var P1 = makeStack();
var P2 = makeStack();
var tempNode = null;
while (P1 != null) {
	tempNode = P1;
	P1 = P1.Next;
	tempNode.Next = P2;
	P2 = tempNode;
	tempNode = null;
}
printStack(P2);

function makeStack() {
	var i, N = parseInt(prompt("N = "));
	var tempNode, top = null;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		tempNode.Next = top;
		top = tempNode;
		tempNode = null;
	}
	return top;
}

function printStack(node) {
	if (node == null) {
		return;
	}
	document.write(node.Data + " ");
	printStack(node.Next);
}
*/
    return false
}

const dynamic9 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
};

var P1 = makeStack();
var P2 = makeStack();
var tempNode = null;
while (P1 != null && P1.Data % 2 != 0) {
	tempNode = P1;
	P1 = P1.Next;
	tempNode.Next = P2;
	P2 = tempNode;
	tempNode = null;
}
document.write("P1 = ");
printStack(P1);
document.write("<br />P2 = ");
printStack(P2);

function makeStack() {
	var i, N = parseInt(prompt("N = "));
	var tempNode, top = null;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		tempNode.Next = top;
		top = tempNode;
		tempNode = null;
	}
	return top;
}

function printStack(node) {
	if (node == null) {
		return;
	}
	document.write(node.Data + " ");
	printStack(node.Next);
}
*/
    return false
}

const dynamic10 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
};

var P1 = makeStack();
var P2 = null, P3 = null;
var tempNode = null;
while (P1 != null) {
	tempNode = P1;
	P1 = P1.Next;
	if (tempNode.Data % 2 == 0) {
		tempNode.Next = P2;
		P2 = tempNode;
	} else {
		tempNode.Next = P3;
		P3 = tempNode;
	}
	tempNode = null;
}
document.write("P2 = ");
printStack(P2);
document.write("<br />P3 = ");
printStack(P3);

function makeStack() {
	var i, N = parseInt(prompt("N = "));
	var tempNode, top = null;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		tempNode.Next = top;
		top = tempNode;
		tempNode = null;
	}
	return top;
}

function printStack(node) {
	if (node == null) {
		return;
	}
	document.write(node.Data + " ");
	printStack(node.Next);
}
*/
    return false
}

const dynamic11 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
};

function TStack() {
	this.Top = null;
};

var P1 = makeStack();
var stack = new TStack();
stack.Top = P1;
var N = parseInt(prompt("N = "));
var i, number;
for (i = 1; i <= N; ++i) {
	number = parseInt(prompt("number" + i + ':'));
	Push(stack, number);
}
for (i = stack.Top; i != null; i = i.Next) {
	document.write(i.Data + ", ");
}

function makeStack() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null;
	var tempNode = null;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		tempNode.Next = first;
		first = tempNode;
		tempNode = null;
	}
	return first;
}

function Push(S, D) {
	var newTop = new TNode();
	newTop.Data = D;
	newTop.Next = S.Top;
	S.Top = newTop;
	newTop = null;
}
*/
    return false
}

const dynamic12 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
};

function TStack() {
	this.Top = null;
};

var P1 = makeStack();
var stack = new TStack();
stack.Top = P1;
var i;
for (i = 0; i < 5; ++i) {
	document.write(Pop(stack) + ", ");
}
document.write("<br />");
for (i = stack.Top; i != null; i = i.Next) {
	document.write(i.Data + ", ");
}

function makeStack() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, tempNode = null;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		tempNode.Next = first;
		first = tempNode;
		tempNode = null;
	}
	return first;
}

function Pop(S) {
	var tempNode = S.Top;
	S.Top = S.Top.Next;
	return tempNode.Data;
}
*/
    return false
}

const dynamic13 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
};

function TStack(_Top) {
	this.Top = _Top;
};

var P1 = makeStack();
var stack = new TStack(P1);
var i;
for (i = 0; !StackIsEmpty(stack) && i < 5; ++i) {
	document.write(Pop(stack) + ", ");
}
document.write("<br />StackIsEmpty: " + StackIsEmpty(stack));
if (!StackIsEmpty(stack)) {
	document.write("<br />Peek: " + Peek(stack) + "<br />");
	for (i = stack.Top; i != null; i = i.Next) {
		document.write(i.Data + ", ");
	}
}

function makeStack() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, tempNode = null;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		tempNode.Next = first;
		first = tempNode;
		tempNode = null;
	}
	return first;
}

function StackIsEmpty(S) {
	return S.Top == null;
}

function Peek(S) {
	return S.Top.Data;
}

function Pop(S) {
	var tempNode = S.Top;
	S.Top = S.Top.Next;
	return tempNode.Data;
}
*/
    return false
}

const dynamic14 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
};

var i, number;
var first = null, last = null;
var tempNode;
for (i = 1; i <= 10; ++i) {
	number = parseInt(prompt("number" + i + ':'));
	tempNode = new TNode();
	tempNode.Data = number;
	if (first == null) {
		first = last = tempNode;
	} else if (first == last) {
		last = tempNode;
		first.Next = last;
	} else {
		last.Next = tempNode;
		last = tempNode;
	}
	tempNode = null;
}
for (i = first; i != null; i = i.Next) {
	document.write(i.Data + ", ");
}
*/
    return false
}

const dynamic15 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
};

var i, number;
var P1 = null, P2 = null;
var P3 = null, P4 = null;
var tempNode;
for (i = 1; i <= 10; ++i) {
	number = parseInt(prompt("number" + i + ':'));
	tempNode = new TNode();
	tempNode.Data = number;
	if (i % 2 != 0) {
		if (P1 == null) {
			P1 = P2 = tempNode;
		} else {
			P2.Next = tempNode;
			P2 = tempNode;
		}
	} else {
		if (P3 == null) {
			P3 = P4 = tempNode;
		} else {
			P4.Next = tempNode;
			P4 = tempNode;
		}
	}
	tempNode = null;
}
printQueue(P1);
printQueue(P3);

function printQueue(node) {
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}
*/
    return false
}

const dynamic16 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
};

var i, number;
var P1 = null, P2 = null;
var P3 = null, P4 = null;
var tempNode;
for (i = 1; i <= 10; ++i) {
	number = parseInt(prompt("number" + i + ':'));
	tempNode = new TNode();
	tempNode.Data = number;
	if (number % 2 != 0) {
		if (P1 == null) {
			P1 = P2 = tempNode;
		} else {
			P2.Next = tempNode;
			P2 = tempNode;
		}
	} else {
		if (P3 == null) {
			P3 = P4 = tempNode;
		} else {
			P4.Next = tempNode;
			P4 = tempNode;
		}
	}
	tempNode = null;
}
printQueue(P1);
printQueue(P3);

function printQueue(node) {
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}
*/
    return false
}

const dynamic17 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
};

var D = parseInt(prompt("D = "));
var P1 = makeQueue();
var P2 = getLast(P1);
var newNode = new TNode();
newNode.Data = D;
if (P1 == null) {
	P1 = P2 = newNode;
} else {
	P2.Next = newNode;
	P2 = newNode;
}
newNode = null;
printQueue(P1);

function makeQueue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function printQueue(node) {
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}
*/
    return false
}

const dynamic18 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
};

var D = parseInt(prompt("D = "));
var P1 = makeQueue();
var P2 = getLast(P1);
var tempNode = new TNode();
tempNode.Data = D;
P2.Next = tempNode;
P2 = tempNode;
tempNode = P1;
P1 = P1.Next;
document.write(tempNode.Data + "<br />");
printQueue(P1);

function makeQueue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function printQueue(node) {
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}
*/
    return false
}

const dynamic19 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
};

var N = parseInt(prompt("N = "));
var P1 = makeQueue();
var P2 = getLast(P1);
var i;
var tempNode;
for (i = 0; (P1 != null) && (i < N); ++i) {
	tempNode = P1;
	document.write(tempNode.Data + ", ");
	P1 = P1.Next;
}
document.write("<br />");
if (P1 == null) {
	P2 = null;
}
printQueue(P1);

function makeQueue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function printQueue(node) {
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}
*/
    return false
}

const dynamic20 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
};

var P1 = makeQueue();
var P2 = getLast(P1);
var tempNode;
while (P1 != null && P1.Data % 2 != 0) {
	tempNode = P1;
	document.write(tempNode.Data + ", ");
	P1 = P1.Next;
}
document.write("<br />");
if (P1 == null) {
	P2 = null;
}
printQueue(P1);

function makeQueue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function printQueue(node) {
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}
*/
    return false
}

const dynamic21 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
};

var P1 = makeQueue();
var P2 = getLast(P1);
var P3 = makeQueue();
var P4 = getLast(P3);
var tempNode;
while (P1 != null) {
	tempNode = P1;
	P1 = P1.Next;
	tempNode.Next = null;
	if (P3 == null) {
		P3 = P4 = tempNode;
	} else {
		P4.Next = tempNode;
		P4 = tempNode;
	}
	tempNode = null;
}
P2 = null;
printQueue(P3);

function makeQueue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function printQueue(node) {
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}
*/
    return false
}

const dynamic22 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
};

var N = parseInt(prompt("N = "));
var P1 = makeQueue();
var P2 = getLast(P1);
var P3 = makeQueue();
var P4 = getLast(P3);
var tempNode;
var i;
for (i = 0; (P1 != null) && (i < N); ++i) {
	tempNode = P1;
	P1 = P1.Next;
	tempNode.Next = null;
	if (P3 == null) {
		P3 = P4 = tempNode;
	} else {
		P4.Next = tempNode;
		P4 = tempNode;
	}
	tempNode = null;
}
if (P1 == null) {
	P2 = null;
}
printQueue(P1);
printQueue(P3);

function makeQueue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function printQueue(node) {
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}
*/
    return false
}

const dynamic23 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
};

var P1 = makeQueue();
var P2 = getLast(P1);
var P3 = makeQueue();
var P4 = getLast(P3);
var tempNode;
while ((P1 != null) && (P1.Data % 2 != 0)) {
	tempNode = P1;
	P1 = P1.Next;
	tempNode.Next = null;
	if (P3 == null) {
		P3 = P4 = tempNode;
	} else {
		P4.Next = tempNode;
		P4 = tempNode;
	}
	tempNode = null;
}
if (P1 == null) {
	P2 = null;
}
printQueue(P1);
printQueue(P3);

function makeQueue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function printQueue(node) {
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}
*/
    return false
}

const dynamic24 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
};

var P1 = makeQueue();
var P2 = getLast(P1);
var P3 = makeQueue();
var P4 = getLast(P3);
var P5 = null, P6 = null;
var tempNode;
var i;
while (P1 != null) {
	for (i = 0; i < 2; ++i) {
		switch (i) {
			case 0: tempNode = P1; P1 = P1.Next; break;
			case 1: tempNode = P3; P3 = P3.Next; break;
		}
		tempNode.Next = null;
		if (P5 == null) {
			P5 = P6 = tempNode;
		} else {
			P6.Next = tempNode;
			P6 = tempNode;
		}
	}
}
P2 = P4 = null;
printQueue(P5);

function makeQueue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function printQueue(node) {
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}
*/
    return false
}

const dynamic25 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
};

var P1 = makeQueue();
var P2 = getLast(P1);
var P3 = makeQueue();
var P4 = getLast(P3);
var P5 = null, P6 = null;
var tempNode;
while (P1 != null || P3 != null) {
	if (P1 == null) {
		tempNode = P3;
		P3 = P3.Next;
	} else if (P3 == null) {
		tempNode = P1;
		P1 = P1.Next;
	} else if (P1.Data < P3.Data) {
		tempNode = P1;
		P1 = P1.Next;
	} else {
		tempNode = P3;
		P3 = P3.Next;
	}
	tempNode.Next = null;
	if (P5 == null) {
		P5 = P6 = tempNode;
	} else {
		P6.Next = tempNode;
		P6 = tempNode;
	}
	tempNode = null;
}
P2 = P4 = null;
printQueue(P5);

function makeQueue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function printQueue(node) {
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}
*/
	return false
}

const dynamic26 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
};

function TQueue(_Head, _Tail) {
	this.Head = _Head;
	this.Tail = _Tail;
};

var P1 = makeQueue();
var P2 = getLast(P1);
var queue = new TQueue(P1, P2);
var N = parseInt(prompt("N = "));
var i, number;
for (i = 1; i <= N; ++i) {
	number = parseInt(prompt("number" + i + ':'));
	Enqueue(queue, number);
}
printQueue(queue.Head);

function makeQueue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function printQueue(node) {
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}

function Enqueue(Q, D) {
	var newNode = new TNode();
	newNode.Data = D;
	if (Q.Head == null) {
		Q.Head = Q.Tail = newNode;
	} else {
		Q.Tail.Next = newNode;
		Q.Tail = newNode;
	}
	newNode = null;
}
*/
	return false
}

const dynamic27 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
};

function TQueue(_Head, _Tail) {
	this.Head = _Head;
	this.Tail = _Tail;
};

var P1 = makeQueue();
var P2 = getLast(P1);
var queue = new TQueue(P1, P2);
var i;
for (i = 0; i < 5; ++i) {
	document.write(Dequeue(queue) + ", ");
}
document.write("<br />");
printQueue(queue.Head);

function makeQueue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function printQueue(node) {
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}

function Dequeue(Q) {
	var tempNode = Q.Head;
	Q.Head = Q.Head.Next;
	return tempNode.Data;
}
*/
	return false
}

const dynamic28 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
};

function TQueue(_Head, _Tail) {
	this.Head = _Head;
	this.Tail = _Tail;
};

var P1 = makeQueue();
var P2 = getLast(P1);
var queue = new TQueue(P1, P2);
var i;
for (i = 0; !QueueIsEmpty(queue) && (i < 5); ++i) {
	document.write(Dequeue(queue) + ", ");
}
document.write("<br />");
printQueue(queue.Head);

function makeQueue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function printQueue(node) {
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}

function Dequeue(Q) {
	var tempNode = Q.Head;
	Q.Head = Q.Head.Next;
	return tempNode.Data;
}

function QueueIsEmpty(Q) {
	return Q.Head == null;
}
*/
	return false
}

const dynamic29 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

var P2 = new TNode();
P2.Data = parseInt(prompt("P2.Data = "));

P2.Prev = new TNode();
P2.Prev.Data = parseInt(prompt("P2.Prev.Data = "));
P2.Prev.Next = P2;

P2.Next = new TNode();
P2.Next.Data = parseInt(prompt("P2.Next.Data = "));
P2.Next.Prev = P2;

document.write("P2.Prev.Data = " + P2.Prev.Data + "<br />");
document.write("P2.Next.Data = " + P2.Next.Data + "<br />");
*/
	return false
}

const dynamic30 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

var P1 = makeQueue();
P1.Prev = null;
var P2 = P1;
while (P2.Next != null) {
	P2.Next.Prev = P2;
	P2 = P2.Next;
}
printIncDequeue(P1);
printDecDequeue(P2);

function makeQueue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var newNode;
	for (i = 1; i <= N; ++i) {
		newNode = new TNode();
		newNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = newNode;
		} else {
			last.Next = newNode;
			last = newNode;
		}
		newNode = null;
	}
	return first;
}

function printIncDequeue(node) {
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}

function printDecDequeue(node) {
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Prev;
	}
	document.write("<br />");
}
*/
	return false
}

const dynamic31 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

var P1 = makeDequeue();

var index = parseInt(prompt("What node?"));
var P0 = getNode(P1, index);
P1 = P0;
while (P1.Prev != null) {
	P1 = P1.Prev;
}
var N = 1;
var P2 = P1;
while (P2.Next != null) {
	P2 = P2.Next;
	++N;
}
document.write("N = " + N + "<br />");
printIncDequeue(P1);
printDecDequeue(P2);

function makeDequeue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			tempNode.Prev = last;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function printIncDequeue(node) {
	node = getFirst(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}

function printDecDequeue(node) {
	node = getLast(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Prev;
	}
	document.write("<br />");
}

function getNode(node, index) {
	node = getFirst(node);
	while (node != null && index > 1) {
		node = node.Next;
		--index;
	}
	return node;
}

function getFirst(node) {
	if (node == null) {
		node;
	}
	while (node.Prev != null) {
		node = node.Prev;
	}
	return node;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}
*/
	return false
}

const dynamic32 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

var D1 = parseInt(prompt("D1 = "));
var D2 = parseInt(prompt("D2 = "));
var P1 = makeDequeue();
var index = parseInt(prompt("What node?"));
var P0 = getNode(P1, index);
P1 = getFirst(P0);
var P2 = getLast(P0);

var newNode = new TNode();
newNode.Data = D1;
newNode.Next = P1;
P1.Prev = newNode;
P1 = newNode;

newNode = new TNode();
newNode.Data = D2;
newNode.Prev = P2;
P2.Next = newNode;
P2 = newNode;
newNode = null;

printIncDequeue(P1);
printDecDequeue(P2);

function makeDequeue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			tempNode.Prev = last;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getFirst(node) {
	if (node == null) {
		return node;
	}
	while (node.Prev != null) {
		node = node.Prev;
	}
	return node;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function getNode(node, index) {
	node = getFirst(node);
	while (node != null && index > 1) {
		node = node.Next;
		--index;
	}
	return node;
}

function printIncDequeue(node) {
	node = getFirst(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}

function printDecDequeue(node) {
	node = getLast(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Prev;
	}
	document.write("<br />");
}
*/
	return false
}

const dynamic33 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

var D = parseInt(prompt("D = "));
var P1 = makeDequeue();
var index = parseInt(prompt("What node?"));
var P0 = getNode(P1, index);
P1 = getFirst(P0);
var P2 = getLast(P0);

var newNode = new TNode();
newNode.Data = D;

newNode.Prev = P0.Prev;
newNode.Next = P0;
if (P0.Prev != null) {
	P0.Prev.Next = newNode;
}
P0.Prev = newNode;
if (P0 == P1) {
	P1 = P1.Prev;
}

printIncDequeue(P1);
printDecDequeue(P2);

function makeDequeue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			tempNode.Prev = last;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getFirst(node) {
	if (node == null) {
		return node;
	}
	while (node.Prev != null) {
		node = node.Prev;
	}
	return node;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function getNode(node, index) {
	node = getFirst(node);
	while (node != null && index > 1) {
		node = node.Next;
		--index;
	}
	return node;
}

function printIncDequeue(node) {
	node = getFirst(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}

function printDecDequeue(node) {
	node = getLast(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Prev;
	}
	document.write("<br />");
}
*/
	return false
}

const dynamic34 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

var D = parseInt(prompt("D = "));
var P1 = makeDequeue();
var index = parseInt(prompt("What node?"));
var P0 = getNode(P1, index);
P1 = getFirst(P0);
var P2 = getLast(P0);

var newNode = new TNode();
newNode.Data = D;

newNode.Prev = P0;
newNode.Next = P0.Next;
if (P0.Next != null) {
	P0.Next.Prev = newNode;
}
P0.Next = newNode;
if (P0 == P2) {
	P2 = P2.Next;
}

printIncDequeue(P1);
printDecDequeue(P2);

function makeDequeue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			tempNode.Prev = last;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getFirst(node) {
	if (node == null) {
		return node;
	}
	while (node.Prev != null) {
		node = node.Prev;
	}
	return node;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function getNode(node, index) {
	node = getFirst(node);
	while (node != null && index > 1) {
		node = node.Next;
		--index;
	}
	return node;
}

function printIncDequeue(node) {
	node = getFirst(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}

function printDecDequeue(node) {
	node = getLast(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Prev;
	}
	document.write("<br />");
}
*/
	return false
}

const dynamic35 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

var P1 = makeDequeue();
var P2 = getLast(P1);

P1.Prev = new TNode();
P1.Prev.Next = P1;
P1.Prev.Data = P1.Data;
P1 = P1.Prev;

var newNode = new TNode();
newNode.Data = P2.Data;
newNode.Next = P2;
newNode.Prev = P2.Prev;
newNode.Next.Prev = newNode;
newNode.Prev.Next = newNode;
newNode = null;

printIncDequeue(P1);
printDecDequeue(P2);

function makeDequeue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			tempNode.Prev = last;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function getFirst(node) {
	if (node == null) {
		return node;
	}
	while (node.Prev != null) {
		node = node.Prev;
	}
	return node;
}

function printIncDequeue(node) {
	node = getFirst(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}

function printDecDequeue(node) {
	node = getLast(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Prev;
	}
	document.write("<br />");
}
*/
	return false
}

const dynamic36 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

var P1 = makeDequeue();
var P2 = getLast(P1);

var newNode = new TNode();
newNode.Data = P1.Data;
newNode.Prev = P1;
newNode.Next = P1.Next;
newNode.Next.Prev = newNode;
newNode.Prev.Next = newNode;
newNod = null;

P2.Next = new TNode();
P2.Next.Data = P2.Data;
P2.Next.Prev = P2;
P2 = P2.Next;

printIncDequeue(P1);
printDecDequeue(P2);

function makeDequeue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			tempNode.Prev = last;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function getFirst(node) {
	if (node == null) {
		return node;
	}
	while (node.Prev != null) {
		node = node.Prev;
	}
	return node;
}

function printIncDequeue(node) {
	node = getFirst(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}

function printDecDequeue(node) {
	node = getLast(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Prev;
	}
	document.write("<br />");
}
*/
	return false
}

const dynamic37 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

var P1 = makeDequeue();
var P2 = getLast(P1);

var i, tempNode;
var j;
for (i = P1; i != null;) {
	tempNode = new TNode();
	tempNode.Data = i.Data;
	tempNode.Next = i;
	tempNode.Prev = i.Prev;
	if (i.Prev != null) {
		i.Prev.Next = tempNode;
	}
	i.Prev = tempNode;
	tempNode = null;
	for (j = 0; i != null && j < 2; ++j) {
		i = i.Next;
	}
}
P1 = getFirst(P1);

printIncDequeue(P1);
printDecDequeue(P2);

function makeDequeue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			tempNode.Prev = last;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function getFirst(node) {
	if (node == null) {
		return node;
	}
	while (node.Prev != null) {
		node = node.Prev;
	}
	return node;
}

function printIncDequeue(node) {
	node = getFirst(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}

function printDecDequeue(node) {
	node = getLast(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Prev;
	}
	document.write("<br />");
}
*/
	return false
}

const dynamic38 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

var P1 = makeDequeue();
var P2 = getLast(P1);

var j;
var i, tempNode;
for (i = P1; i != null;) {
	tempNode = new TNode();
	tempNode.Data = i.Data;
	tempNode.Prev = i;
	tempNode.Next = i.Next;
	if (i.Next != null) {
		i.Next.Prev = tempNode;
	}
	i.Next = tempNode;
	tempNode = null;
	for (j = 0; i != null && j < 3; ++j) {
		i = i.Next;
	}
}
P2 = getLast(P2);

printIncDequeue(P1);
printDecDequeue(P2);

function makeDequeue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			tempNode.Prev = last;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function getFirst(node) {
	if (node == null) {
		return node;
	}
	while (node.Prev != null) {
		node = node.Prev;
	}
	return node;
}

function printIncDequeue(node) {
	node = getFirst(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}

function printDecDequeue(node) {
	node = getLast(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Prev;
	}
	document.write("<br />");
}
*/
	return false
}

const dynamic39 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

var P1 = makeDequeue();
var P2 = getLast(P1);

var i, tempNode;
for (i = P1; i != null; i = i.Next) {
	if (i.Data % 2 == 0) {
		continue;
	}
	tempNode = new TNode();
	tempNode.Data = i.Data;
	tempNode.Next = i;
	tempNode.Prev = i.Prev;
	if (i.Prev != null) {
		i.Prev.Next = tempNode;
	}
	i.Prev = tempNode;
	tempNode = null;
}
P1 = getFirst(P1);

printIncDequeue(P1);
printDecDequeue(P2);

function makeDequeue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			tempNode.Prev = last;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function getFirst(node) {
	if (node == null) {
		return node;
	}
	while (node.Prev != null) {
		node = node.Prev;
	}
	return node;
}

function printIncDequeue(node) {
	node = getFirst(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}

function printDecDequeue(node) {
	node = getLast(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Prev;
	}
	document.write("<br />");
}
*/
	return false
}

const dynamic40 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

var P1 = makeDequeue();
var P2 = getLast(P1);

var j;
var i, tempNode;
for (i = P1; i != null;) {
	if (i.Data % 2 == 0) {
		i = i.Next;
		continue;
	}
	tempNode = new TNode();
	tempNode.Data = i.Data;
	tempNode.Prev = i;
	tempNode.Next = i.Next;
	if (i.Next != null) {
		i.Next.Prev = tempNode;
	}
	i.Next = tempNode;
	tempNode = null;
	for (j = 0; i != null && j < 2; ++j) {
		i = i.Next;
	}
}
P2 = getLast(P2);

printIncDequeue(P1);
printDecDequeue(P2);

function makeDequeue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			tempNode.Prev = last;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function getFirst(node) {
	if (node == null) {
		return node;
	}
	while (node.Prev != null) {
		node = node.Prev;
	}
	return node;
}

function printIncDequeue(node) {
	node = getFirst(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}

function printDecDequeue(node) {
	node = getLast(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Prev;
	}
	document.write("<br />");
}
*/
	return false
}

const dynamic41 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

var P1 = makeDequeue();
var P2 = getLast(P1);
var index = parseInt(prompt("What node?"));
var P0 = getNode(P1, index);
if (P0 == P1) {
	P1 = P1.Next;
	P1.Prev = null;
	P0.Next = null;
} else if (P0 == P2) {
	P2 = P2.Prev;
	P2.Next = null;
	P0.Prev = null;
} else {
	P0.Next.Prev = P0.Prev;
	P0.Prev.Next = P0.Next;
	P0.Next = P0.Prev = null;
}
printIncDequeue(P1);
printDecDequeue(P2);

function makeDequeue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			tempNode.Prev = last;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function getFirst(node) {
	if (node == null) {
		return node;
	}
	while (node.Prev != null) {
		node = node.Prev;
	}
	return node;
}

function getNode(node, index) {
	node = getFirst(node);
	while (node != null && index > 1) {
		node = node.Next;
		--index;
	}
	return node;
}

function printIncDequeue(node) {
	node = getFirst(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}

function printDecDequeue(node) {
	node = getLast(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Prev;
	}
	document.write("<br />");
}
*/
	return false
}

const dynamic42 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

var P1 = makeDequeue();
var i;
var tempNode;
for (i = P1; i != null;) {
	tempNode = i;
	if (i == P1) {
		P1 = P1.Next;
	}
	i = i.Next;
	if (tempNode.Prev != null) {
		tempNode.Prev.Next = i;
	}
	if (i != null) {
		i.Prev = tempNode.Prev;
		i = i.Next;
	}
	tempNode.Prev = tempNode.Next = null;
}
var P2 = getLast(P1);
printIncDequeue(P1);
printDecDequeue(P2);

function makeDequeue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			tempNode.Prev = last;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function getFirst(node) {
	if (node == null) {
		return node;
	}
	while (node.Prev != null) {
		node = node.Prev;
	}
	return node;
}

function printIncDequeue(node) {
	node = getFirst(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}

function printDecDequeue(node) {
	node = getLast(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Prev;
	}
	document.write("<br />");
}
*/
	return false
}

const dynamic43 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

var P1 = makeDequeue();
var i;
var tempNode;
for (i = P1; i != null;) {
	if (i.Data % 2 == 0) {
		i = i.Next;
		continue;
	}
	tempNode = i;
	if (i == P1) {
		P1 = P1.Next;
	}
	i = i.Next;
	if (tempNode.Prev != null) {
		tempNode.Prev.Next = i;
	}
	if (i != null) {
		i.Prev = tempNode.Prev;
	}
	tempNode.Prev = tempNode.Next = null;
}
var P2 = getLast(P1);
printIncDequeue(P1);
printDecDequeue(P2);

function makeDequeue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			tempNode.Prev = last;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function getFirst(node) {
	if (node == null) {
		return node;
	}
	while (node.Prev != null) {
		node = node.Prev;
	}
	return node;
}

function printIncDequeue(node) {
	node = getFirst(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}

function printDecDequeue(node) {
	node = getLast(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Prev;
	}
	document.write("<br />");
}
*/
	return false
}

const dynamic44 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

var P1 = makeDequeue();
var P2 = getLast(P1);
var index = parseInt(prompt("What node?"));
var P0 = getNode(P1, index);
if (P0 != P2) {
	if (P0 == P1) {
		P1 = P1.Next;
	}
	if (P0.Prev != null) {
		P0.Prev.Next = P0.Next;
	}
	if (P0.Next != null) {
		P0.Next.Prev = P0.Prev;
	}
	P2.Next = P0;
	P0.Prev = P2;
	P0.Next = null;
	P2 = P0;
}
printIncDequeue(P1);
printDecDequeue(P2);

function makeDequeue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			tempNode.Prev = last;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function getNode(node, index) {
	node = getFirst(node);
	while (node != null && index > 1) {
		node = node.Next;
		--index;
	}
	return node;
}

function getFirst(node) {
	if (node == null) {
		return node;
	}
	while (node.Prev != null) {
		node = node.Prev;
	}
	return node;
}

function printIncDequeue(node) {
	node = getFirst(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}

function printDecDequeue(node) {
	node = getLast(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Prev;
	}
	document.write("<br />");
}
*/
	return false
}

const dynamic45 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

var P1 = makeDequeue();
var P2 = getLast(P1);
var index = parseInt(prompt("What node?"));
var P0 = getNode(P1, index);
if (P0 != P1) {
	if (P0 == P2) {
		P2 = P2.Prev;
	}
	if (P0.Prev != null) {
		P0.Prev.Next = P0.Next;
	}
	if (P0.Next != null) {
		P0.Next.Prev = P0.Prev;
	}
	P1.Prev = P0;
	P0.Next = P1;
	P0.Prev = null;
	P1 = P0;
}
printIncDequeue(P1);
printDecDequeue(P2);

function makeDequeue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			tempNode.Prev = last;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function getNode(node, index) {
	node = getFirst(node);
	while (node != null && index > 1) {
		node = node.Next;
		--index;
	}
	return node;
}

function getFirst(node) {
	if (node == null) {
		return node;
	}
	while (node.Prev != null) {
		node = node.Prev;
	}
	return node;
}

function printIncDequeue(node) {
	node = getFirst(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}

function printDecDequeue(node) {
	node = getLast(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Prev;
	}
	document.write("<br />");
}
*/
	return false
}

const dynamic46 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

var K = parseInt(prompt("K = "));
var P1 = makeDequeue();
var index = parseInt(prompt("What element?"));
var P0 = getNode(P1, index);
var i = 1;
if (P0 == P1 && K > 0) {
	P1 = P1.Next;
}
while (P0.Next != null && i <= K) {
	if (P0.Prev != null) {
		P0.Prev.Next = P0.Next;
	}
	P0.Next.Prev = P0.Prev;
	
	P0.Prev = P0.Next;
	P0.Next = P0.Next.Next;
	
	P0.Prev.Next = P0;
	if (P0.Next != null) {
		P0.Next.Prev = P0;
	}
	++i;
}
var P2 = getLast(P0);
printIncDequeue(P1);
printDecDequeue(P2);

function makeDequeue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			tempNode.Prev = last;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getNode(node, nomer) {
	node = getFirst(node);
	while (node != null && nomer > 1) {
		node = node.Next;
		--nomer;
	}
	return node;
}

function getFirst(node) {
	if (node == null) {
		return node;
	}
	while (node.Prev != null) {
		node = node.Prev;
	}
	return node;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function printIncDequeue(node) {
	node = getFirst(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}

function printDecDequeue(node) {
	node = getLast(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Prev;
	}
	document.write("<br />");
}
*/
	return false
}

const dynamic47 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

var K = parseInt(prompt("K = "));
var P1 = makeDequeue();
var index = parseInt(prompt("What element?"));
var P0 = getNode(P1, index);
var i = 1;
while (P0.Prev != null && i <= K) {
	P0.Prev.Next = P0.Next;
	if (P0.Next != null) {
		P0.Next.Prev = P0.Prev;
	}
	
	P0.Next = P0.Prev;
	P0.Prev = P0.Prev.Prev;
	
	P0.Next.Prev = P0;
	if (P0.Prev != null) {
		P0.Prev.Next = P0;
	}
	++i;
}
if (P1.Prev != null) {
	P1 = P1.Prev;
}
var P2 = getLast(P0);
printIncDequeue(P1);
printDecDequeue(P2);

function makeDequeue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			tempNode.Prev = last;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getNode(node, nomer) {
	node = getFirst(node);
	while (node != null && nomer > 1) {
		node = node.Next;
		--nomer;
	}
	return node;
}

function getFirst(node) {
	if (node == null) {
		return node;
	}
	while (node.Prev != null) {
		node = node.Prev;
	}
	return node;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function printIncDequeue(node) {
	node = getFirst(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}

function printDecDequeue(node) {
	node = getLast(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Prev;
	}
	document.write("<br />");
}
*/
	return false
}

const dynamic48 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

var P1 = makeDequeue();
var x = parseInt(prompt("What node for x?"));
var Px = getNode(P1, x);
var y = parseInt(prompt("What node for y?"));
var Py = getNode(P1, y);
if (Px.Next == Py) {
	Px.Next.Prev = Px.Prev;
	if (Px.Prev != null) {
		Px.Prev.Next = Px.Next;
	}
	
	Px.Prev = Px.Next;
	Px.Next = Px.Next.Next;
	
	Px.Prev.Next = Px;
	if (Px.Next != null) {
		Px.Next.Prev = Px;
	}
} else {
	var next = Px.Next;
	var prev = Px.Prev;
	
	Px.Next = Py.Next;
	Px.Prev = Py.Prev;
	Py.Prev.Next = Px;
	if (Py.Next != null) {
		Py.Next.Prev = Px;
	}
	
	Py.Next = next;
	Py.Prev = prev;
	Py.Next.Prev = Py;
	if (Py.Prev != null) {
		Py.Prev.Next = Py;
	}
}
P1 = getFirst(Py);
var P2 = getLast(Px);
printIncDequeue(P1);
printDecDequeue(P2);

function makeDequeue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			tempNode.Prev = last;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getNode(node, nomer) {
	node = getFirst(node);
	while (node != null && nomer > 1) {
		node = node.Next;
		--nomer;
	}
	return node;
}

function getFirst(node) {
	if (node == null) {
		return node;
	}
	while (node.Prev != null) {
		node = node.Prev;
	}
	return node;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function printIncDequeue(node) {
	node = getFirst(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}

function printDecDequeue(node) {
	node = getLast(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Prev;
	}
	document.write("<br />");
}
*/
	return false
}

const dynamic49 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

var P1 = makeDequeue();
var i = P1;
var P2 = P1;
if (P1.Next != null) {
	P1 = P1.Next;
}
var N = 1;
while (P2.Next != null) {
	P2 = P2.Next;
	++N;
}
var tempNode;
var j = parseInt(N/2) + (N%2 != 0 ? 1 : 0);
while (j > 0) {
	tempNode = i;
	tempNode.Next.Prev = tempNode.Prev;
	if (tempNode.Prev != null) {
		tempNode.Prev.Next = tempNode.Next;
	}
	i = i.Next.Next;
	tempNode.Next = null;
	tempNode.Prev = P2;
	P2.Next = tempNode;
	P2 = tempNode;
	tempNode = null;
	--j;
}
printIncDequeue(P1);
printDecDequeue(P2);

function makeDequeue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			tempNode.Prev = last;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function getFirst(node) {
	if (node == null) {
		return node;
	}
	while (node.Prev != null) {
		node = node.Prev;
	}
	return node;
}

function printIncDequeue(node) {
	node = getFirst(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}

function printDecDequeue(node) {
	node = getLast(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Prev;
	}
	document.write("<br />");
}
*/
	return false
}

const dynamic50 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

var P1 = makeDequeue();
var i = P1;
var P2 = P1;
var N = 1;
var hasOdd = P1.Data % 2 != 0;
while (P2.Next != null) {
	P2 = P2.Next;
	if (!hasOdd && (P2.Data % 2 != 0)) {
		hasOdd = true;
	}
	++N;
}
if (hasOdd) {
	var tempNode;
	while (N > 0) {
		--N;
		tempNode = i;
		i = i.Next;
		if (tempNode.Data % 2 == 0) {
			continue;
		}
		if (tempNode == P1) {
			P1 = P1.Next;
		}
		tempNode.Next.Prev = tempNode.Prev;
		if (tempNode.Prev != null) {
			tempNode.Prev.Next = tempNode.Next;
		}
		tempNode.Next = null;
		tempNode.Prev = P2;
		P2.Next = tempNode;
		P2 = tempNode;
		tempNode = null;
	}
}
printIncDequeue(P1);
printDecDequeue(P2);

function makeDequeue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			tempNode.Prev = last;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function getFirst(node) {
	if (node == null) {
		return node;
	}
	while (node.Prev != null) {
		node = node.Prev;
	}
	return node;
}

function printIncDequeue(node) {
	node = getFirst(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}

function printDecDequeue(node) {
	node = getLast(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Prev;
	}
	document.write("<br />");
}
*/
	return false
}

const dynamic51 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

var P1 = makeDequeue();
var P2 = getLast(P1);
var P3 = makeDequeue();
var nomer = parseInt(prompt("What node?"));
var P0 = getNode(P3, nomer);
var tempNode;
while (P1 != null) {
	tempNode = P1;
	P1 = P1.Next;
	tempNode.Next = P0;
	tempNode.Prev = P0.Prev;
	if (P0.Prev != null) {
		P0.Prev.Next = tempNode;
	}
	P0.Prev = tempNode;
	tempNode = null;
}
P2 = null;
P3 = getFirst(P0);
var P4 = getLast(P0);
printIncDequeue(P3);
printDecDequeue(P4);

function makeDequeue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			tempNode.Prev = last;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function getFirst(node) {
	if (node == null) {
		return node;
	}
	while (node.Prev != null) {
		node = node.Prev;
	}
	return node;
}

function getNode(node, index) {
	node = getFirst(node);
	while (node != null && index > 1) {
		node = node.Next;
		--index;
	}
	return node;
}

function printIncDequeue(node) {
	node = getFirst(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}

function printDecDequeue(node) {
	node = getLast(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Prev;
	}
	document.write("<br />");
}
*/
	return false
}

const dynamic52 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

var P1 = makeDequeue();
var P2 = getLast(P1);
var P3 = makeDequeue();
var nomer = parseInt(prompt("What node?"));
var P0 = getNode(P3, nomer);
var tempNode;
while (P1 != null) {
	tempNode = P1;
	P1 = P1.Next;
	tempNode.Prev = P0;
	tempNode.Next = P0.Next;
	if (tempNode.Next != null) {
		tempNode.Next.Prev = tempNode;
	}
	tempNode.Prev.Next = tempNode;
	P0 = tempNode;
	tempNode = null;
}
P2 = null;
P3 = getFirst(P0);
var P4 = getLast(P0);
printIncDequeue(P3);
printDecDequeue(P4);

function makeDequeue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			tempNode.Prev = last;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function getFirst(node) {
	if (node == null) {
		return node;
	}
	while (node.Prev != null) {
		node = node.Prev;
	}
	return node;
}

function getNode(node, index) {
	node = getFirst(node);
	while (node != null && index > 1) {
		node = node.Next;
		--index;
	}
	return node;
}

function printIncDequeue(node) {
	node = getFirst(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}

function printDecDequeue(node) {
	node = getLast(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Prev;
	}
	document.write("<br />");
}
*/
	return false
}

const dynamic53 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

var P1 = makeDequeue();
var x = parseInt(prompt("What node for x?"));
var Px = getNode(P1, x);
var y = parseInt(prompt("What node for y?"));
var Py = getNode(P1, y);
var prev = Px.Prev;
var next = Py.Next;
if (prev != null) {
	prev.Next = next;
}
if (next != null) {
	next.Prev = prev;
}
Px.Prev = Py.Next = null;
document.write("First dequeue:<br />");
printIncDequeue(P1);
var P2 = getLast(P1);
printDecDequeue(P2);
document.write("Second dequeue:<br />");
printIncDequeue(Px);
printDecDequeue(Py);

function makeDequeue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			tempNode.Prev = last;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getNode(node, index) {
	node = getFirst(node);
	while (node != null && index > 1) {
		node = node.Next;
		--index;
	}
	return node;
}

function getFirst(node) {
	if (node == null) {
		return node;
	}
	while (node.Prev != null) {
		node = node.Prev;
	}
	return node;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function printIncDequeue(node) {
	node = getFirst(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}

function printDecDequeue(node) {
	node = getLast(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Prev;
	}
	document.write("<br />");
}
*/
	return false
}

const dynamic54 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

var P1 = makeDequeue();
var x = parseInt(prompt("What node for x?"));
var Px = getNode(P1, x);
var y = parseInt(prompt("What node for y?"));
var Py = getNode(P1, y);

if (Px.Next != Py) {
	Px = Px.Next;
	Py = Py.Prev;
	var prev = Px.Prev;
	var next = Py.Next;
	if (prev != null) {
		prev.Next = next;
	}
	if (next != null) {
		next.Prev = prev;
	}
	Px.Prev = Py.Next = null;
} else {
	Px = Py = null;
}

document.write("First dequeue:<br />");
printIncDequeue(P1);
var P2 = getLast(P1);
printDecDequeue(P2);
document.write("Second dequeue:<br />");
printIncDequeue(Px);
printDecDequeue(Py);

function makeDequeue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			tempNode.Prev = last;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getNode(node, index) {
	node = getFirst(node);
	while (node != null && index > 1) {
		node = node.Next;
		--index;
	}
	return node;
}

function getFirst(node) {
	if (node == null) {
		return node;
	}
	while (node.Prev != null) {
		node = node.Prev;
	}
	return node;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function printIncDequeue(node) {
	node = getFirst(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}

function printDecDequeue(node) {
	node = getLast(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Prev;
	}
	document.write("<br />");
}
*/
	return false
}

const dynamic55 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

var P1 = makeDequeue();
var P2 = getLast(P1);
P1.Prev = P2;
P2.Next = P1;
var i = P1;
do {
	document.write(i.Data + ", ");
	i = i.Next;
} while (i == P1);

function makeDequeue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			tempNode.Prev = last;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}
*/
	return false
}

const dynamic56 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

var P1 = makeDequeue();
var P2 = P1;
var N = 1;
while (P2.Next != null) {
	P2 = P2.Next;
	++N;
}
var half = parseInt(N/2);
var P3 = getNode(P1, half);
var P4 = getNode(P1, half+1);
P3.Next = P1;
P1.Prev = P3;
P4.Prev = P2;
P2.Next = P4;
printCircleDequeue(P1);
printCircleDequeue(P4);

function makeDequeue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			tempNode.Prev = last;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getNode(node, index) {
	node = getFirst(node);
	while (node != null && index > 1) {
		node = node.Next;
		--index;
	}
	return node;
}

function getFirst(node) {
	if (node == null) {
		return node;
	}
	while (node.Prev != null) {
		node = node.Prev;
	}
	return node;
}

function printCircleDequeue(node) {
	var i = node;
	do {
		document.write(i.Data + ", ");
		i = i.Next;
	} while (i != node);
	document.write("<br />");
}
*/
	return false
}

const dynamic57 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

var K = parseInt(prompt("K = "));
var P1 = makeDequeue();
var P2 = getLast(P1);
P1.Prev = P2;
P2.Next = P1;
var i;
for (i = 0; i < K; ++i) {
	P1 = P1.Prev;
}
P2 = P1.Prev;
printCircleDequeue(P1);

function makeDequeue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			tempNode.Prev = last;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function printCircleDequeue(node) {
	var i = node;
	do {
		document.write(i.Data + ", ");
		i = i.Next;
	} while (i != node);
	document.write("<br />");
}
*/
	return false
}

const dynamic58 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

var K = parseInt(prompt("K = "));
var P1 = makeDequeue();
var P2 = getLast(P1);
P1.Prev = P2;
P2.Next = P1;
var i;
for (i = 0; i < K; ++i) {
	P1 = P1.Next;
}
P2 = P1.Prev;
printCircleDequeue(P1);

function makeDequeue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			tempNode.Prev = last;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function printCircleDequeue(node) {
	var i = node;
	do {
		document.write(i.Data + ", ");
		i = i.Next;
	} while (i != node);
	document.write("<br />");
}
*/
	return false
}

const dynamic59 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

function TList(_First, _Last, _Current) {
	this.First = _First;
	this.Last = _Last;
	this.Current = _Current;
};

var P1 = makeDequeue();
var P2 = getLast(P1);
var nomer = parseInt(prompt("Which node is current?"));
var P3 = getNode(P1, nomer);
var list = new TList(P1, P2, P3);
var N = parseInt(prompt("N = "));
var i, number;
for (i = 1; i <= N; ++i) {
	number = parseInt(prompt("number" + i + ':'));
	InsertLast(list, number);
}
printIncList(list.First);
printDecList(list.Last);

function makeDequeue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			tempNode.Prev = last;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function getFirst(node) {
	if (node == null) {
		return node;
	}
	while (node.Prev != null) {
		node = node.Prev;
	}
	return node;
}

function getNode(node, index) {
	node = getFirst(node);
	while (node != null && index > 1) {
		node = node.Next;
		--index;
	}
	return node;
}

function InsertLast(L, D) {
	var newNode = new TNode();
	newNode.Data = D;
	if (L.First == null) {
		L.First = L.Last = L.Current = newNode;
	} else {
		L.Last.Next = newNode;
		newNode.Prev = L.Last;
		L.Last = L.Current = newNode;
	}
}

function printIncList(node) {
	node = getFirst(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}

function printDecList(node) {
	node = getLast(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Prev;
	}
	document.write("<br />");
}
*/
	return false
}

const dynamic60 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

function TList(_First, _Last, _Current) {
	this.First = _First;
	this.Last = _Last;
	this.Current = _Current;
};

var P1 = makeDequeue();
var P2 = getLast(P1);
var nomer = parseInt(prompt("Which node is current?"));
var P3 = getNode(P1, nomer);
var list = new TList(P1, P2, P3);
var N = parseInt(prompt("N = "));
var i, number;
for (i = 1; i <= N; ++i) {
	number = parseInt(prompt("number" + i + ':'));
	InsertFirst(list, number);
}
printIncList(list.First);
printDecList(list.Last);

function makeDequeue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			tempNode.Prev = last;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function getFirst(node) {
	if (node == null) {
		return node;
	}
	while (node.Prev != null) {
		node = node.Prev;
	}
	return node;
}

function getNode(node, index) {
	node = getFirst(node);
	while (node != null && index > 1) {
		node = node.Next;
		--index;
	}
	return node;
}

function InsertFirst(L, D) {
	var newNode = new TNode();
	newNode.Data = D;
	if (L.First == null) {
		L.First = L.Last = L.Current = newNode;
	} else {
		L.First.Prev = newNode;
		newNode.Next = L.First;
		L.First = L.Current = newNode;
	}
}

function printIncList(node) {
	node = getFirst(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}

function printDecList(node) {
	node = getLast(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Prev;
	}
	document.write("<br />");
}
*/
	return false
}

const dynamic61 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

function TList(_First, _Last, _Current) {
	this.First = _First;
	this.Last = _Last;
	this.Current = _Current;
};

var P1 = makeDequeue();
var P2 = getLast(P1);
var nomer = parseInt(prompt("Which node is current?"));
var P3 = getNode(P1, nomer);
var list = new TList(P1, P2, P3);
var i, number;
for (i = 1; i <= 5; ++i) {
	number = parseInt(prompt("number" + i + ':'));
	InsertBefore(list, number);
}
printIncList(list.First);
printDecList(list.Last);

function makeDequeue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			tempNode.Prev = last;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function getFirst(node) {
	if (node == null) {
		return node;
	}
	while (node.Prev != null) {
		node = node.Prev;
	}
	return node;
}

function getNode(node, index) {
	node = getFirst(node);
	while (node != null && index > 1) {
		node = node.Next;
		--index;
	}
	return node;
}

function InsertBefore(L, D) {
	var newNode = new TNode();
	newNode.Data = D;
	if (L.First == null) {
		L.First = L.Last = L.Current = newNode;
	} else {
		newNode.Next = L.Current;
		newNode.Prev = L.Current.Prev;
		if (newNode.Prev != null) {
			newNode.Prev.Next = newNode;
		}
		newNode.Next.Prev = newNode;
		L.Current = newNode;
	}
}

function printIncList(node) {
	node = getFirst(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}

function printDecList(node) {
	node = getLast(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Prev;
	}
	document.write("<br />");
}
*/
	return false
}

const dynamic62 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

function TList(_First, _Last, _Current) {
	this.First = _First;
	this.Last = _Last;
	this.Current = _Current;
};

var P1 = makeDequeue();
var P2 = getLast(P1);
var nomer = parseInt(prompt("Which node is current?"));
var P3 = getNode(P1, nomer);
var list = new TList(P1, P2, P3);
var i, number;
for (i = 1; i <= 5; ++i) {
	number = parseInt(prompt("number" + i + ':'));
	InsertAfter(list, number);
}
printIncList(list.First);
printDecList(list.Last);

function makeDequeue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			tempNode.Prev = last;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function getFirst(node) {
	if (node == null) {
		return node;
	}
	while (node.Prev != null) {
		node = node.Prev;
	}
	return node;
}

function getNode(node, index) {
	node = getFirst(node);
	while (node != null && index > 1) {
		node = node.Next;
		--index;
	}
	return node;
}

function InsertAfter(L, D) {
	var newNode = new TNode();
	newNode.Data = D;
	if (L.First == null) {
		L.First = L.Last = L.Current = newNode;
	} else {
		newNode.Prev = L.Current;
		newNode.Next = L.Current.Next;
		if (newNode.Next != null) {
			newNode.Next.Prev = newNode;
		}
		newNode.Prev.Next = newNode;
		L.Current = newNode;
	}
}

function printIncList(node) {
	node = getFirst(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}

function printDecList(node) {
	node = getLast(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Prev;
	}
	document.write("<br />");
}
*/
	return false
}

const dynamic63 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

function TList(_First, _Last, _Current) {
	this.First = _First;
	this.Last = _Last;
	this.Current = _Current;
};

var P1 = makeDequeue();
var P2 = getLast(P1);
var nomer = parseInt(prompt("Which node is current?"));
var P3 = getNode(P1, nomer);
var list = new TList(P1, P2, P3);
var N = 0;
for (ToFirst(list); list.Current != null; ToNext(list)) {
	++N;
	if (N % 2 != 0) {
		SetData(list, 0);
	}
	if (IsLast(list)) {
		break;
	}
}
document.write("Nodes: " + N + "<br />");
printIncList(list.First);
printDecList(list.Last);

function makeDequeue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			tempNode.Prev = last;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function getFirst(node) {
	if (node == null) {
		return node;
	}
	while (node.Prev != null) {
		node = node.Prev;
	}
	return node;
}

function getNode(node, index) {
	node = getFirst(node);
	while (node != null && index > 1) {
		node = node.Next;
		--index;
	}
	return node;
}

function ToFirst(L) {
	L.Current = L.First;
}

function ToNext(L) {
	if (L.Current != null && L.Current.Next != null) {
		L.Current = L.Current.Next;
	}
}

function SetData(L, D) {
	if (L.Current != null) {
		L.Current.Data = D;
	}
}

function IsLast(L) {
	return L.Current != null && L.Current == L.Last;
}

function printIncList(node) {
	node = getFirst(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}

function printDecList(node) {
	node = getLast(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Prev;
	}
	document.write("<br />");
}
*/
	return false
}

const dynamic64 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

function TList(_First, _Last, _Current) {
	this.First = _First;
	this.Last = _Last;
	this.Current = _Current;
};

var P1 = makeDequeue();
var P2 = getLast(P1);
var nomer = parseInt(prompt("Which node is current?"));
var P3 = getNode(P1, nomer);
var list = new TList(P1, P2, P3);
var data, N = 0;
for (ToLast(list); list.Current != null; ToPrev(list)) {
	++N;
	data = GetData(list);
	if (data % 2 == 0) {
		document.write(data + ", ");
	}
	if (IsFirst(list)) {
		break;
	}
}
document.write("<br />Nodes: " + N);

for (ToFirst(list); list.Current != null; ToNext(list)) {
	++N;
	if (N % 2 != 0) {
		SetData(list, 0);
	}
	if (IsLast(list)) {
		break;
	}
}
document.write("Nodes: " + N + "<br />");
printIncList(list.First);
printDecList(list.Last);

function makeDequeue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			tempNode.Prev = last;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function getFirst(node) {
	if (node == null) {
		return node;
	}
	while (node.Prev != null) {
		node = node.Prev;
	}
	return node;
}

function getNode(node, index) {
	node = getFirst(node);
	while (node != null && index > 1) {
		node = node.Next;
		--index;
	}
	return node;
}

function ToLast(L) {
	L.Current = L.Last;
}

function ToPrev(L) {
	if (L.Current != null && L.Current.Prev != null) {
		L.Current = L.Current.Prev;
	}
}

function GetData(L, D) {
	if (L.Current != null) {
		return L.Current.Data;
	}
}

function IsFirst(L) {
	return L.Current != null && L.Current == L.First;
}

function printIncList(node) {
	node = getFirst(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}

function printDecList(node) {
	node = getLast(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Prev;
	}
	document.write("<br />");
}
*/
	return false
}

const dynamic65 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

function TList(_First, _Last, _Current) {
	this.First = _First;
	this.Last = _Last;
	this.Current = _Current;
};

var P1 = makeDequeue();
var P2 = getLast(P1);
var nomer = parseInt(prompt("Which node is current?"));
var P3 = getNode(P1, nomer);
var list = new TList(P1, P2, P3);
var i;
for (i = 0; i < 5; ++i) {
	document.write(DeleteCurrent(list) + ", ");
}
document.write("<br />");
printIncList(list.First);
printDecList(list.Last);

function makeDequeue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			tempNode.Prev = last;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function getFirst(node) {
	if (node == null) {
		return node;
	}
	while (node.Prev != null) {
		node = node.Prev;
	}
	return node;
}

function getNode(node, index) {
	node = getFirst(node);
	while (node != null && index > 1) {
		node = node.Next;
		--index;
	}
	return node;
}

function DeleteCurrent(L) {
	var tempNode = L.Current;
	
	if (L.First == L.Last) {
		L.First = L.Last = L.Current = null;
	
	} else if (L.Current == L.First) {
		L.First = L.Current = L.First.Next;
		L.First.Prev = null;
	
	} else if (L.Current == L.Last) {
		L.Last = L.Current = L.Last.Prev;
		L.Last.Next = null;
	
	} else {
		L.Current.Next.Prev = L.Current.Prev;
		L.Current.Prev.Next = L.Current.Next;
		L.Current = L.Current.Next;
	}
	tempNode.Next = tempNode.Prev = null;
	return tempNode.Data;
}

function printIncList(node) {
	node = getFirst(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}

function printDecList(node) {
	node = getLast(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Prev;
	}
	document.write("<br />");
}
*/
	return false
}

const dynamic66 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

function TList(_First, _Last, _Current) {
	this.First = _First;
	this.Last = _Last;
	this.Current = _Current;
};

var P1 = makeDequeue();
var P2 = getLast(P1);
var nomer = parseInt(prompt("Which node is current?"));
var P3 = getNode(P1, nomer);
var list1 = new TList(P1, P2, P3);
var list2 = new TList(null, null, null);
SplitList(list1, list2);
document.write("List 1:<br />");
printIncList(list1.First);
printDecList(list1.Last);
document.write("List 2:<br />");
printIncList(list2.First);
printDecList(list2.Last);

function makeDequeue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			tempNode.Prev = last;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function getFirst(node) {
	if (node == null) {
		return node;
	}
	while (node.Prev != null) {
		node = node.Prev;
	}
	return node;
}

function getNode(node, index) {
	node = getFirst(node);
	while (node != null && index > 1) {
		node = node.Next;
		--index;
	}
	return node;
}

function SplitList(L1, L2) {
	L2.First = L2.Current = L1.Current;
	L2.Last = L1.Last;
	if (L1.Current == L1.First) {
		L1.First = L1.Last = L1.Current = null;
	} else {
		L1.Last = L1.Current.Prev;
		L1.Last.Next = null;
		L1.Current = L1.First;
	}
	L2.First.Prev = null;
}

function printIncList(node) {
	node = getFirst(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}

function printDecList(node) {
	node = getLast(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Prev;
	}
	document.write("<br />");
}
*/
	return false
}

const dynamic67 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

function TList(_First, _Last, _Current) {
	this.First = _First;
	this.Last = _Last;
	this.Current = _Current;
};

var P1 = makeDequeue();
var P2 = getLast(P1);
var nomer = parseInt(prompt("Which node is current?"));
var P3 = getNode(P1, nomer);

var P4 = makeDequeue();
var P5 = getLast(P4);
nomer = parseInt(prompt("Which node is current?"));
var P6 = getNode(P4, nomer);

var list1 = new TList(P1, P2, P3);
var list2 = new TList(P4, P5, P6);
AddList(list1, list2);
printIncList(list2.First);
printDecList(list2.Last);

function makeDequeue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			tempNode.Prev = last;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function getFirst(node) {
	if (node == null) {
		return node;
	}
	while (node.Prev != null) {
		node = node.Prev;
	}
	return node;
}

function getNode(node, index) {
	node = getFirst(node);
	while (node != null && index > 1) {
		node = node.Next;
		--index;
	}
	return node;
}

function AddList(L1, L2) {
	L2.Last.Next = L1.First;
	L1.First.Prev = L2.Last;
	L2.Current = L1.First;
	L2.Last = L1.Last;
	L1.First = L1.Last = L1.Current = null;
}

function printIncList(node) {
	node = getFirst(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}

function printDecList(node) {
	node = getLast(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Prev;
	}
	document.write("<br />");
}
*/
	return false
}

const dynamic68 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

function TList(_First, _Last, _Current) {
	this.First = _First;
	this.Last = _Last;
	this.Current = _Current;
};

var P1 = makeDequeue();
var P2 = getLast(P1);
var nomer = parseInt(prompt("Which node is current?"));
var P3 = getNode(P1, nomer);
var list1 = new TList(P1, P2, P3);

var P4 = makeDequeue();
var P5 = getLast(P4);
nomer = parseInt(prompt("Which node is current?"));
var P6 = getNode(P4, nomer);
var list2 = new TList(P4, P5, P6);

InsertList(list1, list2);

printIncList(list2.First);
printDecList(list2.Last);

function makeDequeue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			tempNode.Prev = last;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getFirst(node) {
	if (node == null) {
		return node;
	}
	while (node.Prev != null) {
		node = node.Prev;
	}
	return node;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function getNode(node, index) {
	node = getFirst(node);
	while (node != null && index > 1) {
		node = node.Next;
		--index;
	}
	return node;
}

function InsertList(L1, L2) {
	L1.First.Prev = L2.Current.Prev;
	L1.Last.Next = L2.Current;
	if (L2.Current.Prev != null) {
		L2.Current.Prev.Next = L1.First;
	}
	L2.Current.Prev = L1.Last;
	if (L2.First.Prev != null) {
		L2.First = L1.First;
	}
	L2.Current = L1.First;
	L1.First = L1.Last = L1.Current = null;
}

function printIncList(node) {
	node = getFirst(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Next;
	}
	document.write("<br />");
}

function printDecList(node) {
	node = getLast(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Prev;
	}
	document.write("<br />");
}
*/
	return false
}

const dynamic69 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

function TList(_First, _Last, _Current) {
	this.First = _First;
	this.Last = _Last;
	this.Current = _Current;
};

var P1 = makeDequeue();
var P2 = getLast(P1);
var nomer = parseInt(prompt("Which node is current"));
var P3 = getNode(P1, nomer);
var list1 = new TList(P1, P2, P3);

var P4 = makeDequeue();
var P5 = getLast(P4);
nomer = parseInt(prompt("Which node is current"));
var P6 = getNode(P4, nomer);
var list2 = new TList(P4, P5, P6);

MoveCurrent(list1, list2);

document.write("List 1:<br />");
printIncList(list1.First);
printDecList(list1.Last);

document.write("List 2:<br />");
printIncList(list2.First);
printDecList(list2.Last);

function makeDequeue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			tempNode.Prev = last;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getFirst(node) {
	if (node == null) {
		return node;
	}
	while (node.Prev != null) {
		node = node.Prev;
	}
	return node;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function getNode(node, index) {
	node = getFirst(node);
	while (node != null && index > 1) {
		node = node.Next;
		--index;
	}
	return node;
}

function printIncList(node) {
	node = getFirst(node);
	while (node != null) {
		document.write(node.Data + ", ")
		node = node.Next;
	}
	document.write("<br />");
}

function printDecList(node) {
	node = getLast(node);
	while (node != null) {
		document.write(node.Data + ", ");
		node = node.Prev;
	}
	document.write("<br />");
}

function MoveCurrent(L1, L2) {
	if (L1.First == null || L2.First == null) {
		return;
	}
	tempNode = L1.Current;
	if (L1.First == L1.Last) {
		L1.First = L1.Last = L1.Current = null;
	} else if (L1.Current == L1.First) {
		L1.First = L1.Current = L1.First.Next;
		L1.First.Prev = null;
	} else if (L1.Current == L1.Last) {
		L1.Last = L1.Current = L1.Last.Prev;
		L1.Last.Next = null;
	} else {
		L1.Current.Prev.Next = L1.Current.Next;
		L1.Current.Next.Prev = L1.Current.Prev;
		L1.Current = L1.Current.Next;
	}
	tempNode.Prev = L2.Current;
	tempNode.Next = L2.Current.Next;
	if (tempNode.Next != null) {
		tempNode.Next.Prev = tempNode;
	}
	tempNode.Prev.Next = tempNode;
	L2.Current = L2.Current.Next;
	if (L2.Last.Next != null) {
		L2.Last = L2.Last.Next;
	}
}
*/
	return false
}

const dynamic70 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

var P1 = makeDequeue();
var P2 = getLast(P1);
var barrier = new TNode();
if (P1 == null) {
	barrier.Next = barrier.Prev = barrier;
} else {
	barrier.Next = P1;
	barrier.Prev = P2;
	P1.Prev = P2.Next = barrier;
}
printCircleList(barrier);

function makeDequeue() {
	var i, N = parseInt(prompt("How many nodes?"));
	var first = null, last = null;
	var tempNode;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (first == null) {
			first = last = tempNode;
		} else {
			last.Next = tempNode;
			tempNode.Prev = last;
			last = tempNode;
		}
		tempNode = null;
	}
	return first;
}

function getLast(node) {
	if (node == null) {
		return node;
	}
	while (node.Next != null) {
		node = node.Next;
	}
	return node;
}

function getFirst(node) {
	if (node == null) {
		return node;
	}
	while (node.Prev != null) {
		node = node.Prev;
	}
	return node;
}

function printCircleList(node) {
	var i = node;
	do {
		document.write(i.Data + ", ");
		i = i.Next;
	} while (i != node);
	document.write("<br />");
}
*/
	return false
}

const dynamic71 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

var P1 = makeBarrierList();
var nomer = parseInt(prompt("Which node is current?"));
var P2 = getNode(P1, nomer);
var P3 = new TNode();
P3.Next = P3.Prev = P3;
if (P2 != P1) {
	var prev = P1.Prev;
	P1.Prev = P2.Prev;
	P2.Prev.Next = P1;
	
	P3.Next = P2;
	P3.Prev = prev;
	P2.Prev = P3;
	P3.Prev.Next = P3;
}
printCircleList(P1);
printCircleList(P3);

function makeBarrierList() {
	var i, N = parseInt(prompt("How many nodes?"));
	var tempNode, barrier = null;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (barrier == null) {
			barrier = tempNode;
			barrier.Prev = barrier.Next = barrier;
		} else {
			tempNode.Next = barrier;
			tempNode.Prev = barrier.Prev;
			barrier.Prev.Next = tempNode;
			barrier.Prev = tempNode;
		}
		tempNode = null;
	}
	return barrier;
}

function getNode(barrier, index) {
	var node = barrier;
	while (index > 1) {
		node = node.Next;
		--index;
	}
	return node;
}

function printCircleList(node) {
	var i = node;
	do {
		document.write(i.Data + ", ");
		i = i.Next;
	} while (i != node);
	document.write("<br />");
}
*/
	return false
}

const dynamic72 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

var P1 = makeBarrierList();
var P2 = makeBarrierList();
if (P2.Next != P2) {
	var first = P2.Next;
	var last = P2.Prev;
	first.Prev = P1.Prev;
	last.Next = P1;
	P1.Prev.Next = first;
	P1.Prev = last;
}
P2 = null;
printCircleList(P1);

function makeBarrierList() {
	var i, N = parseInt(prompt("How many nodes?"));
	var tempNode, barrier = null;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (barrier == null) {
			barrier = tempNode;
			barrier.Prev = barrier.Next = barrier;
		} else {
			tempNode.Next = barrier;
			tempNode.Prev = barrier.Prev;
			barrier.Prev.Next = tempNode;
			barrier.Prev = tempNode;
		}
		tempNode = null;
	}
	return barrier;
}

function printCircleList(node) {
	var i = node;
	do {
		document.write(i.Data + ", ");
		i = i.Next;
	} while (i != node);
	document.write("<br />");
}
*/
	return false
}

const dynamic73 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

var P1 = makeBarrierList();
var P2 = makeBarrierList();
if (P1.Next != P1) {
	var first = P1.Next;
	var last = P1.Prev;
	last.Next = P2.Next;
	P2.Next.Prev = last;
	P2.Next = first;
	first.Prev = P2;
}
P1 = null;
printCircleList(P2);

function makeBarrierList() {
	var i, N = parseInt(prompt("How many nodes?"));
	var tempNode, barrier = null;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (barrier == null) {
			barrier = tempNode;
			barrier.Prev = barrier.Next = barrier;
		} else {
			tempNode.Next = barrier;
			tempNode.Prev = barrier.Prev;
			barrier.Prev.Next = tempNode;
			barrier.Prev = tempNode;
		}
		tempNode = null;
	}
	return barrier;
}

function printCircleList(node) {
	var i = node;
	do {
		document.write(i.Data + ", ");
		i = i.Next;
	} while (i != node);
	document.write("<br />");
}
*/
	return false
}

const dynamic74 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

function TListB(_Barrier, _Current) {
	this.Barrier = _Barrier;
	this.Current = _Current;
};

var P1 = makeBarrierList();
var nomer = parseInt(prompt("Which node is current?"));
var P2 = getNode(P1, nomer);
var lb = new TListB(P1, P2);
var N = parseInt(prompt("N = "));
var i, number;
for (i = 1; i <= N; ++i) {
	number = parseInt(prompt("number" + i + ':'));
	LBInsertLast(lb, number);
}
printCircleList(lb.Barrier);

function makeBarrierList() {
	var i, N = parseInt(prompt("How many nodes?"));
	var tempNode, barrier = null;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (barrier == null) {
			barrier = tempNode;
			barrier.Next = barrier.Prev = barrier;
		} else {
			tempNode.Next = barrier;
			tempNode.Prev = barrier.Prev;
			barrier.Prev.Next = tempNode;
			barrier.Prev = tempNode;
		}
		tempNode = null;
	}
	return barrier;
}

function getNode(barrier, index) {
	var node = barrier;
	while (index > 1) {
		node = node.Next;
		--index;
	}
	return node;
}

function LBInsertLast(L, D) {
	var newNode = new TNode();
	newNode.Data = D;
	newNode.Next = L.Barrier;
	newNode.Prev = L.Barrier.Prev;
	L.Barrier.Prev.Next = newNode;
	L.Barrier.Prev = newNode;
	L.Current = newNode;
	newNode = null;
}

function printCircleList(node) {
	var i = node;
	do {
		document.write(i.Data + ", ")
		i = i.Next;
	} while (i != node);
}
*/
	return false
}

const dynamic75 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

function TListB(_Barrier, _Current) {
	this.Barrier = _Barrier;
	this.Current = _Current;
};

var P1 = makeBarrierList();
var nomer = parseInt(prompt("Which node is current?"));
var P2 = getNode(P1, nomer);
var lb = new TListB(P1, P2);
var N = parseInt(prompt("N = "));
var i, number;
for (i = 1; i <= N; ++i) {
	number = parseInt(prompt("number" + i + ':'));
	LBInsertFirst(lb, number);
}
printCircleList(lb.Barrier);

function makeBarrierList() {
	var i, N = parseInt(prompt("How many nodes?"));
	var tempNode, barrier = null;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (barrier == null) {
			barrier = tempNode;
			barrier.Next = barrier.Prev = barrier;
		} else {
			tempNode.Next = barrier;
			tempNode.Prev = barrier.Prev;
			barrier.Prev.Next = tempNode;
			barrier.Prev = tempNode;
		}
		tempNode = null;
	}
	return barrier;
}

function getNode(barrier, index) {
	var node = barrier;
	while (index > 1) {
		node = node.Next;
		--index;
	}
	return node;
}

function LBInsertFirst(L, D) {
	var newNode = new TNode();
	newNode.Data = D;
	newNode.Prev = L.Barrier;
	newNode.Next = L.Barrier.Next;
	L.Barrier.Next.Prev = newNode;
	L.Barrier.Next = newNode;
	L.Current = newNode;
	newNode = null;
}

function printCircleList(node) {
	var i = node;
	do {
		document.write(i.Data + ", ")
		i = i.Next;
	} while (i != node);
}
*/
	return false
}

const dynamic76 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

function TListB(_Barrier, _Current) {
	this.Barrier = _Barrier;
	this.Current = _Current;
};

var P1 = makeBarrierList();
var nomer = parseInt(prompt("Which node is current?"));
var P2 = getNode(P1, nomer);
var lb = new TListB(P1, P2);
var i, number;
for (i = 1; i <= 5; ++i) {
	number = parseInt(prompt("number" + i + ':'));
	LBInsertBefore(lb, number);
}
printCircleList(lb.Barrier);

function makeBarrierList() {
	var i, N = parseInt(prompt("How many nodes?"));
	var tempNode, barrier = null;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (barrier == null) {
			barrier = tempNode;
			barrier.Next = barrier.Prev = barrier;
		} else {
			tempNode.Next = barrier;
			tempNode.Prev = barrier.Prev;
			barrier.Prev.Next = tempNode;
			barrier.Prev = tempNode;
		}
		tempNode = null;
	}
	return barrier;
}

function getNode(barrier, index) {
	var node = barrier;
	while (index > 1) {
		node = node.Next;
		--index;
	}
	return node;
}

function printCircleList(node) {
	var i = node;
	do {
		document.write(i.Data + ", ");
		i = i.Next;
	} while (i != node);
	document.write("<br />");
}

function LBInsertBefore(L, D) {
	var newNode = new TNode();
	newNode.Data = D;
	newNode.Next = L.Current;
	newNode.Prev = L.Current.Prev;
	L.Current.Prev.Next = newNode;
	L.Current.Prev = newNode;
	L.Current = newNode;
	newNode = null;
}
*/
	return false
}

const dynamic77 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

function TListB(_Barrier, _Current) {
	this.Barrier = _Barrier;
	this.Current = _Current;
};

var P1 = makeBarrierList();
var nomer = parseInt(prompt("Which node is current?"));
var P2 = getNode(P1, nomer);
var lb = new TListB(P1, P2);
var i, number;
for (i = 1; i <= 5; ++i) {
	number = parseInt(prompt("number" + i + ':'));
	LBInsertAfter(lb, number);
}
printCircleList(lb.Barrier);

function makeBarrierList() {
	var i, N = parseInt(prompt("How many nodes?"));
	var tempNode, barrier = null;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (barrier == null) {
			barrier = tempNode;
			barrier.Next = barrier.Prev = barrier;
		} else {
			tempNode.Next = barrier;
			tempNode.Prev = barrier.Prev;
			barrier.Prev.Next = tempNode;
			barrier.Prev = tempNode;
		}
		tempNode = null;
	}
	return barrier;
}

function getNode(barrier, index) {
	var node = barrier;
	while (index > 1) {
		node = node.Next;
		--index;
	}
	return node;
}

function printCircleList(node) {
	var i = node;
	do {
		document.write(i.Data + ", ");
		i = i.Next;
	} while (i != node);
	document.write("<br />");
}

function LBInsertAfter(L, D) {
	var newNode = new TNode();
	newNode.Data = D;
	newNode.Prev = L.Current;
	newNode.Next = L.Current.Next;
	L.Current.Next.Prev = newNode;
	L.Current.Next = newNode;
	L.Current = newNode;
	newNode = null;
}
*/
	return false
}

const dynamic78 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

function TListB(_Barrier, _Current) {
	this.Barrier = _Barrier;
	this.Current = _Current;
};

var P1 = makeBarrierList();
var nomer = parseInt(prompt("Which node is current?"));
var P2 = getNode(P1, nomer);
var lb = new TListB(P1, P2);
var i, elements = 0;
for (LBToFirst(lb); !IsBarrier(lb);) {
	LBSetData(lb, 0);
	for (i = 0; !IsBarrier(lb) && i < 2; ++i) {
		LBToNext(lb);
		++elements;
	}
}
document.write("Elements count: " + elements + "<br />");
printCircleList(lb.Barrier);

function makeBarrierList() {
	var i, N = parseInt(prompt("How many nodes?"));
	var tempNode, barrier = null;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (barrier == null) {
			barrier = tempNode;
			barrier.Next = barrier.Prev = barrier;
		} else {
			tempNode.Next = barrier;
			tempNode.Prev = barrier.Prev;
			barrier.Prev.Next = tempNode;
			barrier.Prev = tempNode;
		}
		tempNode = null;
	}
	return barrier;
}

function getNode(barrier, index) {
	var node = barrier;
	while (index > 1) {
		node = node.Next;
		--index;
	}
	return node;
}

function printCircleList(node) {
	var i = node;
	do {
		document.write(i.Data + ", ");
		i = i.Next;
	} while (i != node);
	document.write("<br />");
}

function LBToFirst(L) {
	L.Current = L.Barrier.Next;
}

function LBToNext(L) {
	L.Current = L.Current.Next;
}

function LBSetData(L, D) {
	if (!IsBarrier(L)) {
		L.Current.Data = D;
	}
}

function IsBarrier(L) {
	return L.Current == L.Barrier;
}
*/
	return false
}

const dynamic79 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

function TListB(_Barrier, _Current) {
	this.Barrier = _Barrier;
	this.Current = _Current;
};

var P1 = makeBarrierList();
var nomer = parseInt(prompt("Which node is current?"));
var P2 = getNode(P1, nomer);
var lb = new TListB(P1, P2);
var i, data, elements = 0;
for (LBToLast(lb); !IsBarrier(lb); LBToPrev(lb)) {
	++elements;
	data = LBGetData(lb, 0);
	if (data % 2 == 0) {
		document.write(data + ", ");
	}
}
document.write("<br />Elements count: " + elements + "<br />");
printCircleList(lb.Barrier);

function makeBarrierList() {
	var i, N = parseInt(prompt("How many nodes?"));
	var tempNode, barrier = null;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (barrier == null) {
			barrier = tempNode;
			barrier.Next = barrier.Prev = barrier;
		} else {
			tempNode.Next = barrier;
			tempNode.Prev = barrier.Prev;
			barrier.Prev.Next = tempNode;
			barrier.Prev = tempNode;
		}
		tempNode = null;
	}
	return barrier;
}

function getNode(barrier, index) {
	var node = barrier;
	while (index > 1) {
		node = node.Next;
		--index;
	}
	return node;
}

function printCircleList(node) {
	var i = node;
	do {
		document.write(i.Data + ", ");
		i = i.Next;
	} while (i != node);
	document.write("<br />");
}

function LBToLast(L) {
	L.Current = L.Barrier.Prev;
}

function LBToPrev(L) {
	L.Current = L.Current.Prev;
}

function LBGetData(L) {
	return L.Current.Data;
}

function IsBarrier(L) {
	return L.Current == L.Barrier;
}
*/
	return false
}

const dynamic80 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Next = null;
	this.Prev = null;
};

function TListB(_Barrier, _Current) {
	this.Barrier = _Barrier;
	this.Current = _Current;
};

var P1 = makeBarrierList();
var nomer = parseInt(prompt("Which node is current?"));
var P2 = getNode(P1, nomer);
var lb = new TListB(P1, P2);
var i;
for (i = 0; !IsBarrier(lb) && i < 5; ++i) {
	document.write(LBDeleteCurrent(lb) + ", ");
}
document.write("<br />");
printCircleList(lb.Barrier);

function makeBarrierList() {
	var i, N = parseInt(prompt("How many nodes?"));
	var tempNode, barrier = null;
	for (i = 1; i <= N; ++i) {
		tempNode = new TNode();
		tempNode.Data = parseInt(prompt("Data" + i + ':'));
		if (barrier == null) {
			barrier = tempNode;
			barrier.Next = barrier.Prev = barrier;
		} else {
			tempNode.Next = barrier;
			tempNode.Prev = barrier.Prev;
			barrier.Prev.Next = tempNode;
			barrier.Prev = tempNode;
		}
		tempNode = null;
	}
	return barrier;
}

function getNode(barrier, index) {
	var node = barrier;
	while (index > 1) {
		node = node.Next;
		--index;
	}
	return node;
}

function printCircleList(node) {
	var i = node;
	do {
		document.write(i.Data + ", ");
		i = i.Next;
	} while (i != node);
	document.write("<br />");
}

function LBDeleteCurrent(L) {
	if (IsBarrier(L)) {
		return 0;
	}
	var tempNode = L.Current;
	tempNode.Next.Prev = tempNode.Prev;
	tempNode.Prev.Next = tempNode.Next;
	L.Current = L.Current.Next != L.Barrier ? L.Current.Next : L.Current.Prev;
	return tempNode.Data;
}

function IsBarrier(L) {
	return L.Current == L.Barrier;
}
*/
	return false
}