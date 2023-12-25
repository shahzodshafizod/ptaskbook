import {
	initScanner
	, initChecker
	, inputsEOF
	, outputsEOF
} from './io.js'

(async () => {
	for (let taskNo = 1; taskNo <= 30; taskNo++) {
		const taskNoStr = String(taskNo).padStart(3, 0)
		for (let testNo = 1; testNo <= 100; testNo++) {
			const testNoStr = String(testNo).padStart(3, 0)
			const filePath = `../data/kt-gov2/17recur/Recur${taskNoStr}/${testNoStr}/${testNoStr}`

			await initScanner(filePath + '.dat')
			await initChecker(filePath + '.ans')

			let ok = true
			switch (taskNo) {
				case 1: ok = recur1(); break
                case 2: ok = recur2(); break
                case 3: ok = recur3(); break
                case 4: ok = recur4(); break
                case 5: ok = recur5(); break
                case 6: ok = recur6(); break
                case 7: ok = recur7(); break
                case 8: ok = recur8(); break
                case 9: ok = recur9(); break
                case 10: ok = recur10(); break
                case 11: ok = recur11(); break
                case 12: ok = recur12(); break
                case 13: ok = recur13(); break
                case 14: ok = recur14(); break
                case 15: ok = recur15(); break
                case 16: ok = recur16(); break
                case 17: ok = recur17(); break
                case 18: ok = recur18(); break
                case 19: ok = recur19(); break
                case 20: ok = recur20(); break
                case 21: ok = recur21(); break
                case 22: ok = recur22(); break
                case 23: ok = recur23(); break
                case 24: ok = recur24(); break
                case 25: ok = recur25(); break
                case 26: ok = recur26(); break
                case 27: ok = recur27(); break
                case 28: ok = recur28(); break
                case 29: ok = recur29(); break
                case 30: ok = recur30(); break
			}

			if (!ok) {
				console.log("Recur" + taskNoStr + " #" + testNoStr + " has NOT been tested!")
				return
			}

			if (!inputsEOF()) {
				console.log("Recur" + taskNoStr + " #" + testNoStr + " input NOT EOF!")
				return
			}

			if (!outputsEOF()) {
				console.log("Recur" + taskNoStr + " #" + testNoStr + " output NOT EOF!")
				return
			}
		}
		console.log("Recur" + taskNoStr + " has been tested!")
	}
	console.log("Recur has been tested!")
})()

const recur1 = () => {
/*
var i, n;
for (i = 1; i <= 5; ++i) {
	n = parseInt(prompt("n" + i + " = "));
	alert(Fact(n));
}

function Fact(N) {
	if (N <= 1) {
		return 1;
	}
	return N * Fact(N - 1);
}
*/
    return false
}

const recur2 = () => {
/*
var i, n;
for (i = 1; i < 6; ++i) {
	n = parseInt(prompt("n" + i + " = "));
	alert(Fact2(n));
}

function Fact2(N) {
	if (N <= 1) {
		return 1;
	}
	return N * Fact2(N - 2);
}
*/
    return false
}

const recur3 = () => {
/*
var x = parseFloat(prompt("X = "));
var i, n, result;
for (i = 1; i <= 5; ++i) {
	n = parseInt(prompt("N" + i + " = "));
	result = PowerN(x, n);
	alert(result);
}

function PowerN(X, N) {
	if (N == 0) {
		return 1;
	}
	if (N < 0) {
		return 1 / PowerN(X, -N);
	}
	if (N % 2 == 0) {
		var half = PowerN(X, N/2);
		return half * half;
	}
	return X * PowerN(X, N-1);
}
*/
    return false
}

const recur4 = () => {
/*
var count;

var i, n, result;
for (i = 1; i < 6; ++i) {
	n = parseInt(prompt("N" + i + " = "));
	count = 0;
	result = Fib1(n);
	alert("Fib(" + n + ") = " + result + "\ncall count: " + count);
}

function Fib1(N) {
	++count;
	if (N <= 2) {
		return 1;
	}
	return Fib1(N-2) + Fib1(N-1);
}
*/
    return false
}

const recur5 = () => {
/*
var arr = new Array(20);
var count;

var i, j, n, result;
for (i = 1; i <= 5; ++i) {
	n = parseInt(prompt("N" + i + " = "));
	for (j = 0; j < n; ++j) {
		arr[j] = 0;
	}
	count = 0;
	result = Fib2(n);
	alert("Fib2(" + n + ") = " + result + "\ncall count: " + count);
}

function Fib2(N) {
	++count;
	if (N <= 2) {
		arr[N-1] = 1;
		return arr[N-1];
	}
	if (arr[N-1] == 0) {
		arr[N-1] = Fib2(N - 2) + Fib2(N - 1);
	}
	return arr[N-1];
}
*/
    return false
}

