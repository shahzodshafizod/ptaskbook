import {
	initScanner
	, initChecker
	, inputsEOF
	, outputsEOF
} from './io.js'

(async () => {
	for (let taskNo = 1; taskNo <= 70; taskNo++) {
		const taskNoStr = String(taskNo).padStart(3, 0)
		for (let testNo = 1; testNo <= 100; testNo++) {
			const testNoStr = String(testNo).padStart(3, 0)
			const filePath = `../data/kt-gov2/16param/Param${taskNoStr}/${testNoStr}/${testNoStr}`

			await initScanner(filePath + '.dat')
			await initChecker(filePath + '.ans')

			let ok = true
			switch (taskNo) {
				case 1: ok = param1(); break
				case 2: ok = param2(); break
				case 3: ok = param3(); break
				case 4: ok = param4(); break
				case 5: ok = param5(); break
				case 6: ok = param6(); break
				case 7: ok = param7(); break
				case 8: ok = param8(); break
				case 9: ok = param9(); break
				case 10: ok = param10(); break
				case 11: ok = param11(); break
				case 12: ok = param12(); break
				case 13: ok = param13(); break
				case 14: ok = param14(); break
				case 15: ok = param15(); break
				case 16: ok = param16(); break
				case 17: ok = param17(); break
				case 18: ok = param18(); break
				case 19: ok = param19(); break
				case 20: ok = param20(); break
				case 21: ok = param21(); break
				case 22: ok = param22(); break
				case 23: ok = param23(); break
				case 24: ok = param24(); break
				case 25: ok = param25(); break
				case 26: ok = param26(); break
				case 27: ok = param27(); break
				case 28: ok = param28(); break
				case 29: ok = param29(); break
				case 30: ok = param30(); break
				case 31: ok = param31(); break
				case 32: ok = param32(); break
				case 33: ok = param33(); break
				case 34: ok = param34(); break
				case 35: ok = param35(); break
				case 36: ok = param36(); break
				case 37: ok = param37(); break
				case 38: ok = param38(); break
				case 39: ok = param39(); break
				case 40: ok = param40(); break
				case 41: ok = param41(); break
				case 42: ok = param42(); break
				case 43: ok = param43(); break
				case 44: ok = param44(); break
				case 45: ok = param45(); break
				case 46: ok = param46(); break
				case 47: ok = param47(); break
				case 48: ok = param48(); break
				case 49: ok = param49(); break
				case 50: ok = param50(); break
				case 51: ok = param51(); break
				case 52: ok = param52(); break
				case 53: ok = param53(); break
				case 54: ok = param54(); break
				case 55: ok = param55(); break
				case 56: ok = param56(); break
				case 57: ok = param57(); break
				case 58: ok = param58(); break
				case 59: ok = param59(); break
				case 60: ok = param60(); break
				case 61: ok = param61(); break
				case 62: ok = param62(); break
				case 63: ok = param63(); break
				case 64: ok = param64(); break
				case 65: ok = param65(); break
				case 66: ok = param66(); break
				case 67: ok = param67(); break
				case 68: ok = param68(); break
				case 69: ok = param69(); break
				case 70: ok = param70(); break
			}

			if (!ok) {
				console.log("Param" + taskNoStr + " #" + testNoStr + " has NOT been tested!")
				return
			}

			if (!inputsEOF()) {
				console.log("Param" + taskNoStr + " #" + testNoStr + " input NOT EOF!")
				return
			}

			if (!outputsEOF()) {
				console.log("Param" + taskNoStr + " #" + testNoStr + " output NOT EOF!")
				return
			}
		}
		console.log("Param" + taskNoStr + " has been tested!")
	}
	console.log("Param has been tested!")
})()

// { "no": 1, "dat": "", "ans": "" }
const param1 = () => {
	/*
let i, j, N
let arr = null
for (i = 1; i <= 3; ++i) {
	N = parseInt(prompt("N" + i + " = "))
	arr = new Array(N)
	for (j = 0; j < N; ++j) {
		arr[j] = parseInt(prompt("arr[" + j + "] = "))
	}
	alert(MinElem(arr, N))
	arr.splice(0, arr.length)
}

function MinElem(A, N) {
	let minimal = A[0]
	let i
	for (i = 1; i < N; ++i) {
		if (A[i] < minimal) {
			minimal = A[i]
		}
	}
	return minimal
}
	*/
	return false
}

