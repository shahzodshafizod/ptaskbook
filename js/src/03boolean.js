import {
    nextFloat
    , nextInt
    , compareBoolean
    , initScanner
    , initChecker
    , inputsEOF
    , outputsEOF
} from './io.js'

(async () => {
    for (let taskNo = 1; taskNo <= 40; taskNo++) {
        const taskNoStr = String(taskNo).padStart(3, 0)
        for (let testNo = 1; testNo <= 100; testNo++) {
            const testNoStr = String(testNo).padStart(3, 0)
            const filePath = `../data/kt-gov2/03boolean/Boolean${taskNoStr}/${testNoStr}/${testNoStr}`

            await initScanner(filePath + '.dat')
            await initChecker(filePath + '.ans')

            let ok = true
            switch (taskNo) {
                case 1: ok = boolean1(); break
                case 2: ok = boolean2(); break
                case 3: ok = boolean3(); break
                case 4: ok = boolean4(); break
                case 5: ok = boolean5(); break
                case 6: ok = boolean6(); break
                case 7: ok = boolean7(); break
                case 8: ok = boolean8(); break
                case 9: ok = boolean9(); break
                case 10: ok = boolean10(); break
                case 11: ok = boolean11(); break
                case 12: ok = boolean12(); break
                case 13: ok = boolean13(); break
                case 14: ok = boolean14(); break
                case 15: ok = boolean15(); break
                case 16: ok = boolean16(); break
                case 17: ok = boolean17(); break
                case 18: ok = boolean18(); break
                case 19: ok = boolean19(); break
                case 20: ok = boolean20(); break
                case 21: ok = boolean21(); break
                case 22: ok = boolean22(); break
                case 23: ok = boolean23(); break
                case 24: ok = boolean24(); break
                case 25: ok = boolean25(); break
                case 26: ok = boolean26(); break
                case 27: ok = boolean27(); break
                case 28: ok = boolean28(); break
                case 29: ok = boolean29(); break
                case 30: ok = boolean30(); break
                case 31: ok = boolean31(); break
                case 32: ok = boolean32(); break
                case 33: ok = boolean33(); break
                case 34: ok = boolean34(); break
                case 35: ok = boolean35(); break
                case 36: ok = boolean36(); break
                case 37: ok = boolean37(); break
                case 38: ok = boolean38(); break
                case 39: ok = boolean39(); break
                case 40: ok = boolean40(); break
            }

            if (!ok) {
                console.log("Boolean" + taskNoStr + " #" + testNoStr + " has NOT been tested!")
                return
            }

            if (!inputsEOF()) {
                console.log("Boolean" + taskNoStr + " #" + testNoStr + " input NOT EOF!")
                return
            }

            if (!outputsEOF()) {
                console.log("Boolean" + taskNoStr + " #" + testNoStr + " output NOT EOF!")
                return
            }
        }
        console.log("Boolean" + taskNoStr + " has been tested!")
    }
    console.log("Boolean has been tested!")
})()

// { "no": 1, "dat": "", "ans": "" }
const boolean1 = () => {
    let a = nextInt()
    let isPositive = a > 0
    return compareBoolean(isPositive)
}

// { "no": 2, "dat": "", "ans": "" }
const boolean2 = () => {
    let a = nextInt()
    let isOdd = a % 2 != 0
    return compareBoolean(isOdd)
}

// { "no": 3, "dat": "", "ans": "" }
const boolean3 = () => {
    let a = nextInt()
    let isEven = a % 2 == 0
    return compareBoolean(isEven)
}

// { "no": 4, "dat": "", "ans": "" }
const boolean4 = () => {
    let a = nextInt()
    let b = nextInt()
    let result = a > 2 && b <= 3
    return compareBoolean(result)
}

// { "no": 5, "dat": "", "ans": "" }
const boolean5 = () => {
    let a = nextInt()
    let b = nextInt()
    let result = a >= 0 || b < -2
    return compareBoolean(result)
}

// { "no": 6, "dat": "", "ans": "" }
const boolean6 = () => {
    let a = nextInt()
    let b = nextInt()
    let c = nextInt()
    let result = a < b && b < c
    return compareBoolean(result)
}

// { "no": 7, "dat": "", "ans": "" }
const boolean7 = () => {
    let a = nextInt()
    let b = nextInt()
    let c = nextInt()
    let result = a < b && b < c || a > b && b > c
    return compareBoolean(result)
}

// { "no": 8, "dat": "", "ans": "" }
const boolean8 = () => {
    let a = nextInt()
    let b = nextInt()
    let eachOdd = (a % 2 != 0) && (b % 2 != 0)
    return compareBoolean(eachOdd)
}

