import {
	initScanner
	, initChecker
	, inputsEOF
	, outputsEOF
} from './io.js'

(async () => {
	for (let taskNo = 1; taskNo <= 100; taskNo++) {
		const taskNoStr = String(taskNo).padStart(3, 0)
		for (let testNo = 1; testNo <= 100; testNo++) {
			const testNoStr = String(testNo).padStart(3, 0)
			const filePath = `../data/kt-gov2/19tree/Tree${taskNoStr}/${testNoStr}/${testNoStr}`

			await initScanner(filePath + '.dat')
			await initChecker(filePath + '.ans')

			let ok = true
			switch (taskNo) {
				case 1: ok = tree1(); break
                case 2: ok = tree2(); break
                case 3: ok = tree3(); break
                case 4: ok = tree4(); break
                case 5: ok = tree5(); break
                case 6: ok = tree6(); break
                case 7: ok = tree7(); break
                case 8: ok = tree8(); break
                case 9: ok = tree9(); break
                case 10: ok = tree10(); break
                case 11: ok = tree11(); break
                case 12: ok = tree12(); break
                case 13: ok = tree13(); break
                case 14: ok = tree14(); break
                case 15: ok = tree15(); break
                case 16: ok = tree16(); break
                case 17: ok = tree17(); break
                case 18: ok = tree18(); break
                case 19: ok = tree19(); break
                case 20: ok = tree20(); break
                case 21: ok = tree21(); break
                case 22: ok = tree22(); break
                case 23: ok = tree23(); break
                case 24: ok = tree24(); break
                case 25: ok = tree25(); break
                case 26: ok = tree26(); break
                case 27: ok = tree27(); break
                case 28: ok = tree28(); break
                case 29: ok = tree29(); break
                case 30: ok = tree30(); break
                case 31: ok = tree31(); break
                case 32: ok = tree32(); break
                case 33: ok = tree33(); break
                case 34: ok = tree34(); break
                case 35: ok = tree35(); break
                case 36: ok = tree36(); break
                case 37: ok = tree37(); break
                case 38: ok = tree38(); break
                case 39: ok = tree39(); break
                case 40: ok = tree40(); break
                case 41: ok = tree41(); break
                case 42: ok = tree42(); break
                case 43: ok = tree43(); break
                case 44: ok = tree44(); break
                case 45: ok = tree45(); break
                case 46: ok = tree46(); break
                case 47: ok = tree47(); break
                case 48: ok = tree48(); break
                case 49: ok = tree49(); break
                case 50: ok = tree50(); break
                case 51: ok = tree51(); break
                case 52: ok = tree52(); break
                case 53: ok = tree53(); break
                case 54: ok = tree54(); break
                case 55: ok = tree55(); break
                case 56: ok = tree56(); break
                case 57: ok = tree57(); break
                case 58: ok = tree58(); break
                case 59: ok = tree59(); break
                case 60: ok = tree60(); break
                case 61: ok = tree61(); break
                case 62: ok = tree62(); break
                case 63: ok = tree63(); break
                case 64: ok = tree64(); break
                case 65: ok = tree65(); break
                case 66: ok = tree66(); break
                case 67: ok = tree67(); break
                case 68: ok = tree68(); break
                case 69: ok = tree69(); break
                case 70: ok = tree70(); break
                case 71: ok = tree71(); break
                case 72: ok = tree72(); break
                case 73: ok = tree73(); break
                case 74: ok = tree74(); break
                case 75: ok = tree75(); break
                case 76: ok = tree76(); break
                case 77: ok = tree77(); break
                case 78: ok = tree78(); break
                case 79: ok = tree79(); break
                case 80: ok = tree80(); break
                case 81: ok = tree81(); break
                case 82: ok = tree82(); break
                case 83: ok = tree83(); break
                case 84: ok = tree84(); break
                case 85: ok = tree85(); break
                case 86: ok = tree86(); break
                case 87: ok = tree87(); break
                case 88: ok = tree88(); break
                case 89: ok = tree89(); break
                case 90: ok = tree90(); break
                case 91: ok = tree91(); break
                case 92: ok = tree92(); break
                case 93: ok = tree93(); break
                case 94: ok = tree94(); break
                case 95: ok = tree95(); break
                case 96: ok = tree96(); break
                case 97: ok = tree97(); break
                case 98: ok = tree98(); break
                case 99: ok = tree99(); break
                case 100: ok = tree100(); break
			}

			if (!ok) {
				console.log("Tree" + taskNoStr + " #" + testNoStr + " has NOT been tested!")
				return
			}

			if (!inputsEOF()) {
				console.log("Tree" + taskNoStr + " #" + testNoStr + " input NOT EOF!")
				return
			}

			if (!outputsEOF()) {
				console.log("Tree" + taskNoStr + " #" + testNoStr + " output NOT EOF!")
				return
			}
		}
		console.log("Tree" + taskNoStr + " has been tested!")
	}
	console.log("Tree has been tested!")
})()

const tree1 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Left = null;
	this.Right = null;
	this.Parent = null;
};