const recur6 = () => {
/*
var count;

var n = parseInt(prompt("N = "));
var i, k, comb;
for (i = 1; i < 6; ++i) {
	k = parseInt(prompt("K" + i + " = "));
	count = 0;
	comb = Combin1(n, k);
	alert("Combin1(" + n + ", " + k + ") = " + comb + "\ncall count: " + count);
}

function Combin1(N, K) {
	++count;
	if (K == 0 || K == N) {
		return 1;
	}
	return Combin1(N-1, K) + Combin1(N-1, K-1);
}
*/
    return false
}

const recur7 = () => {
/*
var count;
var matr = new Array(20);
for (i = 0; i < 20; ++i) {
	matr[i] = new Array(20);
}

var n = parseInt(prompt("N = "));
var i, a, b, k, comb;
for (i = 1; i < 6; ++i) {
	k = parseInt(prompt("K" + i + " = "));
	for (a = 0; a < n; ++a) {
		for (b = 0; b < k; ++b) {
			matr[a][b] = 0;
		}
	}
	count = 0;
	comb = Combin2(n, k);
	alert("Combin2(" + n + ", " + k + ") = " + comb + "\ncall count: " + count);
}

function Combin2(N, K) {
	++count;
	if (K == 0 || K == N) {
		matr[N-1][K-1] = 1;
		return matr[N-1][K-1];
	}
	if (matr[N-1][K-1] == 0) {
		matr[N-1][K-1] = Combin2(N-1, K) + Combin2(N-1, K-1);
	}
	return matr[N-1][K-1];
}
*/
    return false
}

const recur8 = () => {
/*
var x = parseFloat(prompt("X = "));
var k = parseInt(prompt("K = "));
var i, n, result;
for (i = 1; i < 7; ++i) {
	n = parseInt(prompt("N" + i + " = "));
	result = RootK(x, k, n);
	alert(result.toFixed(8));
}

function PowerN(X, N) {
	if (N == 0) {
		return 1;
	}
	if (N < 0) {
		return 1 / PowerN(X, -N);
	}
	if (N % 2 == 0) {
		var half = PowerN(X, N/2);
		return half * half;
	}
	return X * PowerN(X, N-1);
}

function RootK(X, K, N) {
	if (N == 0) {
		return 1;
	}
	var Yn = RootK(X, K, N-1);
	return  Yn - (Yn - X / PowerN(Yn, K-1)) / K;
}
*/
    return false
}

const recur9 = () => {
/*
var a = parseInt(prompt("A = "));
var i, b;
for (i = 1; i < 4; ++i) {
	b = parseInt(prompt("B" + i + " = "));
	alert(GCD(a, b));
}

function GCD(A, B) {
	if (B == 0) {
		return A;
	}
	return GCD(B, A%B);
}
*/
    return false
}

const recur10 = () => {
/*
var i, k;
for (i = 1; i < 6; ++i) {
	k = parseInt(prompt("K" + i + " = "));
	alert(DigitSum(k));
}

function DigitSum(K) {
	if (K == 0) {
		return 0;
	}
	if (K < 0) {
		K = -K;
	}
	return K%10 + DigitSum(parseInt(K/10));
}
*/
    return false
}

const recur11 = () => {
/*
var N;
var arr = null;
var i, j;
for (i = 1; i < 4; i++)  {
	N = parseInt(prompt("N" + i + " = "));
	arr = new Array(N);
	for (j = 0; j < N; j++)
		arr[j] = parseInt(prompt("arr[" + j + "] = "));
	alert(MaxElem(arr, N));
	arr = null;
}

function MaxElem(A, N) {
	if (N == 0)
		return A[N];
	var elem = MaxElem(A, N-1);
	return A[N-1] > elem ? A[N-1] : elem;
}
*/
    return false
}

const recur12 = () => {
/*
var str;
var i;
for (i = 1; i <= 5; i++) {
	str = prompt("String" + i + ":");
	alert(DigitCount(str, 0));
}

function DigitCount(S, index) {
	if (index == S.length) {
		return 0;
	}
	var count = S.charAt(index) >= '0' && S.charAt(index) <= '9' ? 1 : 0;
	return count + DigitCount(S, index+1);
}
*/
    return false
}

