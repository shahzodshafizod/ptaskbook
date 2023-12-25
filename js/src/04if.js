import {
    nextFloat
    , compareFloat
    , nextInt
    , compareInt
    , compareLine
    , initScanner
    , initChecker
    , inputsEOF
    , outputsEOF
} from './io.js'

(async () => {
    for (let taskNo = 1; taskNo <= 30; taskNo++) {
        const taskNoStr = String(taskNo).padStart(3, 0)
        for (let testNo = 1; testNo <= 100; testNo++) {
            const testNoStr = String(testNo).padStart(3, 0)
            const filePath = `../data/kt-gov2/04if/If${taskNoStr}/${testNoStr}/${testNoStr}`

            await initScanner(filePath + '.dat')
            await initChecker(filePath + '.ans')

            let ok = true
            switch (taskNo) {
                case 1: ok = if1(); break
                case 2: ok = if2(); break
                case 3: ok = if3(); break
                case 4: ok = if4(); break
                case 5: ok = if5(); break
                case 6: ok = if6(); break
                case 7: ok = if7(); break
                case 8: ok = if8(); break
                case 9: ok = if9(); break
                case 10: ok = if10(); break
                case 11: ok = if11(); break
                case 12: ok = if12(); break
                case 13: ok = if13(); break
                case 14: ok = if14(); break
                case 15: ok = if15(); break
                case 16: ok = if16(); break
                case 17: ok = if17(); break
                case 18: ok = if18(); break
                case 19: ok = if19(); break
                case 20: ok = if20(); break
                case 21: ok = if21(); break
                case 22: ok = if22(); break
                case 23: ok = if23(); break
                case 24: ok = if24(); break
                case 25: ok = if25(); break
                case 26: ok = if26(); break
                case 27: ok = if27(); break
                case 28: ok = if28(); break
                case 29: ok = if29(); break
                case 30: ok = if30(); break
            }

            if (!ok) {
                console.log("If" + taskNoStr + " #" + testNoStr + " has NOT been tested!")
                return
            }

            if (!inputsEOF()) {
                console.log("If" + taskNoStr + " #" + testNoStr + " input NOT EOF!")
                return
            }

            if (!outputsEOF()) {
                console.log("If" + taskNoStr + " #" + testNoStr + " output NOT EOF!")
                return
            }
        }
        console.log("If" + taskNoStr + " has been tested!")
    }
    console.log("If has been tested!")
})()

// { "no": 1, "dat": "", "ans": "" }
const if1 = () => {
    let number = nextInt()
    if (number > 0) {
        number -= 8
    }
    return compareInt(number)
}

// { "no": 2, "dat": "", "ans": "" }
const if2 = () => {
    let number = nextInt()
    if (number > 0) {
        number -= 8
    } else {
        number += 6
    }
    return compareInt(number)
}

// { "no": 3, "dat": "", "ans": "" }
const if3 = () => {
    let number = nextInt()
    if (number > 0) {
        number -= 8
    } else if (number < 0) {
        number += 6
    } else {
        number = 10
    }
    return compareInt(number)
}

// { "no": 4, "dat": "", "ans": "" }
const if4 = () => {
    let a = nextInt()
    let b = nextInt()
    let c = nextInt()
    let positivesCount = 0
    if (a > 0) {
        positivesCount++
    }
    if (b > 0) {
        positivesCount++
    }
    if (c > 0) {
        positivesCount++
    }
    return compareInt(positivesCount)
}

// { "no": 5, "dat": "", "ans": "" }
const if5 = () => {
    let a = nextInt()
    let b = nextInt()
    let c = nextInt()
    let positivesCount = 0, negativesCount = 0
    if (a > 0) {
        positivesCount++
    } else if (a < 0) {
        negativesCount++
    }
    if (b > 0) {
        positivesCount++
    } else if (b < 0) {
        negativesCount++
    }
    if (c > 0) {
        positivesCount++
    } else if (c < 0) {
        negativesCount++
    }
    return compareInt(positivesCount, negativesCount)
}

// { "no": 6, "dat": "2", "ans": "2" }
const if6 = () => {
    let a = nextFloat()
    let b = nextFloat()
    let larger = a
    if (b > a) {
        larger = b
    }
    return compareFloat(2, larger)
}

