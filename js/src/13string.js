import {
	initScanner
	, initChecker
	, inputsEOF
	, outputsEOF
	, nextInt
	, compareInt
	, nextChar
	, compareChar
	, compareWord
	, compareLine
	, nextLine
	, compareBoolean
} from './io.js'

(async () => {
	for (let taskNo = 1; taskNo <= 70; taskNo++) {
		const taskNoStr = String(taskNo).padStart(3, 0)
		for (let testNo = 1; testNo <= 100; testNo++) {
			const testNoStr = String(testNo).padStart(3, 0)
			const filePath = `../data/kt-gov2/13string/String${taskNoStr}/${testNoStr}/${testNoStr}`

			await initScanner(filePath + '.dat')
			await initChecker(filePath + '.ans')

			let ok = true
			switch (taskNo) {
				case 1: ok = string1(); break
				case 2: ok = string2(); break
				case 3: ok = string3(); break
				case 4: ok = string4(); break
				case 5: ok = string5(); break
				case 6: ok = string6(); break
				case 7: ok = string7(); break
				case 8: ok = string8(); break
				case 9: ok = string9(); break
				case 10: ok = string10(); break
				case 11: ok = string11(); break
				case 12: ok = string12(); break
				case 13: ok = string13(); break
				case 14: ok = string14(); break
				case 15: ok = string15(); break
				case 16: ok = string16(); break
				case 17: ok = string17(); break
				case 18: ok = string18(); break
				case 19: ok = string19(); break
				case 20: ok = string20(); break
				case 21: ok = string21(); break
				case 22: ok = string22(); break
				case 23: ok = string23(); break
				case 24: ok = string24(); break
				case 25: ok = string25(); break
				case 26: ok = string26(); break
				case 27: ok = string27(); break
				case 28: ok = string28(); break
				case 29: ok = string29(); break
				case 30: ok = string30(); break
				case 31: ok = string31(); break
				case 32: ok = string32(); break
				case 33: ok = string33(); break
				case 34: ok = string34(); break
				case 35: ok = string35(); break
				case 36: ok = string36(); break
				case 37: ok = string37(); break
				case 38: ok = string38(); break
				case 39: ok = string39(); break
				case 40: ok = string40(); break
				case 41: ok = string41(); break
				case 42: ok = string42(); break
				case 43: ok = string43(); break
				case 44: ok = string44(); break
				case 45: ok = string45(); break
				case 46: ok = string46(); break
				case 47: ok = string47(); break
				case 48: ok = string48(); break
				case 49: ok = string49(); break
				case 50: ok = string50(); break
				case 51: ok = string51(); break
				case 52: ok = string52(); break
				case 53: ok = string53(); break
				case 54: ok = string54(); break
				case 55: ok = string55(); break
				case 56: ok = string56(); break
				case 57: ok = string57(); break
				case 58: ok = string58(); break
				case 59: ok = string59(); break
				case 60: ok = string60(); break
				case 61: ok = string61(); break
				case 62: ok = string62(); break
				case 63: ok = string63(); break
				case 64: ok = string64(); break
				case 65: ok = string65(); break
				case 66: ok = string66(); break
				case 67: ok = string67(); break
				case 68: ok = string68(); break
				case 69: ok = string69(); break
				case 70: ok = string70(); break
			}

			if (!ok) {
				console.log("String" + taskNoStr + " #" + testNoStr + " has NOT been tested!")
				return
			}

			if (!inputsEOF()) {
				console.log("String" + taskNoStr + " #" + testNoStr + " input NOT EOF!")
				return
			}

			if (!outputsEOF()) {
				console.log("String" + taskNoStr + " #" + testNoStr + " output NOT EOF!")
				return
			}
		}
		console.log("String" + taskNoStr + " has been tested!")
	}
	console.log("String has been tested!")
})()

const string1 = () => {
	let c = nextChar()
	let code = c.charCodeAt(0)
	return compareInt(code)
}

const string2 = () => {
	let n = nextInt()
	let c = String.fromCharCode(n)
	return compareChar(c)
}

const string3 = () => {
	let c = nextChar()
	let code = c.charCodeAt(0)
	let prev = String.fromCharCode(code - 1)
	let next = String.fromCharCode(code + 1)
	return compareChar(prev, next)
}

