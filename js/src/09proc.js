import {
    nextFloat
    , compareFloat
    , nextInt
    , compareInt
    , initScanner
    , initChecker
    , inputsEOF
    , outputsEOF
    , compareBoolean
} from './io.js'

(async () => {
    for (let taskNo = 1; taskNo <= 60; taskNo++) {
        const taskNoStr = String(taskNo).padStart(3, 0)
        for (let testNo = 1; testNo <= 100; testNo++) {
            const testNoStr = String(testNo).padStart(3, 0)
            const filePath = `../data/kt-gov2/09proc/Proc${taskNoStr}/${testNoStr}/${testNoStr}`

            await initScanner(filePath + '.dat')
            await initChecker(filePath + '.ans')

            let ok = true
            switch (taskNo) {
                case 1: ok = proc1(); break
                case 2: ok = proc2(); break
                case 3: ok = proc3(); break
                case 4: ok = proc4(); break
                case 5: ok = proc5(); break
                case 6: ok = proc6(); break
                case 7: ok = proc7(); break
                case 8: ok = proc8(); break
                case 9: ok = proc9(); break
                case 10: ok = proc10(); break
                case 11: ok = proc11(); break
                case 12: ok = proc12(); break
                case 13: ok = proc13(); break
                case 14: ok = proc14(); break
                case 15: ok = proc15(); break
                case 16: ok = proc16(); break
                case 17: ok = proc17(); break
                case 18: ok = proc18(); break
                case 19: ok = proc19(); break
                case 20: ok = proc20(); break
                case 21: ok = proc21(); break
                case 22: ok = proc22(); break
                case 23: ok = proc23(); break
                case 24: ok = proc24(); break
                case 25: ok = proc25(); break
                case 26: ok = proc26(); break
                case 27: ok = proc27(); break
                case 28: ok = proc28(); break
                case 29: ok = proc29(); break
                case 30: ok = proc30(); break
                case 31: ok = proc31(); break
                case 32: ok = proc32(); break
                case 33: ok = proc33(); break
                case 34: ok = proc34(); break
                case 35: ok = proc35(); break
                case 36: ok = proc36(); break
                case 37: ok = proc37(); break
                case 38: ok = proc38(); break
                case 39: ok = proc39(); break
                case 40: ok = proc40(); break
                case 41: ok = proc41(); break
                case 42: ok = proc42(); break
                case 43: ok = proc43(); break
                case 44: ok = proc44(); break
                case 45: ok = proc45(); break
                case 46: ok = proc46(); break
                case 47: ok = proc47(); break
                case 48: ok = proc48(); break
                case 49: ok = proc49(); break
                case 50: ok = proc50(); break
                case 51: ok = proc51(); break
                case 52: ok = proc52(); break
                case 53: ok = proc53(); break
                case 54: ok = proc54(); break
                case 55: ok = proc55(); break
                case 56: ok = proc56(); break
                case 57: ok = proc57(); break
                case 58: ok = proc58(); break
                case 59: ok = proc59(); break
                case 60: ok = proc60(); break
            }

            if (!ok) {
                console.log("Proc" + taskNoStr + " #" + testNoStr + " has NOT been tested!")
                return
            }

            if (!inputsEOF()) {
                console.log("Proc" + taskNoStr + " #" + testNoStr + " input NOT EOF!")
                return
            }

            if (!outputsEOF()) {
                console.log("Proc" + taskNoStr + " #" + testNoStr + " output NOT EOF!")
                return
            }
        }
        console.log("Proc" + taskNoStr + " has been tested!")
    }
    console.log("Proc has been tested!")
})()

class TNumber {
    constructor() {
        this.value = 0
    }
}

// { "no": 1, "dat": "2", "ans": "2" }
const proc1 = () => {
    let b = new TNumber()
    for (let i = 0; i < 5; i++) {
        let a = nextFloat()
        powerA3(a, b)
        if (!compareFloat(2, b.value)) {
            return false
        }
    }
    return true
}

const powerA3 = (a, b) => {
    b.value = a * a * a
}

// { "no": 2, "dat": "2", "ans": "2" }
const proc2 = () => {
    let b = new TNumber()
    let c = new TNumber()
    let d = new TNumber()
    for (let i = 0; i < 5; i++) {
        let a = nextFloat()
        powerA234(a, b, c, d)
        if (!compareFloat(2, b.value, c.value, d.value)) {
            return false
        }
    }
    return true
}

const powerA234 = (a, b, c, d) => {
    b.value = a * a
    c.value = b.value * a
    d.value = c.value * a
}