// { "no": 7, "dat": "2", "ans": "" }
const if7 = () => {
    let a = nextFloat()
    let b = nextFloat()
    let smallerIndex = 1
    if (b < a) {
        smallerIndex = 2
    }
    return compareInt(smallerIndex)
}

// { "no": 8, "dat": "2", "ans": "2" }
const if8 = () => {
    let a = nextFloat()
    let b = nextFloat()
    let larger = a
    if (b > a) {
        larger = b
    }
    let smaller = a + b - larger
    return compareFloat(2, larger, smaller)
}

// { "no": 9, "dat": "2", "ans": "2" }
const if9 = () => {
    let a = nextFloat()
    let b = nextFloat()
    if (a > b) {
        a += b
        b = a - b
        a -= b
    }
    return compareFloat(2, a, b)
}

// { "no": 10, "dat": "", "ans": "" }
const if10 = () => {
    let a = nextInt()
    let b = nextInt()
    if (a != b) {
        b = a += b
    } else {
        a = b = 0
    }
    return compareInt(a, b)
}

// { "no": 11, "dat": "", "ans": "" }
const if11 = () => {
    let a = nextInt()
    let b = nextInt()
    if (a != b) {
        let larger = a
        if (b > a) {
            larger = b
        }
        a = b = larger
    } else {
        a = b = 0
    }
    return compareInt(a, b)
}

// { "no": 12, "dat": "2", "ans": "2" }
const if12 = () => {
    let a = nextFloat()
    let b = nextFloat()
    let c = nextFloat()
    let minimal
    if (a < b && a < c) {
        minimal = a
    } else if (b < c) {
        minimal = b
    } else {
        minimal = c
    }
    return compareFloat(2, minimal)
}

// { "no": 13, "dat": "2", "ans": "2" }
const if13 = () => {
    let a = nextFloat()
    let b = nextFloat()
    let c = nextFloat()
    let minimal, maximal
    if (a < b && a < c) {
        minimal = a
    } else if (b < c) {
        minimal = b
    } else {
        minimal = c
    }
    if (a > b && a > c) {
        maximal = a
    } else if (b > c) {
        maximal = b
    } else {
        maximal = c
    }
    let between = a + b + c - minimal - maximal
    return compareFloat(2, between)
}

// { "no": 14, "dat": "2", "ans": "2" }
const if14 = () => {
    let a = nextFloat()
    let b = nextFloat()
    let c = nextFloat()
    let minimal, maximal
    if (a < b && a < c) {
        minimal = a
    } else if (b < c) {
        minimal = b
    } else {
        minimal = c
    }
    if (a > b && a > c) {
        maximal = a
    } else if (b > c) {
        maximal = b
    } else {
        maximal = c
    }
    return compareFloat(2, minimal, maximal)
}

// { "no": 15, "dat": "2", "ans": "2" }
const if15 = () => {
    let a = nextFloat()
    let b = nextFloat()
    let c = nextFloat()
    let sum = a + b + c
    if (a < b && a < c) {
        sum -= a
    } else if (b < c) {
        sum -= b
    } else {
        sum -= c
    }
    return compareFloat(2, sum)
}

// { "no": 16, "dat": "2", "ans": "2" }
const if16 = () => {
    let a = nextFloat()
    let b = nextFloat()
    let c = nextFloat()
    let factor
    if (a < b && b < c) {
        factor = 2
    } else {
        factor = -1
    }
    a *= factor
    b *= factor
    c *= factor
    return compareFloat(2, a, b, c)
}

// { "no": 17, "dat": "2", "ans": "2" }
const if17 = () => {
    let a = nextFloat()
    let b = nextFloat()
    let c = nextFloat()
    let factor
    if (a < b && b < c || a > b && b > c) {
        factor = 2
    } else {
        factor = -1
    }
    a *= factor
    b *= factor
    c *= factor
    return compareFloat(2, a, b, c)
}

// { "no": 18, "dat": "", "ans": "" }
const if18 = () => {
    let a = nextInt()
    let b = nextInt()
    let c = nextInt()
    let differIndex
    if (a == b) {
        differIndex = 3
    } else if (a == c) {
        differIndex = 2
    } else {
        differIndex = 1
    }
    return compareInt(differIndex)
}

// { "no": 19, "dat": "", "ans": "" }
const if19 = () => {
    let a = nextInt()
    let b = nextInt()
    let c = nextInt()
    let d = nextInt()
    let differIndex
    if (a == b && b == c) {
        differIndex = 4
    } else if (d != c) {
        differIndex = 3
    } else if (d != b) {
        differIndex = 2
    } else {
        differIndex = 1
    }
    return compareInt(differIndex)
}