const string4 = () => {
	let n = nextInt()
	let code = "A".charCodeAt(0)
	for (let i = 0; i < n; i++) {
		let c = String.fromCharCode(code++)
		if (!compareChar(c)) {
			return false
		}
	}
	return true
}

const string5 = () => {
	let n = nextInt()
	let code = "z".charCodeAt(0)
	for (let i = 0; i < n; i++) {
		let c = String.fromCharCode(code--)
		if (!compareChar(c)) {
			return false
		}
	}
	return true
}

const string6 = () => {
	let c = nextChar()
	let result = c.match(/\d/) ? "digit" : c.match(/[A-Z]/) ? "capital" : "small"
	return compareWord(result)
}

const string7 = () => {
	let str = nextLine()
	let firstCode = str.charCodeAt(0)
	let lastCode = str.charCodeAt(str.length-1)
	return compareInt(firstCode, lastCode)
}

const string8 = () => {
	let n = nextInt()
	let c = nextChar()
	let str = ""
	for (let i = 0; i < n; i++) {
		str += c
	}
	return compareLine(str)
}

const string9 = () => {
	let n = nextInt()
	let c1 = nextChar()
	let c2 = nextChar()
	let str = ""
	for (let i = parseInt(n/2); i > 0; i--) {
		str += c1 + c2
	}
	return compareLine(str)
}

const string10 = () => {
	let str = nextLine()
	let newStr = ""
	for (let i = str.length-1; i >= 0; i--) {
		newStr += str.charAt(i)
	}
	return compareLine(newStr)
}

const string11 = () => {
	let s = nextLine()
	let newS = ""
	let i = 0, len = s.length
	while (i < len) {
		newS += s.charAt(i)
		if (i < len-1) {
			newS += ' '
		}
		i++
	}
	return compareLine(newS)
}

const string12 = () => {
	let s = nextLine()
	let n = nextInt()
	let newS = ""
	let i = 0, len = s.length
	while (i < len) {
		newS += s.charAt(i)
		if (i < len-1) {
			for (let j = 0; j < n; j++) {
				newS += '*'
			}
		}
		i++
	}
	return compareLine(newS)
}

const string13 = () => {
	let str = nextLine()
	let digitsCount = str.match(/\d/g).length
	return compareInt(digitsCount)
}

const string14 = () => {
	let str = nextLine()
	let pattern = /[A-Z]{1}/g
	let latsCount = str.match(pattern).length
	return compareInt(latsCount)
}

const string15 = () => {
	let str = nextLine()
	let pattern = /[a-zа-я]{1}/g
	let lowersCount = str.match(pattern).length
	return compareInt(lowersCount)
}

const string16 = () => {
	let str = nextLine()
	let len = str.length
	let newStr = ""
	for (let i = 0; i < len; i++) {
		let c = str.charAt(i)
		if (c.search(/[a-z]{1}/) == 0) {
			let farq = c.charCodeAt(0) - 'a'.charCodeAt(0)
			c = String.fromCharCode('A'.charCodeAt(0) + farq)
		}
		newStr += c
	}
	return compareLine(newStr)
}

const string17 = () => {
	let str = nextLine()
	let farq, len = str.length
	let newStr = ""
	for (let i = 0; i < len; i++) {
		let c = str.charAt(i)
		if (c.search(/[а-я]{1}/) == 0) {
			farq = c.charCodeAt(0) - 'а'.charCodeAt(0)
			c = String.fromCharCode('А'.charCodeAt(0) + farq)
		} else if (c.search(/[a-z]{1}/) == 0) {
			farq = c.charCodeAt(0) - 'a'.charCodeAt(0)
			c = String.fromCharCode('A'.charCodeAt(0) + farq)
		}
		newStr += c
	}
	return compareLine(newStr)
}

