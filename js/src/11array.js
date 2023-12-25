import {
    nextFloat
    , compareFloat
    , nextInt
    , compareInt
    , initScanner
    , initChecker
    , inputsEOF
    , outputsEOF
} from './io.js'

(async () => {
    for (let taskNo = 1; taskNo <= 140; taskNo++) {
        const taskNoStr = String(taskNo).padStart(3, 0)
        for (let testNo = 1; testNo <= 100; testNo++) {
            const testNoStr = String(testNo).padStart(3, 0)
            const filePath = `../data/kt-gov2/11array/Array${taskNoStr}/${testNoStr}/${testNoStr}`

            await initScanner(filePath + '.dat')
            await initChecker(filePath + '.ans')

            let ok = true
            switch (taskNo) {
                case 1: ok = array1(); break
                case 2: ok = array2(); break
                case 3: ok = array3(); break
                case 4: ok = array4(); break
                case 5: ok = array5(); break
                case 6: ok = array6(); break
                case 7: ok = array7(); break
                case 8: ok = array8(); break
                case 9: ok = array9(); break
                case 10: ok = array10(); break
                case 11: ok = array11(); break
                case 12: ok = array12(); break
                case 13: ok = array13(); break
                case 14: ok = array14(); break
                case 15: ok = array15(); break
                case 16: ok = array16(); break
                case 17: ok = array17(); break
                case 18: ok = array18(); break
                case 19: ok = array19(); break
                case 20: ok = array20(); break
                case 21: ok = array21(); break
                case 22: ok = array22(); break
                case 23: ok = array23(); break
                case 24: ok = array24(); break
                case 25: ok = array25(); break
                case 26: ok = array26(); break
                case 27: ok = array27(); break
                case 28: ok = array28(); break
                case 29: ok = array29(); break
                case 30: ok = array30(); break
                case 31: ok = array31(); break
                case 32: ok = array32(); break
                case 33: ok = array33(); break
                case 34: ok = array34(); break
                case 35: ok = array35(); break
                case 36: ok = array36(); break
                case 37: ok = array37(); break
                case 38: ok = array38(); break
                case 39: ok = array39(); break
                case 40: ok = array40(); break
                case 41: ok = array41(); break
                case 42: ok = array42(); break
                case 43: ok = array43(); break
                case 44: ok = array44(); break
                case 45: ok = array45(); break
                case 46: ok = array46(); break
                case 47: ok = array47(); break
                case 48: ok = array48(); break
                case 49: ok = array49(); break
                case 50: ok = array50(); break
                case 51: ok = array51(); break
                case 52: ok = array52(); break
                case 53: ok = array53(); break
                case 54: ok = array54(); break
                case 55: ok = array55(); break
                case 56: ok = array56(); break
                case 57: ok = array57(); break
                case 58: ok = array58(); break
                case 59: ok = array59(); break
                case 60: ok = array60(); break
                case 61: ok = array61(); break
                case 62: ok = array62(); break
                case 63: ok = array63(); break
                case 64: ok = array64(); break
                case 65: ok = array65(); break
                case 66: ok = array66(); break
                case 67: ok = array67(); break
                case 68: ok = array68(); break
                case 69: ok = array69(); break
                case 70: ok = array70(); break
                case 71: ok = array71(); break
                case 72: ok = array72(); break
                case 73: ok = array73(); break
                case 74: ok = array74(); break
                case 75: ok = array75(); break
                case 76: ok = array76(); break
                case 77: ok = array77(); break
                case 78: ok = array78(); break
                case 79: ok = array79(); break
                case 80: ok = array80(); break
                case 81: ok = array81(); break
                case 82: ok = array82(); break
                case 83: ok = array83(); break
                case 84: ok = array84(); break
                case 85: ok = array85(); break
                case 86: ok = array86(); break
                case 87: ok = array87(); break
                case 88: ok = array88(); break
                case 89: ok = array89(); break
                case 90: ok = array90(); break
                case 91: ok = array91(); break
                case 92: ok = array92(); break
                case 93: ok = array93(); break
                case 94: ok = array94(); break
                case 95: ok = array95(); break
                case 96: ok = array96(); break
                case 97: ok = array97(); break
                case 98: ok = array98(); break
                case 99: ok = array99(); break
                case 100: ok = array100(); break
                case 101: ok = array101(); break
                case 102: ok = array102(); break
                case 103: ok = array103(); break
                case 104: ok = array104(); break
                case 105: ok = array105(); break
                case 106: ok = array106(); break
                case 107: ok = array107(); break
                case 108: ok = array108(); break
                case 109: ok = array109(); break
                case 110: ok = array110(); break
                case 111: ok = array111(); break
                case 112: ok = array112(); break
                case 113: ok = array113(); break
                case 114: ok = array114(); break
                case 115: ok = array115(); break
                case 116: ok = array116(); break
                case 117: ok = array117(); break
                case 118: ok = array118(); break
                case 119: ok = array119(); break
                case 120: ok = array120(); break
                case 121: ok = array121(); break
                case 122: ok = array122(); break
                case 123: ok = array123(); break
                case 124: ok = array124(); break
                case 125: ok = array125(); break
                case 126: ok = array126(); break
                case 127: ok = array127(); break
                case 128: ok = array128(); break
                case 129: ok = array129(); break
                case 130: ok = array130(); break
                case 131: ok = array131(); break
                case 132: ok = array132(); break
                case 133: ok = array133(); break
                case 134: ok = array134(); break
                case 135: ok = array135(); break
                case 136: ok = array136(); break
                case 137: ok = array137(); break
                case 138: ok = array138(); break
                case 139: ok = array139(); break
                case 140: ok = array140(); break
            }

            if (!ok) {
                console.log("Array" + taskNoStr + " #" + testNoStr + " has NOT been tested!")
                return
            }

            if (!inputsEOF()) {
                console.log("Array" + taskNoStr + " #" + testNoStr + " input NOT EOF!")
                return
            }

            if (!outputsEOF()) {
                console.log("Array" + taskNoStr + " #" + testNoStr + " output NOT EOF!")
                return
            }
        }
        console.log("Array" + taskNoStr + " has been tested!")
    }
    console.log("Array has been tested!")
})()

// { "no": 1, "dat": "", "ans": "" }
const array1 = () => {
    let n = nextInt()
    let array = new Array(n)
    for (let i = 0; i < n; i++) {
        array[i] = 2 * i + 1
    }
    return compareInt(...array)
}

// { "no": 2, "dat": "", "ans": "" }
const array2 = () => {
    let n = nextInt()
    let mass = new Array(n)
    for (let i = 0; i < n; i++) {
        mass[i] = Math.pow(2, (i + 1))
    }
    return compareInt(...mass)
}

// { "no": 3, "dat": "2", "ans": "2" }
const array3 = () => {
    let n = nextInt()
    let a = nextFloat()
    let d = nextFloat()
    let mass = new Array(n)
    for (let i = 0; i < n; i++) {
        mass[i] = a + i * d
    }
    return compareFloat(2, ...mass)
}

// { "no": 4, "dat": "2", "ans": "2" }
const array4 = () => {
    let n = nextInt()
    let a = nextFloat()
    let r = nextFloat()
    let mass = new Array(n)
    mass[0] = a
    for (let i = 1; i < n; i++) {
        mass[i] = mass[i - 1] * r
    }
    return compareFloat(2, ...mass)
}

// { "no": 5, "dat": "", "ans": "" }
const array5 = () => {
    let n = nextInt()
    let mass = new Array(n)
    mass[0] = mass[1] = 1
    for (let i = 2; i < n; i++) {
        mass[i] = mass[i - 1] + mass[i - 2]
    }
    return compareInt(...mass)
}

// { "no": 6, "dat": "", "ans": "" }
const array6 = () => {
    let n = nextInt()
    let arr = new Array(n)
    arr[0] = nextInt()
    arr[1] = nextInt()
    let sum = arr[0] + arr[1]
    for (let i = 2; i < n; i++) {
        arr[i] = sum
        sum += arr[i]
    }
    return compareInt(...arr)
}

// { "no": 7, "dat": "2", "ans": "2" }
const array7 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
    }
    for (let i = n - 1; i >= 0; i--) {
        if (!compareFloat(2, arr[i])) {
            return false
        }
    }
    return true
}

// { "no": 8, "dat": "", "ans": "" }
const array8 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextInt()
    }
    let k = 0
    for (let i = 0; i < n; i++) {
        if (arr[i] % 2 != 0) {
            if (!compareInt(arr[i])) {
                return false
            }
            k++
        }
    }
    return compareInt(k)
}

// { "no": 9, "dat": "", "ans": "" }
const array9 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextInt()
    }
    let k = 0
    for (let i = n - 1; i >= 0; i--) {
        if (arr[i] % 2 == 0) {
            if (!compareInt(arr[i])) {
                return false
            }
            k++
        }
    }
    return compareInt(k)
}

// { "no": 10, "dat": "", "ans": "" }
const array10 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextInt()
    }
    for (let i = 0; i < n; i++) {
        if (arr[i] % 2 == 0) {
            if (!compareInt(arr[i])) {
                return false
            }
        }
    }
    for (let i = n - 1; i >= 0; i--) {
        if (arr[i] % 2 != 0) {
            if (!compareInt(arr[i])) {
                return false
            }
        }
    }
    return true
}

// { "no": 11, "dat": "2", "ans": "2" }
const array11 = () => {
    let n = nextInt()
    let a = new Array(n)
    for (let i = 0; i < n; i++) {
        a[i] = nextFloat()
    }
    let k = nextInt()
    for (let i = k - 1; i < n; i += k) {
        if (!compareFloat(2, a[i])) {
            return false
        }
    }
    return true
}