// { "no": 2, "dat": "2", "ans": "" }
const param2 = () => {
	/*
let i, j, N
let arr = null
for (i = 1; i <= 3; ++i) {
	N = parseInt(prompt("N" + i + " = "))
	arr = new Array(N)
	for (j = 0; j < N; ++j) {
		arr[j] = parseFloat(prompt("arr" + j + "] = "))
	}
	alert(MaxNum(arr, N))
	arr.splice(0, arr.length)
}

function MaxNum(A, N) {
	let i, num = 0
	for (i = 1; i < N; ++i) {
		if (A[i] > A[num]) {
			num = i
		}
	}
	return num+1
}
	*/
	return false
}

// { "no": 3, "dat": "2", "ans": "" }
const param3 = () => {
	/*
function TInt() {
	this.value = 0
}

let i, j, N
let arr = null
let minIndex = new TInt()
let maxIndex = new TInt()
for (i = 1; i <= 3; ++i) {
	N = parseInt(prompt("N" + i + " = "))
	arr = new Array(N)
	for (j = 0; j < N; ++j) {
		arr[j] = parseFloat(prompt("arr[" + j + "] = "))
	}
	MinmaxNum(arr, N, minIndex, maxIndex)
	console.log("NMin = " + minIndex.value + "  NMax = " + maxIndex.value)
	arr.splice(0, arr.length)
}

function MinmaxNum(A, N, NMin, NMax) {
	NMin.value = NMax.value = 0
	let i
	for (i = 1; i < N; ++i) {
		if (A[i] > A[NMax.value]) {
			NMax.value = i
		}
		if (A[i] < A[NMin.value]) {
			NMin.value = i
		}
	}
	NMin.value++
	NMax.value++
}
	*/
	return false
}

// { "no": 4, "dat": "2", "ans": "2" }
const param4 = () => {
	/*
let i, j, N
let arr = null
for (i = 1; i <= 3; ++i) {
	N = parseInt(prompt("N" + i + " = "))
	arr = new Array(N)
	for (j = 0; j < N; ++j) {
		arr[j] = parseFloat(prompt("arr[" + j + "] = "))
	}
	Inv(arr, N)
	for (j = 0; j < N; ++j) {
		console.log(arr[j].toFixed(2) + " ")
	}
	console.log("\n")
	arr.splice(0, arr.length)
}

function Inv(A, N) {
	A.reverse()
}
	*/
	return false
}

// { "no": 5, "dat": "2", "ans": "2" }
const param5 = () => {
	/*
let N = parseInt(prompt("N = "))
let arr = new Array(N)
let i
for (i = 0; i < N; ++i) {
	arr[i] = parseFloat(prompt("arr[" + i + "] = "))
}
for (i = 0; i < 5; ++i) {
	Smooth1(arr, N)
	printArr(arr, N)
}

function Smooth1(A, N) {
	let i, sum = 0
	for (i = 0; i < N; ++i) {
		sum += A[i]
		A[i] = sum / (i+1)
	}
}

function printArr(A, N) {
	let i
	for (i = 0; i < N; ++i) {
		console.log(A[i].toFixed(2) + " ")
	}
	console.log("\n")
}
	*/
	return false
}

// { "no": 6, "dat": "2", "ans": "2" }
const param6 = () => {
	/*
let i, N = parseInt(prompt("N = "))
let arr = new Array(N)
for (i = 0; i < N; ++i) {
	arr[i] = parseFloat(prompt("arr[" + i + "] = "))
}
for (i = 0; i < 5; ++i) {
	Smooth2(arr, N)
	printArr(arr, N)
}

function Smooth2(A, N) {
	let i, temp, sum = A[0]
	for (i = 1; i < N; ++i) {
		temp = A[i]
		sum += A[i]
		A[i] = sum / 2
		sum = temp
	}
}

function printArr(A, N) {
	let i
	for (i = 0; i < N; ++i) {
		console.log(A[i].toFixed(2) + " ")
	}
	console.log("\n")
}
	*/
	return false
}

