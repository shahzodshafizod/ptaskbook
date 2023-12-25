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
    , roundFixed
} from './io.js'

(async () => {
    for (let taskNo = 1; taskNo <= 40; taskNo++) {
        const taskNoStr = String(taskNo).padStart(3, 0)
        for (let testNo = 1; testNo <= 100; testNo++) {
            const testNoStr = String(testNo).padStart(3, 0)
            const filePath = `../data/kt-gov2/08series/Series${taskNoStr}/${testNoStr}/${testNoStr}`

            await initScanner(filePath + '.dat')
            await initChecker(filePath + '.ans')

            let ok = true
            switch (taskNo) {
                case 1: ok = series1(); break
                case 2: ok = series2(); break
                case 3: ok = series3(); break
                case 4: ok = series4(); break
                case 5: ok = series5(); break
                case 6: ok = series6(); break
                case 7: ok = series7(); break
                case 8: ok = series8(); break
                case 9: ok = series9(); break
                case 10: ok = series10(); break
                case 11: ok = series11(); break
                case 12: ok = series12(); break
                case 13: ok = series13(); break
                case 14: ok = series14(); break
                case 15: ok = series15(); break
                case 16: ok = series16(); break
                case 17: ok = series17(); break
                case 18: ok = series18(); break
                case 19: ok = series19(); break
                case 20: ok = series20(); break
                case 21: ok = series21(); break
                case 22: ok = series22(); break
                case 23: ok = series23(); break
                case 24: ok = series24(); break
                case 25: ok = series25(); break
                case 26: ok = series26(); break
                case 27: ok = series27(); break
                case 28: ok = series28(); break
                case 29: ok = series29(); break
                case 30: ok = series30(); break
                case 31: ok = series31(); break
                case 32: ok = series32(); break
                case 33: ok = series33(); break
                case 34: ok = series34(); break
                case 35: ok = series35(); break
                case 36: ok = series36(); break
                case 37: ok = series37(); break
                case 38: ok = series38(); break
                case 39: ok = series39(); break
                case 40: ok = series40(); break
            }

            if (!ok) {
                console.log("Series" + taskNoStr + " #" + testNoStr + " has NOT been tested!")
                return
            }

            if (!inputsEOF()) {
                console.log("Series" + taskNoStr + " #" + testNoStr + " input NOT EOF!")
                return
            }

            if (!outputsEOF()) {
                console.log("Series" + taskNoStr + " #" + testNoStr + " output NOT EOF!")
                return
            }
        }
        console.log("Series" + taskNoStr + " has been tested!")
    }
    console.log("Series has been tested!")
})()

// { "no": 1, "dat": "2", "ans": "2" }
const series1 = () => {
    let number, sum = 0
    for (let i = 1; i <= 10; i++) {
        number = nextFloat()
        sum += number
    }
    return compareFloat(2, sum)
}

// { "no": 2, "dat": "2", "ans": "2" }
const series2 = () => {
    let number, mult = 1
    for (let i = 1; i <= 10; i++) {
        number = nextFloat()
        mult *= number
    }
    return compareFloat(2, mult)
}

// { "no": 3, "dat": "2", "ans": "2" }
const series3 = () => {
    let number, sum = 0
    for (let i = 1; i <= 10; i++) {
        number = nextFloat()
        sum += number
    }
    let aMean = sum / 10
    return compareFloat(2, aMean)
}

// { "no": 4, "dat": "2", "ans": "2" }
const series4 = () => {
    let n = nextInt()
    let sum = 0
    let mul = 1
    for (let i = 1; i <= n; i++) {
        let number = nextFloat()
        sum += number
        mul *= number
    }
    return compareFloat(2, sum, mul)
}

// { "no": 5, "dat": "2", "ans": "2" }
const series5 = () => {
    let n = nextInt()
    let sum = 0
    for (let i = 1; i <= n; i++) {
        let number = nextFloat()
        let butun = parseInt(number)
        if (!compareFloat(2, butun)) {
            return false
        }
        sum += butun
    }
    return compareFloat(2, sum)
}

// { "no": 6, "dat": "2", "ans": "2{N}, 6" }
const series6 = () => {
    let n = nextInt()
    let mul = 1
    for (let i = 1; i <= n; i++) {
        let number = nextFloat()
        let kasri = number - parseInt(number)
        if (!compareFloat(2, kasri)) {
            return false
        }
        mul *= kasri
    }
    return compareFloat(6, mul)
}