// { "no": 12, "dat": "2", "ans": "2" }
const array12 = () => {
    let n = nextInt()
    let a = new Array(n)
    for (let i = 0; i < n; i++) {
        a[i] = nextFloat()
    }
    for (let i = 1; i < n; i += 2) {
        if (!compareFloat(2, a[i])) {
            return false
        }
    }
    return true
}

// { "no": 13, "dat": "2", "ans": "2" }
const array13 = () => {
    let n = nextInt()
    let a = new Array(n)
    for (let i = 0; i < n; i++) {
        a[i] = nextFloat()
    }
    for (let i = n - 1; i >= 0; i -= 2) {
        if (!compareFloat(2, a[i])) {
            return false
        }
    }
    return true
}

// { "no": 14, "dat": "2", "ans": "2" }
const array14 = () => {
    let n = nextInt()
    let a = new Array(n)
    for (let i = 0; i < n; i++) {
        a[i] = nextFloat()
    }
    for (let i = 1; i < n; i += 2) {
        if (!compareFloat(2, a[i])) {
            return false
        }
    }
    for (let i = 0; i < n; i += 2) {
        if (!compareFloat(2, a[i])) {
            return false
        }
    }
    return true
}

// { "no": 15, "dat": "2", "ans": "2" }
const array15 = () => {
    let n = nextInt()
    let a = new Array(n)
    for (let i = 0; i < n; i++) {
        a[i] = nextFloat()
    }
    for (let i = 0; i < n; i += 2) {
        if (!compareFloat(2, a[i])) {
            return false
        }
    }
    for (let i = n - 1 - n % 2; i >= 0; i -= 2) {
        if (!compareFloat(2, a[i])) {
            return false
        }
    }
    return true
}

// { "no": 16, "dat": "2", "ans": "2" }
const array16 = () => {
    let n = nextInt()
    let a = new Array(n)
    for (let i = 0; i < n; i++) {
        a[i] = nextFloat()
    }
    let start = 0, finish = n - 1
    while (true) {
        if (start - 1 == finish) {
            break
        } else {
            if (!compareFloat(2, a[start++])) {
                return false
            }
        }
        if (start - 1 == finish) {
            break
        } else {
            if (!compareFloat(2, a[finish--])) {
                return false
            }
        }
    }
    return true
}

// { "no": 17, "dat": "2", "ans": "2" }
const array17 = () => {
    let n = nextInt()
    let a = new Array(n)
    for (let i = 0; i < n; i++) {
        a[i] = nextFloat()
    }
    let start = 0, finish = n - 1
    while (true) {
        if (start - 1 == finish) {
            break
        } else {
            if (!compareFloat(2, a[start++])) {
                return false
            }
        }
        if (start - 1 == finish) {
            break
        } else {
            if (!compareFloat(2, a[start++])) {
                return false
            }
        }
        if (start - 1 == finish) {
            break
        } else {
            if (!compareFloat(2, a[finish--])) {
                return false
            }
        }
        if (start - 1 == finish) {
            break
        } else {
            if (!compareFloat(2, a[finish--])) {
                return false
            }
        }
    }
    return true
}

// { "no": 18, "dat": "", "ans": "" }
const array18 = () => {
    let a = new Array(10)
    for (let i = 0; i < 10; i++) {
        a[i] = nextInt()
    }
    let element = 0
    for (let i = 0; i < 9; i++) {
        if (a[i] < a[9]) {
            element = a[i]
            break
        }
    }
    return compareFloat(2, element)
}

// { "no": 19, "dat": "", "ans": "" }
const array19 = () => {
    let a = new Array(10)
    for (let i = 0; i < 10; i++) {
        a[i] = nextInt()
    }
    let index = -1
    for (let i = 1; i < 9; i++) {
        if ((a[i] > a[0]) && (a[i] < a[9])) {
            index = i
        }
    }
    index++
    return compareInt(index)
}

// { "no": 20, "dat": "2", "ans": "2" }
const array20 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
    }
    let k = nextInt()
    let l = nextInt()
    let sum = 0
    for (let i = k - 1; i < l; i++) {
        sum += arr[i]
    }
    return compareFloat(2, sum)
}

// { "no": 21, "dat": "2", "ans": "2" }
const array21 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
    }
    let k = nextInt()
    let l = nextInt()
    let sum = 0
    for (let i = k - 1; i < l; i++) {
        sum += arr[i]
    }
    let aMean = sum / (l - k + 1)
    return compareFloat(2, aMean)
}

// { "no": 22, "dat": "2", "ans": "2" }
const array22 = () => {
    let n = nextInt()
    let arr = new Array(n)
    let sum = 0
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
        sum += arr[i]
    }
    let k = nextInt()
    let l = nextInt()
    for (let i = k - 1; i < l; i++) {
        sum -= arr[i]
    }
    return compareFloat(2, sum)
}

// { "no": 23, "dat": "2", "ans": "2" }
const array23 = () => {
    let n = nextInt()
    let arr = new Array(n)
    let sum = 0
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
        sum += arr[i]
    }
    let k = nextInt()
    let l = nextInt()
    for (let i = k - 1; i < l; i++) {
        sum -= arr[i]
    }
    let aMean = sum / (n - (l - k + 1))
    return compareFloat(2, aMean)
}

// { "no": 24, "dat": "", "ans": "" }
const array24 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextInt()
    }
    let farq = arr[1] - arr[0]
    for (let i = 1; i < n; i++) {
        if (arr[i] - arr[i - 1] != farq) {
            farq = 0
            break
        }
    }
    return compareInt(farq)
}

// { "no": 25, "dat": "", "ans": "" }
const array25 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextInt()
    }
    let maxraj = parseInt(arr[1] / arr[0])
    for (let i = 1; i < n; i++) {
        if ((arr[i] / arr[i - 1] - parseInt(arr[i] / arr[i - 1]) != 0)
            || (parseInt(arr[i] / arr[i - 1]) != maxraj)) {
            maxraj = 0
            break
        }
    }
    return compareInt(maxraj)
}

// { "no": 26, "dat": "", "ans": "" }
const array26 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextInt()
    }
    let index = -1
    for (let i = 1; i < n; i++) {
        if ((arr[i] + arr[i - 1]) % 2 == 0) {
            index = i
            break
        }
    }
    index++
    return compareInt(index)
}

// { "no": 27, "dat": "", "ans": "" }
const array27 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextInt()
    }
    let index = -1
    for (let i = 1; i < n; i++) {
        if (arr[i] * arr[i - 1] > 0) {
            index = i
            break
        }
    }
    index++
    return compareInt(index)
}

// { "no": 28, "dat": "2", "ans": "2" }
const array28 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
    }
    let minimal = 0
    for (let i = 1; i < n; i += 2) {
        if (i == 1 || arr[i] < minimal) {
            minimal = arr[i]
        }
    }
    return compareFloat(2, minimal)
}

// { "no": 29, "dat": "2", "ans": "2" }
const array29 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
    }
    let maximal = 0
    for (let i = 0; i < n; i += 2) {
        if (i == 0 || arr[i] > maximal) {
            maximal = arr[i]
        }
    }
    return compareFloat(2, maximal)
}

// { "no": 30, "dat": "2", "ans": "" }
const array30 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
    }
    let count = 0
    for (let i = 0; i < n - 1; i++) {
        if (arr[i] > arr[i + 1]) {
            if (!compareInt(i + 1)) {
                return false
            }
            count++
        }
    }
    return compareInt(count)
}

// { "no": 31, "dat": "2", "ans": "" }
const array31 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
    }
    let count = 0
    for (let i = n - 1; i > 0; i--) {
        if (arr[i] > arr[i - 1]) {
            if (!compareInt(i + 1)) {
                return false
            }
            count++
        }
    }
    return compareInt(count)
}

// { "no": 32, "dat": "2", "ans": "" }
const array32 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
    }
    let localMinimum = 0
    let index = -1
    let found = false
    //check the first element
    if (arr[0] < arr[1]) {
        localMinimum = arr[0]
        index = 0
        found = true
    }
    //check elements between if until not found
    for (let i = 1; !found && (i < n - 1); i++) {
        if ((arr[i] < arr[i - 1]) && (arr[i] < arr[i + 1])) {
            localMinimum = arr[i]
            index = i
            found = true
        }
    }
    //check the last element
    if (!found && (arr[n - 1] < arr[n - 2])) {
        localMinimum = arr[n - 1]
        index = n - 1
    }
    index++
    return compareInt(index)
}

// { "no": 33, "dat": "2", "ans": "" }
const array33 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
    }
    let localMaximum = 0
    let index = -1
    //check the first element
    if (arr[0] > arr[1]) {
        localMaximum = arr[0]
        index = 0
    }
    //check elements between
    for (let i = 1; i < n - 1; i++) {
        if ((arr[i] > arr[i - 1]) && (arr[i] > arr[i + 1])) {
            localMaximum = arr[i]
            index = i
        }
    }
    //check the last element
    if (arr[n - 1] > arr[n - 2]) {
        localMaximum = arr[n - 1]
        index = n - 1
    }
    index++
    return compareInt(index)
}

// { "no": 34, "dat": "2", "ans": "2" }
const array34 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
    }
    let maxLocalMin = 0
    let initialized = false
    if (arr[0] < arr[1]) {
        maxLocalMin = arr[0]
        initialized = true
    }
    for (let i = 1; i < n - 1; i++) {
        if ((arr[i] < arr[i - 1]) && (arr[i] < arr[i + 1])) {
            if (!initialized) {
                maxLocalMin = arr[i]
                initialized = true
            } else if (arr[i] > maxLocalMin) {
                maxLocalMin = arr[i]
            }
        }
    }
    if (arr[n - 1] < arr[n - 2]) {
        if (!initialized) {
            maxLocalMin = arr[n - 1]
            initialized = true
        } else if (arr[n - 1] > maxLocalMin) {
            maxLocalMin = arr[n - 1]
        }
    }
    return compareFloat(2, maxLocalMin)
}

