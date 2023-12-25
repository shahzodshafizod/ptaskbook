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
    for (let taskNo = 1; taskNo <= 30; taskNo++) {
        const taskNoStr = String(taskNo).padStart(3, 0)
        for (let testNo = 1; testNo <= 100; testNo++) {
            const testNoStr = String(testNo).padStart(3, 0)
            const filePath = `../data/kt-gov2/10minmax/Minmax${taskNoStr}/${testNoStr}/${testNoStr}`

            await initScanner(filePath + '.dat')
            await initChecker(filePath + '.ans')

            let ok = true
            switch (taskNo) {
                case 1: ok = minmax1(); break
                case 2: ok = minmax2(); break
                case 3: ok = minmax3(); break
                case 4: ok = minmax4(); break
                case 5: ok = minmax5(); break
                case 6: ok = minmax6(); break
                case 7: ok = minmax7(); break
                case 8: ok = minmax8(); break
                case 9: ok = minmax9(); break
                case 10: ok = minmax10(); break
                case 11: ok = minmax11(); break
                case 12: ok = minmax12(); break
                case 13: ok = minmax13(); break
                case 14: ok = minmax14(); break
                case 15: ok = minmax15(); break
                case 16: ok = minmax16(); break
                case 17: ok = minmax17(); break
                case 18: ok = minmax18(); break
                case 19: ok = minmax19(); break
                case 20: ok = minmax20(); break
                case 21: ok = minmax21(); break
                case 22: ok = minmax22(); break
                case 23: ok = minmax23(); break
                case 24: ok = minmax24(); break
                case 25: ok = minmax25(); break
                case 26: ok = minmax26(); break
                case 27: ok = minmax27(); break
                case 28: ok = minmax28(); break
                case 29: ok = minmax29(); break
                case 30: ok = minmax30(); break
            }

            if (!ok) {
                console.log("Minmax" + taskNoStr + " #" + testNoStr + " has NOT been tested!")
                return
            }

            if (!inputsEOF()) {
                console.log("Minmax" + taskNoStr + " #" + testNoStr + " input NOT EOF!")
                return
            }

            if (!outputsEOF()) {
                console.log("Minmax" + taskNoStr + " #" + testNoStr + " output NOT EOF!")
                return
            }
        }
        console.log("Minmax" + taskNoStr + " has been tested!")
    }
    console.log("Minmax has been tested!")
})()

// { "no": 1, "dat": "2", "ans": "2" }
const minmax1 = () => {
    let n = nextInt()
    let minimal, maximal
    for (let i = 1; i <= n; i++) {
        let number = nextFloat()
        if (i == 1) {
            minimal = maximal = number
        } else if (number < minimal) {
            minimal = number
        } else if (number > maximal) {
            maximal = number
        }
    }
    return compareFloat(2, minimal, maximal)
}

// { "no": 2, "dat": "2", "ans": "2" }
const minmax2 = () => {
    let n = nextInt()
    let minS = 0
    for (let i = 1; i <= n; i++) {
        let a = nextFloat()
        let b = nextFloat()
        let s = a * b
        if (i == 1 || s < minS) {
            minS = s
        }
    }
    return compareFloat(2, minS)
}

// { "no": 3, "dat": "2", "ans": "2" }
const minmax3 = () => {
    let n = nextInt()
    let maxP = 0
    for (let i = 1; i <= n; i++) {
        let a = nextFloat()
        let b = nextFloat()
        let p = 2 * (a + b)
        if (i == 1 || p > maxP) {
            maxP = p
        }
    }
    return compareFloat(2, maxP)
}

// { "no": 4, "dat": "2", "ans": "" }
const minmax4 = () => {
    let n = nextInt()
    let minimal = 0, minIndex = 0
    for (let i = 1; i <= n; i++) {
        let number = nextFloat()
        if (i == 1 || number <= minimal) {
            minimal = number
            minIndex = i
        }
    }
    return compareInt(minIndex)
}

// { "no": 5, "dat": "2", "ans": "2" }
const minmax5 = () => {
    let n = nextInt()
    let maxP = 0, maxIndex = 1
    for (let i = 1; i <= n; i++) {
        let m = nextFloat()
        let v = nextFloat()
        let p = m / v
        if (i == 1 || p > maxP) {
            maxP = p
            maxIndex = i
        }
    }
    return compareInt(maxIndex) && compareFloat(2, maxP)
}

// { "no": 6, "dat": "", "ans": "" }
const minmax6 = () => {
    let n = nextInt()
    let minIndex = 0, maxIndex = 0
    let minimal = 0, maximal = 0
    for (let i = 1; i <= n; i++) {
        let number = nextFloat()
        if (i == 1) {
            minimal = maximal = number
            minIndex = maxIndex = i
        } else if (number < minimal) {
            minimal = number
            minIndex = i
        } else if (number >= maximal) {
            maximal = number
            maxIndex = i
        }
    }
    return compareInt(minIndex, maxIndex)
}