// { "no": 20, "dat": "2", "ans": "2" }
const if20 = () => {
    let a = nextFloat()
    let b = nextFloat()
    let c = nextFloat()
    let distance, point
    if (Math.abs(b - a) < Math.abs(c - a)) {
        point = b
    } else {
        point = c
    }
    distance = Math.abs(point - a)
    return compareFloat(2, point, distance)
}

// { "no": 21, "dat": "", "ans": "" }
const if21 = () => {
    let x = nextInt()
    let y = nextInt()
    let result
    if (x == 0 && y == 0) {
        result = 0
    } else if (y == 0) {
        result = 1
    } else if (x == 0) {
        result = 2
    } else {
        result = 3
    }
    return compareInt(result)
}

// { "no": 22, "dat": "2", "ans": "" }
const if22 = () => {
    let x = nextFloat()
    let y = nextFloat()
    let quarter
    if (x > 0) {
        if (y > 0) {
            quarter = 1
        } else {
            quarter = 4
        }
    } else if (y > 0) {
        quarter = 2
    } else {
        quarter = 3
    }
    return compareInt(quarter)
}

// { "no": 23, "dat": "", "ans": "" }
const if23 = () => {
    let x1 = nextInt()
    let y1 = nextInt()
    let x2 = nextInt()
    let y2 = nextInt()
    let x3 = nextInt()
    let y3 = nextInt()
    let x4, y4
    if (x1 == x2) {
        x4 = x3
    } else if (x1 == x3) {
        x4 = x2
    } else {
        x4 = x1
    }
    if (y1 == y2) {
        y4 = y3
    } else if (y1 == y3) {
        y4 = y2
    } else {
        y4 = y1
    }
    return compareInt(x4, y4)
}

// { "no": 24, "dat": "2", "ans": "2" }
const if24 = () => {
    let x = nextFloat()
    let fx
    if (x > 0) {
        fx = 2 * Math.sin(x)
    } else {
        fx = 6 - x
    }
    return compareFloat(2, fx)
}

// { "no": 25, "dat": "", "ans": "" }
const if25 = () => {
    let x = nextInt()
    let fx
    if (x < -2 || x > 2) {
        fx = 2 * x
    } else {
        fx = -3 * x
    }
    return compareInt(fx)
}

// { "no": 26, "dat": "2", "ans": "2" }
const if26 = () => {
    let x = nextFloat()
    let fx
    if (x <= 0) {
        fx = -x
    } else if (x >= 2) {
        fx = 4
    } else {
        fx = x * x
    }
    return compareFloat(2, fx)
}

// { "no": 27, "dat": "2", "ans": "" }
const if27 = () => {
    let x = nextFloat()
    let fx
    if (x < 0) {
        fx = 0
    } else if (parseInt(x) % 2 == 0) {
        fx = 1
    } else {
        fx = -1
    }
    return compareInt(fx)
}

// { "no": 28, "dat": "", "ans": "" }
const if28 = () => {
    let year = nextInt()
    let daysAmount
    if (year % 400 == 0) {
        daysAmount = 366
    } else if (year % 100 == 0) {
        daysAmount = 365
    } else if (year % 4 == 0) {
        daysAmount = 366
    } else {
        daysAmount = 365
    }
    return compareInt(daysAmount)
}

// { "no": 29, "dat": "", "ans": "" }
const if29 = () => {
    let number = nextInt()
    let description = ""
    if (number == 0) {
        description += "zero"
    } else {
        if (number > 0) {
            description += "positive"
        } else {
            description += "negative"
        }
        if (number % 2 == 0) {
            description += " even"
        } else {
            description += " odd"
        }
    }
    description += " number"
    return compareLine(description)
}

// { "no": 30, "dat": "", "ans": "" }
const if30 = () => {
    let number = nextInt()
    let description = ""
    if (number > 0 && number < 1000) {
        if (number < 10) {
            description += "one"
        } else if (number < 100) {
            description += "two"
        } else {
            description += "three"
        }
        description += "-digit "
        if (number % 2 == 0) {
            description += "even"
        } else {
            description += "odd"
        }
        description += " number"
    }
    return compareLine(description)
}