var P1 = new TNode();
P1 = new TNode();
P1.Data = parseInt(prompt("P1.Data = "));
P1.Left = new TNode();
P1.Left.Data = parseInt(prompt("P1.Left.Data = "));
P1.Left.Parent = P1;
P1.Right = new TNode();
P1.Right.Data = parseInt(prompt("P1.Right.Data = "));
P1.Right.Parent = P1;

document.write("Root's data = " + P1.Data + "<br />");
document.write("Left data = " + P1.Left.Data + "<br />");
document.write("Right data = " + P1.Right.Data + "<br />");
document.write("Left = " + P1.Left + "<br />");
document.write("Right = " + P1.Right + "<br />");
*/
	return false
}

const tree2 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Left = null;
	this.Right = null;
	this.Parent = null;
};

var P1 = new TNode();
P1 = new TNode();
P1.Data = parseInt(prompt("root's data = "));
//P1.Left = P1.Right = P1.Parent = null;
makeTree(P1);
alert(getNodeCount(P1));

function makeTree(node) {
	var answer = parseInt(prompt("Where to go? (0-exit; 1-left; 2-right; 3-parent):"));
	if (answer == 0) {
		return;
	}
	if (answer < 0 || answer > 3) {
		makeTree(node);
		return;
	}
	if (answer == 3) {
		makeTree(node.Parent);
		return;
	}
	tempNode = new TNode();
	tempNode.Data = parseInt(prompt("data = "));
	tempNode.Left = tempNode.Right = null;
	tempNode.Parent = node;
	switch (answer) {
		case 1: node.Left = tempNode; break;
		case 2: node.Right = tempNode; break;
	}
	makeTree(tempNode);
}

function getNodeCount(node) {
	if (node == null) {
		return 0;
	}
	var count = 1;
	count += getNodeCount(node.Left);
	count += getNodeCount(node.Right);
	return count;
}
*/
	return false
}

const tree3 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Left = null;
	this.Right = null;
	this.Parent = null;
};

var P1 = new TNode();
P1 = new TNode();
P1.Data = parseInt(prompt("root's data = "));
//P1.Left = P1.Right = P1.Parent = null;
makeTree(P1);
var K = parseInt(prompt("K = "));
alert(getNodeCount(P1, K));

function makeTree(node) {
	var answer = parseInt(prompt("Where to go? (0-exit; 1-left; 2-right; 3-parent):"));
	if (answer == 0) {
		return;
	}
	if (answer < 0 || answer > 3) {
		makeTree(node);
		return;
	}
	if (answer == 3) {
		makeTree(node.Parent);
		return;
	}
	tempNode = new TNode();
	tempNode.Data = parseInt(prompt("data = "));
	tempNode.Left = tempNode.Right = null;
	tempNode.Parent = node;
	switch (answer) {
		case 1: node.Left = tempNode; break;
		case 2: node.Right = tempNode; break;
	}
	makeTree(tempNode);
}

function getNodeCount(node, K) {
	if (node == null) {
		return 0;
	}
	var count = node.Data == K ? 1 : 0;
	count += getNodeCount(node.Left, K);
	count += getNodeCount(node.Right, K);
	return count;
}
*/
	return false
}

const tree4 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Left = null;
	this.Right = null;
	this.Parent = null;
};

var P1 = new TNode();
P1 = new TNode();
P1.Data = parseInt(prompt("root's data = "));
//P1.Left = P1.Right = P1.Parent = null;
makeTree(P1);
alert(getNodeSum(P1));

function makeTree(node) {
	var answer = parseInt(prompt("Where to go? (0-exit; 1-left; 2-right; 3-parent):"));
	if (answer == 0) {
		return;
	}
	if (answer < 0 || answer > 3) {
		makeTree(node);
		return;
	}
	if (answer == 3) {
		makeTree(node.Parent);
		return;
	}
	tempNode = new TNode();
	tempNode.Data = parseInt(prompt("data = "));
	tempNode.Left = tempNode.Right = null;
	tempNode.Parent = node;
	switch (answer) {
		case 1: node.Left = tempNode; break;
		case 2: node.Right = tempNode; break;
	}
	makeTree(tempNode);
}

function getNodeSum(node) {
	if (node == null) {
		return 0;
	}
	var sum = node.Data;
	sum += getNodeSum(node.Left);
	sum += getNodeSum(node.Right);
	return sum;
}
*/
	return false
}

const tree5 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Left = null;
	this.Right = null;
	this.Parent = null;
};

var P1 = new TNode();
P1 = new TNode();
P1.Data = parseInt(prompt("root's data = "));
//P1.Left = P1.Right = P1.Parent = null;
makeTree(P1);
alert(getLeftCount(P1));

function makeTree(node) {
	var answer = parseInt(prompt("Where to go? (0-exit; 1-left; 2-right; 3-parent):"));
	if (answer == 0) {
		return;
	}
	if (answer < 0 || answer > 3) {
		makeTree(node);
		return;
	}
	if (answer == 3) {
		makeTree(node.Parent);
		return;
	}
	tempNode = new TNode();
	tempNode.Data = parseInt(prompt("data = "));
	tempNode.Left = tempNode.Right = null;
	tempNode.Parent = node;
	switch (answer) {
		case 1: node.Left = tempNode; break;
		case 2: node.Right = tempNode; break;
	}
	makeTree(tempNode);
}