// { "no": 3, "dat": "2", "ans": "2" }
const proc3 = () => {
    let a = nextFloat()
    let b = nextFloat()
    let c = nextFloat()
    let d = nextFloat()
    let am = new TNumber()
    let gm = new TNumber()
    mean(a, b, am, gm)
    if (!compareFloat(2, am.value, gm.value)) {
        return false
    }
    mean(a, c, am, gm)
    if (!compareFloat(2, am.value, gm.value)) {
        return false
    }
    mean(a, d, am, gm)
    return compareFloat(2, am.value, gm.value)
}

const mean = (x, y, aMean, gMean) => {
    aMean.value = (x + y) / 2
    gMean.value = Math.sqrt(x * y)
}

// { "no": 4, "dat": "2", "ans": "2" }
const proc4 = () => {
    let p = new TNumber()
    let s = new TNumber()
    for (let i = 0; i < 3; i++) {
        let a = nextFloat()
        trianglePS(a, p, s)
        if (!compareFloat(2, p.value, s.value)) {
            return false
        }
    }
    return true
}

const trianglePS = (a, p, s) => {
    p.value = 3 * a
    s.value = a * a * Math.sqrt(3) / 4
}

// { "no": 5, "dat": "2", "ans": "2" }
const proc5 = () => {
    let p = new TNumber()
    let s = new TNumber()
    for (let i = 0; i < 3; i++) {
        let x1 = nextFloat()
        let y1 = nextFloat()
        let x2 = nextFloat()
        let y2 = nextFloat()
        rectPS(x1, y1, x2, y2, p, s)
        if (!compareFloat(2, p.value, s.value)) {
            return false
        }
    }
    return true
}

const rectPS = (x1, y1, x2, y2, p, s) => {
    let a = Math.abs(x2 - x1)
    let b = Math.abs(y2 - y1)
    p.value = 2 * (a + b)
    s.value = a * b
}

// { "no": 6, "dat": "", "ans": "" }
const proc6 = () => {
    let c = new TNumber()
    let s = new TNumber()
    for (let i = 0; i < 5; i++) {
        let k = nextInt()
        digitCountSum(k, c, s)
        if (!compareInt(c.value, s.value)) {
            return false
        }
    }
    return true
}

const digitCountSum = (k, c, s) => {
    c.value = s.value = 0
    while (k > 0) {
        c.value++
        s.value += k % 10
        k = parseInt(k / 10)
    }
}

// { "no": 7, "dat": "", "ans": "" }
const proc7 = () => {
    let k = new TNumber()
    for (let i = 0; i < 5; i++) {
        k.value = nextInt()
        invDigits(k)
        if (!compareInt(k.value)) {
            return false
        }
    }
    return true
}

const invDigits = (k) => {
    let inved = 0
    while (k.value > 0) {
        inved = inved * 10 + k.value % 10
        k.value = parseInt(k.value / 10)
    }
    k.value = inved
}

// { "no": 8, "dat": "", "ans": "" }
const proc8 = () => {
    let k = new TNumber()
    k.value = nextInt()
    for (let i = 0; i < 2; i++) {
        let d = nextInt()
        addRightDigit(d, k)
        if (!compareInt(k.value)) {
            return false
        }
    }
    return true
}

const addRightDigit = (d, k) => {
    k.value = k.value * 10 + d
}

// { "no": 9, "dat": "", "ans": "" }
const proc9 = () => {
    let k = new TNumber()
    k.value = nextInt()
    for (let i = 0; i < 2; i++) {
        let d = nextInt()
        addLeftDigit(d, k)
        if (!compareInt(k.value)) {
            return false
        }
    }
    return true
}

const addLeftDigit = (d, k) => {
    let temp = k.value
    while (temp > 0) {
        d *= 10
        temp = parseInt(temp / 10)
    }
    k.value += d
}

// { "no": 10, "dat": "2", "ans": "2" }
const proc10 = () => {
    let a = new TNumber()
    a.value = nextFloat()
    let b = new TNumber()
    b.value = nextFloat()
    let c = new TNumber()
    c.value = nextFloat()
    let d = new TNumber()
    d.value = nextFloat()
    swap(a, b)
    swap(c, d)
    swap(b, c)
    return compareFloat(2, a.value, b.value, c.value, d.value)
}

const swap = (a, b) => {
    a.value += b.value
    b.value = a.value - b.value
    a.value -= b.value
}

