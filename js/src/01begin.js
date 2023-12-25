import {
    nextFloat
    , compareFloat
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
            const filePath = `../data/kt-gov2/01begin/Begin${taskNoStr}/${testNoStr}/${testNoStr}`

            await initScanner(filePath + '.dat')
            await initChecker(filePath + '.ans')

            let ok = true
            switch (taskNo) {
                case 1: ok = begin1(); break
                case 2: ok = begin2(); break
                case 3: ok = begin3(); break
                case 4: ok = begin4(); break
                case 5: ok = begin5(); break
                case 6: ok = begin6(); break
                case 7: ok = begin7(); break
                case 8: ok = begin8(); break
                case 9: ok = begin9(); break
                case 10: ok = begin10(); break
                case 11: ok = begin11(); break
                case 12: ok = begin12(); break
                case 13: ok = begin13(); break
                case 14: ok = begin14(); break
                case 15: ok = begin15(); break
                case 16: ok = begin16(); break
                case 17: ok = begin17(); break
                case 18: ok = begin18(); break
                case 19: ok = begin19(); break
                case 20: ok = begin20(); break
                case 21: ok = begin21(); break
                case 22: ok = begin22(); break
                case 23: ok = begin23(); break
                case 24: ok = begin24(); break
                case 25: ok = begin25(); break
                case 26: ok = begin26(); break
                case 27: ok = begin27(); break
                case 28: ok = begin28(); break
                case 29: ok = begin29(); break
                case 30: ok = begin30(); break
                case 31: ok = begin31(); break
                case 32: ok = begin32(); break
                case 33: ok = begin33(); break
                case 34: ok = begin34(); break
                case 35: ok = begin35(); break
                case 36: ok = begin36(); break
                case 37: ok = begin37(); break
                case 38: ok = begin38(); break
                case 39: ok = begin39(); break
                case 40: ok = begin40(); break
            }

            if (!ok) {
                console.log("Begin" + taskNoStr + " #" + testNoStr + " has NOT been tested!")
                return
            }

            if (!inputsEOF()) {
                console.log("Begin" + taskNoStr + " #" + testNoStr + " output NOT EOF!")
                return
            }

            if (!outputsEOF()) {
                console.log("Begin" + taskNoStr + " #" + testNoStr + " output NOT EOF!")
                return
            }
        }
        console.log("Begin" + taskNoStr + " has been tested!")
    }
    console.log("Begin has been tested!")
})()

// { "no": 1, "dat": "2", "ans": "2" }
const begin1 = () => {
    let a = nextFloat()
    const p = 4 * a
    return compareFloat(2, p)
}

// { "no": 2, "dat": "2", "ans": "2" }
const begin2 = () => {
    let a = nextFloat()
    let s = a * a
    return compareFloat(2, s)
}

// { "no": 3, "dat": "2", "ans": "2" }
const begin3 = () => {
    let a = nextFloat()
    let b = nextFloat()
    let s = a * b
    let p = 2 * (a + b)
    return compareFloat(2, s, p)
}

// { "no": 4, "dat": "2", "ans": "2" }
const begin4 = () => {
    let d = nextFloat()
    const PI = 3.14
    let l = PI * d
    return compareFloat(2, l)
}

// { "no": 5, "dat": "2", "ans": "3" }
const begin5 = () => {
    let a = nextFloat()
    let v = a * a * a
    let s = 6 * a * a
    return compareFloat(3, v, s)
}

// { "no": 6, "dat": "2", "ans": "3" }
const begin6 = () => {
    let a = nextFloat()
    let b = nextFloat()
    let c = nextFloat()
    let v = a * b * c
    let s = 2 * (a * b + b * c + c * a)
    return compareFloat(3, v, s)
}

// { "no": 7, "dat": "2", "ans": "3" }
const begin7 = () => {
    let r = nextFloat()
    const PI = 3.14
    let l = 2 * PI * r
    let s = PI * r * r
    return compareFloat(3, l, s)
}

// { "no": 8, "dat": "2", "ans": "2" }
const begin8 = () => {
    let a = nextFloat()
    let b = nextFloat()
    let average = (a + b) / 2
    return compareFloat(2, average)
}

// { "no": 9, "dat": "2", "ans": "2" }
const begin9 = () => {
    let a = nextFloat()
    let b = nextFloat()
    let gmean = Math.sqrt(a * b)
    return compareFloat(2, gmean)
}

// { "no": 10, "dat": "2", "ans": "2" }
const begin10 = () => {
    let a = nextFloat()
    let b = nextFloat()
    a *= a
    b *= b
    let sum = a + b
    let dif = a - b
    let pro = a * b
    let div = a / b
    return compareFloat(2, sum, dif, pro, div)
}

