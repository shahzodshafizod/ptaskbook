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
    for (let taskNo = 1; taskNo <= 40; taskNo++) {
        const taskNoStr = String(taskNo).padStart(3, 0)
        for (let testNo = 1; testNo <= 100; testNo++) {
            const testNoStr = String(testNo).padStart(3, 0)
            const filePath = `../data/kt-gov2/06for/For${taskNoStr}/${testNoStr}/${testNoStr}`

            await initScanner(filePath + '.dat')
            await initChecker(filePath + '.ans')

            let ok = true
            switch (taskNo) {
                case 1: ok = for1(); break
                case 2: ok = for2(); break
                case 3: ok = for3(); break
                case 4: ok = for4(); break
                case 5: ok = for5(); break
                case 6: ok = for6(); break
                case 7: ok = for7(); break
                case 8: ok = for8(); break
                case 9: ok = for9(); break
                case 10: ok = for10(); break
                case 11: ok = for11(); break
                case 12: ok = for12(); break
                case 13: ok = for13(); break
                case 14: ok = for14(); break
                case 15: ok = for15(); break
                case 16: ok = for16(); break
                case 17: ok = for17(); break
                case 18: ok = for18(); break
                case 19: ok = for19(); break
                case 20: ok = for20(); break
                case 21: ok = for21(); break
                case 22: ok = for22(); break
                case 23: ok = for23(); break
                case 24: ok = for24(); break
                case 25: ok = for25(); break
                case 26: ok = for26(); break
                case 27: ok = for27(); break
                case 28: ok = for28(); break
                case 29: ok = for29(); break
                case 30: ok = for30(); break
                case 31: ok = for31(); break
                case 32: ok = for32(); break
                case 33: ok = for33(); break
                case 34: ok = for34(); break
                case 35: ok = for35(); break
                case 36: ok = for36(); break
                case 37: ok = for37(); break
                case 38: ok = for38(); break
                case 39: ok = for39(); break
                case 40: ok = for40(); break
            }

            if (!ok) {
                console.log("For" + taskNoStr + " #" + testNoStr + " has NOT been tested!")
                return
            }

            if (!inputsEOF()) {
                console.log("For" + taskNoStr + " #" + testNoStr + " input NOT EOF!")
                return
            }

            if (!outputsEOF()) {
                console.log("For" + taskNoStr + " #" + testNoStr + " output NOT EOF!")
                return
            }
        }
        console.log("For" + taskNoStr + " has been tested!")
    }
    console.log("For has been tested!")
})()

// { "no": 1, "dat": "", "ans": "" }
const for1 = () => {
    let k = nextInt()
    let n = nextInt()
    for (let i = 0; i < n; i++) {
        if (!compareInt(k)) {
            return false
        }
    }
    return true
}

// { "no": 2, "dat": "", "ans": "" }
const for2 = () => {
    let a = nextInt()
    let b = nextInt()
    let n = 0
    for (let i = a; i <= b; i++) {
        if (!compareInt(i)) {
            return false
        }
        n++
    }
    return compareInt(n)
}

// { "no": 3, "dat": "", "ans": "" }
const for3 = () => {
    let a = nextInt()
    let b = nextInt()
    let n = 0
    for (let i = b - 1; i > a; i--) {
        if (!compareInt(i)) {
            return false
        }
        n++
    }
    return compareInt(n)
}

// { "no": 4, "dat": "2", "ans": "2" }
const for4 = () => {
    let price = nextFloat()
    for (let i = 1; i <= 10; i++) {
        let cost = i * price
        if (!compareFloat(2, cost)) {
            return false
        }
    }
    return true
}

// { "no": 5, "dat": "2", "ans": "2" }
const for5 = () => {
    let price = nextFloat()
    for (let i = 0.1; i <= 1; i += 0.1) {
        let cost = i * price
        if (!compareFloat(2, cost)) {
            return false
        }
    }
    return true
}

