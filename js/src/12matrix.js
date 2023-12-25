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
    for (let taskNo = 1; taskNo <= 100; taskNo++) {
        const taskNoStr = String(taskNo).padStart(3, 0)
        for (let testNo = 1; testNo <= 100; testNo++) {
            const testNoStr = String(testNo).padStart(3, 0)
            const filePath = `../data/kt-gov2/12matrix/Matrix${taskNoStr}/${testNoStr}/${testNoStr}`

            await initScanner(filePath + '.dat')
            await initChecker(filePath + '.ans')

            let ok = true
            switch (taskNo) {
                case 1: ok = matrix1(); break
                case 2: ok = matrix2(); break
                case 3: ok = matrix3(); break
                case 4: ok = matrix4(); break
                case 5: ok = matrix5(); break
                case 6: ok = matrix6(); break
                case 7: ok = matrix7(); break
                case 8: ok = matrix8(); break
                case 9: ok = matrix9(); break
                case 10: ok = matrix10(); break
                case 11: ok = matrix11(); break
                case 12: ok = matrix12(); break
                case 13: ok = matrix13(); break
                case 14: ok = matrix14(); break
                case 15: ok = matrix15(); break
                case 16: ok = matrix16(); break
                case 17: ok = matrix17(); break
                case 18: ok = matrix18(); break
                case 19: ok = matrix19(); break
                case 20: ok = matrix20(); break
                case 21: ok = matrix21(); break
                case 22: ok = matrix22(); break
                case 23: ok = matrix23(); break
                case 24: ok = matrix24(); break
                case 25: ok = matrix25(); break
                case 26: ok = matrix26(); break
                case 27: ok = matrix27(); break
                case 28: ok = matrix28(); break
                case 29: ok = matrix29(); break
                case 30: ok = matrix30(); break
                case 31: ok = matrix31(); break
                case 32: ok = matrix32(); break
                case 33: ok = matrix33(); break
                case 34: ok = matrix34(); break
                case 35: ok = matrix35(); break
                case 36: ok = matrix36(); break
                case 37: ok = matrix37(); break
                case 38: ok = matrix38(); break
                case 39: ok = matrix39(); break
                case 40: ok = matrix40(); break
                case 41: ok = matrix41(); break
                case 42: ok = matrix42(); break
                case 43: ok = matrix43(); break
                case 44: ok = matrix44(); break
                case 45: ok = matrix45(); break
                case 46: ok = matrix46(); break
                case 47: ok = matrix47(); break
                case 48: ok = matrix48(); break
                case 49: ok = matrix49(); break
                case 50: ok = matrix50(); break
                case 51: ok = matrix51(); break
                case 52: ok = matrix52(); break
                case 53: ok = matrix53(); break
                case 54: ok = matrix54(); break
                case 55: ok = matrix55(); break
                case 56: ok = matrix56(); break
                case 57: ok = matrix57(); break
                case 58: ok = matrix58(); break
                case 59: ok = matrix59(); break
                case 60: ok = matrix60(); break
                case 61: ok = matrix61(); break
                case 62: ok = matrix62(); break
                case 63: ok = matrix63(); break
                case 64: ok = matrix64(); break
                case 65: ok = matrix65(); break
                case 66: ok = matrix66(); break
                case 67: ok = matrix67(); break
                case 68: ok = matrix68(); break
                case 69: ok = matrix69(); break
                case 70: ok = matrix70(); break
                case 71: ok = matrix71(); break
                case 72: ok = matrix72(); break
                case 73: ok = matrix73(); break
                case 74: ok = matrix74(); break
                case 75: ok = matrix75(); break
                case 76: ok = matrix76(); break
                case 77: ok = matrix77(); break
                case 78: ok = matrix78(); break
                case 79: ok = matrix79(); break
                case 80: ok = matrix80(); break
                case 81: ok = matrix81(); break
                case 82: ok = matrix82(); break
                case 83: ok = matrix83(); break
                case 84: ok = matrix84(); break
                case 85: ok = matrix85(); break
                case 86: ok = matrix86(); break
                case 87: ok = matrix87(); break
                case 88: ok = matrix88(); break
                case 89: ok = matrix89(); break
                case 90: ok = matrix90(); break
                case 91: ok = matrix91(); break
                case 92: ok = matrix92(); break
                case 93: ok = matrix93(); break
                case 94: ok = matrix94(); break
                case 95: ok = matrix95(); break
                case 96: ok = matrix96(); break
                case 97: ok = matrix97(); break
                case 98: ok = matrix98(); break
                case 99: ok = matrix99(); break
                case 100: ok = matrix100(); break
            }

            if (!ok) {
                console.log("Matrix" + taskNoStr + " #" + testNoStr + " has NOT been tested!")
                return
            }

            if (!inputsEOF()) {
                console.log("Matrix" + taskNoStr + " #" + testNoStr + " input NOT EOF!")
                return
            }

            if (!outputsEOF()) {
                console.log("Matrix" + taskNoStr + " #" + testNoStr + " output NOT EOF!")
                return
            }
        }
        console.log("Matrix" + taskNoStr + " has been tested!")
    }
    console.log("Matrix has been tested!")
})()

// { "no": 1, "dat": "", "ans": "" }
const matrix1 = () => {
    let m = nextInt()
    let n = nextInt()
    let matrix = new Array(m)
    for (let i = 0; i < m; i++) {
        matrix[i] = new Array(n)
    }
    for (let i = 0; i < m; i++) {
        let value = (i + 1) * 10
        for (let j = 0; j < n; j++) {
            matrix[i][j] = value
        }
    }
    for (let i = 0; i < m; i++) {
        if (!compareInt(...matrix[i])) {
            return false
        }
    }
    return true
}

// { "no": 2, "dat": "", "ans": "" }
const matrix2 = () => {
    let m = nextInt()
    let n = nextInt()
    let matrix = new Array(m)
    for (let i = 0; i < m; i++) {
        matrix[i] = new Array(n)
    }
    for (let j = 0; j < n; j++) {
        let value = (j + 1) * 5
        for (let i = 0; i < m; i++) {
            matrix[i][j] = value
        }
    }
    for (let i = 0; i < m; i++) {
        if (!compareInt(...matrix[i])) {
            return false
        }
    }
    return true
}

// { "no": 3, "dat": "2", "ans": "2" }
const matrix3 = () => {
    let m = nextInt()
    let n = nextInt()
    let matrix = new Array(m)
    for (let i = 0; i < m; i++) {
        matrix[i] = new Array(n)
    }
    for (let i = 0; i < m; i++) {
        matrix[i][0] = nextFloat()
        for (let j = 1; j < n; j++) {
            matrix[i][j] = matrix[i][0]
        }
    }
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...matrix[i])) {
            return false
        }
    }
    return true
}

// { "no": 4, "dat": "2", "ans": "2" }
const matrix4 = () => {
    let m = nextInt()
    let n = nextInt()
    let matrix = new Array(m)
    for (let i = 0; i < m; i++) {
        matrix[i] = new Array(n)
    }
    for (let j = 0; j < n; j++) {
        matrix[0][j] = nextFloat()
        for (let i = 1; i < m; i++) {
            matrix[i][j] = matrix[0][j]
        }
    }
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...matrix[i])) {
            return false
        }
    }
    return true
}

// { "no": 5, "dat": "2", "ans": "2" }
const matrix5 = () => {
    let m = nextInt()
    let n = nextInt()
    let d = nextFloat()
    let matrix = new Array(m)
    for (let i = 0; i < m; i++) {
        matrix[i] = new Array(n)
    }
    for (let i = 0; i < m; i++) {
        matrix[i][0] = nextFloat()
    }
    for (let i = 0; i < m; i++) {
        for (let j = 1; j < n; j++) {
            matrix[i][j] = d + matrix[i][j - 1]
        }
    }
    for (let i = 0; i < m; i++) {
        for (let j = 0; j < n; j++) {
            if (!compareFloat(2, matrix[i][j])) {
                return false
            }
        }
    }
    return true
}

// { "no": 6, "dat": "2", "ans": "2" }
const matrix6 = () => {
    let m = nextInt()
    let n = nextInt()
    let d = nextFloat()
    let matrix = new Array(m)
    for (let i = 0; i < m; i++) {
        matrix[i] = new Array(n)
    }
    for (let j = 0; j < n; j++) {
        matrix[0][j] = nextFloat()
    }
    for (let i = 1; i < m; i++) {
        for (let j = 0; j < n; j++) {
            matrix[i][j] = d * matrix[i - 1][j]
        }
    }
    for (let i = 0; i < m; i++) {
        for (let j = 0; j < n; j++) {
            if (!compareFloat(2, matrix[i][j])) {
                return false
            }
        }
    }
    return true
}