// { "no": 7, "dat": "2", "ans": "" }
const series7 = () => {
    let n = nextInt()
    let sum = 0
    for (let index = 0; index < n; index++) {
        let number = nextFloat()
        let rounded = roundFixed(number, 0)
        if (!compareInt(rounded)) {
            return false
        }
        sum += rounded
    }
    return compareInt(sum)
}

// { "no": 8, "dat": "", "ans": "" }
const series8 = () => {
    let n = nextInt()
    let k = 0
    for (let i = 1; i <= n; i++) {
        let number = nextInt()
        if (number % 2 == 0) {
            if (!compareInt(number)) {
                return false
            }
            k++
        }
    }
    return compareInt(k)
}

// { "no": 9, "dat": "", "ans": "" }
const series9 = () => {
    let n = nextInt()
    let k = 0
    for (let i = 1; i <= n; i++) {
        let number = nextInt()
        if (number % 2 != 0) {
            if (!compareInt(i)) {
                return false
            }
            k++
        }
    }
    return compareInt(k)
}

// { "no": 10, "dat": "", "ans": "" }
const series10 = () => {
    let n = nextInt()
    let hasPositive = false
    for (let i = 1; i <= n; i++) {
        let number = nextInt()
        if (!hasPositive && number > 0) {
            hasPositive = true
        }
    }
    return compareBoolean(hasPositive)
}

// { "no": 11, "dat": "", "ans": "" }
const series11 = () => {
    let k = nextInt()
    let n = nextInt()
    let lessK = false
    for (let i = 1; i <= n; i++) {
        let number = nextInt()
        if (!lessK && number < k) {
            lessK = true
        }
    }
    return compareBoolean(lessK)
}

// { "no": 12, "dat": "", "ans": "" }
const series12 = () => {
    let number, counter = 0
    while (number = nextInt()) {
        counter++
    }
    return compareInt(counter)
}

// { "no": 13, "dat": "", "ans": "" }
const series13 = () => {
    let number
    let sum = 0
    while (number = nextInt()) {
        if ((number > 0) && (number % 2 == 0)) {
            sum += number
        }
    }
    return compareInt(sum)
}

// { "no": 14, "dat": "", "ans": "" }
const series14 = () => {
    let k = nextInt()
    let number
    let counter = 0
    while (number = nextInt()) {
        if (number < k) {
            counter++
        }
    }
    return compareInt(counter)
}

// { "no": 15, "dat": "", "ans": "" }
const series15 = () => {
    let k = nextInt()
    let number, index = 0
    for (let i = 1; number = nextInt(); i++) {
        if (index == 0 && number > k) {
            index = i
        }
    }
    return compareInt(index)
}

// { "no": 16, "dat": "", "ans": "" }
const series16 = () => {
    let k = nextInt()
    let number, index = 0
    for (let i = 1; number = nextInt(); i++) {
        if (number > k) {
            index = i
        }
    }
    return compareInt(index)
}

// { "no": 17, "dat": "2", "ans": "2" }
const series17 = () => {
    let b = nextFloat()
    let n = nextInt()
    let outed = false
    for (let i = 1; i <= n; i++) {
        let number = nextFloat()
        if (!outed && b < number) {
            if (!compareFloat(2, b)) {
                return false
            }
            outed = true
        }
        if (!compareFloat(2, number)) {
            return false
        }
    }
    if (!outed) {
        if (!compareFloat(2, b)) {
            return false
        }
    }
    return true
}

// { "no": 18, "dat": "", "ans": "" }
const series18 = () => {
    let n = nextInt()
    let a
    for (let i = 1; i <= n; i++) {
        let b = nextInt()
        if (i == 1 || b != a) {
            if (!compareInt(b)) {
                return false
            }
        }
        a = b
    }
    return true
}

// { "no": 19, "dat": "", "ans": "" }
const series19 = () => {
    let n = nextInt()
    let a
    let k = 0
    for (let i = 1; i <= n; i++) {
        let b = nextInt()
        if (i != 1 && b < a) {
            if (!compareInt(b)) {
                return false
            }
            k++
        }
        a = b
    }
    return compareInt(k)
}

// { "no": 20, "dat": "", "ans": "" }
const series20 = () => {
    let n = nextInt()
    let a = nextInt()
    let k = 0
    for (let i = 2; i <= n; i++) {
        let b = nextInt()
        if (a < b) {
            if (!compareInt(a)) {
                return false
            }
            k++
        }
        a = b
    }
    return compareInt(k)
}