const string18 = () => {
	let str = nextLine()
	let farq, len = str.length
	let newStr = ""
	for (let i = 0; i < len; i++) {
		c = str.charAt(i)
		if (c.search(/[a-z]{1}/) == 0) {
			farq = c.charCodeAt(0) - 'a'.charCodeAt(0)
			c = String.fromCharCode('A'.charCodeAt(0) + farq)
		} else if (c.search(/[A-Z]{1}/) == 0) {
			farq = c.charCodeAt(0) - 'A'.charCodeAt(0)
			c = String.fromCharCode('a'.charCodeAt(0) + farq)
		} else if (c.search(/[а-я]{1}/) == 0) {
			farq = c.charCodeAt(0) - 'а'.charCodeAt(0)
			c = String.fromCharCode('А'.charCodeAt(0) + farq)
		} else if (c.search(/[А-Я]{1}/) == 0) {
			farq = c.charCodeAt(0) - 'А'.charCodeAt(0)
			c = String.fromCharCode('а'.charCodeAt(0) + farq)
		}
		newStr += c
	}
	return compareLine(newStr)
}

const string19 = () => {
	let str = nextLine()
	let butun = /^[-]?[\d]+$/
	let haqiqi = /^[-]?[\d]+\.[\d]+$/
	let result
	if (butun.test(str)) {
		result = 1
	} else if (haqiqi.test(str)) {
		result = 2
	} else {
		result = 0
	}
	return compareInt(result)
}

const string20 = () => {
	let number = nextInt()
	let str = new String(number)
	let len = str.length
	for (let i = 0; i < len; i++) {
		if (!compareChar(str.charAt(i))) {
			return false
		}
	}
	return true
}

const string21 = () => {
	let number = nextInt()
	while (number > 0) {
		let c = String.fromCharCode('0'.charCodeAt(0) + number%10)
		if (!compareChar(c)) {
			return false
		}
		number = parseInt(number / 10)
	}
	return true
}

const string22 = () => {
	let str = nextLine()
	let number = parseInt(str)
	let sum = 0
	while (number > 0) {
		sum += number % 10
		number = parseInt(number / 10)
	}
	return compareInt(sum)
}

const string23 = () => {
	let str = nextLine()
	let res = parseInt(str)
	let len = str.length
	for (let i = 1; i < len; i += 2) {
		let digit = parseInt(str.substr(i+1, 1))
		switch (str.charAt(i)) {
			case '+': res += digit; break
			case '-': res -= digit; break
		}
	}
	return compareInt(res)
}

const string24 = () => {
	let str = nextLine()
	let dahi = 0, degree2 = 1
	for (let i = str.length-1; i >= 0; i--) {
		if (str.charAt(i) == '1') {
			dahi += degree2
		}
		degree2 *= 2
	}
	return compareInt(dahi)
}

const string25 = () => {
	let str = nextLine()
	let number = parseInt(str)
	let dui = ""
	while (number > 0) {
		let r = number % 2
		switch (r) {
			case 0: dui = '0' + dui; break
			case 1: dui = '1' + dui; break
		}
		number = parseInt(number / 2)
	}
	return compareWord(dui)
}

const string26 = () => {
	let n = nextInt()
	let s = nextLine()
	let newStr = ""
	for (let j = 0, i = s.length-1; j < n; i--, j++) {
		if (i < 0) {
			newStr = '.' + newStr
		} else {
			newStr = s.charAt(i) + newStr
		}
	}
	return compareLine(newStr)
}

const string27 = () => {
	let n1 = nextInt()
	let n2 = nextInt()
	let s1 = nextLine()
	let s2 = nextLine()
	let s = ""
	for (let i = 0; i < n1; i++) {
		s += s1.charAt(i)
	}
	for (let i = s2.length-n2, j = 0; j < n2; i++, j++) {
		s += s2.charAt(i)
	}
	return compareLine(s)
}

const string28 = () => {
	let c = nextChar()
	let s = nextLine()
	let len = s.length
	let newS = ""
	for (let i = 0; i < len; i++) {
		let symbol = s.charAt(i)
		newS += symbol
		if (symbol == c) {
			newS += symbol
		}
	}
	return compareLine(newS)
}

const string29 = () => {
	let c = nextChar()
	let s = nextLine()
	let s0 = nextLine()
	let len = s.length
	let newS = ""
	for (let i = 0; i < len; i++) {
		let symbol = s.charAt(i)
		if (symbol == c) {
			newS += s0
		}
		newS += symbol
	}
	return compareLine(newS)
}