// { "no": 11, "dat": "2", "ans": "2" }
const proc11 = () => {
    let a = new TNumber()
    a.value = nextFloat()
    let b = new TNumber()
    b.value = nextFloat()
    let c = new TNumber()
    c.value = nextFloat()
    let d = new TNumber()
    d.value = nextFloat()
    minmax(a, b)
    minmax(c, d)
    minmax(a, c)
    minmax(b, d)
    return compareFloat(2, a.value, d.value)
}

const minmax = (x, y) => {
    if (x.value > y.value) {
        swap(x, y)
    }
}

// { "no": 12, "dat": "2", "ans": "2" }
const proc12 = () => {
    for (let i = 0; i < 2; i++) {
        let a = new TNumber()
        a.value = nextFloat()
        let b = new TNumber()
        b.value = nextFloat()
        let c = new TNumber()
        c.value = nextFloat()
        sortInc3(a, b, c)
        if (!compareFloat(2, a.value, b.value, c.value)) {
            return false
        }
    }
    return true
}

const sortInc3 = (a, b, c) => {
    minmax(a, b)
    minmax(a, c)
    minmax(b, c)
}

// { "no": 13, "dat": "2", "ans": "2" }
const proc13 = () => {
    for (let i = 0; i < 2; i++) {
        let a = new TNumber()
        a.value = nextFloat()
        let b = new TNumber()
        b.value = nextFloat()
        let c = new TNumber()
        c.value = nextFloat()
        sortDec3(a, b, c)
        if (!compareFloat(2, a.value, b.value, c.value)) {
            return false
        }
    }
    return true
}

const sortDec3 = (a, b, c) => {
    sortInc3(c, b, a)
}

// { "no": 14, "dat": "2", "ans": "2" }
const proc14 = () => {
    for (let i = 0; i < 2; i++) {
        let a = new TNumber()
        a.value = nextFloat()
        let b = new TNumber()
        b.value = nextFloat()
        let c = new TNumber()
        c.value = nextFloat()
        shiftRight3(a, b, c)
        if (!compareFloat(2, a.value, b.value, c.value)) {
            return false
        }
    }
    return true
}

const shiftRight3 = (a, b, c) => {
    let temp = a.value
    a.value = c.value
    c.value = b.value
    b.value = temp
}

// { "no": 15, "dat": "2", "ans": "2" }
const proc15 = () => {
    for (let i = 0; i < 2; i++) {
        let a = new TNumber()
        a.value = nextFloat()
        let b = new TNumber()
        b.value = nextFloat()
        let c = new TNumber()
        c.value = nextFloat()
        shiftLeft3(a, b, c)
        if (!compareFloat(2, a.value, b.value, c.value)) {
            return false
        }
    }
    return true
}

const shiftLeft3 = (a, b, c) => {
    let temp = a.value
    a.value = b.value
    b.value = c.value
    c.value = temp
}

// { "no": 16, "dat": "2", "ans": "" }
const proc16 = () => {
    let a = nextFloat()
    let b = nextFloat()
    let result = sign(a) + sign(b)
    return compareInt(result)
}

const sign = (x) => x < 0 ? -1 : x > 0 ? 1 : 0

// { "no": 17, "dat": "2", "ans": "" }
const proc17 = () => {
    for (let i = 0; i < 3; i++) {
        let a = nextFloat()
        let b = nextFloat()
        let c = nextFloat()
        let rc = rootCount(a, b, c)
        if (!compareInt(rc)) {
            return false
        }
    }
    return true
}

const rootCount = (a, b, c) => {
    let d = b * b - 4 * a * c
    if (d > 0) {
        return 2
    } else if (d == 0) {
        return 1
    }
    return 0
}

// { "no": 18, "dat": "2", "ans": "2" }
const proc18 = () => {
    for (let i = 0; i < 3; i++) {
        let r = nextFloat()
        let s = circleS(r)
        if (!compareFloat(2, s)) {
            return false
        }
    }
    return true
}

const circleS = (r) => {
    const PI = 3.14
    return PI * r * r
}

// { "no": 19, "dat": "2", "ans": "2" }
const proc19 = () => {
    for (let i = 0; i < 3; i++) {
        let r1 = nextFloat()
        let r2 = nextFloat()
        let s = ringS(r1, r2)
        if (!compareFloat(2, s)) {
            return false
        }
    }
    return true
}

const ringS = (r1, r2) => {
    const PI = 3.14
    let s1 = PI * r1 * r1
    let s2 = PI * r2 * r2
    return Math.abs(s1 - s2)
}

// { "no": 20, "dat": "2", "ans": "2" }
const proc20 = () => {
    for (let i = 0; i < 3; i++) {
        let a = nextFloat()
        let h = nextFloat()
        let p = triangleP(a, h)
        if (!compareFloat(2, p)) {
            return false
        }
    }
    return true
}