function getLeftCount(node) {
	if (node == null) {
		return 0;
	}
	var count = node.Left != null ? 1 : 0;
	count += getLeftCount(node.Left);
	count += getLeftCount(node.Right);
	return count;
}
*/
	return false
}

const tree6 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Left = null;
	this.Right = null;
	this.Parent = null;
};

var P1 = new TNode();
P1 = new TNode();
P1.Data = parseInt(prompt("root's data = "));
//P1.Left = P1.Right = P1.Parent = null;
makeTree(P1);
alert(getLeafCount(P1));

function makeTree(node) {
	var answer = parseInt(prompt("Where to go? (0-exit; 1-left; 2-right; 3-parent):"));
	if (answer == 0) {
		return;
	}
	if (answer < 0 || answer > 3) {
		makeTree(node);
		return;
	}
	if (answer == 3) {
		makeTree(node.Parent);
		return;
	}
	tempNode = new TNode();
	tempNode.Data = parseInt(prompt("data = "));
	tempNode.Left = tempNode.Right = null;
	tempNode.Parent = node;
	switch (answer) {
		case 1: node.Left = tempNode; break;
		case 2: node.Right = tempNode; break;
	}
	makeTree(tempNode);
}

function getLeafCount(node) {
	if (node == null) {
		return 0;
	}
	var count = node.Left == null && node.Right == null ? 1 : 0;
	count += getLeafCount(node.Left);
	count += getLeafCount(node.Right);
	return count;
}
*/
	return false
}

const tree7 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Left = null;
	this.Right = null;
	this.Parent = null;
};

var P1 = new TNode();
P1 = new TNode();
P1.Data = parseInt(prompt("root's data = "));
//P1.Left = P1.Right = P1.Parent = null;
makeTree(P1);
alert(getLeafSum(P1));

function makeTree(node) {
	var answer = parseInt(prompt("Where to go? (0-exit; 1-left; 2-right; 3-parent):"));
	if (answer == 0) {
		return;
	}
	if (answer < 0 || answer > 3) {
		makeTree(node);
		return;
	}
	if (answer == 3) {
		makeTree(node.Parent);
		return;
	}
	tempNode = new TNode();
	tempNode.Data = parseInt(prompt("data = "));
	tempNode.Left = tempNode.Right = null;
	tempNode.Parent = node;
	switch (answer) {
		case 1: node.Left = tempNode; break;
		case 2: node.Right = tempNode; break;
	}
	makeTree(tempNode);
}

function getLeafSum(node) {
	if (node == null) {
		return 0;
	}
	var sum = node.Left == null && node.Right == null ? node.Data : 0;
	sum += getLeafSum(node.Left);
	sum += getLeafSum(node.Right);
	return sum;
}
*/
	return false
}

const tree8 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Left = null;
	this.Right = null;
	this.Parent = null;
};

var P1 = new TNode();
P1 = new TNode();
P1.Data = parseInt(prompt("root's data = "));
//P1.Left = P1.Right = P1.Parent = null;
makeTree(P1);
alert(getRightLeafCount(P1));

function makeTree(node) {
	var answer = parseInt(prompt("Where to go? (0-exit; 1-left; 2-right; 3-parent):"));
	if (answer == 0) {
		return;
	}
	if (answer < 0 || answer > 3) {
		makeTree(node);
		return;
	}
	if (answer == 3) {
		makeTree(node.Parent);
		return;
	}
	tempNode = new TNode();
	tempNode.Data = parseInt(prompt("data = "));
	tempNode.Left = tempNode.Right = null;
	tempNode.Parent = node;
	switch (answer) {
		case 1: node.Left = tempNode; break;
		case 2: node.Right = tempNode; break;
	}
	makeTree(tempNode);
}

function getRightLeafCount(node) {
	if (node == null) {
		return 0;
	}
	var count = node.Right != null && node.Right.Left == null && node.Right.Right == null ? 1 : 0;
	count += getRightLeafCount(node.Left);
	count += getRightLeafCount(node.Right);
	return count;
}
*/
	return false
}

const tree9 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Left = null;
	this.Right = null;
	this.Parent = null;
};

var P1 = new TNode();
P1 = new TNode();
P1.Data = parseInt(prompt("root's data = "));
//P1.Left = P1.Right = P1.Parent = null;
makeTree(P1);
alert(getDeep(P1));

function makeTree(node) {
	var answer = parseInt(prompt("Where to go? (0-exit; 1-left; 2-right; 3-parent):"));
	if (answer == 0) {
		return;
	}
	if (answer < 0 || answer > 3) {
		makeTree(node);
		return;
	}
	if (answer == 3) {
		makeTree(node.Parent);
		return;
	}
	tempNode = new TNode();
	tempNode.Data = parseInt(prompt("data = "));
	tempNode.Left = tempNode.Right = null;
	tempNode.Parent = node;
	switch (answer) {
		case 1: node.Left = tempNode; break;
		case 2: node.Right = tempNode; break;
	}
	makeTree(tempNode);
}

function getDeep(node) {
	if (node == null) {
		return 0;
	}
	var level = node.Parent == 0 ? 0 : 1;
	var left = level + getDeep(node.Left);
	var right = level + getDeep(node.Right);
	return left > right ? left : right;
}
*/
	return false
}