// { "no": 7, "dat": "2", "ans": "2" }
const matrix7 = () => {
    let m = nextInt()
    let n = nextInt()
    let matrix = new Array(m)
    for (let i = 0; i < m; i++) {
        matrix[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matrix[i][j] = nextFloat()
        }
    }
    let k = nextInt()
    for (let j = 0; j < n; j++) {
        if (!compareFloat(2, matrix[k - 1][j])) {
            return false
        }
    }
    return true
}

// { "no": 8, "dat": "2", "ans": "2" }
const matrix8 = () => {
    let m = nextInt()
    let n = nextInt()
    let matrix = new Array(m)
    for (let i = 0; i < m; i++) {
        matrix[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matrix[i][j] = nextFloat()
        }
    }
    let k = nextInt()
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, matrix[i][k - 1])) {
            return false
        }
    }
    return true
}

// { "no": 9, "dat": "2", "ans": "2" }
const matrix9 = () => {
    let m = nextInt()
    let n = nextInt()
    let matrix = new Array(m)
    for (let i = 0; i < m; i++) {
        matrix[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matrix[i][j] = nextFloat()
        }
    }
    for (let i = 1; i < m; i += 2) {
        if (!compareFloat(2, ...matrix[i])) {
            return false
        }
    }
    return true
}

// { "no": 10, "dat": "2", "ans": "2" }
const matrix10 = () => {
    let m = nextInt()
    let n = nextInt()
    let matrix = new Array(m)
    for (let i = 0; i < m; i++) {
        matrix[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matrix[i][j] = nextFloat()
        }
    }
    for (let j = 0; j < n; j += 2) {
        for (let i = 0; i < m; i++) {
            if (!compareFloat(2, matrix[i][j])) {
                return false
            }
        }
    }
    return true
}

// { "no": 11, "dat": "2", "ans": "2" }
const matrix11 = () => {
    let m = nextInt()
    let n = nextInt()
    let matrix = new Array(m)
    for (let i = 0; i < m; i++) {
        matrix[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matrix[i][j] = nextFloat()
        }
    }
    for (let i = 0; i < m; i++) {
        if (i % 2 == 0) {
            for (let j = 0; j < n; j++) {
                if (!compareFloat(2, matrix[i][j])) {
                    return false
                }
            }
        } else {
            for (let j = n - 1; j >= 0; j--) {
                if (!compareFloat(2, matrix[i][j])) {
                    return false
                }
            }
        }
    }
    return true
}

// { "no": 12, "dat": "2", "ans": "2" }
const matrix12 = () => {
    let m = nextInt()
    let n = nextInt()
    let matrix = new Array(m)
    for (let i = 0; i < m; i++) {
        matrix[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matrix[i][j] = nextFloat()
        }
    }
    for (let j = 0; j < n; j++) {
        if (j % 2 == 0) {
            for (let i = 0; i < m; i++) {
                if (!compareFloat(2, matrix[i][j])) {
                    return false
                }
            }
        } else {
            for (let i = m - 1; i >= 0; i--) {
                if (!compareFloat(2, matrix[i][j])) {
                    return false
                }
            }
        }
    }
    return true
}

// { "no": 13, "dat": "2", "ans": "2" }
const matrix13 = () => {
    let m = nextInt()
    let a = new Array(m)
    for (let i = 0; i < m; i++) {
        a[i] = new Array(m)
        for (let j = 0; j < m; j++) {
            a[i][j] = nextFloat()
        }
    }
    for (let i = 0; i < m; i++) {
        for (let j = 0; j < m - i; j++) {
            if (!compareFloat(2, a[i][j])) {
                return false
            }
        }
        for (let j = i + 1; j < m; j++) {
            if (!compareFloat(2, a[j][m - 1 - i])) {
                return false
            }
        }
    }
    return true
}

// { "no": 14, "dat": "2", "ans": "2" }
const matrix14 = () => {
    let m = nextInt()
    let a = new Array(m)
    for (let i = 0; i < m; i++) {
        a[i] = new Array(m)
        for (let j = 0; j < m; j++) {
            a[i][j] = nextFloat()
        }
    }
    for (let j = 0; j < m; j++) {
        for (let i = 0; i < m - j; i++) {
            if (!compareFloat(2, a[i][j])) {
                return false
            }
        }
        for (let i = j + 1; i < m; i++) {
            if (!compareFloat(2, a[m - 1 - j][i])) {
                return false
            }
        }
    }
    return true
}

// { "no": 15, "dat": "2", "ans": "2" }
const matrix15 = () => {
    let m = nextInt()
    let a = new Array(m)
    for (let i = 0; i < m; i++) {
        a[i] = new Array(m)
        for (let j = 0; j < m; j++) {
            a[i][j] = nextFloat()
        }
    }
    let index, limit = parseInt(m / 2) + m % 2
    for (let i = 0; i < limit; i++) {
        index = m - i - 1
        for (let j = i; j <= index; j++) {
            if (!compareFloat(2, a[i][j])) {
                return false
            }
        }
        for (let j = i + 1; j <= index; j++) {
            if (!compareFloat(2, a[j][index])) {
                return false
            }
        }
        for (let j = index - 1; j >= i; j--) {
            if (!compareFloat(2, a[index][j])) {
                return false
            }
        }
        for (let j = index - 1; j > i; j--) {
            if (!compareFloat(2, a[j][i])) {
                return false
            }
        }
    }
    return true
}

// { "no": 16, "dat": "2", "ans": "2" }
const matrix16 = () => {
    let m = nextInt()
    let a = new Array(m)
    for (let i = 0; i < m; i++) {
        a[i] = new Array(m)
        for (let j = 0; j < m; j++) {
            a[i][j] = nextFloat()
        }
    }
    let index, limit = parseInt(m / 2) + m % 2
    for (let j = 0; j < limit; j++) {
        index = m - 1 - j
        for (let i = j; i <= index; i++) {
            if (!compareFloat(2, a[i][j])) {
                return false
            }
        }
        for (let i = j + 1; i <= index; i++) {
            if (!compareFloat(2, a[index][i])) {
                return false
            }
        }
        for (let i = index - 1; i >= j; i--) {
            if (!compareFloat(2, a[i][index])) {
                return false
            }
        }
        for (let i = index - 1; i > j; i--) {
            if (!compareFloat(2, a[j][i])) {
                return false
            }
        }
    }
    return true
}

// { "no": 17, "dat": "2", "ans": "2" }
const matrix17 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
        }
    }
    let k = nextInt()
    let sum = 0, zarb = 1
    for (let j = 0; j < n; j++) {
        sum += matr[k - 1][j]
        zarb *= matr[k - 1][j]
    }
    return compareFloat(2, sum, zarb)
}

// { "no": 18, "dat": "2", "ans": "2" }
const matrix18 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
        }
    }
    let k = nextInt()
    let sum = 0, zarb = 1
    for (let i = 0; i < m; i++) {
        sum += matr[i][k - 1]
        zarb *= matr[i][k - 1]
    }
    return compareFloat(2, sum, zarb)
}

// { "no": 19, "dat": "2", "ans": "2" }
const matrix19 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
        }
    }
    for (let i = 0; i < m; i++) {
        let sum = 0
        for (let j = 0; j < n; j++) {
            sum += matr[i][j]
        }
        if (!compareFloat(2, sum)) {
            return false
        }
    }
    return true
}

// { "no": 20, "dat": "2", "ans": "2" }
const matrix20 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
        }
    }
    for (let j = 0; j < n; j++) {
        let zarb = 1
        for (let i = 0; i < m; i++) {
            zarb *= matr[i][j]
        }
        if (!compareFloat(2, zarb)) {
            return false
        }
    }
    return true
}

// { "no": 21, "dat": "2", "ans": "2" }
const matrix21 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
        }
    }
    for (let i = 0; i < m; i += 2) {
        let sum = 0
        for (let j = 0; j < n; j++) {
            sum += matr[i][j]
        }
        let aMean = sum / n
        if (!compareFloat(2, aMean)) {
            return false
        }
    }
    return true
}