const triangleP = (a, h) => {
    let b = Math.sqrt(Math.pow(a / 2, 2) + h * h)
    return a + 2 * b
}

// { "no": 21, "dat": "", "ans": "" }
const proc21 = () => {
    let a = nextInt()
    let b = nextInt()
    let c = nextInt()
    let ab = sumRange(a, b)
    let bc = sumRange(b, c)
    return compareInt(ab, bc)
}

const sumRange = (a, b) => {
    let sum = 0
    for (let i = a; i <= b; i++) {
        sum += i
    }
    return sum
}

// { "no": 22, "dat": "2", "ans": "2" }
const proc22 = () => {
    let a = nextFloat()
    let b = nextFloat()
    for (let i = 0; i < 3; i++) {
        let n = nextInt()
        let result = calc(a, b, n)
        if (!compareFloat(2, result)) {
            return false
        }
    }
    return true
}

const calc = (a, b, op) => {
    let result
    switch (op) {
        case 1: result = a - b; break
        case 2: result = a * b; break
        case 3: result = a / b; break
        default: result = a + b; break
    }
    return result
}

// { "no": 23, "dat": "2", "ans": "" }
const proc23 = () => {
    for (let i = 0; i < 3; i++) {
        let x = nextFloat()
        let y = nextFloat()
        let part = quarter(x, y)
        if (!compareInt(part)) {
            return false
        }
    }
    return true
}

const quarter = (x, y) => x > 0 ? y > 0 ? 1 : 4 : y > 0 ? 2 : 3

// { "no": 24, "dat": "", "ans": "" }
const proc24 = () => {
    let counter = 0
    for (let i = 1; i <= 10; i++) {
        let k = nextInt()
        counter += +even(k)
    }
    return compareInt(counter)
}

const even = (k) => k % 2 == 0

// { "no": 25, "dat": "", "ans": "" }
const proc25 = () => {
    let counter = 0
    for (let i = 1; i <= 10; i++) {
        let k = nextInt()
        counter += +isSquare(k)
    }
    return compareInt(counter)
}

const isSquare = (k) => {
    let i = 1
    while (i * i < k) {
        i++
    }
    return i * i == k
}

// { "no": 26, "dat": "", "ans": "" }
const proc26 = () => {
    let counter = 0
    for (let i = 1; i <= 10; i++) {
        let k = nextInt()
        counter += +isPower5(k)
    }
    return compareInt(counter)
}

const isPower5 = (k) => {
    let number = 1
    while (number < k) {
        number *= 5
    }
    return number == k
}

// { "no": 27, "dat": "", "ans": "" }
const proc27 = () => {
    let n = nextInt()
    let counter = 0
    for (let i = 1; i < 11; i++) {
        let k = nextInt()
        counter += +isPowerN(k, n)
    }
    return compareInt(counter)
}

const isPowerN = (k, n) => {
    let number = 1
    while (number < k) {
        number *= n
    }
    return number == k
}

// { "no": 28, "dat": "", "ans": "" }
const proc28 = () => {
    let counter = 0
    for (let i = 1; i <= 10; i++) {
        let n = nextInt()
        counter += +isPrime(n)
    }
    return compareInt(counter)
}

const isPrime = (n) => {
    let prime = true
    for (let i = parseInt(Math.sqrt(n)); i >= 2; i--) {
        if (n % i == 0) {
            prime = false
            break
        }
    }
    return prime
}

// { "no": 29, "dat": "", "ans": "" }
const proc29 = () => {
    for (let i = 1; i < 6; i++) {
        let k = nextInt()
        let dc = digitCount(k)
        if (!compareInt(dc)) {
            return false
        }
    }
    return true
}

const digitCount = (k) => {
    let counter = 0
    while (k > 0) {
        counter++
        k = parseInt(k / 10)
    }
    return counter
}

// { "no": 30, "dat": "", "ans": "" }
const proc30 = () => {
    for (let i = 1; i <= 5; i++) {
        let k = nextInt()
        for (let j = 1; j <= 5; j++) {
            let digit = digitN(k, j)
            if (!compareInt(digit)) {
                return false
            }
        }
    }
    return true
}

const digitN = (k, n) => {
    let i = 0
    while (k > 0) {
        i++
        let digit = k % 10
        if (i == n) {
            return digit
        }
        k = parseInt(k / 10)
    }
    return -1
}

// { "no": 31, "dat": "", "ans": "" }
const proc31 = () => {
    let counter = 0
    for (let i = 1; i <= 10; i++) {
        let n = nextInt()
        counter += +isPalindrome(n)
    }
    return compareInt(counter)
}