const tree10 = () => {
/*
function TInt() {
	this.value = 0;
};

function TNode() {
	this.Data = 0;
	this.Left = null;
	this.Right = null;
	this.Parent = null;
};

var P1 = new TNode();
P1 = new TNode();
P1.Data = parseInt(prompt("root's data = "));
//P1.Left = P1.Right = P1.Parent = null;
makeTree(P1);
var N = getHeight(P1);
var arr = new Array(N);
var i;
for (i = 0; i < N; ++i) {
	arr[i] = 0;
}
var i = new TInt();
levelNodeCount(P1, arr, i);
document.write(arr.join(", "));

function makeTree(node) {
	var answer = parseInt(prompt("Where to go? (0-exit; 1-left; 2-right; 3-parent):"));
	if (answer == 0) {
		return;
	}
	if (answer < 0 || answer > 3) {
		makeTree(node);
		return;
	}
	if (answer == 3) {
		makeTree(node.Parent);
		return;
	}
	tempNode = new TNode();
	tempNode.Data = parseInt(prompt("data = "));
	tempNode.Left = tempNode.Right = null;
	tempNode.Parent = node;
	switch (answer) {
		case 1: node.Left = tempNode; break;
		case 2: node.Right = tempNode; break;
	}
	makeTree(tempNode);
}

function getHeight(node) {
	if (node == null) {
		return 0;
	}
	var left = 1 + getHeight(node.Left);
	var right = 1 + getHeight(node.Right);
	return left > right ? left : right;
}

function levelNodeCount(node, arr, index) {
	if (node == null) {
		return;
	}
	arr[index.value++]++;
	levelNodeCount(node.Left, arr, index);
	levelNodeCount(node.Right, arr, index);
	index.value--;
}
*/
	return false
}

const tree11 = () => {
/*
function TInt() {
	this.value = 0;
};

function TNode() {
	this.Data = 0;
	this.Left = null;
	this.Right = null;
	this.Parent = null;
};

var P1 = new TNode();
P1.Data = parseInt(prompt("root's data:"));
makeTree(P1);
var N = getHeight(P1);
var arr = new Array(N);
var i;
for (i = 0; i < N; ++i) {
	arr[i] = 0;
}
var index = new TInt();
levelNodeSum(P1, arr, index);
alert(arr.join(" | "));

function makeTree(node) {
	var answer = parseInt(prompt("Where to go? (0-exit; 1-left; 2-right; 3-parent):"));
	if (answer == 0) {
		return;
	}
	if (answer < 0 || answer > 3) {
		makeTree(node);
		return;
	}
	if (answer == 3) {
		makeTree(node.Parent);
		return;
	}
	var tempNode = new TNode();
	tempNode.Data = parseInt(prompt("data:"));
	tempNode.Parent = node;
	switch (answer) {
		case 1: node.Left = tempNode; break;
		case 2: node.Right = tempNode; break;
	}
	makeTree(tempNode);
}

function getHeight(node) {
	if (node == null) {
		return 0;
	}
	var left = 1 + getHeight(node.Left);
	var right = 1 + getHeight(node.Right);
	return left > right ? left : right;
}

function levelNodeSum(node, A, i) {
	if (node == null) {
		return;
	}
	A[i.value++] += node.Data;
	levelNodeSum(node.Left, A, i);
	levelNodeSum(node.Right, A, i);
	i.value--;
}
*/
	return false
}

const tree12 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Left = null;
	this.Right = null;
	this.Parent = null;
};

var P1 = new TNode();
P1.Data = parseInt(prompt("root's data:"));
makeTree(P1);
printInfix(P1);

function makeTree(node) {
	var answer = parseInt(prompt("Where to go? (0-exit; 1-left; 2-right; 3-parent)"));
	if (answer == 0) {
		return;
	}
	if (answer > 3 || answer < 0) {
		makeTree(node);
		return;
	}
	if (answer == 3) {
		makeTree(node.Parent);
		return;
	}
	var newLeaf = new TNode();
	newLeaf.Data = parseInt(prompt("Data:"));
	newLeaf.Parent = node;
	switch (answer) {
		case 1: node.Left = newLeaf; break;
		case 2: node.Right = newLeaf; break;
	}
	makeTree(newLeaf);
}

function printInfix(node) {
	if (node == null) {
		return;
	}
	printInfix(node.Left);
	document.write(node.Data + ", ");
	printInfix(node.Right);
}
*/
	return false
}

const tree13 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Left = null;
	this.Right = null;
	this.Parent = null;
};

var P1 = new TNode();
P1.Data = parseInt(prompt("root's data:"));
makeTree(P1);
printPrefix(P1);

function makeTree(node) {
	var answer = parseInt(prompt("Where to go? (0-exit; 1-left; 2-right; 3-parent)"));
	if (answer == 0) {
		return;
	}
	if (answer > 3 || answer < 0) {
		makeTree(node);
		return;
	}
	if (answer == 3) {
		makeTree(node.Parent);
		return;
	}
	var newLeaf = new TNode();
	newLeaf.Data = parseInt(prompt("Data:"));
	newLeaf.Parent = node;
	switch (answer) {
		case 1: node.Left = newLeaf; break;
		case 2: node.Right = newLeaf; break;
	}
	makeTree(newLeaf);
}