// { "no": 6, "dat": "2", "ans": "2" }
const for6 = () => {
    let price = nextFloat()
    for (let i = 1.2; i <= 2; i += 0.2) {
        let cost = i * price
        if (!compareFloat(2, cost)) {
            return false
        }
    }
    return true
}

// { "no": 7, "dat": "", "ans": "" }
const for7 = () => {
    let a = nextInt()
    let b = nextInt()
    let sum = 0
    for (let i = a; i <= b; i++) {
        sum += i
    }
    return compareInt(sum)
}

// { "no": 8, "dat": "", "ans": "" }
const for8 = () => {
    let a = nextInt()
    let b = nextInt()
    let mult = 1
    for (let i = a; i <= b; i++) {
        mult *= i
    }
    return compareInt(mult)
}

// { "no": 9, "dat": "", "ans": "" }
const for9 = () => {
    let a = nextInt()
    let b = nextInt()
    let sum = 0
    for (let i = a; i <= b; i++) {
        sum += i * i
    }
    return compareInt(sum)
}

// { "no": 10, "dat": "", "ans": "6" }
const for10 = () => {
    let n = nextInt()
    let sum = 0
    for (let i = 1; i <= n; i++) {
        sum += 1 / i
    }
    return compareFloat(6, sum)
}

// { "no": 11, "dat": "", "ans": "" }
const for11 = () => {
    let n = nextInt()
    let sum = 0
    for (let i = 0; i <= n; i++) {
        sum += parseInt(Math.pow(n + i, 2))
    }
    return compareInt(sum)
}

// { "no": 12, "dat": "", "ans": "6" }
const for12 = () => {
    let n = nextInt()
    let mult = 1
    let number = 1.1
    for (let i = 0; i < n; i++) {
        mult *= number
        number += 0.1
    }
    return compareFloat(6, mult)
}

// { "no": 13, "dat": "", "ans": "2" }
const for13 = () => {
    let n = nextInt()
    let factor = 1
    let sum = 0
    let number = 1.1
    for (let i = 0; i < n; i++) {
        sum += factor * number
        factor *= -1
        number += 0.1
    }
    return compareFloat(2, sum)
}

// { "no": 14, "dat": "", "ans": "" }
const for14 = () => {
    let n = nextInt()
    let sum = 0
    for (let i = 1; i <= n; i++) {
        sum += 2 * i - 1
        if (!compareInt(sum)) {
            return false
        }
    }
    return true
}

// { "no": 15, "dat": "2", "ans": "2" }
const for15 = () => {
    let a = nextFloat()
    let n = nextInt()
    let mult = 1
    for (let i = 0; i < n; i++) {
        mult *= a
    }
    return compareFloat(2, mult)
}

// { "no": 16, "dat": "2", "ans": "2" }
const for16 = () => {
    let a = nextFloat()
    let n = nextInt()
    let mult = 1
    for (let i = 0; i < n; i++) {
        mult *= a
        if (!compareFloat(2, mult)) {
            return false
        }
    }
    return true
}

// { "no": 17, "dat": "2", "ans": "2" }
const for17 = () => {
    let a = nextFloat()
    let n = nextInt()
    let sum = 0
    let daraja = 1
    for (let i = 0; i <= n; i++) {
        sum += daraja
        daraja *= a
    }
    return compareFloat(2, sum)
}

// { "no": 18, "dat": "2", "ans": "2" }
const for18 = () => {
    let a = nextFloat()
    let n = nextInt()
    let factor = 1
    let sum = 0
    let daraja = 1
    for (let i = 0; i <= n; i++) {
        sum += factor * daraja
        factor *= -1
        daraja *= a
    }
    return compareFloat(2, sum)
}

// { "no": 19, "dat": "", "ans": "2" }
const for19 = () => {
    let n = nextInt()
    let fact = 1
    for (let i = 2; i <= n; i++) {
        fact *= i
    }
    return compareFloat(2, fact)
}

// { "no": 20, "dat": "", "ans": "2" }
const for20 = () => {
    let n = nextInt()
    let fact = 1
    let sum = 0
    for (let i = 1; i <= n; i++) {
        fact *= i
        sum += fact
    }
    return compareFloat(2, sum)
}