const isPalindrome = (k) => {
    let temp = k
    let chappa = 0
    while (temp > 0) {
        chappa = chappa * 10 + temp % 10
        temp = parseInt(temp / 10)
    }
    return chappa == k
}

// { "no": 32, "dat": "2", "ans": "2" }
const proc32 = () => {
    for (let i = 1; i < 6; i++) {
        let d = nextFloat()
        let r = degToRad(d)
        if (!compareFloat(2, r)) {
            return false
        }
    }
    return true
}

const degToRad = (d) => {
    const PI = 3.14
    let r = d * PI / 180
    return r
}

// { "no": 33, "dat": "2", "ans": "2" }
const proc33 = () => {
    for (let i = 1; i < 6; i++) {
        let r = nextFloat()
        let d = radToDeg(r)
        if (!compareFloat(2, d)) {
            return false
        }
    }
    return true
}

const radToDeg = (r) => {
    const PI = 3.14
    let d = r * 180 / PI
    return d
}

// { "no": 34, "dat": "", "ans": "2" }
const proc34 = () => {
    for (let i = 1; i < 6; i++) {
        let n = nextInt()
        let f = fact(n)
        if (!compareFloat(2, f)) {
            return false
        }
    }
    return true
}

const fact = (n) => {
    if (n <= 1) {
        return 1
    }
    return n * fact(n - 1)
}

// { "no": 35, "dat": "", "ans": "2" }
const proc35 = () => {
    for (let i = 1; i < 6; i++) {
        let n = nextInt()
        let f = fact2(n)
        if (!compareFloat(2, f)) {
            return false
        }
    }
    return true
}

const fact2 = (n) => {
    if (n <= 1) {
        return 1
    }
    return n * fact2(n - 2)
}

// { "no": 36, "dat": "", "ans": "" }
const proc36 = () => {
    for (let i = 1; i < 6; i++) {
        let n = nextInt()
        let f = fib(n)
        if (!compareInt(f)) {
            return false
        }
    }
    return true
}

const fib = (n) => {
    if (n <= 2) {
        return 1
    }
    return fib(n - 1) + fib(n - 2)
}

// { "no": 37, "dat": "2", "ans": "2" }
const proc37 = () => {
    let p = nextFloat()
    for (let i = 1; i <= 3; i++) {
        let a = nextFloat()
        let pow = power1(a, p)
        if (!compareFloat(2, pow)) {
            return false
        }
    }
    return true
}

const power1 = (a, b) => {
    if (a <= 0) {
        return 0
    }
    return Math.exp(Math.log(a) * b)
}

// { "no": 38, "dat": "2", "ans": "2" }
const proc38 = () => {
    let a = nextFloat()
    for (let i = 1; i <= 3; i++) {
        let n = nextInt()
        let res = power2(a, n)
        if (!compareFloat(2, res)) {
            return false
        }
    }
    return true
}

const power2 = (a, n) => {
    let result = 1
    let isPositive = false
    if (n < 0) {
        n *= -1
        isPositive = true
    }
    for (let i = 1; i <= n; i++) {
        result *= a
    }
    if (isPositive) {
        result = 1 / result
    }
    return result
}

// { "no": 39, "dat": "2", "ans": "2" }
const proc39 = () => {
    let p = nextFloat()
    for (let i = 1; i <= 3; i++) {
        let a = nextFloat()
        let res = power3(a, p)
        if (!compareFloat(2, res)) {
            return false
        }
    }
    return true
}

const power3 = (a, b) => b - parseInt(b) == 0 ? power2(a, b) : power1(a, b)

// { "no": 40, "dat": "7", "ans": "7" }
const proc40 = () => {
    let x = nextFloat()
    for (let i = 1; i < 7; i++) {
        let e = nextFloat()
        let res = exp1(x, e)
        if (!compareFloat(7, res)) {
            return false
        }
    }
    return true
}

const exp1 = (x, e) => {
    let daraja = 1
    let fact = 1
    let sum = 0
    let i = 1
    while (daraja / fact > e) {
        sum += daraja / fact
        daraja *= x
        fact *= i++
    }
    return sum
}

// { "no": 41, "dat": "7", "ans": "7" }
const proc41 = () => {
    let x = nextFloat()
    for (let i = 1; i < 7; i++) {
        let e = nextFloat()
        let res = sin1(x, e)
        if (!compareFloat(7, res)) {
            return false
        }
    }
    return true
}

