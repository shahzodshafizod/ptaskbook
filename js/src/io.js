import fs from 'node:fs'

let inputs, outputs

//FLOAT/////////////////////////////////////////////////////////////////////////
export function nextFloat() {
    return parseFloat(__getWord("inputs"));
}

export function roundFixed(number, n) {
    return Number(number.toLocaleString('en-US', {
        minimumFractionDigits: n,
        maximumFractionDigits: n,
    }).replaceAll(',', ''))
}

export function compareFloat(fractionDigits, ...values) {
    for (let i = 0; i < values.length; i++) {
        let value = roundFixed(parseFloat(values[i]), fractionDigits)
        let correctValue = roundFixed(parseFloat(__getWord("outputs")), fractionDigits)
        if (value !== correctValue) {
            console.log('values[i] = ' + values[i])
            console.log("value = " + value)
            console.log("correctValue = " + correctValue)
            return false
        }
    }
    return true
}
////////////////////////////////////////////////////////////////////////////////

//INTEGER///////////////////////////////////////////////////////////////////////
export function nextInt() {
    return parseInt(__getWord("inputs"))
}

export function compareInt(...values) {
    for (let i = 0; i < values.length; i++) {
        let value = parseInt(values[i])
        let correctValue = parseInt(__getWord("outputs"))
        if (value !== correctValue) {
            console.log("value = " + value)
            console.log("correctValue = " + correctValue)
            return false
        }
    }
    return true
}
////////////////////////////////////////////////////////////////////////////////

//BOOLEAN///////////////////////////////////////////////////////////////////////
export function nextBoolean() {
    return __parseBoolean(__getWord("inputs"))
}

export function compareBoolean(...values) {
    for (let i = 0; i < values.length; i++) {
        let value = __parseBoolean(values[i])
        let correctValue = __parseBoolean(__getWord("outputs"))
        if (value !== correctValue) {
            console.log("value = " + value)
            console.log("correctValue = " + correctValue)
            return false
        }
    }
    return true
}

const __parseBoolean = (value) => {
    let result
    switch (String(value).toLowerCase()) {
        case "true": result = true; break
        case "false": result = false; break
        default: {
            console.log("not a boolean value: " + value)
            result = false
            break
        }
    }
    return result
}
////////////////////////////////////////////////////////////////////////////////

//CHAR//////////////////////////////////////////////////////////////////////////
export function nextChar() {
    return __getChar("inputs")
}

export function compareChar(...values) {
    for (let i = 0; i < values.length; i++) {
        let value = values[i]
        let correctValue = __getChar("outputs")
        if (value !== correctValue) {
            console.log("value = " + value)
            console.log("correctValue = " + correctValue)
            return false
        }
    }
    return true
}
////////////////////////////////////////////////////////////////////////////////

//WORD//////////////////////////////////////////////////////////////////////////
export function nextWord() {
    return __getWord("inputs")
}

export function compareWord(...values) {
    for (let i = 0; i < values.length; i++) {
        let value = values[i]
        let correctValue = __getWord("outputs")
        if (value !== correctValue) {
            console.log("value = " + value)
            console.log("correctValue = " + correctValue)
            return false
        }
    }
    return true
}
////////////////////////////////////////////////////////////////////////////////

//LINE//////////////////////////////////////////////////////////////////////////
export function nextLine() {
    return __getLine("inputs")
}

export function compareLine(...values) {
    for (let i = 0; i < values.length; i++) {
        let value = values[i]
        let correctValue = __getLine("outputs")
        if (value !== correctValue) {
            console.log("value = " + value)
            console.log("correctValue = " + correctValue)
            return false
        }
    }
    return true
}
////////////////////////////////////////////////////////////////////////////////

const __getChar = (source) => {
    let char = null
    switch (source) {
        case "inputs":
            if (inputs.length > 0) {
                char = inputs.substring(0, 1)
                inputs = inputs.substring(2)
            }
            break
        case "outputs":
            if (outputs.length > 0) {
                char = outputs.substring(0, 1)
                outputs = outputs.substring(2)
            }
            break
    }
    if (char == null) {
        console.log("__getChar " + source + " EOF")
    }
    return char
}

const __getWord = (source) => {
    let word = null
    let spaceIndex
    switch (source) {
        case "inputs":
            if (inputs.trimLeft().length > 0) {
                inputs = inputs.trimLeft()
                spaceIndex = inputs.search(/\s/)
                if (spaceIndex < 0) {
                    spaceIndex = inputs.length
                }
                word = inputs.substring(0, spaceIndex)
                inputs = inputs.substring(spaceIndex+1)
            }
            break
        case "outputs":
            if (outputs.trimLeft().length > 0) {
                outputs = outputs.trimLeft()
                spaceIndex = outputs.search(/\s/)
                if (spaceIndex < 0) {
                    spaceIndex = outputs.length
                }
                word = outputs.substring(0, spaceIndex)
                outputs = outputs.substring(spaceIndex+1)
            }
            break
    }
    if (word == null) {
        console.log("__getWord " + source + " EOF")
    }
    return word
}

const __getLine = (source) => {
    let word = null
    let spaceIndex
    switch (source) {
        case "inputs":
            if (inputs.trimLeft().length > 0) {
                inputs = inputs.trimLeft()
                spaceIndex = inputs.search(/\r\n|\n/)
                if (spaceIndex < 0) {
                    spaceIndex = inputs.length
                }
                word = inputs.substring(0, spaceIndex)
                inputs = inputs.substring(spaceIndex+1)
            }
        case "outputs":
            if (outputs.trimLeft().length > 0) {
                outputs = outputs.trimLeft()
                spaceIndex = outputs.search(/\r\n|\n/)
                if (spaceIndex < 0) {
                    spaceIndex = outputs.length
                }
                word = outputs.substring(0, spaceIndex)
                outputs = outputs.substring(spaceIndex+1)
            }
    }
    if (word == null) {
        console.log("__getLine " + source + " EOF")
    }
    return word
}

const __readFile = (fileName) => {
    try {
        const data = fs.readFileSync(fileName, 'utf8')
        return data
    } catch (err) {
        console.error(err)
    }
}

export async function initScanner(fileName) {
    inputs = __readFile(fileName)
}

export async function initChecker(fileName) {
    outputs = __readFile(fileName)
}

export function inputsEOF() { return inputs.trimLeft().length == 0; }

export function outputsEOF() { return outputs.trimLeft().length == 0; }