// { "no": 11, "dat": "2", "ans": "2" }
const begin11 = () => {
    let a = nextFloat()
    let b = nextFloat()
    a = Math.abs(a)
    b = Math.abs(b)
    let sum = a + b
    let dif = a - b
    let pro = a * b
    let div = a / b
    return compareFloat(2, sum, dif, pro, div)
}

// { "no": 12, "dat": "2", "ans": "2" }
const begin12 = () => {
    let a = nextFloat()
    let b = nextFloat()
    let c = Math.sqrt(a * a + b * b)
    let p = a + b + c
    return compareFloat(2, c, p)
}

// { "no": 13, "dat": "2", "ans": "2" }
const begin13 = () => {
    let r1 = nextFloat()
    let r2 = nextFloat()
    const PI = 3.14
    let s1 = PI * r1 * r1
    let s2 = PI * r2 * r2
    let s3 = s1 - s2
    return compareFloat(2, s1, s2, s3)
}

// { "no": 14, "dat": "2", "ans": "2" }
const begin14 = () => {
    let l = nextFloat()
    const PI = 3.14
    let r = l / (2 * PI)
    let s = PI * r * r
    return compareFloat(2, r, s)
}

// { "no": 15, "dat": "2", "ans": "2" }
const begin15 = () => {
    let s = nextFloat()
    const PI = 3.14
    let d = Math.sqrt(4 * s / PI)
    let l = PI * d
    return compareFloat(2, d, l)
}

// { "no": 16, "dat": "2", "ans": "2" }
const begin16 = () => {
    let x1 = nextFloat()
    let x2 = nextFloat()
    let distance = Math.abs(x2 - x1)
    return compareFloat(2, distance)
}

// { "no": 17, "dat": "2", "ans": "2" }
const begin17 = () => {
    let a = nextFloat()
    let b = nextFloat()
    let c = nextFloat()
    let ac = Math.abs(c - a)
    let bc = Math.abs(c - b)
    let sum = ac + bc
    return compareFloat(2, ac, bc, sum)
}

// { "no": 18, "dat": "2", "ans": "2" }
const begin18 = () => {
    let a = nextFloat()
    let b = nextFloat()
    let c = nextFloat()
    let ac = Math.abs(c - a)
    let bc = Math.abs(b - c)
    let pro = ac * bc
    return compareFloat(2, pro)
}

// { "no": 19, "dat": "2", "ans": "2" }
const begin19 = () => {
    let x1 = nextFloat()
    let y1 = nextFloat()
    let x2 = nextFloat()
    let y2 = nextFloat()
    let a = Math.abs(x2 - x1)
    let b = Math.abs(y2 - y1)
    let p = 2 * (a + b)
    let s = a * b
    return compareFloat(2, p, s)
}

// { "no": 20, "dat": "2", "ans": "2" }
const begin20 = () => {
    let x1 = nextFloat()
    let y1 = nextFloat()
    let x2 = nextFloat()
    let y2 = nextFloat()
    let distance = Math.sqrt(Math.pow(x2 - x1, 2) + Math.pow(y2 - y1, 2))
    return compareFloat(2, distance)
}

// { "no": 21, "dat": "2", "ans": "2" }
const begin21 = () => {
    let x1 = nextFloat()
    let y1 = nextFloat()
    let x2 = nextFloat()
    let y2 = nextFloat()
    let x3 = nextFloat()
    let y3 = nextFloat()
    let a = Math.sqrt(Math.pow(x2 - x1, 2) + Math.pow(y2 - y1, 2)); // 39,69+30,25=69,94
    let b = Math.sqrt(Math.pow(x3 - x2, 2) + Math.pow(y3 - y2, 2)); // 72,25+1,96=74,21
    let c = Math.sqrt(Math.pow(x3 - x1, 2) + Math.pow(y3 - y1, 2)); // 4,84+16,81=21,65
    let p = a + b + c; // 69,94+74,21+21,65=165,8
    let halfP = p / 2
    let s = Math.sqrt(halfP * (halfP - a) * (halfP - b) * (halfP - c))
    return compareFloat(2, p, s)
}

// { "no": 22, "dat": "2", "ans": "2" }
const begin22 = () => {
    let a = nextFloat()
    let b = nextFloat()
    a += b
    b = a - b
    a -= b
    return compareFloat(2, a, b)
}

// { "no": 23, "dat": "2", "ans": "2" }
const begin23 = () => {
    let a = nextFloat()
    let b = nextFloat()
    let c = nextFloat()
    let temp = c
    c = b
    b = a
    a = temp
    return compareFloat(2, a, b, c)
}

// { "no": 24, "dat": "2", "ans": "2" }
const begin24 = () => {
    let a = nextFloat()
    let b = nextFloat()
    let c = nextFloat()
    let temp = a
    a = b
    b = c
    c = temp
    return compareFloat(2, a, b, c)
}