const string30 = () => {
	let c = nextChar()
	let s = nextLine()
	let s0 = nextLine()
	let len = s.length
	let newS = ""
	for (let i = 0; i < len; i++) {
		let symbol = s.charAt(i)
		newS += symbol
		if (symbol == c) {
			newS += s0
		}
	}
	return compareLine(newS)
}

const string31 = () => {
	let s = nextLine()
	let s0 = nextLine()
	let exists = s.search(s0) >= 0
	return compareBoolean(exists)
}

const string32 = () => {
	let s = nextLine()
	let s0 = nextLine()
	let count = s.split(s0).length-1
	return compareInt(count)
}

const string33 = () => {
	let s = nextLine()
	let s0 = nextLine()
	let newS = s.replace(s0, "")
	return compareLine(newS)
}

const string34 = () => {
	let s = nextLine()
	let s0 = nextLine()
	let firstIndex = s.indexOf(s0)
	let lastIndex = s.lastIndexOf(s0)
	let newS = s
	if (lastIndex >= 0) {
		if (firstIndex == lastIndex) {
			newS = s.replace(s0, "")
		} else {
			newS = s.substr(0, lastIndex)
			newS += s.substr(lastIndex).replace(s0, "")
		}
	}
	return compareLine(newS)
}

const string35 = () => {
	let s = nextLine()
	let s0 = nextLine()
	let arr = s.split(s0)
	let newS = ""
	let len = arr.length
	for (let i = 0; i < len; i++) {
		newS += arr[i]
	}
	return compareLine(newS)
}

const string36 = () => {
	let s = nextLine()
	let s1 = nextLine()
	let s2 = nextLine()
	let newS = s.replace(s1, s2)
	return compareLine(newS)
}

const string37 = () => {
	let s = nextLine()
	let s1 = nextLine()
	let s2 = nextLine()
	let lastIndex = s.lastIndexOf(s1)
	let newS = s.substr(0, lastIndex)
	newS += s.substr(lastIndex).replace(s1, s2)
	return compareLine(newS)
}

const string38 = () => {
	let s = nextLine()
	let s1 = nextLine()
	let s2 = nextLine()
	let arr = s.split(s1)
	let len = arr.length
	let newS = ""
	for (let i = 0; i < len; i++) {
		newS += arr[i]
		if (i < len-1) {
			newS += s2
		}
	}
	return compareLine(newS)
}

const string39 = () => {
	let str = nextLine()
	let arr = str.split(' ')
	let subS = ""
	if (arr.length > 2) {
		subS = arr[1]
	}
	return compareLine(subS)
}

const string40 = () => {
	let str = nextLine()
	let firstIndex = str.indexOf(' ')
	let lastIndex = str.lastIndexOf(' ')
	let subS = ""
	if (firstIndex != lastIndex) {
		subS = str.substring(firstIndex, lastIndex)
	}
	return compareLine(subS)
}

const string41 = () => {
	let str = nextLine()
	str = str.replace(/ */, "")
	let len = str.length
	let i = len-1
	for (; i >= 0; i--) {
		if (str.charAt(i) != ' ') {
			break
		}
	}
	if (i < len-1) {
		str = str.substr(0, i+1)
	}
	let arr = str.split(/ {1,}/g)
	return compareInt(arr.length)
}

const string42 = () => {
	let str = nextLine()
	str = str.replace(/ */, "")
	let len = str.length
	let i = len-1
	for (; i >= 0; i--) {
		if (str.charAt(i) != ' ') {
			break
		}
	}
	if (i < len-1) {
		str = str.substr(0, i+1)
	}
	let arr = str.split(/ +/g)
	let count = 0
	for (let i = 0; i < arr.length; i++) {
		if (arr[i].charAt(0) == arr[i].charAt(arr[i].length-1)) {
			count++
		}
	}
	return compareInt(count)
}