function printPrefix(node) {
	if (node == null) {
		return;
	}
	document.write(node.Data + ", ");
	printPrefix(node.Left);
	printPrefix(node.Right);
}
*/
	return false
}

const tree14 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Left = null;
	this.Right = null;
	this.Parent = null;
};

var P1 = new TNode();
P1.Data = parseInt(prompt("root's data:"));
makeTree(P1);
printPostfix(P1);

function makeTree(node) {
	var answer = parseInt(prompt("Where to go? (0-exit; 1-left; 2-right; 3-parent)"));
	if (answer == 0) {
		return;
	}
	if (answer > 3 || answer < 0) {
		makeTree(node);
		return;
	}
	if (answer == 3) {
		makeTree(node.Parent);
		return;
	}
	var newLeaf = new TNode();
	newLeaf.Data = parseInt(prompt("Data:"));
	newLeaf.Parent = node;
	switch (answer) {
		case 1: node.Left = newLeaf; break;
		case 2: node.Right = newLeaf; break;
	}
	makeTree(newLeaf);
}

function printPostfix(node) {
	if (node == null) {
		return;
	}
	printPostfix(node.Left);
	printPostfix(node.Right);
	document.write(node.Data + ", ");
}
*/
	return false
}

const tree15 = () => {
/*
function TInt() {
	this.value = 0;
};

function TNode() {
	this.Data = 0;
	this.Left = null;
	this.Right = null;
	this.Parent = null;
};

var P1 = new TNode();
P1.Data = parseInt(prompt("root's data:"));
makeTree(P1);
var N = new TNode();
N.value = parseInt(prompt("N = "));
printInfix(P1, N);

function makeTree(node) {
	var answer = parseInt(prompt("Where to go? (0-exit; 1-left; 2-right; 3-parent)"));
	if (answer == 0) {
		return;
	}
	if (answer > 3 || answer < 0) {
		makeTree(node);
		return;
	}
	if (answer == 3) {
		makeTree(node.Parent);
		return;
	}
	var newLeaf = new TNode();
	newLeaf.Data = parseInt(prompt("Data:"));
	newLeaf.Parent = node;
	switch (answer) {
		case 1: node.Left = newLeaf; break;
		case 2: node.Right = newLeaf; break;
	}
	makeTree(newLeaf);
}

function printInfix(node, N) {
	if (node == null || N.value < 1) {
		return;
	}
	printInfix(node.Left, N);
	if (N.value > 0) {
		document.write(node.Data + ", ");
		N.value--;
	}
	printInfix(node.Right, N);
}
*/
	return false
}

const tree16 = () => {
/*
function TInt() {
	this.value = 0;
};

function TNode() {
	this.Data = 0;
	this.Left = null;
	this.Right = null;
	this.Parent = null;
};

var P1 = new TNode();
P1.Data = parseInt(prompt("root's data:"));
makeTree(P1);
var nomer = new TInt();
nomer.value = 1;
var N = parseInt(prompt("N = "));
printPostfix(P1, nomer, N);

function makeTree(node) {
	var answer = parseInt(prompt("Where to go? (0-exit; 1-left; 2-right; 3-parent)"));
	if (answer == 0) {
		return;
	}
	if (answer > 3 || answer < 0) {
		makeTree(node);
		return;
	}
	if (answer == 3) {
		makeTree(node.Parent);
		return;
	}
	var newLeaf = new TNode();
	newLeaf.Data = parseInt(prompt("Data:"));
	newLeaf.Parent = node;
	switch (answer) {
		case 1: node.Left = newLeaf; break;
		case 2: node.Right = newLeaf; break;
	}
	makeTree(newLeaf);
}

function printPostfix(node, index, N) {
	if (node == null) {
		return;
	}
	printPostfix(node.Left, index, N);
	printPostfix(node.Right, index, N);
	if (index.value >= N) {
		document.write(node.Data + ", ");
	}
	index.value++;
}
*/
	return false
}

const tree17 = () => {
/*
function TInt() {
	this.value = 0;
};

function TNode() {
	this.Data = 0;
	this.Left = null;
	this.Right = null;
	this.Parent = null;
};

var P1 = new TNode();
P1.Data = parseInt(prompt("root's data:"));
makeTree(P1);
var N1 = parseInt(prompt("N1 = "));
var N2 = parseInt(prompt("N2 = "));
var nomer = new TInt();
nomer.value = 1;
printPrefix(P1, nomer, N1, N2);

function makeTree(node) {
	var answer = parseInt(prompt("Where to go? (0-exit; 1-left; 2-right; 3-parent)"));
	if (answer == 0) {
		return;
	}
	if (answer > 3 || answer < 0) {
		makeTree(node);
		return;
	}
	if (answer == 3) {
		makeTree(node.Parent);
		return;
	}
	var newLeaf = new TNode();
	newLeaf.Data = parseInt(prompt("Data:"));
	newLeaf.Parent = node;
	switch (answer) {
		case 1: node.Left = newLeaf; break;
		case 2: node.Right = newLeaf; break;
	}
	makeTree(newLeaf);
}

function printPrefix(node, index, N1, N2) {
	if (node == null) {
		return;
	}
	if (N1 <= index.value && index.value <= N2) {
		document.write(node.Data + ", ")
	}
	index.value++;
	printPrefix(node.Left, index, N1, N2);
	printPrefix(node.Right, index, N1, N2);
}
*/
	return false
}