// { "no": 21, "dat": "", "ans": "8" }
const for21 = () => {
    let n = nextInt()
    let sum = 0
    let fact = 1
    for (let i = 0; i <= n;) {
        sum += 1 / fact
        fact *= ++i
    }
    return compareFloat(8, sum)
}

// { "no": 22, "dat": "8", "ans": "8" }
const for22 = () => {
    let x = nextFloat()
    let n = nextInt()
    let sum = 0
    let fact = 1
    let daraja = 1
    for (let i = 0; i <= n;) {
        sum += daraja / fact
        daraja *= x
        fact *= ++i
    }
    return compareFloat(8, sum)
}

// { "no": 23, "dat": "8", "ans": "8" }
const for23 = () => {
    let x = nextFloat()
    let n = nextInt()
    let sum = 0
    let factor = 1
    let daraja = 1
    let fact = 1
    for (let i = 0; i <= n;) {
        daraja *= x
        sum += factor * daraja / fact
        factor *= -1
        daraja *= x
        ++i
        fact *= (2 * i) * (2 * i + 1)
    }
    return compareFloat(8, sum)
}

// { "no": 24, "dat": "8", "ans": "8" }
const for24 = () => {
    let x = nextFloat()
    let n = nextInt()
    let sum = 0
    let factor = 1
    let fact = 1
    let daraja = 1
    for (let i = 0; i <= n;) {
        sum += factor * daraja / fact
        factor *= -1
        daraja *= x * x
        ++i
        fact *= (2 * i - 1) * (2 * i)
    }
    return compareFloat(8, sum)
}

// { "no": 25, "dat": "8", "ans": "8" }
const for25 = () => {
    let x = nextFloat()
    let n = nextInt()
    let daraja = 1
    let sum = 0
    let factor = 1
    for (let i = 1; i <= n; i++) {
        daraja *= x
        sum += factor * daraja / i
        factor *= -1
    }
    return compareFloat(8, sum)
}

// { "no": 26, "dat": "8", "ans": "8" }
const for26 = () => {
    let x = nextFloat()
    let n = nextInt()
    let sum = 0
    let daraja = 1
    let factor = 1
    for (let i = 0; i <= n; i++) {
        daraja *= x
        sum += factor * daraja / (2 * i + 1)
        factor *= -1
        daraja *= x
    }
    return compareFloat(8, sum)
}

// { "no": 27, "dat": "8", "ans": "8" }
const for27 = () => {
    let x = nextFloat()
    let n = nextInt()
    let sum = 0
    let daraja = 1
    let toqho = 1
    let juftho = 1
    for (let i = 0; i <= n;) {
        daraja *= x
        sum += toqho * daraja / (juftho * (2 * i + 1))
        daraja *= x
        ++i
        toqho *= 2 * i - 1
        juftho *= 2 * i
    }
    return compareFloat(8, sum)
}

// { "no": 28, "dat": "8", "ans": "8" }
const for28 = () => {
    let x = nextFloat()
    let n = nextInt()
    let sum = 1
    let factor = 1
    let daraja = 1
    let toqho = 1
    let juftho = 1
    for (let i = 1; i <= n;) {
        daraja *= x
        juftho *= 2 * i
        sum += factor * toqho * daraja / juftho
        factor *= -1
        ++i
        toqho *= 2 * i - 3
    }
    return compareFloat(8, sum)
}

// { "no": 29, "dat": "2", "ans": "5" }
const for29 = () => {
    let n = nextInt()
    let a = nextFloat()
    let b = nextFloat()
    let h = (b - a) / n
    if (!compareFloat(5, h)) {
        return false
    }
    let point = a
    for (let index = 0; index <= n; index++) {
        if (!compareFloat(5, point)) {
            return false
        }
        point += h
    }
    return true
}