// { "no": 21, "dat": "2", "ans": "" }
const series21 = () => {
    let n = nextInt()
    let a = nextFloat()
    let incr = true
    for (let i = 2; i <= n; i++) {
        let b = nextFloat()
        if (a > b) {
            incr = false
        }
        a = b
    }
    return compareBoolean(incr)
}

// { "no": 22, "dat": "2", "ans": "" }
const series22 = () => {
    let n = nextInt()
    let a = nextFloat()
    let index = 0
    for (let i = 2; i <= n; i++) {
        let b = nextFloat()
        if (index == 0 && a < b) {
            index = i
        }
        a = b
    }
    return compareInt(index)
}

// { "no": 23, "dat": "2", "ans": "" }
const series23 = () => {
    let n = nextInt()
    let prev = nextFloat()
    let curr = nextFloat()
    let errorIndex = 0
    for (let i = 3; i <= n; i++) {
        let next = nextFloat()
        if ((errorIndex == 0) && !((curr > prev) && (curr > next) || (curr < prev) && (curr < next))) {
            errorIndex = i - 1
        }
        prev = curr
        curr = next
    }
    return compareInt(errorIndex)
}

// { "no": 24, "dat": "", "ans": "" }
const series24 = () => {
    let n = nextInt()
    let sum = 0, lastSum = 0
    for (let i = 1; i <= n; i++) {
        let number = nextInt()
        sum += number
        if (number == 0) {
            lastSum = sum
            sum = 0
        }
    }
    return compareInt(lastSum)
}

// { "no": 25, "dat": "", "ans": "" }
const series25 = () => {
    let n = nextInt()
    let sum = 0, allSum = 0
    let started = false
    for (let i = 1; i <= n; i++) {
        let number = nextInt()
        sum += number
        if (number == 0) {
            if (started == false)
                started = true
            else
                allSum += sum
            sum = 0
        }
    }
    return compareInt(allSum)
}

// { "no": 26, "dat": "2", "ans": "2" }
const series26 = () => {
    let k = nextInt()
    let n = nextInt()
    for (let i = 1; i <= n; i++) {
        let number = nextFloat()
        let degree = 1
        for (let j = 1; j <= k; j++) {
            degree *= number
        }
        if (!compareFloat(2, degree)) {
            return false
        }
    }
    return true
}

// { "no": 27, "dat": "2", "ans": "2" }
const series27 = () => {
    let n = nextInt()
    for (let i = 1; i <= n; i++) {
        let number = nextFloat()
        let degree = 1
        for (let j = 1; j <= i; j++) {
            degree *= number
        }
        if (!compareFloat(2, degree)) {
            return false
        }
    }
    return true
}

// { "no": 28, "dat": "2", "ans": "2" }
const series28 = () => {
    let n = nextInt()
    for (let i = 1; i <= n; i++) {
        let number = nextFloat()
        let degree = 1
        for (let j = i; j <= n; j++) {
            degree *= number
        }
        if (!compareFloat(2, degree)) {
            return false
        }
    }
    return true
}

// { "no": 29, "dat": "", "ans": "" }
const series29 = () => {
    let k = nextInt()
    let n = nextInt()
    let sum = 0
    for (let i = 1; i <= k; i++) {
        for (let j = 1; j <= n; j++) {
            let number = nextInt()
            sum += number
        }
    }
    return compareInt(sum)
}

// { "no": 30, "dat": "", "ans": "" }
const series30 = () => {
    let k = nextInt()
    let n = nextInt()
    for (let i = 1; i <= k; i++) {
        let sum = 0
        for (let j = 1; j <= n; j++) {
            let number = nextInt()
            sum += number
        }
        if (!compareInt(sum)) {
            return false
        }
    }
    return true
}

// { "no": 31, "dat": "", "ans": "" }
const series31 = () => {
    let k = nextInt()
    let n = nextInt()
    let counts = 0
    for (let i = 1; i <= k; i++) {
        let has2 = false
        for (let j = 1; j <= n; j++) {
            let number = nextInt()
            if (number == 2) {
                has2 = true
            }
        }
        counts += +has2
    }
    return compareInt(counts)
}

// { "no": 32, "dat": "", "ans": "" }
const series32 = () => {
    let k = nextInt()
    let n = nextInt()
    for (let i = 1; i <= k; i++) {
        let index = 0
        for (let j = 1; j <= n; j++) {
            let number = nextInt()
            if (index == 0 && number == 2) {
                index = j
            }
        }
        if (!compareInt(index)) {
            return false
        }
    }
    return true
}

