import java.util.Scanner;

class MinmaxTask {

	private boolean minmax1() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("N = ");
		int N = scanner.nextInt();
		double number = scanner.nextDouble();
		double minimal = number;
		double maximal = number;
		for (int i = 2; i <= N; i++) {
			number = scanner.nextDouble();
			if (number < minimal) {
				minimal = number;
			}
			if (number > maximal) {
				maximal = number;
			}
		}
		System.out.printf("minimal = %1$.2f\n", minimal);
		System.out.printf("maximal = %1$.2f\n", maximal);
		return false;
	}

	private boolean minmax2() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("N = ");
		int N = scanner.nextInt();
		double a;
		double b;
		double square;
		double minSquare = 0;
		for (int i = 1; i <= N; i++) {
			a = scanner.nextDouble();
			b = scanner.nextDouble();
			square = a * b;
			if (minSquare == 0) {
				minSquare = square;
			} else if (square < minSquare) {
				minSquare = square;
			}
		}
		System.out.printf("minimal square:\t%1$.2f\n", minSquare);
		return false;
	}

	private boolean minmax3() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("N = ");
		int N = scanner.nextInt();
		double a = scanner.nextDouble();
		double b = scanner.nextDouble();
		double P = 2 * (a + b);
		double maxP = P;
		for (int i = 2; i <= N; i++) {
			a = scanner.nextDouble();
			b = scanner.nextDouble();
			P = 2 * (a + b);
			if (P > maxP) {
				maxP = P;
			}
		}
		System.out.printf("maximal perimeter:\t%1$.2f\n", maxP);
		return false;
	}

	private boolean minmax4() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("N = ");
		int N = scanner.nextInt();
		double number = scanner.nextDouble();
		double minimal = number;
		int minIndex = 1;
		for (int i = 2; i <= N; i++) {
			number = scanner.nextDouble();
			if (number < minimal) {
				minimal = number;
				minIndex = i;
			}
		}
		System.out.println("minimal index:\t" + minIndex);
		return false;
	}

	private boolean minmax5() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("N = ");
		int N = scanner.nextInt();
		double m = scanner.nextDouble();
		double v = scanner.nextDouble();
		double P = m / v;
		double maxP = P;
		int maxPno = 1;
		for (int i = 2; i <= N; i++) {
			m = scanner.nextDouble();
			v = scanner.nextDouble();
			P = m / v;
			if (P > maxP) {
				maxP = P;
				maxPno = i;
			}
		}
		System.out.printf("maxPno = %d\tmaxP = %2$.2f\n", maxPno, maxP);
		return false;
	}

	private boolean minmax6() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("N = ");
		int N = scanner.nextInt();
		int number = scanner.nextInt();
		int minimal = number;
		int maximal = number;
		int firstMinIndex = 0;
		int lastMaxIndex = 0;
		for (int i = 1; i < N; i++) {
			number = scanner.nextInt();
			if (number < minimal) {
				minimal = number;
				firstMinIndex = i;
			}
			if (number >= maximal) {
				maximal = number;
				lastMaxIndex = i;
			}
		}
		firstMinIndex++;
		lastMaxIndex++;
		System.out.println("first minimal index:\t" + firstMinIndex);
		System.out.println("last maximal index:\t" + lastMaxIndex);
		return false;
	}

	private boolean minmax7() {
        int N = getInt();
        int number, maximal, minimal, minIndex, maxIndex;
        number = getInt();
        minimal = maximal = number;
        minIndex = maxIndex = 1;
        for (int i = 2; i <= N; i++) {
            number = getInt();
            if (number > maximal) {
                maximal = number;
                maxIndex = i;
            }
            if (number <= minimal) {
                if (number < minimal) {
                    minimal = number;
                }
                minIndex = i;
            }
        }
        put(maxIndex);
        put(minIndex);
        return false;
    }

	private boolean minmax8() {
        int N = getInt();
        int number, minimal, minIndex1, minIndexN;
        number = getInt();
        minimal = number;
        minIndex1 = minIndexN = 1;
        for (int i = 2; i <= N; i++) {
            number = getInt();
            if (number <= minimal) {
                if (number < minimal) {
                    minimal = number;
                    minIndex1 = minIndexN = i;
                } else {
                    minIndexN = i;
                }
            }
        }
        put(minIndex1);
        put(minIndexN);
        return false;
    }

	private boolean minmax9() {
        int N = getInt();
        int number, maximal, maxIndex1, maxIndexN;
        number = getInt();
        maximal = number;
        maxIndex1 = maxIndexN = 1;
        for (int i = 2; i <= N; i++) {
            number = getInt();
            if (number >= maximal) {
                if (number > maximal) {
                    maximal = number;
                    maxIndex1 = maxIndexN = i;
                } else {
                    maxIndexN = i;
                }
            }
        }
        put(maxIndex1);
        put(maxIndexN);
        return false;
    }

	private boolean minmax10() {
        int N = getInt();
        int number, minimal, maximal, minIndex, maxIndex;
        number = getInt();
        minimal = maximal = number;
        minIndex = maxIndex = 1;
        for (int i = 2; i <= N; i++) {
            number = getInt();
            if (number > maximal) {
                maximal = number;
                maxIndex = i;
            }
            if (number < minimal) {
                minimal = number;
                minIndex = i;
            }
        }
        int index = minIndex < maxIndex ? minIndex : maxIndex;
        put(index);
        return false;
    }

	private boolean minmax11() {
        int N = getInt();
        int number, maximal, minimal, minIndex, maxIndex;
        number = getInt();
        maximal = minimal = number;
        minIndex = maxIndex = 1;
        for (int i = 2; i <= N; i++) {
            number = getInt();
            if (number <= minimal) {
                if (number < minimal) {
                    minimal = number;
                }
                minIndex = i;
            }
            if (number >= maximal) {
                if (number > maximal) {
                    maximal = number;
                }
                maxIndex = i;
            }
        }
        int index = minIndex > maxIndex ? minIndex : maxIndex;
        put(index);
        return false;
    }

	private boolean minmax12() {
        int N = getInt();
        double number, minimal = 0;
        boolean inited = false;
        for (int i = 1; i <= N; i++) {
            number = getDouble();
            if (number > 0) {
                if (!inited) {
                    minimal = number;
                    inited = true;
                } else if (number < minimal) {
                    minimal = number;
                }
            }
        }
        put(minimal);
        return false;
    }

	private boolean minmax13() {
        int N = getInt();
        int number, maximal = 0, index = 0;
        boolean inited = false;
        for (int i = 1; i <= N; i++) {
            number = getInt();
            if (number % 2 != 0) {
                if (!inited) {
                    maximal = number;
                    index = i;
                    inited = true;
                } else if (number > maximal) {
                    maximal = number;
                    index = i;
                }
            }
        }
        put(index);
        return false;
    }

	private boolean minmax14() {
        double B = getDouble();
        double number, minimal = 0;
        int index = 0;
        boolean inited = false;
        for (int i = 1; i <= 10; i++) {
            number = getDouble();
            if (number > B) {
                if (!inited) {
                    minimal = number;
                    index = i;
                    inited = true;
                } else if (number < minimal) {
                    minimal = number;
                    index = i;
                }
            }
        }
        put(minimal);
        put(index);
        return false;
    }

	private boolean minmax15() {
        double B, C, number, maximal = 0;
        int index = 0;
        B = getDouble();
        C = getDouble();
        boolean inited = false;
        for (int i = 1; i <= 10; i++) {
            number = getDouble();
            if ( (number > B) && (number < C) ) {
                if (!inited) {
                    maximal = number;
                    index = i;
                    inited = true;
                } else if (number > maximal) {
                    maximal = number;
                    index = i;
                }
            }
        }
        put(maximal);
        put(index);
        return false;
    }

	private boolean minmax16() {
        int N = getInt();
        int number, minimal, firstMinIndex;
        number = getInt();
        minimal = number;
        firstMinIndex = 1;
        for (int i = 2; i <= N; i++) {
            number = getInt();
            if (number < minimal) {
                minimal = number;
                firstMinIndex = i;
            }
        }
        int elementsBeforeFirstMin = firstMinIndex - 1;
        put(elementsBeforeFirstMin);
        return false;
    }

	private boolean minmax17() {
        int N = getInt();
        int number, maximal, lastMaxIndex;
        number = getInt();
        maximal = number;
        lastMaxIndex = 1;
        for (int i = 2; i <= N; i++) {
            number = getInt();
            if (number >= maximal) {
                if (number > maximal) {
                    maximal = number;
                }
                lastMaxIndex = i;
            }
        }
        int elementsAfterLastMax = N - lastMaxIndex;
        put(elementsAfterLastMax);
        return false;
    }

	private boolean minmax18() {
        int N = getInt();
        int number, maximal, index1, indexN;
        number = getInt();
        maximal = number;
        index1 = indexN = 1;
        for (int i = 2; i <= N; i++) {
            number = getInt();
            if (number > maximal) {
                maximal = number;
                index1 = indexN = i;
            } else if (number == maximal) {
                indexN = i;
            }
        }
        int elementsBetween = index1 == indexN ? 0 : indexN - index1 - 1;
        put(elementsBetween);
        return false;
    }

	private boolean minmax19() {
        int N = getInt();
        int number, minimal, mins;
        number = getInt();
        minimal = number;
        mins = 1;
        for (int i = 2; i <= N; i++) {
            number = getInt();
            if (number < minimal) {
                minimal = number;
                mins = 1;
            } else if (number == minimal) {
                mins++;
            }
        }
        put(mins);
        return false;
    }

	private boolean minmax20() {
        int N = getInt();
        int number, minimal, maximal, mins, maxs;
        number = getInt();
        minimal = maximal = number;
        mins = maxs = 1;
        for (int i = 2; i <= N; i++) {
            number = getInt();
            if (number == minimal) {
                mins++;
            } else if (number < minimal) {
                minimal = number;
                mins = 1;
            }
            if (number == maximal) {
                maxs++;
            } else if (number > maximal) {
                maximal = number;
                maxs = 1;
            }
        }
        int sum = maxs;
        sum += minimal != maximal ? mins : 0;
        put(sum);
        return false;
    }

	private boolean minmax21() {
        int N = getInt();
        double number, minimal, maximal, sum = 0;
        number = getDouble();
        sum += number;
        minimal = maximal = number;
        for (int i = 2; i <= N; i++) {
            number = getDouble();
            sum += number;
            if (number > maximal) {
                maximal = number;
            }
            if (number < minimal) {
                minimal = number;
            }
        }
        double amean = (sum - minimal - maximal) / (N - 2);
        put(amean);
        return false;
    }

	private boolean minmax22() {
        int N = getInt();
        double temp, number, number1 = getDouble(), number2 = getDouble();
        //sort inc;
        if (number1 > number2) {
            temp = number1;
            number1 = number2;
            number2 = temp;
        }
        for (int i = 3; i <= N; i++) {
            number = getDouble();
            if (number < number2) {
                number2 = number;
                //sort inc;
                if (number1 > number2) {
                    temp = number1;
                    number1 = number2;
                    number2 = temp;
                }
            }
        }
        put(number1);
        put(number2);
        return false;
    }

	private boolean minmax23() {
        int N = getInt();
        double temp, number, number1 = getDouble(), number2 = getDouble(), number3 = getDouble();
        //sort dec;
        if (number1 < number2) {
            temp = number1;
            number1 = number2;
            number2 = temp;
        }
        if (number1 < number3) {
            temp = number1;
            number1= number3;
            number3 = temp;
        }
        if (number2 < number3) {
            temp = number2;
            number2 = number3;
            number3 = temp;
        }
        for (int i = 4; i <= N; i++) {
            number = getDouble();
            if (number > number3) {
                number3 = number;
                //sort dec;
                sortDec(number1, number2, number3);
                if (number1 < number2) {
                    temp = number1;
                    number1 = number2;
                    number2 = temp;
                }
                if (number1 < number3) {
                    temp = number1;
                    number1= number3;
                    number3 = temp;
                }
                if (number2 < number3) {
                    temp = number2;
                    number2 = number3;
                    number3 = temp;
                }
            }
        }
        put(number1);
        put(number2);
        put(number3);
        return false;
    }
    
    static void sortDec(Double number1, Double number2, Double number3) {
        double temp;
        if (number1 < number2) {
            temp = number1;
            number1 = number2;
            number2 = temp;
        }
        if (number1 < number3) {
            temp = number1;
            number1= number3;
            number3 = temp;
        }
        if (number2 < number3) {
            temp = number2;
            number2 = number3;
            number3 = temp;
        }
    }

	private boolean minmax24() {
        int N = getInt();
        double number1, number2, sum, maxSum = 0;
        number1 = getDouble();
        boolean inited = false;
        for (int i = 2; i <= N; i++) {
            number2 = getDouble();
            sum = number1 + number2;
            if (!inited) {
                maxSum = sum;
                inited = true;
            } else if (sum > maxSum) {
                maxSum = sum;
            }
            number1 = number2;
        }
        put(maxSum);
        return false;
    }

	private boolean minmax25() {
        int N = getInt();
        double number1, number2, zarb, minZarb = 0;
        int index1 = 0, index2 = 0;
        number1 = getDouble();
        boolean inited = false;
        for (int i = 2; i <= N; i++) {
            number2 = getDouble();
            zarb = number1 * number2;
            if (!inited) {
                minZarb = zarb;
                inited = true;
                index1 = i-1;
                index2 = i;
            } else if (zarb < minZarb) {
                minZarb = zarb;
                index1 = i-1;
                index2 = i;
            }
            number1 = number2;
        }
        put(index1);
        put(index2);
        return false;
    }

	private boolean minmax26() {
        int N = getInt();
        int number1, number2, count, maxCount = 0;
        number1 = getInt();
        count = number1 % 2 == 0 ? 1 : 0;
        for (int i = 2; i <= N; i++) {
            number2 = getInt();
            if (number2 % 2 == 0) {
                if (number1 % 2 == 0) {
                    count++;
                } else {
                    count = 1;
                }
            } else if (count != 0) {
                if (count > maxCount) {
                    maxCount = count;
                }
                count = 0;
            }
            number1 = number2;
        }
        if ( (count != 0) && (count > maxCount) ) {
            maxCount = count;
        }
        put(maxCount);
        return false;
    }

	private boolean minmax27() {
        int N = getInt();
        int number1, number2, count1, count0, maxCount1, maxCount0, index1, index0;
        number1 = getInt();
        maxCount1 = maxCount0 = 0;
        index1 = index0 = 0;
        count1 = number1 == 1 ? 1 : 0;
        count0 = number1 == 0 ? 1 : 0;
        for (int i = 2; i <= N; i++) {
            number2 = getInt();
            if (number2 == 1) {
                if (number1 == 1) {
                    count1++;
                } else {
                    count1 = 1;
                }
                if ( (count0 != 0) && (count0 > maxCount0) ) {
                    maxCount0 = count0;
                    index0 = i - maxCount0;
                }
            } else {
                if (number1 == 0) {
                    count0++;
                } else {
                    count0 = 1;
                }
                if ( (count1 != 0) && (count1 > maxCount1) ) {
                    maxCount1 = count1;
                    index1 = i - maxCount1;
                }
            }
            number1 = number2;
        }
        if ( (count0 != 0) && (count0 > maxCount0) ) {
            maxCount0 = count0;
            index0 = N - maxCount0 + 1;
        }
        if ( (count1 != 0) && (count1 > maxCount1) ) {
            maxCount1 = count1;
            index1 = N - maxCount1 + 1;
        }
        int maxCount, index;
        if (maxCount1 == maxCount0) {
            if (index1 < index0) {
                index = index1;
                maxCount = maxCount1;
            } else {
                index = index0;
                maxCount = maxCount0;
            }
        }
        else if (maxCount1 > maxCount0) {
            maxCount = maxCount1;
            index = index1;
        } else {
            maxCount = maxCount0;
            index = index0;
        }
        put(index);
        put(maxCount);
        return false;
    }

	private boolean minmax28() {
        int N = getInt();
        int number1, number2, count, index, maxCount = 0;
        number1 = getInt();
        count = number1 == 1 ? 1 : 0;
        index = 0;
        for (int i = 2; i <= N; i++) {
            number2 = getInt();
            if (number2 == 1) {
                if (number1 == 1) {
                    count++;
                } else {
                    count = 1;
                }
            } else if (count != 0) {
                if (count >= maxCount) {
                    maxCount = count;
                    index = i - maxCount;
                }
                count = 0;
            }
            number1 = number2;
        }
        if ( (count != 0) && (count >= maxCount) ) {
            maxCount = count;
            index = N - maxCount + 1;
        }
        put(index);
        put(maxCount);
        return false;
    }

	private boolean minmax29() {
        int N = getInt();
        int number1, number2, minimal, count, maxCount = 0;
        number1 = getInt();
        minimal = number1;
        count = 1;
        for (int i = 2; i <= N; i++) {
            number2 = getInt();
            if (number2 < minimal) {
                minimal = number2;
                count = 1;
                maxCount = 0;
            } else if (number2 == minimal) {
                if (number1 == minimal) {
                    count++;
                } else {
                    count = 1;
                }
            } else if (count != 0) {
                if (count > maxCount) {
                    maxCount = count;
                }
                count = 0;
            }
            number1 = number2;
        }
        if ( (count != 0) && (count > maxCount) ) {
            maxCount = count;
        }
        put(maxCount);
        return false;
    }

	private boolean minmax30() {
        int N = getInt();
        int number1, number2, maximal, count, minCount;
        number1 = getInt();
        maximal = number1;
        count = 1;
        minCount = -1;
        for (int i = 2; i <= N; i++) {
            number2 = getInt();
            if (number2 > maximal) {
                maximal = number2;
                count = 1;
                minCount = -1;
            } else if (number2 == maximal) {
                if (number1 == maximal) {
                    count++;
                } else {
                    count = 1;
                }
            } else if (count != 0) {
                if (minCount < 0) {
                    minCount = count;
                } else if (count < minCount) {
                    minCount = count;
                }
                count = 0;
            }
            number1 = number2;
        }
        if (count != 0) {
            if (minCount < 0) {
                minCount = count;
            } else if (count < minCount) {
                minCount = count;
            }
        }
        put(minCount);
        return false;
    }
}