const string43 = () => {
	let str = nextLine()
	str = str.replace(/ */, "")
	let len = str.length
	let i = len-1
	for (; i >= 0; i--) {
		if (str.charAt(i) != ' ') {
			break
		}
	}
	if (i < len-1) {
		str = str.substr(0, i+1)
	}
	let arr = str.split(/ +/g)
	let count = 0
	for (let i = 0; i < arr.length; i++) {
		if (arr[i].search('А') >= 0) {
			count++
		}
	}
	return compareInt(count)
}

const string44 = () => {
	let str = nextLine()
	// delete spaces at the start and the end
	str = str.replace(/ */, "")
	let len = str.length
	let i = len-1
	for (; i >= 0; i--) {
		if (str.charAt(i) != ' ') {
			break
		}
	}
	if (i < len-1) {
		str = str.substr(0, i+1)
	}
	let arr = str.split(/ {1,}/g)
	let count = 0
	for (let i = 0; i < arr.length; i++) {
		let mass = arr[i].match(/[А]{1}/g)
		if (mass != null && mass.length == 3) {
			count++
		}
	}
	return compareInt(count)
}

const string45 = () => {
	let str = nextLine()
	// delete spaces at the start and the end
	str = str.replace(/ */, "")
	let len = str.length
	let i = len-1
	for (; i >= 0; i--) {
		if (str.charAt(i) != ' ') {
			break
		}
	}
	if (i < len-1) {
		str = str.substr(0, i+1)
	}
	let arr = str.split(/ +/g)
	let minLen = arr[0].length
	for (let i = 1; i < arr.length; i++) {
		len = arr[i].length
		if (len < minLen) {
			minLen = len
		}
	}
	return compareInt(minLen)
}

const string46 = () => {
	let str = nextLine()
	// delete spaces at the start and the end
	str = str.replace(/ */, ' ')
	let len = str.length
	let i = len-1
	for (; i >= 0; i--) {
		if (str.charAt(i) != ' ') {
			break
		}
	}
	if (i < len-1) {
		str = str.substr(0, i+1)
	}
	let arr = str.split(/ +/g)
	let maxLen = arr[0].length
	for (let i = 1; i < arr.length; i++) {
		len = arr[i].length
		if (len > maxLen) {
			maxLen = len
		}
	}
	return compareInt(maxLen)
}

const string47 = () => {
	let str = nextLine()
	// delete spaces at the start and the end
	str = str.replace(/ */, "")
	let len = str.length
	let i = len-1
	for (; i >= 0;) {
		if (str.charAt(i) != ' ') {
			break
		}
	}
	if (i < len-1) {
		str = str.substr(0, i+1)
	}
	let arr = str.split(/ +/g)
	let newStr = ""
	len = arr.length
	for (let i = 0; i < len; i++) {
		newStr += arr[i]
		if (i < len-1) {
			newStr += '.'
		}
	}
	return compareLine(newStr)
}

const string48 = () => {
	let str = nextLine()
	let word, newStr = ""
	let start
	while ((start = str.search(/\S/)) >= 0) {
		newStr += str.substr(0, start)
		str = str.substr(start)
		start = 0
		let end = str.search(/\s/)
		if (end < 0) {
			end = str.length
		}
		word = str.substring(start-1, end)
		c = word.charAt(0)
		word = c + word.substr(1).replace(c, '.')
		newStr += word
		str = str.substr(end)
	}
	return compareLine(newStr)
}

const string49 = () => {
	let str = nextLine()
	let newStr = ""
	let start
	while ((start = str.search(/\S/)) >= 0) {
		newStr += str.substr(0, start)
		str = str.substr(start)
		start = 0
		let end = str.search(/\s/)
		if (end < 0) {
			end = str.length
		}
		let word = str.substring(start-1, end)
		let c = word.charAt(word.length-1)
		word = word.substr(0, word.length-1).replace(c, '.') + c
		newStr += word
		str = str.substr(end)
	}
	return compareLine(newStr)
}

const string50 = () => {
	let str = nextLine()
	// delete spaces at the start and the end
	let i = str.search(/\S/)
	if (i > 0) {
		str = str.substr(i)
	}
	i = len-1
	let len = str.length
	for (; i >= 0; i--) {
		if (str.charAt(i) != ' ') {
			break
		}
	}
	if (i < len-1) {
		str = str.substr(0, i+1)
	}
	let arr = str.split(/ {1,}/)
	let newStr = ""
	for (let i = arr.length-1; i >= 0; i--) {
		newStr += arr[i]
		if (i > 0) {
			newStr += ' '
		}
	}
	return compareLine(newStr)
}

