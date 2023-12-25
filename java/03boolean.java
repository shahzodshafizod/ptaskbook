import java.util.Scanner;

class BooleanTask {

	private boolean boolean1() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("A = ");
		int A = scanner.nextInt();
		boolean isPositive = A > 0;
		System.out.println("isPositive = " + isPositive);
		return false;
	}

	private boolean boolean2() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("A = ");
		int A = scanner.nextInt();
		boolean isOdd = A % 2 != 0;
		System.out.println("isOdd = " + isOdd);
		return false;
	}

	private boolean boolean3() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("A = ");
		int A = scanner.nextInt();
		boolean isEven = (A & 1) == 0;
		System.out.println("is even:\t" + isEven);
		return false;
	}

	private boolean boolean4() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("A = ");
		int A = scanner.nextInt();
		System.out.print("B = ");
		int B = scanner.nextInt();
		boolean nobarobari = A > 2 && B <= 3;
		System.out.println("A > 2 AND B <= 3:\t" + nobarobari);
		return false;
	}

	private boolean boolean5() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("A = ");
		int A = scanner.nextInt();
		System.out.print("B = ");
		int B = scanner.nextInt();
		boolean nobarobari = A >= 0 || B < -2;
		System.out.println("A >= 0 OR B < -2:\t" + nobarobari);
		return false;
	}

	private boolean boolean6() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("A = ");
		int A = scanner.nextInt();
		System.out.print("B = ");
		int B = scanner.nextInt();
		System.out.print("C = ");
		int C = scanner.nextInt();
		boolean nobarobari = A < B && B < C;
		System.out.println("A < B < C:\t" + nobarobari);
		return false;
	}

	private boolean boolean7() {
        int A = getInt();
        int B = getInt();
        int C = getInt();
        boolean isBetween = (A < B) && (B < C) || (A > B) && (B > C);
        put(isBetween);
        return false;
    }

	private boolean boolean8() {
        int A = getInt();
        int B = getInt();
        boolean areOdd = (A % 2 != 0) && (B % 2 != 0);
        put(areOdd);
        return false;
    }

	private boolean boolean9() {
        int A = getInt();
        int B = getInt();
        boolean oneOdd = (A % 2 != 0) || (B % 2 != 0);
        put(oneOdd);
        return false;
    }

	private boolean boolean10() {
        int A = getInt();
        int B = getInt();
        boolean oneOdd = (A + B) % 2 != 0;
        put(oneOdd);
        return false;
    }

	private boolean boolean11() {
        int A = getInt();
        int B = getInt();
        boolean sameEvenity = (A + B) % 2 == 0;
        put(sameEvenity);
        return false;
    }

	private boolean boolean12() {
        int A = getInt();
        int B = getInt();
        int C = getInt();
        boolean arePositive = (A > 0) && (B > 0) && (C > 0);
        put(arePositive);
        return false;
    }

	private boolean boolean13() {
        int A = getInt();
        int B = getInt();
        int C = getInt();
        boolean onePositive = (A > 0) || (B > 0) || (C > 0);
        put(onePositive);
        return false;
    }

	private boolean boolean14() {
        int A = getInt();
        int B = getInt();
        int C = getInt();
        boolean onePositive = (A > 0) && (B <= 0) && (C <= 0) ||
                                (B > 0) && (A <= 0) && (C <= 0) ||
                                (C > 0) && (A <= 0) && (B <= 0);
        put(onePositive);
        return false;
    }

	private boolean boolean15() {
        int A = getInt();
        int B = getInt();
        int C = getInt();
        boolean twoPositives = (A > 0) && (B > 0) && (C <= 0) ||
                                (A > 0) && (C > 0) && (B <= 0) ||
                                (B > 0) && (C > 0) && (A <= 0);
        put(twoPositives);
        return false;
    }

	private boolean boolean16() {
		int number = getInt();
		boolean evenDouble = (number % 2 == 0) && (number > 9) && (number < 100);
		put(evenDouble);
		return false;
	}

	private boolean boolean17() {
        int number = getInt();
        boolean oddThrible = (number % 2 != 0) && (number > 99) && (number < 1000);
        put(oddThrible);
        return false;
    }

	private boolean boolean18() {
        int a = getInt();
        int b = getInt();
        int c = getInt();
        boolean hasSame = (a == b) || (a == c) || (b == c);
        put(hasSame);
        return false;
    }

	private boolean boolean19() {
        int a = getInt();
        int b = getInt();
        int c = getInt();
        boolean hasBaraks = (a == -b) || (a == -c) || (b == -c);
        put(hasBaraks);
        return false;
    }

	private boolean boolean20() {
        int number = getInt();
        int sadi = number / 100;
        int dahi = number / 10 % 10;
        int vohid = number % 10;
        boolean areDifferent = (sadi != dahi) && (dahi != vohid) && (sadi != vohid);
        put(areDifferent);
        return false;
    }

	private boolean boolean21() {
        int number = getInt();
        int sadi = number / 100;
        int dahi = number / 10 % 10;
        int vohid = number % 10;
        boolean areProgress = (sadi < dahi) && (dahi < vohid);
        put(areProgress);
        return false;
    }

	private boolean boolean22() {
        int number = getInt();
        int sadi = number / 100;
        int dahi = number / 10 % 10;
        int vohid = number % 10;
        boolean isProgReg = (sadi < dahi) && (dahi < vohid) ||
                            (sadi > dahi) && (dahi > vohid);
        put(isProgReg);
        return false;
    }

	private boolean boolean23() {
        int number = getInt();
        int hazori = number / 1000;
        int sadi = number / 100 % 10;
        boolean isPalindrome = sadi * 10 + hazori == number % 100;
        put(isPalindrome);
        return false;
    }

	private boolean boolean24() {
        double A = getDouble();
        double B = getDouble();
        double C = getDouble();
        double D = B*B - 4*A*C;
        boolean hasRoot = D >= 0;
        put(hasRoot);
        return false;
    }

	private boolean boolean25() {
        double x = getDouble();
        double y = getDouble();
        boolean two = (x < 0) && (y > 0);
        put(two);
        return false;
    }

	private boolean boolean26() {
        double x = getDouble();
        double y = getDouble();
        boolean four = (x > 0) && (y < 0);
        put(four);
        return false;
    }

	private boolean boolean27() {
        double x = getDouble();
        double y = getDouble();
        boolean twoOrThree = x < 0;
        put(twoOrThree);
        return false;
    }

	private boolean boolean28() {
        double x = getDouble();
        double y = getDouble();
        boolean oneOrTwo = x * y > 0;
        put(oneOrTwo);
        return false;
    }

	private boolean boolean29() {
        double x = getDouble();
        double y = getDouble();
        double x1 = getDouble();
        double y1 = getDouble();
        double x2 = getDouble();
        double y2 = getDouble();
        boolean isIn = (x > x1) && (x < x2) && (y > y2) && (y < y1);
        put(isIn);
        return false;
    }

	private boolean boolean30() {
		int a = getInt();
		int b = getInt();
		int c = getInt();
		boolean isBarobarTaraf = (a == b) && (b == c);
		put(isBarobarTaraf);
		return false;
	}

	private boolean boolean31() {
        int a = getInt();
        int b = getInt();
        int c = getInt();
        boolean isBarobarPahlu = (a == b) || (a == c) || (b == c);
        put(isBarobarPahlu);
        return false;
    }

	private boolean boolean32() {
        int a = getInt();
        int b = getInt();
        int c = getInt();
        boolean isRectTriangle = (a*a + b*b == c*c) || (a*a + c*c == b*b) ||
                                    (b*b + c*c == a*a);
        put(isRectTriangle);
        return false;
    }

	private boolean boolean33() {
        int a = getInt();
        int b = getInt();
        int c = getInt();
        boolean isTriangle = (a + b > c) && (a + c > b) && (b + c > a);
        put(isTriangle);
        return false;
    }

	private boolean boolean34() {
		int x = getInt();
		int y = getInt();
		boolean isBlack = (x + y) % 2 != 0;
		put(isBlack);
		return false;
	}

	private boolean boolean35() {
        int x1 = getInt();
        int y1 = getInt();
        int x2 = getInt();
        int y2 = getInt();
        boolean sameColor = (x1 + y1) % 2 == (y2 + x2) % 2;
        put(sameColor);
        return false;
    }

	private boolean boolean36() {
        int x1 = getInt();
        int y1 = getInt();
        int x2 = getInt();
        int y2 = getInt();
        boolean rux = (x1 == x2) || (y1 == y2);
        put(rux);
        return false;
    }

	private boolean boolean37() {
        int x1 = getInt();
        int y1 = getInt();
        int x2 = getInt();
        int y2 = getInt();
        boolean shoh = (Math.abs(x2 - x1) < 2) && (Math.abs(y2 - y1) < 2);
        put(shoh);
        return false;
    }

	private boolean boolean38() {
        int x1 = getInt();
        int y1 = getInt();
        int x2 = getInt();
        int y2 = getInt();
        boolean fil = Math.abs(x2 - x1) == Math.abs(y2 - y1);
        put(fil);
        return false;
    }

	private boolean boolean39() {
        int x1 = getInt();
        int y1 = getInt();
        int x2 = getInt();
        int y2 = getInt();
        boolean farzin = (x1 == x2) || (y1 == y2) || (Math.abs(x2 - x1) == Math.abs(y2 - y1));
        put(farzin);
        return false;
    }

	private boolean boolean40() {
        int x1 = getInt();
        int y1 = getInt();
        int x2 = getInt();
        int y2 = getInt();
        boolean asp = (Math.abs(x2 - x1) == 1) && (Math.abs(y2 - y1) == 2)
                    || (Math.abs(x2 - x1) == 2) && (Math.abs(y2 - y1) == 1);
        put(asp);
        return false;
    }
}