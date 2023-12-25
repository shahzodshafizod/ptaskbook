import {
    nextFloat
    , compareFloat
    , nextInt
    , compareInt
    , nextChar
    , compareChar
    , compareWord
    , compareLine
    , initScanner
    , initChecker
    , inputsEOF
    , outputsEOF
} from './io.js'

(async () => {
    for (let taskNo = 1; taskNo <= 20; taskNo++) {
        const taskNoStr = String(taskNo).padStart(3, 0)
        for (let testNo = 1; testNo <= 100; testNo++) {
            const testNoStr = String(testNo).padStart(3, 0)
            const filePath = `../data/kt-gov2/05case/Case${taskNoStr}/${testNoStr}/${testNoStr}`

            await initScanner(filePath + '.dat')
            await initChecker(filePath + '.ans')

            let ok = true
            switch (taskNo) {
                case 1: ok = case1(); break
                case 2: ok = case2(); break
                case 3: ok = case3(); break
                case 4: ok = case4(); break
                case 5: ok = case5(); break
                case 6: ok = case6(); break
                case 7: ok = case7(); break
                case 8: ok = case8(); break
                case 9: ok = case9(); break
                case 10: ok = case10(); break
                case 11: ok = case11(); break
                case 12: ok = case12(); break
                case 13: ok = case13(); break
                case 14: ok = case14(); break
                case 15: ok = case15(); break
                case 16: ok = case16(); break
                case 17: ok = case17(); break
                case 18: ok = case18(); break
                case 19: ok = case19(); break
                case 20: ok = case20(); break
            }

            if (!ok) {
                console.log("Case" + taskNoStr + " #" + testNoStr + " has NOT been tested!")
                return
            }

            if (!inputsEOF()) {
                console.log("Case" + taskNoStr + " #" + testNoStr + " input NOT EOF!")
                return
            }

            if (!outputsEOF()) {
                console.log("Case" + taskNoStr + " #" + testNoStr + " output NOT EOF!")
                return
            }
        }
        console.log("Case" + taskNoStr + " has been tested!")
    }
    console.log("Case has been tested!")
})()

// { "no": 1, "dat": "", "ans": "" }
const case1 = () => {
    let weekDayNo = nextInt()
    let weekDayName = ""
    switch (weekDayNo) {
        case 1: weekDayName = "Monday"; break
        case 2: weekDayName = "Tuesday"; break
        case 3: weekDayName = "Wednesday"; break
        case 4: weekDayName = "Thursday"; break
        case 5: weekDayName = "Friday"; break
        case 6: weekDayName = "Saturday"; break
        case 7: weekDayName = "Sunday"; break
    }
    return compareWord(weekDayName)
}

// { "no": 2, "dat": "", "ans": "" }
const case2 = () => {
    let k = nextInt()
    let mark = ""
    switch (k) {
        case 1: mark = "bad"; break
        case 2: mark = "unsatisfactory"; break
        case 3: mark = "mediocre"; break
        case 4: mark = "good"; break
        case 5: mark = "excellent"; break
        default: mark = "error"; break
    }
    return compareWord(mark)
}

// { "no": 3, "dat": "", "ans": "" }
const case3 = () => {
    let monthNo = nextInt()
    let season = ""
    switch (monthNo) {
        case 1:
        case 2:
        case 12: season = "Winter"; break
        case 3:
        case 4:
        case 5: season = "Spring"; break
        case 6:
        case 7:
        case 8: season = "Summer"; break
        case 9:
        case 10:
        case 11: season = "Autumn"; break
    }
    return compareWord(season)
}

// { "no": 4, "dat": "", "ans": "" }
const case4 = () => {
    let monthNo = nextInt()
    let monthDays = 0
    switch (monthNo) {
        case 1:
        case 3:
        case 5:
        case 7:
        case 8:
        case 10:
        case 12: monthDays = 31; break
        case 4:
        case 6:
        case 9:
        case 11: monthDays = 30; break
        case 2: monthDays = 28; break
    }
    return compareInt(monthDays)
}