// { "no": 7, "dat": "2", "ans": "2" }
const param7 = () => {
	/*
let i, N = parseInt(prompt("N = "))
let arr = new Array(N)
for (i = 0; i < N; ++i) {
	arr[i] = parseFloat(prompt("arr[" + i + "] = "))
}
for (i = 0; i < 5; ++i) {
	Smooth3(arr, N)
	printArr(arr, N)
}

function Smooth3(A, N) {
	let tempA = A.slice(0, N)
	let i
	if (N > 1) {
		A[0] = (tempA[0] + tempA[1]) / 2
		A[N-1] = (tempA[N-2] + tempA[N-1]) / 2
	}
	for (i = 1; i < N-1; ++i) {
		A[i] = (tempA[i-1] + tempA[i] + tempA[i+1]) / 3
	}
}

function printArr(A, N) {
	let i
	for (i = 0; i < N; ++i) {
		console.log(A[i].toFixed(2) + " ")
	}
	console.log("\n")
}
	*/
	return false
}

// { "no": 8, "dat": "", "ans": "" }
const param8 = () => {
	/*
function TInt() {
	this.value = 0
}

let i, j, x
let n = new TInt()
let arr = null
for (i = 1; i <= 3; ++i) {
	x = parseInt(prompt("X" + i + " = "))
	n.value = parseInt(prompt("N" + i + " = "))
	arr = new Array(n.value)
	for (j = 0; j < n.value; ++j) {
		arr[j] = parseFloat(prompt("arr[" + j + "] = "))
	}
	RemoveX(arr, n, x)
	console.log("N" + i + " = " + n.value + " : " + arr.join(", ") + "\n")
}

function RemoveX(A, N, X) {
	let i
	for (i = 0; i < N.value;) {
		if (A[i] == X) {
			A.splice(i, 1)
			N.value--
		} else {
			++i
		}
	}
}
	*/
	return false
}

// { "no": 9, "dat": "2", "ans": "2" }
const param9 = () => {
	/*
function TInt() {
	this.value = 0
}

let n = new TInt()
let i, j
let arr = null
for (i = 1; i <= 3; ++i) {
	n.value = parseInt(prompt("N" + i + " = "))
	arr = new Array(n.value)
	for (j = 0; j < n.value; ++j) {
		arr[j] = parseFloat(prompt("arr[" + j + "] = "))
	}
	RemoveForInc(arr, n)
	console.log("N" + i + " = " + n.value + " : ")
	printArr(arr, n.value)
	arr.splice(0, arr.length)
}

function RemoveForInc(A, N) {
	let i
	for (i = 1; i < N.value;) {
		if (A[i] < A[i-1]) {
			A.splice(i, 1)
			N.value--
		} else {
			++i
		}
	}
}

function printArr(A, N) {
	let i
	for (i = 0; i < N; ++i) {
		console.log(A[i].toFixed(2) + " ")
	}
	console.log("\n")
}
	*/
	return false
}

// { "no": 10, "dat": "", "ans": "" }
const param10 = () => {
	/*
function TInt() {
	this.value = 0
}

let x, i, j
let n = new TInt()
let arr = null
for (i = 1; i <= 3; ++i) {
	x = parseInt(prompt("X" + i + " = "))
	n.value = parseInt(prompt("N" + i + " = "))
	arr = new Array(n.value)
	for (j = 0; j < n.value; ++j) {
		arr[j] = parseInt(prompt("arr[" + j + "] = "))
	}
	DoubleX(arr, n, x)
	console.log("N" + i + " = " + n.value + " : " + arr.join(", ") + "\n")
	arr.splice(0, arr.length)
}

function DoubleX(A, N, X) {
	let i
	for (i = 0; i < N.value; ++i) {
		if (A[i] == X) {
			A.splice(i, 0, A[i])
			++i
			N.value++
		}
	}
}
	*/
	return false
}

// { "no": 11, "dat": "2", "ans": "2" }
const param11 = () => {
	/*
let i, j, n
let arr = null
for (i = 1; i < 4; ++i) {
	n = parseInt(prompt("N" + i + " = "))
	arr = new Array(n)
	for (j = 0; j < n; ++j) {
		arr[j] = parseFloat(prompt("arr[" + j + "] = "))
	}
	SortArray(arr, n)
	printArray(arr, n)
	arr.splice(0, arr.length)
}

function SortArray(A, N) {
	let i, j, temp
	for (i = 0; i < N-1; ++i) {
		for (j = 1; j < N-i; ++j) {
			if (A[j-1] > A[j]) {
				temp = A[j-1]
				A[j-1] = A[j]
				A[j] = temp
			}
		}
	}
}

function printArray(A, N) {
	let i
	for (i = 0; i < N; ++i) {
		console.log(A[i].toFixed(2) + " ")
	}
	console.log("\n")
}
	*/
	return false
}