// { "no": 35, "dat": "2", "ans": "2" }
const array35 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
    }
    let minLocalMax = 0
    let inited = false
    if (arr[0] > arr[1]) {
        minLocalMax = arr[0]
        inited = true
    }
    for (let i = 1; i < n - 1; i++) {
        if ((arr[i] > arr[i - 1]) && (arr[i] > arr[i + 1])) {
            if (!inited) {
                minLocalMax = arr[i]
                inited = true
            } else if (arr[i] < minLocalMax) {
                minLocalMax = arr[i]
            }
        }
    }
    if (arr[n - 1] > arr[n - 2]) {
        if (!inited) {
            minLocalMax = arr[n - 1]
            inited = true
        } else if (arr[n - 1] < minLocalMax) {
            minLocalMax = arr[n - 1]
        }
    }
    return compareFloat(2, minLocalMax)
}

// { "no": 36, "dat": "2", "ans": "2" }
const array36 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
    }
    let maximal = 0
    let inited = false
    for (let i = 1; i < n - 1; i++) {
        if ((arr[i] > arr[i - 1]) && (arr[i] < arr[i + 1]) || (arr[i] < arr[i - 1]) && (arr[i] > arr[i + 1])) {
            if (!inited) {
                maximal = arr[i]
                inited = true
            } else if (arr[i] > maximal) {
                maximal = arr[i]
            }
        }
    }
    return compareFloat(2, maximal)
}

// { "no": 37, "dat": "2", "ans": "" }
const array37 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
    }
    let seriesCount = 0
    let started = false
    for (let i = 1; i < n; i++) {
        if (arr[i - 1] > arr[i]) {
            if (started) {
                seriesCount++
                started = false
            }
        } else {
            started = true
        }
    }
    seriesCount += started ? 1 : 0
    return compareInt(seriesCount)
}

// { "no": 38, "dat": "2", "ans": "" }
const array38 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
    }
    let seriesCount = 0
    let started = false
    for (let i = 1; i < n; i++) {
        if (arr[i - 1] < arr[i]) {
            if (started) {
                seriesCount++
                started = false
            }
        } else {
            started = true
        }
    }
    seriesCount += started ? 1 : 0
    return compareInt(seriesCount)
}

// { "no": 39, "dat": "2", "ans": "" }
const array39 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
    }
    let isProgress = false, isRegress = false
    let progressesCount = 0, regressesCount = 0
    for (let i = 1; i < n; i++) {
        if (arr[i - 1] < arr[i]) {
            isProgress = true
            if (isRegress) {
                regressesCount++
                isRegress = false
            }
        } else {
            isRegress = true
            if (isProgress) {
                progressesCount++
                isProgress = false
            }
        }
    }
    progressesCount += +isProgress
    regressesCount += +isRegress
    let totalCount = progressesCount + regressesCount
    return compareInt(totalCount)
}

// { "no": 40, "dat": "2", "ans": "2" }
const array40 = () => {
    let r = nextFloat()
    let n = nextInt()
    let a = new Array(n)
    for (let i = 0; i < n; i++) {
        a[i] = nextFloat()
    }
    let distance, minDistance = Math.abs(r - a[0])
    let index = 0
    for (let i = 1; i < n; i++) {
        distance = Math.abs(r - a[i])
        if (distance < minDistance) {
            minDistance = distance
            index = i
        }
    }
    return compareFloat(2, a[index])
}

// { "no": 41, "dat": "2", "ans": "2" }
const array41 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
    }
    let sum, maxSum = arr[0] + arr[1]
    let index = 0
    for (let i = 1; i < n - 1; i++) {
        sum = arr[i] + arr[i + 1]
        if (sum > maxSum) {
            maxSum = sum
            index = i
        }
    }
    return compareFloat(2, arr[index], arr[index + 1])
}

// { "no": 42, "dat": "2", "ans": "2" }
const array42 = () => {
    let r = nextFloat()
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
    }
    let sum = arr[0] + arr[1]
    let minDistance = Math.abs(r - sum)
    let index = 0
    for (let i = 1; i < n - 1; i++) {
        sum = arr[i] + arr[i + 1]
        if (Math.abs(r - sum) < minDistance) {
            minDistance = Math.abs(r - sum)
            index = i
        }
    }
    return compareFloat(2, arr[index], arr[index + 1])
}

// { "no": 43, "dat": "", "ans": "" }
const array43 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextInt()
    }
    let seriesCount = 1
    for (let i = 1; i < n; i++) {
        if (arr[i] != arr[i - 1]) {
            seriesCount++
        }
    }
    return compareInt(seriesCount)
}

// { "no": 44, "dat": "", "ans": "" }
const array44 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextInt()
    }
    let one = -1, two = 0
    let found = false
    while (!found) {
        one++
        for (two = one + 1; two < n; two++) {
            if (arr[one] == arr[two]) {
                found = true
                break
            }
        }
    }
    one++
    two++
    return compareInt(one, two)
}

// { "no": 45, "dat": "2", "ans": "" }
const array45 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
    }
    let one = 0, two = 1
    let distance, minDistance = Math.abs(arr[one] - arr[two])
    let lastOne = one, lastTwo = two
    while (one < n) {
        for (two = 0; two < n; two++) {
            if (one == two) {
                continue
            }
            distance = Math.abs(arr[one] - arr[two])
            if (distance < minDistance) {
                minDistance = distance
                lastOne = one
                lastTwo = two
            }
        }
        one++
    }
    if (lastOne > lastTwo) {
        let temp = lastOne
        lastOne = lastTwo
        lastTwo = temp
    }
    lastOne++
    lastTwo++
    return compareInt(lastOne, lastTwo)
}

// { "no": 46, "dat": "2", "ans": "2" }
const array46 = () => {
    let r = nextFloat()
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
    }
    let one = 0, two = 1
    let lastOne = one, lastTwo = two
    let sum, closerSum = Math.abs(arr[one] + arr[two] - r)
    while (one < n) {
        for (two = 0; two < n; two++) {
            if (one == two) {
                continue
            }
            sum = arr[one] + arr[two]
            if (Math.abs(r - sum) < closerSum) {
                closerSum = Math.abs(r - sum)
                lastOne = one
                lastTwo = two
            }
        }
        one++
    }
    if (lastOne > lastTwo) {
        lastOne += lastTwo
        lastTwo = lastOne - lastTwo
        lastOne -= lastTwo
    }
    return compareFloat(2, arr[lastOne], arr[lastTwo])
}

// { "no": 47, "dat": "", "ans": "" }
const array47 = () => {
    let n = nextInt()
    let arr = new Array(n)
    let hast = new Array(101)
    for (let i = 0; i <= 100; i++) {
        hast[i] = false
    }
    for (let i = 0; i < n; i++) {
        arr[i] = nextInt()
        hast[arr[i]] = true
    }
    let differentsCount = 0
    for (let i = 1; i <= 100; i++) {
        if (hast[i]) {
            differentsCount++
        }
    }
    return compareInt(differentsCount)
}

// { "no": 48, "dat": "", "ans": "" }
const array48 = () => {
    let n = nextInt()
    let arr = new Array(n)
    let count = new Array(101)
    for (let i = 0; i <= 100; i++) {
        count[i] = 0
    }
    for (let i = 0; i < n; i++) {
        arr[i] = nextInt()
        count[arr[i]]++
    }
    let maxCount = 0
    for (let i = 1; i <= 100; i++) {
        if (count[i] > 0) {
            if (maxCount == 0) {
                maxCount = count[i]
            } else if (count[i] > maxCount) {
                maxCount = count[i]
            }
        }
    }
    return compareInt(maxCount)
}

// { "no": 49, "dat": "", "ans": "" }
const array49 = () => {
    let n = nextInt()
    let arr = new Array(n)
    let hast = new Array(n + 1)
    for (let i = 0; i <= n; i++) {
        hast[i] = false
    }
    let errorNo = -1
    for (let i = 0; i < n; i++) {
        arr[i] = nextInt()
    }
    for (let i = 0; i < n; i++) {
        if ((arr[i] >= 1) && (arr[i] <= n)) {
            if (hast[arr[i]]) {
                errorNo = i
                break
            } else {
                hast[arr[i]] = true
            }
        } else {
            errorNo = i
            break
        }
    }
    errorNo++
    return compareInt(errorNo)
}

// { "no": 50, "dat": "", "ans": "" }
const array50 = () => {
    let n = nextInt()
    let arr = new Array(n)
    let count = 0
    for (let i = 0; i < n; i++) {
        arr[i] = nextInt()
        for (let j = 0; j < i; j++) {
            if (arr[i] < arr[j]) {
                count++
            }
        }
    }
    return compareInt(count)
}

// { "no": 51, "dat": "2", "ans": "2" }
const array51 = () => {
    let n = nextInt()
    let a = new Array(n)
    for (let i = 0; i < n; i++) {
        a[i] = nextFloat()
    }
    let b = new Array(n)
    for (let i = 0; i < n; i++) {
        b[i] = nextFloat()
    }
    for (let i = 0; i < n; i++) {
        a[i] += b[i]
        b[i] = a[i] - b[i]
        a[i] -= b[i]
    }
    return compareFloat(2, ...a) && compareFloat(2, ...b)
}

// { "no": 52, "dat": "2", "ans": "2" }
const array52 = () => {
    let n = nextInt()
    let a = new Array(n)
    let b = new Array(n)
    for (let i = 0; i < n; i++) {
        a[i] = nextFloat()
        b[i] = a[i] < 5 ? 2 * a[i] : a[i] / 2
    }
    return compareFloat(2, ...b)
}