// { "no": 7, "dat": "", "ans": "" }
const minmax7 = () => {
    let n = nextInt()
    let maxIndex = 0, minIndex = 0
    let minimal = 0, maximal = 0
    for (let i = 1; i <= n; i++) {
        let number = nextFloat()
        if (i == 1) {
            minimal = maximal = number
            minIndex = maxIndex = i
        } else if (number <= minimal) {
            minimal = number
            minIndex = i
        } else if (number > maximal) {
            maximal = number
            maxIndex = i
        }
    }
    return compareInt(maxIndex, minIndex)
}

// { "no": 8, "dat": "", "ans": "" }
const minmax8 = () => {
    let n = nextInt()
    let firstIndex = 0, lastIndex = 0
    let minimal = 0
    for (let i = 1; i <= n; i++) {
        let number = nextFloat()
        if (i == 1) {
            minimal = number
            firstIndex = lastIndex = i
        } else if (number <= minimal) {
            if (number < minimal) {
                minimal = number
                firstIndex = i
            }
            lastIndex = i
        }
    }
    return compareInt(firstIndex, lastIndex)
}

// { "no": 9, "dat": "", "ans": "" }
const minmax9 = () => {
    let n = nextInt()
    let firstIndex = 0, lastIndex = 0
    let maximal = 0
    for (let i = 1; i <= n; i++) {
        let number = nextFloat()
        if (i == 1) {
            maximal = number
            firstIndex = lastIndex = i
        } else if (number >= maximal) {
            if (number > maximal) {
                maximal = number
                firstIndex = i
            }
            lastIndex = i
        }
    }
    return compareInt(firstIndex, lastIndex)
}

// { "no": 10, "dat": "", "ans": "" }
const minmax10 = () => {
    let n = nextInt()
    let maximal = 0, minimal = 0
    let maxIndex = 0, minIndex = 0
    for (let i = 1; i <= n; i++) {
        let number = nextFloat()
        if (i == 1) {
            maximal = minimal = number
            maxIndex = minIndex = i
        } else if (number > maximal) {
            maximal = number
            maxIndex = i
        } else if (number < minimal) {
            minimal = number
            minIndex = i
        }
    }
    let firstExtremalIndex = minIndex < maxIndex ? minIndex : maxIndex
    return compareInt(firstExtremalIndex)
}

// { "no": 11, "dat": "", "ans": "" }
const minmax11 = () => {
    let n = nextInt()
    let maximal = 0, minimal = 0
    let maxIndex = 0, minIndex = 0
    for (let i = 1; i <= n; i++) {
        let number = nextFloat()
        if (i == 1) {
            maximal = minimal = number
            maxIndex = minIndex = i
        } else if (number >= maximal) {
            maximal = number
            maxIndex = i
        } else if (number <= minimal) {
            minimal = number
            minIndex = i
        }
    }
    let lastExtremalIndex = minIndex > maxIndex ? minIndex : maxIndex
    return compareInt(lastExtremalIndex)
}

// { "no": 12, "dat": "2", "ans": "2" }
const minmax12 = () => {
    let n = nextInt()
    let minimalPositive = 0
    for (let i = 1; i <= n; i++) {
        let number = nextFloat()
        if (number > 0) {
            if (minimalPositive <= 0 || number < minimalPositive) {
                minimalPositive = number
            }
        }
    }
    return compareFloat(2, minimalPositive)
}

// { "no": 13, "dat": "", "ans": "" }
const minmax13 = () => {
    let n = nextInt()
    let maxOdd = 0, index = 0
    for (let i = 1; i <= n; i++) {
        let number = nextFloat()
        if (number % 2 != 0) {
            if (index == 0 || number > maxOdd) {
                maxOdd = number
                index = i
            }
        }
    }
    return compareInt(index)
}

// { "no": 14, "dat": "2", "ans": "2" }
const minmax14 = () => {
    let b = nextFloat()
    let minimal = 0, minIndex = 0
    for (let i = 1; i <= 10; i++) {
        let number = nextFloat()
        if (number > b) {
            if (minIndex == 0 || number < minimal) {
                minimal = number
                minIndex = i
            }
        }
    }
    return compareFloat(2, minimal) && compareInt(minIndex)
}