// { "no": 12, "dat": "2", "ans": "" }
const param12 = () => {
	/*
let i, j, n
let arr = null
let indexs = null
for (i = 1; i < 4; ++i) {
	n = parseInt(prompt("N" + i + " = "))
	arr = new Array(n)
	indexs = new Array(n)
	for (j = 0; j < n; ++j) {
		arr[j] = parseFloat(prompt("arr[" + j + "] = "))
	}
	SortIndex(arr, n, indexs)
	alert(indexs.join(", "))
	arr.splice(0, n)
	indexs.splice(0, n)
	arr = indexs = null
}

function SortIndex(A, N, I) {
	let i, j, temp
	for (i = 0; i < N; ++i) {
		I[i] = i
	}
	for (i = 0; i < N-1; ++i) {
		for (j = 1; j < N-i; ++j) {
			if (A[I[j-1]] > A[I[j]]) {
				temp = I[j-1]
				I[j-1] = I[j]
				I[j] = temp
			}
		}
	}
	for (i = 0; i < N; ++i) {
		I[i]++
	}
}
	*/
	return false
}

// { "no": 13, "dat": "2", "ans": "2" }
const param13 = () => {
	/*
let i, j, n
let arr = null
for (i = 1; i < 4; ++i) {
	n = parseInt(prompt("N" + i + " = "))
	arr = new Array(n)
	for (j = 0; j < n; ++j) {
		arr[j] = parseFloat(prompt("arr[" + j + "] = "))
	}
	Hill(arr, n)
	printArray(arr, n)
	arr.splice(0, n)
	arr = null
}

function Hill(A, N) {
	let temp, minIndex, i, j, k
	for (i = 0, j = N-1; i < j;) {
		minIndex = i
		for (k = i+1; k <= j; ++k) {
			if (A[k] < A[minIndex]) {
				minIndex = k
			}
		}
		temp = A[i]
		A[i] = A[minIndex]
		A[minIndex] = temp
		++i
		if (i == j) {
			break
		}
		
		minIndex = i
		for (k = i+1; k <= j; ++k) {
			if (A[k] < A[minIndex]) {
				minIndex = k
			}
		}
		temp = A[j]
		A[j] = A[minIndex]
		A[minIndex] = temp
		--j
	}
}

function printArray(A, N) {
	let i
	for (i = 0; i < N; ++i) {
		console.log(A[i].toFixed(2) + " ")
	}
	console.log("\n")
}
	*/
	return false
}

// { "no": 14, "dat": "2", "ans": "2" }
const param14 = () => {
	/*
function TInt() {
	this.value = 0
}

function TArray() {
	this.array = null
}

let na = parseInt(prompt("NA = "))
a = new Array(na)
let i
for (i = 0; i < na; ++i) {
	a[i] = parseFloat(prompt("A[" + i + "] = "))
}
let b = new TArray()
let c = new TArray()
let nb = new TInt()
let nc = new TInt()
Split1(a, na, b, nb, c, nc)
console.log("NB = " + nb.value + " : ")
printArray(b.array, nb.value)
console.log("\nNC = " + nc.value + " : ")
printArray(c.array, nc.value)
b.array.splice(0, nb.value)
c.array.splice(0, nc.value)

function Split1(A, NA, B, NB, C, NC) {
	NB.value = parseInt(NA/2) + (NA%2 != 0 ? 1 : 0)
	NC.value = parseInt(NA/2)
	B.array = new Array(NB.value)
	C.array = new Array(NC.value)
	let i, index
	for (i = 0, index = 0; i < NA; i += 2, ++index) {
		B.array[index] = A[i]
	}
	for (i = 1, index = 0; i < NA; i += 2, ++index) {
		C.array[index] = A[i]
	}
}

function printArray(A, N) {
	let i
	for (i = 0; i < N; ++i) {
		console.log(A[i].toFixed(2) + " ")
	}
	console.log("\n")
}
	*/
	return false
}