// { "no": 30, "dat": "2", "ans": "5" }
const for30 = () => {
    let n = nextInt()
    let a = nextFloat()
    let b = nextFloat()
    let h = (b - a) / n
    if (!compareFloat(5, h)) {
        return false
    }
    let value
    for (let i = a; i <= b; i += h) {
        value = 1 - Math.sin(i)
        if (!compareFloat(5, value)) {
            return false
        }
    }
    return true
}

// { "no": 31, "dat": "", "ans": "6" }
const for31 = () => {
    let n = nextInt()
    let a0 = 2
    let ak
    for (let i = 1; i <= n; i++) {
        ak = 2 + 1 / a0
        if (!compareFloat(6, ak)) {
            return false
        }
        a0 = ak
    }
    return true
}

// { "no": 32, "dat": "", "ans": "6" }
const for32 = () => {
    let n = nextInt()
    let a0 = 1
    let ak
    for (let i = 1; i <= n; ++i) {
        ak = (a0 + 1) / i
        if (!compareFloat(6, ak)) {
            return false
        }
        a0 = ak
    }
    return true
}

// { "no": 33, "dat": "", "ans": "" }
const for33 = () => {
    let n = nextInt()
    let f1 = 1
    let f2 = 1
    if (!compareInt(f1, f2)) {
        return false
    }
    let fk
    for (let i = 3; i <= n; i++) {
        fk = f1 + f2
        if (!compareInt(fk)) {
            return false
        }
        f1 = f2
        f2 = fk
    }
    return true
}

// { "no": 34, "dat": "", "ans": "6" }
const for34 = () => {
    let n = nextInt()
    let a1 = 1
    let a2 = 2
    if (!compareFloat(6, a1, a2)) {
        return false
    }
    let ak
    for (let i = 3; i <= n; i++) {
        ak = (a1 + 2 * a2) / 3
        if (!compareFloat(6, ak)) {
            return false
        }
        a1 = a2
        a2 = ak
    }
    return true
}

// { "no": 35, "dat": "", "ans": "" }
const for35 = () => {
    let n = nextInt()
    let a1 = 1
    let a2 = 2
    let a3 = 3
    if (!compareInt(a1, a2, a3)) {
        return false
    }
    let ak
    for (let i = 4; i <= n; i++) {
        ak = a3 + a2 - 2 * a1
        if (!compareInt(ak)) {
            return false
        }
        a1 = a2
        a2 = a3
        a3 = ak
    }
    return true
}

// { "no": 36, "dat": "", "ans": "2" }
const for36 = () => {
    let n = nextInt()
    let k = nextInt()
    let daraja
    let sum = 0
    for (let i = 1; i <= n; i++) {
        daraja = 1
        for (let j = 1; j <= k; j++) {
            daraja *= i
        }
        sum += daraja
    }
    return compareFloat(2, sum)
}

// { "no": 37, "dat": "", "ans": "2" }
const for37 = () => {
    let n = nextInt()
    let sum = 0
    let daraja
    for (let i = 1; i <= n; i++) {
        daraja = 1
        for (let j = 1; j <= i; j++) {
            daraja *= i
        }
        sum += daraja
    }
    return compareFloat(2, sum)
}

// { "no": 38, "dat": "", "ans": "2" }
const for38 = () => {
    let n = nextInt()
    let sum = 0
    let daraja
    for (let i = 1; i <= n; i++) {
        daraja = 1
        for (let j = n - i + 1; j > 0; j--) {
            daraja *= i
        }
        sum += daraja
    }
    return compareFloat(2, sum)
}

// { "no": 39, "dat": "", "ans": "" }
const for39 = () => {
    let a = nextInt()
    let b = nextInt()
    for (let i = a; i <= b; i++) {
        for (let j = 1; j <= i; j++) {
            if (!compareInt(i)) {
                return false
            }
        }
    }
    return true
}

// { "no": 40, "dat": "", "ans": "" }
const for40 = () => {
    let a = nextInt()
    let b = nextInt()
    for (let i = a, j = 1; i <= b; i++, j++) {
        for (let k = 1; k <= j; k++) {
            if (!compareInt(i)) {
                return false
            }
        }
    }
    return true
}