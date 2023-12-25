import java.util.Scanner;

class ArrayTask {

	private boolean array1() {
		Scanner scanner = new Scanner(System.in);
		int N = scanner.nextInt();
		int[] arr = new int [N];
		for (int i = 0; i < N; i++) {
			arr[i] = 2 * i + 1;
		}
		for (int i = 0; i < N; i++) {
			System.out.print(arr[i] + "\t");
		}
		System.out.println();
		return false;
	}

	private boolean array2() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("N = ");
		int N = scanner.nextInt();
		int[] arr = new int [N];
		for (int i = 0; i < N; i++) {
			arr[i] = (int)Math.pow(2, i+1);
		}
		for (int i = 0; i < N; i++) {
			System.out.print(arr[i] + "\t");
		}
		System.out.println();
		return false;
	}

	private boolean array3() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("N = ");
		int N = scanner.nextInt();
		System.out.print("A = ");
		double A = scanner.nextDouble();
		System.out.print("D = ");
		double D = scanner.nextDouble();
		double[] arr = new double [N];
		for (int i = 0; i < N; i++) {
			arr[i] = A + i * D;
		}
		for (int i = 0; i < N; i++) {
			System.out.printf("%1$.2f\t", arr[i]);
		}
		System.out.println();
		return false;
	}

	private boolean array4() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("N = ");
		int N = scanner.nextInt();
		System.out.print("A = ");
		double A = scanner.nextDouble();
		System.out.print("D = ");
		double D = scanner.nextDouble();
		double[] arr = new double [N];
		for (int i = 0; i < N; i++) {
			arr[i] = A * pow(D, i);
		}
		for (int i = 0; i < N; i++) {
			System.out.printf("%1$.2f\t", arr[i]);
		}
		System.out.println();
		return false;
	}

	private boolean array5() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("N = ");
		int N = scanner.nextInt();
		int[] facts = new int [N];
		for (int i = 0; i < N; i++) {
			facts[i] = i == 1 || i == 0 ? 1 : facts[i - 1] + facts[i - 2];
		}
		for (int i = 0; i < N; i++) {
			System.out.print(facts[i] + "\t");
		}
		System.out.println();
		return false;
	}

	private boolean array6() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("N = ");
		int N = scanner.nextInt();
		System.out.print("A = ");
		int A = scanner.nextInt();
		System.out.print("B = ");
		int B = scanner.nextInt();
		int[] arr = new int [N];
		int sum = 0;
		for (int i = 0; i < N; i++) {
			if (i > 1) {
				arr[i] = sum;
			} else if (i == 0) {
				arr[i] = A;
			} else {
				arr[i] = B;
			}
			sum += arr[i];
		}
		for (int i = 0; i < N; i++) {
			System.out.print(arr[i] + "\t");
		}
		System.out.println();
		return false;
	}

	private boolean array7() {
        int N = getInt();
        double[] arr = new double [N];
        for (int i = 0; i < N; i++)
            arr[i] = getDouble();
        for (int i = N-1; i >= 0; i--)
            put(arr[i]);
        return false;
    }

	private boolean array8() {
        int N = getInt();
        int[] arr = new int [N];
        for (int i = 0; i < N; i++)
            arr[i] = getInt();
        int K = 0;
        for (int i = 0; i < N; i++) {
            if (arr[i] % 2 != 0)
            {
                put(arr[i]);
                K++;
            }
        }
        put(K);
        return false;
    }

	private boolean array9() {
        int N = getInt();
        int[] arr = new int [N];
        for (int i = 0; i < N; i++)
            arr[i] = getInt();
        int K = 0;
        for (int i = N-1; i >= 0; i--)
            if (arr[i] % 2 == 0)
            {
                put(arr[i]);
                K++;
            }
        put(K);
        return false;
    }

	private boolean array10() {
        int N = getInt();
        int[] arr = new int [N];
        for (int i = 0; i < N; i++)
            arr[i] = getInt();
        for (int i = 0; i < N; i++)
            if (arr[i] % 2 == 0)
                put(arr[i]);
        for (int i = N-1; i >= 0; i--)
            if (arr[i] % 2 != 0)
                put(arr[i]);
        return false;
    }

	private boolean array11() {
        int N = getInt();
        double[] A = new double [N];
        for (int i = 0; i < N; i++)
            A[i] = getDouble();
        int K = getInt();
        for (int i = K-1; i < N; i += K)
            put(A[i]);
        return false;
    }

	private boolean array12() {
        int N = getInt();
        double[] A = new double [N];
        for (int i = 0; i < N; i++)
            A[i] = getDouble();
        for (int i = 1; i < N; i += 2)
            put(A[i]);
        return false;
    }

	private boolean array13() {
        int N = getInt();
        double[] A = new double [N];
        for (int i = 0; i < N; i++)
            A[i] = getDouble();
        for (int i = N-1; i >= 0; i -= 2)
            put(A[i]);
        return false;
    }

	private boolean array14() {
        int N = getInt();
        double[] A = new double [N];
        for (int i = 0; i < N; i++)
            A[i] = getDouble();
        for (int i = 1; i < N; i += 2)
            put(A[i]);
        for (int i = 0; i < N; i += 2)
            put(A[i]);
        return false;
    }

	private boolean array15() {
        int N = getInt();
        double[] A = new double [N];
        for (int i = 0; i < N; i++)
            A[i] = getDouble();
        for (int i = 0; i < N; i += 2)
            put(A[i]);
        for (int i = N-1-N%2; i >= 0; i -= 2)
            put(A[i]);
        return false;
    }

	private boolean array16() {
        int N = getInt();
        double[] A = new double[N];
        for (int i = 0; i < N; i++)
            A[i] = getDouble();
        int start = 0, finish = N-1;
        while (true) {
            if (start-1 == finish) break;
            else put(A[start++]);
            
            if (start-1 == finish) break;
            else put(A[finish--]);
        }
        return false;
    }

	private boolean array17() {
        int N = getInt();
        double[] A = new double [N];
        for (int i = 0; i < N; i++)
            A[i] = getDouble();
        int start = 0, finish = N-1;
        while (true) {
            if (start-1 == finish) break;
            else put(A[start++]);
            
            if (start-1 == finish) break;
            else put(A[start++]);
            
            if (start-1 == finish) break;
            else put(A[finish--]);
            
            if (start-1 == finish) break;
            else put(A[finish--]);
        }
        return false;
    }

	private boolean array18() {
        int[] A = new int [10];
        for (int i = 0; i < 10; i++)
            A[i] = getInt();
        int element = 0;
        for (int i = 0; i < 9; i++) {
            if (A[i] < A[9])
            {
                element = A[i];
                break;
            }
        }
        put(element);
        return false;
    }

	private boolean array19() {
        int[] A = new int [10];
        for (int i = 0; i < 10; i++)
            A[i] = getInt();
        int index = -1;
        for (int i = 1; i < 9; i++)
            if ( (A[i] > A[0]) && (A[i] < A[9]) )
                index = i;
        put(1+index);
        return false;
    }

	private boolean array20() {
        int N = getInt();
        double[] arr = new double [N];
        for (int i = 0; i < N; i++)
            arr[i] = getDouble();
        int K = getInt(), L = getInt();
        double sum = 0;
        for (int i = K-1; i < L; i++)
            sum += arr[i];
        put(sum);
        return false;
    }

	private boolean array21() {
        int N = getInt();
        double[] arr = new double [N];
        for (int i = 0; i < N; i++)
            arr[i] = getDouble();
        int K = getInt(), L = getInt();
        double sum = 0;
        for (int i = K-1; i < L; i++)
            sum += arr[i];
        double amean = sum / (L - K + 1);
        put(amean);
        return false;
    }

	private boolean array22() {
        int N = getInt();
        double[] arr = new double[N];
        double sum = 0;
        for (int i = 0; i < N; i++) {
            arr[i] = getDouble();
            sum += arr[i];
        }
        int K = getInt(), L = getInt();
        for (int i = K-1; i < L; i++)
            sum -= arr[i];
        put(sum);
        return false;
    }

	private boolean array23() {
        int N = getInt();
        double[] arr = new double [N];
        double sum = 0;
        for (int i = 0; i < N; i++) {
            arr[i] = getDouble();
            sum += arr[i];
        }
        int K = getInt(), L = getInt();
        for (int i = K-1; i < L; i++)
            sum -= arr[i];
        double amean = sum / (N - (L - K + 1));
        put(amean);
        return false;
    }

	private boolean array24() {
        int N = getInt();
        int[] arr = new int [N];
        for (int i = 0; i < N; i++)
            arr[i] = getInt();
        int farq = arr[1] - arr[0];
        for (int i = 1; i < N; i++) {
            if (arr[i] - arr[i-1] != farq)
            {
                farq = 0;
                break;
            }
        }
        put(farq);
        return false;
    }

	private boolean array25() {
        int N = getInt();
        int[] arr = new int [N];
        for (int i = 0; i < N; i++)
            arr[i] = getInt();
        int maxraj = arr[1] / arr[0];
        for (int i = 1; i < N; i++)
            if ( ( arr[i] / (double)arr[i-1] - arr[i] / arr[i-1] != 0 ) ||
                    (arr[i] / arr[i-1] != maxraj) )
            {
                maxraj = 0;
                break;
            }
        put(maxraj);
        return false;
    }

	private boolean array26() {
        int N = getInt();
        int[] arr = new int [N];
        for (int i = 0; i < N; i++)
            arr[i] = getInt();
        int index = -1;
        for (int i = 1; i < N; i++)
            if ((arr[i] + arr[i-1]) % 2 == 0)
            {
                index = i;
                break;
            }
        put(index+1);
        return false;
    }

	private boolean array27() {
        int N = getInt();
        int[] arr = new int [N];
        for (int i = 0; i < N; i++)
            arr[i] = getInt();
        int index = -1;
        for (int i = 1; i < N; i++)
            if (arr[i] * arr[i-1] > 0)
            {
                index = i;
                break;
            }
        put(1+index);
        return false;
    }

	private boolean array28() {
        int N = getInt();
        double[] arr = new double [N];
        for (int i = 0; i < N; i++)
            arr[i] = getDouble();
        double minimal = 0;
        boolean initialized = false;
        for (int i = 1; i < N; i += 2) {
            if (!initialized)
            {
                minimal = arr[i];
                initialized = true;
            }
            else if (arr[i] < minimal)
                minimal = arr[i];
        }
        put(minimal);
        return false;
    }

	private boolean array29() {
        int N = getInt();
        double[] arr = new double [N];
        for (int i = 0; i < N; i++)
            arr[i] = getDouble();
        double maximal = 0;
        boolean initialized = false;
        for (int i = 0; i < N; i += 2) {
            if (!initialized)
            {
                maximal = arr[i];
                initialized = true;
            }
            else if (arr[i] > maximal)
                maximal = arr[i];
        }
        put(maximal);
        return false;
    }

	private boolean array30() {
        int N = getInt();
        double[] arr = new double [N];
        for (int i = 0; i < N; i++)
            arr[i] = getDouble();
        int count = 0;
        for (int i = 0; i < N-1; i++) {
            if (arr[i] > arr[i+1])
            {
                put(1+i);
                count++;
            }
        }
        put(count);
        return false;
    }

	private boolean array31() {
        int N = getInt();
        double[] arr = new double [N];
        for (int i = 0; i < N; i++)
            arr[i] = getDouble();
        int count = 0;
        for (int i = N-1; i > 0; i--) {
            if (arr[i] > arr[i-1])
            {
                put(1+i);
                count++;
            }
        }
        put(count);
        return false;
    }

	private boolean array32() {
        int N = getInt();
        double[] arr = new double [N];
        for (int i = 0; i < N; i++)
            arr[i] = getDouble();
        double localMinimum = 0;
        int index = -1;
        boolean found = false;
        //check the first element
        if (arr[0] < arr[1]) {
            localMinimum = arr[0];
            index = 0;
            found = true;
        }
        //check elements between if until not found
        for (int i = 1; !found && (i < N-1); i++) {
            if ( (arr[i] < arr[i-1]) && (arr[i] < arr[i+1]) )
            {
                localMinimum = arr[i];
                index = i;
                found = true;
            }
        }
        //check the last element
        if ( !found && (arr[N-1] < arr[N-2]) ) {
            localMinimum = arr[N-1];
            index = N-1;
        }
        put(1 + index);
        return false;
    }

	private boolean array33() {
        int N = getInt();
        double[] arr = new double [N];
        for (int i = 0; i < N; i++)
            arr[i] = getDouble();
        double localMaximum = 0;
        int index = -1;
        //check the first element;
        if (arr[0] > arr[1]) {
            localMaximum = arr[0];
            index = 0;
        }
        //check elements between;
        for (int i = 1; i < N-1; i++)
            if ( (arr[i] > arr[i-1]) && (arr[i] > arr[i+1]) )
            {
                localMaximum = arr[i];
                index = i;
            }
        //check the last element;
        if (arr[N-1] > arr[N-2]) {
            localMaximum = arr[N-1];
            index = N-1;
        }
        put(1+index);
        return false;
    }

	private boolean array34() {
        int N = getInt();
        double[] arr = new double [N];
        for (int i = 0; i < N; i++)
            arr[i] = getDouble();
        double maxLocalMin = 0;
        boolean initialized = false;
        if (arr[0] < arr[1]) {
            maxLocalMin = arr[0];
            initialized = true;
        }
        for (int i = 1; i < N-1; i++) {
            if ( (arr[i] < arr[i-1]) && (arr[i] < arr[i+1]) )
            {
                if (!initialized)
                {
                    maxLocalMin = arr[i];
                    initialized = true;
                }
                else if (arr[i] > maxLocalMin)
                    maxLocalMin = arr[i];
            }
        }
        if (arr[N-1] < arr[N-2]) {
            if (!initialized)
            {
                maxLocalMin = arr[N-1];
                initialized = true;
            }
            else if (arr[N-1] > maxLocalMin)
                maxLocalMin = arr[N-1];
        }
        put(maxLocalMin);
        return false;
    }

	private boolean array35() {
        int N = getInt();
        double[] arr = new double [N];
        for (int i = 0; i < N; i++)
            arr[i] = getDouble();
        double minLocalMax = 0;
        boolean inited = false;
        if (arr[0] > arr[1]) {
            minLocalMax = arr[0];
            inited = true;
        }
        for (int i = 1; i < N-1; i++) {
            if ( (arr[i] > arr[i-1]) && (arr[i] > arr[i+1]) )
            {
                if (!inited)
                {
                    minLocalMax = arr[i];
                    inited = true;
                }
                else if (arr[i] < minLocalMax)
                    minLocalMax = arr[i];
            }
        }
        if (arr[N-1] > arr[N-2]) {
            if (!inited)
            {
                minLocalMax = arr[N-1];
                inited = true;
            }
            else if (arr[N-1] < minLocalMax)
                minLocalMax = arr[N-1];
        }
        put(minLocalMax);
        return false;
    }

	private boolean array36() {
        int N = getInt();
        double[] arr = new double [N];
        for (int i = 0; i < N; i++)
            arr[i] = getDouble();
        double maximal = 0;
        boolean inited = false;
        for (int i = 1; i < N-1; i++) {
            if ( (arr[i] > arr[i-1]) && (arr[i] < arr[i+1]) ||
                    (arr[i] < arr[i-1]) && (arr[i] > arr[i+1]) )
            {
                if (!inited)
                {
                    maximal = arr[i];
                    inited = true;
                }
                else if (arr[i] > maximal)
                    maximal =  arr[i];
            }
        }
        put(maximal);
        return false;
    }

	private boolean array37() {
        int N = getInt();
        double[] arr = new double [N];
        for (int i = 0; i < N; i++)
            arr[i] = getDouble();
        int series = 0;
        boolean started = false;
        for (int i = 1; i < N; i++) {
            if (arr[i-1] > arr[i])
            {
                if (started)
                {
                    series++;
                    started = false;
                }
            }
            else
                started = true;
        }
        series += started ? 1 : 0;
        put(series);
        return false;
    }

	private boolean array38() {
        int N = getInt();
        double[] arr = new double [N];
        for (int i = 0; i < N; i++)
            arr[i] = getDouble();
        int series = 0;
        boolean started = false;
        for (int i = 1; i < N; i++) {
            if (arr[i-1] < arr[i])
            {
                if (started)
                {
                    series++;
                    started = false;
                }
            }
            else
                started = true;
        }
        series += started ? 1 : 0;
        put(series);
        return false;
    }

	private boolean array39() {
        int N = getInt();
        double[] arr = new double [N];
        for (int i = 0; i < N; i++)
            arr[i] = getDouble();
        boolean isProgress = false, isRegress = false;
        int progresses = 0, regresses = 0;
        for (int i = 1; i < N; i++) {
            if (arr[i-1] < arr[i])
            {
                isProgress = true;
                if (isRegress)
                {
                    regresses++;
                    isRegress = false;
                }
            }
            else
            {
                isRegress = true;
                if (isProgress)
                {
                    progresses++;
                    isProgress = false;
                }
            }
        }
        progresses += isProgress ? 1 : 0;
        regresses += isRegress ? 1 : 0;
        put(progresses + regresses);
        return false;
    }

	private boolean array40() {
        double R = getDouble();
        int N = getInt();
        double[] A = new double [N];
        for (int i = 0; i < N; i++)
            A[i] = getDouble();
        double distance, minDistance = Math.abs(R-A[0]);
        int index = 0;
        for (int i = 1; i < N; i++) {
            distance = Math.abs(R - A[i]);
            if (distance < minDistance)
            {
                minDistance = distance;
                index = i;
            }
        }
        put(A[index]);
        return false;
    }

	private boolean array41() {
        int N = getInt();
        double[] arr = new double [N];
        for (int i = 0; i < N; i++)
            arr[i] = getDouble();
        double sum, maxSum = arr[0] + arr[1];
        int index = 0;
        for (int i = 1; i < N-1; i++) {
            sum = arr[i] + arr[i+1];
            if (sum > maxSum)
            {
                maxSum = sum;
                index = i;
            }
        }
        put(arr[index]);
        put(arr[index+1]);
        return false;
    }

	private boolean array42() {
        double R = getDouble();
        int N = getInt();
        double[] arr = new double [N];
        for (int i = 0; i < N; i++)
            arr[i] = getDouble();
        double sum = arr[0] + arr[1];
        double minDistance = Math.abs(R - sum);
        int index = 0;
        for (int i = 1; i < N-1; i++) {
            sum = arr[i] + arr[i+1];
            if (Math.abs(R - sum) < minDistance)
            {
                minDistance = Math.abs(R - sum);
                index = i;
            }
        }
        put(arr[index]);
        put(arr[index+1]);
        return false;
    }

	private boolean array43() {
        int N = getInt();
        int[] arr = new int [N];
        for (int i = 0; i < N; i++)
            arr[i] = getInt();
        int series = 1;
        for (int i = 1; i < N; i++)
            if (arr[i] != arr[i-1])
                series++;
        put(series);
        return false;
    }

	private boolean array44() {
        int N = getInt();
        int[] arr = new int [N];
        for (int i = 0; i < N; i++)
            arr[i] = getInt();
        int one = -1, two = 0;
        boolean found = false;
        while (!found) {
            one++;
            for (two = one+1; two < N; two++)
                if (arr[one] == arr[two])
                {
                    found = true;
                    break;
                }
        }
        put(one+1);
        put(two+1);
        return false;
    }

	private boolean array45() {
        int N = getInt();
        double[] arr = new double [N];
        for (int i = 0; i < N; i++)
            arr[i] = getDouble();
        int one = 0, two = 1;
        double distance, minDistance = Math.abs(arr[one] - arr[two]);
        int lastOne = one, lastTwo = two;
        while (one < N) {
            for (two = 0; two < N; two++)
            {
                if (one == two)
                    continue;
                distance = Math.abs(arr[one] - arr[two]);
                if (distance < minDistance)
                {
                    minDistance = distance;
                    lastOne = one;
                    lastTwo = two;
                }
            }
            one++;
        }
        if (lastOne > lastTwo) {
            int temp = lastOne;
            lastOne = lastTwo;
            lastTwo = temp;
        }
        put(lastOne+1);
        put(lastTwo+1);
        return false;
    }

	private boolean array46() {
        double R = getDouble();
        int N = getInt();
        double[] arr = new double [N];
        for (int i = 0; i < N; i++)
            arr[i] = getDouble();
        int one = 0, two = 1;
        int lastOne = one, lastTwo = two;
        double sum, closerSum = Math.abs(arr[one] + arr[two] - R);
        while (one < N) {
            for (two = 0; two < N; two++)
            {
                if (one == two)
                    continue;
                sum = arr[one] + arr[two];
                if (Math.abs(R - sum) < closerSum)
                {
                    closerSum = Math.abs(R - sum);
                    lastOne = one;
                    lastTwo = two;
                }
            }
            one++;
        }
        if (lastOne > lastTwo) {
            int temp = lastOne;
            lastOne = lastTwo;
            lastTwo = temp;
        }
        put(arr[lastOne]);
        put(arr[lastTwo]);
        return false;
    }

	private boolean array47() {
        int N = getInt();
        int[] arr = new int [N];
        boolean[] hast = new boolean [101];
        for (int i = 0; i <= 100; i++)
            hast[i] = false;
        for (int i = 0; i < N; i++) {
            arr[i] = getInt();
            hast[arr[i]] = true;
        }
        int differents = 0;
        for (int i = 1; i <= 100; i++)
            if (hast[i])
                differents++;
        put(differents);
        return false;
    }

	private boolean array48() {
        int N = getInt();
        int[] arr = new int [N];
        int[] count = new int [101];
        for (int i = 0; i <= 100; i++)
            count[i] = 0;
        for (int i = 0; i < N; i++) {
            arr[i] = getInt();
            count[arr[i]]++;
        }
        int maxCount = 0;
        for (int i = 1; i <= 100; i++) {
            if (count[i] > 0)
            {
                if (maxCount == 0)
                    maxCount = count[i];
                else if (count[i] > maxCount)
                    maxCount = count[i];
            }
        }
        put(maxCount);
        return false;
    }

	private boolean array49() {
        int N = getInt();
        int[] arr = new int [N];
        boolean[] hast = new boolean [N+1];
        for (int i = 0; i <= N; i++)
            hast[i] = false;
        int errorNo = -1;
        for (int i = 0; i < N; i++) {
            arr[i] = getInt();
            if ( (arr[i] >= 1) && (arr[i] <= N) )
            {
                if (hast[arr[i]])
                {
                    errorNo = i;
                    break;
                }
                else
                    hast[arr[i]] = true;
            }
            else
            {
                errorNo = i;
                break;
            }
        }
        errorNo++;
        put(errorNo);
        return false;
    }

	private boolean array50() {
        int N = getInt();
        int[] arr = new int [N];
        int count = 0;
        for (int i = 0; i < N; i++) {
            arr[i] = getInt();
            for (int j = 0; j < i; j++)
                if (arr[i] < arr[j])
                    count++;
        }
        put(count);
        return false;
    }

	private boolean array51() {
        int N = getInt();
        double[] A = new double [N];
        for (int i = 0; i < N; i++)
            A[i] = getDouble();
        double[] B = new double [N];
        for (int i = 0; i < N; i++)
            B[i] = getDouble();
        double temp;
        for (int i = 0; i < N; i++) {
            temp = A[i];
            A[i] = B[i];
            B[i] = temp;
        }
        for (int i = 0; i < N; i++)
            put(A[i]);
        for (int i = 0; i < N; i++)
            put(B[i]);
        return false;
    }

	private boolean array52() {
        int N = getInt();
        double[] A = new double [N];
        double[] B = new double [N];
        for (int i = 0; i < N; i++) {
            A[i] = getDouble();
            B[i] = A[i] < 5 ? 2 * A[i] : A[i] / 2;
        }
        for (int i = 0; i < N; i++)
            put(B[i]);
        return false;
    }

	private boolean array53() {
        int N = getInt();
        double[] A = new double [N];
        for (int i = 0; i < N; i++)
            A[i] = getDouble();
        double[] B = new double [N];
        for (int i = 0; i < N; i++)
            B[i] = getDouble();
        double[] C = new double [N];
        for (int i = 0; i < N; i++)
            C[i] = A[i] > B[i] ? A[i] : B[i];
        for (int i = 0; i < N; i++)
            put(C[i]);
        return false;
    }

	private boolean array54() {
        int N = getInt();
        int[] A = new int [N];
        int size = 0;
        for (int i = 0; i < N; i++) {
            A[i] = getInt();
            if (A[i] % 2 == 0)
                size++;
        }
        int[] B = new int [size];
        int index = 0;
        for (int i = 0; i < N; i++)
            if (A[i] % 2 == 0)
                B[index++] = A[i];
        put(size);
        for (int i = 0; i < size; i++)
            put(B[i]);
        return false;
    }

	private boolean array55() {
        int N = getInt();
        int[] A = new int [N];
        int size = N/2 + N%2;
        int[] B = new int [size];
        int index = 0;
        for (int i = 0; i < N; i++) {
            A[i] = getInt();
            if (i % 2 == 0)
                B[index++] = A[i];
        }
        put(size);
        for (int i = 0; i < size; i++)
            put(B[i]);
        return false;
    }

	private boolean array56() {
        int N = getInt();
        int[] A = new int [N];
        for (int i = 0; i < N; i++)
            A[i] = getInt();
        int size = N / 3;
        int[] B = new int [size];
        int index = 0;
        for (int i = 2; i < N; i += 3)
            B[index++] = A[i];
        put(size);
        for (int i = 0; i < size; i++)
            put(B[i]);
        return false;
    }

	private boolean array57() {
        int N = getInt();
        int[] A = new int [N];
        for (int i = 0; i < N; i++)
            A[i] = getInt();
        int[] B = new int [N];
        int index = 0;
        for (int i = 1; i < N; i += 2)
            B[index++] = A[i];
        for (int i = 0; i < N; i += 2)
            B[index++] = A[i];
        for (int i = 0; i < N; i++)
            put(B[i]);
        return false;
    }

	private boolean array58() {
        int N = getInt();
        double[] A = new double [N];
        double[] B = new double [N];
        double sum = 0;
        for (int i = 0; i < N; i++) {
            A[i] = getDouble();
            sum += A[i];
            B[i] = sum;
        }
        for (int i = 0; i < N; i++)
            put(B[i]);
        return false;
    }

	private boolean array59() {
        int N = getInt();
        double[] A = new double [N];
        double[] B = new double [N];
        double sum = 0;
        for (int i = 0; i < N; i++) {
            A[i] = getDouble();
            sum += A[i];
            B[i] = sum / (1+i);
        }
        for (int i = 0; i < N; i++)
            put(B[i]);
        return false;
    }

	private boolean array60() {
        int N = getInt();
        double[] A = new double [N];
        for (int i = 0; i < N; i++)
            A[i] = getDouble();
        double[] B = new double [N];
        double sum = 0;
        for (int i = N-1; i >= 0; i--) {
            sum += A[i];
            B[i] = sum;
        }
        for (int i = 0; i < N; i++)
            put(B[i]);
        return false;
    }

	private boolean array61() {
        int N = getInt();
        double[] A = new double [N];
        for (int i = 0; i < N; i++)
            A[i] = getDouble();
        double[] B = new double [N];
        double sum = 0;
        int count = 0;
        for (int i = N-1; i >= 0; i--) {
            count++;
            sum += A[i];
            B[i] = sum / count;
        }
        for (int i = 0; i < N; i++)
            put(B[i]);
        return false;
    }

	private boolean array62() {
        int N = getInt();
        double[] A = new double [N];
        int pluses = 0, minuses = 0;
        for (int i = 0; i < N; i++) {
            A[i] = getDouble();
            if (A[i] > 0)
                pluses++;
            else
                minuses++;
        }
        double[] B = new double [pluses];
        double[] C = new double [minuses];
        for (int i = 0, iB = 0, iC = 0; i < N; i++) {
            if (A[i] > 0)
                B[iB++] = A[i];
            else
                C[iC++] = A[i];
        }
        put(pluses);
        for (int i = 0; i < pluses; i++)
            put(B[i]);
        put(minuses);
        for (int i = 0; i < minuses; i++)
            put(C[i]);
        return false;
    }

	private boolean array63() {
        double[] A = new double [5], B = new double [5], C = new double [10];
        for (int i = 0; i < 5; i++)
            A[i] = getDouble();
        for (int i = 0; i < 5; i++)
            B[i] = getDouble();
        for (int i = 0, iA = 0, iB = 0; i < 10; i++) {
            if (iA >= 5)
                C[i] = B[iB++];
            else if (iB >= 5)
                C[i] = A[iA++];
            else
                C[i] = A[iA] < B[iB] ? A[iA++] : B[iB++];
        }
        for (int i = 0; i < 10; i++)
            put(C[i]);
        return false;
    }

	private boolean array64() {
        int NA = getInt();
        int[] A = new int [NA];
        for (int i = 0; i < NA; i++)
            A[i] = getInt();
        
        int NB = getInt();
        int[] B = new int [NB];
        for (int i = 0; i < NB; i++)
            B[i] = getInt();
        
        int NC = getInt();
        int[] C = new int [NC];
        for (int i = 0; i < NC; i++)
            C[i] = getInt();
        
        //yakjoyakunii massivhoi A va B: A + B = T;
        int NT = NA + NB;
        int[] T = new int [NT];
        for (int i = 0, iA = 0, iB = 0; i < NT; i++) {
            if (iA >= NA)
                T[i] = B[iB++];
            else if (iB >= NB)
                T[i] = A[iA++];
            else
                T[i] = A[iA] > B[iB] ? A[iA++] : B[iB++];
        }
        
        //yakjoyakunii massivhoi T va C: T + C = D;
        int ND = NT + NC;
        int[] D = new int [ND];
        for (int i = 0, iT = 0, iC = 0; i < ND; i++) {
            if (iT >= NT)
                D[i] = C[iC++];
            else if (iC >= NC)
                D[i] = T[iT++];
            else
                D[i] = T[iT] > C[iC] ? T[iT++] : C[iC++];
        }
        
        for (int i = 0; i < ND; i++)
            put(D[i]);
    }

	private boolean array65() {
        int N = getInt();
        double[] A = new double [N];
        for (int i = 0; i < N; i++)
            A[i] = getDouble();
        int K = getInt();
        double AK = A[K-1];
        for (int i = 0; i < N; i++)
            A[i] += AK;
        for (int i = 0; i < N; i++)
            put(A[i]);
        return false;
    }

	private boolean array66() {
        int N = getInt();
        int[] arr = new int [N];
        int index = -1;
        for (int i = 0; i < N; i++) {
            arr[i] = getInt();
            if (arr[i] % 2 == 0)
            {
                if (index < 0)
                    index = i;
                else
                    arr[i] += arr[index];
            }
        }
        if (index >= 0)
            arr[index] += arr[index];
        for (int i = 0; i < N; i++)
            put(arr[i]);
        return false;
    }

	private boolean array67() {
        int N = getInt();
        int[] arr = new int [N];
        int index = -1;
        for (int i = 0; i < N; i++) {
            arr[i] = getInt();
            if (arr[i] % 2 != 0)
                index = i;
        }
        if (index >= 0) {
            int element = arr[index];
            for (int i = 0; i < N; i++)
                if (arr[i] % 2 != 0)
                    arr[i] += element;
        }
        for (int i = 0; i < N; i++)
            put(arr[i]);
        return false;
    }

	private boolean array68() {
        int N = getInt();
        double[] arr = new double [N];
        int maxIndex = 0, minIndex = 0;
        for (int i = 0; i < N; i++) {
            arr[i] = getDouble();
            if (arr[i] > arr[maxIndex])
                maxIndex = i;
            if (arr[i] < arr[minIndex])
                minIndex = i;
        }
        double temp = arr[maxIndex];
        arr[maxIndex] = arr[minIndex];
        arr[minIndex] = temp;
        for (int i = 0; i < N; i++)
            put(arr[i]);
        return false;
    }

	private boolean array69() {
        int N = getInt();
        double[] arr = new double [N];
        for (int i = 0; i < N; i++)
            arr[i] = getDouble();
        double temp;
        for (int i = 1; i < N; i += 2) {
            temp = arr[i];
            arr[i] = arr[i-1];
            arr[i-1] = temp;
        }
        for (int i = 0; i < N; i++)
            put(arr[i]);
        return false;
    }

	private boolean array70() {
        int N = getInt();
        double[] arr = new double [N];
        for (int i = 0; i < N; i++)
            arr[i] = getDouble();
        double temp;
        for (int a = 0, b = N/2 + N%2; b < N; a++, b++) {
            temp = arr[a];
            arr[a] = arr[b];
            arr[b] = temp;
        }
        for (int i = 0; i < N; i++)
            put(arr[i]);
        return false;
    }

	private boolean array71() {
        int N = getInt();
        double[] arr = new double [N];
        for (int i = 0; i < N; i++)
            arr[i] = getDouble();
        double temp;
        for (int a = 0, b = N-1; a < b; a++, b--) {
            temp = arr[a];
            arr[a] = arr[b];
            arr[b] = temp;
        }
        for (int i = 0; i < N; i++)
            put(arr[i]);
        return false;
    }

	private boolean array72() {
        int N = getInt();
        double[] A = new double [N];
        for (int i = 0; i < N; i++)
            A[i] = getDouble();
        int K = getInt(), L = getInt();
        double temp;
        for (int a = K-1, b = L-1; a < b; a++, b--) {
            temp = A[a];
            A[a] = A[b];
            A[b] = temp;
        }
        for (int i = 0; i < N; i++)
            put(A[i]);
        return false;
    }

	private boolean array73() {
        int N = getInt();
        double[] A = new double [N];
        for (int i = 0; i < N; i++)
            A[i] = getDouble();
        int K = getInt(), L = getInt();
        double temp;
        for (int a = K, b = L-2; a < b; a++, b--) {
            temp = A[a];
            A[a] = A[b];
            A[b] = temp;
        }
        for (int i = 0; i < N; i++)
            put(A[i]);
        return false;
    }

	private boolean array74() {
        int N = getInt();
        double[] arr = new double [N];
        int minIndex = 0, maxIndex = 0;
        for (int i = 0; i < N; i++) {
            arr[i] = getDouble();
            if (arr[i] > arr[maxIndex])
                maxIndex = i;
            if (arr[i] < arr[minIndex])
                minIndex = i;
        }
        int from = minIndex < maxIndex ? minIndex : maxIndex;
        int to = minIndex > maxIndex ? minIndex : maxIndex;
        for (from++; from < to; from++)
            arr[from] = 0;
        for (int i = 0; i < N; i++)
            put(arr[i]);
        return false;
    }

	private boolean array75() {
        int N = getInt();
        double[] arr = new double [N];
        int minIndex = 0, maxIndex = 0;
        for (int i = 0; i < N; i++) {
            arr[i] = getDouble();
            if (arr[i] > arr[maxIndex])
                maxIndex = i;
            if (arr[i] < arr[minIndex])
                minIndex = i;
        }
        int from = minIndex < maxIndex ? minIndex : maxIndex;
        int to = minIndex > maxIndex ? minIndex : maxIndex;
        double temp;
        for (; from < to; from++, to--) {
            temp = arr[from];
            arr[from] = arr[to];
            arr[to] = temp;
        }
        for (int i = 0; i < N; i++)
            put(arr[i]);
        return false;
    }

	private boolean array76() {
        int N = getInt();
        double[] arr = new double [N], copy = new double [N];
        for (int i = 0; i < N; i++) {
            arr[i] = getDouble();
            copy[i] = arr[i];
        }
        if (copy[0] > copy[1])
            arr[0] = 0;
        for (int i = 1; i < N-1; i++) {
            if ( (copy[i] > copy[i-1]) && (copy[i] > copy[i+1]) )
                arr[i] = 0;
        }
        if (copy[N-1] > copy[N-2])
            arr[N-1] = 0;
        for (int i = 0; i < N; i++)
            put(arr[i]);
        return false;
    }

	private boolean array77() {
        int N = getInt();
        double[] arr = new double [N], copy = new double [N];
        for (int i = 0; i < N; i++) {
            arr[i] = getDouble();
            copy[i] = arr[i];
        }
        if (copy[0] < copy[1])
            arr[0] *= arr[0];
        for (int i = 1; i < N-1; i++) {
            if ( (copy[i] < copy[i-1]) && (copy[i] < copy[i+1]) )
                arr[i] *= arr[i];
        }
        if (copy[N-1] < copy[N-2])
            arr[N-1] *= arr[N-1];
        for (int i = 0; i < N; i++)
            put(arr[i]);
        return false;
    }

	private boolean array78() {
        int N = getInt();
        double[] arr = new double [N], copy = new double [N];
        for (int i = 0; i < N; i++) {
            arr[i] = getDouble();
            copy[i] = arr[i];
        }
        arr[0] = (copy[0] + copy[1]) / 2;
        for (int i = 1; i < N-1; i++)
            arr[i] = (copy[i-1] + copy[i] + copy[i+1]) / 3;
        arr[N-1] = (copy[N-1] + copy[N-2]) / 2;
        for (int i = 0; i < N; i++)
            put(arr[i]);
        return false;
    }

	private boolean array79() {
        int N = getInt();
        double[] arr = new double [N];
        for (int i = 0; i < N; i++)
            arr[i] = getDouble();
        for (int i = N-1; i > 0; i--)
            arr[i] = arr[i-1];
        arr[0] = 0;
        for (int i = 0; i < N; i++)
            put(arr[i]);
        return false;
    }

	private boolean array80() {
        int N = getInt();
        double[] arr = new double [N];
        for (int i = 0; i < N; i++)
            arr[i] = getDouble();
        for (int i = 0; i < N-1; i++)
            arr[i] = arr[i+1];
        arr[N-1] = 0;
        for (int i = 0; i < N; i++)
            put(arr[i]);
        return false;
    }

	private boolean array81() {
        int N = getInt();
        double[] arr = new double [N];
        for (int i = 0; i < N; i++)
            arr[i] = getDouble();
        int K = getInt();
        for (int i = N-1; i >= K; i--) {
            arr[i] = arr[i - K];
        }
        for (int i = K-1; i >= 0; i--) {
            arr[i] = 0;
        }
        for (int i = 0; i < N; i++) {
            put(arr[i]);
        }
        return false;
    }

	private boolean array82() {
        int N = getInt();
        double[] arr = new double [N];
        for (int i = 0; i < N; i++) {
            arr[i] = getDouble();
        }
        int K = getInt();
        for (int i = 0; i < N-K; i++) {
            arr[i] = arr[i+K];
        }
        for (int i = N-K; i < N; i++) {
            arr[i] = 0;
        }
        for (int i = 0; i < N; i++) {
            put(arr[i]);
        }
        return false;
    }

	private boolean array83() {
        int N = getInt();
        double[] arr = new double [N];
        for (int i = 0; i < N; i++) {
            arr[i] = getDouble();
        }
        double temp = arr[N-1];
        for (int i = N-1; i > 0; i--) {
            arr[i] = arr[i-1];
        }
        arr[0] = temp;
        for (int i = 0; i < N; i++) {
            put(arr[i]);
        }
        return false;
    }

	private boolean array84() {
        int N = getInt();
        double[] arr = new double [N];
        for (int i = 0; i < N; i++) {
            arr[i] = getDouble();
        }
        double temp = arr[0];
        for (int i = 0; i < N-1; i++) {
            arr[i] = arr[i+1];
        }
        arr[N-1] = temp;
        for (int i = 0; i < N; i++) {
            put(arr[i]);
        }
        return false;
    }

	private boolean array85() {
        int N = getInt();
        double[] arr = new double [N];
        for (int i = 0; i < N; i++) {
            arr[i] = getDouble();
        }
        int K = getInt();
        double[] temp = new double [K];
        for (int i = N-K, j = 0; i < N; i++, j++) {
            temp[j] = arr[i];
        }
        for (int i = N-1; i >= K; i--) {
            arr[i] = arr[i-K];
        }
        System.arraycopy(temp, 0, arr, 0, K);
        /*
        for (int i = 0; i < K; i++) {
            arr[i] = temp[i];
        }
        */
        for (int i = 0; i < N; i++) {
            put(arr[i]);
        }
        return false;
    }

	private boolean array86() {
        int N = getInt();
        double[] arr = new double [N];
        for (int i = 0; i < N; i++) {
            arr[i] = getDouble();
        }
        int K = getInt();
        double[] temp = new double [K];
        System.arraycopy(arr, 0, temp, 0, K);
        /*
        for (int i = 0; i < K; i++) {
            temp[i] = arr[i];
        }
        */
        for (int i = 0; i < N-K; i++) {
            arr[i] = arr[i+K];
        }
        for (int i = 0, j = N-K; i < K; i++, j++) {
            arr[j] = temp[i];
        }
        for (int i = 0; i < N; i++) {
            put(arr[i]);
        }
        return false;
    }

	private boolean array87() {
        int N = getInt();
        double[] arr = new double [N];
        boolean inited = false;
        int index = -1;
        for (int i = 0; i < N; i++) {
            arr[i] = getDouble();
            if ( (index < 0) && (i != 0) ) {
                if (arr[i] > arr[0]) {
                    index = i;
                }
            }
        }
        index = index < 0 ? N-1 : index-1;
        double temp = arr[0];
        for (int i = 0; i < index; i++) {
            arr[i] = arr[i+1];
        }
        arr[index] = temp;
        for (int i = 0; i < N; i++) {
            put(arr[i]);
        }
        return false;
    }

	private boolean array88() {
        int N = getInt();
        double[] arr = new double [N];
        for (int i = 0; i < N; i++) {
            arr[i] = getDouble();
        }
        int index = N;
        for (int i = N-2; i >= 0; i--) {
            if (arr[i] < arr[N-1]) {
                index = i + 1;
                break;
            }
        }
        if (index == N) {
            index = 0;
        }
        double temp = arr[N-1];
        for (int i = N-1; i > index; i--) {
            arr[i] = arr[i-1];
        }
        arr[index] = temp;
        for (int i = 0; i < N; i++) {
            put(arr[i]);
        }
        return false;
    }

	private boolean array89() {
        int N = getInt();
        double[] arr = new double [N];
        for (int i = 0; i < N; i++) {
            arr[i] = getDouble();
        }
        double temp;
        for (int i = 0; i < N-1; i++) {
            if (arr[i] < arr[i+1]) {
                temp = arr[i];
                arr[i] = arr[i+1];
                arr[i+1] = temp;
            }
        }
        for (int i = N-1; i > 0; i--) {
            if (arr[i] > arr[i-1]) {
                temp = arr[i];
                arr[i] = arr[i-1];
                arr[i-1] = temp;
            }
        }
        for (int i = 0; i < N; i++) {
            put(arr[i]);
        }
        return false;
    }

	private boolean array90() {
        int N = getInt();
        double[] arr = new double [N];
        for (int i = 0; i < N; i++) {
            arr[i] = getDouble();
        }
        int K = getInt();
        double[] temp = new double [N-1];
        for (int iA = 0, iT = 0; iA < N; iA++) {
            if (iA == K-1) {
                continue;
            }
            temp[iT++] = arr[iA];
        }
        arr = temp;
        N--;
        for (int i = 0; i < N; i++) {
            put(arr[i]);
        }
        return false;
    }

	private boolean array91() {
        int N = getInt();
        double[] arr = new double [N];
        for (int i = 0; i < N; i++) {
            arr[i] = getDouble();
        }
        int K = getInt(), L = getInt();
        int size = N - (L - K + 1);
        double[] temp = new double [size];
        for (int i = 0, j = 0; i < N; i++) {
            if ( (i >= K-1) && (i <= L-1) ) {
                continue;
            }
            temp[j++] = arr[i];
        }
        arr = temp;
        N = size;
        put(N);
        for (int i = 0; i < N; i++) {
            put(arr[i]);
        }
        return false;
    }

	private boolean array92() {
        int N = getInt();
        int[] arr = new int [N];
        int size = N;
        for (int i = 0; i < N; i++) {
            arr[i] = getInt();
            if (arr[i] % 2 != 0) {
                size--;
            }
        }
        if (size != N) {
            int[] temp = new int [size];
            for (int i = 0, j = 0; i < N; i++) {
                if (arr[i] % 2 != 0) {
                    continue;
                }
                temp[j++] = arr[i];
            }
            arr = temp;
            N = size;
        }
        put(N);
        for (int i = 0; i < N; i++) {
            put(arr[i]);
        }
        return false;
    }

	private boolean array93() {
        int N = getInt();
        int[] arr = new int [N];
        for (int i = 0; i < N; i++) {
            arr[i] = getInt();
        }
        int size = N/2 + N%2;
        int[] temp = new int [size];
        for (int i = 0, j = 0; i < N; i += 2, j++) {
            temp[j] = arr[i];
        }
        arr = temp;
        N = size;
        for (int i = 0; i < N; i++) {
            put(arr[i]);
        }
        return false;
    }

	private boolean array94() {
        int N = getInt();
        int[] arr = new int [N];
        for (int i = 0; i < N; i++) {
            arr[i] = getInt();
        }
        int size = N/2;
        int[] temp = new int [size];
        for (int i = 1, j = 0; i < N; i += 2, j++) {
            temp[j] = arr[i];
        }
        arr = temp;
        N = size;
        for (int i = 0; i < N; i++) {
            put(arr[i]);
        }
        return false;
    }

	private boolean array95() {
        int N = getInt();
        int[] arr = new int [N];
        int size = 1;
        for (int i = 0; i < N; i++) {
            arr[i] = getInt();
            if ( (i > 0) && (arr[i] != arr[i-1]) ) {
                size++;
            }
        }
        int[] temp = new int [size];
        temp[0] = arr[0];
        for (int i = 1, j = 1; i < N; i++) {
            if (arr[i] != arr[i-1]) {
                temp[j++] = arr[i];
            }
        }
        arr = temp;
        N = size;
        for (int i = 0; i < N; i++) {
            put(arr[i]);
        }
        return false;
    }

	private boolean array96() {
        int N = getInt();
        int[] arr = new int [N];
        boolean[] hast = new boolean [101];
        for (int i = 1; i <= 100; i++) {
            hast[i] = false;
        }
        int size = 0;
        for (int i = 0; i < N; i++) {
            arr[i] = getInt();
            if (!hast[arr[i]]) {
                hast[arr[i]] = true;
                size++;
            }
        }
        int[] temp = new int [size];
        for (int i = 0, j = 0; i < N; i++) {
            if (hast[arr[i]]) {
                temp[j++] = arr[i];
                hast[arr[i]] = false;
            }
        }
        arr = temp;
        N = size;
        for (int i = 0; i < N; i++) {
            put(arr[i]);
        }
        return false;
    }

	private boolean array97() {
        int N = getInt();
        int[] arr = new int [N];
        boolean[] hast = new boolean [101];
        for (int i = 1; i <= 100; i++) {
            hast[i] = false;
        }
        int size = 0;
        for (int i = 0; i < N; i++) {
            arr[i] = getInt();
            if (!hast[arr[i]]) {
                hast[arr[i]] = true;
                size++;
            }
        }
        int[] temp = new int [size];
        for (int i = N-1, j = size-1; i >= 0; i--) {
            if (hast[arr[i]]) {
                temp[j--] = arr[i];
                hast[arr[i]] = false;
            }
        }
        arr = temp;
        N = size;
        for (int i = 0; i < N; i++) {
            put(arr[i]);
        }
        return false;
    }

	private boolean array98() {
        int[] count = new int [101];
        for (int i = 1; i <= 100; i++) {
            count[i] = 0;
        }
        int N = getInt();
        int[] arr = new int [N];
        for (int i = 0; i < N; i++) {
            arr[i] = getInt();
            count[arr[i]]++;
        }
        int size = 0;
        for (int i = 1; i <= 100; i++) {
            if (count[i] >= 3) {
                size += count[i];
            }
        }
        int[] temp = new int [size];
        for (int i = 0, j = 0; i < N; i++) {
            if (count[arr[i]] < 3) {
                continue;
            }
            temp[j++] = arr[i];
        }
        arr = temp;
        N = size;
        put(N);
        for (int i = 0; i < N; i++) {
            put(arr[i]);
        }
        return false;
    }

	private boolean array99() {
        int[] count = new int [101];
        for (int i = 1; i <= 100; i++) {
            count[i] = 0;
        }
        int N = getInt();
        int[] arr = new int [N];
        for (int i = 0; i < N; i++) {
            arr[i] = getInt();
            count[arr[i]]++;
        }
        int size = N;
        for (int i = 1; i <= 100; i++) {
            if (count[i] > 2) {
                size -= count[i];
            }
        }
        int[] temp = new int [size];
        for (int i = 0, j = 0; i < N; i++) {
            if (count[arr[i]] > 2) {
                continue;
            }
            temp[j++] = arr[i];
        }
        arr = temp;
        N = size;
        put(N);
        for (int i = 0; i < N; i++) {
            put(arr[i]);
        }
        return false;
    }

	private boolean array100() {
        int[] count = new int [101];
        for (int i = 1; i <= 100; i++) {
            count[i] = 0;
        }
        int N = getInt();
        int[] arr = new int [N];
        for (int i = 0; i < N; i++) {
            arr[i] = getInt();
            count[arr[i]]++;
        }
        int size = N;
        for (int i = 1; i <= 100; i++) {
            if (count[i] == 2) {
                size -= 2;
            }
        }
        int[] temp = new int [size];
        for (int i = 0, j = 0; i < N; i++) {
            if (count[arr[i]] == 2) {
                continue;
            }
            temp[j++] = arr[i];
        }
        arr = temp;
        temp = null;
        N = size;
        put(N);
        for (int i = 0; i < N; i++) {
            put(arr[i]);
        }
        return false;
    }

	private boolean array101() {
        int N = getInt();
        double[] arr = new double [N];
        for (int i = 0; i < N; i++) {
            arr[i] = getDouble();
        }
        int K = getInt();
        double[] temp = new double [N+1];
        for (int i = 0, j = 0; i < N; i++) {
            if (i == K-1) {
                temp[j++] = 0;
            }
            temp[j++] = arr[i];
        }
        arr = temp;
        temp = null;
        N++;
        for (int i = 0; i < N; i++) {
            put(arr[i]);
        }
        return false;
    }

	private boolean array102() {
        int N = getInt();
        double[] arr = new double [N];
        for (int i = 0; i < N; i++) {
            arr[i] = getDouble();
        }
        int K = getInt();
        double[] temp = new double [N+1];
        for (int i = 0, j = 0; i < N; i++) {
            temp[j++] = arr[i];
            if (i == K-1) {
                temp[j++] = 0;
            }
        }
        arr = temp;
        temp = null;
        N++;
        for (int i = 0; i < N; i++) {
            put(arr[i]);
        }
        return false;
    }

	private boolean array103() {
        int N = getInt();
        double[] arr = new double [N];
        int maximalIndex = -1, minimalIndex = -1;
        for (int i = 0; i < N; i++) {
            arr[i] = getDouble();
            
            if (maximalIndex < 0) {
                maximalIndex = i;
            } else if (arr[i] > arr[maximalIndex]) {
                maximalIndex = i;
            }
            
            if (minimalIndex < 0) {
                minimalIndex = i;
            } else if (arr[i] < arr[minimalIndex]) {
                minimalIndex = i;
            }
        }
        double[] temp = new double [N+2];
        for (int i = 0, j = 0; i < N; i++) {
            if (i == minimalIndex) {
                temp[j++] = 0;
            }
            temp[j++] = arr[i];
            if (i == maximalIndex) {
                temp[j++] = 0;
            }
        }
        arr = temp;
        temp = null;
        N += 2;
        for (int i = 0; i < N; i++) {
            put(arr[i]);
        }
        return false;
    }

	private boolean array104() {
        int N = getInt();
        double[] arr = new double [N];
        for (int i = 0; i < N; i++) {
            arr[i] = getDouble();
        }
        int K = getInt();
        int M = getInt();
        double[] temp = new double [N+M];
        for (int i = 0, j = 0; i < N; i++) {
            if (i == K-1) {
                for (int k = 0; k < M; k++) {
                    temp[j++] = 0;
                }
            }
            temp[j++] = arr[i];
        }
        arr = temp;
        temp = null;
        N += M;
        for (int i = 0; i < N; i++) {
            put(arr[i]);
        }
        return false;
    }

	private boolean array105() {
        int N = getInt();
        double[] arr = new double [N];
        for (int i = 0; i < N; i++) {
            arr[i] = getDouble();
        }
        int K = getInt();
        int M = getInt();
        double[] temp = new double [N+M];
        for (int i = 0, j = 0; i < N; i++) {
            temp[j++] = arr[i];
            if (i == K-1) {
                for (int k = 0; k < M; k++) {
                    temp[j++] = 0;
                }
            }
        }
        arr = temp;
        temp = null;
        N += M;
        for (int i = 0; i < N; i++) {
            put(arr[i]);
        }
        return false;
    }

	private boolean array106() {
        int N = getInt();
        double[] arr = new double [N];
        for (int i = 0; i < N; i++) {
            arr[i] = getDouble();
        }
        double[] temp = new double [N + N/2];
        for (int i = 0, j = 0; i < N; i += 2, j += 3) {
            temp[j] = arr[i];
        }
        for (int i = 1, j = 1; i < N; i += 2, j += 3) {
            temp[j] = temp[j+1] = arr[i];
        }
        arr = temp;
        temp = null;
        N += N/2;
        for (int i = 0; i < N; i++) {
            put(arr[i]);
        }
        return false;
    }

	private boolean array107() {
        int N = getInt();
        double[] arr = new double [N];
        for (int i = 0; i < N; i++) {
            arr[i] = getDouble();
        }
        int newSize = N + 2 * (N/2 + N%2);
        double[] temp = new double [newSize];
        for (int i = 1, j = 3; i < N; i += 2, j += 4) {
            temp[j] = arr[i];
        }
        for (int i = 0, j = 0; i < N; i += 2, j += 4) {
            temp[j] = temp[j+1] = temp[j+2] = arr[i];
        }
        arr = temp;
        temp = null;
        N = newSize;
        for (int i = 0; i < N; i++) {
            put(arr[i]);
        }
        return false;
    }

	private boolean array108() {
        int N = getInt();
        double[] arr = new double [N];
        int newSize = N;
        for (int i = 0; i < N; i++) {
            arr[i] = getDouble();
            if (arr[i] > 0) {
                newSize++;
            }
        }
        if (newSize != N) {
            double[] temp = new double [newSize];
            for (int i = 0, j = 0; i < N; i++) {
                if (arr[i] > 0) {
                    temp[j++] = 0;
                }
                temp[j++] = arr[i];
            }
            arr = temp;
            temp = null;
            N = newSize;
        }
        for (int i = 0; i < N; i++) {
            put(arr[i]);
        }
        return false;
    }

	private boolean array109() {
        int N = getInt();
        double[] arr = new double [N];
        int newSize = N;
        for (int i = 0; i < N; i++) {
            arr[i] = getDouble();
            if (arr[i] < 0) {
                newSize++;
            }
        }
        if (newSize != N) {
            double[] temp = new double [newSize];
            for (int i = 0, j = 0; i < N; i++) {
                temp[j++] = arr[i];
                if (arr[i] < 0) {
                    temp[j++] = 0;
                }
            }
            arr = temp;
            temp = null;
            N = newSize;                
        }
        for (int i = 0; i < N; i++) {
            put(arr[i]);
        }
        return false;
    }

	private boolean array110() {
        int N = getInt();
        int[] arr = new int [N];
        int newSize = N;
        for (int i = 0; i < N; i++) {
            arr[i] = getInt();
            if (arr[i] % 2 == 0) {
                newSize++;
            }
        }
        if (newSize != N) {
            int[] temp = new int [newSize];
            for (int i = 0, j = 0; i < N; i++) {
                temp[j++] = arr[i];
                if (arr[i] % 2 == 0) {
                    temp[j++] = arr[i];
                }
            }
            arr = temp;
            temp = null;
            N = newSize;
        }
        for (int i = 0; i < N; i++) {
            put(arr[i]);
        }
        return false;
    }

	private boolean array111() {
        int N = getInt();
        int[] arr = new int [N];
        int newSize = N;
        for (int i = 0; i < N; i++) {
            arr[i] = getInt();
            if (arr[i] % 2 != 0) {
                newSize += 2;
            }
        }
        if (newSize != N) {
            int[] temp = new int [newSize];
            for (int i = 0, j = 0; i < N; i++) {
                temp[j++] = arr[i];
                if (arr[i] % 2 != 0) {
                    for (int k = 1; k <= 2; k++) {
                        temp[j++] = arr[i];
                    }
                }
            }
            arr = temp;
            temp = null;
            N = newSize;
        }
        for (int i = 0; i < N; i++) {
            put(arr[i]);
        }
        return false;
    }

	private boolean array112() {
        int N = getInt();
        double[] A = new double [N];
        for (int i = 0; i < N; i++) {
            A[i] = getDouble();
        }
        double temp;
        for (int i = 0; i < N-1; i++) {
            for (int j = 1; j < N-i; j++) {
                if (A[j] < A[j-1]) {
                    temp = A[j];
                    A[j] = A[j-1];
                    A[j-1] = temp;
                }
            }
            for (int j = 0; j < N; j++) {
                put(A[j]);
            }
        }
        return false;
    }

	private boolean array113() {
        int N = getInt();
        double[] A = new double [N];
        for (int i = 0; i < N; i++) {
            A[i] = getDouble();
        }
        double temp;
        int maximalIndex;
        for (int i = N-1; i > 0; i--) {
            maximalIndex = i;
            for (int j = i-1; j >= 0; j--) {
                if (A[j] > A[maximalIndex]) {
                    maximalIndex = j;
                }
            }
            if (maximalIndex != i) {
                temp = A[i];
                A[i] = A[maximalIndex];
                A[maximalIndex] = temp;
            }
            for (int j = 0; j < N; j++) {
                put(A[j]);
            }
        }
        return false;
    }

	private boolean array114() {
        int N = getInt();
        double[] A  = new double [N];
        for (int i = 0; i < N; i++) {
            A[i] = getDouble();
        }
        double temp;
        for (int i = 1; i < N; i++) {
            for (int j = i; j > 0; j--) {
                if (A[j] > A[j-1]) {
                    break;
                }
                temp = A[j];
                A[j] = A[j-1];
                A[j-1] = temp;
            }
            for (int j = 0; j < N; j++) {
                put(A[j]);
            }
        }
        return false;
    }

	private boolean array115() {
        int N = getInt();
        double[] A = new double [N];
        int[] I = new int [N];
        for (int i = 0; i < N; i++) {
            A[i] = getDouble();
            I[i] = i;
        }
        int temp;
        for (int i = 0; i < N-1; i++) {
            for (int j = 1; j < N-i; j++) {
                if (A[I[j]] < A[I[j-1]]) {
                    temp = I[j];
                    I[j] = I[j-1];
                    I[j-1] = temp;
                }
            }
        }
        for (int i = 0; i < N; i++) {
            put(I[i]+1);
        }
        return false;
    }

	private boolean array116() {
        int N = getInt();
        int[] A = new int [N];
        int series = 0;
        for (int i = 0; i < N; i++) {
            A[i] = getInt();
            if (series == 0) {
                series++;
            } else if (A[i] != A[i-1]) {
                series++;
            }
        }
        int[] B = new int [series], C = new int [series];
        int count = 1, index = 0;
        for (int i = 1; i < N; i++) {
            if (A[i] != A[i-1]) {
                B[index] = count;
                C[index] = A[i-1];
                index++;
                count = 1;
            } else {
                count++;
            }
        }
        B[index] = count;
        C[index] = A[N-1];
        for (int i = 0; i < series; i++) {
            put(B[i]);
        }
        for (int i = 0; i < series; i++) {
            put(C[i]);
        }
        return false;
    }

	private boolean array117() {
        int N = getInt();
        int[] arr = new int [N];
        int series = 0;
        for (int i = 0; i < N; i++) {
            arr[i] = getInt();
            if (series == 0) {
                series++;
            } else if (arr[i] != arr[i-1]) {
                series++;
            }
        }
        int[] temp = new int [N + series];
        int index = 0;
        temp[index++] = 0;
        temp[index++] = arr[0];
        for (int i = 1; i < N; i++) {
            if (arr[i] != arr[i-1]) {
                temp[index++] = 0;
            }
            temp[index++] = arr[i];
        }
        arr = temp;
        temp = null;
        N += series;
        for (int i = 0; i < N; i++) {
            put(arr[i]);
        }
        return false;
    }

	private boolean array118() {
        int N = getInt();
        int[] arr = new int [N];
        int series = 0;
        for (int i = 0; i < N; i++) {
            arr[i] = getInt();
            if (series == 0) {
                series++;
            } else if (arr[i] != arr[i-1]) {
                series++;
            }
        }
        int[] temp = new int [N + series];
        int index = 0;
        temp[index++] = arr[0];
        for (int i = 1; i < N; i++) {
            if (arr[i] != arr[i-1]) {
                temp[index++] = 0;
            }
            temp[index++] = arr[i];
        }
        temp[index] = 0;
        arr = temp;
        temp = null;
        N += series;
        for (int i = 0; i < N; i++) {
            put(arr[i]);
        }
        return false;
    }

	private boolean array119() {
        int N = getInt();
        int[] arr = new int [N];
        int series = 0;
        for (int i = 0; i < N; i++) {
            arr[i] = getInt();
            if (series == 0) {
                series++;
            } else if (arr[i] != arr[i-1]) {
                series++;
            }
        }
        int[] temp = new int [N + series];
        int index = 0;
        temp[index++] = arr[0];
        for (int i = 1; i < N; i++) {
            if (arr[i] != arr[i-1]) {
                temp[index++] = arr[i-1];
            }
            temp[index++] = arr[i];
        }
        temp[index] = arr[N-1];
        arr = temp;
        temp = null;
        N += series;
        for (int i = 0; i < N; i++) {
            put(arr[i]);
        }
        return false;
    }

	private boolean array120() {
        int N = getInt();
        int[] arr = new int [N];
        int series = 0;
        for (int i = 0; i < N; i++) {
            arr[i] = getInt();
            if (series == 0) {
                series++;
            } else if (arr[i] != arr[i-1]) {
                series++;
            }
        }
        int[] temp = new int [N - series];
        int index = 0;
        for (int i = 1; i < N; i++) {
            if (arr[i] == arr[i-1]) {
                temp[index++] = arr[i];
            }
        }
        arr = temp;
        temp = null;
        N -= series;
        for (int i = 0; i < N; i++) {
            put(arr[i]);
        }
        return false;
    }

	private boolean array121() {
        int K = getInt();
        int N = getInt();
        int[] arr = new int [N];
        int series = 0, count = 0;
        int newSize = N;
        for (int i = 0; i < N; i++) {
            arr[i] = getInt();
            if (series == 0) {
                series++;
                count = 1;
            } else if (arr[i] != arr[i-1]) {
                if (series == K) {
                    newSize += count;
                }
                series++;
                count = 1;
            } else {
                count++;
            }
        }
        if (series == K) {
            newSize += count;
        }
        if (newSize != N) {
            int[] temp = new int [newSize];
            int index = 0;
            series = 0;
            for (int i = 0; i < N; i++) {
                temp[index++] = arr[i];
                if (series == 0) {
                    series++;
                } else if (arr[i] != arr[i-1]) {
                    series++;
                }
                if (series == K) {
                    temp[index++] = arr[i];
                }
            }
            arr = temp;
            temp = null;
            N = newSize;
        }
        for (int i = 0; i < N; i++) {
            put(arr[i]);
        }
        return false;
    }

	private boolean array122() {
        int K = getInt();
        int N = getInt();
        int[] arr = new int [N];
        int series = 0;
        int count = 0;
        int newSize = N;
        for (int i = 0; i < N; i++) {
            arr[i] = getInt();
            if (series == 0) {
                series++;
                count = 1;
            } else if (arr[i] != arr[i-1]) {
                if (series == K) {
                    newSize -= count;
                }
                series++;
                count = 1;
            } else {
                count++;
            }
        }
        if (series == K) {
            newSize -= count;
        }
        if (newSize != N) {
            int[] temp = new int [newSize];
            int index = 0;
            series = 0;
            for (int i = 0; i < N; i++) {
                if (series == 0) {
                    series++;
                } else if (arr[i] != arr[i-1]) {
                    series++;
                }
                if (series == K) {
                    continue;
                }
                temp[index++] = arr[i];
            }
            arr = temp;
            temp = null;
            N = newSize;
        }
        for (int i = 0; i < N; i++) {
            put(arr[i]);
        }
        return false;
    }

	private boolean array123() {
        int K = getInt();
        int N = getInt();
        int[] arr = new int [N];
        int series = 0;
        int start1 = -1, finish1 = -1, startK = -1, finishK = -1;
        for (int i = 0; i < N; i++) {
            arr[i] = getInt();
            if (series == 0) {
                series++;
                start1 = i;
            } else if (arr[i] != arr[i-1]) {
                if ( (finish1 < 0) && (series == 1) ) {
                    finish1 = i-1;
                }
                if ( (finishK < 0) && (series == K) ) {
                    finishK = i-1;
                }
                series++;
                if ( (startK < 0) && (series == K) ) {
                    startK = i;
                }
            }
        }
        if ( (finish1 < 0) && (series == 1) ) {
            finish1 = N-1;
        }
        if ( (finishK < 0) && (series == K) ) {
            finishK = N-1;
        }
        if ( (startK >= 0) && (finishK >= 0) ) {
            int[] temp = new int [N];
            int index = 0;
            for (int i = startK; i <= finishK; i++) {
                temp[index++] = arr[i];
            }
            for (int i = finish1+1; i < startK; i++) {
                temp[index++] = arr[i];
            }
            for (int i = start1; i <= finish1; i++) {
                temp[index++] = arr[i];
            }
            for (int i = finishK+1; i < N; i++) {
                temp[index++] = arr[i];
            }
            arr = temp;
            temp = null;
        }
        for (int i = 0; i < N; i++) {
            put(arr[i]);
        }
        return false;
    }

	private boolean array124() {
        int K = getInt();
        int N = getInt();
        int[] arr = new int [N];
        int series = 0;
        int startN = -1, finishN = -1, startK = -1, finishK = -1;
        for (int i = 0; i < N; i++) {
            arr[i] = getInt();
            if (series == 0) {
                series++;
                if ( (startK < 0) && (series == K) ) {
                    startK = i;
                }
            } else if (arr[i] != arr[i-1]) {
                if ( (finishK < 0) && (series == K) ) {
                    finishK = i-1;
                }
                series++;
                if ( (startK < 0) && (series == K) ) {
                    startK = i;
                }
                startN = i;
            }
        }
        if ( (finishK < 0) && (series == K) ) {
            finishK = N-1;
        }
        finishN = N-1;
        if ( (startK >= 0) && (finishK >= 0) && (startK != startN) && (finishK != finishN) ) {
            int[] temp = new int [N];
            int index = 0;
            for (int i = 0; i < startK; i++) {
                temp[index++] = arr[i];
            }
            for (int i = startN; i <= finishN; i++) {
                temp[index++] = arr[i];
            }
            for (int i = finishK + 1; i < startN; i++) {
                temp[index++] = arr[i];
            }
            for (int i = startK; i <= finishK; i++) {
                temp[index++] = arr[i];
            }
            arr = temp;
            temp = null;
        }
        show(startK+1);
        show(finishK+1);
        show(startN+1);
        show(finishN+1);
        for (int i = 0; i < N; i++) {
            put(arr[i]);
        }
        return false;
    }

	private boolean array125() {
        int L = getInt();
        int N = getInt();
        int[] arr = new int [N];
        for (int i = 0; i < N; i++) {
            arr[i] = getInt();
        }
        int count = 1, series = 1;
        for (int i = 1; i < N; i++) {
            if (arr[i] != arr[i-1]) {
                series++;
                if (count < L) {
                    arr[i-count] = 0;
                    for (int ba = i-count+1, az = i; az < N; az++, ba++) {
                        arr[ba] = arr[az];
                    }
                    N -= count-1;
                    i -= count-1;
                }
                count = 1;
            } else {
                count++;
            }
        }
        if (count < L) {
            arr[N-count] = 0;
            N -= count-1;
        }
        for (int i = 0; i < N; i++) {
            put(arr[i]);
        }
        return false;
    }

	private boolean array126() {
        int L = getInt();
        int N = getInt();
        int[] arr = new int [N];
        for (int i = 0; i < N; i++) {
            arr[i] = getInt();
        }
        int series = 1, count = 1;
        for (int i = 1; i < N; i++) {
            if (arr[i] != arr[i-1]) {
                series++;
                if (count == L) {
                    arr[i-count] = 0;
                    for (int ba = i-count+1, az = i; az < N; ba++, az++) {
                        arr[ba] = arr[az];
                    }
                    N -= count-1;
                    i -= count-1;
                }
                count = 1;
            } else {
                count++;
            }
        }
        if (count == L) {
            arr[N-count] = 0;
            N -= count-1;
        }
        for (int i = 0; i < N; i++) {
            put(arr[i]);
        }
        return false;
    }

	private boolean array127() {
        int L = getInt();
        int N = getInt();
        int[] arr = new int [N];
        for (int i = 0; i < N; i++) {
            arr[i] = getInt();
        }
        int series = 1, count = 1;
        for (int i = 1; i < N; i++) {
            if (arr[i] != arr[i-1]) {
                series++;
                if (count > L) {
                    arr[i-count] = 0;
                    for (int ba = i-count+1, az = i; az < N; ba++, az++) {
                        arr[ba] = arr[az];
                    }
                    N -= count-1;
                    i -= count-1;
                }
                count = 1;
            } else {
                count++;
            }
        }
        if (count > L) {
            arr[N-count] = 0;
            N -= count-1;
        }
        for (int i = 0; i < N; i++) {
            put(arr[i]);
        }
        return false;
    }

	private boolean array128() {
        int N = getInt();
        int[] arr = new int [N];
        int series = 0, count = 0, maxCount = 0, index = -1;
        for (int i = 0; i < N; i++) {
            arr[i] = getInt();
            if (series == 0) {
                series++;
                count = 1;
            } else if (arr[i] != arr[i-1]) {
                series++;
                if (count > maxCount) {
                    maxCount = count;
                    index = i-1;
                }
                count = 1;
            } else {
                count++;
            }
        }
        if (count > maxCount) {
            index = N-1;
        }
        int[] temp = new int [N+1];
        for (int i = 0, j = 0; i < N; i++) {
            temp[j++] = arr[i];
            if (index == i) {
                temp[j++] = arr[i];
            }
        }
        arr = temp;
        temp = null;
        N++;
        for (int i = 0; i < N; i++) {
            put(arr[i]);
        }
        return false;
    }

	private boolean array129() {
        int N = getInt();
        int[] arr = new int [N];
        int series = 0, count = 0, maxCount = 0, index = -1;
        for (int i = 0; i < N; i++) {
            arr[i] = getInt();
            if (series == 0) {
                series++;
                count = 1;
            } else if (arr[i] != arr[i-1]) {
                series++;
                if (count >= maxCount) {
                    maxCount = count;
                    index = i-1;
                }
                count = 1;
            } else {
                count++;
            }
        }
        if (count >= maxCount) {
            index = N-1;
        }
        int[] temp = new int [N+1];
        for (int i = 0, j = 0; i < N; i++) {
            temp[j++] = arr[i];
            if (i == index) {
                temp[j++] = arr[i];
            }
        }
        arr = temp;
        temp = null;
        N++;
        for (int i = 0; i < N; i++) {
            put(arr[i]);
        }
        return false;
    }

	private boolean array130() {
        int N = getInt();
        int[] arr = new int [N];
        int series = 0, count = 0, maxCount = 0;
        for (int i = 0; i < N; i++) {
            arr[i] = getInt();
            if (series == 0) {
                series++;
                count = 1;
            } else if (arr[i] != arr[i-1]) {
                series++;
                if (count > maxCount) {
                    maxCount = count;
                }
                count = 1;
            } else {
                count++;
            }
        }
        if (count > maxCount) {
            maxCount = count;
        }
        count = 1;
        for (int i = 1; i < N; i++) {
            if (arr[i] != arr[i-1]) {
                if (count == maxCount) {
                    int[] temp = new int [N+1];
                    for (int a = 0, t = 0; a < N; a++) {
                        if (a == i) {
                            temp[t++] = arr[a-1];
                        }
                        temp[t++] = arr[a];
                    }
                    arr = temp;
                    temp = null;
                    N++;
                    i++;
                }
                count = 1;
            } else {
                count++;
            }
        }
        if (count == maxCount) {
            int[] temp = new int [N+1];
            for (int a = 0, t = 0; a < N; a++) {
                if (a == N-1) {
                    temp[t++] = arr[a];
                }
                temp[t++] = arr[a];
            }
            arr = temp;
            temp = null;
            N++;
        }
        for (int i = 0; i < N; i++) {
            put(arr[i]);
        }
        return false;
    }

	private boolean array131() {
        int N = getInt();
        Point[] A = new Point [N];
        for (int i = 0; i < N; i++) {
            A[i] = new Point();
            A[i].x = getDouble();
            A[i].y = getDouble();
        }
        Point B = new Point();
        B.x = getDouble();
        B.y = getDouble();
        int i = 0, index = 0;
        double R, minR = Math.sqrt(Math.pow(A[i].x - B.x, 2) + Math.pow(A[i].y - B.y, 2));
        for (i++; i < N; i++) {
            R = Math.sqrt(Math.pow(A[i].x - B.x, 2) + Math.pow(A[i].y - B.y, 2));
            if (R < minR) {
                minR = R;
                index = i;
            }
        }
        put(A[index].x);
        put(A[index].y);
        return false;
    }

	private boolean array132() {
        int N = getInt();
        Point[] A = new Point[N];
        Point nuqta = new Point();
        nuqta.x = nuqta.y = 0;
        double R = 0, maxR = 0;
        for (int i = 0; i < N; i++) {
            A[i] = new Point();
            A[i].x = getDouble();
            A[i].y = getDouble();
            if ( (A[i].x < 0) && (A[i].y > 0) ) {
                if (nuqta.x == 0.0) {
                    nuqta.x = A[i].x;
                    nuqta.y = A[i].y;
                    maxR = Math.sqrt(Math.pow(A[i].x, 2) + Math.pow(A[i].y, 2));
                } else {
                    R = Math.sqrt(Math.pow(A[i].x, 2) + Math.pow(A[i].y, 2));
                    if (R > maxR) {
                        nuqta.x = A[i].x;
                        nuqta.y = A[i].y;
                        maxR = R;
                    }
                }
            }
        }
        put(nuqta.x);
        put(nuqta.y);
        return false;
    }

	private boolean array133() {
        int N = getInt();
        Point[] A = new Point[N];
        Point nuqta = new Point();
        nuqta.x = nuqta.y = 0;
        double R = 0, minR = 0;
        for (int i = 0; i < N; i++) {
            A[i] = new Point();
            A[i].x = getDouble();
            A[i].y = getDouble();
            if (A[i].x * A[i].y > 0) {
                if (nuqta.x == 0.0) {
                    nuqta.x = A[i].x;
                    nuqta.y = A[i].y;
                    minR = Math.sqrt(A[i].x * A[i].x + A[i].y * A[i].y);
                } else {
                    R = Math.sqrt(A[i].x * A[i].x + A[i].y * A[i].y);
                    if (R < minR) {
                        minR = R;
                        nuqta.x = A[i].x;
                        nuqta.y = A[i].y;
                    }
                }
            }
        }
        put(nuqta.x);
        put(nuqta.y);
        return false;
    }

	private boolean array134() {
        int N = getInt();
        Point[] A = new Point[N];
        double R = 0, maxR = 0;
        int one = -1, two = -1;
        for (int i = 0; i < N; i++) {
            A[i] = new Point();
            A[i].x = getDouble();
            A[i].y = getDouble();
            for (int j = 0; j < i; j++) {
                R = Math.sqrt( Math.pow(A[i].x - A[j].x, 2) + Math.pow(A[i].y - A[j].y, 2) );
                if (one < 0) {
                    maxR = R;
                    one = j;
                    two = i;
                } else if (R > maxR) {
                    maxR = R;
                    one = j;
                    two = i;
                }
            }
        }
        put(A[one].x);
        put(A[one].y);
        put(A[two].x);
        put(A[two].y);
        put(maxR);
        return false;
    }

	private boolean array135() {
        int N1 = getInt();
        Point[] A = new Point [N1];
        for (int i = 0; i < N1; i++) {
            A[i] = new Point();
            A[i].x = getDouble();
            A[i].y = getDouble();
        }
        int N2 = getInt();
        Point[] B = new Point [N2];
        double R = 0, minR = 0;
        int one = -1, two = -1;
        for (int i = 0; i < N2; i++) {
            B[i] = new Point();
            B[i].x = getDouble();
            B[i].y = getDouble();
            for (int a = 0; a < N1; a++) {
                R = Math.sqrt( Math.pow(A[a].x - B[i].x, 2) + Math.pow(A[a].y - B[i].y, 2) );
                if (one < 0) {
                    minR = R;
                    one = a;
                    two = i;
                } else if (R < minR) {
                    minR = R;
                    one = a;
                    two = i;
                }
            }
        }
        put(minR);
        put(A[one].x);
        put(A[one].y);
        put(B[two].x);
        put(B[two].y);
        return false;
    }

	private boolean array136() {
        int N = getInt();
        Point[] A = new Point [N];
        double[] sum = new double [N];
        for (int i = 0; i < N; i++) {
            sum[i] = 0;
        }
        double R;
        for (int i = 0; i < N; i++) {
            A[i] = new Point();
            A[i].x = getDouble();
            A[i].y = getDouble();
            for (int j = 0; j < i; j++) {
                R = Math.sqrt( Math.pow(A[i].x - A[j].x, 2) + Math.pow(A[i].y - A[j].y, 2) );
                sum[i] += R;
                sum[j] += R;
            }
        }
        int minIndex = 0;
        for (int i = 1; i < N; i++) {
            if (sum[i] < sum[minIndex]) {
                minIndex = i;
            }
        }
        put(A[minIndex].x);
        put(A[minIndex].y);
        put(sum[minIndex]);
        return false;
    }

	private boolean array137() {
        int N = getInt();
        Point[] A = new Point [N];
        for (int i = 0; i < N; i++) {
            A[i] = new Point();
            A[i].x = getDouble();
            A[i].y = getDouble();
        }
        double P = 0, maxP = -1;
        int index1 = 0, index2 = 0, index3 = 0;
        for (int a = 0; a < N; a++) {
            for (int b = a+1; b < N; b++) {
                for (int c = b+1; c < N; c++) {
                    P = GetPerim(A[a], A[b], A[c]);
                    if (P > maxP) {
                        maxP = P;
                        index1 = a;
                        index2 = b;
                        index3 = c;
                    }
                }
            }
        }
        put(maxP);
        put(A[index1].x);
        put(A[index1].y);
        put(A[index2].x);
        put(A[index2].y);
        put(A[index3].x);
        put(A[index3].y);
		return false;
    }
    
    public static double GetPerim(Point a, Point b, Point c) {
        double one = Math.sqrt( Math.pow(a.x - b.x, 2) + Math.pow(a.y - b.y, 2) );
        double two = Math.sqrt( Math.pow(a.x - c.x, 2) + Math.pow(a.y - c.y, 2) );
        double three = Math.sqrt( Math.pow(b.x - c.x, 2) + Math.pow(b.y - c.y, 2) );
        double P = one + two + three;
        return P;
    }

	private boolean array138() {
		int N = getInt();
		Point[] A = new Point [N];
		for (int i = 0; i < N; i++) {
			A[i] = new Point();
			A[i].x = getDouble();
			A[i].y = getDouble();
		}
		double P = 0, minP = 0;
		boolean firstTime = true;
		int index1 = 0, index2 = 0, index3 = 0;
		for (int a = 0; a < N; a++) {
			for (int b = a+1; b < N; b++) {
				for (int c = b+1; c < N; c++) {
					P = getPerim(A[a], A[b], A[c]);
					if (firstTime) {
						minP = P;
						index1 = a;
						index2 = b;
						index3 = c;
						firstTime = false;
					} else if (P < minP) {
						minP = P;
						index1 = a;
						index2 = b;
						index3 = c;
					}
				}
			}
		}
		put(minP);
		put(A[index1].x);
		put(A[index1].y);
		put(A[index2].x);
		put(A[index2].y);
		put(A[index3].x);
		put(A[index3].y);
		return false;
	}
	
	public static double getPerim(Point a, Point b, Point c) {
		double one = Math.sqrt( Math.pow(a.x - b.x, 2) + Math.pow(a.y - b.y, 2) );
		double two = Math.sqrt( Math.pow(a.x - c.x, 2) + Math.pow(a.y - c.y, 2) );
		double three = Math.sqrt( Math.pow(b.x - c.x, 2) + Math.pow(b.y - c.y, 2) );
		double P = one + two + three;
		return P;
	}

	private boolean array139() {
        int N = getInt();
        Point139[] A = new Point139 [N];
        for (int i = 0; i < N; i++) {
            A[i] = new Point139();
            A[i].x = getInt();
            A[i].y = getInt();
            if (i > 0) {
                for (int a = i; a > 0; a--) {
                    if (!sorted(A[a-1], A[a])) {
                        swap(A[a-1], A[a]);
                    }
                }
            }
        }
        for (int i = 0; i < N; i++) {
            put(A[i].x);
            put(A[i].y);
        }
        return false;
    }
    
    public static boolean sorted(Point139 a, Point139 b) {
        boolean areSorted = (a.x < b.x) && (a.y < b.y) || (a.x < b.x) || (a.x == b.x) && (a.y < b.y);
        return areSorted;
    }
    
    static void swap(Point139 a, Point139 b) {
        Point139 temp = new Point139();
        temp.x = a.x;
        temp.y = a.y;
        a.x = b.x;
        a.y = b.y;
        b.x = temp.x;
        b.y = temp.y;
    }

	private boolean array140() {
        int N = getInt();
        Point139[] A = new Point139 [N];
        for (int i = 0; i < N; i++) {
            A[i] = new Point139();
            A[i].x = getInt();
            A[i].y = getInt();
            for (int a = i; a > 0; a--) {
                if (sorted(A[a-1], A[a])) {
                    swap(A[a-1], A[a]);
                }
            }
        }
        for (int i = 0; i < N; i++) {
            put(A[i].x);
            put(A[i].y);
        }
        return false;
    }

    public static boolean sorted(Point139 a, Point139 b) {
        boolean areSorted = (a.x < b.x) && (a.y < b.y) ||
                (a.x + a.y < b.x + b.y) ||
                (a.x + a.y == b.x + b.y) && (a.x < b.x);
        return areSorted;
    }

    static void swap(Point139 a, Point139 b) {
        Point139 temp = new Point139();
        temp.x = a.x;
        temp.y = a.y;
        a.x = b.x;
        a.y = b.y;
        b.x = temp.x;
        b.y = temp.y;
    }
}

class Point {
    public double x;
    public double y;
}

class Point139 {
    public int x;
    public int y;
}

class Triangle {
    public Point a;
    public Point b;
    public Point c;
}