// { "no": 5, "dat": "2", "ans": "2" }
const case5 = () => {
    let n = nextInt()
    let a = nextFloat()
    let b = nextFloat()
    let result = 0
    switch (n) {
        case 1: result = a + b; break
        case 2: result = a - b; break
        case 3: result = a * b; break
        case 4: result = a / b; break
    }
    return compareFloat(2, result)
}

// { "no": 6, "dat": "2", "ans": "5" }
const case6 = () => {
    let n = nextInt()
    let l = nextFloat()
    let meters = 0
    switch (n) {
        case 1: meters = l / 10; break
        case 2: meters = l * 1000; break
        case 3: meters = l; break
        case 4: meters = l / 1000; break
        case 5: meters = l / 100; break
    }
    return compareFloat(5, meters)
}

// { "no": 7, "dat": "2", "ans": "8" }
const case7 = () => {
    let n = nextInt()
    let m = nextFloat()
    let kg = 0
    switch (n) {
        case 1: kg = m; break
        case 2: kg = m / 1000000; break
        case 3: kg = m / 1000; break
        case 4: kg = m * 1000; break
        case 5: kg = m * 100; break
    }
    return compareFloat(8, kg)
}

// { "no": 8, "dat": "", "ans": "" }
const case8 = () => {
    let d = nextInt()
    let m = nextInt()
    switch (d) {
        case 1:
            switch (m) {
                case 1:
                    d = 31
                    m = 13
                    break
                case 3:
                    d = 28
                    break
                case 2: case 4: case 6: case 8: case 9: case 11:
                    d = 31
                    break
                case 5: case 7: case 10: case 12:
                    d = 30
                    break
            }
            m--
            break
        default:
            d--
            break
    }
    return compareInt(d, m)
}

// { "no": 9, "dat": "", "ans": "" }
const case9 = () => {
    let d = nextInt()
    let m = nextInt()
    switch (d) {
        case 31:
            d = 0
            switch (m) {
                case 12: m = 0; break
            }
            break
        case 30:
            switch (m) {
                case 4: case 6: case 9: case 11:
                    d = 0
                    break
            }
            break
        case 28:
            switch (m) {
                case 2: d = 0; break
            }
            break
    }
    if (d == 0) {
        m++
    }
    d++
    return compareInt(d, m)
}

// { "no": 10, "dat": "", "ans": "" }
const case10 = () => {
    let c = nextChar()
    let n = nextInt()
    switch (n) {
        case 1:
            switch (c) {
                case 'N': c = 'W'; break
                case 'E': c = 'N'; break
                case 'S': c = 'E'; break
                case 'W': c = 'S'; break
            }
            break
        case -1:
            switch (c) {
                case 'N': c = 'E'; break
                case 'E': c = 'S'; break
                case 'S': c = 'W'; break
                case 'W': c = 'N'; break
            }
            break
    }
    return compareChar(c)
}

// { "no": 11, "dat": "", "ans": "" }
const case11 = () => {
    let c = nextChar()
    let n1 = nextInt()
    let n2 = nextInt()
    /*
    1+1 = 2 (180)
    -1-1 = -2 (180)

    1-1 = 0 (cont)
    -1+1 = 0 (cont)
    2+2 = 4 (cont)

    1+2 = 3 (right)
    2+1 = 3 (right)

    -1+2 = 1 (left)
    2-1 = 1 (left)

    1: left
    2: 180
    3: right
    */
    let commands = n1 + n2
    commands *= commands < 0 ? -1 : 1
    switch (c) {
        case 'N':
            switch (commands) {
                case 1: c = 'W'; break
                case 2: c = 'S'; break
                case 3: c = 'E'; break
            }
            break
        case 'E':
            switch (commands) {
                case 1: c = 'N'; break
                case 2: c = 'W'; break
                case 3: c = 'S'; break
            }
            break
        case 'W':
            switch (commands) {
                case 1: c = 'S'; break
                case 2: c = 'E'; break
                case 3: c = 'N'; break
            }
            break
        case 'S':
            switch (commands) {
                case 1: c = 'E'; break
                case 2: c = 'N'; break
                case 3: c = 'W'; break
            }
            break
    }
    return compareChar(c)
}

