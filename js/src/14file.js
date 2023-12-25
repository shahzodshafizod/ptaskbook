import {
	initScanner
	, initChecker
	, inputsEOF
	, outputsEOF
} from './io.js'

(async () => {
	for (let taskNo = 1; taskNo <= 90; taskNo++) {
		const taskNoStr = String(taskNo).padStart(3, 0)
		for (let testNo = 1; testNo <= 100; testNo++) {
			const testNoStr = String(testNo).padStart(3, 0)
			const filePath = `../data/kt-gov2/14file/File${taskNoStr}/${testNoStr}/${testNoStr}`

			await initScanner(filePath + '.dat')
			await initChecker(filePath + '.ans')

			let ok = true
			switch (taskNo) {
				case 1: ok = file1(); break
				case 2: ok = file2(); break
				case 3: ok = file3(); break
				case 4: ok = file4(); break
				case 5: ok = file5(); break
				case 6: ok = file6(); break
				case 7: ok = file7(); break
				case 8: ok = file8(); break
				case 9: ok = file9(); break
				case 10: ok = file10(); break
				case 11: ok = file11(); break
				case 12: ok = file12(); break
				case 13: ok = file13(); break
				case 14: ok = file14(); break
				case 15: ok = file15(); break
				case 16: ok = file16(); break
				case 17: ok = file17(); break
				case 18: ok = file18(); break
				case 19: ok = file19(); break
				case 20: ok = file20(); break
				case 21: ok = file21(); break
				case 22: ok = file22(); break
				case 23: ok = file23(); break
				case 24: ok = file24(); break
				case 25: ok = file25(); break
				case 26: ok = file26(); break
				case 27: ok = file27(); break
				case 28: ok = file28(); break
				case 29: ok = file29(); break
				case 30: ok = file30(); break
				case 31: ok = file31(); break
				case 32: ok = file32(); break
				case 33: ok = file33(); break
				case 34: ok = file34(); break
				case 35: ok = file35(); break
				case 36: ok = file36(); break
				case 37: ok = file37(); break
				case 38: ok = file38(); break
				case 39: ok = file39(); break
				case 40: ok = file40(); break
				case 41: ok = file41(); break
				case 42: ok = file42(); break
				case 43: ok = file43(); break
				case 44: ok = file44(); break
				case 45: ok = file45(); break
				case 46: ok = file46(); break
				case 47: ok = file47(); break
				case 48: ok = file48(); break
				case 49: ok = file49(); break
				case 50: ok = file50(); break
				case 51: ok = file51(); break
				case 52: ok = file52(); break
				case 53: ok = file53(); break
				case 54: ok = file54(); break
				case 55: ok = file55(); break
				case 56: ok = file56(); break
				case 57: ok = file57(); break
				case 58: ok = file58(); break
				case 59: ok = file59(); break
				case 60: ok = file60(); break
				case 61: ok = file61(); break
				case 62: ok = file62(); break
				case 63: ok = file63(); break
				case 64: ok = file64(); break
				case 65: ok = file65(); break
				case 66: ok = file66(); break
				case 67: ok = file67(); break
				case 68: ok = file68(); break
				case 69: ok = file69(); break
				case 70: ok = file70(); break
				case 71: ok = file71(); break
				case 72: ok = file72(); break
				case 73: ok = file73(); break
				case 74: ok = file74(); break
				case 75: ok = file75(); break
				case 76: ok = file76(); break
				case 77: ok = file77(); break
				case 78: ok = file78(); break
				case 79: ok = file79(); break
				case 80: ok = file80(); break
				case 81: ok = file81(); break
				case 82: ok = file82(); break
				case 83: ok = file83(); break
				case 84: ok = file84(); break
				case 85: ok = file85(); break
				case 86: ok = file86(); break
				case 87: ok = file87(); break
				case 88: ok = file88(); break
				case 89: ok = file89(); break
				case 90: ok = file90(); break
			}

			if (!ok) {
				console.log("File" + taskNoStr + " #" + testNoStr + " has NOT been tested!")
				return
			}

			if (!inputsEOF()) {
				console.log("File" + taskNoStr + " #" + testNoStr + " input NOT EOF!")
				return
			}

			if (!outputsEOF()) {
				console.log("File" + taskNoStr + " #" + testNoStr + " output NOT EOF!")
				return
			}
		}
		console.log("File" + taskNoStr + " has been tested!")
	}
	console.log("File has been tested!")
})()