// { "no": 53, "dat": "2", "ans": "2" }
const array53 = () => {
    let n = nextInt()
    let a = new Array(n)
    for (let i = 0; i < n; i++) {
        a[i] = nextFloat()
    }
    let b = new Array(n)
    for (let i = 0; i < n; i++) {
        b[i] = nextFloat()
    }
    let c = new Array(n)
    for (let i = 0; i < n; i++) {
        c[i] = a[i] > b[i] ? a[i] : b[i]
    }
    return compareFloat(2, ...c)
}

// { "no": 54, "dat": "", "ans": "" }
const array54 = () => {
    let n = nextInt()
    let a = new Array(n)
    let size = 0
    for (let i = 0; i < n; i++) {
        a[i] = nextInt()
        if (a[i] % 2 == 0) {
            size++
        }
    }
    let b = new Array(size)
    let index = 0
    for (let i = 0; i < n; i++) {
        if (a[i] % 2 == 0) {
            b[index++] = a[i]
        }
    }
    return compareInt(size) && compareFloat(2, ...b)
}

// { "no": 55, "dat": "", "ans": "" }
const array55 = () => {
    let n = nextInt()
    let a = new Array(n)
    let size = parseInt(n / 2) + n % 2
    let b = new Array(size)
    let index = 0
    for (let i = 0; i < n; i++) {
        a[i] = nextInt()
        if (i % 2 == 0) {
            b[index++] = a[i]
        }
    }
    return compareInt(size) && compareFloat(2, ...b)
}

// { "no": 56, "dat": "", "ans": "" }
const array56 = () => {
    let n = nextInt()
    let a = new Array(n)
    for (let i = 0; i < n; i++) {
        a[i] = nextInt()
    }
    let size = parseInt(n / 3)
    let b = new Array(size)
    let index = 0
    for (let i = 2; i < n; i += 3) {
        b[index++] = a[i]
    }
    return compareInt(size) && compareFloat(2, ...b)
}

// { "no": 57, "dat": "", "ans": "" }
const array57 = () => {
    let n = nextInt()
    let a = new Array(n)
    for (let i = 0; i < n; i++) {
        a[i] = nextInt()
    }
    let b = new Array(n)
    let index = 0
    for (let i = 1; i < n; i += 2) {
        b[index++] = a[i]
    }
    for (let i = 0; i < n; i += 2) {
        b[index++] = a[i]
    }
    return compareFloat(2, ...b)
}

// { "no": 58, "dat": "2", "ans": "2" }
const array58 = () => {
    let n = nextInt()
    let a = new Array(n)
    let b = new Array(n)
    let sum = 0
    for (let i = 0; i < n; i++) {
        a[i] = nextFloat()
        sum += a[i]
        b[i] = sum
    }
    return compareFloat(2, ...b)
}

// { "no": 59, "dat": "2", "ans": "2" }
const array59 = () => {
    let n = nextInt()
    let a = new Array(n)
    let b = new Array(n)
    let sum = 0
    for (let i = 0; i < n; i++) {
        a[i] = nextFloat()
        sum += a[i]
        b[i] = sum / (1 + i)
    }
    return compareFloat(2, ...b)
}

// { "no": 60, "dat": "2", "ans": "2" }
const array60 = () => {
    let n = nextInt()
    let a = new Array(n)
    for (let i = 0; i < n; i++) {
        a[i] = nextFloat()
    }
    let b = new Array(n)
    let sum = 0
    for (let i = n - 1; i >= 0; i--) {
        sum += a[i]
        b[i] = sum
    }
    return compareFloat(2, ...b)
}

// { "no": 61, "dat": "2", "ans": "2" }
const array61 = () => {
    let n = nextInt()
    let a = new Array(n)
    let sum = 0
    for (let i = 0; i < n; i++) {
        a[i] = nextFloat()
        sum += a[i]
    }
    let b = new Array(n)
    for (let i = 0; i < n; i++) {
        b[i] = sum / (n - i)
        sum -= a[i]
    }
    return compareFloat(2, ...b)
}

// { "no": 62, "dat": "2", "ans": "2" }
const array62 = () => {
    let n = nextInt()
    let a = new Array(n)
    let plusesCount = 0, minusesCount = 0
    for (let i = 0; i < n; i++) {
        a[i] = nextFloat()
        if (a[i] >= 0) {
            plusesCount++
        } else {
            minusesCount++
        }
    }
    let b = new Array(plusesCount)
    let c = new Array(minusesCount)
    for (let i = 0, iB = 0, iC = 0; i < n; i++) {
        if (a[i] >= 0) {
            b[iB++] = a[i]
        } else {
            c[iC++] = a[i]
        }
    }
    return compareInt(plusesCount) && compareFloat(2, ...b) && compareInt(minusesCount) && compareFloat(2, ...c)
}

// { "no": 63, "dat": "2", "ans": "2" }
const array63 = () => {
    let a = new Array(5)
    let b = new Array(5)
    let c = new Array(10)
    for (let i = 0; i < 5; i++) {
        a[i] = nextFloat()
    }
    for (let i = 0; i < 5; i++) {
        b[i] = nextFloat()
    }
    for (let i = 0, iA = 0, iB = 0; i < 10; i++) {
        if (iA >= 5) {
            c[i] = b[iB++]
        } else if (iB >= 5) {
            c[i] = a[iA++]
        } else {
            c[i] = a[iA] < b[iB] ? a[iA++] : b[iB++]
        }
    }
    return compareFloat(2, ...c)
}

// { "no": 64, "dat": "", "ans": "" }
const array64 = () => {
    let nA = nextInt()
    let a = new Array(nA)
    for (let i = 0; i < nA; i++) {
        a[i] = nextInt()
    }
    let nB = nextInt()
    let b = new Array(nB)
    for (let i = 0; i < nB; i++) {
        b[i] = nextInt()
    }
    let nC = nextInt()
    let c = new Array(nC)
    for (let i = 0; i < nC; i++) {
        c[i] = nextInt()
    }
    // yakjoyakunii massivhoi A va B: A + B = T
    let nT = nA + nB
    let t = new Array(nT)
    for (let i = 0, iA = 0, iB = 0; i < nT; i++) {
        if (iA >= nA) {
            t[i] = b[iB++]
        } else if (iB >= nB) {
            t[i] = a[iA++]
        } else {
            t[i] = a[iA] > b[iB] ? a[iA++] : b[iB++]
        }
    }
    // yakjoyakunii massivhoi T va C: T + C = D
    let nD = nT + nC
    let d = new Array(nD)
    for (let i = 0, iT = 0, iC = 0; i < nD; i++) {
        if (iT >= nT) {
            d[i] = c[iC++]
        } else if (iC >= nC) {
            d[i] = t[iT++]
        } else {
            d[i] = t[iT] > c[iC] ? t[iT++] : c[iC++]
        }
    }
    return compareInt(...d)
}

// { "no": 65, "dat": "2", "ans": "2" }
const array65 = () => {
    let n = nextInt()
    let a = new Array(n)
    for (let i = 0; i < n; i++) {
        a[i] = nextFloat()
    }
    let k = nextInt()
    let aK = a[k - 1]
    for (let i = 0; i < n; i++) {
        a[i] += aK
    }
    return compareFloat(2, ...a)
}

// { "no": 66, "dat": "", "ans": "" }
const array66 = () => {
    let n = nextInt()
    let arr = new Array(n)
    let index = -1
    for (let i = 0; i < n; i++) {
        arr[i] = nextInt()
        if (arr[i] % 2 == 0) {
            if (index < 0) {
                index = i
            } else {
                arr[i] += arr[index]
            }
        }
    }
    if (index >= 0) {
        arr[index] += arr[index]
    }
    return compareInt(...arr)
}

// { "no": 67, "dat": "", "ans": "" }
const array67 = () => {
    let n = nextInt()
    let arr = new Array(n)
    let index = -1
    for (let i = 0; i < n; i++) {
        arr[i] = nextInt()
        if (arr[i] % 2 != 0) {
            index = i
        }
    }
    if (index >= 0) {
        let element = arr[index]
        for (let i = 0; i < n; i++) {
            if (arr[i] % 2 != 0) {
                arr[i] += element
            }
        }
    }
    return compareFloat(2, ...arr)
}

// { "no": 68, "dat": "2", "ans": "2" }
const array68 = () => {
    let n = nextInt()
    let arr = new Array(n)
    let maxIndex = -1, minIndex = -1
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
        if (maxIndex < 0 || arr[i] >= arr[maxIndex]) {
            maxIndex = i
        }
        if (minIndex < 0 || arr[i] <= arr[minIndex]) {
            minIndex = i
        }
    }
    if (maxIndex >= 0 && minIndex >= 0) {
        arr[maxIndex] += arr[minIndex]
        arr[minIndex] = arr[maxIndex] - arr[minIndex]
        arr[maxIndex] -= arr[minIndex]
    }
    return compareFloat(2, ...arr)
}

// { "no": 69, "dat": "2", "ans": "2" }
const array69 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
    }
    for (let i = 1; i < n; i += 2) {
        arr[i] += arr[i - 1]
        arr[i - 1] = arr[i] - arr[i - 1]
        arr[i] -= arr[i - 1]
    }
    return compareFloat(2, ...arr)
}

// { "no": 70, "dat": "2", "ans": "2" }
const array70 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
    }
    for (let i = 0, j = parseInt(n / 2) + n % 2; j < n; i++, j++) {
        arr[i] += arr[j]
        arr[j] = arr[i] - arr[j]
        arr[i] -= arr[j]
    }
    return compareFloat(2, ...arr)
}

// { "no": 71, "dat": "2", "ans": "2" }
const array71 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
    }
    for (let i = 0, j = n - 1; i < j; i++, j--) {
        arr[i] += arr[j]
        arr[j] = arr[i] - arr[j]
        arr[i] -= arr[j]
    }
    return compareFloat(2, ...arr)
}

