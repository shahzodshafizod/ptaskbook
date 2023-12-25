import {
    nextInt
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
            const filePath = `../data/kt-gov2/02integer/Integer${taskNoStr}/${testNoStr}/${testNoStr}`

            await initScanner(filePath + '.dat')
            await initChecker(filePath + '.ans')

            let ok = true
            switch (taskNo) {
                case 1: ok = integer1(); break
                case 2: ok = integer2(); break
                case 3: ok = integer3(); break
                case 4: ok = integer4(); break
                case 5: ok = integer5(); break
                case 6: ok = integer6(); break
                case 7: ok = integer7(); break
                case 8: ok = integer8(); break
                case 9: ok = integer9(); break
                case 10: ok = integer10(); break
                case 11: ok = integer11(); break
                case 12: ok = integer12(); break
                case 13: ok = integer13(); break
                case 14: ok = integer14(); break
                case 15: ok = integer15(); break
                case 16: ok = integer16(); break
                case 17: ok = integer17(); break
                case 18: ok = integer18(); break
                case 19: ok = integer19(); break
                case 20: ok = integer20(); break
                case 21: ok = integer21(); break
                case 22: ok = integer22(); break
                case 23: ok = integer23(); break
                case 24: ok = integer24(); break
                case 25: ok = integer25(); break
                case 26: ok = integer26(); break
                case 27: ok = integer27(); break
                case 28: ok = integer28(); break
                case 29: ok = integer29(); break
                case 30: ok = integer30(); break
            }

            if (!ok) {
                console.log("Integer" + taskNoStr + " #" + testNoStr + " has NOT been tested!")
                return
            }

            if (!inputsEOF()) {
                console.log("Integer" + taskNoStr + " #" + testNoStr + " input NOT EOF!")
                return
            }

            if (!outputsEOF()) {
                console.log("Integer" + taskNoStr + " #" + testNoStr + " output NOT EOF!")
                return
            }
        }
        console.log("Integer" + taskNoStr + " has been tested!")
    }
    console.log("Integer has been tested!")
})()

const integer1 = () => {
    let l = nextInt()
    let meters = parseInt(l / 100)
    return compareInt(meters)
}

const integer2 = () => {
    let m = nextInt()
    let tons = parseInt(m / 1000)
    return compareInt(tons)
}

const integer3 = () => {
    let bytes = nextInt()
    let kBytes = parseInt(bytes / 1024)
    return compareInt(kBytes)
}

const integer4 = () => {
    let a = nextInt()
    let b = nextInt()
    let segmentsCount = parseInt(a / b)
    return compareInt(segmentsCount)
}

const integer5 = () => {
    let a = nextInt()
    let b = nextInt()
    let unusedSegment = a % b
    return compareInt(unusedSegment)
}

const integer6 = () => {
    let number = nextInt()
    let tens = parseInt(number / 10)
    let ones = number % 10
    return compareInt(tens, ones)
}

const integer7 = () => {
    let number = nextInt()
    let tens = parseInt(number / 10)
    let ones = number % 10
    let sum = tens + ones
    let pro = tens * ones
    return compareInt(sum, pro)
}

const integer8 = () => {
    let number = nextInt()
    let tens = parseInt(number / 10)
    let ones = number % 10
    let rebmun = ones * 10 + tens
    return compareInt(rebmun)
}

const integer9 = () => {
    let number = nextInt()
    let hundreds = parseInt(number / 100)
    return compareInt(hundreds)
}

const integer10 = () => {
    let number = nextInt()
    let ones = number % 10
    let tens = parseInt(number / 10) % 10
    return compareInt(ones, tens)
}

const integer11 = () => {
    let number = nextInt()
    let hundreds = parseInt(number / 100)
    let tens = parseInt(number / 10) % 10
    let ones = number % 10
    let sum = hundreds + tens + ones
    let pro = hundreds * tens * ones
    return compareInt(sum, pro)
}

const integer12 = () => {
    let number = nextInt()
    let rebmun = (number % 10) * 100 + (parseInt(number / 10) % 10) * 10 + parseInt(number / 100)
    return compareInt(rebmun)
}

const integer13 = () => {
    let number = nextInt()
    number = (number % 100) * 10 + parseInt(number / 100)
    return compareInt(number)
}

const integer14 = () => {
    let number = nextInt()
    number = number % 10 * 100 + parseInt(number / 10)
    return compareInt(number)
}

const integer15 = () => {
    let number = nextInt()
    number = parseInt(number / 10) % 10 * 100 + parseInt(number / 100) * 10 + number % 10
    return compareInt(number)
}

const integer16 = () => {
    let number = nextInt()
    number = number - number % 100 + number % 10 * 10 + parseInt(number / 10) % 10
    return compareInt(number)
}

const integer17 = () => {
    let number = nextInt()
    let hundreds = parseInt(number / 100) % 10
    return compareInt(hundreds)
}

const integer18 = () => {
    let number = nextInt()
    let thousands = parseInt(number / 1000) % 10
    return compareInt(thousands)
}

const integer19 = () => {
    let n = nextInt()
    let minutes = parseInt(n / 60)
    return compareInt(minutes)
}

const integer20 = () => {
    let n = nextInt()
    let hours = parseInt(n / 3600)
    return compareInt(hours)
}

const integer21 = () => {
    let n = nextInt()
    let seconds = n % 60
    return compareInt(seconds)
}

const integer22 = () => {
    let n = nextInt()
    let seconds = n % 3600
    return compareInt(seconds)
}

const integer23 = () => {
    let n = nextInt()
    let minutes = parseInt(n % 3600 / 60)
    return compareInt(minutes)
}

const integer24 = () => {
    let k = nextInt()
    let weekNo = k % 7
    return compareInt(weekNo)
}

const integer25 = () => {
    let k = nextInt()
    let weekNo = (k + 3) % 7
    return compareInt(weekNo)
}

const integer26 = () => {
    let k = nextInt()
    let weekNo = k % 7 + 1
    return compareInt(weekNo)
}

const integer27 = () => {
    let k = nextInt()
    let weekNo = (k + 4) % 7 + 1
    return compareInt(weekNo)
}

const integer28 = () => {
    let k = nextInt()
    let n = nextInt()
    let weekNo = (k + n - 2) % 7 + 1
    return compareInt(weekNo)
}

const integer29 = () => {
    let a = nextInt()
    let b = nextInt()
    let c = nextInt()
    let cInA = parseInt(a / c)
    let cInB = parseInt(b / c)
    let squaresCount = cInA * cInB
    let unusedPart = a * b - squaresCount * c * c
    return compareInt(squaresCount, unusedPart)
}

const integer30 = () => {
    let year = nextInt()
    let century = parseInt((year - 1) / 100) + 1
    return compareInt(century)
}