// { "no": 15, "dat": "", "ans": "" }
const param15 = () => {
	/*
function TInt() {
	this.value = 0
}

function TArray() {
	this.array = null
}

let na = parseInt(prompt("NA = "))
let a = new Array(na)
let i
for (i = 0; i < na; ++i) {
	a[i] = parseInt(prompt("A[" + i + "] = "))
}
let nb = new TInt()
let nc = new TInt()
let b = new TArray()
let c = new TArray()
Split2(a, na, b, nb, c, nc)
console.log("NB = " + nb.value + "; B: " + b.array.join(", ") + "\n")
console.log("NC = " + nc.value + "; C: " + c.array.join(", ") + "\n")

function Split2(A, NA, B, NB, C, NC) {
	NB.value = NC.value = 0
	let i
	for (i = 0; i < NA; ++i) {
		if (A[i] % 2 == 0) {
			NB.value++
		} else {
			NC.value++
		}
	}
	B.array = new Array(NB.value)
	C.array = new Array(NC.value)
	let iB = 0, iC = 0
	for (i = 0; i < NA; ++i) {
		if (A[i] % 2 == 0) {
			B.array[iB++] = A[i]
		} else {
			C.array[iC++] = A[i]
		}
	}
}
	*/
	return false
}

// { "no": 16, "dat": "2", "ans": "2" }
const param16 = () => {
	/*
function TMatrix() {
	this.matrix = null
}

let k = parseInt(prompt("K = "))
let a = new Array(k)
let i, j
for (i = 0; i < k; ++i) {
	a[i] = parseFloat(prompt("A[" + i + "] = "))
}
let m = parseInt(prompt("M = "))
let n = parseInt(prompt("N = "))
let b = new TMatrix()
ArrayToMatrRow(a, k, m, n, b)
for (i = 0; i < m; ++i) {
	for (j = 0; j < n; ++j) {
		console.log(b.matrix[i][j].toFixed(2) + " ")
	}
	console.log("\n")
}
for (i = 0; i < m; ++i) {
	b.matrix.splice(0, n)
}
b.matrix.splice(0, m)
b.matrix = null

function ArrayToMatrRow(A, K, M, N, B) {
	let i, j
	B.matrix = new Array(M)
	for (i = 0; i < M; ++i) {
		B.matrix[i] = new Array(N)
	}
	let index = 0
	for (i = 0; i < M; ++i) {
		for (j = 0; j < N; ++j) {
			B.matrix[i][j] = index < K ? A[index++] : 0
		}
	}
}
	*/
	return false
}

// { "no": 17, "dat": "2", "ans": "2" }
const param17 = () => {
	/*
function TMatrix() {
	this.matrix = null
}

let k = parseInt(prompt("K = "))
let a = new Array(k)
let i, j
for (i = 0; i < k; ++i) {
	a[i] = parseFloat(prompt("A[" + i + "] = "))
}
let m = parseInt(prompt("M = "))
let n = parseInt(prompt("N = "))
let b = new TMatrix()
ArrayToMatrCol(a, k, m, n, b)
for (i = 0; i < m; ++i) {
	for (j = 0; j < n; ++j) {
		console.log(b.matrix[i][j].toFixed(2) + " ")
	}
	console.log("\n")
}
for (i = 0; i < m; ++i) {
	b.matrix[i].splice(0, n)
}
b.matrix.splice(0, m)
b.matrix = null

function ArrayToMatrCol(A, K, M, N, B) {
	let i, j
	B.matrix = new Array(M)
	for (i = 0; i < M; ++i) {
		B.matrix[i] = new Array(N)
	}
	let index = 0
	for (j = 0; j < N; ++j) {
		for (i = 0; i < M; ++i) {
			B.matrix[i][j] = index < K ? A[index++] : 0
		}
	}
}
	*/
	return false
}

// { "no": 18, "dat": "", "ans": "" }
const param18 = () => {
	/*
function TMatrix() {
	this.matrix = null
}

let M = parseInt(prompt("M = "))
let N = parseInt(prompt("N = "))
let matr = new TMatrix()
Chessboard(M, N, matr)
let i, j
for (i = 0; i < M; ++i) {
	for (j = 0; j < N; ++j) {
		console.log(matr.matrix[i][j] + ", ")
	}
	console.log("\n")
}

function Chessboard(M, N, A) {
	A.matrix = new Array(M)
	let i, j
	for (i = 0; i < M; ++i) {
		A.matrix[i] = new Array(N)
		for (j = 0; j < N; ++j) {
			A.matrix[i][j] = (i+j)%2 == 0 ? 0 : 1
		}
	}
}
	*/
	return false
}

