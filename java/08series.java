import java.util.Scanner;

class SeriesTask {

	private boolean series1() {
		Scanner scanner = new Scanner(System.in);
		double number;
		double sum = 0;
		for (int i = 1; i <= 10; i++) {
			number = scanner.nextDouble();
			sum += number;
		}
		System.out.printf("sum = %1$.2f\n", sum);
		return false;
	}

	private boolean series2() {
		Scanner scanner = new Scanner(System.in);
		double number;
		double mult = 1;
		for (int i = 1; i <= 10; i++) {
			number = scanner.nextDouble();
			mult *= number;
		}
		System.out.printf("mult = %1$.2f\n", mult);
		return false;
	}

	private boolean series3() {
		Scanner scanner = new Scanner(System.in);
		double number;
		double sum = 0;
		for (int i = 1; i <= 10; i++) {
			number = scanner.nextDouble();
			sum += number;
		}
		double aMean = sum / 10;
		System.out.printf("ariphmetical mean:\t%1$.2f\n", aMean);
		return false;
	}

	private boolean series4() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("N = ");
		int N = scanner.nextInt();
		double number;
		double sum = 0;
		double mult = 1;
		for (int i = 1; i <= N; i++) {
			number = scanner.nextDouble();
			sum += number;
			mult *= number;
		}
		System.out.printf("sum = %1$.2f\t\tmult = %2$.2f\n", sum, mult);
		return false;
	}

	private boolean series5() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("N = ");
		int N = scanner.nextInt();
		double number;
		double intPart;
		double sum = 0;
		for (int i = 1; i <= N; i++) {
			number = scanner.nextDouble();
			intPart = (int)number;
			System.out.printf("%1$.2f\t", intPart);
			sum += intPart;
		}
		System.out.printf("\nsum = %1$.2f\n", sum);
		return false;
	}

	private boolean series6() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("N = ");
		int N = scanner.nextInt();
		double number;
		double precPart;
		double mult = 1;
		for (int i = 0; i < N; i++) {
			number = scanner.nextDouble();
			precPart = number - (int)number;
			mult *= precPart;
			System.out.printf("%1$.2f\t", precPart);
		}
		System.out.printf("\nmult = %1$.6f\n", mult);
		return false;
	}

	private boolean series7() {
        int N = getInt();
        double number;
        int intPart, round, sum = 0;
        for (int i = 1; i <= N; i++) {
            number = getDouble();
            intPart = (int)number;
            round = (int)Math.round(number);
            put(round);
            sum += round;
        }
        put(sum);
        return false;
    }

	private boolean series8() {
        int N = getInt();
        int number, K = 0;
        for (int i = 1; i <= N; i++) {
            number = getInt();
            if (number % 2 == 0) {
                put(number);
                K++;
            }
        }
        put(K);
        return false;
    }

	private boolean series9() {
        int N = getInt();
        int number, K = 0;
        for (int i = 1; i <= N; i++) {
            number = getInt();
            if (number % 2 != 0) {
                put(i);
                K++;
            }
        }
        put(K);
        return false;
    }

	private boolean series10() {
        int N = getInt();
        int number;
        boolean hasPositive = false;
        for (int i = 1; i <= N; i++) {
            number = getInt();
            if ( !hasPositive && (number > 0) ) {
                hasPositive = true;
            }
        }
        put(hasPositive);
        return false;
    }

	private boolean series11() {
        int K = getInt();
        int N = getInt();
        int number;
        boolean hasLessThanK = false;
        for (int i = 1; i <= N; i++) {
            number = getInt();
            if ( !hasLessThanK && (number < K) ) {
                hasLessThanK = true;
            }
        }
        put(hasLessThanK);
        return false;
    }

	private boolean series12() {
        int number, count = 0;
        while (true) {
            number = getInt();
            if (number == 0) {
                break;
            }
            count++;
        }
        put(count);
        return false;
    }

	private boolean series13() {
        int number, sum = 0;
        while (true) {
            number = getInt();
            if (number == 0) {
                break;
            }
            if ( (number > 0) && (number % 2 == 0) ) {
                sum += number;
            }
        }
        put(sum);
        return false;
    }

	private boolean series14() {
        int K = getInt();
        int number, count = 0;
        while (true) {
            number = getInt();
            if (number == 0) {
                break;
            }
            if (number < K) {
                count++;
            }
        }
        put(count);
        return false;
    }

	private boolean series15() {
        int K = getInt();
        int number, index = 0;
        for (int i = 1; true; i++) {
            number = getInt();
            if (number == 0) {
                break;
            }
            if ( (index == 0) && (number > K) ) {
                index = i;
            }
        }
        put(index);
        return false;
    }

	private boolean series16() {
        int K = getInt();
        int number;
        int index = 0;
        for (int i = 1; true; i++) {
            number = getInt();
            if (number == 0) {
                break;
            }
            if (number > K) {
                index = i;
            }
        }
        put(index);
        return false;
    }

	private boolean series17() {
        double B = getDouble();
        int N = getInt();
        double number;
        boolean outed = false;
        for (int i = 1; i <= N; i++) {
            number = getDouble();
            if ( !outed && (B < number) ) {
                put(B);
                outed = true;
            }
            put(number);
        }
        if (!outed) {
            put(B);
        }
        return false;
    }

	private boolean series18() {
        int N = getInt();
        int number1, number2;
        number1 = getInt();
        put(number1);
        for (int i = 2; i <= N; i++) {
            number2 = getInt();
            if (number1 != number2) {
                put(number2);
            }
            number1 = number2;
        }
        return false;
    }

	private boolean series19() {
        int N = getInt();
        int number1, number2, K = 0;
        number1 = getInt();
        for (int i = 2; i <= N; i++) {
            number2 = getInt();
            if (number2 < number1) {
                put(number2);
                K++;
            }
            number1 = number2;
        }
        put(K);
        return false;
    }

	private boolean series20() {
        int N = getInt();
        int number1, number2, K = 0;
        number1 = getInt();
        for (int i = 2; i <= N; i++) {
            number2 = getInt();
            if (number1 < number2) {
                put(number1);
                K++;
            }
            number1 = number2;
        }
        put(K);
        return false;
    }

	private boolean series21() {
        int N = getInt();
        double number1, number2;
        number1 = getDouble();
        boolean progress = true;
        for (int i = 2; i <= N; i++) {
            number2 = getDouble();
            if ( progress && (number1 > number2) ) {
                progress = false;
            }
            number1 = number2;
        }
        put(progress);
        return false;
    }

	private boolean series22() {
        int N = getInt();
        double number1, number2;
        int index = 0;
        number1 = getDouble();
        for (int i = 2; i <= N; i++) {
            number2 = getDouble();
            if ( (index == 0) && (number1 < number2) ) {
                index = i;
            }
            number1 = number2;
        }
        put(index);
        return false;
    }

	private boolean series23() {
        int N = getInt();
        double number1, number2, number3;
        number1 = getDouble();
        number2 = getDouble();
        boolean arrashakl;
        int index = 0;
        for (int i = 3; i <= N; i++) {
            number3 = getDouble();
            arrashakl = (number2 > number1) && (number2 > number3) ||
                        (number2 < number1) && (number2 < number3);
            if ( !arrashakl && (index == 0) ) {
                index = i-1;
            }
            number1 = number2;
            number2 = number3;
        }
        put(index);
        return false;
    }

	private boolean series24() {
        int N = getInt();
        int number, partSum = 0, sum = 0;
        for (int i = 1; i <= N; i++) {
            number = getInt();
            if (number == 0) {
                sum = partSum;
                partSum = 0;
            } else {
                partSum += number;
            }
        }
        put(sum);
        return false;
    }

	private boolean series25() {
        int N = getInt();
        int number, partSum = 0, sum = 0;
        boolean started = false;
        for (int i = 1; i <= N; i++) {
            number = getInt();
            if (number == 0) {
                started = true;
                sum += partSum;
                partSum = 0;
            } else if (started) {
                partSum += number;
            }
        }
        put(sum);
        return false;
    }

	private boolean series26() {
        int K = getInt(), N = getInt();
        double number, degree;
        for (int i = 1; i <= N; i++) {
            number = getDouble();
            degree = 1;
            for (int j = 1; j <= K; j++) {
                degree *= number;
            }
            put(degree);
        }
        return false;
    }

	private boolean series27() {
        int N = getInt();
        double number, degree;
        for (int i = 1; i <= N; i++) {
            number = getDouble();
            degree = 1;
            for (int j = 1; j <= i; j++) {
                degree *= number;
            }
            put(degree);
        }
        return false;
    }

	private boolean series28() {
        int N = getInt();
        double number, degree;
        for (int i = N; i > 0; i--) {
            number = getDouble();
            degree = 1;
            for (int j = 1; j <= i; j++) {
                degree *= number;
            }
            put(degree);
        }
        return false;
    }

	private boolean series29() {
        int K = getInt(), N = getInt();
        int number, sum = 0;
        for (int i = 1; i <= K; i++) {
            for (int j = 1; j <= N; j++) {
                number = getInt();
                sum += number;
            }
        }
        put(sum);
        return false;
    }

	private boolean series30() {
        int K = getInt(), N = getInt();
        int number, sum;
        for (int i = 1; i <= K; i++) {
            sum = 0;
            for (int j = 1; j <= N; j++) {
                number = getInt();
                sum += number;
            }
            put(sum);
        }
        return false;
    }

	private boolean series31() {
        int K = getInt(), N = getInt();
        int number, count = 0;
        boolean has2;
        for (int i = 1; i <= K; i++) {
            has2 = false;
            for (int j = 1; j <= N; j++) {
                number = getInt();
                if (number == 2) {
                    has2 = true;
                }
            }
            if (has2) {
                count++;
            }
        }
        put(count);
        return false;
    }

	private boolean series32() {
        int K = getInt(), N = getInt();
        int number, index;
        for (int i = 1; i <= K; i++) {
            index = 0;
            for (int j = 1; j <= N; j++) {
                number = getInt();
                if ( (index == 0) && (number == 2) ) {
                    index = j;
                }
            }
            put(index);
        }
        return false;
    }

	private boolean series33() {
        int K = getInt(), N = getInt();
        int number, index;
        for (int i = 1; i <= K; i++) {
            index = 0;
            for (int j = 1; j <= N; j++) {
                number = getInt();
                if (number == 2) {
                    index = j;
                }
            }
            put(index);
        }
        return false;
    }

	private boolean series34() {
        int K = getInt(), N = getInt();
        int number, count, sum;
        boolean has2;
        for (int i = 1; i <= K; i++) {
            has2 = false;
            sum = 0;
            for (int j = 1; j <= N; j++) {
                number = getInt();
                sum += number;
                if (number == 2) {
                    has2 = true;
                }
            }
            if (has2) {
                put(sum);
            } else {
                put(0);
            }
        }
        return false;
    }

	private boolean series35() {
        int K = getInt();
        int number, count, totalCount = 0;
        for (int i = 1; i <= K; i++) {
            count = 0;
            while (true) {
                number = getInt();
                if (number == 0) {
                    break;
                }
                count++;
            }
            put(count);
            totalCount += count;
        }
        put(totalCount);
        return false;
    }

	private boolean series36() {
        int K = getInt();
        int number1, number2, count = 0;
        boolean progress;
        for (int i = 1; i <= K; i++) {
            number1 = getInt();
            progress = true;
            while (true) {
                number2 = getInt();
                if (number2 == 0) {
                    break;
                }
                if (number1 > number2) {
                    progress = false;
                }
                number1 = number2;
            }
            if (progress) {
                count++;
            }
        }
        put(count);
        return false;
    }

	private boolean series37() {
        int K = getInt();
        int number1, number2, count = 0;
        boolean progress, regress;
        for (int i = 1; i <= K; i++) {
            number1 = getInt();
            progress = regress = true;
            while (true) {
                number2 = getInt();
                if (number2 == 0) {
                    break;
                }
                if (number1 < number2) {
                    regress = false;
                } else if (number1 > number2) {
                    progress = false;
                }
                number1 = number2;
            }
            if (progress || regress) {
                count++;
            }
        }
        put(count);
        return false;
    }

	private boolean series38() {
        int K = getInt();
        int number1, number2;
        boolean progress, regress;
        for (int i = 1; i <= K; i++) {
            number1 = getInt();
            progress = regress = true;
            while (true) {
                number2 = getInt();
                if (number2 == 0) {
                    break;
                }
                if (number1 < number2) {
                    regress = false;
                }
                else if (number1 > number2) {
                    progress = false;
                }
                number1 = number2;
            }
            if (progress) {
                put(1);
            } else if (regress) {
                put(-1);
            } else {
                put(0);
            }
        }
        return false;
    }

	private boolean series39() {
        int K = getInt();
        int number1, number2, number3, count = 0;
        boolean arrashakl;
        for (int i = 1; i <= K; i++) {
            number1 = getInt();
            number2 = getInt();
            arrashakl = true;
            while (true) {
                number3 = getInt();
                if (number3 == 0) {
                    break;
                }
                if ( (number1 > number2) && (number2 > number3) || 
                        (number1 < number2) && (number2 < number3) ) {
                    arrashakl = false;
                }
                number1 = number2;
                number2 = number3;
            }
            if (arrashakl) {
                count++;
            }
        }
        put(count);
        return false;
    }

	private boolean series40() {
        int K = getInt();
        int number1, number2, number3, count, errorNo;
        boolean arrashakl;
        for (int i = 1; i <= K; i++) {
            number1 = getInt();
            number2 = getInt();
            count = 2;
            errorNo = 0;
            while (true) {
                number3 = getInt();
                if (number3 == 0) {
                    break;
                }
                count++;
                arrashakl = (number2 > number1) && (number2 > number3) || 
                            (number2 < number1) && (number2 < number3);
                if ( !arrashakl && (errorNo == 0) ) {
                    errorNo = count-1;
                }
                number1 = number2;
                number2 = number3;
            }
            if (errorNo == 0) {
                put(count);
            } else {
                put(errorNo);
            }
        }
        return false;
    }
}