const string51 = () => {
	let str = nextLine()
	// delete spaces at the start and the end
	let i = str.search(/\S/)
	if (i > 0) {
		str = str.substr(i)
	}
	let len = str.length
	for (i = len-1; i >= 0; i--) {
		if (str.charAt(i) != ' ') {
			break
		}
	}
	if (i < len-1) {
		str = str.substr(0, i+1)
	}
	let arr = str.split(/ {1,}/)
	sortTextArr(arr)
	let newStr = ""
	for (let i = arr.length-1; i >= 0; i--) {
		newStr += arr[i]
		if (i > 0) {
			newStr += ' '
		}
	}
	return compareLine(newStr)
}

const sortTextArr = (arr) => {
	let len = arr.length
	for (let i = 0; i < len-1; i++) {
		for (let j = 1; j < len; j++) {
			if (arr[j-1] < arr[j]) {
				let temp = arr[j-1]
				arr[j-1] = arr[j]
				arr[j] = temp
			}
		}
	}
}

const string52 = () => {
	let str = nextLine()
	let start
	let newStr = ""
	while ((start = str.search(/\S/)) >= 0) {
		newStr += str.substr(0, start)
		str = str.substr(start)
		start = 0
		let end = str.search(/\s/)
		if (end < 0) {
			end = str.length
		}
		let word = str.substr(0, end)
		let c = word.charAt(0)
		if (/[а-я]{1}/.test(c)) {
			let farq = c.charCodeAt(0) - 'а'.charCodeAt(0)
			let code = 'А'.charCodeAt(0) + farq
			c = String.fromCharCode(code)
			word = c + word.substr(1)
		}
		newStr += word
		str = str.substr(end)
	}
	return compareLine(newStr)
}

const string53 = () => {
	let str = nextLine()
	let alomatho = [/,{1}/g, /\.{1}/g, /:{1}/g, /;{1}/g, /\?{1}/g, /!{1}/g, /-{1}/g]
	let count = 0
	let len = alomatho.length
	for (let i = 0; i < len; i++) {
		let arr = str.match(alomatho[i])
		if (arr != null) {
			count += arr.length
		}
	}
	return compareInt(count)
}

const string54 = () => {
	let str = nextLine()
	let sadonokho = [/а/g, /о/g, /ы/g, /и/g, /у/g, /э/g, /е/g, /ю/g, /я/g, /ё/g]
	let len = sadonokho.length
	let count = 0
	for (let i = 0; i < len; i++) {
		let arr = str.match(sadonokho[i])
		if (arr != null) {
			count += arr.length
		}
	}
	return compareInt(count)
}

const string55 = () => {
	let str = nextLine()
	let i = str.search(/\S/)
	if (i > 0) {
		str = str.substr(i)
	}
	let len = str.length
	i = len-1
	for (; i >= 0; i--) {
		if (str.charAt(i) != ' ') {
			break
		}
	}
	if (i < len-1) {
		str = str.substr(0, i+1)
	}
	let arr = str.split(/ +/g)
	len = arr.length
	let maxDarozi = arr[0].length
	let maxIndex
	for (let i = 1; i < len; i++) {
		let darozi = arr[i].length
		if (/[^а-я]/.test(arr[i].charAt(arr[i].length-1))) {
			darozi--
		}
		if (darozi > maxDarozi) {
			maxIndex = i
			maxDarozi = darozi
		}
	}
	return compareWord(arr[maxIndex])
}

const string56 = () => {
	let str = nextLine()
	let i = str.search(/\S/)
	if (i > 0) {
		str = str.substr(i)
	}
	let len = str.length
	for (let i = len-1; i >= 0;) {
		if (str.charAt(i) == ' ') {
			i--
		} else {
			break
		}
	}
	if (i < len-1) {
		str = str.substr(0, i+1)
	}
	let arr = str.split(/ +/g)
	len = arr.length
	let minDarozi = arr[0].length
	let minIndex = 0
	for (let i = 1; i < len; i++) {
		let darozi = arr[i].length
		if (/[^а-я]/.test(arr[i].charAt(arr[i].length-1))) {
			darozi--
		}
		if (darozi <= minDarozi) {
			minIndex = i
			minDarozi = darozi
		}
	}
	return compareWord(arr[minIndex])
}