// { "no": 19, "dat": "2", "ans": "2" }
const param19 = () => {
	/*
let M = parseInt(prompt("M = "))
let N = parseInt(prompt("N = "))
let matrix = new Array(M)
let i, j, k
for (i = 0; i < M; ++i) {
	matrix[i] = new Array(N)
	for (j = 0; j < N; ++j) {
		matrix[i][j] = parseFloat(prompt("matrix[" + i + "][" + j + "] = "))
	}
}

for (k = 1; k <= M; ++k) {
	console.log(Norm1(matrix, k, N).toFixed(2) + ", ")
}

function Norm1(A, M, N) {
	let sum, maxSum = 0
	let i, j
	for (j = 0; j < N; ++j) {
		sum = 0
		for (i = 0; i < M; ++i) {
			sum += A[i][j]
		}
		if (j == 0) {
			maxSum = sum
		} else if (sum > maxSum) {
			maxSum = sum
		}
	}
	return maxSum
}
	*/
	return false
}

// { "no": 20, "dat": "2", "ans": "2" }
const param20 = () => {
	/*
let M = parseInt(prompt("M = "))
let N = parseInt(prompt("N = "))
let matrix = new Array(M)
let i, j, k
for (i = 0; i < M; ++i) {
	matrix[i] = new Array(N)
	for (j = 0; j < N; ++j) {
		matrix[i][j] = parseFloat(prompt("matrix[" + i + "][" + j + "] = "))
	}
}
for (k = 1; k <= M; ++k) {
	console.log(Norm2(matrix, k, N).toFixed(2) + "\n")
}

function Norm2(A, M, N) {
	let i, j
	let sum, maxSum = 0
	for (i = 0; i < M; ++i) {
		sum = 0
		for (j = 0; j < N; ++j) {
			sum += A[i][j]
		}
		if (i == 0) {
			maxSum = sum
		} else if (sum > maxSum) {
			maxSum = sum
		}
	}
	return maxSum
}
	*/
	return false
}

// { "no": 21, "dat": "2", "ans": "2" }
const param21 = () => {
	/*
let M = parseInt(prompt("M = "))
let N = parseInt(prompt("N = "))
let matrix = new Array(M)
let i, j
for (i = 0; i < M; ++i) {
	matrix[i] = new Array(N)
	for (j = 0; j < N; ++j) {
		matrix[i][j] = parseFloat(prompt("matrix[" + i + "][" + j + "] = "))
	}
}
let k
for (i = 1; i <= 3; ++i) {
	k = parseInt(prompt('K' + i + ':'))
	console.log(SumRow(matrix, M, N, k).toFixed(2) + "\n")
}

function SumRow(A, M, N, K) {
	let sum = 0
	if (1 <= K && K <= M) {
		let j
		for (j = 0; j < N; ++j) {
			sum += A[K-1][j]
		}
	}
	return sum
}
	*/
	return false
}

// { "no": 22, "dat": "2", "ans": "2" }
const param22 = () => {
	/*
let M = parseInt(prompt("M = "))
let N = parseInt(prompt("N = "))
let matrix = new Array(M)
let i, j
for (i = 0; i < M; ++i) {
	matrix[i] = new Array(N)
	for (j = 0; j < N; ++j) {
		matrix[i][j] = parseFloat(prompt("matrix[" + i + "][" + j + "] = "))
	}
}
let k
for (i = 1; i < 4; ++i) {
	k = parseInt(prompt('K' + i + ':'))
	console.log(SumCol(matrix, M, N, k).toFixed(2) + "\n")
}

function SumCol(A, M, N, K) {
	let sum = 0
	if (1 <= K && K <= N) {
		let i
		for (i = 0; i < M; ++i) {
			sum += A[i][K-1]
		}
	}
	return sum
}
	*/
	return false
}

// { "no": 23, "dat": "2", "ans": "2" }
const param23 = () => {
	return false
}

// { "no": 24, "dat": "2", "ans": "2" }
const param24 = () => {
	return false
}