// { "no": 25, "dat": "2", "ans": "2" }
const begin25 = () => {
    let x = nextFloat()
    let y = 3 * Math.pow(x, 6) - 6 * Math.pow(x, 2) - 7
    return compareFloat(2, y)
}

// { "no": 26, "dat": "2", "ans": "2" }
const begin26 = () => {
    let x = nextFloat()
    let thirdDegree = Math.pow(x - 3, 3)
    let y = 4 * thirdDegree * thirdDegree - 7 * thirdDegree + 2
    return compareFloat(2, y)
}

// { "no": 27, "dat": "2", "ans": "2" }
const begin27 = () => {
    let a = nextFloat()
    let degree = a * a
    if (!compareFloat(2, degree)) {
        return false
    }
    degree *= degree
    if (!compareFloat(2, degree)) {
        return false
    }
    degree *= degree
    return compareFloat(2, degree)
}

// { "no": 28, "dat": "2", "ans": "2" }
const begin28 = () => {
    let a = nextFloat()
    let degree1 = a * a
    if (!compareFloat(2, degree1)) {
        return false
    }
    let degree2 = degree1 * a
    if (!compareFloat(2, degree2)) {
        return false
    }
    degree1 *= degree2
    if (!compareFloat(2, degree1)) {
        return false
    }
    degree2 = degree1 * degree1
    if (!compareFloat(2, degree2)) {
        return false
    }
    degree1 *= degree2
    return compareFloat(2, degree1)
}

// { "no": 29, "dat": "2", "ans": "2" }
const begin29 = () => {
    let degrees = nextFloat()
    const PI = 3.14
    let radians = degrees * PI / 180
    return compareFloat(2, radians)
}

// { "no": 30, "dat": "2", "ans": "2" }
const begin30 = () => {
    let radians = nextFloat()
    const PI = 3.14
    let degrees = radians * 180 / PI
    return compareFloat(2, degrees)
}

// { "no": 31, "dat": "2", "ans": "2" }
const begin31 = () => {
    let tf = nextFloat()
    let tc = (tf - 32) / 9 * 5
    return compareFloat(2, tc)
}

// { "no": 32, "dat": "2", "ans": "2" }
const begin32 = () => {
    let tc = nextFloat()
    let tf = tc / 5 * 9 + 32
    return compareFloat(2, tf)
}

// { "no": 33, "dat": "3,2,3", "ans": "2" }
const begin33 = () => {
    let x = nextFloat()
    let a = nextFloat()
    let y = nextFloat()
    let price = a / x
    let yKgPrice = price * y
    return compareFloat(2, price, yKgPrice)
}

// { "no": 34, "dat": "3,2,3,2", "ans": "2" }
const begin34 = () => {
    let x = nextFloat()
    let a = nextFloat()
    let y = nextFloat()
    let b = nextFloat()
    let chocoPrice = a / x
    let candyPrice = b / y
    let expensiver = chocoPrice / candyPrice
    return compareFloat(2, chocoPrice, candyPrice, expensiver)
}

// { "no": 35, "dat": "2", "ans": "2" }
const begin35 = () => {
    let v = nextFloat()
    let u = nextFloat()
    let t1 = nextFloat()
    let t2 = nextFloat()
    let s = t1 * v + t2 * (v - u)
    return compareFloat(2, s)
}

// { "no": 36, "dat": "2", "ans": "2" }
const begin36 = () => {
    let v1 = nextFloat()
    let v2 = nextFloat()
    let s = nextFloat()
    let t = nextFloat()
    s += t * (v1 + v2)
    return compareFloat(2, s)
}

// { "no": 37, "dat": "2", "ans": "2" }
const begin37 = () => {
    let v1 = nextFloat()
    let v2 = nextFloat()
    let s = nextFloat()
    let t = nextFloat()
    s = Math.abs(s - t * (v1 + v2))
    return compareFloat(2, s)
}

// { "no": 38, "dat": "2", "ans": "2" }
const begin38 = () => {
    let a = nextFloat()
    let b = nextFloat()
    let x = -b / a
    return compareFloat(2, x)
}

// { "no": 39, "dat": "2", "ans": "2" }
const begin39 = () => {
    let a = nextFloat()
    let b = nextFloat()
    let c = nextFloat()
    let dRoot = Math.sqrt(b * b - 4 * a * c)
    let x1 = (-b - dRoot) / 2 / a
    let x2 = (-b + dRoot) / 2 / a
    return compareFloat(2, x1, x2)
}

// { "no": 40, "dat": "2", "ans": "2" }
const begin40 = () => {
    let a1 = nextFloat()
    let b1 = nextFloat()
    let c1 = nextFloat()
    let a2 = nextFloat()
    let b2 = nextFloat()
    let c2 = nextFloat()
    let d = a1 * b2 - a2 * b1
    let x = (c1 * b2 - c2 * b1) / d
    let y = (a1 * c2 - a2 * c1) / d
    return compareFloat(2, x, y)
}