const tree18 = () => {
/*
function TInt() {
	this.value = 0;
}

function TNode() {
	this.Data = 0;
	this.Left = null;
	this.Right = null;
	this.Parent = null;
};

var P1 = new TNode();
P1.Data = parseInt(prompt("Root's Data:"));
makeTree(P1);
printInfix(P1);
document.write("<br />");
var L = parseInt(prompt("L = "));
var i = new TInt();
i.value = -1;
var N = levelNodes(P1, i, L);
document.write("<br />N = " + N);

function makeTree(node) {
	var answer = parseInt(prompt("Where to go? (0-exit; 1-left; 2-right; 3-parent)"));
	if (answer == 0) {
		return;
	}
	if (answer > 3 || answer < 0) {
		makeTree(node);
		return;
	}
	if (answer == 3) {
		makeTree(node.Parent);
		return;
	}
	var newLeaf = new TNode();
	newLeaf.Data = parseInt(prompt("Data:"));
	newLeaf.Parent = node;
	switch (answer) {
		case 1: node.Left = newLeaf; break;
		case 2: node.Right = newLeaf; break;
	}
	makeTree(newLeaf);
}

function levelNodes(node, i, L) {
	if (node == null) {
		return 0;
	}
	i.value++;
	var count = 0;
	if (i.value == L) {
		document.write(node.Data + ", ");
		++count;
	}
	count += levelNodes(node.Left, i, L);
	count += levelNodes(node.Right, i, L);
	i.value--;
	return count;
}

function printInfix(node) {
	if (node == null) {
		return;
	}
	printInfix(node.Left);
	document.write(node.Data + ", ");
	printInfix(node.Right);
}
*/
	return false
}

const tree19 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Left = null;
	this.Right = null;
	this.Parent = null;
};

var P1 = new TNode();
P1.Data = parseInt(prompt("Root's Data:"));
makeTree(P1);
printInfix(P1);
var maximal = getMaxData(P1);
var maxCount = getMaxCount(P1, maximal);
document.write("<br />maximal = " + maximal);
document.write("<br />maximal count: " + maxCount);

function makeTree(node) {
	var answer = parseInt(prompt("Where to go? (0-exit; 1-left; 2-right; 3-parent)"));
	if (answer == 0) {
		return;
	}
	if (answer < 0 || answer > 3) {
		makeTree(node);
		return;
	}
	if (answer == 3) {
		makeTree(node.Parent);
		return;
	}
	var newLeaf = new TNode();
	newLeaf.Data = parseInt(prompt("Data:"));
	newLeaf.Parent = node;
	switch (answer) {
		case 1: node.Left = newLeaf; break;
		case 2: node.Right = newLeaf; break;
	}
	makeTree(newLeaf);
}

function printInfix(node) {
	if (node == null) {
		return;
	}
	printInfix(node.Left);
	document.write(node.Data + ", ");
	printInfix(node.Right);
}

function getMaxData(node) {
	var data, maximal = node.Data;
	if (node.Left != null) {
		data = getMaxData(node.Left);
		if (data > maximal) {
			maximal = data;
		}
	}
	if (node.Right != null) {
		data = getMaxData(node.Right);
		if (data > maximal) {
			maximal = data;
		}
	}
	return maximal;
}

function getMaxCount(node, maximal) {
	if (node == null) {
		return 0;
	}
	var count = node.Data == maximal ? 1 : 0;
	count += getMaxCount(node.Left, maximal);
	count += getMaxCount(node.Right, maximal);
	return count;
}
*/
	return false
}

const tree20 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Left = null;
	this.Right = null;
	this.Parent = null;
};

var P1 = new TNode();
P1.Data = parseInt(prompt("Root's Data:"));
makeTree(P1);
printInfix(P1);
var minimal = getMinData(P1);
var maxCount = getMinLeavesCount(P1, minimal);
document.write("<br />minimal = " + minimal);
document.write("<br />minimal count: " + maxCount);

function makeTree(node) {
	var answer = parseInt(prompt("Where to go? (0-exit; 1-left; 2-right; 3-parent)"));
	if (answer == 0) {
		return;
	}
	if (answer < 0 || answer > 3) {
		makeTree(node);
		return;
	}
	if (answer == 3) {
		makeTree(node.Parent);
		return;
	}
	var newLeaf = new TNode();
	newLeaf.Data = parseInt(prompt("Data:"));
	newLeaf.Parent = node;
	switch (answer) {
		case 1: node.Left = newLeaf; break;
		case 2: node.Right = newLeaf; break;
	}
	makeTree(newLeaf);
}

function printInfix(node) {
	if (node == null) {
		return;
	}
	printInfix(node.Left);
	document.write(node.Data + ", ");
	printInfix(node.Right);
}