// { "no": 72, "dat": "2", "ans": "2" }
const array72 = () => {
    let n = nextInt()
    let a = new Array(n)
    for (let i = 0; i < n; i++) {
        a[i] = nextFloat()
    }
    let k = nextInt()
    let l = nextInt()
    for (let i = k - 1, j = l - 1; i < j; i++, j--) {
        a[i] += a[j]
        a[j] = a[i] - a[j]
        a[i] -= a[j]
    }
    return compareFloat(2, ...a)
}

// { "no": 73, "dat": "2", "ans": "2" }
const array73 = () => {
    let n = nextInt()
    let a = new Array(n)
    for (let i = 0; i < n; i++) {
        a[i] = nextFloat()
    }
    let k = nextInt()
    let l = nextInt()
    for (let i = k, j = l - 2; i < j; i++, j--) {
        a[i] += a[j]
        a[j] = a[i] - a[j]
        a[i] -= a[j]
    }
    return compareFloat(2, ...a)
}

// { "no": 74, "dat": "2", "ans": "2" }
const array74 = () => {
    let n = nextInt()
    let arr = new Array(n)
    let minIndex = 0, maxIndex = 0
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
        if (arr[i] > arr[maxIndex]) {
            maxIndex = i
        } else if (arr[i] < arr[minIndex]) {
            minIndex = i
        }
    }
    let from = minIndex < maxIndex ? minIndex : maxIndex
    let to = minIndex > maxIndex ? minIndex : maxIndex
    for (from++; from < to; from++) {
        arr[from] = 0
    }
    return compareFloat(2, ...arr)
}

// { "no": 75, "dat": "2", "ans": "2" }
const array75 = () => {
    let n = nextInt()
    let arr = new Array(n)
    let minIndex = 0, maxIndex = 0
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
        if (arr[i] > arr[maxIndex]) {
            maxIndex = i
        } else if (arr[i] < arr[minIndex]) {
            minIndex = i
        }
    }
    let from = minIndex < maxIndex ? minIndex : maxIndex
    let to = minIndex > maxIndex ? minIndex : maxIndex
    for (; from < to; from++, to--) {
        arr[from] += arr[to]
        arr[to] = arr[from] - arr[to]
        arr[from] -= arr[to]
    }
    return compareFloat(2, ...arr)
}

// { "no": 76, "dat": "2", "ans": "2" }
const array76 = () => {
    let n = nextInt()
    let arr = new Array(n)
    let copy = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
        copy[i] = arr[i]
    }
    if (copy[0] > copy[1]) {
        arr[0] = 0
    }
    for (let i = 1; i < n - 1; i++) {
        if ((copy[i] > copy[i - 1]) && (copy[i] > copy[i + 1])) {
            arr[i] = 0
        }
    }
    if (copy[n - 1] > copy[n - 2]) {
        arr[n - 1] = 0
    }
    return compareFloat(2, ...arr)
}

// { "no": 77, "dat": "2", "ans": "2" }
const array77 = () => {
    let n = nextInt()
    let arr = new Array(n)
    let copy = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
        copy[i] = arr[i]
    }
    if (copy[0] < copy[1]) {
        arr[0] *= arr[0]
    }
    for (let i = 1; i < n - 1; i++) {
        if ((copy[i] < copy[i - 1]) && (copy[i] < copy[i + 1])) {
            arr[i] *= arr[i]
        }
    }
    if (copy[n - 1] < copy[n - 2]) {
        arr[n - 1] *= arr[n - 1]
    }
    return compareFloat(2, ...arr)
}

// { "no": 78, "dat": "2", "ans": "2" }
const array78 = () => {
    let n = nextInt()
    let arr = new Array(n)
    let copy = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
        copy[i] = arr[i]
    }
    arr[0] = (copy[0] + copy[1]) / 2
    for (let i = 1; i < n - 1; i++) {
        arr[i] = (copy[i - 1] + copy[i] + copy[i + 1]) / 3
    }
    arr[n - 1] = (copy[n - 1] + copy[n - 2]) / 2
    return compareFloat(2, ...arr)
}

// { "no": 79, "dat": "2", "ans": "2" }
const array79 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
    }
    for (let i = n - 1; i > 0; i--) {
        arr[i] = arr[i - 1]
    }
    arr[0] = 0
    return compareFloat(2, ...arr)
}

// { "no": 80, "dat": "2", "ans": "2" }
const array80 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
    }
    for (let i = 0; i < n - 1; i++) {
        arr[i] = arr[i + 1]
    }
    arr[n - 1] = 0
    return compareFloat(2, ...arr)
}

// { "no": 81, "dat": "2", "ans": "2" }
const array81 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
    }
    let k = nextInt()
    for (let i = n - 1; i >= k; i--) {
        arr[i] = arr[i - k]
    }
    for (let i = k - 1; i >= 0; i--) {
        arr[i] = 0
    }
    return compareFloat(2, ...arr)
}

// { "no": 82, "dat": "2", "ans": "2" }
const array82 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
    }
    let k = nextInt()
    for (let i = 0; i < n - k; i++) {
        arr[i] = arr[i + k]
    }
    for (let i = n - k; i < n; i++) {
        arr[i] = 0
    }
    return compareFloat(2, ...arr)
}

// { "no": 83, "dat": "2", "ans": "2" }
const array83 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
    }
    let temp = arr[n - 1]
    for (let i = n - 1; i > 0; i--) {
        arr[i] = arr[i - 1]
    }
    arr[0] = temp
    return compareFloat(2, ...arr)
}

// { "no": 84, "dat": "2", "ans": "2" }
const array84 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
    }
    let temp = arr[0]
    for (let i = 0; i < n - 1; i++) {
        arr[i] = arr[i + 1]
    }
    arr[n - 1] = temp
    return compareFloat(2, ...arr)
}

// { "no": 85, "dat": "2", "ans": "2" }
const array85 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
    }
    let k = nextInt()
    let temp = new Array(k)
    for (let i = n - k, j = 0; i < n; i++, j++) {
        temp[j] = arr[i]
    }
    for (let i = n - 1; i >= k; i--) {
        arr[i] = arr[i - k]
    }
    for (let i = 0; i < k; i++) {
        arr[i] = temp[i]
    }
    return compareFloat(2, ...arr)
}

// { "no": 86, "dat": "2", "ans": "2" }
const array86 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
    }
    let k = nextInt()
    let temp = new Array(k)
    for (let i = 0; i < k; i++) {
        temp[i] = arr[i]
    }
    for (let i = 0; i < n - k; i++) {
        arr[i] = arr[i + k]
    }
    for (let i = 0, j = n - k; j < n; j++, i++) {
        arr[j] = temp[i]
    }
    return compareFloat(2, ...arr)
}

// { "no": 87, "dat": "2", "ans": "2" }
const array87 = () => {
    let n = nextInt()
    let arr = new Array(n)
    let index = -1
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
        if ((index < 0) && (i != 0)) {
            if (arr[i] > arr[0]) {
                index = i
            }
        }
    }
    index = index < 0 ? n - 1 : index - 1
    let temp = arr[0]
    for (let i = 0; i < index; i++) {
        arr[i] = arr[i + 1]
    }
    arr[index] = temp
    return compareFloat(2, ...arr)
}

// { "no": 88, "dat": "2", "ans": "2" }
const array88 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
    }
    let index = n
    for (let i = n - 2; i >= 0; i--) {
        if (arr[i] < arr[n - 1]) {
            index = i + 1
            break
        }
    }
    if (index == n) {
        index = 0
    }
    let temp = arr[n - 1]
    for (let i = n - 1; i > index; i--) {
        arr[i] = arr[i - 1]
    }
    arr[index] = temp
    return compareFloat(2, ...arr)
}

// { "no": 89, "dat": "2", "ans": "2" }
const array89 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
    }
    for (let i = 0; i < n - 1; i++) {
        if (arr[i] < arr[i + 1]) {
            arr[i] += arr[i + 1]
            arr[i + 1] = arr[i] - arr[i + 1]
            arr[i] -= arr[i + 1]
        }
    }
    for (let i = n - 1; i > 0; i--) {
        if (arr[i] > arr[i - 1]) {
            arr[i] += arr[i - 1]
            arr[i - 1] = arr[i] - arr[i - 1]
            arr[i] -= arr[i - 1]
        }
    }
    return compareFloat(2, ...arr)
}

// { "no": 90, "dat": "2", "ans": "2" }
const array90 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
    }
    let k = nextInt()
    arr.splice(k - 1, 1)
    n--
    return compareFloat(2, ...arr)
}

// { "no": 91, "dat": "2", "ans": "2" }
const array91 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
    }
    let k = nextInt()
    let l = nextInt()
    arr.splice(k - 1, l - k + 1)
    n = arr.length
    return compareInt(n) && compareFloat(2, ...arr)
}

// { "no": 92, "dat": "", "ans": "" }
const array92 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextInt()
    }
    for (let i = 0; i < n;) {
        if (arr[i] % 2 != 0) {
            arr.splice(i, 1)
            n--
        } else {
            i++
        }
    }
    return compareInt(n, ...arr)
}

// { "no": 93, "dat": "", "ans": "" }
const array93 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextInt()
    }
    for (let i = 1; i < n; i++) {
        arr.splice(i, 1)
    }
    n = arr.length
    return compareInt(...arr)
}

// { "no": 94, "dat": "", "ans": "" }
const array94 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextInt()
    }
    for (let i = 0; i < n; i++) {
        arr.splice(i, 1)
    }
    n = arr.length
    return compareInt(...arr)
}

// { "no": 95, "dat": "", "ans": "" }
const array95 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextInt()
    }
    for (let i = 1; i < n;) {
        if (arr[i] == arr[i - 1]) {
            arr.splice(i, 1)
            n--
        } else {
            i++
        }
    }
    return compareInt(...arr)
}