// { "no": 22, "dat": "2", "ans": "2" }
const matrix22 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
        }
    }
    for (let j = 1; j < n; j += 2) {
        let sum = 0
        for (let i = 0; i < m; i++) {
            sum += matr[i][j]
        }
        if (!compareFloat(2, sum)) {
            return false
        }
    }
    return true
}

// { "no": 23, "dat": "2", "ans": "2" }
const matrix23 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
        }
    }
    for (let i = 0; i < m; i++) {
        let minimal = matr[i][0]
        for (let j = 1; j < n; j++) {
            if (matr[i][j] < minimal) {
                minimal = matr[i][j]
            }
        }
        if (!compareFloat(2, minimal)) {
            return false
        }
    }
    return true
}

// { "no": 24, "dat": "2", "ans": "2" }
const matrix24 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
        }
    }
    for (let j = 0; j < n; j++) {
        let maximal = matr[0][j]
        for (let i = 1; i < m; i++) {
            if (matr[i][j] > maximal) {
                maximal = matr[i][j]
            }
        }
        if (!compareFloat(2, maximal)) {
            return false
        }
    }
    return true
}

// { "no": 25, "dat": "2", "ans": "2" }
const matrix25 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
        }
    }
    let maxSum = 0
    let maxSumRow = -1
    for (let i = 0; i < m; i++) {
        let sum = 0
        for (let j = 0; j < n; j++) {
            sum += matr[i][j]
        }
        if (maxSumRow < 0) {
            maxSum = sum
            maxSumRow = i
        } else if (sum > maxSum) {
            maxSum = sum
            maxSumRow = i
        }
    }
    maxSumRow++
    return compareInt(maxSumRow) && compareFloat(2, maxSum)
}

// { "no": 26, "dat": "2", "ans": "2" }
const matrix26 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
        }
    }
    let minZarb = 0
    let minZarbCol = -1
    for (let j = 0; j < n; j++) {
        let zarb = 1
        for (let i = 0; i < m; i++) {
            zarb *= matr[i][j]
        }
        if (minZarbCol < 0) {
            minZarbCol = j
            minZarb = zarb
        } else if (zarb < minZarb) {
            minZarbCol = j
            minZarb = zarb
        }
    }
    minZarbCol++
    return compareInt(minZarbCol) && compareFloat(2, minZarb)
}

// { "no": 27, "dat": "2", "ans": "2" }
const matrix27 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
        }
    }
    let minimal, maximalMin = 0
    for (let i = 0; i < m; i++) {
        minimal = matr[i][0]
        for (let j = 1; j < n; j++) {
            if (matr[i][j] < minimal) {
                minimal = matr[i][j]
            }
        }
        if (i == 0 || minimal > maximalMin) {
            maximalMin = minimal
        }
    }
    return compareFloat(2, maximalMin)
}

// { "no": 28, "dat": "2", "ans": "2" }
const matrix28 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
        }
    }
    let maximal, minimalMax = 0
    for (let j = 0; j < n; j++) {
        maximal = matr[0][j]
        for (let i = 1; i < m; i++) {
            if (matr[i][j] > maximal) {
                maximal = matr[i][j]
            }
        }
        if (j == 0 || maximal < minimalMax) {
            minimalMax = maximal
        }
    }
    return compareFloat(2, minimalMax)
}

// { "no": 29, "dat": "2", "ans": "" }
const matrix29 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
        }
    }
    for (let i = 0; i < m; i++) {
        let sum = 0
        for (let j = 0; j < n; j++) {
            sum += matr[i][j]
        }
        let aMean = sum / n
        let count = 0
        for (let j = 0; j < n; j++) {
            if (matr[i][j] < aMean) {
                count++
            }
        }
        if (!compareInt(count)) {
            return false
        }
    }
    return true
}

// { "no": 30, "dat": "2", "ans": "" }
const matrix30 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
        }
    }
    for (let j = 0; j < n; j++) {
        let sum = 0
        for (let i = 0; i < m; i++) {
            sum += matr[i][j]
        }
        let aMean = sum / m
        let count = 0
        for (let i = 0; i < m; i++) {
            if (matr[i][j] > aMean) {
                count++
            }
        }
        if (!compareInt(count)) {
            return false
        }
    }
    return true
}

// { "no": 31, "dat": "2", "ans": "" }
const matrix31 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    let sum = 0
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
            sum += matr[i][j]
        }
    }
    let aMean = sum / (m * n)
    let distance, closer = -1
    let row = 0, col = 0
    for (let i = 0; i < m; i++) {
        for (let j = 0; j < n; j++) {
            distance = Math.abs(matr[i][j] - aMean)
            if (closer < 0 || distance <= closer) {
                closer = distance
                row = i
                col = j
            }
        }
    }
    row++
    col++
    return compareInt(row, col)
}

// { "no": 32, "dat": "", "ans": "" }
const matrix32 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextInt()
        }
    }
    let positives, negatives, row = -1
    for (let i = 0; i < m; i++) {
        positives = negatives = 0
        for (let j = 0; j < n; j++) {
            if (matr[i][j] > 0) {
                positives++
            } else if (matr[i][j] < 0) {
                negatives++
            }
        }
        if (positives == negatives) {
            row = i
            break
        }
    }
    row++
    return compareInt(row)
}

// { "no": 33, "dat": "", "ans": "" }
const matrix33 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextInt()
        }
    }
    let positives, negatives, col = -1
    for (let j = 0; j < n; j++) {
        positives = negatives = 0
        for (let i = 0; i < m; i++) {
            if (matr[i][j] > 0) {
                positives++
            } else if (matr[i][j] < 0) {
                negatives++
            }
        }
        if (positives == negatives) {
            col = j
        }
    }
    col++
    return compareInt(col)
}

// { "no": 34, "dat": "", "ans": "" }
const matrix34 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextInt()
        }
    }
    let row = -1
    let hasOdd
    for (let i = 0; i < m; i++) {
        hasOdd = false
        for (let j = 0; j < n; j++) {
            if (matr[i][j] % 2 != 0) {
                hasOdd = true
                break
            }
        }
        if (!hasOdd) {
            row = i
        }
    }
    row++
    return compareInt(row)
}

// { "no": 35, "dat": "", "ans": "" }
const matrix35 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextInt()
        }
    }
    let col = -1
    let hasEven
    for (let j = 0; j < n; j++) {
        hasEven = false
        for (let i = 0; i < m; i++) {
            if (matr[i][j] % 2 == 0) {
                hasEven = true
                break
            }
        }
        if (!hasEven) {
            col = j
            break
        }
    }
    col++
    return compareInt(col)
}

// { "no": 36, "dat": "", "ans": "" }
const matrix36 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextInt()
        }
    }
    let first = new Array(101)
    let other = new Array(101)
    for (let i = 0; i <= 100; i++) {
        first[i] = other[i] = false
    }
    for (let j = 0; j < n; j++) {
        first[matr[0][j]] = true
    }
    let same, rowsCount = 0
    for (let i = 1; i < m; i++) {
        for (let k = 0; k <= 100; k++) {
            other[k] = false
        }
        for (let j = 0; j < n; j++) {
            other[matr[i][j]] = true
        }
        same = true
        for (let k = 0; k <= 100; k++) {
            if (first[k] != other[k]) {
                same = false
                break
            }
        }
        rowsCount += +same
    }
    return compareInt(rowsCount)
}

// { "no": 37, "dat": "", "ans": "" }
const matrix37 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextInt()
        }
    }
    let last = new Array(101)
    let other = new Array(101)
    for (let k = 0; k < 101; k++) {
        last[k] = false
    }
    for (let i = 0; i < m; i++) {
        last[matr[i][n - 1]] = true
    }
    let colsCount = 0
    for (let j = n - 2; j >= 0; j--) {
        for (let k = 0; k < 101; k++) {
            other[k] = false
        }
        for (let i = 0; i < m; i++) {
            other[matr[i][j]] = true
        }
        let same = true
        for (let k = 0; k < 101; k++) {
            if (last[k] != other[k]) {
                same = false
                break
            }
        }
        colsCount += +same
    }
    return compareInt(colsCount)
}

