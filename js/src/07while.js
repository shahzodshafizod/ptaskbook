import {
    nextFloat
    , compareFloat
    , nextInt
    , compareInt
    , compareBoolean
    , initScanner
    , initChecker
    , inputsEOF
    , outputsEOF
} from './io.js'

(async () => {
    for (let taskNo = 1; taskNo <= 30; taskNo++) {
        const taskNoStr = String(taskNo).padStart(3, 0)
        for (let testNo = 1; testNo <= 2; testNo++) {
            const testNoStr = String(testNo).padStart(3, 0)
            const filePath = `../data/kt-gov2/07while/While${taskNoStr}/${testNoStr}/${testNoStr}`

            await initScanner(filePath + '.dat')
            await initChecker(filePath + '.ans')

            let ok = true
            switch (taskNo) {
                case 1: ok = while1(); break
                case 2: ok = while2(); break
                case 3: ok = while3(); break
                case 4: ok = while4(); break
                case 5: ok = while5(); break
                case 6: ok = while6(); break
                case 7: ok = while7(); break
                case 8: ok = while8(); break
                case 9: ok = while9(); break
                case 10: ok = while10(); break
                case 11: ok = while11(); break
                case 12: ok = while12(); break
                case 13: ok = while13(); break
                case 14: ok = while14(); break
                case 15: ok = while15(); break
                case 16: ok = while16(); break
                case 17: ok = while17(); break
                case 18: ok = while18(); break
                case 19: ok = while19(); break
                case 20: ok = while20(); break
                case 21: ok = while21(); break
                case 22: ok = while22(); break
                case 23: ok = while23(); break
                case 24: ok = while24(); break
                case 25: ok = while25(); break
                case 26: ok = while26(); break
                case 27: ok = while27(); break
                case 28: ok = while28(); break
                case 29: ok = while29(); break
                case 30: ok = while30(); break
            }

            if (!ok) {
                console.log("While" + taskNoStr + " #" + testNoStr + " has NOT been tested!")
                return
            }

            if (!inputsEOF()) {
                console.log("While" + taskNoStr + " #" + testNoStr + " input NOT EOF!")
                return
            }

            if (!outputsEOF()) {
                console.log("While" + taskNoStr + " #" + testNoStr + " output NOT EOF!")
                return
            }
        }
        console.log("While" + taskNoStr + " has been tested!")
    }
    console.log("While has been tested!")
})()

// { "no": 1, "dat": "2", "ans": "2" }
const while1 = () => {
    let a = nextFloat()
    let b = nextFloat()
    while (a >= b) {
        a -= b
    }
    return compareFloat(2, a)
}

// { "no": 2, "dat": "2", "ans": "" }
const while2 = () => {
    let a = nextFloat()
    let b = nextFloat()
    let n = 0
    while (a >= b) {
        a -= b
        n++
    }
    return compareInt(n)
}

// { "no": 3, "dat": "", "ans": "" }
const while3 = () => {
    let n = nextInt()
    let k = nextInt()
    let counter = 0
    while (n >= k) {
        n -= k
        ++counter
    }
    return compareInt(counter, n)
}

// { "no": 4, "dat": "", "ans": "" }
const while4 = () => {
    let n = nextInt()
    let daraja = 1
    while (daraja < n) {
        daraja *= 3
    }
    let isDegree = daraja == n
    return compareBoolean(isDegree)
}

// { "no": 5, "dat": "", "ans": "" }
const while5 = () => {
    let n = nextInt()
    let k = 0
    let daraja = 1
    while (daraja < n) {
        daraja *= 2
        k++
    }
    return compareInt(k)
}

// { "no": 6, "dat": "", "ans": "2" }
const while6 = () => {
    let n = nextInt()
    let fact = 1
    while (n > 0) {
        fact *= n
        n -= 2
    }
    return compareFloat(2, fact)
}

// { "no": 7, "dat": "", "ans": "" }
const while7 = () => {
    let n = nextInt()
    let k = 1
    while (k * k <= n) {
        k++
    }
    return compareInt(k)
}

// { "no": 8, "dat": "", "ans": "" }
const while8 = () => {
    let n = nextInt()
    let k = 1
    while (k * k <= n) {
        k++
    }
    k--
    return compareInt(k)
}

// { "no": 9, "dat": "", "ans": "" }
const while9 = () => {
    let n = nextInt()
    let daraja = 1
    let k = 0
    while (daraja <= n) {
        daraja *= 3
        k++
    }
    return compareInt(k)
}

// { "no": 10, "dat": "", "ans": "" }
const while10 = () => {
    let n = nextInt()
    let daraja = 1
    let k = 0
    while (daraja < n) {
        daraja *= 3
        k++
    }
    k--
    return compareInt(k)
}

// { "no": 11, "dat": "", "ans": "" }
const while11 = () => {
    let n = nextInt()
    let k = 0
    let sum = 0
    while (sum < n) {
        sum += ++k
    }
    return compareInt(k, sum)
}

// { "no": 12, "dat": "", "ans": "" }
const while12 = () => {
    let n = nextInt()
    let k = 0
    let sum = 0
    while (sum <= n) {
        sum += ++k
    }
    sum -= k
    k--
    return compareInt(k, sum)
}

// { "no": 13, "dat": "5", "ans": "5" }
const while13 = () => {
    let a = nextFloat()
    let sum = 0
    let k = 0
    while (sum <= a) {
        sum += 1 / ++k
    }
    return compareInt(k) && compareFloat(5, sum)
}