const string57 = () => {
	let str = nextLine()
	let i = str.search(/\S/)
	if (i > 0) {
		str = str.substr(i)
	}
	let len = str.length
	i = len-1
	for (; i >= 0; i--) {
		if (str.charAt(i) != ' ') {
			break
		}
	}
	if (i < len-1) {
		str = str.substr(0, i+1)
	}
	let arr = str.split(/ +/g)
	let newStr = ""
	len = arr.length
	for (let i = 0; i < len; i++) {
		newStr += arr[i]
		if (i < len-1) {
			newStr += ' '
		}
	}
	return compareWord(newStr)
}

const string58 = () => {
	let str = nextLine()
	let start = str.lastIndexOf('\\') + 1
	let end = str.lastIndexOf('.')
	if (end < 0) {
		end = str.length
	} else {
		end--
	}
	let fileName = str.substr(start, end-start+1)
	return compareLine(fileName)
}

const string59 = () => {
	let str = nextLine()
	let point = str.lastIndexOf('.')
	let newStr = ""
	if (point >= 0) {
		newStr = str.substr(point+1)
	}
	return compareLine(newStr)
}

const string60 = () => {
	let str = nextLine()
	let start = str.indexOf('\\')+1
	let end = str.substr(start).indexOf('\\') + start
	let newStr = "\\"
	if (end > start) {
		end--
		newStr = str.substr(start, end-start+1)
	}
	return compareLine(newStr)
}

const string61 = () => {
	let str = nextLine()
	let end = str.lastIndexOf('\\')-1
	let start = str.substr(0, end).lastIndexOf('\\')
	let newStr = ""
	if (start > 0) {
		start++
		newStr = str.substr(start, end-start+1)
	}
	return compareLine(newStr)
}

const string62 = () => {
	let str = nextLine()
	let len = str.length
	let start = 'а'.charCodeAt(0)
	let end = 'я'.charCodeAt(0)
	let START = 'А'.charCodeAt(0)
	let END = 'Я'.charCodeAt(0)
	for (let i = 0; i < len; i++) {
		let tempStr = str.substr(0, i)
		let c = str.charAt(i)
		let code = c.charCodeAt(0)
		if (/[А-Я]/.test(c)) {
			code = code+1 > END ? START : code+1
		} else if (/[а-я]/.test(c)) {
			code = code+1 > end ? start : code+1
		}
		c = String.fromCharCode(code)
		tempStr += c
		tempStr += str.substr(i+1)
		str = tempStr
	}
	return compareLine(str)
}

const string63 = () => {
	let str = nextLine()
	let k = nextInt()
	let len = str.length
	let start = 'а'.charCodeAt(0)
	let end = 'я'.charCodeAt(0)
	let START = 'А'.charCodeAt(0)
	let END = 'Я'.charCodeAt(0)
	for (let i = 0; i < len; i++) {
		let tempStr = str.substr(0, i)
		let c = str.charAt(i)
		let code = c.charCodeAt(0)
		if (/[А-Я]/.test(c)) {
			code = code+k <= END ? code+k : START + (k - (END - code + 1))
		} else if (/[а-я]/.test(c)) {
			code = code+k <= end ? code+k : start + (k - (end - code + 1))
		}
		c = String.fromCharCode(code)
		tempStr += c
		tempStr += str.substr(i+1)
		str = tempStr
	}
	return compareLine(str)
}

const string64 = () => {
	let str = nextLine()
	let k = nextInt()
	let len = str.length
	let start = 'а'.charCodeAt(0)
	let end = 'я'.charCodeAt(0)
	let START = 'А'.charCodeAt(0)
	let END = 'Я'.charCodeAt(0)
	for (let i = 0; i < len; i++) {
		let tempStr = str.substr(0, i)
		let c = str.charAt(i)
		let code = c.charCodeAt(0)
		if (/[А-Я]/.test(c)) {
			code = code-k >= START ? code-k : END - (k - (code - START + 1))
		} else if (/[а-я]/.test(c)) {
			code = code-k >= start ? code-k : end - (k - (code - start + 1))
		}
		c = String.fromCharCode(code)
		tempStr += c
		tempStr += str.substr(i+1)
		str = tempStr
	}
	return compareLine(str)
}