// { "no": 96, "dat": "", "ans": "" }
const array96 = () => {
    let n = nextInt()
    let arr = new Array(n)
    let hast = new Array(101)
    for (let i = 1; i <= 100; i++) {
        hast[i] = false
    }
    for (let i = 0; i < n; i++) {
        arr[i] = nextInt()
        if (!hast[arr[i]]) {
            hast[arr[i]] = true
        }
    }
    for (let i = 0; i < n;) {
        if (hast[arr[i]]) {
            hast[arr[i]] = false
            i++
        } else {
            arr.splice(i, 1)
            n--
        }
    }
    return compareInt(...arr)
}

// { "no": 97, "dat": "", "ans": "" }
const array97 = () => {
    let n = nextInt()
    let arr = new Array(n)
    let hast = new Array(101)
    for (let i = 1; i <= 100; i++) {
        hast[i] = false
    }
    for (let i = 0; i < n; i++) {
        arr[i] = nextInt()
        if (!hast[arr[i]]) {
            hast[arr[i]] = true
        }
    }
    for (let i = n - 1; i >= 0; i--) {
        if (hast[arr[i]]) {
            hast[arr[i]] = false
        } else {
            arr.splice(i, 1)
        }
    }
    return compareInt(...arr)
}

// { "no": 98, "dat": "", "ans": "" }
const array98 = () => {
    let count = new Array(101)
    for (let i = 1; i <= 100; i++) {
        count[i] = 0
    }
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextInt()
        count[arr[i]]++
    }
    let size = 0
    for (let i = 1; i <= 100; i++) {
        if (count[i] >= 3) {
            size += count[i]
        }
    }
    let temp = new Array(size)
    for (let i = 0, j = 0; i < n; i++) {
        if (count[arr[i]] < 3) {
            continue
        }
        temp[j++] = arr[i]
    }
    arr.splice(0, arr.length)
    arr = temp
    temp = null
    n = size
    return compareInt(n, ...arr)
}

// { "no": 99, "dat": "", "ans": "" }
const array99 = () => {
    let count = new Array(101)
    for (let i = 1; i <= 100; i++) {
        count[i] = 0
    }
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextInt()
        count[arr[i]]++
    }
    let size = n
    for (let i = 1; i <= 100; i++) {
        if (count[i] > 2) {
            size -= count[i]
        }
    }
    let temp = new Array(size)
    for (let i = 0, j = 0; i < n; i++) {
        if (count[arr[i]] > 2) {
            continue
        }
        temp[j++] = arr[i]
    }
    arr.splice(0, arr.length)
    arr = temp
    temp = null
    n = size
    return compareInt(n, ...arr)
}

// { "no": 100, "dat": "", "ans": "" }
const array100 = () => {
    let count = new Array(101)
    for (let i = 1; i <= 100; i++) {
        count[i] = 0
    }
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextInt()
        count[arr[i]]++
    }
    for (let i = 0; i < n;) {
        if (count[arr[i]] == 2) {
            arr.splice(i, 1)
            n--
        } else {
            i++
        }
    }
    return compareInt(n, ...arr)
}

// { "no": 101, "dat": "2", "ans": "2" }
const array101 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
    }
    let k = nextInt()
    arr.splice(k - 1, 0, 0)
    n++
    return compareFloat(2, ...arr)
}

// { "no": 102, "dat": "2", "ans": "2" }
const array102 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
    }
    let k = nextInt()
    if (k == n) {
        arr.push(0)
    } else {
        arr.splice(k, 0, 0)
    }
    n++
    return compareFloat(2, ...arr)
}

// { "no": 103, "dat": "2", "ans": "2" }
const array103 = () => {
    let n = nextInt()
    let arr = new Array(n)
    let maximalIndex = -1, minimalIndex = -1
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
        if (maximalIndex < 0 || arr[i] >= arr[maximalIndex]) {
            maximalIndex = i
        }
        if (minimalIndex < 0 || arr[i] <= arr[minimalIndex]) {
            minimalIndex = i
        }
    }
    arr.splice(minimalIndex, 0, 0)
    n++
    if (minimalIndex < maximalIndex) {
        maximalIndex++
    }
    if (maximalIndex == n - 1) {
        arr.push(0)
    } else {
        arr.splice(maximalIndex + 1, 0, 0)
    }
    n++
    return compareFloat(2, ...arr)
}

// { "no": 104, "dat": "2", "ans": "2" }
const array104 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
    }
    let k = nextInt()
    let m = nextInt()
    for (let i = 0; i < m; i++) {
        arr.splice(k - 1, 0, 0)
    }
    n += m
    return compareFloat(2, ...arr)
}

// { "no": 105, "dat": "2", "ans": "2" }
const array105 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
    }
    let k = nextInt()
    let m = nextInt()
    for (let i = 0; i < m; i++) {
        if (k == n) {
            arr.push(0)
        } else {
            arr.splice(k, 0, 0)
        }
    }
    n += m
    return compareFloat(2, ...arr)
}

// { "no": 106, "dat": "1", "ans": "1" }
const array106 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
    }
    for (let i = 1; i < n;) {
        arr.splice(i, 0, arr[i])
        n++
        i += 3
    }
    return compareFloat(2, ...arr)
}

// { "no": 107, "dat": "1", "ans": "1" }
const array107 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
    }
    for (let i = 0; i < n;) {
        arr.splice(i, 0, arr[i], arr[i])
        n += 2
        i += 4
    }
    return compareFloat(2, ...arr)
}

// { "no": 108, "dat": "2", "ans": "2" }
const array108 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
    }
    for (let i = 0; i < n; i++) {
        if (arr[i] >= 0) {
            arr.splice(i, 0, 0)
            i++
            n++
        }
    }
    return compareFloat(2, ...arr)
}

// { "no": 109, "dat": "2", "ans": "2" }
const array109 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextFloat()
    }
    for (let i = 0; i < n; i++) {
        if (arr[i] < 0) {
            if (i == n - 1) {
                arr.push(0)
            } else {
                arr.splice(i + 1, 0, 0)
            }
            i++
            n++
        }
    }
    return compareFloat(2, ...arr)
}

// { "no": 110, "dat": "", "ans": "" }
const array110 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextInt()
    }
    for (let i = 0; i < n; i++) {
        if (arr[i] % 2 == 0) {
            arr.splice(i, 0, arr[i])
            i++
            n++
        }
    }
    return compareInt(...arr)
}

// { "no": 111, "dat": "", "ans": "" }
const array111 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextInt()
    }
    for (let i = 0; i < n; i++) {
        if (arr[i] % 2 != 0) {
            arr.splice(i, 0, arr[i], arr[i])
            i += 2
            n += 2
        }
    }
    return compareInt(...arr)
}

// { "no": 112, "dat": "2", "ans": "2" }
const array112 = () => {
    let n = nextInt()
    let a = new Array(n)
    for (let i = 0; i < n; i++) {
        a[i] = nextFloat()
    }
    for (let i = 0; i < n - 1; i++) {
        for (let j = 1; j < n - i; j++) {
            if (a[j] < a[j - 1]) {
                a[j] += a[j - 1]
                a[j - 1] = a[j] - a[j - 1]
                a[j] -= a[j - 1]
            }
        }
        if (!compareFloat(2, ...a)) {
            return false
        }
    }
    return true
}

// { "no": 113, "dat": "2", "ans": "2" }
const array113 = () => {
    let n = nextInt()
    let a = new Array(n)
    for (let i = 0; i < n; i++) {
        a[i] = nextFloat()
    }
    let maximalIndex
    for (let i = n - 1; i > 0; i--) {
        maximalIndex = i
        for (let j = i - 1; j >= 0; j--) {
            if (a[j] > a[maximalIndex]) {
                maximalIndex = j
            }
        }
        if (maximalIndex != i) {
            a[i] += a[maximalIndex]
            a[maximalIndex] = a[i] - a[maximalIndex]
            a[i] -= a[maximalIndex]
        }
        if (!compareFloat(2, ...a)) {
            return false
        }
    }
    return true
}

// { "no": 114, "dat": "2", "ans": "2" }
const array114 = () => {
    let n = nextInt()
    let a = new Array(n)
    for (let i = 0; i < n; i++) {
        a[i] = nextFloat()
    }
    for (let i = 1; i < n; i++) {
        for (let j = i; j > 0; j--) {
            if (a[j] > a[j - 1]) {
                break
            }
            a[j] += a[j - 1]
            a[j - 1] = a[j] - a[j - 1]
            a[j] -= a[j - 1]
        }
        if (!compareFloat(2, ...a)) {
            return false
        }
    }
    return true
}

// { "no": 115, "dat": "2", "ans": "" }
const array115 = () => {
    let n = nextInt()
    let a = new Array(n)
    let indexes = new Array(n)
    for (let i = 0; i < n; i++) {
        a[i] = nextFloat()
        indexes[i] = i
    }
    for (let i = 0; i < n - 1; i++) {
        for (let j = 1; j < n - i; j++) {
            if (a[indexes[j]] < a[indexes[j - 1]]) {
                indexes[j] += indexes[j - 1]
                indexes[j - 1] = indexes[j] - indexes[j - 1]
                indexes[j] -= indexes[j - 1]
            }
        }
    }
    for (let i = 0; i < n; i++) {
        indexes[i]++
    }
    return compareInt(...indexes)
}