// { "no": 12, "dat": "2", "ans": "2" }
const case12 = () => {
    const PI = 3.14
    let index = nextInt()
    let value = nextFloat()
    let r, d, l, s
    switch (index) {
        case 1:
            r = value
            d = 2 * r
            l = 2 * PI * r
            s = PI * r * r
            break
        case 2:
            d = value
            r = d / 2
            l = 2 * PI * r
            s = PI * r * r
            break
        case 3:
            l = value
            r = l / (2 * PI)
            d = 2 * r
            s = PI * r * r
            break
        case 4:
            s = value
            r = Math.sqrt(s / PI)
            d = 2 * r
            l = 2 * PI * r
            break
    }
    if (index != 1 && !compareFloat(2, r)) {
        return false
    }
    if (index != 2 && !compareFloat(2, d)) {
        return false
    }
    if (index != 3 && !compareFloat(2, l)) {
        return false
    }
    if (index != 4 && !compareFloat(2, s)) {
        return false
    }
    return true
}

// { "no": 13, "dat": "2", "ans": "2" }
const case13 = () => {
    let index = nextInt()
    let value = nextFloat()
    let a, c, h, s
    switch (index) {
        case 1:
            a = value
            c = a * Math.sqrt(2.0)
            h = c / 2
            s = c * h / 2
            break
        case 2:
            c = value
            a = c / Math.sqrt(2.0)
            h = c / 2
            s = c * h / 2
            break
        case 3:
            h = value
            c = 2 * h
            a = c / Math.sqrt(2.0)
            s = c * h / 2
            break
        case 4:
            s = value
            h = Math.sqrt(s)
            c = 2 * h
            a = c / Math.sqrt(2.0)
            break
    }
    if (index != 1 && !compareFloat(2, a)) {
        return false
    }
    if (index != 2 && !compareFloat(2, c)) {
        return false
    }
    if (index != 3 && !compareFloat(2, h)) {
        return false
    }
    if (index != 4 && !compareFloat(2, s)) {
        return false
    }
    return true
}

// { "no": 14, "dat": "2", "ans": "2" }
const case14 = () => {
    let index = nextInt()
    let value = nextFloat()
    let a, r1, r2, s
    switch (index) {
        case 1:
            a = value
            r1 = a * Math.sqrt(3.0) / 6
            r2 = 2 * r1
            s = a * a * Math.sqrt(3.0) / 4
            break
        case 2:
            r1 = value
            r2 = 2 * r1
            a = 6 * r1 / Math.sqrt(3.0)
            s = a * a * Math.sqrt(3.0) / 4
            break
        case 3:
            r2 = value
            r1 = r2 / 2
            a = 6 * r1 / Math.sqrt(3.0)
            s = a * a * Math.sqrt(3.0) / 4
            break
        case 4:
            s = value
            a = Math.sqrt(4 * s / Math.sqrt(3.0))
            r1 = a * Math.sqrt(3.0) / 6
            r2 = 2 * r1
            break
    }
    if (index != 1 && !compareFloat(2, a)) {
        return false
    }
    if (index != 2 && !compareFloat(2, r1)) {
        return false
    }
    if (index != 3 && !compareFloat(2, r2)) {
        return false
    }
    if (index != 4 && !compareFloat(2, s)) {
        return false
    }
    return true
}

// { "no": 15, "dat": "", "ans": "" }
const case15 = () => {
    let cardValue = nextInt()
    let cardSuit = nextInt()
    let description = ""
    switch (cardValue) {
        case 6: description += "six"; break
        case 7: description += "seven"; break
        case 8: description += "eight"; break
        case 9: description += "nine"; break
        case 10: description += "ten"; break
        case 11: description += "jack"; break
        case 12: description += "queen"; break
        case 13: description += "king"; break
        case 14: description += "ace"; break
    }
    description += " of "
    switch (cardSuit) {
        case 1: description += "spades"; break
        case 2: description += "clubs"; break
        case 3: description += "diamonds"; break
        case 4: description += "hearts"; break
    }
    return compareLine(description)
}