// { "no": 9, "dat": "", "ans": "" }
const boolean9 = () => {
    let a = nextInt()
    let b = nextInt()
    let anyOdd = (a % 2 != 0) || (b % 2 != 0)
    return compareBoolean(anyOdd)
}

// { "no": 10, "dat": "", "ans": "" }
const boolean10 = () => {
    let a = nextInt()
    let b = nextInt()
    let oneOdd = (a + b) % 2 != 0
    return compareBoolean(oneOdd)
}

// { "no": 11, "dat": "", "ans": "" }
const boolean11 = () => {
    let a = nextInt()
    let b = nextInt()
    let equalParity = (a + b) % 2 == 0
    return compareBoolean(equalParity)
}

// { "no": 12, "dat": "", "ans": "" }
const boolean12 = () => {
    let a = nextInt()
    let b = nextInt()
    let c = nextInt()
    let eachPositive = a > 0 && b > 0 && c > 0
    return compareBoolean(eachPositive)
}

// { "no": 13, "dat": "", "ans": "" }
const boolean13 = () => {
    let a = nextInt()
    let b = nextInt()
    let c = nextInt()
    let anyPositive = a > 0 || b > 0 || c > 0
    return compareBoolean(anyPositive)
}

// { "no": 14, "dat": "", "ans": "" }
const boolean14 = () => {
    let a = nextInt()
    let b = nextInt()
    let c = nextInt()
    let onePositive = a > 0 && b <= 0 && c <= 0 || b > 0 && a <= 0 && c <= 0 || c > 0 && a <= 0 && b <= 0
    return compareBoolean(onePositive)
}

// { "no": 15, "dat": "", "ans": "" }
const boolean15 = () => {
    let a = nextInt()
    let b = nextInt()
    let c = nextInt()
    let onePositive = a > 0 && b > 0 && c <= 0 || a > 0 && b <= 0 && c > 0 || a <= 0 && b > 0 && c > 0
    return compareBoolean(onePositive)
}

// { "no": 16, "dat": "", "ans": "" }
const boolean16 = () => {
    let number = nextInt()
    let result = number > 9 && number < 100 && number % 2 == 0
    return compareBoolean(result)
}

// { "no": 17, "dat": "", "ans": "" }
const boolean17 = () => {
    let number = nextInt()
    let result = number > 99 && number < 1000 && number % 2 != 0
    return compareBoolean(result)
}

// { "no": 18, "dat": "", "ans": "" }
const boolean18 = () => {
    let a = nextInt()
    let b = nextInt()
    let c = nextInt()
    let pair = a == b || b == c || c == a
    return compareBoolean(pair)
}

// { "no": 19, "dat": "", "ans": "" }
const boolean19 = () => {
    let a = nextInt()
    let b = nextInt()
    let c = nextInt()
    let pair = a == -b || b == -c || c == -a
    return compareBoolean(pair)
}

// { "no": 20, "dat": "", "ans": "" }
const boolean20 = () => {
    let number = nextInt()
    let hundreds = parseInt(number / 100)
    let tens = parseInt(number / 10) % 10
    let ones = number % 10
    let allDifferent = hundreds != tens && tens != ones && ones != hundreds
    return compareBoolean(allDifferent)
}

// { "no": 21, "dat": "", "ans": "" }
const boolean21 = () => {
    let number = nextInt()
    let hundreds = parseInt(number / 100)
    let tens = parseInt(number / 10) % 10
    let ones = number % 10
    let ascOrdered = hundreds < tens && tens < ones
    return compareBoolean(ascOrdered)
}

// { "no": 22, "dat": "", "ans": "" }
const boolean22 = () => {
    let number = nextInt()
    let hundreds = parseInt(number / 100)
    let tens = parseInt(number / 10) % 10
    let ones = number % 10
    let isSorted = hundreds < tens && tens < ones || hundreds > tens && tens > ones
    return compareBoolean(isSorted)
}

// { "no": 23, "dat": "", "ans": "" }
const boolean23 = () => {
    let number = nextInt()
    let isPalindrome = parseInt(number / 100) == number % 10 * 10 + parseInt(number / 10) % 10
    return compareBoolean(isPalindrome)
}

// { "no": 24, "dat": "2", "ans": "" }
const boolean24 = () => {
    let a = nextFloat()
    let b = nextFloat()
    let c = nextFloat()
    let d = b * b - 4 * a * c
    let hasRoot = d >= 0
    return compareBoolean(hasRoot)
}

