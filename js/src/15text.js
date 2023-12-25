import {
	initScanner
	, initChecker
	, inputsEOF
	, outputsEOF
} from './io.js'

(async () => {
	for (let taskNo = 1; taskNo <= 60; taskNo++) {
		const taskNoStr = String(taskNo).padStart(3, 0)
		for (let testNo = 1; testNo <= 100; testNo++) {
			const testNoStr = String(testNo).padStart(3, 0)
			const filePath = `../data/kt-gov2/15text/Text${taskNoStr}/${testNoStr}/${testNoStr}`

			await initScanner(filePath + '.dat')
			await initChecker(filePath + '.ans')

			let ok = true
			switch (taskNo) {
				case 1: ok = text1(); break
				case 2: ok = text2(); break
				case 3: ok = text3(); break
				case 4: ok = text4(); break
				case 5: ok = text5(); break
				case 6: ok = text6(); break
				case 7: ok = text7(); break
				case 8: ok = text8(); break
				case 9: ok = text9(); break
				case 10: ok = text10(); break
				case 11: ok = text11(); break
				case 12: ok = text12(); break
				case 13: ok = text13(); break
				case 14: ok = text14(); break
				case 15: ok = text15(); break
				case 16: ok = text16(); break
				case 17: ok = text17(); break
				case 18: ok = text18(); break
				case 19: ok = text19(); break
				case 20: ok = text20(); break
				case 21: ok = text21(); break
				case 22: ok = text22(); break
				case 23: ok = text23(); break
				case 24: ok = text24(); break
				case 25: ok = text25(); break
				case 26: ok = text26(); break
				case 27: ok = text27(); break
				case 28: ok = text28(); break
				case 29: ok = text29(); break
				case 30: ok = text30(); break
				case 31: ok = text31(); break
				case 32: ok = text32(); break
				case 33: ok = text33(); break
				case 34: ok = text34(); break
				case 35: ok = text35(); break
				case 36: ok = text36(); break
				case 37: ok = text37(); break
				case 38: ok = text38(); break
				case 39: ok = text39(); break
				case 40: ok = text40(); break
				case 41: ok = text41(); break
				case 42: ok = text42(); break
				case 43: ok = text43(); break
				case 44: ok = text44(); break
				case 45: ok = text45(); break
				case 46: ok = text46(); break
				case 47: ok = text47(); break
				case 48: ok = text48(); break
				case 49: ok = text49(); break
				case 50: ok = text50(); break
				case 51: ok = text51(); break
				case 52: ok = text52(); break
				case 53: ok = text53(); break
				case 54: ok = text54(); break
				case 55: ok = text55(); break
				case 56: ok = text56(); break
				case 57: ok = text57(); break
				case 58: ok = text58(); break
				case 59: ok = text59(); break
				case 60: ok = text60(); break
			}

			if (!ok) {
				console.log("Text" + taskNoStr + " #" + testNoStr + " has NOT been tested!")
				return
			}

			if (!inputsEOF()) {
				console.log("Text" + taskNoStr + " #" + testNoStr + " input NOT EOF!")
				return
			}

			if (!outputsEOF()) {
				console.log("Text" + taskNoStr + " #" + testNoStr + " output NOT EOF!")
				return
			}
		}
		console.log("Text" + taskNoStr + " has been tested!")
	}
	console.log("Text has been tested!")
})()

const text1 = () => {
	return false
}

const text2 = () => {
	return false
}

const text3 = () => {
	return false
}

const text4 = () => {
	return false
}

const text5 = () => {
	return false
}

const text6 = () => {
	return false
}

const text7 = () => {
	return false
}

const text8 = () => {
	return false
}

const text9 = () => {
	return false
}

const text10 = () => {
	return false
}

const text11 = () => {
	return false
}

const text12 = () => {
	return false
}

const text13 = () => {
	return false
}

const text14 = () => {
	return false
}

const text15 = () => {
	return false
}

const text16 = () => {
	return false
}

const text17 = () => {
	return false
}

const text18 = () => {
	return false
}

const text19 = () => {
	return false
}

const text20 = () => {
	return false
}

const text21 = () => {
	return false
}

const text22 = () => {
	return false
}

const text23 = () => {
	return false
}

const text24 = () => {
	return false
}

const text25 = () => {
	return false
}

const text26 = () => {
	return false
}

const text27 = () => {
	return false
}

const text28 = () => {
	return false
}

const text29 = () => {
	return false
}

const text30 = () => {
	return false
}

const text31 = () => {
	return false
}

const text32 = () => {
	return false
}

const text33 = () => {
	return false
}

const text34 = () => {
	return false
}

const text35 = () => {
	return false
}

const text36 = () => {
	return false
}

const text37 = () => {
	return false
}

const text38 = () => {
	return false
}

const text39 = () => {
	return false
}

const text40 = () => {
	return false
}

const text41 = () => {
	return false
}

const text42 = () => {
	return false
}

const text43 = () => {
	return false
}

const text44 = () => {
	return false
}

const text45 = () => {
	return false
}

const text46 = () => {
	return false
}

const text47 = () => {
	return false
}

const text48 = () => {
	return false
}

const text49 = () => {
	return false
}

const text50 = () => {
	return false
}

const text51 = () => {
	return false
}

const text52 = () => {
	return false
}

const text53 = () => {
	return false
}

const text54 = () => {
	return false
}

const text55 = () => {
	return false
}

const text56 = () => {
	return false
}

const text57 = () => {
	return false
}

const text58 = () => {
	return false
}

const text59 = () => {
	return false
}

const text60 = () => {
	return false
}