const recur13 = () => {
/*
var str;
var i;
for (i = 1; i < 6; i++) {
	str = prompt("String" + i + ":");
	alert(Palindrome(str, 0, str.length-1));
}

function Palindrome(S, start, end) {
	if (start >= end)
		return true;
	var areEqual = S.charAt(start) == S.charAt(end);
	return (areEqual ? Palindrome(S, start+1, end-1) : false );
}
*/
    return false
}

const recur14 = () => {
/*
var str = prompt("String:");
alert(Rec(str, str.length-1));

function Rec(S, index) {
	var digit = parseInt(S.charAt(index));
	if (index <= 0)
		return digit;
	if (S.charAt(index-1) == '+')
		return Rec(S, index-2) + digit;
	return Rec(S, index-2) - digit;
}
*/
    return false
}

const recur15 = () => {
/*
function TInt() {
	this.value = 0;
};

var str = prompt("String:");
var index = new TInt();
index.value = str.length-1;
alert(Rec(str, index));

function Rec(S, index) {
	var number = (index.value != 0) && (S.charAt(index.value-1) != '*') ? parseInt(S.charAt(index.value)) : Zarb(S, index);
	if (index.value == 0)
		return number;
	index.value -= 2;
	if (S.charAt(index.value+1) == '+')
		return Rec(S, index) + number;
	return Rec(S, index) - number;
}

function Zarb(S, index) {
	var number = parseInt(S.charAt(index.value));
	if ( (index.value == 0) || (S.charAt(index.value-1) != '*') )
		return number;
	index.value -= 2;
	return Zarb(S, index) * number;
}
*/
    return false
}

const recur16 = () => {
/*
function TInt() {
	this.value = 0;
};

var str = prompt("String:");
var index = new TInt();
index.value = str.length - 1;
alert(Rec(str, index));

function Rec(S, index) {
	var number;
	if (S.charAt(index.value) == ')') {
		index.value--;
		number = Rec(S, index);
	} else {
		number = parseInt(S.charAt(index.value));
	}
	if (S.charAt(index.value-1) == '*') {
		index.value -= 2;
		number *= Zarb(S, index);
	}
	if (index.value == 0)
		return number;
	if (S.charAt(index.value-1) == '(') {
		index.value--;
		return number;
	}
	index.value -= 2;
	if (S.charAt(index.value+1) == '+')
		return Rec(S, index) + number;
	return Rec(S, index) - number;
}

function Zarb(S, index) {
	var number;
	if (S.charAt(index.value) == ')') {
		index.value--;
		number = Rec(S, index);
	} else {
		number = parseInt(S.charAt(index.value));
	}
	if ( (index.value == 0) || (S.charAt(index.value-1) != '*') )
		return number;
	index.value -= 2;
	return Zarb(S, index) * number;
}
*/
    return false
}

const recur17 = () => {
/*
function TInt() {
	this.value = 0;
};

var str = prompt("String:");
var index = new TInt();
index.value = str.length - 1;
alert(Rec(str, index));

function Rec(S, index) {
	var number;
	if (S.charAt(index.value) == ')') {
		index.value--;
		number = Rec(S, index);
	} else {
		number = parseInt(S.charAt(index.value));
	}
	if (S.charAt(index.value-1) == '*') {
		index.value -= 2;
		number *= Zarb(S, index);
	}
	if (index.value == 0)
		return number;
	if (S.charAt(index.value-1) == '(') {
		index.value--;
		return number;
	}
	index.value -= 2;
	if (S.charAt(index.value+1) == '+')
		return Rec(S, index) + number;
	return Rec(S, index) - number;
}

function Zarb(S, index) {
	var number;
	if (S.charAt(index.value) == ')') {
		index.value--;
		number = Rec(S, index);
	} else {
		number = parseInt(S.charAt(index.value));
	}
	if ( (index.value == 0) || (S.charAt(index.value-1) != '*') )
		return number;
	index.value -= 2;
	return Zarb(S, index) * number;
}
*/
    return false
}

const recur18 = () => {
    return false
}

const recur19 = () => {
    return false
}

const recur20 = () => {
    return false
}

const recur21 = () => {
    return false
}

const recur22 = () => {
    return false
}

const recur23 = () => {
    return false
}

const recur24 = () => {
    return false
}

const recur25 = () => {
	return false
}

const recur26 = () => {
	return false
}

const recur27 = () => {
	return false
}

const recur28 = () => {
	return false
}

const recur29 = () => {
	return false
}

const recur30 = () => {
	return false
}