// { "no": 38, "dat": "", "ans": "" }
const matrix38 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextInt()
        }
    }
    let rowsCount = 0
    let arr = new Array(101)
    for (let i = 0; i < m; i++) {
        for (let k = 0; k < 101; k++) {
            arr[k] = false
        }
        let different = true
        for (let j = 0; j < n; j++) {
            if (arr[matr[i][j]]) {
                different = false
                break
            }
            arr[matr[i][j]] = true
        }
        if (different) {
            rowsCount++
        }
    }
    return compareInt(rowsCount)
}

// { "no": 39, "dat": "", "ans": "" }
const matrix39 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextInt()
        }
    }
    let colsCount = 0
    let arr = new Array(101)
    for (let j = 0; j < n; j++) {
        for (let k = 0; k < 101; k++) {
            arr[k] = false
        }
        let different = true
        for (let i = 0; i < m; i++) {
            if (arr[matr[i][j]]) {
                different = false
                break
            }
            arr[matr[i][j]] = true
        }
        if (different) {
            colsCount++
        }
    }
    return compareInt(colsCount)
}

// { "no": 40, "dat": "", "ans": "" }
const matrix40 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextInt()
        }
    }
    let arr = new Array(101)
    let MAXIMAL = 0, rowIndex = -1
    for (let i = 0; i < m; i++) {
        for (let k = 0; k < 101; k++) {
            arr[k] = 0
        }
        for (let j = 0; j < n; j++) {
            arr[matr[i][j]]++
        }
        let maximal = arr[0]
        for (let k = 1; k < 101; k++) {
            if (arr[k] > maximal) {
                maximal = arr[k]
            }
        }
        if (rowIndex < 0 || maximal >= MAXIMAL) {
            rowIndex = i
            MAXIMAL = maximal
        }
    }
    rowIndex++
    return compareInt(rowIndex)
}

// { "no": 41, "dat": "", "ans": "" }
const matrix41 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextInt()
        }
    }
    let colIndex = -1
    let arr = new Array(101)
    let MAXIMAL = 0
    for (let j = 0; j < n; j++) {
        for (let k = 0; k < 101; k++) {
            arr[k] = 0
        }
        for (let i = 0; i < m; i++) {
            arr[matr[i][j]]++
        }
        let maximal = arr[0]
        for (let k = 1; k < 101; k++) {
            if (arr[k] > maximal) {
                maximal = arr[k]
            }
        }
        if (colIndex < 0 || maximal > MAXIMAL) {
            colIndex = j
            MAXIMAL = maximal
        }
    }
    colIndex++
    return compareInt(colIndex)
}

// { "no": 42, "dat": "2", "ans": "" }
const matrix42 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
        }
    }
    let sortedRowsCount = 0
    for (let i = 0; i < m; i++) {
        let progress = true
        for (let j = 1; j < n; j++) {
            if (matr[i][j - 1] > matr[i][j]) {
                progress = false
                break
            }
        }
        sortedRowsCount += +progress
    }
    return compareInt(sortedRowsCount)
}

// { "no": 43, "dat": "2", "ans": "" }
const matrix43 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
        }
    }
    let sortedColsCount = 0
    for (let j = 0; j < n; j++) {
        let regress = true
        for (let i = 1; i < m; i++) {
            if (matr[i - 1][j] < matr[i][j]) {
                regress = false
                break
            }
        }
        sortedColsCount += +regress
    }
    return compareInt(sortedColsCount)
}

// { "no": 44, "dat": "2", "ans": "2" }
const matrix44 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
        }
    }
    let MINIMAL = 0
    for (let i = 0; i < m; i++) {
        let progress = true, regress = true
        let minimal = matr[i][0]
        for (let j = 1; j < n; j++) {
            if (matr[i][j] < minimal) {
                minimal = matr[i][j]
            }
            if (matr[i][j - 1] > matr[i][j]) {
                progress = false
            }
            if (matr[i][j - 1] < matr[i][j]) {
                regress = false
            }
        }
        if (progress || regress) {
            if (MINIMAL == 0 || minimal < MINIMAL) {
                MINIMAL = minimal
            }
        }
    }
    return compareFloat(2, MINIMAL)
}

// { "no": 45, "dat": "2", "ans": "2" }
const matrix45 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
        }
    }
    let MAXIMAL = 0
    for (let j = 0; j < n; j++) {
        let progress = true, regress = true
        let maximal = matr[0][j]
        for (let i = 1; i < m; i++) {
            if (matr[i][j] > maximal) {
                maximal = matr[i][j]
            }
            if (matr[i - 1][j] < matr[i][j]) {
                regress = false
            }
            if (matr[i - 1][j] > matr[i][j]) {
                progress = false
            }
        }
        if (progress || regress) {
            if (MAXIMAL == 0 || maximal > MAXIMAL) {
                MAXIMAL = maximal
            }
        }
    }
    return compareFloat(2, MAXIMAL)
}

// { "no": 46, "dat": "", "ans": "" }
const matrix46 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextInt()
        }
    }
    let maximal = new Array(m)
    for (let i = 0; i < m; i++) {
        maximal[i] = matr[i][0]
        for (let j = 1; j < n; j++) {
            if (matr[i][j] > maximal[i]) {
                maximal[i] = matr[i][j]
            }
        }
    }
    let minimal = new Array(n)
    for (let j = 0; j < n; j++) {
        minimal[j] = matr[0][j]
        for (let i = 1; i < m; i++) {
            if (matr[i][j] < minimal[j]) {
                minimal[j] = matr[i][j]
            }
        }
    }
    let minMax = 0
    for (let i = 0; i < m; i++) {
        for (let j = 0; j < n; j++) {
            if ((matr[i][j] == maximal[i]) && (matr[i][j] == minimal[j])) {
                minMax = matr[i][j]
            }
        }
    }
    return compareInt(minMax)
}

// { "no": 47, "dat": "2", "ans": "2" }
const matrix47 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
        }
    }
    let k1 = nextInt()
    let k2 = nextInt()
    for (let j = 0; j < n; j++) {
        matr[k1 - 1][j] += matr[k2 - 1][j]
        matr[k2 - 1][j] = matr[k1 - 1][j] - matr[k2 - 1][j]
        matr[k1 - 1][j] -= matr[k2 - 1][j]
    }
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...matr[i])) {
            return false
        }
    }
    return true
}

// { "no": 48, "dat": "2", "ans": "2" }
const matrix48 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
        }
    }
    let k1 = nextInt()
    let k2 = nextInt()
    for (let i = 0; i < m; i++) {
        matr[i][k1 - 1] += matr[i][k2 - 1]
        matr[i][k2 - 1] = matr[i][k1 - 1] - matr[i][k2 - 1]
        matr[i][k1 - 1] -= matr[i][k2 - 1]
    }
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...matr[i])) {
            return false
        }
    }
    return true
}

// { "no": 49, "dat": "2", "ans": "2" }
const matrix49 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        let minIndex = -1, maxIndex = -1
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
            if (minIndex < 0 || matr[i][j] <= matr[i][minIndex]) {
                minIndex = j
            }
            if (maxIndex < 0 || matr[i][j] >= matr[i][maxIndex]) {
                maxIndex = j
            }
        }
        matr[i][minIndex] += matr[i][maxIndex]
        matr[i][maxIndex] = matr[i][minIndex] - matr[i][maxIndex]
        matr[i][minIndex] -= matr[i][maxIndex]
    }
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...matr[i])) {
            return false
        }
    }
    return true
}

// { "no": 50, "dat": "2", "ans": "2" }
const matrix50 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
        }
    }
    for (let j = 0; j < n; j++) {
        let minIndex = 0, maxIndex = 0
        for (let i = 1; i < m; i++) {
            if (matr[i][j] > matr[maxIndex][j]) {
                maxIndex = i
            }
            if (matr[i][j] < matr[minIndex][j]) {
                minIndex = i
            }
        }
        if (minIndex == maxIndex) {
            continue
        }
        matr[minIndex][j] += matr[maxIndex][j]
        matr[maxIndex][j] = matr[minIndex][j] - matr[maxIndex][j]
        matr[minIndex][j] -= matr[maxIndex][j]
    }
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...matr[i])) {
            return false
        }
    }
    return true
}