// { "no": 116, "dat": "", "ans": "" }
const array116 = () => {
    let n = nextInt()
    let a = new Array(n)
    let seriesCount = 0
    for (let i = 0; i < n; i++) {
        a[i] = nextInt()
        if (seriesCount == 0) {
            seriesCount++
        } else if (a[i] != a[i - 1]) {
            seriesCount++
        }
    }
    let b = new Array(seriesCount)
    let c = new Array(seriesCount)
    let count = 1, index = 0
    for (let i = 1; i < n; i++) {
        if (a[i] != a[i - 1]) {
            b[index] = count
            c[index] = a[i - 1]
            index++
            count = 1
            continue
        }
        count++
    }
    b[index] = count
    c[index] = a[n - 1]
    return compareInt(...b) && compareInt(...c)
}

// { "no": 117, "dat": "", "ans": "" }
const array117 = () => {
    let n = nextInt()
    let arr = new Array(n)
    let seriesCount = 0
    for (let i = 0; i < n; i++) {
        arr[i] = nextInt()
        if (seriesCount == 0 || arr[i] != arr[i - 1]) {
            seriesCount++
        }
    }
    arr.splice(0, 0, 0)
    n++
    for (let i = 2; i < n; i++) {
        if (arr[i] != arr[i - 1]) {
            arr.splice(i, 0, 0)
            i++
            n++
        }
    }
    return compareInt(...arr)
}

// { "no": 118, "dat": "", "ans": "" }
const array118 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextInt()
    }
    for (let i = 1; i < n; i++) {
        if (arr[i] != arr[i - 1]) {
            arr.splice(i, 0, 0)
            i++
            n++
        }
    }
    arr.push(0)
    return compareInt(...arr)
}

// { "no": 119, "dat": "", "ans": "" }
const array119 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextInt()
    }
    arr.splice(0, 0, arr[0])
    n++
    for (let i = 2; i < n; i++) {
        if (arr[i] != arr[i - 1]) {
            arr.splice(i, 0, arr[i])
            i++
            n++
        }
    }
    return compareInt(...arr)
}

// { "no": 120, "dat": "", "ans": "" }
const array120 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextInt()
    }
    for (let i = 0; i < n - 1; i++) {
        if (arr[i] != arr[i + 1]) {
            arr.splice(i, 1)
            i--
            n--
        }
    }
    arr.splice(n - 1, 1)
    n--
    return compareInt(...arr)
}

// { "no": 121, "dat": "", "ans": "" }
const array121 = () => {
    let k = nextInt()
    let n = nextInt()
    let seriesCount = 0
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextInt()
        if (seriesCount == 0) {
            seriesCount++
        } else if (arr[i] != arr[i - 1]) {
            seriesCount++
        }
    }
    if (seriesCount >= k) {
        seriesCount = 0
        let count = 0, start = -1
        for (let i = 0; i < n; i++) {
            if (seriesCount == 0) {
                seriesCount++
            } else if (arr[i] != arr[i - 1]) {
                seriesCount++
            }
            if (seriesCount == k) {
                count++
            } else if (seriesCount > k) {
                start = i - count
                break
            }
        }
        if (start < 0) {
            start = n - count
        }
        for (let i = 0; i < count; i++) {
            arr.splice(start, 0, arr[start])
        }
    }
    n = arr.length
    return compareInt(...arr)
}

// { "no": 122, "dat": "", "ans": "" }
const array122 = () => {
    let k = nextInt()
    let n = nextInt()
    let i, series = 0
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextInt()
        if (series == 0) {
            series++
        } else if (arr[i] != arr[i - 1]) {
            series++
        }
    }
    if (series >= k) {
        series = 0
        let count = 0, start = -1
        for (let i = 0; i < n; i++) {
            if (series == 0) {
                series++
            } else if (arr[i] != arr[i - 1]) {
                series++
            }
            if (series == k) {
                count++
            }
            if (series > k) {
                start = i - count
                break
            }
        }
        if (start < 0) {
            start = n - count
        }
        arr.splice(start, count)
    }
    n = arr.length
    return compareInt(...arr)
}

// { "no": 123, "dat": "", "ans": "" }
const array123 = () => {
    let k = nextInt()
    let n = nextInt()
    let arr = new Array(n)
    let series = 0
    let start1 = -1, finish1 = -1, startK = -1, finishK = -1
    for (let i = 0; i < n; i++) {
        arr[i] = nextInt()
        if (series == 0) {
            series++
            start1 = i
        } else if (arr[i] != arr[i - 1]) {
            if ((finish1 < 0) && (series == 1)) {
                finish1 = i - 1
            }
            if ((finishK < 0) && (series == k)) {
                finishK = i - 1
            }
            series++
            if ((startK < 0) && (series == k)) {
                startK = i
            }
        }
    }
    if ((finish1 < 0) && (series == 1)) {
        finish1 = n - 1
    }
    if ((finishK < 0) && (series == k)) {
        finishK = n - 1
    }
    if ((startK >= 0) && (finishK >= 0)) {
        let temp = new Array(n)
        let index = 0
        for (let i = startK; i <= finishK; i++) {
            temp[index++] = arr[i]
        }
        for (let i = finish1 + 1; i < startK; i++) {
            temp[index++] = arr[i]
        }
        for (let i = start1; i <= finish1; i++) {
            temp[index++] = arr[i]
        }
        for (let i = finishK + 1; i < n; i++) {
            temp[index++] = arr[i]
        }
        arr.splice(0, arr.length)
        arr = temp
        temp = null
    }
    return compareInt(...arr)
}

// { "no": 124, "dat": "", "ans": "" }
const array124 = () => {
    let k = nextInt()
    let n = nextInt()
    let arr = new Array(n)
    let series = 0
    let startN = -1, finishN = -1, startK = -1, finishK = -1
    for (let i = 0; i < n; i++) {
        arr[i] = nextInt()
        if (series == 0) {
            series++
            if ((startK < 0) && (series == k)) {
                startK = i
            }
        } else if (arr[i] != arr[i - 1]) {
            if ((finishK < 0) && (series == k)) {
                finishK = i - 1
            }
            series++
            if ((startK < 0) && (series == k)) {
                startK = i
            }
            startN = i
        }
    }
    if ((finishK < 0) && (series == k)) {
        finishK = n - 1
    }
    finishN = n - 1
    if ((startK >= 0) && (finishK >= 0) && (startK != startN) && (finishK != finishN)) {
        let temp = new Array(n)
        let index = 0
        for (let i = 0; i < startK; i++) {
            temp[index++] = arr[i]
        }
        for (let i = startN; i <= finishN; i++) {
            temp[index++] = arr[i]
        }
        for (let i = finishK + 1; i < startN; i++) {
            temp[index++] = arr[i]
        }
        for (let i = startK; i <= finishK; i++) {
            temp[index++] = arr[i]
        }
        arr.splice(0, arr.length)
        arr = temp
        temp = null
    }
    return compareInt(...arr)
}

// { "no": 125, "dat": "", "ans": "" }
const array125 = () => {
    let l = nextInt()
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextInt()
    }
    let count = 1
    for (let i = 1; i < n; i++) {
        if (arr[i] != arr[i - 1]) {
            if (count < l) {
                arr.splice(i - count, count, 0)
                i -= count - 1
                n -= count - 1
            }
            count = 1
        } else {
            count++
        }
    }
    if (count < l) {
        arr.splice(n - count, count, 0)
        n -= count - 1
    }
    return compareInt(...arr)
}

// { "no": 126, "dat": "", "ans": "" }
const array126 = () => {
    let l = nextInt()
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextInt()
    }
    let count = 1
    for (let i = 1; i < n; i++) {
        if (arr[i] != arr[i - 1]) {
            if (count == l) {
                arr.splice(i - count, count, 0)
                i -= count - 1
                n -= count - 1
            }
            count = 1
        } else {
            count++
        }
    }
    if (count == l) {
        arr.splice(n - count, count, 0)
        n -= count - 1
    }
    return compareInt(...arr)
}

// { "no": 127, "dat": "", "ans": "" }
const array127 = () => {
    let l = nextInt()
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextInt()
    }
    let count = 1
    for (let i = 1; i < n; i++) {
        if (arr[i] != arr[i - 1]) {
            if (count > l) {
                arr.splice(i - count, count, 0)
                i -= count - 1
                n -= count - 1
            }
            count = 1
        } else {
            count++
        }
    }
    if (count > l) {
        arr.splice(n - count, count, 0)
        n -= count - 1
    }
    return compareInt(...arr)
}

// { "no": 128, "dat": "", "ans": "" }
const array128 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextInt()
    }
    let count = 1, maxCount = 0, start = -1
    for (let i = 1; i < n; i++) {
        if (arr[i] != arr[i - 1]) {
            if (count > maxCount) {
                maxCount = count
                start = i - count
            }
            count = 1
        } else {
            count++
        }
    }
    if (count > maxCount) {
        maxCount = count
        start = n - count
    }
    arr.splice(start, 0, arr[start])
    return compareInt(...arr)
}

// { "no": 129, "dat": "", "ans": "" }
const array129 = () => {
    let n = nextInt()
    let arr = new Array(n)
    for (let i = 0; i < n; i++) {
        arr[i] = nextInt()
    }
    let count = 1, maxCount = 0, start = -1
    for (let i = 1; i < n; i++) {
        if (arr[i] != arr[i - 1]) {
            if (count >= maxCount) {
                maxCount = count
                start = i - count
            }
            count = 1
        } else {
            count++
        }
    }
    if (count >= maxCount) {
        maxCount = count
        start = n - count
    }
    arr.splice(start, 0, arr[start])
    return compareInt(...arr)
}

// { "no": 130, "dat": "", "ans": "" }
const array130 = () => {
    let n = nextInt()
    let arr = new Array(n)
    let count = 0, maxCount = 0
    for (let i = 0; i < n; i++) {
        arr[i] = nextInt()
        if (count == 0) {
            count++
        } else if (arr[i] != arr[i - 1]) {
            if (count > maxCount) {
                maxCount = count
            }
            count = 1
        } else {
            count++
        }
    }
    if (count > maxCount) {
        maxCount = count
    }
    let start
    count = 1
    for (let i = 1; i < n; i++) {
        if (arr[i] != arr[i - 1]) {
            if (count == maxCount) {
                start = i - count
                arr.splice(start, 0, arr[start])
                i++
                n++
            }
            count = 1
        } else {
            count++
        }
    }
    if (count == maxCount) {
        start = n - count
        arr.splice(start, 0, arr[start])
        n++
    }
    return compareInt(...arr)
}