const sin1 = (x, e) => {
    let daraja = x
    let fact = 1
    let sum = 0
    let factor = 1
    let i = 1
    while (Math.abs(daraja / fact) > e) {
        sum += factor * daraja / fact
        factor *= -1
        i += 2
        fact *= (i - 1) * i
        daraja *= x * x
    }
    return sum
}

// { "no": 42, "dat": "7", "ans": "7" }
const proc42 = () => {
    let x = nextFloat()
    for (let i = 1; i < 7; i++) {
        let e = nextFloat()
        let res = cos1(x, e)
        if (!compareFloat(7, res)) {
            return false
        }
    }
    return true
}

const cos1 = (x, e) => {
    let daraja = 1
    let fact = 1
    let sum = 0
    let i = 0
    let factor = 1
    while (Math.abs(daraja / fact) > e) {
        sum += factor * daraja / fact
        factor *= -1
        daraja *= x * x
        i += 2
        fact *= (i - 1) * i
    }
    return sum
}

// { "no": 43, "dat": "7", "ans": "7" }
const proc43 = () => {
    let x = nextFloat()
    for (let i = 1; i < 7; i++) {
        let e = nextFloat()
        let res = ln1(x, e)
        if (!compareFloat(7, res)) {
            return false
        }
    }
    return true
}

const ln1 = (x, e) => {
    let daraja = x
    let number = 1
    let sum = 0
    let factor = 1
    while (Math.abs(daraja / number) > e) {
        sum += factor * daraja / number++
        daraja *= x
        factor *= -1
    }
    return sum
}

// { "no": 44, "dat": "7", "ans": "7" }
const proc44 = () => {
    let x = nextFloat()
    for (let i = 1; i < 7; i++) {
        let e = nextFloat()
        let res = arctg1(x, e)
        if (!compareFloat(7, res)) {
            return false
        }
    }
    return true
}

const arctg1 = (x, e) => {
    let daraja = x
    let number = 1
    let sum = 0
    let factor = 1
    while (Math.abs(daraja / number) > e) {
        sum += factor * daraja / number
        daraja *= x * x
        number += 2
        factor *= -1
    }
    return sum
}

// { "no": 45, "dat": "7", "ans": "7" }
const proc45 = () => {
    let x = nextFloat()
    let a = nextFloat()
    for (let i = 1; i < 7; i++) {
        let e = nextFloat()
        let res = power4(x, a, e)
        if (!compareFloat(7, res)) {
            return false
        }
    }
    return true
}

const power4 = (x, a, e) => {
    let aho = 1
    let daraja = 1
    let fact = 1
    let sum = 0
    let number = 0
    while (Math.abs(aho * daraja / fact) > e) {
        sum += aho * daraja / fact
        aho *= (a - number++)
        fact *= number
        daraja *= x
    }
    return sum
}

// { "no": 46, "dat": "", "ans": "" }
const proc46 = () => {
    let a = nextInt()
    for (let i = 1; i <= 3; i++) {
        let b = nextInt()
        let gcd = gcd2(a, b)
        if (!compareInt(gcd)) {
            return false
        }
    }
    return true
}

const gcd2 = (a, b) => {
    while (b != 0) {
        let r = a % b
        a = b
        b = r
    }
    return a
}

// { "no": 47, "dat": "", "ans": "" }
const proc47 = () => {
    let a = nextInt()
    let b = nextInt()
    for (let i = 1; i < 4; i++) {
        let c = nextInt()
        let d = nextInt()
        let p = new TNumber()
        let q = new TNumber()
        fraq1(d * a + b * c, b * d, p, q)
        if (!compareInt(p.value, q.value)) {
            return false
        }
    }
    return true
}

const fraq1 = (a, b, p, q) => {
    let gcd = gcd2(a, b)
    p.value = a / gcd
    q.value = b / gcd
    if (q.value < 0) {
        p.value *= -1
        q.value *= -1
    }
}

// { "no": 48, "dat": "", "ans": "" }
const proc48 = () => {
    let a = nextInt()
    for (let i = 1; i < 4; i++) {
        let b = nextInt()
        let lcm = lcm2(a, b)
        if (!compareInt(lcm)) {
            return false
        }
    }
    return true
}

const lcm2 = (a, b) => b / gcd2(a, b) * a

// { "no": 49, "dat": "", "ans": "" }
const proc49 = () => {
    let a = nextInt()
    let b = nextInt()
    let c = nextInt()
    let d = nextInt()
    return compareInt(gcd3(a, b, c)) && compareInt(gcd3(a, c, d)) && compareInt(gcd3(b, c, d))
}

const gcd3 = (a, b, c) => gcd2(a, gcd2(b, c))