// { "no": 25, "dat": "2", "ans": "" }
const boolean25 = () => {
    let x = nextFloat()
    let y = nextFloat()
    let inSecond = x < 0 && y > 0
    return compareBoolean(inSecond)
}

// { "no": 26, "dat": "2", "ans": "" }
const boolean26 = () => {
    let x = nextFloat()
    let y = nextFloat()
    let inFourth = x > 0 && y < 0
    return compareBoolean(inFourth)
}

// { "no": 27, "dat": "2", "ans": "" }
const boolean27 = () => {
    let x = nextFloat()
    let y = nextFloat()
    let result = x < 0
    return compareBoolean(result)
}

// { "no": 28, "dat": "2", "ans": "" }
const boolean28 = () => {
    let x = nextFloat()
    let y = nextFloat()
    let result = x * y > 0
    return compareBoolean(result)
}

// { "no": 29, "dat": "2", "ans": "" }
const boolean29 = () => {
    let x = nextFloat()
    let y = nextFloat()
    let x1 = nextFloat()
    let y1 = nextFloat()
    let x2 = nextFloat()
    let y2 = nextFloat()
    let isInside = x > x1 && x < x2 && y > y2 && y < y1
    return compareBoolean(isInside)
}

// { "no": 30, "dat": "", "ans": "" }
const boolean30 = () => {
    let a = nextInt()
    let b = nextInt()
    let c = nextInt()
    let isEquilateral = a == b && b == c
    return compareBoolean(isEquilateral)
}

// { "no": 31, "dat": "", "ans": "" }
const boolean31 = () => {
    let a = nextInt()
    let b = nextInt()
    let c = nextInt()
    let isIsosceles = a == b || b == c || c == a
    return compareBoolean(isIsosceles)
}

// { "no": 32, "dat": "", "ans": "" }
const boolean32 = () => {
    let a = nextInt()
    let b = nextInt()
    let c = nextInt()
    let rightTriangle = a * a + b * b == c * c || a * a + c * c == b * b || b * b + c * c == a * a
    return compareBoolean(rightTriangle)
}

// { "no": 33, "dat": "", "ans": "" }
const boolean33 = () => {
    let a = nextInt()
    let b = nextInt()
    let c = nextInt()
    let exists = a + b > c && a + c > b && b + c > a
    return compareBoolean(exists)
}

// { "no": 34, "dat": "", "ans": "" }
const boolean34 = () => {
    let x = nextInt()
    let y = nextInt()
    let isWhite = (x + y) % 2 != 0
    return compareBoolean(isWhite)
}

// { "no": 35, "dat": "", "ans": "" }
const boolean35 = () => {
    let x1 = nextInt()
    let y1 = nextInt()
    let x2 = nextInt()
    let y2 = nextInt()
    let sameColor = (x1 + y1) % 2 == (x2 + y2) % 2
    return compareBoolean(sameColor)
}

// { "no": 36, "dat": "", "ans": "" }
const boolean36 = () => {
    let x1 = nextInt()
    let y1 = nextInt()
    let x2 = nextInt()
    let y2 = nextInt()
    let rux = x1 == x2 || y1 == y2; // rook
    return compareBoolean(rux)
}

// { "no": 37, "dat": "", "ans": "" }
const boolean37 = () => {
    let x1 = nextInt()
    let y1 = nextInt()
    let x2 = nextInt()
    let y2 = nextInt()
    let shoh = Math.abs(x2 - x1) < 2 && Math.abs(y2 - y1) < 2; // king
    return compareBoolean(shoh)
}

// { "no": 38, "dat": "", "ans": "" }
const boolean38 = () => {
    let x1 = nextInt()
    let y1 = nextInt()
    let x2 = nextInt()
    let y2 = nextInt()
    let fil = Math.abs(x2 - x1) == Math.abs(y2 - y1); // bishop
    return compareBoolean(fil)
}

// { "no": 39, "dat": "", "ans": "" }
const boolean39 = () => {
    let x1 = nextInt()
    let y1 = nextInt()
    let x2 = nextInt()
    let y2 = nextInt()
    let farzin = x1 == x2 || y1 == y2 || Math.abs(x2 - x1) == Math.abs(y2 - y1); // queen
    return compareBoolean(farzin)
}

// { "no": 40, "dat": "", "ans": "" }
const boolean40 = () => {
    let x1 = nextInt()
    let y1 = nextInt()
    let x2 = nextInt()
    let y2 = nextInt()
    let asp = Math.abs(x2 - x1) == 1 && Math.abs(y2 - y1) == 2 || Math.abs(x2 - x1) == 2 && Math.abs(y2 - y1) == 1; // knight
    return compareBoolean(asp)
}