// { "no": 33, "dat": "", "ans": "" }
const series33 = () => {
    let k = nextInt()
    let n = nextInt()
    for (let i = 0; i < k; i++) {
        let index = 0
        for (let j = 1; j <= n; j++) {
            let number = nextInt()
            if (number == 2) {
                index = j
            }
        }
        if (!compareInt(index)) {
            return false
        }
    }
    return true
}

// { "no": 34, "dat": "", "ans": "" }
const series34 = () => {
    let k = nextInt()
    let n = nextInt()
    for (let i = 0; i < k; i++) {
        let sum = 0
        let has2 = false
        for (let j = 1; j <= n; j++) {
            let number = nextInt()
            if (number == 2) {
                has2 = true
            }
            sum += number
        }
        if (!compareInt(has2 ? sum : 0)) {
            return false
        }
    }
    return true
}

// { "no": 35, "dat": "", "ans": "" }
const series35 = () => {
    let k = nextInt()
    let totalCounts = 0
    for (let i = 0; i < k; i++) {
        let counts = 0
        do {
            let number = nextInt()
            if (number == 0) {
                break
            }
            counts++
        } while (true)
        if (!compareInt(counts)) {
            return false
        }
        totalCounts += counts
    }
    return compareInt(totalCounts)
}

// { "no": 36, "dat": "", "ans": "" }
const series36 = () => {
    let k = nextInt()
    let progressesCount = 0
    for (let i = 1; i <= k; i++) {
        let isProgress = true
        let prev = nextInt()
        let j = 2
        do {
            let number = nextInt()
            if (number == 0) {
                break
            }
            if (number < prev) {
                isProgress = false
            }
            prev = number
            j++
        } while (true)
        progressesCount += +isProgress
    }
    return compareInt(progressesCount)
}

// { "no": 37, "dat": "", "ans": "" }
const series37 = () => {
    let k = nextInt()
    let counts = 0
    for (let i = 0; i < k; i++) {
        let isProgress = true, isRegress = true
        let prev = nextInt()
        let j = 1
        do {
            j++
            let curr = nextInt()
            if (curr == 0) {
                break
            }
            if (prev > curr) {
                isProgress = false
            }
            if (prev < curr) {
                isRegress = false
            }
            prev = curr
        } while (true)
        counts += +(isProgress || isRegress)
    }
    return compareInt(counts)
}

// { "no": 38, "dat": "", "ans": "" }
const series38 = () => {
    let k = nextInt()
    for (let i = 1; i <= k; i++) {
        let isProgress = true, isRegress = true
        let prev = nextInt()
        let j = 1
        do {
            ++j
            let curr = nextInt()
            if (curr == 0) {
                break
            }
            if (prev < curr) {
                isRegress = false
            } else if (prev > curr) {
                isProgress = false
            }
            prev = curr
        } while (true)
        let result = isProgress ? 1 : isRegress ? -1 : 0
        if (!compareInt(result)) {
            return false
        }
    }
    return true
}

// { "no": 39, "dat": "", "ans": "" }
const series39 = () => {
    let k = nextInt()
    let counts = 0
    for (let i = 1; i <= k; i++) {
        let ok = true
        let prev = nextInt()
        let curr = nextInt()
        let j = 2
        do {
            j++
            let next = nextInt()
            if (next == 0) {
                break
            }
            if ((prev > curr) && (curr > next) || (prev < curr) && (curr < next)) {
                ok = false
            }
            prev = curr
            curr = next
        } while (true)
        counts += +ok
    }
    return compareInt(counts)
}

// { "no": 40, "dat": "", "ans": "" }
const series40 = () => {
    let k = nextInt()
    for (let i = 1; i <= k; i++) {
        let errorIndex = 0
        let counts = 2
        let prev = nextInt()
        let curr = nextInt()
        let j = 2
        do {
            j++
            let next = nextInt()
            if (next == 0) {
                break
            }
            counts++
            if ((errorIndex == 0) && ((prev > curr) && (curr > next) || (prev < curr) && (curr < next))) {
                errorIndex = counts - 1
            }
            prev = curr
            curr = next
        } while (true)
        let result = errorIndex == 0 ? counts : errorIndex
        if (!compareInt(result)) {
            return false
        }
    }
    return true
}