// { "no": 50, "dat": "", "ans": "" }
const proc50 = () => {
    for (let i = 1; i < 6; i++) {
        let t = nextInt()
        let h = new TNumber()
        let m = new TNumber()
        let s = new TNumber()
        timeToHMS(t, h, m, s)
        if (!compareInt(h.value, m.value, s.value)) {
            return false
        }
    }
    return true
}

const timeToHMS = (t, h, m, s) => {
    h.value = parseInt(t / 3600)
    t %= 3600
    m.value = parseInt(t / 60)
    s.value = t % 60
}

// { "no": 51, "dat": "", "ans": "" }
const proc51 = () => {
    let h = new TNumber()
    h.value = nextInt()
    let m = new TNumber()
    m.value = nextInt()
    let s = new TNumber()
    s.value = nextInt()
    let t = nextInt()
    incTime(h, m, s, t)
    return compareInt(h.value, m.value, s.value)
}

const incTime = (h, m, s, t) => {
    t += h.value * 3600 + m.value * 60 + s.value
    h.value = parseInt(t / 3600)
    t %= 3600
    m.value = parseInt(t / 60)
    s.value = t % 60
}

// { "no": 52, "dat": "", "ans": "" }
const proc52 = () => {
    for (let i = 1; i < 6; i++) {
        let y = nextInt()
        if (!compareBoolean(isLeapYear(y))) {
            return false
        }
    }
    return true
}

const isLeapYear = (y) => y % 400 == 0 ? true : y % 100 == 0 ? false : y % 4 == 0 ? true : false

// { "no": 53, "dat": "", "ans": "" }
const proc53 = () => {
    let y = nextInt()
    for (let i = 1; i < 4; i++) {
        let m = nextInt()
        let days = monthDays(m, y)
        if (!compareInt(days)) {
            return false
        }
    }
    return true
}

const monthDays = (m, y) => {
    let days = 0
    switch (m) {
        case 1: case 3: case 5: case 7: case 8: case 10: case 12:
            days = 31
            break
        case 4: case 6: case 9: case 11:
            days = 30
            break
        case 2:
            days = 28 + +isLeapYear(y)
            break
    }
    return days
}

// { "no": 54, "dat": "", "ans": "" }
const proc54 = () => {
    for (let i = 1; i <= 3; i++) {
        let d = new TNumber()
        d.value = nextInt()
        let m = new TNumber()
        m.value = nextInt()
        let y = new TNumber()
        y.value = nextInt()
        prevDate(d, m, y)
        if (!compareInt(d.value, m.value, y.value)) {
            return false
        }
    }
    return true
}

const prevDate = (d, m, y) => {
    if (d.value == 1 && m.value == 1) {
        d.value = 31
        m.value = 12
        y.value--
    } else if (d.value == 1) {
        m.value--
        d.value = monthDays(m.value, y.value)
    } else {
        d.value--
    }
}

// { "no": 55, "dat": "", "ans": "" }
const proc55 = () => {
    for (let i = 1; i < 4; i++) {
        let d = new TNumber()
        d.value = nextInt()
        let m = new TNumber()
        m.value = nextInt()
        let y = new TNumber()
        y.value = nextInt()
        nextDate(d, m, y)
        if (!compareInt(d.value, m.value, y.value)) {
            return false
        }
    }
    return true
}

const nextDate = (d, m, y) => {
    if (d.value == 31 && m.value == 12) {
        d.value = m.value = 1
        y.value++
    } else if (d.value == monthDays(m.value, y.value)) {
        d.value = 1
        m.value++
    } else {
        d.value++
    }
}

class Point {
    constructor() {
        this.x = 0
        this.y = 0
    }
}

// { "no": 56, "dat": "2", "ans": "2" }
const proc56 = () => {
    let a = new Point()
    a.x = nextFloat()
    a.y = nextFloat()
    for (let i = 1; i <= 3; i++) {
        let b = new Point()
        b.x = nextFloat()
        b.y = nextFloat()
        let darozi = leng(a.x, a.y, b.x, b.y)
        if (!compareFloat(2, darozi)) {
            return false
        }
    }
    return true
}

const leng = (xA, yA, xB, yB) => Math.sqrt(Math.pow(xB - xA, 2) + Math.pow(yB - yA, 2))