// { "no": 14, "dat": "5", "ans": "5" }
const while14 = () => {
    let a = nextFloat()
    let k = 0
    let sum = 0
    while (sum < a) {
        sum += 1 / ++k
    }
    sum -= 1 / k--
    return compareInt(k) && compareFloat(5, sum)
}

// { "no": 15, "dat": "2", "ans": "2" }
const while15 = () => {
    let s = 1000
    let p = nextFloat()
    let k = 0
    while (s <= 1100) {
        s += p * s / 100
        k++
    }
    return compareInt(k) && compareFloat(2, s)
}

// { "no": 16, "dat": "2", "ans": "3" }
const while16 = () => {
    let run = 10
    let p = nextFloat()
    let s = 10
    let k = 1
    while (s <= 200) {
        run += run * p / 100
        s += run
        k++
    }
    return compareInt(k) && compareFloat(3, s)
}

// { "no": 17, "dat": "", "ans": "" }
const while17 = () => {
    let n = nextInt()
    while (n > 0) {
        let digit = n % 10
        if (!compareInt(digit)) {
            return false
        }
        n = parseInt(n / 10)
    }
    return true
}

// { "no": 18, "dat": "", "ans": "" }
const while18 = () => {
    let n = nextInt()
    let counter = 0
    let sum = 0
    while (n > 0) {
        counter++
        sum += n % 10
        n = parseInt(n / 10)
    }
    return compareInt(counter, sum)
}

// { "no": 19, "dat": "", "ans": "" }
const while19 = () => {
    let n = nextInt()
    let chappa = 0
    while (n > 0) {
        chappa = chappa * 10 + n % 10
        n = parseInt(n / 10)
    }
    return compareInt(chappa)
}

// { "no": 20, "dat": "", "ans": "" }
const while20 = () => {
    let n = nextInt()
    let has2 = false
    while (n > 0) {
        if (n % 10 == 2) {
            has2 = true
            break
        }
        n = parseInt(n / 10)
    }
    return compareBoolean(has2)
}

// { "no": 21, "dat": "", "ans": "" }
const while21 = () => {
    let n = nextInt()
    let hasOdd = false
    while (n > 0) {
        if (n % 10 % 2 != 0) {
            hasOdd = true
            break
        }
        n = parseInt(n / 10)
    }
    return compareBoolean(hasOdd)
}

// { "no": 22, "dat": "", "ans": "" }
const while22 = () => {
    let n = nextInt()
    let isPrime = true
    let i = parseInt(Math.sqrt(n))
    while (i > 1) {
        if (n % i == 0) {
            isPrime = false
            break
        }
        i--
    }
    return compareBoolean(isPrime)
}

// { "no": 23, "dat": "", "ans": "" }
const while23 = () => {
    let a = nextInt()
    let b = nextInt()
    while (b != 0) {
        let r = a % b
        a = b
        b = r
    }
    return compareInt(a)
}

// { "no": 24, "dat": "", "ans": "" }
const while24 = () => {
    let n = nextInt()
    let f1 = 1
    let f2 = 1
    let fk
    do {
        fk = f1 + f2
        f1 = f2
        f2 = fk
    } while (fk < n)
    let isFibonacci = fk == n
    return compareBoolean(isFibonacci)
}

// { "no": 25, "dat": "", "ans": "" }
const while25 = () => {
    let n = nextInt()
    let f1 = 1
    let f2 = 1
    let fk
    do {
        fk = f1 + f2
        f1 = f2
        f2 = fk
    } while (fk <= n)
    return compareInt(fk)
}

// { "no": 26, "dat": "", "ans": "" }
const while26 = () => {
    let n = nextInt()
    let f1 = 1
    let f2 = 1
    let fk
    do {
        fk = f1 + f2
        f1 = f2
        f2 = fk
    } while (fk < n)
    fk = f1 + f2
    return compareInt(f1, fk)
}

// { "no": 27, "dat": "", "ans": "" }
const while27 = () => {
    let n = nextInt()
    let f1 = 1
    let f2 = 1
    let fk
    let k = 2
    do {
        fk = f1 + f2
        k++
        f1 = f2
        f2 = fk
    } while (fk < n)
    return compareInt(k)
}

// { "no": 28, "dat": "6", "ans": "6" }
const while28 = () => {
    let e = nextFloat()
    let prev = 2
    let curr = 2 + 1 / prev
    let k = 2
    while (Math.abs(curr - prev) >= e) {
        prev = curr
        curr = 2 + 1 / prev
        k++
    }
    return compareInt(k) && compareFloat(6, prev, curr)
}

// { "no": 29, "dat": "6", "ans": "6" }
const while29 = () => {
    let e = nextFloat()
    let prev = 1
    let curr = 2
    let next = (prev + 2 * curr) / 3
    let k = 3
    while (Math.abs(next - curr) >= e) {
        prev = curr
        curr = next
        next = (prev + 2 * curr) / 3
        k++
    }
    return compareInt(k) && compareFloat(6, curr, next)
}

// { "no": 30, "dat": "2", "ans": "" }
const while30 = () => {
    let a = nextFloat()
    let b = nextFloat()
    let c = nextFloat()
    let cInA = 0
    while (a >= c) {
        a -= c
        cInA++
    }
    let cInB = 0
    while (b >= c) {
        b -= c
        cInB++
    }
    let kvads = 0
    let i = 0
    while (i < cInA) {
        kvads += cInB
        ++i
    }
    return compareInt(kvads)
}