function getMinData(node) {
	var data, minimal = node.Data;
	if (node.Left != null) {
		data = getMinData(node.Left);
		if (data < minimal) {
			minimal = data;
		}
	}
	if (node.Right != null) {
		data = getMinData(node.Right);
		if (data < minimal) {
			minimal = data;
		}
	}
	return minimal;
}

function getMinLeavesCount(node, minimal) {
	if (node == null) {
		return 0;
	}
	var count = 0;
	if (node.Data == minimal && node.Left == null && node.Right == null) {
		++count;
	}
	count += getMinLeavesCount(node.Left, minimal);
	count += getMinLeavesCount(node.Right, minimal);
	return count;
}
*/
	return false
}

const tree21 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Left = null;
	this.Right = null;
};

var P1 = new TNode();
P1.Data = parseInt(prompt("Root's Data"));
makeTree(P1);
printInfix(P1);
document.write("<br />");
document.write("minimal from leaves: " + getMinLeaf(P1));

function makeTree(node) {
	var answer = parseInt(prompt("Where to go? (0-exit; 1-left; 2-right; 3-parent)"));
	if (answer == 0) {
		return;
	}
	if (answer > 3 || answer < 0) {
		makeTree(node);
		return;
	}
	if (answer == 3) {
		makeTree(node.Parent);
		return;
	}
	var newLeaf = new TNode();
	newLeaf.Data = parseInt(prompt("Data:"));
	newLeaf.Parent = node;
	switch (answer) {
		case 1: node.Left = newLeaf; break;
		case 2: node.Right = newLeaf; break;
	}
	makeTree(newLeaf);
}

function printInfix(node) {
	if (node == null) {
		return;
	}
	printInfix(node.Left);
	document.write(node.Data + ", ");
	printInfix(node.Right);
}

function getMinLeaf(node) {
	var arr = [];
	var index = 0;
	if (node.Left == null && node.Right == null) {
		arr[index++] = node.Data;
	}
	if (node.Left != null) {
		arr[index++] = getMinLeaf(node.Left);
	}
	if (node.Right != null) {
		arr[index++] = getMinLeaf(node.Right);
	}
	var i, minimal = arr[0];
	for (i = 1; i < arr.length; ++i) {
		if (arr[i] < minimal) {
			minimal = arr[i];
		}
	}
	return minimal;
}
*/
	return false
}

const tree22 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Left = null;
	this.Right = null;
	this.Parent = null;
};

var P1 = new TNode();
P1.Data = parseInt(prompt("Root's Data:"));
makeTree(P1);
printInfix(P1);
document.write("<br />maximal = " + getMax(P1));

function makeTree(node) {
	var answer = parseInt(prompt("Where to go? (0-exit; 1-left; 2-right; 3-parent)"));
	if (answer == 0) {
		return;
	}
	if (answer < 0 || answer > 3) {
		makeTree(node);
		return;
	}
	if (answer == 3) {
		makeTree(node.Parent);
		return;
	}
	var newLeaf = new TNode();
	newLeaf.Data = parseInt(prompt("Data:"));
	newLeaf.Parent = node;
	switch (answer) {
		case 1: node.Left = newLeaf; break;
		case 2: node.Right = newLeaf; break;
	}
	makeTree(newLeaf);
}

function printInfix(node) {
	if (node == null) {
		return;
	}
	printInfix(node.Left);
	document.write(node.Data + ", ");
	printInfix(node.Right);
}

function getMax(node) {
	var arr = [];
	var index = 0;
	if (node.Left != null || node.Right != null) {
		arr[index++] = node.Data;
	}
	if (node.Left != null && (node.Left.Left != null || node.Left.Right != null)) {
		arr[index++] = getMax(node.Left);
	}
	if (node.Right != null && (node.Right.Left != null || node.Right.Right != null)) {
		arr[index++] = getMax(node.Right);
	}
	var i, maximal = arr[0];
	for (i = 1; i < arr.length; ++i) {
		if (arr[i] > maximal) {
			maximal = arr[i];
		}
	}
	return maximal;
}
*/
	return false
}

const tree23 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Left = null;
	this.Right = null;
	this.Parent = null;
};

var P1 = new TNode();
P1.Data = parseInt(prompt("Root's Data:"));
makeTree(P1);
printInfix(P1);
var minimal = getMin(P1);
var minNode = getMinNode(P1, minimal);
document.write("<br />minimal = " + minNode.Data);

function makeTree(node) {
	var answer = parseInt(prompt("Where to go? (0-exit; 1-left; 2-right; 3-parent)"));
	if (answer == 0) {
		return;
	}
	if (answer > 3 || answer < 0) {
		makeTree(node);
		return;
	}
	if (answer == 3) {
		makeTree(node.Parent);
		return;
	}
	var newLeaf = new TNode();
	newLeaf.Data = parseInt(prompt("Data:"));
	newLeaf.Parent = node;
	switch (answer) {
		case 1: node.Left = newLeaf; break;
		case 2: node.Right = newLeaf; break;
	}
	makeTree(newLeaf);
}

function printInfix(node) {
	if (node == null) {
		return;
	}
	printInfix(node.Left);
	document.write(node.Data + ", ");
	printInfix(node.Right);
}