// { "no": 25, "dat": "2", "ans": "2" }
const param25 = () => {
	return false
}

// { "no": 26, "dat": "2", "ans": "2" }
const param26 = () => {
	return false
}

// { "no": 27, "dat": "2", "ans": "2" }
const param27 = () => {
	return false
}

// { "no": 28, "dat": "2", "ans": "2" }
const param28 = () => {
	return false
}

// { "no": 29, "dat": "", "ans": "" }
const param29 = () => {
	return false
}

// { "no": 30, "dat": "", "ans": "" }
const param30 = () => {
	return false
}

// { "no": 31, "dat": "", "ans": "" }
const param31 = () => {
	return false
}

// { "no": 32, "dat": "", "ans": "" }
const param32 = () => {
	return false
}

// { "no": 33, "dat": "", "ans": "" }
const param33 = () => {
	return false
}

// { "no": 34, "dat": "", "ans": "" }
const param34 = () => {
	return false
}

// { "no": 35, "dat": "", "ans": "" }
const param35 = () => {
	return false
}

// { "no": 36, "dat": "", "ans": "" }
const param36 = () => {
	return false
}

// { "no": 37, "dat": "", "ans": "" }
const param37 = () => {
	return false
}

// { "no": 38, "dat": "", "ans": "" }
const param38 = () => {
	return false
}

// { "no": 39, "dat": "", "ans": "" }
const param39 = () => {
	return false
}

// { "no": 40, "dat": "", "ans": "" }
const param40 = () => {
	return false
}

// { "no": 41, "dat": "", "ans": "" }
const param41 = () => {
	return false
}

// { "no": 42, "dat": "", "ans": "" }
const param42 = () => {
	return false
}

// { "no": 43, "dat": "", "ans": "" }
const param43 = () => {
	return false
}

// { "no": 44, "dat": "", "ans": "" }
const param44 = () => {
	return false
}

// { "no": 45, "dat": "", "ans": "" }
const param45 = () => {
	return false
}

// { "no": 46, "dat": "", "ans": "" }
const param46 = () => {
	return false
}

// { "no": 47, "dat": "", "ans": "" }
const param47 = () => {
	return false
}

// { "no": 48, "dat": "", "ans": "" }
const param48 = () => {
	return false
}

// { "no": 49, "dat": "", "ans": "" }
const param49 = () => {
	return false
}

// { "no": 50, "dat": "", "ans": "" }
const param50 = () => {
	return false
}

// { "no": 51, "dat": "", "ans": "" }
const param51 = () => {
	return false
}

// { "no": 52, "dat": "", "ans": "" }
const param52 = () => {
	return false
}

// { "no": 53, "dat": "", "ans": "" }
const param53 = () => {
	return false
}

// { "no": 54, "dat": "", "ans": "" }
const param54 = () => {
	return false
}

// { "no": 55, "dat": "", "ans": "" }
const param55 = () => {
	return false
}

// { "no": 56, "dat": "", "ans": "" }
const param56 = () => {
	return false
}

// { "no": 57, "dat": "", "ans": "" }
const param57 = () => {
	return false
}

// { "no": 58, "dat": "", "ans": "" }
const param58 = () => {
	return false
}

// { "no": 59, "dat": "", "ans": "" }
const param59 = () => {
	return false
}

// { "no": 60, "dat": "", "ans": "" }
const param60 = () => {
	return false
}

// { "no": 61, "dat": "", "ans": "" }
const param61 = () => {
	return false
}

// { "no": 62, "dat": "", "ans": "" }
const param62 = () => {
	return false
}

// { "no": 63, "dat": "", "ans": "" }
const param63 = () => {
	return false
}

// { "no": 64, "dat": "2", "ans": "2" }
const param64 = () => {
	return false
}

// { "no": 65, "dat": "2", "ans": "2" }
const param65 = () => {
	return false
}

// { "no": 66, "dat": "2", "ans": "2" }
const param66 = () => {
	return false
}

// { "no": 67, "dat": "2", "ans": "2" }
const param67 = () => {
	return false
}

// { "no": 68, "dat": "2", "ans": "2" }
const param68 = () => {
	return false
}

// { "no": 69, "dat": "2", "ans": "2" }
const param69 = () => {
	return false
}

// { "no": 70, "dat": "2", "ans": "2" }
const param70 = () => {
	return false
}