// { "no": 16, "dat": "", "ans": "" }
const case16 = () => {
    let age = nextInt()
    let alphabetic = ""
    switch (parseInt(age / 10)) {
        case 2: alphabetic += "twenty"; break
        case 3: alphabetic += "thirty"; break
        case 4: alphabetic += "fourty"; break
        case 5: alphabetic += "fifty"; break
        case 6: alphabetic += "sixty"; break
    }
    switch (age % 10) {
        case 1: alphabetic += "-one"; break
        case 2: alphabetic += "-two"; break
        case 3: alphabetic += "-three"; break
        case 4: alphabetic += "-four"; break
        case 5: alphabetic += "-five"; break
        case 6: alphabetic += "-six"; break
        case 7: alphabetic += "-seven"; break
        case 8: alphabetic += "-eight"; break
        case 9: alphabetic += "-nine"; break
    }
    alphabetic += " years"
    return compareLine(alphabetic)
}

// { "no": 17, "dat": "", "ans": "" }
const case17 = () => {
    let taskIndex = nextInt()
    let tens = parseInt(taskIndex / 10)
    let ones = taskIndex % 10
    let alphabetic = "the "
    switch (tens) {
        case 1:
            switch (ones) {
                case 0: alphabetic += "tenth"; break
                case 1: alphabetic += "eleventh"; break
                case 2: alphabetic += "twelfth"; break
                case 3: alphabetic += "thirteenth"; break
                case 4: alphabetic += "fourteenth"; break
                case 5: alphabetic += "fifteenth"; break
                case 6: alphabetic += "sixteenth"; break
                case 7: alphabetic += "seventeenth"; break
                case 8: alphabetic += "eighteenth"; break
                case 9: alphabetic += "nineteenth"; break
            }
            break
        case 2: alphabetic += "twent"; break
        case 3: alphabetic += "thirt"; break
        case 4: alphabetic += "fort"; break
    }
    if (tens != 1) {
        switch (ones) {
            case 1: alphabetic += "y-first"; break
            case 2: alphabetic += "y-second"; break
            case 3: alphabetic += "y-third"; break
            case 4: alphabetic += "y-fourth"; break
            case 5: alphabetic += "y-fifth"; break
            case 6: alphabetic += "y-sixth"; break
            case 7: alphabetic += "y-seventh"; break
            case 8: alphabetic += "y-eighth"; break
            case 9: alphabetic += "y-ninth"; break
            default: alphabetic += "ieth"; break
        }
    }
    alphabetic += " task"
    return compareLine(alphabetic)
}

// { "no": 18, "dat": "", "ans": "" }
const case18 = () => {
    let number = nextInt()
    let hundreds = parseInt(number / 100)
    let tens = parseInt(number / 10) % 10
    let ones = number % 10
    let alphabetic = ""
    switch (hundreds) {
        case 1: alphabetic += "one"; break
        case 2: alphabetic += "two"; break
        case 3: alphabetic += "three"; break
        case 4: alphabetic += "four"; break
        case 5: alphabetic += "five"; break
        case 6: alphabetic += "six"; break
        case 7: alphabetic += "seven"; break
        case 8: alphabetic += "eight"; break
        case 9: alphabetic += "nine"; break
    }
    if (hundreds != 0) {
        alphabetic += " hundred"
    }
    if (tens != 0 || ones != 0) {
        alphabetic += " and "
    }
    switch (tens) {
        case 1:
            switch (ones) {
                case 0: alphabetic += "ten"; break
                case 1: alphabetic += "eleven"; break
                case 2: alphabetic += "twelve"; break
                case 3: alphabetic += "thirteen"; break
                case 4: alphabetic += "fourteen"; break
                case 5: alphabetic += "fifteen"; break
                case 6: alphabetic += "sixteen"; break
                case 7: alphabetic += "seventeen"; break
                case 8: alphabetic += "eighteen"; break
                case 9: alphabetic += "nineteen"; break
            }
            break
        case 2: alphabetic += "twenty"; break
        case 3: alphabetic += "thirty"; break
        case 4: alphabetic += "forty"; break
        case 5: alphabetic += "fifty"; break
        case 6: alphabetic += "sixty"; break
        case 7: alphabetic += "seventy"; break
        case 8: alphabetic += "eighty"; break
        case 9: alphabetic += "ninety"; break
    }
    if (tens > 1 && ones != 0) {
        alphabetic += "-"
    }
    if (tens != 1) {
        switch (ones) {
            case 1: alphabetic += "one"; break
            case 2: alphabetic += "two"; break
            case 3: alphabetic += "three"; break
            case 4: alphabetic += "four"; break
            case 5: alphabetic += "five"; break
            case 6: alphabetic += "six"; break
            case 7: alphabetic += "seven"; break
            case 8: alphabetic += "eight"; break
            case 9: alphabetic += "nine"; break
        }
    }
    return compareLine(alphabetic)
}