function getMin(node) {
	var minimal = node.Data;
	var data;
	if (node.Left != null) {
		data = getMin(node.Left);
		if (data < minimal) {
			minimal = data;
		}
	}
	if (node.Right != null) {
		data = getMin(node.Right);
		if (data < minimal) {
			minimal = data;
		}
	}
	return minimal;
}

function getMinNode(node, data) {
	if (node == null) {
		return node;
	}
	if (node.Data == data) {
		return node;
	}
	var minNode = getMinNode(node.Left, data);
	if (minNode == null) {
		minNode = getMinNode(node.Right, data);
	}
	return minNode;
}
*/
	return false
}

const tree24 = () => {
/*
function TNode() {
	this.Data = 0;
	this.Left = null;
	this.Right = null;
	this.Parent = null;
};

var P1 = new TNode();
P1.Data = parseInt(prompt("Root's Data:"));
makeTree(P1);
printInfix(P1);
document.write("<br />");
var node = null;
if (hasEvenData(P1)) {
	node = getMaxEvenDataNode(P1);
}

function makeTree(node) {
	var answer = parseInt(prompt("Where to go? (0-exit; 1-left; 2-right; 3-parent)"));
	if (answer == 0) { return ; }
	if (answer > 3 || answer < 0) { makeTree(node); return; }
	if (answer == 3) { makeTree(node.Parent); return; }
	var newLeaf = new TNode();
	newLeaf.Data = parseInt(prompt("Data:"));
	newLeaf.Parent = node;
	switch (answer) {
		case 1: node.Left = newLeaf; break;
		case 2: node.Right = newLeaf; break;
	}
	makeTree(newLeaf);
}

function printInfix(node) {
	if (node == null) { return; }
	printInfix(node.Left);
	document.write(node.Data + ", ");
	printInfix(node.Right);
}

function hasEvenData(node) {
	if (node == null) {
		return false;
	}
	if (node.Data % 2 == 0) {
		return true;
	}
	var even = hasEvenData(node.Left);
	if (!even) {
		even = hasEvenData(node.Right);
	}
	return even;
}

function getMaxEvenDataNode(node) {
	//continue;
}
*/
    return false
}

const tree25 = () => {
	return false
}

const tree26 = () => {
	return false
}

const tree27 = () => {
	return false
}

const tree28 = () => {
	return false
}

const tree29 = () => {
	return false
}

const tree30 = () => {
	return false
}

const tree31 = () => {
	return false
}

const tree32 = () => {
	return false
}

const tree33 = () => {
	return false
}

const tree34 = () => {
	return false
}

const tree35 = () => {
	return false
}

const tree36 = () => {
	return false
}

const tree37 = () => {
	return false
}

const tree38 = () => {
	return false
}

const tree39 = () => {
	return false
}

const tree40 = () => {
	return false
}

const tree41 = () => {
	return false
}

const tree42 = () => {
	return false
}

const tree43 = () => {
	return false
}

const tree44 = () => {
	return false
}

const tree45 = () => {
	return false
}

const tree46 = () => {
	return false
}

const tree47 = () => {
	return false
}

const tree48 = () => {
	return false
}

const tree49 = () => {
	return false
}

const tree50 = () => {
	return false
}

const tree51 = () => {
	return false
}

const tree52 = () => {
	return false
}

const tree53 = () => {
	return false
}

const tree54 = () => {
	return false
}

const tree55 = () => {
	return false
}

const tree56 = () => {
	return false
}

const tree57 = () => {
	return false
}

const tree58 = () => {
	return false
}

const tree59 = () => {
	return false
}

const tree60 = () => {
	return false
}

const tree61 = () => {
	return false
}

const tree62 = () => {
	return false
}

const tree63 = () => {
	return false
}

const tree64 = () => {
	return false
}

const tree65 = () => {
	return false
}

const tree66 = () => {
	return false
}

const tree67 = () => {
	return false
}

const tree68 = () => {
	return false
}

const tree69 = () => {
	return false
}

const tree70 = () => {
	return false
}

const tree71 = () => {
	return false
}

const tree72 = () => {
	return false
}

const tree73 = () => {
	return false
}

const tree74 = () => {
	return false
}

const tree75 = () => {
	return false
}

const tree76 = () => {
	return false
}

const tree77 = () => {
	return false
}

const tree78 = () => {
	return false
}

const tree79 = () => {
	return false
}

const tree80 = () => {
	return false
}

const tree81 = () => {
	return false
}

const tree82 = () => {
	return false
}

const tree83 = () => {
	return false
}

const tree84 = () => {
	return false
}

const tree85 = () => {
	return false
}

const tree86 = () => {
	return false
}

const tree87 = () => {
	return false
}

const tree88 = () => {
	return false
}

const tree89 = () => {
	return false
}

const tree90 = () => {
	return false
}

const tree91 = () => {
	return false
}

const tree92 = () => {
	return false
}

const tree93 = () => {
	return false
}

const tree94 = () => {
	return false
}

const tree95 = () => {
	return false
}

const tree96 = () => {
	return false
}

const tree97 = () => {
	return false
}

const tree98 = () => {
	return false
}

const tree99 = () => {
	return false
}

const tree100 = () => {
	return false
}