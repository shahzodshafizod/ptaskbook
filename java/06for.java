import java.util.Scanner;

class ForTask {

	private boolean for1() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("K = ");
		int K = scanner.nextInt();
		System.out.print("N = ");
		int N = scanner.nextInt();
		for (int i = 1; i <= N; i++) {
			System.out.printf("%d\t", K);
		}
		System.out.println();
		return false;
	}

	private boolean for2() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("A = ");
		int A = scanner.nextInt();
		System.out.print("B = ");
		int B = scanner.nextInt();
		int N = 0;
		for (int i = A; i <= B; i++) {
			System.out.print(i + "\t");
			N++;
		}
		System.out.println("\nN = " + N);
		return false;
	}

	private boolean for3() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("A = ");
		int A = scanner.nextInt();
		System.out.print("B = ");
		int B = scanner.nextInt();
		int N = 0;
		for (int i = B - 1; i > A; i--) {
			System.out.print(i + "\t");
			N++;
		}
		System.out.println("\nN = " + N);
		return false;
	}

	private boolean for4() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("one kg price:\t");
		double oneKgPrice = scanner.nextDouble();
		for (int i = 1; i <= 10; i++) {
			System.out.printf("%1$.2f\t", i * oneKgPrice);
		}
		return false;
	}

	private boolean for5() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("one kg price:\t");
		double oneKgPrice = scanner.nextDouble();
		for (double i = 0.1; i <= 1; i += 0.1) {
			System.out.printf("%1$.2f kg costs %2$.2f somoni\n", i, i * oneKgPrice);
		}
		return false;
	}

	private boolean for6() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("one kg price:\t");
		double oneKgPrice = scanner.nextDouble();
		for (double i = 1.2; i <= 2; i += 0.2) {
			System.out.printf("%1$.2f kg costs $%2$.2f somoni\n", i, i * oneKgPrice);
		}
		return false;
	}

	private boolean for7() {
        int A = getInt(), B = getInt();
        int sum = 0;
        for (int i = A; i <= B; i++) {
            sum += i;
        }
        put(sum);
        return false;
    }

	private boolean for8() {
        int A = getInt(), B = getInt();
        int mult = 1;
        for (int i = A; i <= B; i++) {
            mult *= i;
        }
        put(mult);
        return false;
    }

	private boolean for9() {
        int A = getInt(), B = getInt();
        int sum = 0;
        for (int i = A; i <= B; i++) {
            sum += i*i;
        }
        put(sum);
        return false;
    }

	private boolean for10() {
        int N = getInt();
        double sum = 0;
        for (int i = 1; i <= N; i++) {
            sum += 1 / (double)i;
        }
        put(sum);
        return false;
    }

	private boolean for11() {
        int N = getInt();
        int sum = 0;
        for (int i = 2*N; i >= N; i--) {
            sum += i*i;
        }
        put(sum);
        return false;
    }

	private boolean for12() {
        int N = getInt();
        double number = 1.1, mult = 1;
        for (int i = 1; i <= N; i++, number += 0.1) {
            mult *= number;
        }
        put(mult);
        return false;
    }

	private boolean for13() {
        int N = getInt();
        double number = 1.1, sum = 0;
        for (int factor = 1, i = 1; i <= N; i++, factor *= -1, number += 0.1) {
            sum += factor * number;
        }
        put(sum);
        return false;
    }

	private boolean for14() {
        int N = getInt();
        int sum = 0;
        for (int i = 1; i <= N; i++) {
            sum += 2 * i - 1;
            put(sum);
        }
        return false;
    }

	private boolean for15() {
        double A = getDouble();
        int N = getInt();
        double degreeA = 1;
        for (int i = 1; i <= N; i++) {
            degreeA *= A;
        }
        put(degreeA);
        return false;
    }

	private boolean for16() {
        double A = getDouble();
        int N = getInt();
        double degreeA = 1;
        for (int i = 1; i <= N; i++) {
            degreeA *= A;
            put(degreeA);
        }
        return false;
    }

	private boolean for17() {
        double A = getDouble();
        int N = getInt();
        double degreeA = 1;
        double sum = 0;
        for (int i = 0; i <= N; i++) {
            sum += degreeA;
            degreeA *= A;
        }
        put(sum);
        return false;
    }

	private boolean for18() {
        double A = getDouble();
        int N = getInt();
        double degreeA = 1, sum = 0;
        int factor = 1;
        for (int i = 0; i <= N; i++) {
            sum += factor * degreeA;
            degreeA *= A;
            factor *= -1;
        }
        put(sum);
        return false;
    }

	private boolean for19() {
        int N = getInt();
        double fact = 1;
        for (int i = 1; i <= N; i++) {
            fact *= i;
        }
        put(fact);
        return false;
    }

	private boolean for20() {
        int N = getInt();
        double fact = 1, sum = 0;
        for (int i = 1; i <= N; i++) {
            fact *= i;
            sum += fact;
        }
        put(sum);
        return false;
    }

	private boolean for21() {
        int N = getInt();
        double sum = 0, fact = 1;
        for (int i = 0; i <= N; ) {
            sum += 1 / fact;
            i++;
            fact *= i;
        }
        put(sum);
        return false;
    }

	private boolean for22() {
        double X = getDouble();
        int N = getInt();
        double sum = 0, fact = 1;
        double degreeX = 1;
        for (int i = 0; i <= N; ) {
            sum += degreeX / fact;
            degreeX *= X;
            i++;
            fact *= i;
        }
        put(sum);
        return false;
    }

	private boolean for23() {
        double X = getDouble();
        int N = getInt();
        double degreeX = 1, sum = 0;
        double fact = 1;
        int factor = 1;
        for (int i = 1; i <= N+1; i++) {
            degreeX *= X;
            sum += factor * degreeX / fact;
            fact *= (2*i) * (2*i+1);
            degreeX *= X;
            factor *= -1;
        }
        put(sum);
        return false;
    }

	private boolean for24() {
        double X = getDouble();
        int N = getInt();
        double degreeX = 1, sum = 1;
        double fact = 1;
        int factor = 1;
        for (int i = 1; i <= N; i++) {
            factor *= -1;
            degreeX *= X * X;
            fact *= (2 * i - 1) * (2 * i);
            sum += factor * degreeX / fact;
        }
        put(sum);
        return false;
    }

	private boolean for25() {
        double X = getDouble();
        int N = getInt();
        double degreeX = 1, sum = 0;
        int factor = 1;
        for (int i = 1; i <= N; i++) {
            degreeX *= X;
            sum += factor * degreeX / i;
            factor *= -1;
        }
        put(sum);
        return false;
    }

	private boolean for26() {
        double X = getDouble();
        int N = getInt();
        double degreeX = 1, sum = 0;
        int factor = 1, maxraj = 1;
        for (int i = 0; i <= N; i++) {
            degreeX *= X;
            sum += factor * degreeX / (2 * i + 1);
            degreeX *= X;
            factor *= -1;
        }
        put(sum);
        return false;
    }

	private boolean for27() {
        double X = getDouble();
        int N = getInt();
        double sum = X, degreeX = X;
        double toqs = 1, jufts = 1;
        for (int i = 1; i <= N; i++) {
            jufts *= 2 * i;
            toqs *= 2 * i - 1;
            degreeX *= X * X;
            sum += toqs * degreeX / (jufts * (2 * i + 1));
        }
        put(sum);
        return false;
    }

	private boolean for28() {
        double X = getDouble();
        int N = getInt();
        double degreeX = 1;
        double jufts = 1, toqs = 1;
        int factor = 1;
        double sum = 1;
        for (int i = 1; i <= N; ) {
            jufts *= 2 * i;
            degreeX *= X;
            sum += factor * toqs * degreeX / jufts;
            i++;
            toqs *= 2 * i - 3;
            factor *= -1;
        }
        put(sum);
        return false;
    }

	private boolean for29() {
        int N = getInt();
        double A = getDouble();
        double B = getDouble();
        double H = (B - A) / N;
        put(H);
        double X = A;
        for (int i = 0; i <= N; X += H, i++) {
            put(X);
        }
        return false;
    }

	private boolean for30() {
        int N = getInt();
        double A = getDouble();
        double B = getDouble();
        double H = (B - A) / N;
        double F;
        put(H);
        double X = A;
        for (int i = 0; i <= N; X += H, i++) {
            F = 1 - Math.sin(X);
            put(F);
        }
        return false;
    }

	private boolean for31() {
        int N = getInt();
        double A0 = 2;
        double Ak;
        for (int K = 1; K <= N; K++) {
            Ak = 2 + 1 / A0;
            put(Ak);
            A0 = Ak;
        }
        return false;
    }

	private boolean for32() {
        int N = getInt();
        double A0 = 1;
        double Ak;
        for (int K = 1; K <= N; K++) {
            Ak = (A0 + 1) / K;
            put(Ak);
            A0 = Ak;
        }
        return false;
    }

	private boolean for33() {
        int N = getInt();
        int a = 1, b = 1;
        put(a); put(b);
        int c;
        for (int K = 3; K <= N; K++) {
            c = a + b;
            put(c);
            a = b;
            b = c;
        }
        return false;
    }

	private boolean for34() {
        int N = getInt();
        double A1 = 1, A2 = 2;
        put(A1); put(A2);
        double Ak;
        for (int K = 3; K <= N; K++) {
            Ak = (A1 + 2 * A2) / 3;
            put(Ak);
            A1 = A2;
            A2 = Ak;
        }
        return false;
    }

	private boolean for35() {
        int N = getInt();
        int A1 = 1, A2 = 2, A3 = 3;
        put(A1); put(A2); put(A3);
        int Ak;
        for (int K = 4; K <= N; K++) {
            Ak = A3 + A2 - 2 * A1;
            put(Ak);
            A1 = A2;
            A2 = A3;
            A3 = Ak;
        }
        return false;
    }

	private boolean for36() {
        int N = getInt(), K = getInt();
        double degree, sum = 0;
        for (int i = 1; i <= N; i++) {
            degree = 1;
            for (int j = 1; j <= K; j++) {
                degree *= i;
            }
            sum += degree;
        }
        put(sum);
        return false;
    }

	private boolean for37() {
        int N = getInt();
        double degree, sum = 0;
        for (int i = 1; i <= N; i++) {
            degree = 1;
            for (int j = 1; j <= i; j++) {
                degree *= i;
            }
            sum += degree;
        }
        put(sum);
        return false;
    }

	private boolean for38() {
        int N = getInt();
        double degree, sum = 0;
        for (int i = 1, j = N; i <= N; i++, j--) {
            degree = 1;
            for (int k = 1; k <= j; k++) {
                degree *= i;
            }
            sum += degree;
        }
        put(sum);
        return false;
    }

	private boolean for39() {
        int A = getInt(), B = getInt();
        for (int i = A; i <= B; i++) {
            for (int j = 1; j <= i; j++) {
                put(i);
            }
        }
        return false;
    }

	private boolean for40() {
        int A = getInt(), B = getInt();
        for (int i = A, j = 1; i <= B; i++, j++) {
            for (int k = 1; k <= j; k++) {
                put(i);
            }
        }
        return false;
    }
}