// { "no": 19, "dat": "", "ans": "" }
const case19 = () => {
    let year = nextInt()
    year -= 4
    if (year < 0) {
        year = -year
    }
    year %= 60
    let yearName = "The "
    // colors
    if (year < 12) {
        yearName += "Green "
    } else if (year < 24) {
        yearName += "Red "
    } else if (year < 36) {
        yearName += "Yellow "
    } else if (year < 48) {
        yearName += "White "
    } else {
        yearName += "Black "
    }
    // animals
    switch (year % 60 % 12) {
        case 0: yearName += "Rat"; break
        case 1: yearName += "Cow"; break
        case 2: yearName += "Tiger"; break
        case 3: yearName += "Hare"; break
        case 4: yearName += "Dragon"; break
        case 5: yearName += "Snake"; break
        case 6: yearName += "Horse"; break
        case 7: yearName += "Sheep"; break
        case 8: yearName += "Monkey"; break
        case 9: yearName += "Hen"; break
        case 10: yearName += "Dog"; break
        case 11: yearName += "Pig"; break
    }
    yearName += "'s year"
    return compareLine(yearName)
}

// { "no": 20, "dat": "", "ans": "" }
const case20 = () => {
    let d = nextInt()
    let m = nextInt()
    let zodiacalName = ""
    switch (m) {
        case 1:
            if (d < 20) {
                zodiacalName = "Capricorn"
            } else {
                zodiacalName = "Aquarius"
            }
            break
        case 2:
            if (d < 19) {
                zodiacalName = "Aquarius"
            } else {
                zodiacalName = "Pisces"
            }
            break
        case 3:
            if (d < 21) {
                zodiacalName = "Pisces"
            } else {
                zodiacalName = "Aries"
            }
            break
        case 4:
            if (d < 20) {
                zodiacalName = "Aries"
            } else {
                zodiacalName = "Taurus"
            }
            break
        case 5:
            if (d < 21) {
                zodiacalName = "Taurus"
            } else {
                zodiacalName = "Gemini"
            }
            break
        case 6:
            if (d < 22) {
                zodiacalName = "Gemini"
            } else {
                zodiacalName = "Cancer"
            }
            break
        case 7:
            if (d < 23) {
                zodiacalName = "Cancer"
            } else {
                zodiacalName = "Leo"
            }
            break
        case 8:
            if (d < 23) {
                zodiacalName = "Leo"
            } else {
                zodiacalName = "Virgo"
            }
            break
        case 9:
            if (d < 23) {
                zodiacalName = "Virgo"
            } else {
                zodiacalName = "Libra"
            }
            break
        case 10:
            if (d < 23) {
                zodiacalName = "Libra"
            } else {
                zodiacalName = "Scorpio"
            }
            break
        case 11:
            if (d < 23) {
                zodiacalName = "Scorpio"
            } else {
                zodiacalName = "Sagittarius"
            }
            break
        case 12:
            if (d < 22) {
                zodiacalName = "Sagittarius"
            } else {
                zodiacalName = "Capricorn"
            }
            break
    }
    return compareWord(zodiacalName)
}