// { "no": 51, "dat": "2", "ans": "2" }
const matrix51 = () => {
    let m = nextInt()
    let n = nextFloat()
    let matr = new Array(m)
    let minRowIndex = -1, maxRowIndex = -1
    let minimal = 0, maximal = 0
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
            if (i == 0 && j == 0) {
                minRowIndex = maxRowIndex = i
                minimal = maximal = matr[i][j]
            } else if (matr[i][j] <= minimal) {
                minRowIndex = i
                minimal = matr[i][j]
            } else if (matr[i][j] >= maximal) {
                maxRowIndex = i
                maximal = matr[i][j]
            }
        }
    }
    if (minRowIndex != maxRowIndex && minRowIndex >= 0 && maxRowIndex >= 0) {
        for (let j = 0; j < n; j++) {
            matr[minRowIndex][j] += matr[maxRowIndex][j]
            matr[maxRowIndex][j] = matr[minRowIndex][j] - matr[maxRowIndex][j]
            matr[minRowIndex][j] -= matr[maxRowIndex][j]
        }
    }
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...matr[i])) {
            return false
        }
    }
    return true
}

// { "no": 52, "dat": "2", "ans": "2" }
const matrix52 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
        }
    }
    let minColIndex = -1, maxColIndex = -1
    let minimal = 0, maximal = 0
    for (let j = 0; j < n; j++) {
        for (let i = 0; i < m; i++) {
            if (minColIndex < 0 || matr[i][j] < minimal) {
                minColIndex = j
                minimal = matr[i][j]
            }
            if (maxColIndex < 0 || matr[i][j] > maximal) {
                maxColIndex = j
                maximal = matr[i][j]
            }
        }
    }
    if (minColIndex != maxColIndex && minColIndex >= 0 && maxColIndex >= 0) {
        for (let i = 0; i < m; i++) {
            matr[i][minColIndex] += matr[i][maxColIndex]
            matr[i][maxColIndex] = matr[i][minColIndex] - matr[i][maxColIndex]
            matr[i][minColIndex] -= matr[i][maxColIndex]
        }
    }
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...matr[i])) {
            return false
        }
    }
    return true
}

// { "no": 53, "dat": "2", "ans": "2" }
const matrix53 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
        }
    }
    let positiveCol = 0
    for (let j = 1; j < n; j++) {
        let hasNegative = false
        for (let i = 0; i < m; i++) {
            if (matr[i][j] < 0) {
                hasNegative = true
                break
            }
        }
        if (!hasNegative) {
            positiveCol = j
        }
    }
    if (positiveCol > 0) {
        for (let i = 0; i < m; i++) {
            matr[i][0] += matr[i][positiveCol]
            matr[i][positiveCol] = matr[i][0] - matr[i][positiveCol]
            matr[i][0] -= matr[i][positiveCol]
        }
    }
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...matr[i])) {
            return false
        }
    }
    return true
}

// { "no": 54, "dat": "2", "ans": "2" }
const matrix54 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
        }
    }
    let positiveCol = n - 1
    for (let j = 0; j < n - 1; j++) {
        let hasPositive = false
        for (let i = 0; i < m; i++) {
            if (matr[i][j] > 0) {
                hasPositive = true
                break
            }
        }
        if (!hasPositive) {
            positiveCol = j
            break
        }
    }
    if (positiveCol < n - 1) {
        for (let i = 0; i < m; i++) {
            matr[i][n - 1] += matr[i][positiveCol]
            matr[i][positiveCol] = matr[i][n - 1] - matr[i][positiveCol]
            matr[i][n - 1] -= matr[i][positiveCol]
        }
    }
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...matr[i])) {
            return false
        }
    }
    return true
}

// { "no": 55, "dat": "2", "ans": "2" }
const matrix55 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
        }
    }
    for (let a = 0, b = parseInt(m / 2); b < m; a++, b++) {
        for (let j = 0; j < n; j++) {
            matr[a][j] += matr[b][j]
            matr[b][j] = matr[a][j] - matr[b][j]
            matr[a][j] -= matr[b][j]
        }
    }
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...matr[i])) {
            return false
        }
    }
    return true
}

// { "no": 56, "dat": "2", "ans": "2" }
const matrix56 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
        }
    }
    for (let a = 0, b = parseInt(n / 2); b < n; a++, b++) {
        for (let i = 0; i < m; i++) {
            matr[i][a] += matr[i][b]
            matr[i][b] = matr[i][a] - matr[i][b]
            matr[i][a] -= matr[i][b]
        }
    }
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...matr[i])) {
            return false
        }
    }
    return true
}

// { "no": 57, "dat": "2", "ans": "2" }
const matrix57 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
        }
    }
    for (let i = 0, k = parseInt(m / 2); k < m; i++, k++) {
        for (let j = 0, l = parseInt(n / 2); l < n; j++, l++) {
            matr[i][j] += matr[k][l]
            matr[k][l] = matr[i][j] - matr[k][l]
            matr[i][j] -= matr[k][l]
        }
    }
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...matr[i])) {
            return false
        }
    }
    return true
}

// { "no": 58, "dat": "2", "ans": "2" }
const matrix58 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
        }
    }
    for (let i = parseInt(m / 2), k = 0; i < m; i++, k++) {
        for (let j = 0, l = parseInt(n / 2); l < n; j++, l++) {
            matr[i][j] += matr[k][l]
            matr[k][l] = matr[i][j] - matr[k][l]
            matr[i][j] -= matr[k][l]
        }
    }
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...matr[i])) {
            return false
        }
    }
    return true
}

// { "no": 59, "dat": "2", "ans": "2" }
const matrix59 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
        }
    }
    for (let start = 0, finish = m - 1; start < finish; start++, finish--) {
        for (let j = 0; j < n; j++) {
            matr[start][j] += matr[finish][j]
            matr[finish][j] = matr[start][j] - matr[finish][j]
            matr[start][j] -= matr[finish][j]
        }
    }
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...matr[i])) {
            return false
        }
    }
    return true
}

// { "no": 60, "dat": "2", "ans": "2" }
const matrix60 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
        }
    }
    for (let start = 0, finish = n - 1; start < finish; start++, finish--) {
        for (let i = 0; i < m; i++) {
            matr[i][start] += matr[i][finish]
            matr[i][finish] = matr[i][start] - matr[i][finish]
            matr[i][start] -= matr[i][finish]
        }
    }
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...matr[i])) {
            return false
        }
    }
    return true
}

// { "no": 61, "dat": "2", "ans": "2" }
const matrix61 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
        }
    }
    let k = nextInt()
    matr[k - 1].splice(0, matr[k - 1].length)
    matr.splice(k - 1, 1)
    m--
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...matr[i])) {
            return false
        }
    }
    return true
}

// { "no": 62, "dat": "2", "ans": "2" }
const matrix62 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
        }
    }
    let k = nextInt()
    for (let i = 0; i < m; i++) {
        matr[i].splice(k - 1, 1)
    }
    n--
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...matr[i])) {
            return false
        }
    }
    return true
}

// { "no": 63, "dat": "2", "ans": "2" }
const matrix63 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    let minimal = 0
    let minimalRow = -1
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
            if (minimalRow < 0 || matr[i][j] <= minimal) {
                minimal = matr[i][j]
                minimalRow = i
            }
        }
    }
    matr[minimalRow].splice(0, matr[minimalRow].length)
    matr.splice(minimalRow, 1)
    m--
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...matr[i])) {
            return false
        }
    }
    return true
}

// { "no": 64, "dat": "2", "ans": "2" }
const matrix64 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
        }
    }
    let maximal = 0
    let maximalColIndex = -1
    for (let j = 0; j < n; j++) {
        for (let i = 0; i < m; i++) {
            if (maximalColIndex < 0 || matr[i][j] >= maximal) {
                maximalColIndex = j
                maximal = matr[i][j]
            }
        }
    }
    for (let i = 0; i < m; i++) {
        matr[i].splice(maximalColIndex, 1)
    }
    n--
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...matr[i])) {
            return false
        }
    }
    return true
}

// { "no": 65, "dat": "2", "ans": "2" }
const matrix65 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
        }
    }
    let positiveCol = -1
    for (let j = 0; j < n; j++) {
        let isPositiveCol = true
        for (let i = 0; i < m; i++) {
            if (matr[i][j] < 0) {
                isPositiveCol = false
                break
            }
        }
        if (isPositiveCol) {
            positiveCol = j
            break
        }
    }
    if (positiveCol >= 0) {
        for (let i = 0; i < m; i++) {
            matr[i].splice(positiveCol, 1)
        }
        n--
    }
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...matr[i])) {
            return false
        }
    }
    return true
}