// { "no": 15, "dat": "2", "ans": "2" }
const minmax15 = () => {
    let b = nextFloat()
    let c = nextFloat()
    let maximal = 0, maxIndex = 0
    for (let i = 1; i <= 10; i++) {
        let number = nextFloat()
        if (b < number && number < c) {
            if (maxIndex == 0) {
                maximal = number
                maxIndex = i
            } else if (number > maximal) {
                maximal = number
                maxIndex = i
            }
        }
    }
    return compareFloat(2, maximal) && compareInt(maxIndex)
}

// { "no": 16, "dat": "", "ans": "" }
const minmax16 = () => {
    let n = nextInt()
    let minimal = 0, minIndex = 0
    for (let i = 1; i <= n; i++) {
        let number = nextInt()
        if (i == 1 || number < minimal) {
            minimal = number
            minIndex = i
        }
    }
    let counter = minIndex - 1
    return compareInt(counter)
}

// { "no": 17, "dat": "", "ans": "" }
const minmax17 = () => {
    let n = nextInt()
    let maximal = 0, maxIndex = 0
    for (let i = 1; i <= n; i++) {
        let number = nextInt()
        if (i == 1 || number >= maximal) {
            maximal = number
            maxIndex = i
        }
    }
    let counter = n - maxIndex
    return compareInt(counter)
}

// { "no": 18, "dat": "", "ans": "" }
const minmax18 = () => {
    let n = nextInt()
    let maximal = 0
    let firstIndex = 0, lastIndex = 0
    for (let i = 1; i <= n; i++) {
        let number = nextInt()
        if (i == 1) {
            maximal = number
            firstIndex = lastIndex = i
        } else if (number >= maximal) {
            if (number > maximal) {
                maximal = number
                firstIndex = i
            }
            lastIndex = i
        }
    }
    let counter = firstIndex == lastIndex ? 0 : lastIndex - firstIndex - 1
    return compareInt(counter)
}

// { "no": 19, "dat": "", "ans": "" }
const minmax19 = () => {
    let n = nextInt()
    let minimal = 0
    let counter = 0
    for (let i = 1; i <= n; i++) {
        let number = nextInt()
        if (i == 1) {
            minimal = number
            counter = 1
        } else if (number <= minimal) {
            if (number < minimal) {
                minimal = number
                counter = 0
            }
            counter++
        }
    }
    return compareInt(counter)
}

// { "no": 20, "dat": "", "ans": "" }
const minmax20 = () => {
    let n = nextInt()
    let minimal, maximal, minsCount = 0, maxsCount = 0
    for (let i = 0; i < n; i++) {
        let number = nextInt()
        if (i == 0) {
            minimal = maximal = number
        }
        if (number == maximal) {
            maxsCount++
        } else if (number > maximal) {
            maximal = number
            maxsCount = 1
        }
        if (number == minimal) {
            minsCount++
        } else if (number < minimal) {
            minimal = number
            minsCount = 1
        }
    }
    let externalsCount = minsCount + maxsCount
    if (minimal == maximal) {
        externalsCount = minsCount
    }
    return compareInt(externalsCount)
}

// { "no": 21, "dat": "2", "ans": "2" }
const minmax21 = () => {
    let n = nextInt()
    let number = nextFloat()
    let maximal = number, minimal = number
    let sum = number
    for (let i = 2; i <= n; i++) {
        number = nextFloat()
        sum += number
        if (number > maximal) {
            maximal = number
        } else if (number < minimal) {
            minimal = number
        }
    }
    sum -= minimal + maximal
    let aMean = sum / (n - 2)
    return compareFloat(2, aMean)
}

class TNumber {
    constructor() {
        this.value = 0
    }
}

// { "no": 22, "dat": "2", "ans": "2" }
const minmax22 = () => {
    let n = nextInt()
    let min1 = new TNumber()
    min1.value = nextFloat()
    let min2 = new TNumber()
    min2.value = nextFloat()
    minMax2(min1, min2)
    for (let i = 3; i <= n; i++) {
        let number = nextFloat()
        if (number < min2.value) {
            min2.value = number
            minMax2(min1, min2)
        }
    }
    return compareFloat(2, min1.value, min2.value)
}

const minMax2 = (a, b) => {
    if (a.value > b.value) {
        a.value += b.value
        b.value = a.value - b.value
        a.value -= b.value
    }
}

// { "no": 23, "dat": "2", "ans": "2" }
const minmax23 = () => {
    let n = nextInt()
    let max1 = new TNumber()
    max1.value = nextFloat()
    let max2 = new TNumber()
    max2.value = nextFloat()
    let max3 = new TNumber()
    max3.value = nextFloat()
    minMax3(max1, max2, max3)
    for (let i = 4; i <= n; i++) {
        let number = nextFloat()
        if (number > max1.value) {
            max1.value = number
            minMax3(max1, max2, max3)
        }
    }
    return compareFloat(2, max3.value, max2.value, max1.value)
}

