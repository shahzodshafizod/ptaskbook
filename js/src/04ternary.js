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
                case 1: ok = ternary1(); break
                case 2: ok = ternary2(); break
                case 3: ok = ternary3(); break
                case 4: ok = ternary4(); break
                case 5: ok = ternary5(); break
                case 6: ok = ternary6(); break
                case 7: ok = ternary7(); break
                case 8: ok = ternary8(); break
                case 9: ok = ternary9(); break
                case 10: ok = ternary10(); break
                case 11: ok = ternary11(); break
                case 12: ok = ternary12(); break
                case 13: ok = ternary13(); break
                case 14: ok = ternary14(); break
                case 15: ok = ternary15(); break
                case 16: ok = ternary16(); break
                case 17: ok = ternary17(); break
                case 18: ok = ternary18(); break
                case 19: ok = ternary19(); break
                case 20: ok = ternary20(); break
                case 21: ok = ternary21(); break
                case 22: ok = ternary22(); break
                case 23: ok = ternary23(); break
                case 24: ok = ternary24(); break
                case 25: ok = ternary25(); break
                case 26: ok = ternary26(); break
                case 27: ok = ternary27(); break
                case 28: ok = ternary28(); break
                case 29: ok = ternary29(); break
                case 30: ok = ternary30(); break
            }

            if (!ok) {
                console.log("Ternary" + taskNoStr + " #" + testNoStr + " has NOT been tested!")
                return
            }

            if (!inputsEOF()) {
                console.log("Ternary" + taskNoStr + " #" + testNoStr + " input NOT EOF!")
                return
            }

            if (!outputsEOF()) {
                console.log("Ternary" + taskNoStr + " #" + testNoStr + " output NOT EOF!")
                return
            }
        }
        console.log("Ternary" + taskNoStr + " has been tested!")
    }
    console.log("Ternary has been tested!")
})()

// { "no": 1, "dat": "", "ans": "" }
const ternary1 = () => {
    let number = nextInt()
    number -= number > 0 ? 8 : 0
    return compareInt(number)
}

// { "no": 2, "dat": "", "ans": "" }
const ternary2 = () => {
    let number = nextInt()
    number += number > 0 ? -8 : 6
    return compareInt(number)
}

// { "no": 3, "dat": "", "ans": "" }
const ternary3 = () => {
    let number = nextInt()
    number += number > 0 ? -8 : number < 0 ? 6 : -number + 10
    return compareInt(number)
}

// { "no": 4, "dat": "", "ans": "" }
const ternary4 = () => {
    let a = nextInt()
    let b = nextInt()
    let c = nextInt()
    let positivesCount = +(a > 0) + +(b > 0) + +(c > 0)
    return compareInt(positivesCount)
}

// { "no": 5, "dat": "", "ans": "" }
const ternary5 = () => {
    let a = nextInt()
    let b = nextInt()
    let c = nextInt()
    let positivesCount = +(a > 0) + +(b > 0) + +(c > 0)
    let negativesCount = +(a < 0) + +(b < 0) + +(c < 0)
    return compareInt(positivesCount, negativesCount)
}

// { "no": 6, "dat": "2", "ans": "2" }
const ternary6 = () => {
    let a = nextFloat()
    let b = nextFloat()
    let larger = a > b ? a : b
    return compareFloat(2, larger)
}

// { "no": 7, "dat": "2", "ans": "" }
const ternary7 = () => {
    let a = nextFloat()
    let b = nextFloat()
    let smallerIndex = a < b ? 1 : 2
    return compareInt(smallerIndex)
}

// { "no": 8, "dat": "2", "ans": "2" }
const ternary8 = () => {
    let a = nextFloat()
    let b = nextFloat()
    let larger = a > b ? a : b
    let smaller = a + b - larger
    return compareFloat(2, larger, smaller)
}

// { "no": 9, "dat": "2", "ans": "2" }
const ternary9 = () => {
    let a = nextFloat()
    let b = nextFloat()
    let larger = a > b ? a : b
    let smaller = a + b - larger
    a = smaller
    b = larger
    return compareFloat(2, a, b)
}

// { "no": 10, "dat": "", "ans": "" }
const ternary10 = () => {
    let a = nextInt()
    let b = nextInt()
    a = b = a != b ? a + b : 0
    return compareInt(a, b)
}

// { "no": 11, "dat": "", "ans": "" }
const ternary11 = () => {
    let a = nextInt()
    let b = nextInt()
    a = b = a != b ? a > b ? a : b : 0
    return compareInt(a, b)
}

// { "no": 12, "dat": "2", "ans": "2" }
const ternary12 = () => {
    let a = nextFloat()
    let b = nextFloat()
    let c = nextFloat()
    let minimal = a < b && a < c ? a : b < c ? b : c
    return compareFloat(2, minimal)
}

// { "no": 13, "dat": "2", "ans": "2" }
const ternary13 = () => {
    let a = nextFloat()
    let b = nextFloat()
    let c = nextFloat()
    let minimal = a < b && a < c ? a : b < c ? b : c
    let maximal = a > b && a > c ? a : b > c ? b : c
    let between = a + b + c - minimal - maximal
    return compareFloat(2, between)
}

