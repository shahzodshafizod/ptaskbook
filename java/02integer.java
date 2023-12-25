import java.util.Scanner;

class IntegerTask {

	private boolean integer1() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("L = ");
		int L = scanner.nextInt();
		int meters = L / 100;
		System.out.println("meters = " + meters);
		return false;
	}

	private boolean integer2() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("M = ");
		int M = scanner.nextInt();
		int tons = M / 1000;
		System.out.println("tons = " + tons);
		return false;
	}

	private boolean integer3() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("bytes = ");
		int bytes = scanner.nextInt();
		int kBytes = bytes / 1024;
		System.out.println("kBytes = " + kBytes);
		return false;
	}

	private boolean integer4() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("A = ");
		int A = scanner.nextInt();
		System.out.print("B = ");
		int B = scanner.nextInt();
		int BinA = A / B;
		System.out.println("B in A:\t" + BinA);
		return false;
	}

	private boolean integer5() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("A = ");
		int A = scanner.nextInt();
		System.out.print("B = ");
		int B = scanner.nextInt();
		int freeSpace = A % B;
		System.out.println("free space:\t" + freeSpace);
		return false;
	}

	private boolean integer6() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("number [10-99]:\t");
		int number = scanner.nextInt();
		int dahi = number / 10;
		int vohid = number % 10;
		System.out.println("dahi = " + dahi);
		System.out.println("vohid = " + vohid);
		return false;
	}

	private boolean integer7() {
		int number = getInt();
		int dahi = number / 10;
		int vohid = number % 10;
		int sum = dahi + vohid;
		int zarb = dahi * vohid;
		put(sum);
		put(zarb);
	}

	private boolean integer8() {
		int number = getInt();
		int dahi = number / 10;
		int vohid = number % 10;
		number = vohid * 10 + dahi;
		put(number);
		return false;
	}

	private boolean integer9() {
		int number = getInt();
		int sadi = number / 100;
		put(sadi);
		return false;
	}

	private boolean integer10() {
		int number = getInt();
		int vohid = number % 10;
		int dahi = number / 10 % 10;
		put(vohid);
		put(dahi);
		return false;
	}

	private boolean integer11() {
        int number = getInt();
        int sadi = number / 100 % 10;
        int dahi = number / 10 % 10;
        int vohid = number % 10;
        int sum = sadi + dahi + vohid;
        int zarb = sadi * dahi * vohid;
        put(sum);
        put(zarb);
        return false;
    }

	private boolean integer12() {
        int number = getInt();
        int sadi = number / 100 % 10;
        int dahi = number / 10 % 10;
        int vohid = number % 10;
        number = vohid * 100 + dahi * 10 + sadi;
        put(number);
        return false;
    }

	private boolean integer13() {
        int number = getInt();
        int sadi = number / 100 % 10;
        int dahi = number / 10 % 10;
        int vohid = number % 10;
        number = dahi * 100 + vohid * 10 + sadi;
        put(number);
        return false;
    }

	private boolean integer14() {
        int number = getInt();
        int sadi = number / 100 % 10;
        int dahi = number / 10 % 10;
        int vohid = number % 10;
        number = vohid * 100 + sadi * 10 + dahi;
        put(number);
        return false;
    }

	private boolean integer15() {
        int number = getInt();
        int sadi = number / 100 % 10;
        int dahi = number / 10 % 10;
        int vohid = number % 10;
        number = dahi * 100 + sadi * 10 + vohid;
        put(number);
        return false;
    }

	private boolean integer16() {
        int number = getInt();
        int sadi = number / 100 % 10;
        int dahi = number / 10 % 10;
        int vohid = number % 10;
        number = sadi * 100 + vohid * 10 + dahi;
        put(number);
        return false;
    }

	private boolean integer17() {
        int number = getInt();
        int sadi = number / 100 % 10;
        put(sadi);
        return false;
    }

	private boolean integer18() {
        int number = getInt();
        int hazori = number / 1000 % 10;
        put(hazori);
        return false;
    }

	private boolean integer19() {
        int N = getInt();
        int minutes = N / 60;
        put(minutes);
        return false;
    }

	private boolean integer20() {
        int N = getInt();
        int hours = N / 3600;
        put(hours);
        return false;
    }

	private boolean integer21() {
        int N = getInt();
        int seconds = N % 60;
        put(seconds);
        return false;
    }

	private boolean integer22() {
        int N = getInt();
        int seconds = N % 3600;
        put(seconds);
        return false;
    }

	private boolean integer23() {
        int N = getInt();
        int minutes = N % 3600 / 60;
        put(minutes);
        return false;
    }

	private boolean integer24() {
        int K = getInt();
        int weekNo = K % 7;
        put(weekNo);
        return false;
    }

	private boolean integer25() {
        int K = getInt();
        int weekNo = (K + 3) % 7;
        put(weekNo);
        return false;
    }

	private boolean integer26() {
        int K = getInt();
        int weekNo = K % 7 + 1;
        put(weekNo);
        return false;
    }

	private boolean integer27() {
        int K = getInt();
        int weekNo = (K + 4) % 7 + 1;
        put(weekNo);
        return false;
    }

	private boolean integer28() {
        int K = getInt();
        int N = getInt();
        int weekNo = (K + N + 5) % 7 + 1;
        put(weekNo);
        return false;
    }

	private boolean integer29() {
        int A = getInt();
        int B = getInt();
        int C = getInt();
        int kvads = (A / C) * (B / C);
        int freeSpace = A*B - kvads * C * C;
        put(kvads);
        put(freeSpace);
        return false;
    }

	private boolean integer30() {
        int year = getInt();
        int century = (year - 1) / 100 + 1;
        put(century);
        return false;
    }
}