const minMax3 = (a, b, c) => {
    minMax2(a, b)
    minMax2(a, c)
    minMax2(b, c)
}

// { "no": 24, "dat": "2", "ans": "2" }
const minmax24 = () => {
    let n = nextInt()
    let a = nextFloat()
    let maxSum = 0
    for (let i = 2; i <= n; i++) {
        let b = nextFloat()
        if (i == 2 || a + b > maxSum) {
            maxSum = a + b
        }
        a = b
    }
    return compareFloat(2, maxSum)
}

// { "no": 25, "dat": "2", "ans": "" }
const minmax25 = () => {
    let n = nextInt()
    let a = nextFloat()
    let minMult = 0
    let index = 0
    for (let i = 2; i <= n; i++) {
        let b = nextFloat()
        if (i == 2 || a * b < minMult) {
            minMult = a * b
            index = i
        }
        a = b
    }
    return compareInt(index - 1, index)
}

// { "no": 26, "dat": "", "ans": "" }
const minmax26 = () => {
    let n = nextInt()
    let counts = 0, prevIndex = 0, maxCounts = 0
    for (let i = 1; i <= n; i++) {
        let number = nextInt()
        if (number % 2 == 0) {
            if (prevIndex == 0) {
                counts = 1
            } else if (i - 1 == prevIndex) {
                counts++
            } else {
                if (maxCounts == 0 || counts > maxCounts) {
                    maxCounts = counts
                }
                counts = 1
            }
            prevIndex = i
        }
    }
    if (maxCounts == 0 || counts > maxCounts) {
        maxCounts = counts
    }
    return compareInt(maxCounts)
}

// { "no": 27, "dat": "", "ans": "" }
const minmax27 = () => {
    let n = nextInt()
    let prev
    let counts = 0, maxCounts = 0, index = 0
    for (let i = 1; i <= n; i++) {
        let curr = nextFloat()
        if (i == 1) {
            counts = index = i
        } else if (curr == prev) {
            counts++
        } else {
            if (maxCounts == 0) {
                maxCounts = counts
                index = i - maxCounts
            } else if (counts > maxCounts) {
                maxCounts = counts
                index = i - maxCounts
            }
            counts = 1
        }
        prev = curr
    }
    if (maxCounts == 0 || counts > maxCounts) {
        maxCounts = counts
        index = n - maxCounts + 1
    }
    return compareInt(index, maxCounts)
}

// { "no": 28, "dat": "", "ans": "" }
const minmax28 = () => {
    let n = nextInt()
    let counts = 0, maxCountsIndex = 0, maxCounts = 0
    for (let i = 1; i <= n; i++) {
        let number = nextFloat()
        if (number == 1) {
            counts++
        } else {
            if (counts == 0)
                continue
            if (maxCounts == 0) {
                maxCounts = counts
                maxCountsIndex = i - counts
            } else if (counts >= maxCounts) {
                maxCounts = counts
                maxCountsIndex = i - counts
            }
            counts = 0
        }
    }
    if (counts != 0) {
        if (maxCounts == 0 || counts >= maxCounts) {
            maxCounts = counts
            maxCountsIndex = n - counts + 1
        }
    }
    return compareInt(maxCountsIndex, maxCounts)
}

// { "no": 29, "dat": "", "ans": "" }
const minmax29 = () => {
    let n = nextInt()
    let minimal = 0, counts = 0, maxCounts = 0
    for (let i = 1; i <= n; i++) {
        let number = nextFloat()
        if (i == 1) {
            counts = 1
            minimal = number
        } else if (number == minimal) {
            counts++
        } else if (number < minimal) {
            minimal = number
            counts = 1
            maxCounts = 0
        } else {
            if (maxCounts == 0 || counts > maxCounts) {
                maxCounts = counts
            }
            counts = 0
        }
    }
    if (maxCounts == 0 || counts > maxCounts) {
        maxCounts = counts
    }
    return compareInt(maxCounts)
}

// { "no": 30, "dat": "", "ans": "" }
const minmax30 = () => {
    let n = nextInt()
    let number = 0, maximal = 0, counts = 0, maxCounts = 0
    for (let i = 1; i <= n; i++) {
        number = nextFloat()
        if (i == 1) {
            counts = 1
            maximal = number
        } else if (number == maximal) {
            counts++
        } else if (number > maximal) {
            maximal = number
            counts = 1
            maxCounts = 0
        } else {
            if (counts != 0 && (maxCounts == 0 || counts < maxCounts)) {
                maxCounts = counts
            }
            counts = 0
        }
    }
    if (maxCounts == 0 || (counts != 0 && counts < maxCounts)) {
        maxCounts = counts
    }
    return compareInt(maxCounts)
}