const string65 = () => {
	let str = nextLine()
	let c = nextChar()
	let start = 'а'.charCodeAt(0)
	let end = 'я'.charCodeAt(0)
	let START = 'А'.charCodeAt(0)
	let END = 'Я'.charCodeAt(0)
	let k = c.charCodeAt(0) - str.charAt(0).charCodeAt(0)
	if (k > 0) {
		k = END - START + 1 - k
	} else {
		k *= -1
	}
	if (!compareInt(k)) {
		return false
	}
	let len = str.length
	for (let i = 0; i < len; i++) {
		let tempStr = str.substr(0, i)
		let c = str.charAt(i)
		let code = c.charCodeAt(0)
		if (/[А-Я]/.test(c)) {
			code = code-k >= START ? code-k : END - (k - (code - START + 1))
		} else if (/[а-я]/.test(c)) {
			code = code-k >= start ? code-k : end - (k - (code - start + 1))
		}
		c = String.fromCharCode(code)
		tempStr += c
		tempStr += str.substr(i+1)
		str = tempStr
	}
	return compareLine(str)
}

const string66 = () => {
	let str = nextLine()
	let newStr = ""
	let len = str.length
	for (let i = 1; i < len; i += 2) {
		newStr += str.charAt(i)
	}
	for (let i = len-1-(len%2==0 ? 1 : 0); i >= 0; i -= 2) {
		newStr += str.charAt(i)
	}
	return compareLine(newStr)
}

const string67 = () => {
	let str = nextLine()
	let newStr = ""
	for (let start = 0, end = str.length-1; start <= end;) {
		newStr += str.charAt(end--)
		if (start > end) {
			break
		}
		newStr += str.charAt(start++)
	}
	return compareLine(newStr)
}

const string68 = () => {
	let str = nextLine()
	let len = str.length
	let firstTime = true
	let prev
	let errorNo = -1
	for (let i = 0; i < len; i++) {
		let c = str.charAt(i)
		if ((errorNo < 0) && (/[a-z]/.test(c))) {
			if (firstTime) {
				firstTime = false
			} else {
				if (c < prev) {
					if (errorNo < 0) {
						errorNo = i
					}
				}
			}
			prev = c
		}
	}
	errorNo++
	return compareInt(errorNo)
}

const string69 = () => {
	let str = nextLine()
	let arr = []
	let i, len = str.length
	let errorNo = 0
	for (i = 0; i < len; i++) {
		let c = str.charAt(i)
		if (c == '(') {
			arr.push(c)
		} else if (c == ')') {
			if (arr.length > 0) {
				if (arr[arr.length-1] == '(') {
					arr.pop()
				} else {
					errorNo = i+1
					break
				}
			} else {
				errorNo = i+1
				break
			}
		}
	}
	if (errorNo == 0 && i == len && arr.length > 0) {
		errorNo = -1
	}
	return compareInt(errorNo)
}

const string70 = () => {
	let str = nextLine()
	let arr = []
	let i, len = str.length
	let openC
	let errorNo = 0
	for (i = 0; i < len; i++) {
		let c = str.charAt(i)
		if (c == '(' || c == '[' || c == '{') {
			arr.push(c)
		} else if (c == ')' || c == ']' || c == '}') {
			if (c == ')') {
				openc = '('
			} else if (c == ']') {
				openc = '['
			} else if (c == '}') {
				openc = '{'
			}
			if (arr.length > 0) {
				if (arr[arr.length-1] == openC) {
					arr.pop()
				} else {
					errorNo = i+1
					break
				}
			} else {
				errorNo = i+1
				break
			}
		}
	}
	if (errorNo == 0 && i == len && arr.length > 0) {
		errorNo = -1
	}
	return compareInt(errorNo)
}