// { "no": 1, "dat": "", "ans": "" }
const file1 = () => {
	return false
}

// { "no": 2, "dat": "", "ans": "" }
const file2 = () => {
	return false
}

// { "no": 3, "dat": "2", "ans": "2" }
const file3 = () => {
	return false
}

// { "no": 4, "dat": "", "ans": "" }
const file4 = () => {
	return false
}

// { "no": 5, "dat": "", "ans": "" }
const file5 = () => {
	return false
}

// { "no": 6, "dat": "", "ans": "" }
const file6 = () => {
	return false
}

// { "no": 7, "dat": "", "ans": "" }
const file7 = () => {
	return false
}

// { "no": 8, "dat": "2", "ans": "2" }
const file8 = () => {
	return false
}

// { "no": 9, "dat": "2", "ans": "2" }
const file9 = () => {
	return false
}

// { "no": 10, "dat": "", "ans": "" }
const file10 = () => {
	return false
}

// { "no": 11, "dat": "2", "ans": "2" }
const file11 = () => {
	return false
}

// { "no": 12, "dat": "", "ans": "" }
const file12 = () => {
	return false
}

// { "no": 13, "dat": "", "ans": "" }
const file13 = () => {
	return false
}

// { "no": 14, "dat": "2", "ans": "2" }
const file14 = () => {
	return false
}

// { "no": 15, "dat": "2", "ans": "2" }
const file15 = () => {
	return false
}

// { "no": 16, "dat": "", "ans": "" }
const file16 = () => {
	return false
}

// { "no": 17, "dat": "", "ans": "" }
const file17 = () => {
	return false
}

// { "no": 18, "dat": "2", "ans": "2" }
const file18 = () => {
	return false
}

// { "no": 19, "dat": "2", "ans": "2" }
const file19 = () => {
	return false
}

// { "no": 20, "dat": "2", "ans": "" }
const file20 = () => {
	return false
}

// { "no": 21, "dat": "2", "ans": "" }
const file21 = () => {
	return false
}

// { "no": 22, "dat": "2", "ans": "" }
const file22 = () => {
	return false
}

// { "no": 23, "dat": "2", "ans": "" }
const file23 = () => {
	return false
}

// { "no": 24, "dat": "2", "ans": "" }
const file24 = () => {
	return false
}

// { "no": 25, "dat": "2", "ans": "2" }
const file25 = () => {
	return false
}

// { "no": 26, "dat": "2", "ans": "2" }
const file26 = () => {
	return false
}

// { "no": 27, "dat": "", "ans": "" }
const file27 = () => {
	return false
}

// { "no": 28, "dat": "2", "ans": "2" }
const file28 = () => {
	return false
}

// { "no": 29, "dat": "", "ans": "" }
const file29 = () => {
	return false
}

// { "no": 30, "dat": "", "ans": "" }
const file30 = () => {
	return false
}

// { "no": 31, "dat": "", "ans": "" }
const file31 = () => {
	return false
}

// { "no": 32, "dat": "", "ans": "" }
const file32 = () => {
	return false
}

// { "no": 33, "dat": "", "ans": "" }
const file33 = () => {
	return false
}

// { "no": 34, "dat": "", "ans": "" }
const file34 = () => {
	return false
}

// { "no": 35, "dat": "", "ans": "" }
const file35 = () => {
	return false
}

// { "no": 36, "dat": "", "ans": "" }
const file36 = () => {
	return false
}

// { "no": 37, "dat": "", "ans": "" }
const file37 = () => {
	return false
}

// { "no": 38, "dat": "", "ans": "" }
const file38 = () => {
	return false
}

// { "no": 39, "dat": "", "ans": "" }
const file39 = () => {
	return false
}

// { "no": 40, "dat": "", "ans": "" }
const file40 = () => {
	return false
}

// { "no": 41, "dat": "", "ans": "" }
const file41 = () => {
	return false
}

// { "no": 42, "dat": "2", "ans": "2" }
const file42 = () => {
	return false
}

// { "no": 43, "dat": "2", "ans": "2" }
const file43 = () => {
	return false
}

// { "no": 44, "dat": "2", "ans": "2" }
const file44 = () => {
	return false
}

// { "no": 45, "dat": "2", "ans": "2" }
const file45 = () => {
	return false
}