// { "no": 66, "dat": "2", "ans": "2" }
const matrix66 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
        }
    }
    let negativeCol = -1
    for (let j = 0; j < n; j++) {
        let isNegativeCol = true
        for (let i = 0; i < m; i++) {
            if (matr[i][j] >= 0) {
                isNegativeCol = false
                break
            }
        }
        if (isNegativeCol) {
            negativeCol = j
        }
    }
    if (negativeCol >= 0) {
        for (let i = 0; i < m; i++) {
            matr[i].splice(negativeCol, 1)
        }
        n--
    }
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...matr[i])) {
            return false
        }
    }
    return true
}

// { "no": 67, "dat": "2", "ans": "2" }
const matrix67 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
        }
    }
    for (let j = 0; j < n; j++) {
        let isPositiveCol = true
        for (let i = 0; i < m; i++) {
            if (matr[i][j] <= 0) {
                isPositiveCol = false
                break
            }
        }
        if (isPositiveCol) {
            for (let i = 0; i < m; i++) {
                matr[i].splice(j, 1)
            }
            n--
            j--
        }
    }
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...matr[i])) {
            return false
        }
    }
    return true
}

// { "no": 68, "dat": "2", "ans": "2" }
const matrix68 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
        }
    }
    let k = nextInt()
    matr.splice(k - 1, 0, new Array(n))
    for (let j = 0; j < n; j++) {
        matr[k - 1][j] = 0
    }
    m++
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...matr[i])) {
            return false
        }
    }
    return true
}

// { "no": 69, "dat": "2", "ans": "2" }
const matrix69 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
        }
    }
    let k = nextInt()
    for (let i = 0; i < m; i++) {
        if (k == n) {
            matr[i].push(1)
        } else {
            matr[i].splice(k, 0, 1)
        }
    }
    n++
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...matr[i])) {
            return false
        }
    }
    return true
}

// { "no": 70, "dat": "2", "ans": "2" }
const matrix70 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
        }
    }
    let maximal = matr[0][0]
    let maximalRow = 0
    for (let i = 0; i < m; i++) {
        for (let j = 0; j < n; j++) {
            if (matr[i][j] > maximal) {
                maximal = matr[i][j]
                maximalRow = i
            }
        }
    }
    matr.splice(maximalRow, 0, matr[maximalRow])
    m++
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...matr[i])) {
            return false
        }
    }
    return true
}

// { "no": 71, "dat": "2", "ans": "2" }
const matrix71 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
        }
    }
    let minimal = matr[0][0]
    let minimalCol = 0
    for (let j = 0; j < n; j++) {
        for (let i = 0; i < m; i++) {
            if (matr[i][j] < minimal) {
                minimal = matr[i][j]
                minimalCol = j
            }
        }
    }
    for (let i = 0; i < m; i++) {
        matr[i].splice(minimalCol, 0, matr[i][minimalCol])
    }
    n++
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...matr[i])) {
            return false
        }
    }
    return true
}

// { "no": 72, "dat": "2", "ans": "2" }
const matrix72 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
        }
    }
    for (let j = 0; j < n; j++) {
        let isPositiveCol = true
        for (let i = 0; i < m; i++) {
            if (matr[i][j] <= 0) {
                isPositiveCol = false
                break
            }
        }
        if (isPositiveCol) {
            for (let i = 0; i < m; i++) {
                matr[i].splice(j, 0, 1)
            }
            n++
            break
        }
    }
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...matr[i])) {
            return false
        }
    }
    return true
}

// { "no": 73, "dat": "2", "ans": "2" }
const matrix73 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
        }
    }
    let negativeCol = -1
    for (let j = 0; j < n; j++) {
        let isNegativeCol = true
        for (let i = 0; i < m; i++) {
            if (matr[i][j] >= 0) {
                isNegativeCol = false
                break
            }
        }
        if (isNegativeCol) {
            negativeCol = j
        }
    }
    if (negativeCol >= 0) {
        for (let i = 0; i < m; i++) {
            if (negativeCol == n - 1) {
                matr[i].push(0)
            } else {
                matr[i].splice(negativeCol + 1, 0, 0)
            }
        }
        n++
    }
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...matr[i])) {
            return false
        }
    }
    return true
}

// { "no": 74, "dat": "2", "ans": "2" }
const matrix74 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    let temp = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        temp[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
            temp[i][j] = matr[i][j]
        }
    }
    for (let i = 0; i < m; i++) {
        for (let j = 0; j < n; j++) {
            if (isLocalMinimum(temp, m, n, i, j)) {
                matr[i][j] = 0
            }
        }
    }
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...matr[i])) {
            return false
        }
    }
    return true
}

// { "no": 75, "dat": "2", "ans": "2" }
const matrix75 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    let temp = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        temp[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
            temp[i][j] = matr[i][j]
        }
    }
    for (let i = 0; i < m; i++) {
        for (let j = 0; j < n; j++) {
            if (isLocalMaximum(temp, m, n, i, j)) {
                matr[i][j] *= -1
            }
        }
    }
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...matr[i])) {
            return false
        }
    }
    return true
}

const isLocalMinimum = (matrix, rowsCount, colsCount, rowIndex, colIndex) => {
    // check with left
    if (rowIndex > 0 && !(matrix[rowIndex][colIndex] < matrix[rowIndex - 1][colIndex])) {
        return
    }
    // check with right
    if (rowIndex < rowsCount - 1 && !(matrix[rowIndex][colIndex] < matrix[rowIndex + 1][colIndex])) {
        return
    }
    // check with top
    if (colIndex > 0 && !(matrix[rowIndex][colIndex] < matrix[rowIndex][colIndex - 1])) {
        return
    }
    // check with down
    if (colIndex < colsCount - 1 && !(matrix[rowIndex][colIndex] < matrix[rowIndex][colIndex + 1])) {
        return
    }
    // check with top + left
    if (rowIndex > 0 && colIndex > 0 && !(matrix[rowIndex][colIndex] < matrix[rowIndex - 1][colIndex - 1])) {
        return
    }
    // check with top + right
    if (rowIndex > 0 && colIndex < colsCount - 1 && !(matrix[rowIndex][colIndex] < matrix[rowIndex - 1][colIndex + 1])) {
        return
    }
    // check with down + left
    if (rowIndex < rowsCount - 1 && colIndex > 0 && !(matrix[rowIndex][colIndex] < matrix[rowIndex + 1][colIndex - 1])) {
        return
    }
    // check with down + right
    if (rowIndex < rowsCount - 1 && colIndex < colsCount - 1 && !(matrix[rowIndex][colIndex] < matrix[rowIndex + 1][colIndex + 1])) {
        return
    }
    return true
}

const isLocalMaximum = (matrix, rowsCount, colsCount, rowIndex, colIndex) => {
    // check with top
    if (rowIndex > 0 && !(matrix[rowIndex][colIndex] > matrix[rowIndex - 1][colIndex])) {
        return false
    }
    // check with bottom
    if (rowIndex < rowsCount - 1 && !(matrix[rowIndex][colIndex] > matrix[rowIndex + 1][colIndex])) {
        return false
    }
    // check with left
    if (colIndex > 0 && !(matrix[rowIndex][colIndex] > matrix[rowIndex][colIndex - 1])) {
        return false
    }
    // check with right
    if (colIndex < colsCount - 1 && !(matrix[rowIndex][colIndex] > matrix[rowIndex][colIndex + 1])) {
        return false
    }
    // check with top + left
    if (rowIndex > 0 && colIndex > 0 && !(matrix[rowIndex][colIndex] > matrix[rowIndex - 1][colIndex - 1])) {
        return false
    }
    // check with top + right
    if (rowIndex > 0 && colIndex < colsCount - 1 && !(matrix[rowIndex][colIndex] > matrix[rowIndex - 1][colIndex + 1])) {
        return false
    }
    // check with down + left
    if (rowIndex < rowsCount - 1 && colIndex > 0 && !(matrix[rowIndex][colIndex] > matrix[rowIndex + 1][colIndex - 1])) {
        return false
    }
    // check with down + right
    if (rowIndex < rowsCount - 1 && colIndex < colsCount - 1 && !(matrix[rowIndex][colIndex] > matrix[rowIndex + 1][colIndex + 1])) {
        return false
    }
    return true
}

// { "no": 76, "dat": "2", "ans": "2" }
const matrix76 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
        }
    }
    for (let k = 0; k < m - 1; k++) {
        for (let i = 1; i < m - k; i++) {
            if (matr[i][0] < matr[i - 1][0]) {
                let temp = matr[i - 1]
                matr[i - 1] = matr[i]
                matr[i] = temp
            }
        }
    }
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...matr[i])) {
            return false
        }
    }
    return true
}

