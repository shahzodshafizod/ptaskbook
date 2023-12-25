import java.util.Scanner;

class IfTask {

	private boolean if1() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("number = ");
		int number = scanner.nextInt();
		if (number > 0) {
			number -= 8;
		}
		System.out.println("number = " + number);
		return false;
	}

	private boolean if2() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("number = ");
		int number = scanner.nextInt();
		if (number > 0) {
			number -= 8;
		} else {
			number += 6;
		}
		System.out.println("number = " + number);
		return false;
	}

	private boolean if3() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("number = ");
		int number = scanner.nextInt();
		if (number > 0) {
			number -= 8;
		} else if (number < 0) {
			number -= 6;
		} else {
			number = 10;
		}
		System.out.println("number = " + number);
		return false;
	}

	private boolean if4() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("a = ");
		int a = scanner.nextInt();
		System.out.print("b = ");
		int b = scanner.nextInt();
		System.out.print("c = ");
		int c = scanner.nextInt();
		int positiveCount = 0;
		if (a > 0) {
			positiveCount++;
		}
		if (b > 0) {
			positiveCount++;
		}
		if (c > 0) {
			positiveCount++;
		}
		System.out.println("positiveCount = " + positiveCount);
		return false;
	}

	private boolean if5() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("a = ");
		int a = scanner.nextInt();
		System.out.print("b = ");
		int b = scanner.nextInt();
		System.out.print("c = ");
		int c = scanner.nextInt();
		int positiveCount = 0;
		int negativeCount = 0;
		if (a > 0) {
			positiveCount++;
		} else if (a < 0) {
			negativeCount++;
		}
		if (b > 0) {
			positiveCount++;
		} else if (b < 0) {
			negativeCount++;
		}
		if (c > 0) {
			positiveCount++;
		} else if (c < 0) {
			negativeCount++;
		}
		System.out.println("positiveCount = " + positiveCount);
		System.out.println("negativeCount = " + negativeCount);
		return false;
	}

	private boolean if6() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("number1 = ");
		double number1 = scanner.nextDouble();
		System.out.print("number2 = ");
		double number2 = scanner.nextDouble();
		if (number1 > number2) {
			System.out.println(number1);
		} else {
			System.out.println(number2);
		}
		return false;
	}

	private boolean if7() {
        double a = getDouble();
        double b = getDouble();
        if (a < b) {
            put(1);
        } else {
            put(2);
        }
        return false;
    }
	
	private boolean if8() {
        double A = getDouble();
        double B = getDouble();
        if (A > B) {
            put(A);
            put(B);
        } else {
            put(B);
            put(A);
        }
        return false;
    }

	private boolean if9() {
        double A = getDouble();
        double B = getDouble();
        if (A > B) {
            double temp = A;
            A = B;
            B = temp;
        }
        put(A);
        put(B);
        return false;
    }

	private boolean if10() {
        int A = getInt();
        int B = getInt();
        if (A != B) {
            A += B;
            B = A;
        } else {
            A = B = 0;
        }
        put(A);
        put(B);
        return false;
    }

	private boolean if11() {
		int A = getInt();
		int B = getInt();
		if (A != B) {
			if (A > B) {
				B = A;
			} else {
				A = B;
			}
		} else {
			A = B = 0;
		}
		put(A);
		put(B);
		return false;
	}

	private boolean if12() {
        double a = getDouble();
        double b = getDouble();
        double c = getDouble();
        double xurd = 0;
        if ( (a < b) && (a < c) ) {
            xurd = a;
        } else if (b < c) {
            xurd = b;
        } else {
            xurd = c;
        }
        put(xurd);
        return false;
    }

	private boolean if13() {
        double a = getDouble();
        double b = getDouble();
        double c = getDouble();
        double kalon, xurd;
        if ( (a > b) && (a > c) ) {
            kalon = a;
        } else if (b > c) {
            kalon = b;
        } else {
            kalon = c;
        }
        if ( (a < b) && (a < c) ) {
            xurd = a;
        } else if (b < c) {
            xurd = b;
        } else {
            xurd = c;
        }
        double middle = a+b+c - kalon - xurd;
        put(middle);
        return false;
    }

	private boolean if14() {
        double a = getDouble();
        double b = getDouble();
        double c = getDouble();
        double kalon, xurd;
        if ( (a > b) && (a > c) ) {
            kalon = a;
        } else if (b > c) {
            kalon = b;
        } else {
            kalon = c;
        }
        if ( (a < b) && (a < c) ) {
            xurd = a;
        } else if (b < c) {
            xurd = b;
        } else {
            xurd = c;
        }
        put(xurd);
        put(kalon);
        return false;
    }

	private boolean if15() {
        double a = getDouble();
        double b = getDouble();
        double c = getDouble();
        double xurd = 0;
        if ( (a < b) && (a < c) ) {
            xurd = a;
        } else if (b < c) {
            xurd = b;
        } else {
            xurd = c;
        }
        double sum = a + b + c - xurd;
        put(sum);
        return false;
    }

	private boolean if16() {
        double A = getDouble();
        double B = getDouble();
        double C = getDouble();
        if ( (A < B) && (B < C) ) {
            A *= 2;
            B *= 2;
            C *= 2;
        } else {
            A *= -1;
            B *= -1;
            C *= -1;
        }
        put(A);
        put(B);
        put(C);
        return false;
    }

	private boolean if17() {
        double A = getDouble();
        double B = getDouble();
        double C = getDouble();
        if ( (A < B) && (B < C) || (A > B) && (B > C) ) {
            A *= 2;
            B *= 2;
            C *= 2;
        } else {
            A *= -1;
            B *= -1;
            C *= -1;
        }
        put(A);
        put(B);
        put(C);
        return false;
    }

	private boolean if18() {
        int A = getInt();
        int B = getInt();
        int C = getInt();
        if (A == B) {
            put(3);
        } else if (A == C) {
            put(2);
        } else {
            put(1);
        }
        return false;
    }

	private boolean if19() {
        int A = getInt();
        int B = getInt();
        int C = getInt();
        int D = getInt();
        if (A == B) {
            if (C != A) {
                put(3);
            } else {
                put(4);
            }
        } else {
            if (A == C) {
                put(2);
            } else {
                put(1);
            }
        }
        return false;
    }

	private boolean if20() {
        double A = getDouble();
        double B = getDouble();
        double C = getDouble();
        double AB = Math.abs(B - A);
        double AC = Math.abs(C - A);
        if (AB < AC) {
            put(B);
            put(AB);
        } else {
            put(C);
            put(AC);
        }
        return false;
    }

	private boolean if21() {
        int x = getInt();
        int y = getInt();
        if ( (x == 0) && (y == 0) ) {
            put(0);
        } else if (y == 0) {
            put(1);
        } else if (x == 0) {
            put(2);
        } else {
            put(3);
        }
        return false;
    }

	private boolean if22() {
        double x = getDouble();
        double y = getDouble();
        int quarterNo = 0;
        if (x > 0) {
            if (y > 0) {
                quarterNo = 1;
            } else {
                quarterNo = 4;
            }
        } else {
            if (y > 0) {
                quarterNo = 2;
            } else {
                quarterNo = 3;
            }
        }
        put(quarterNo);
        return false;
    }

	private boolean if23() {
        int x1 = getInt();
        int y1 = getInt();
        int x2 = getInt();
        int y2 = getInt();
        int x3 = getInt();
        int y3 = getInt();
        int x4 = 0, y4 = 0;
        if (x1 == x2) {
            x4 = x3;
        } else if (x1 == x3) {
            x4 = x2;
        } else {
            x4 = x1;
        }
        if (y1 == y2) {
            y4 = y3;
        } else if (y1 == y3) {
            y4 = y2;
        } else {
            y4 = y1;
        }
        put(x4);
        put(y4);
        return false;
    }

	private boolean if24() {
        double x = getDouble();
        double y = 0;
        if (x > 0) {
            y = 2 * Math.sin(x);
        } else {
            y = 6 - x;
        }
        put(y);
        return false;
    }

	private boolean if25() {
        int x = getInt();
        int y = 0;
        if ( (x < -2) || (x > 2) ) {
            y = 2*x;
        } else {
            y = -3*x;
        }
        put(y);
        return false;
    }

	private boolean if26() {
        double x = getDouble();
        double y = 0;
        if (x <= 0) {
            y = -x;
        } else if (x >= 2) {
            y = 4;
        } else {
            y = x*x;
        }
        put(y);
        return false;
    }

	private boolean if27() {
        double x = getDouble();
        int y = 0;
        if (x < 0) {
            y = 0;
        } else if ((int)x % 2 == 0) {
            y = 1;
        } else {
            y = -1;
        }
        put(y);
        return false;
    }

	private boolean if28() {
        int year = getInt();
        int days = 0;
        if (year % 400 == 0) {
            days = 366;
        } else if (year % 100 == 0) {
            days = 365;
        } else if (year % 4 == 0) {
            days = 366;
        } else {
            days = 365;
        }
        put(days);
        return false;
    }

	private boolean if29() {
        int number = getInt();
        String descr = "";
        if (number == 0) {
            descr += "нулевое ";
        } else {
            if (number > 0) {
                descr += "положительное ";
            } else {
                descr += "отрицательное ";
            }
            if (number % 2 != 0) {
                descr += "не";
            }
            descr += "четное ";
        }
        descr += "число";
        put(descr);
        return false;
    }

	private boolean if30() {
        int number = getInt();
        String descr = "";
        if (number % 2 != 0) {
            descr += "не";
        }
        descr += "четное ";
        if (number <= 9) {
            descr += "одно";
        } else if (number <= 99) {
            descr += "дву";
        } else {
            descr += "трех";
        }
        descr += "значное число";
        put(descr);
        return false;
    }
}