// { "no": 46, "dat": "2", "ans": "2" }
const file46 = () => {
	return false
}

// { "no": 47, "dat": "2", "ans": "2" }
const file47 = () => {
	return false
}

// { "no": 48, "dat": "", "ans": "" }
const file48 = () => {
	return false
}

// { "no": 49, "dat": "", "ans": "" }
const file49 = () => {
	return false
}

// { "no": 50, "dat": "2", "ans": "2" }
const file50 = () => {
	return false
}

// { "no": 51, "dat": "2", "ans": "2" }
const file51 = () => {
	return false
}

// { "no": 52, "dat": "", "ans": "" }
const file52 = () => {
	return false
}

// { "no": 53, "dat": "", "ans": "" }
const file53 = () => {
	return false
}

// { "no": 54, "dat": "", "ans": "2" }
const file54 = () => {
	return false
}

// { "no": 55, "dat": "", "ans": "" }
const file55 = () => {
	return false
}

// { "no": 56, "dat": "", "ans": "" }
const file56 = () => {
	return false
}

// { "no": 57, "dat": "", "ans": "" }
const file57 = () => {
	return false
}

// { "no": 58, "dat": "", "ans": "" }
const file58 = () => {
	return false
}

// { "no": 59, "dat": "", "ans": "" }
const file59 = () => {
	return false
}

// { "no": 60, "dat": "", "ans": "" }
const file60 = () => {
	return false
}

// { "no": 61, "dat": "", "ans": "" }
const file61 = () => {
	return false
}

// { "no": 62, "dat": "", "ans": "" }
const file62 = () => {
	return false
}

// { "no": 63, "dat": "", "ans": "" }
const file63 = () => {
	return false
}

// { "no": 64, "dat": "", "ans": "" }
const file64 = () => {
	return false
}

// { "no": 65, "dat": "", "ans": "" }
const file65 = () => {
	return false
}

// { "no": 66, "dat": "", "ans": "" }
const file66 = () => {
	return false
}

// { "no": 67, "dat": "", "ans": "" }
const file67 = () => {
	return false
}

// { "no": 68, "dat": "", "ans": "" }
const file68 = () => {
	return false
}

// { "no": 69, "dat": "", "ans": "" }
const file69 = () => {
	return false
}

// { "no": 70, "dat": "", "ans": "" }
const file70 = () => {
	return false
}

// { "no": 71, "dat": "", "ans": "" }
const file71 = () => {
	return false
}

// { "no": 72, "dat": "", "ans": "" }
const file72 = () => {
	return false
}

// { "no": 73, "dat": "", "ans": "" }
const file73 = () => {
	return false
}

// { "no": 74, "dat": "2", "ans": "2" }
const file74 = () => {
	return false
}

// { "no": 75, "dat": "2", "ans": "2" }
const file75 = () => {
	return false
}

// { "no": 76, "dat": "2", "ans": "2" }
const file76 = () => {
	return false
}

// { "no": 77, "dat": "2", "ans": "2" }
const file77 = () => {
	return false
}

// { "no": 78, "dat": "2", "ans": "2" }
const file78 = () => {
	return false
}

// { "no": 79, "dat": "2", "ans": "2" }
const file79 = () => {
	return false
}

// { "no": 80, "dat": "2", "ans": "2" }
const file80 = () => {
	return false
}

// { "no": 81, "dat": "2", "ans": "2" }
const file81 = () => {
	return false
}

// { "no": 82, "dat": "2", "ans": "2" }
const file82 = () => {
	return false
}

// { "no": 83, "dat": "2", "ans": "2" }
const file83 = () => {
	return false
}

// { "no": 84, "dat": "2", "ans": "2" }
const file84 = () => {
	return false
}

// { "no": 85, "dat": "2", "ans": "2" }
const file85 = () => {
	return false
}

// { "no": 86, "dat": "2", "ans": "2" }
const file86 = () => {
	return false
}

// { "no": 87, "dat": "2", "ans": "2" }
const file87 = () => {
	return false
}

// { "no": 88, "dat": "2", "ans": "2" }
const file88 = () => {
	return false
}

// { "no": 89, "dat": "2", "ans": "2" }
const file89 = () => {
	return false
}

// { "no": 90, "dat": "2", "ans": "2" }
const file90 = () => {
	return false
}