// { "no": 77, "dat": "2", "ans": "2" }
const matrix77 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
        }
    }
    for (let k = 0; k < n - 1; k++) {
        for (let j = 1; j < n - k; j++) {
            if (matr[m - 1][j] > matr[m - 1][j - 1]) {
                for (let i = 0; i < m; i++) {
                    matr[i][j] += matr[i][j - 1]
                    matr[i][j - 1] = matr[i][j] - matr[i][j - 1]
                    matr[i][j] -= matr[i][j - 1]
                }
            }
        }
    }
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...matr[i])) {
            return false
        }
    }
    return true
}

// { "no": 78, "dat": "2", "ans": "2" }
const matrix78 = () => {
    let m = nextInt()
    let n = nextInt()
    let minimal = new Array(m)
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
            if (j == 0 || matr[i][j] < minimal[i]) {
                minimal[i] = matr[i][j]
            }
        }
    }
    for (let j = 0; j < m - 1; j++) {
        for (let i = 1; i < m - j; i++) {
            if (minimal[i - 1] < minimal[i]) {
                let temp = matr[i - 1]
                matr[i - 1] = matr[i]
                matr[i] = temp

                minimal[i - 1] += minimal[i]
                minimal[i] = minimal[i - 1] - minimal[i]
                minimal[i - 1] -= minimal[i]
            }
        }
    }
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...matr[i])) {
            return false
        }
    }
    return true
}

// { "no": 79, "dat": "2", "ans": "2" }
const matrix79 = () => {
    let m = nextInt()
    let n = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(n)
        for (let j = 0; j < n; j++) {
            matr[i][j] = nextFloat()
        }
    }
    let maximal = new Array(n)
    for (let j = 0; j < n; j++) {
        maximal[j] = matr[0][j]
        for (let i = 1; i < m; i++) {
            if (matr[i][j] > maximal[j]) {
                maximal[j] = matr[i][j]
            }
        }
    }
    for (let k = 0; k < n - 1; k++) {
        for (let j = 1; j < n - k; j++) {
            if (maximal[j - 1] > maximal[j]) {
                for (let i = 0; i < m; i++) {
                    matr[i][j] += matr[i][j - 1]
                    matr[i][j - 1] = matr[i][j] - matr[i][j - 1]
                    matr[i][j] -= matr[i][j - 1]
                }
                maximal[j] += maximal[j - 1]
                maximal[j - 1] = maximal[j] - maximal[j - 1]
                maximal[j] -= maximal[j - 1]
            }
        }
    }
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...matr[i])) {
            return false
        }
    }
    return true
}

// A1,1,    A2,2,    A3,3,    …,    AM,M.
// { "no": 80, "dat": "2", "ans": "2" }
const matrix80 = () => {
    let m = nextInt()
    let row, col
    let a = new Array(m)
    for (let i = 0; i < m; i++) {
        a[i] = new Array(m)
        for (let j = 0; j < m; j++) {
            a[i][j] = nextFloat()
        }
    }
    let sum = 0
    for (row = 0, col = 0; row < m; row++, col++) {
        sum += a[row][col]
    }
    return compareFloat(2, sum)
}

// A1,M,    A2,M−1,    A3,M−2,    …,    AM,1.
// { "no": 81, "dat": "2", "ans": "2" }
const matrix81 = () => {
    let m = nextInt()
    let row, col
    let a = new Array(m)
    for (let i = 0; i < m; i++) {
        a[i] = new Array(m)
        for (let j = 0; j < m; j++) {
            a[i][j] = nextFloat()
        }
    }
    let sum = 0
    let count = 0
    for (row = 0, col = m - 1; row < m; row++, col--) {
        sum += a[row][col]
        count++
    }
    let aMean = sum / count
    return compareFloat(2, aMean)
}

// { "no": 82, "dat": "2", "ans": "2" }
const matrix82 = () => {
    let m = nextInt()
    let row, col
    let a = new Array(m)
    for (let i = 0; i < m; i++) {
        a[i] = new Array(m)
        for (let j = 0; j < m; j++) {
            a[i][j] = nextFloat()
        }
    }
    for (col = m - 1; col >= 0; col--) {
        let sum = 0
        for (let i = 0, j = col; j < m; i++, j++) {
            sum += a[i][j]
        }
        if (!compareFloat(2, sum)) {
            return false
        }
    }
    for (row = 1; row < m; row++) {
        let sum = 0
        for (let i = row, j = 0; i < m; i++, j++) {
            sum += a[i][j]
        }
        if (!compareFloat(2, sum)) {
            return false
        }
    }
    return true
}

// { "no": 83, "dat": "2", "ans": "2" }
const matrix83 = () => {
    let m = nextInt()
    let row, col
    let a = new Array(m)
    for (let i = 0; i < m; i++) {
        a[i] = new Array(m)
        for (let j = 0; j < m; j++) {
            a[i][j] = nextFloat()
        }
    }
    for (col = 0; col < m; col++) {
        let sum = 0
        for (let i = 0, j = col; j >= 0; i++, j--) {
            sum += a[i][j]
        }
        if (!compareFloat(2, sum)) {
            return false
        }
    }
    for (row = 1; row < m; row++) {
        let sum = 0
        for (let i = row, j = m - 1; i < m; i++, j--) {
            sum += a[i][j]
        }
        if (!compareFloat(2, sum)) {
            return false
        }
    }
    return true
}

// { "no": 84, "dat": "2", "ans": "2" }
const matrix84 = () => {
    let m = nextInt()
    let row, col;
    let a = new Array(m)
    for (let i = 0; i < m; i++) {
        a[i] = new Array(m)
        for (let j = 0; j < m; j++) {
            a[i][j] = nextFloat()
        }
    }
    for (col = m - 1; col >= 0; col--) {
        let sum = 0
        let count = 0
        for (let i = 0, j = col; j < m; i++, j++) {
            sum += a[i][j]
            count++
        }
        let amean = sum / count
        if (!compareFloat(2, amean)) {
            return false
        }
    }
    for (row = 1; row < m; row++) {
        let sum = 0
        let count = 0
        for (let i = row, j = 0; i < m; i++, j++) {
            sum += a[i][j]
            count++
        }
        let amean = sum / count
        if (!compareFloat(2, amean)) {
            return false
        }
    }
    return true
}

// { "no": 85, "dat": "2", "ans": "2" }
const matrix85 = () => {
    let m = nextInt()
    let row, col
    let a = new Array(m)
    for (let i = 0; i < m; i++) {
        a[i] = new Array(m)
        for (let j = 0; j < m; j++) {
            a[i][j] = nextFloat()
        }
    }
    for (col = 0; col < m; col++) {
        let sum = 0
        let count = 0
        for (let i = 0, j = col; j >= 0; i++, j--) {
            sum += a[i][j]
            count++
        }
        let amean = sum / count
        if (!compareFloat(2, amean)) {
            return false
        }
    }
    for (row = 1; row < m; row++) {
        let sum = 0
        let count = 0
        for (let i = row, j = m - 1; i < m; i++, j--) {
            sum += a[i][j]
            count++
        }
        let amean = sum / count
        if (!compareFloat(2, amean)) {
            return false
        }
    }
    return true
}

// { "no": 86, "dat": "2", "ans": "2" }
const matrix86 = () => {
    let m = nextInt()
    let row, col
    let a = new Array(m)
    for (let i = 0; i < m; i++) {
        a[i] = new Array(m)
        for (let j = 0; j < m; j++) {
            a[i][j] = nextFloat()
        }
    }
    for (col = m - 1; col >= 0; col--) {
        let minimal = a[0][col]
        for (let i = 0, j = col; j < m; i++, j++) {
            if (a[i][j] < minimal) {
                minimal = a[i][j]
            }
        }
        if (!compareFloat(2, minimal)) {
            return false
        }
    }
    for (row = 1; row < m; row++) {
        let minimal = a[row][0]
        for (let i = row, j = 0; i < m; i++, j++) {
            if (a[i][j] < minimal) {
                minimal = a[i][j]
            }
        }
        if (!compareFloat(2, minimal)) {
            return false
        }
    }
    return true
}