// { "no": 57, "dat": "2", "ans": "2" }
const proc57 = () => {
    let a = new Point()
    let b = new Point()
    let c = new Point()
    let d = new Point()

    a.x = nextFloat()
    a.y = nextFloat()
    b.x = nextFloat()
    b.y = nextFloat()
    c.x = nextFloat()
    c.y = nextFloat()
    d.x = nextFloat()
    d.y = nextFloat()

    let p = perim(a.x, a.y, b.x, b.y, c.x, c.y)
    if (!compareFloat(2, p)) {
        return false
    }
    p = perim(a.x, a.y, b.x, b.y, d.x, d.y)
    if (!compareFloat(2, p)) {
        return false
    }
    p = perim(a.x, a.y, c.x, c.y, d.x, d.y)
    return compareFloat(2, p)
}

const perim = (xA, yA, xB, yB, xC, yC) => {
    let a = leng(xA, yA, xB, yB)
    let b = leng(xA, yA, xC, yC)
    let c = leng(xB, yB, xC, yC)
    return a + b + c
}

// { "no": 58, "dat": "2", "ans": "2" }
const proc58 = () => {
    let a = new Point()
    let b = new Point()
    let c = new Point()
    let d = new Point()

    a.x = nextFloat()
    a.y = nextFloat()
    b.x = nextFloat()
    b.y = nextFloat()
    c.x = nextFloat()
    c.y = nextFloat()
    d.x = nextFloat()
    d.y = nextFloat()

    let s = area(a.x, a.y, b.x, b.y, c.x, c.y)
    if (!compareFloat(2, s)) {
        return false
    }
    s = area(a.x, a.y, b.x, b.y, d.x, d.y)
    if (!compareFloat(2, s)) {
        return false
    }
    s = area(a.x, a.y, c.x, c.y, d.x, d.y)
    return compareFloat(2, s)
}

const area = (xA, yA, xB, yB, xC, yC) => {
    let p = perim(xA, yA, xB, yB, xC, yC)
    let halfP = p / 2
    let a = new TNumber()
    let b = new TNumber()
    let c = new TNumber()
    triangleSides(xA, yA, xB, yB, xC, yC, a, b, c)
    return Math.sqrt(halfP * (halfP - a.value) * (halfP - b.value) * (halfP - c.value))
}

const triangleSides = (xA, yA, xB, yB, xC, yC, a, b, c) => {
    a.value = leng(xA, yA, xB, yB)
    b.value = leng(xA, yA, xC, yC)
    c.value = leng(xB, yB, xC, yC)
}

// { "no": 59, "dat": "2", "ans": "2" }
const proc59 = () => {
    let p = new Point()
    let a = new Point()
    let b = new Point()
    let c = new Point()

    p.x = nextFloat()
    p.y = nextFloat()
    a.x = nextFloat()
    a.y = nextFloat()
    b.x = nextFloat()
    b.y = nextFloat()
    c.x = nextFloat()
    c.y = nextFloat()

    let d = dist(p.x, p.y, a.x, a.y, b.x, b.y)
    if (!compareFloat(2, d)) {
        return false
    }
    d = dist(p.x, p.y, a.x, a.y, c.x, c.y)
    if (!compareFloat(2, d)) {
        return false
    }
    d = dist(p.x, p.y, b.x, b.y, c.x, c.y)
    return compareFloat(2, d)
}

const dist = (xP, yP, xA, yA, xB, yB) => area(xP, yP, xA, yA, xB, yB) / leng(xA, yA, xB, yB) * 2

// { "no": 60, "dat": "2", "ans": "2" }
const proc60 = () => {
    let a = new Point()
    let b = new Point()
    let c = new Point()
    let d = new Point()

    a.x = nextFloat()
    a.y = nextFloat()
    b.x = nextFloat()
    b.y = nextFloat()
    c.x = nextFloat()
    c.y = nextFloat()
    d.x = nextFloat()
    d.y = nextFloat()

    let hA = new TNumber()
    let hB = new TNumber()
    let hC = new TNumber()
    alts(a.x, a.y, b.x, b.y, c.x, c.y, hA, hB, hC)
    if (!compareFloat(2, hA.value, hB.value, hC.value)) {
        return false
    }
    alts(a.x, a.y, b.x, b.y, d.x, d.y, hA, hB, hC)
    if (!compareFloat(2, hA.value, hB.value, hC.value)) {
        return false
    }
    alts(a.x, a.y, c.x, c.y, d.x, d.y, hA, hB, hC)
    return compareFloat(2, hA.value, hB.value, hC.value)
}

const alts = (xA, yA, xB, yB, xC, yC, hA, hB, hC) => {
    hA.value = dist(xA, yA, xB, yB, xC, yC)
    hB.value = dist(xB, yB, xA, yA, xC, yC)
    hC.value = dist(xC, yC, xA, yA, xB, yB)
}