class Point {
    constructor() {
        this.x = 0
        this.y = 0
    }
}

// { "no": 131, "dat": "2", "ans": "2" }
const array131 = () => {
    let n = nextInt()
    let a = new Array(n)
    for (let i = 0; i < n; i++) {
        a[i] = new Point()
        a[i].x = nextFloat()
        a[i].y = nextFloat()
    }
    let b = new Point()
    b.x = nextFloat()
    b.y = nextFloat()
    let i = 0, index = 0
    let minR = Math.sqrt(Math.pow(a[i].x - b.x, 2) + Math.pow(a[i].y - b.y, 2))
    for (i++; i < n; i++) {
        let r = Math.sqrt(Math.pow(a[i].x - b.x, 2) + Math.pow(a[i].y - b.y, 2))
        if (r < minR) {
            minR = r
            index = i
        }
    }
    return compareFloat(2, a[index].x, a[index].y)
}

// { "no": 132, "dat": "2", "ans": "2" }
const array132 = () => {
    let n = nextInt()
    let a = new Array(n)
    let nuqta = new Point()
    nuqta.x = nuqta.y = 0
    let maxR = 0
    for (let i = 0; i < n; i++) {
        a[i] = new Point()
        a[i].x = nextFloat()
        a[i].y = nextFloat()
        if ((a[i].x < 0) && (a[i].y > 0)) {
            if (nuqta.x == 0.0) {
                nuqta.x = a[i].x
                nuqta.y = a[i].y
                maxR = Math.sqrt(Math.pow(a[i].x, 2) + Math.pow(a[i].y, 2))
            } else {
                let r = Math.sqrt(Math.pow(a[i].x, 2) + Math.pow(a[i].y, 2))
                if (r > maxR) {
                    nuqta.x = a[i].x
                    nuqta.y = a[i].y
                    maxR = r
                }
            }
        }
    }
    return compareFloat(2, nuqta.x, nuqta.y)
}

// { "no": 133, "dat": "2", "ans": "2" }
const array133 = () => {
    let n = nextInt()
    let a = new Array(n)
    let nuqta = new Point()
    nuqta.x = nuqta.y = 0
    let minR = 0
    for (let i = 0; i < n; i++) {
        a[i] = new Point()
        a[i].x = nextFloat()
        a[i].y = nextFloat()
        if (a[i].x * a[i].y > 0) {
            if (nuqta.x == 0.0) {
                nuqta.x = a[i].x
                nuqta.y = a[i].y
                minR = Math.sqrt(a[i].x * a[i].x + a[i].y * a[i].y)
            } else {
                let r = Math.sqrt(a[i].x * a[i].x + a[i].y * a[i].y)
                if (r < minR) {
                    minR = r
                    nuqta.x = a[i].x
                    nuqta.y = a[i].y
                }
            }
        }
    }
    return compareFloat(2, nuqta.x, nuqta.y)
}

// { "no": 134, "dat": "2", "ans": "2" }
const array134 = () => {
    let n = nextInt()
    let a = new Array(n)
    let maxR = 0
    let one = -1, two = -1
    for (let i = 0; i < n; i++) {
        a[i] = new Point()
        a[i].x = nextFloat()
        a[i].y = nextFloat()
        for (let j = 0; j < i; j++) {
            let r = Math.sqrt(Math.pow(a[i].x - a[j].x, 2) + Math.pow(a[i].y - a[j].y, 2))
            if (one < 0) {
                maxR = r
                one = j
                two = i
            } else if (r > maxR) {
                maxR = r
                one = j
                two = i
            }
        }
    }
    return compareFloat(2, a[one].x, a[one].y, a[two].x, a[two].y, maxR)
}

// { "no": 135, "dat": "2", "ans": "2" }
const array135 = () => {
    let n1 = nextInt()
    let a = new Array(n1)
    for (let i = 0; i < n1; i++) {
        a[i] = new Point()
        a[i].x = nextFloat()
        a[i].y = nextFloat()
    }
    let n2 = nextInt()
    let b = new Array(n2)
    let minR = 0
    let one = -1, two = -1
    for (let i = 0; i < n2; i++) {
        b[i] = new Point()
        b[i].x = nextFloat()
        b[i].y = nextFloat()
        for (let j = 0; j < n1; j++) {
            let r = Math.sqrt(Math.pow(a[j].x - b[i].x, 2) + Math.pow(a[j].y - b[i].y, 2))
            if (one < 0) {
                minR = r
                one = j
                two = i
            } else if (r < minR) {
                minR = r
                one = j
                two = i
            }
        }
    }
    return compareFloat(2, minR, a[one].x, a[one].y, b[two].x, b[two].y)
}

// { "no": 136, "dat": "2", "ans": "2" }
const array136 = () => {
    let n = nextInt()
    let a = new Array(n)
    let sum = new Array(n)
    for (let i = 0; i < n; i++) {
        sum[i] = 0
    }
    for (let i = 0; i < n; i++) {
        a[i] = new Point()
        a[i].x = nextFloat()
        a[i].y = nextFloat()
        for (let j = 0; j < i; j++) {
            let r = Math.sqrt(Math.pow(a[i].x - a[j].x, 2) + Math.pow(a[i].y - a[j].y, 2))
            sum[i] += r
            sum[j] += r
        }
    }
    let minIndex = 0
    for (let i = 1; i < n; i++) {
        if (sum[i] < sum[minIndex]) {
            minIndex = i
        }
    }
    return compareFloat(2, a[minIndex].x, a[minIndex].y, sum[minIndex])
}

// { "no": 137, "dat": "2", "ans": "2" }
const array137 = () => {
    let n = nextInt()
    let a = new Array(n)
    for (let i = 0; i < n; i++) {
        a[i] = new Point()
        a[i].x = nextFloat()
        a[i].y = nextFloat()
    }
    let maxP = 0
    let index1 = 0, index2 = 0, index3 = 0
    for (let i = 0; i < n; i++) {
        for (let j = i + 1; j < n; j++) {
            for (let k = j + 1; k < n; k++) {
                let p = getPerim(a[i], a[j], a[k])
                if ((i == 0 && j == 1 && k == 2) || p > maxP) {
                    maxP = p
                    index1 = i
                    index2 = j
                    index3 = k
                }
            }
        }
    }
    return compareFloat(2, maxP, a[index1].x, a[index1].y, a[index2].x, a[index2].y, a[index3].x, a[index3].y)
}

const getPerim = (a, b, c) => {
    let one = Math.sqrt(Math.pow(a.x - b.x, 2) + Math.pow(a.y - b.y, 2))
    let two = Math.sqrt(Math.pow(a.x - c.x, 2) + Math.pow(a.y - c.y, 2))
    let three = Math.sqrt(Math.pow(b.x - c.x, 2) + Math.pow(b.y - c.y, 2))
    let p = one + two + three
    return p
}

// { "no": 138, "dat": "2", "ans": "2" }
const array138 = () => {
    let n = nextInt()
    let a = new Array(n)
    for (let i = 0; i < n; i++) {
        a[i] = new Point()
        a[i].x = nextFloat()
        a[i].y = nextFloat()
    }
    let minP = 0
    let firstTime = true
    let index1 = 0, index2 = 0, index3 = 0
    for (let i = 0; i < n; i++) {
        for (let j = i + 1; j < n; j++) {
            for (let k = j + 1; k < n; k++) {
                let p = getPerim(a[i], a[j], a[k])
                if (firstTime || p < minP) {
                    minP = p
                    index1 = i
                    index2 = j
                    index3 = k
                    firstTime = false
                }
            }
        }
    }
    return compareFloat(2, minP, a[index1].x, a[index1].y, a[index2].x, a[index2].y, a[index3].x, a[index3].y)
}

// { "no": 139, "dat": "", "ans": "" }
const array139 = () => {
    let n = nextInt()
    let a = new Array(n)
    for (let i = 0; i < n; i++) {
        a[i] = new Point()
        a[i].x = nextInt()
        a[i].y = nextInt()
        if (i > 0) {
            for (let j = i; j > 0; j--) {
                let areSorted = a[j - 1].x < a[j].x || a[j - 1].x == a[j].x && a[j - 1].y < a[j].y
                if (!areSorted) {
                    swap(a[j - 1], a[j])
                }
            }
        }
    }
    for (let i = 0; i < n; i++) {
        if (!compareInt(a[i].x, a[i].y)) {
            return false
        }
    }
    return true
}

const swap = (a, b) => {
    a.x += b.x
    b.x = a.x - b.x
    a.x -= b.x
    a.y += b.y
    b.y = a.y - b.y
    a.y -= b.y
}

// { "no": 140, "dat": "", "ans": "" }
const array140 = () => {
    let n = nextInt()
    let a = new Array(n)
    for (let i = 0; i < n; i++) {
        a[i] = new Point()
        a[i].x = nextInt()
        a[i].y = nextInt()
        for (let j = i; j > 0; j--) {
            let areSorted = a[j - 1].x + a[j - 1].y < a[j].x + a[j].y || (a[j - 1].x + a[j - 1].y == a[j].x + a[j].y && a[j - 1].x < a[j].x)
            if (areSorted) {
                swap(a[j - 1], a[j])
            }
        }
    }
    for (let i = 0; i < n; i++) {
        if (!compareInt(a[i].x, a[i].y)) {
            return false
        }
    }
    return true
}