// { "no": 87, "dat": "2", "ans": "2" }
const matrix87 = () => {
    let m = nextInt()
    let row, col
    let a = new Array(m)
    for (let i = 0; i < m; i++) {
        a[i] = new Array(m)
        for (let j = 0; j < m; j++) {
            a[i][j] = nextFloat()
        }
    }
    for (col = 0; col < m; col++) {
        let maximal = a[0][col]
        for (let i = 0, j = col; j >= 0; i++, j--) {
            if (a[i][j] > maximal) {
                maximal = a[i][j]
            }
        }
        if (!compareFloat(2, maximal)) {
            return false
        }
    }
    for (row = 1; row < m; row++) {
        let maximal = a[row][m - 1]
        for (let i = row, j = m - 1; i < m; i++, j--) {
            if (a[i][j] > maximal) {
                maximal = a[i][j]
            }
        }
        if (!compareFloat(2, maximal)) {
            return false
        }
    }
    return true
}

// { "no": 88, "dat": "2", "ans": "2" }
const matrix88 = () => {
    let m = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(m)
        for (let j = 0; j < m; j++) {
            matr[i][j] = nextFloat()
        }
    }
    for (let row = 1; row < m; row++) {
        for (let i = row, j = 0; i < m; i++, j++) {
            matr[i][j] = 0
        }
    }
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...matr[i])) {
            return false
        }
    }
    return true
}

// { "no": 89, "dat": "2", "ans": "2" }
const matrix89 = () => {
    let m = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(m)
        for (let j = 0; j < m; j++) {
            matr[i][j] = nextFloat()
        }
    }
    for (let col = m - 2; col >= 0; col--) {
        for (let i = 0, j = col; j >= 0; i++, j--) {
            matr[i][j] = 0
        }
    }
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...matr[i])) {
            return false
        }
    }
    return true
}

// { "no": 90, "dat": "2", "ans": "2" }
const matrix90 = () => {
    let m = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(m)
        for (let j = 0; j < m; j++) {
            matr[i][j] = nextFloat()
        }
    }
    for (let row = 0; row < m; row++) {
        for (let i = row, j = m - 1; i < m; i++, j--) {
            matr[i][j] = 0
        }
    }
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...matr[i])) {
            return false
        }
    }
    return true
}

// { "no": 91, "dat": "2", "ans": "2" }
const matrix91 = () => {
    let m = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(m)
        for (let j = 0; j < m; j++) {
            matr[i][j] = nextFloat()
        }
    }
    for (let col = 0; col < m; col++) {
        for (let i = 0, j = col; j < m; i++, j++) {
            matr[i][j] = 0
        }
    }
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...matr[i])) {
            return false
        }
    }
    return true
}

// { "no": 92, "dat": "2", "ans": "2" }
const matrix92 = () => {
    let m = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(m)
        for (let j = 0; j < m; j++) {
            matr[i][j] = nextFloat()
        }
    }
    for (let row = 0; row < parseInt(m / 2); row++) {
        for (let col = row + 1; col < m - row - 1; col++) {
            matr[row][col] = 0
        }
    }
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...matr[i])) {
            return false
        }
    }
    return true
}

// { "no": 93, "dat": "2", "ans": "2" }
const matrix93 = () => {
    let m = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(m)
        for (let j = 0; j < m; j++) {
            matr[i][j] = nextFloat()
        }
    }
    for (let col = m - 1; col > parseInt(m / 2); col--) {
        for (let row = m - col; row < col; row++) {
            matr[row][col] = 0
        }
    }
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...matr[i])) {
            return false
        }
    }
    return true
}

// { "no": 94, "dat": "2", "ans": "2" }
const matrix94 = () => {
    let m = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(m)
        for (let j = 0; j < m; j++) {
            matr[i][j] = nextFloat()
        }
    }
    for (let col = 0; col <= parseInt(m / 2); col++) {
        for (let row = col; row < m - col; row++) {
            matr[row][col] = 0
        }
    }
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...matr[i])) {
            return false
        }
    }
    return true
}

// { "no": 95, "dat": "2", "ans": "2" }
const matrix95 = () => {
    let m = nextInt()
    let matr = new Array(m)
    for (let i = 0; i < m; i++) {
        matr[i] = new Array(m)
        for (let j = 0; j < m; j++) {
            matr[i][j] = nextFloat()
        }
    }
    for (let row = m - 1; row >= parseInt(m / 2); row--) {
        for (let col = m - row - 1; col <= row; col++) {
            matr[row][col] = 0
        }
    }
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...matr[i])) {
            return false
        }
    }
    return true
}

// { "no": 96, "dat": "2", "ans": "2" }
const matrix96 = () => {
    let m = nextInt()
    let a = new Array(m)
    for (let i = 0; i < m; i++) {
        a[i] = new Array(m)
        for (let j = 0; j < m; j++) {
            a[i][j] = nextFloat()
        }
    }
    for (let col = 1; col < m; col++) {
        for (let i = 0, j = col; j < m; i++, j++) {
            a[i][j] += a[j][i]
            a[j][i] = a[i][j] - a[j][i]
            a[i][j] -= a[j][i]
        }
    }
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...a[i])) {
            return false
        }
    }
    return true
}

// { "no": 97, "dat": "2", "ans": "2" }
const matrix97 = () => {
    let m = nextInt()
    let a = new Array(m)
    for (let i = 0; i < m; i++) {
        a[i] = new Array(m)
        for (let j = 0; j < m; j++) {
            a[i][j] = nextFloat()
        }
    }
    for (let col = m - 2; col >= 0; col--) {
        for (let i = 0, j = col; j >= 0; i++, j--) {
            a[i][j] += a[m - 1 - j][m - 1 - i]
            a[m - 1 - j][m - 1 - i] = a[i][j] - a[m - 1 - j][m - 1 - i]
            a[i][j] -= a[m - 1 - j][m - 1 - i]
        }
    }
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...a[i])) {
            return false
        }
    }
    return true
}

// { "no": 98, "dat": "2", "ans": "2" }
const matrix98 = () => {
    let m = nextInt()
    let a = new Array(m)
    for (let i = 0; i < m; i++) {
        a[i] = new Array(m)
        for (let j = 0; j < m; j++) {
            a[i][j] = nextFloat()
        }
    }
    for (let i = parseInt(m / 2) + m % 2 - 1; i >= 0; i--) {
        let limit = i == m - 1 - i ? parseInt(m / 2) : m
        for (let j = 0; j < limit; j++) {
            a[i][j] += a[m - 1 - i][m - 1 - j]
            a[m - 1 - i][m - 1 - j] = a[i][j] - a[m - 1 - i][m - 1 - j]
            a[i][j] -= a[m - 1 - i][m - 1 - j]
        }
    }
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...a[i])) {
            return false
        }
    }
    return true
}

// { "no": 99, "dat": "2", "ans": "2" }
const matrix99 = () => {
    let m = nextInt()
    let a = new Array(m)
    for (let i = 0; i < m; i++) {
        a[i] = new Array(m)
        for (let j = 0; j < m; j++) {
            a[i][j] = nextFloat()
        }
    }
    for (let i = 0; i < parseInt(m / 2); i++) {
        for (let j = i; j < m - 1 - i; j++) {
            let temp = a[i][j]
            a[i][j] = a[j][m - 1 - i]
            a[j][m - 1 - i] = a[m - 1 - i][m - 1 - j]
            a[m - 1 - i][m - 1 - j] = a[m - 1 - j][i]
            a[m - 1 - j][i] = temp
        }
    }
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...a[i])) {
            return false
        }
    }
    return true
}

// { "no": 100, "dat": "2", "ans": "2" }
const matrix100 = () => {
    let m = nextInt()
    let a = new Array(m)
    for (let i = 0; i < m; i++) {
        a[i] = new Array(m)
        for (let j = 0; j < m; j++) {
            a[i][j] = nextFloat()
        }
    }
    for (let i = 0; i < parseInt(m / 2); i++) {
        for (let j = i; j < m - 1 - i; j++) {
            let temp = a[i][j]
            a[i][j] = a[m - 1 - j][i]
            a[m - 1 - j][i] = a[m - 1 - i][m - 1 - j]
            a[m - 1 - i][m - 1 - j] = a[j][m - 1 - i]
            a[j][m - 1 - i] = temp
        }
    }
    for (let i = 0; i < m; i++) {
        if (!compareFloat(2, ...a[i])) {
            return false
        }
    }
    return true
}