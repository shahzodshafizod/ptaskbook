import java.util.Scanner;

class WhileTask {

	private boolean while1() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("A = ");
		double A = scanner.nextDouble();
		System.out.print("B = ");
		double B = scanner.nextDouble();
		double freeSpace = A;
		while (freeSpace >= B) {
			freeSpace -= B;
		}
		System.out.printf("free space:\t%1$.2f\n", freeSpace);
		return false;
	}

	private boolean while2() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("A = ");
		double A = scanner.nextDouble();
		System.out.print("B = ");
		double B = scanner.nextDouble();
		int count = 0;
		while (A >= B) {
			A -= B;
			count++;
		}
		System.out.println("count = " + count);
	}

	private boolean while3() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("N = ");
		int N = scanner.nextInt();
		System.out.print("K = ");
		int K = scanner.nextInt();
		int div = 0;
		while (N >= K) {
			N -= K;
			div++;
		}
		int mod = N;
		System.out.println("div = " + div);
		System.out.println("mod = " + mod);
	}

	private boolean while4() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("N = ");
		int N = scanner.nextInt();
		int degree3 = 1;
		while (degree3 < N) {
			degree3 *= 3;
		}
		boolean isDegree3 = degree3 == N;
		System.out.println("is a degree of 3:\t" + isDegree3);
	}

	private boolean while5() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("N = ");
		int N = scanner.nextInt();
		int K = 0;
		int degree2 = 1;
		while (degree2 < N) {
			degree2 *= 2;
			K++;
		}
		System.out.println("K = " + K);
	}

	private boolean while6() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("N = ");
		int N = scanner.nextInt();
		double fact2 = 1;
		int temp = N;
		while (temp > 0) {
			fact2 *= temp;
			temp -= 2;
		}
		System.out.printf("%1$d! = %2$.0f\n", N, fact2);
	}

	private boolean while7() {
        int N = getInt();
        int K = 1;
        while (K*K <= N) {
            K++;
        }
        put(K);
        return false;
    }

	private boolean while8() {
        int N = getInt();
        int K = 1;
        while (K*K <= N) {
            K++;
        }
        K--;
        put(K);
        return false;
    }

	private boolean while9() {
        int N = getInt();
        int K = 0;
        int degree3 = 1;
        while (degree3 <= N) {
            K++;
            degree3 *= 3;
        }
        put(K);
        return false;
    }

	private boolean while10() {
        int N = getInt();
        int K = 0;
        int degree3 = 1;
        while (degree3 < N) {
            K++;
            degree3 *= 3;
        }
        degree3 /= 3;
        K--;
        put(K);
        return false;
    }

	private boolean while11() {
        int N = getInt();
        int sum = 0;
        int K = 0;
        while (sum < N) {
            K++;
            sum += K;
        }
        put(K);
        put(sum);
        return false;
    }

	private boolean while12() {
        int N = getInt();
        int sum = 0;
        int K = 0;
        while (sum <= N) {
            K++;
            sum += K;
        }
        sum -= K;
        K--;
        put(K);
        put(sum);
        return false;
    }

	private boolean while13() {
        double A = getDouble();
        double sum = 0;
        int K = 0;
        while (sum <= A) {
            K++;
            sum += 1.0 / K;
        }
        put(K);
        put(sum);
        return false;
    }

	private boolean while14() {
        double A = getDouble();
        double sum = 0;
        int K = 0;
        while (sum < A) {
            K++;
            sum += 1.0 / K;
        }
        sum -= 1.0 / K;
        K--;
        put(K);
        put(sum);
        return false;
    }

	private boolean while15() {
        double deposit = 1000;
        double P = getDouble();
        int K = 0;
        while (deposit <= 1100) {
            deposit += deposit * P / 100;
            K++;
        }
        put(K);
        put(deposit);
        return false;
    }

	private boolean while16() {
        double runWay = 10;
        double P = getDouble();
        int K = 1;
        double sum = runWay;
        while (sum <= 200) {
            runWay += runWay * P / 100;
            sum += runWay;
            K++;
        }
        put(K);
        put(sum);
        return false;
    }

	private boolean while17() {
        int N = getInt();
        int digit;
        while (N > 0) {
            digit = N % 10;
            put(digit);
            N /= 10;
        }
        return false;
    }

	private boolean while18() {
        int N = getInt();
        int count = 0, sum = 0;
        while (N > 0) {
            sum += N % 10;
            count++;
            N /= 10;
        }
        put(count);
        put(sum);
        return false;
    }

	private boolean while19() {
        int N = getInt();
        int chappa = 0;
        while (N > 0) {
            chappa = chappa * 10 + N % 10;
            N /= 10;
        }
        put(chappa);
        return false;
    }

	private boolean while20() {
        int N = getInt();
        boolean has2 = false;
        while (N > 0) {
            if (N % 10 == 2) {
                has2 = true;
                break;
            }
            N /= 10;
        }
        put(has2);
        return false;
    }

	private boolean while21() {
        int N = getInt();
        boolean hasToq = false;
        while (N > 0) {
            if (N % 10 % 2 != 0) {
                hasToq = true;
                break;
            }
            N /= 10;
        }
        put(hasToq);
        return false;
    }

	private boolean while22() {
        int N = getInt();
        boolean isPrime = true;
        int i = (int)Math.sqrt(N);
        while (i > 1) {
            if (N % i == 0) {
                isPrime = false;
                break;
            }
            i--;
        }
        put(isPrime);
        return false;
    }

	private boolean while23() {
        int A = getInt(), B = getInt();
        while (B != 0) {
            int mod = A % B;
            A = B;
            B = mod;
        }
        put(A);
        return false;
    }

	private boolean while24() {
        int N = getInt();
        int a = 1, b = 1;
        int c;
        do {
            c = a + b;
            if (c >= N) {
                break;
            }
            a = b;
            b = c;
        } while (true);
        boolean isFib = c == N;
        put(isFib);
        return false;
    }

	private boolean while25() {
        int N = getInt();
        int a = 1, b = 1;
        int c;
        do {
            c = a + b;
            if (c > N) {
                break;
            }
            a = b;
            b = c;
        } while (true);
        put(c);
        return false;
    }

	private boolean while26() {
        int N = getInt();
        int a = 1, b = 1;
        int c;
        do {
            c = a + b;
            if (b >= N) {
                break;
            }
            a = b;
            b = c;
        } while (true);
        put(a);
        put(c);
        return false;
    }

	private boolean while27() {
        int N = getInt();
        int a = 1, b = 1;
        int c, K = 2;
        do {
            K++;
            c = a + b;
            if (c >= N) {
                break;
            }
            a = b;
            b = c;
        } while (true);
        put(K);
        return false;
    }

	private boolean while28() {
        double e = getDouble();
        double A1 = 2;
        double Ak;
        int K = 1;
        do {
            K++;
            Ak = 2 + 1 / A1;
            if (Math.abs(Ak - A1) < e) {
                break;
            }
            A1 = Ak;
        } while (true);
        put(K);
        put(A1);
        put(Ak);
        return false;
    }

	private boolean while29() {
        double e = getDouble();
        double A1 = 1, A2 = 2;
        double Ak;
        int K = 2;
        do {
            K++;
            Ak = (A1 + 2*A2) / 3;
            if (Math.abs(Ak - A2) < e) {
                break;
            }
            A1 = A2;
            A2 = Ak;
        } while (true);
        put(K);
        put(A2);
        put(Ak);
        return false;
    }

	private boolean while30() {
        double A = getDouble();
        double B = getDouble();
        double C = getDouble();
        int CinA = 0;
        while (A >= C) {
            A -= C;
            CinA++;
        }
        int CinB = 0;
        while (B >= C) {
            B -= C;
            CinB++;
        }
        int kvads = 0;
        int i = 1;
        while (i <= CinA) {
            int j = 1;
            while (j <= CinB) {
                kvads++;
                j++;
            }
            i++;
        }
        put(kvads);
        return false;
    }
}