// { "no": 14, "dat": "2", "ans": "2" }
const ternary14 = () => {
    let a = nextFloat()
    let b = nextFloat()
    let c = nextFloat()
    let minimal = a < b && a < c ? a : b < c ? b : c
    let maximal = a > b && a > c ? a : b > c ? b : c
    return compareFloat(2, minimal, maximal)
}

// { "no": 15, "dat": "2", "ans": "2" }
const ternary15 = () => {
    let a = nextFloat()
    let b = nextFloat()
    let c = nextFloat()
    let sum = a + b + c
    sum -= a < b && a < c ? a : b < c ? b : c
    return compareFloat(2, sum)
}

// { "no": 16, "dat": "2", "ans": "2" }
const ternary16 = () => {
    let a = nextFloat()
    let b = nextFloat()
    let c = nextFloat()
    let factor = a < b && b < c ? 2 : -1
    a *= factor
    b *= factor
    c *= factor
    return compareFloat(2, a, b, c)
}

// { "no": 17, "dat": "2", "ans": "2" }
const ternary17 = () => {
    let a = nextFloat()
    let b = nextFloat()
    let c = nextFloat()
    let factor = a < b && b < c || a > b && b > c ? 2 : -1
    a *= factor
    b *= factor
    c *= factor
    return compareFloat(2, a, b, c)
}

// { "no": 18, "dat": "", "ans": "" }
const ternary18 = () => {
    let a = nextInt()
    let b = nextInt()
    let c = nextInt()
    let differIndex = a == b ? 3 : a == c ? 2 : 1
    return compareInt(differIndex)
}

// { "no": 19, "dat": "", "ans": "" }
const ternary19 = () => {
    let a = nextInt()
    let b = nextInt()
    let c = nextInt()
    let d = nextInt()
    let differIndex = a == b && b == c ? 4 : d != c ? 3 : d != b ? 2 : 1
    return compareInt(differIndex)
}

// { "no": 20, "dat": "2", "ans": "2" }
const ternary20 = () => {
    let a = nextFloat()
    let b = nextFloat()
    let c = nextFloat()
    let point = Math.abs(b - a) < Math.abs(c - a) ? b : c
    let distance = Math.abs(point - a)
    return compareFloat(2, point, distance)
}

// { "no": 21, "dat": "", "ans": "" }
const ternary21 = () => {
    let x = nextInt()
    let y = nextInt()
    let result = x == 0 && y == 0 ? 0 : y == 0 ? 1 : x == 0 ? 2 : 3
    return compareInt(result)
}

// { "no": 22, "dat": "2", "ans": "" }
const ternary22 = () => {
    let x = nextFloat()
    let y = nextFloat()
    let quarter = x > 0 ? y > 0 ? 1 : 4 : y > 0 ? 2 : 3
    return compareInt(quarter)
}

// { "no": 23, "dat": "", "ans": "" }
const ternary23 = () => {
    let x1 = nextInt()
    let y1 = nextInt()
    let x2 = nextInt()
    let y2 = nextInt()
    let x3 = nextInt()
    let y3 = nextInt()
    let x4 = x1 == x2 ? x3 : x1 == x3 ? x2 : x1
    let y4 = y1 == y2 ? y3 : y1 == y3 ? y2 : y1
    return compareInt(x4, y4)
}

// { "no": 24, "dat": "2", "ans": "2" }
const ternary24 = () => {
    let x = nextFloat()
    let fx = x > 0 ? 2 * Math.sin(x) : 6 - x
    return compareFloat(2, fx)
}

// { "no": 25, "dat": "", "ans": "" }
const ternary25 = () => {
    let x = nextInt()
    let fx = x < -2 || x > 2 ? 2 * x : -3 * x
    return compareInt(fx)
}

// { "no": 26, "dat": "2", "ans": "2" }
const ternary26 = () => {
    let x = nextFloat()
    let fx = x <= 0 ? -x : x >= 2 ? 4 : x * x
    return compareFloat(2, fx)
}

// { "no": 27, "dat": "2", "ans": "" }
const ternary27 = () => {
    let x = nextFloat()
    let fx = x < 0 ? 0 : parseInt(x) % 2 == 0 ? 1 : -1
    return compareInt(fx)
}

// { "no": 28, "dat": "", "ans": "" }
const ternary28 = () => {
    let year = nextInt()
    let daysAmount = year % 400 == 0 ? 366 : year % 100 == 0 ? 365 : year % 4 == 0 ? 366 : 365
    return compareInt(daysAmount)
}

// { "no": 29, "dat": "", "ans": "" }
const ternary29 = () => {
    let number = nextInt()
    let description = number == 0 ? "zero" : number > 0 ? "positive" : "negative"
    description += number == 0 ? "" : number % 2 == 0 ? " even" : " odd"
    description += " number"
    return compareLine(description)
}

// { "no": 30, "dat": "", "ans": "" }
const ternary30 = () => {
    let number = nextInt()
    let description = number < 10 ? "one" : number < 100 ? "two" : "three"
    description += "-digit "
    description += number % 2 == 0 ? "even" : "odd"
    description += " number"
    return compareLine(description)
}