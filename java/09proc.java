import java.util.Scanner;
import static java.lang.Math.sqrt;
import static java.lang.Math.abs;

class TDouble {
	public double value;
}

class TInt {
	public int value;
}

class ProcTask {

	private boolean proc1() {
		Scanner scanner = new Scanner(System.in);
		double a;
		TDouble b = new TDouble();
		for (int i = 1; i < 6; i++) {
			a = scanner.nextDouble();
			PowerA3(a, b);
			System.out.printf("b = %1$.2f\n", b.value);
		}
		return false;
	}
	
	static void PowerA3(double A, TDouble B) {
		B.value = A * A * A;
	}

	private boolean proc2() {
		Scanner scanner = new Scanner(System.in);
		double a;
		TDouble b = new TDouble();
		TDouble c = new TDouble();
		TDouble d = new TDouble();
		for (int i = 1; i <= 5; i++) {
			a = scanner.nextDouble();
			PowerA234(a, b, c, d);
			System.out.printf("b = %1$.2f\n", b.value);
			System.out.printf("c = %1$.2f\n", c.value);
			System.out.printf("d = %1$.2f\n", d.value);
		}
		return false;
	}
	
	static void PowerA234(double A, TDouble B, TDouble C, TDouble D) {
		B.value = A*A;
		C.value = B.value * A;
		D.value = C.value * A;
	}

	private boolean proc3() {
		Scanner scanner = new Scanner(System.in);
		TDouble aMean = new TDouble();
		TDouble gMean = new TDouble();
		System.out.print("A = ");
		double a = scanner.nextDouble();
		System.out.print("B = ");
		double b = scanner.nextDouble();
		System.out.print("C = ");
		double c = scanner.nextDouble();
		System.out.print("D = ");
		double d = scanner.nextDouble();
		
		Mean(a, b, aMean, gMean);
		System.out.printf("Mean(%1$.2f, %2$.2f): AMean = %3$.2f, GMean = %4$.2f\n", a, b, aMean.value, gMean.value);
		
		Mean(a, c, aMean, gMean);
		System.out.printf("Mean(%1$.2f, %2$.2f): AMean = %3$.2f, GMean = %4$.2f\n", a, c, aMean.value, gMean.value);
		
		Mean(a, d, aMean, gMean);
		System.out.printf("Mean(%1$.2f, %2$.2f): AMean = %3$.2f, GMean = %4$.2f\n", a, d, aMean.value, gMean.value);
		return false;
	}
	
	static void Mean(double A, double B, TDouble AMean, TDouble GMean) {
		AMean.value = (A + B) / 2;
		GMean.value = Math.sqrt(A * B);
	}

	private boolean proc4() {
		Scanner scanner = new Scanner(System.in);
		double a;
		TDouble p = new TDouble();
		TDouble s = new TDouble();
		for (int i = 1; i <= 3; i++) {
			a = scanner.nextDouble();
			TrianglePS(a, p, s);
			System.out.printf("P = %1$.2f\t\tS = %2$.2f\n", p.value, s.value);
		}
		return false;
	}
	
	static void TrianglePS(double a, TDouble P, TDouble S) {
		P.value = 3 * a;
		S.value = a * a * sqrt(3) / 4;
	}

	private boolean proc5() {
		Scanner scanner = new Scanner(System.in);
		double x1;
		double y1;
		double x2;
		double y2;
		TDouble P = new TDouble();
		TDouble S = new TDouble();
		for (int i = 1; i <= 3; i++) {
			x1 = scanner.nextDouble();
			y1 = scanner.nextDouble();
			x2 = scanner.nextDouble();
			y2 = scanner.nextDouble();
			RectPS(x1, y1, x2, y2, P, S);
			System.out.printf("P = %1$.2f\nS = %2$.2f\n", P.value, S.value);
		}
		return false;
	}
	
	static void RectPS(double x1, double y1, double x2, double y2, TDouble P, TDouble S) {
		double a = abs(x2 - x1);
		double b = abs(y2 - y1);
		P.value = 2 * (a + b);
		S.value = a * b;
	}

	private boolean proc6() {
		Scanner scanner = new Scanner(System.in);
		int k;
		TInt c = new TInt();
		TInt s = new TInt();
		for (int i = 0; i < 5; i++) {
			k = scanner.nextInt();
			DigitCountSum(k, c, s);
			System.out.println("count = " + c.value);
			System.out.println("sum = " + s.value);
		}
		return false;
	}
	
	static void DigitCountSum(int k, TInt c, TInt s) {
		c.value = s.value = 0;
		while (k > 0) {
			s.value += k % 10;
			k /= 10;
			c.value++;
		}
	}

	private boolean proc7() {
        TInt K = new TInt();
        for (int i = 1; i < 6; i++) {
            K.value = getInt();
            InvDigits(K);
            put(K.value);
        }
		return false;
    }
    
    static void InvDigits(TInt K) {
        int temp = K.value;
        K.value = 0;
        while (temp > 0) {
            K.value = K.value * 10 + temp % 10;
            temp /= 10;
        }
    }

	private boolean proc8() {
        TInt K = new TInt();
        K.value = getInt();
        short D;
        for (int i = 1; i < 3; i++) {
            D = (short)getInt();
            AddRightDigit(D, K);
            put(K.value);
        }
        return false;
    }
    
    static void AddRightDigit(short D, TInt K) {
        K.value = K.value * 10 + D;
    }

	private boolean proc9() {
        TInt K = new TInt();
        K.value = getInt();
        int D;
        for (int i = 0; i < 2; i++) {
            D = getInt();
            AddLeftDigit(D, K);
            put(K.value);
        }
        return false;
    }
    
    static void AddLeftDigit(int D, TInt K) {
        int temp = K.value;
        while (temp > 0) {
            D *= 10;
            temp /= 10;
        }
        K.value += D;
    }

	private boolean proc10() {
        TDouble A = new TDouble(), B = new TDouble(), C = new TDouble(), D = new TDouble();
        A.value = getDouble();
        B.value = getDouble();
        C.value = getDouble();
        D.value = getDouble();
        Swap(A, B);
        Swap(C, D);
        Swap(B, C);
        put(A.value);
        put(B.value);
        put(C.value);
        put(D.value);
        return false;
    }
    
    static void Swap(TDouble X, TDouble Y) {
        double temp = X.value;
        X.value = Y.value;
        Y.value = temp;
    }

	private boolean proc11() {
        TDouble A = new TDouble(), B = new TDouble(), C = new TDouble(), D = new TDouble();
        A.value = getDouble();
        B.value = getDouble();
        C.value = getDouble();
        D.value = getDouble();
        Minmax(A, B);
        Minmax(C, D);
        Minmax(A, C);
        Minmax(B, D);
        put(A.value);
        put(D.value);
        return false;
    }
    
    static void Minmax(TDouble X, TDouble Y) {
        if (X.value > Y.value) {
            Swap(X, Y);
        }
    }
    
    static void Swap(TDouble X, TDouble Y) {
        double temp = X.value;
        X.value = Y.value;
        Y.value = temp;
    }

	private boolean proc12() {
        TDouble A = new TDouble(), B = new TDouble(), C = new TDouble();
        for (int i = 1; i <= 2; i++) {
            A.value = getDouble();
            B.value = getDouble();
            C.value = getDouble();
            SortInc3(A, B, C);
            put(A.value);
            put(B.value);
            put(C.value);
        }
		return false;
    }
    
    static void SortInc3(TDouble A, TDouble B, TDouble C) {
        Minmax(A, B);
        Minmax(A, C);
        Minmax(B, C);
    }
    
    static void Minmax(TDouble A, TDouble B) {
        if (A.value > B.value)
            Swap(A, B);
    }
    
    static void Swap(TDouble A, TDouble B) {
        double temp = A.value;
        A.value = B.value;
        B.value = temp;
    }

	private boolean proc13() {
        TDouble A = new TDouble(), B = new TDouble(), C = new TDouble();
        for (int i = 1; i <= 2; i++) {
            A.value = getDouble();
            B.value = getDouble();
            C.value = getDouble();
            SortDec3(A, B, C);
            put(A.value);
            put(B.value);
            put(C.value);
        }
        return false;
    }
    
    static void SortDec3(TDouble A, TDouble B, TDouble C) {
        Maxmin(A, B);
        Maxmin(A, C);
        Maxmin(B, C);
    }
    
    static void Maxmin(TDouble A, TDouble B) {
        if (A.value < B.value)
            Swap(A, B);
    }
    
    static void Swap(TDouble A, TDouble B) {
        double temp = A.value;
        A.value = B.value;
        B.value = temp;
    }

	private boolean proc14() {
        TDouble A = new TDouble(), B = new TDouble(), C = new TDouble();
        for (int i = 1; i <= 2; i++) {
            A.value = getDouble();
            B.value = getDouble();
            C.value = getDouble();
            ShiftRight3(A, B, C);
            put(A.value);
            put(B.value);
            put(C.value);
        }
        return false;
    }
    
    static void ShiftRight3(TDouble A, TDouble B, TDouble C) {
        TDouble temp = new TDouble();
        temp.value = A.value;
        A.value = C.value;
        C.value = B.value;
        B.value = temp.value;
    }

	private boolean proc15() {
		TDouble A = new TDouble(), B = new TDouble(), C = new TDouble();
		for (int i = 1; i <= 2; i++) {
			A.value = getDouble();
			B.value = getDouble();
			C.value = getDouble();
			ShiftLeft3(A, B, C);
			put(A.value);
			put(B.value);
			put(C.value);
		}
		return false;
	}
	
	static void ShiftLeft3(TDouble A, TDouble B, TDouble C) {
		TDouble D = new TDouble();
		D.value = A.value;
		A.value = B.value;
		B.value = C.value;
		C.value = D.value;
	}

	private boolean proc16() {
        double A = getDouble();
        double B = getDouble();
        int r = Sign(A) + Sign(B);
        put(r);
        return false;
    }

    static int Sign(double X) {
        if (X > 0)
            return 1;
        else if (X < 0)
            return -1;
        else
            return 0;
    }

	private boolean proc17() {
        double A, B, C;
        for (int i = 0; i < 3; i++) {
            A = getDouble();
            B = getDouble();
            C = getDouble();
            put(RootCount(A, B, C));
        }
        return false;
    }
    
    public static int RootCount(double a, double b, double c) {
        double d = b*b - 4 * a * c;
        if (d > 0)
            return 2;
        else if (d == 0)
            return 1;
        else
            return 0;
    }

	private boolean proc18() {
        double R;
        for (int i = 1; i < 4; i++) {
            R = getDouble();
            put(CircleS(R));
        }
        return false;
    }
    
    public static double CircleS(double R) {
        final double PI = 3.14;
        double S = PI * R * R;
        return S;
    }

	private boolean proc19() {
        double r1, r2, s;
        for (int i = 1; i < 4; i++) {
            r1 = getDouble();
            r2 = getDouble();
            s = RingS(r1, r2);
            put(s);
        }
        return false;
    }
    
    public static double RingS(double R1, double R2) {
        final double PI = 3.14;
        double S1 = PI * R1 * R1;
        double S2 = PI * R2 * R2;
        double S3 = S1 - S2;
        return S3;
    }

	private boolean proc20() {
		double a, h, p;
		for (int i = 1; i < 4; i++) {
			a = getDouble();
			h = getDouble();
			p = TriangleP(a, h);
			put(p);
		}
		return false;
	}
	
	public static double TriangleP(double a, double h) {
		double b = Math.sqrt(Math.pow(a/2, 2) + h*h);
		double c = b;
		double P = a + b + c;
		return P;
	}

    private boolean proc21() {
        int A = getInt();
        int B = getInt();
        int C = getInt();
        put(SumRange(A, B));
        put(SumRange(B, C));
        return false;
    }
    
    public static int SumRange(int A, int B) {
        int sum = 0;
        for (int i = A; i <= B; i++) {
            sum += i;
        }
        return sum;
    }

    private boolean proc22() {
        double A = getDouble(), B = getDouble();
        int N;
        for (int i = 0; i < 3; i++) {
            N = getInt();
            put(Calc(A, B, N));
        }
        return false;
    }
    
    public static double Calc(double A, double B, int Op) {
        double result;
        switch (Op) {
            case 1: result = A - B; break;
            case 2: result = A * B; break;
            case 3: result = B != 0 ? A / B : 0; break;
            default: result = A + B; break;
        }
        return result;
    }

    private boolean proc23() {
        double x, y;
        for (int i = 1; i < 4; i++) {
            x = getDouble();
            y = getDouble();
            put(Quarter(x, y));
        }
        return false;
    }
    
    public static int Quarter(double x, double y) {
        if (x > 0) {
            if (y > 0)
                return 1;
            else
                return 4;
        } else {
            if (y > 0)
                return 2;
            else
                return 3;
        }
    }

    private boolean proc24() {
        int number, evenCount = 0;
        for (int i = 0; i < 10; i++) {
            number = getInt();
            evenCount += Even(number) ? 1 : 0;
        }
        put(evenCount);
        return false;
    }
    
    public static boolean Even(int K) {
        return K % 2 == 0;
    }

    private boolean proc25() {
        int number, kvads = 0;
        for (int i = 0; i < 10; i++) {
            number = getInt();
            kvads += IsSquare(number) ? 1 : 0;
        }
        put(kvads);
        return false;
    }
    
    public static boolean IsSquare(int K) {
        int number = 1;
        while (number * number < K) {
            number++;
        }
        return number*number == K;
    }

    private boolean proc26() {
        int number, degrees = 0;
        for (int i = 1; i <= 10; i++) {
            number = getInt();
            degrees += IsPower5(number) ? 1 : 0;
        }
        put(degrees);
        return false;
    }
    
    public static boolean IsPower5(int K) {
        int degree5 = 1;
        while (degree5 < K) {
            degree5 *= 5;
        }
        return degree5 == K;
    }

    private boolean proc27() {
        int number, N = getInt();
        int degrees = 0;
        for (int i = 1; i <= 10; i++) {
            number = getInt();
            degrees += IsPowerN(number, N) ? 1 : 0;
        }
        put(degrees);
        return false;
    }
    
    public static boolean IsPowerN(int K, int N) {
        int degreeN = 1;
        while (degreeN < K) {
            degreeN *= N;
        }
        return degreeN == K;
    }

    private boolean proc28() {
        int number, primes = 0;
        for (int i = 1; i <= 10; i++) {
            number = getInt();
            primes += IsPrime(number) ? 1 : 0;
        }
        put(primes);
        return false;
    }
    
    public static boolean IsPrime(int K) {
        boolean isPrime = true;
        for (int i = (int)Math.sqrt(K); i > 1; i--) {
            if (K % i == 0) {
                isPrime = false;
                break;
            }
        }
        return isPrime;
    }

    private boolean proc29() {
        int K, digits;
        for (int i = 1; i < 6; i++) {
            K = getInt();
            digits = DigitCount(K);
            put(digits);
        }
        return false;
    }
    
    public static int DigitCount(int K) {
        int digits = 0;
        while (K > 0) {
            digits++;
            K /= 10;
        }
        return digits;
    }

    private boolean proc30() {
        int K, digit;
        for (int i = 1; i < 6; i++) {
            K = getInt();
            for (int j = 1; j < 6; j++) {
                digit = DigitN(K, j);
                put(digit);
            }
        }
        return false;
    }
    
    public static int DigitN(int K, int N) {
        int digit = -1;
        for (int i = 1; K > 0; i++) {
            if (i == N) {
                digit = K % 10;
                break;
            }
            K /= 10;
        }
        return digit;
    }

    private boolean proc31() {
        int number, palindromes = 0;
        for (int i = 1; i <= 10; i++) {
            number = getInt();
            palindromes += IsPalindrome(number) ? 1 : 0;
        }
        put(palindromes);
        return false;
    }
    
    public static boolean IsPalindrome(int K) {
        int chappa = 0;
        int temp = K;
        while (temp > 0) {
            chappa = chappa * 10 + temp % 10;
            temp /= 10;
        }
        return K == chappa;
    }

    private boolean proc32() {
        double D, R;
        for (int i = 1; i < 6; i++) {
            D = getDouble();
            R = DegToRad(D);
            put(R);
        }
        return false;
    }
    
    public static double DegToRad(double deg) {
        double rad = deg * 3.14 / 180;
        return rad;
    }

    private boolean proc33() {
        double R, D;
        for (int i = 1; i < 6; i++) {
            R = getDouble();
            D = RadToDeg(R);
            put(D);
        }
        return false;
    }
    
    public static double RadToDeg(double rad) {
        double deg = rad * 180 / 3.14;
        return deg;
    }

    private boolean proc34() {
        int N;
        double fact;
        for (int i = 1; i < 6; i++) {
            N = getInt();
            fact = Fact(N);
            put(fact);
        }
        return false;
    }
    
    public static double Fact(int N) {
        double fact = 1;
        while (N > 0) {
            fact *= N--;
        }
        return fact;
    }

    private boolean proc35() {
        int N;
        double fact;
        for (int i = 0; i < 5; i++) {
            N = getInt();
            fact = Fact2(N);
            put(fact);
        }
        return false;
    }
    
    public static double Fact2(int N) {
        double fact = 1;
        while (N > 0) {
            fact *= N;
            N -= 2;
        }
        return fact;
    }

    private boolean proc36() {
        int N, fib;
        for (int i = 1; i < 6; i++) {
            N = getInt();
            fib = Fib(N);
            put(fib);
        }
        return false;
    }
    
    public static int Fib(int N) {
        if (N <= 2) {
            return 1;
        }
        int a = 1;
        int b = 1;
        int c;
        int K = 3;
        do {
            c = a + b;
            a = b;
            b = c;
            K++;
        } while (K <= N);
        
        return c;
    }

    private boolean proc37() {
        double P = getDouble();
        double A;
        for (int i = 1; i < 4; i++) {
            A = getDouble();
            put( Power1(A, P) );
        }
        return false;
    }
    
    public static double Power1(double A, double B) {
        if (A <= 0) {
            return 0;
        }
        return Math.exp(B * Math.log(A));
    }

    private boolean proc38() {
        double A = getDouble();
        int N;
        for (int i = 1; i < 4; i++) {
            N = getInt();
            put( Power2(A, N) );
        }
        return false;
    }
    
    public static double Power2(double A, int N) {
        if (N == 0) {
            return 1;
        }
        double degreeA = 1;
        boolean isNegative = false;
        if (N < 0) {
            N *= -1;
            isNegative = true;
        }
        for (int i = 1; i <= N; i++) {
            degreeA *= A;
        }
        if (isNegative) {
            return 1 / degreeA;
        } else {
            return degreeA;
        }
    }

    private boolean proc39() {
        double P = getDouble();
        double A;
        for (int i = 1; i <= 3; i++) {
            A = getDouble();
            put( Power3(A, P) );
        }
        return false;
    }
    
    public static double Power1(double A, double B) {
        if (A <= 0) {
            return 0;
        }
        return Math.exp(B * Math.log(A));
    }
    
    public static double Power2(double A, int N) {
        if (N == 0) {
            return 1;
        }
        boolean isNegative = false;
        if (N < 0) {
            N *= -1;
            isNegative = true;
        }
        double degreeA = 1;
        for (int i = 1; i <= N; i++) {
            degreeA *= A;
        }
        if (isNegative) {
            return 1 / degreeA;
        }
        return degreeA;
    }
    
    public static double Power3(double A, double B) {
        int intPart = (int)B;
        if (B == intPart) {
            return Power2(A, intPart);
        } else {
            return Power1(A, B);
        }
    }

    private boolean proc40() {
        double x = getDouble();
        double e;
        for (int i = 1; i < 7; i++) {
            e = getDouble();
            put( Exp1(x, e) );
        }
        return false;
    }
    
    public static double Exp1(double x, double e) {
        double sum = 0;
        double fact = 1, degree = 1;
        int i = 1;
        while (degree / fact > e) {
            sum += degree / fact;
            degree *= x;
            fact *= i++;
        }
        return sum;
    }

    private boolean proc41() {
        double x = getDouble();
        double e;
        for (int i = 1; i < 7; i++) {
            e = getDouble();
            put( Sin1(x, e) );
        }
        return false;
    }
    
    public static double Sin1(double x, double e) {
        double sum = 0;
        int factor = 1;
        double degree = x, fact = 1;
        int i = 1;
        while (Math.abs(degree / fact) > e) {
            sum += factor * degree / fact;
            degree *= x * x;
            factor *= -1;
            i += 2;
            fact *= (i-1) * i;
        }
        return sum;
    }

    private boolean proc42() {
        double x = getDouble();
        double e;
        for (int i = 1; i < 7; i++) {
            e = getDouble();
            put( Cos1(x, e) );
        }
        return false;
    }
    
    public static double Cos1(double x, double e) {
        double sum = 0;
        double degree = 1, fact = 1;
        int factor = 1;
        int i = 0;
        while (Math.abs(degree / fact) > e) {
            sum += factor * degree / fact;
            degree *= x * x;
            factor *= -1;
            i += 2;
            fact *= (i-1) * i;
        }
        return sum;
    }

    private boolean proc43() {
        double x = getDouble();
        double e;
        for (int i = 1; i < 7; i++) {
            e = getDouble();
            put( Ln1(x, e) );
        }
        return false;
    }
    
    public static double Ln1(double x, double e) {
        double sum = 0;
        double degree = 1;
        int factor = 1;
        for (int i = 1; true; i++) {
            degree *= x;
            if (Math.abs(degree / i) > e) {
                sum += factor * degree / i;
            } else {
                break;
            }
            factor *= -1;
        }
        return sum;
    }

    private boolean proc44() {
        double x = getDouble();
        double e;
        for (int i = 1; i < 7; i++) {
            e = getDouble();
            put( Arctg1(x, e) );
        }
        return false;
    }
    
    public static double Arctg1(double x, double e) {
        double sum = 0;
        double degree = 1;
        int factor = 1;
        for (int i = 1; true; i += 2) {
            degree *= x;
            if (Math.abs(degree / i) > e) {
                sum += factor * degree / i;
            } else {
                break;
            }
            degree *= x;
            factor *= -1;
        }
        return sum;
    }

    private boolean proc45() {
        double x = getDouble(), a = getDouble();
        double e;
        for (int i = 1; i < 7; i++) {
            e = getDouble();
            put( Power4(x, a, e) );
        }
        return false;
    }
    
    public static double Power4(double x, double a, double e) {
        double sum = 0;
        double degree = 1;
        double factorA = 1;
        double fact = 1;
        int i = 0;
        while ( Math.abs(factorA * degree / fact) > e ) {
            sum += factorA * degree / fact;
            factorA *= (a - i);
            degree *= x;
            fact *= ++i;
        }
        return sum;
    }

    private boolean proc46() {
        int A = getInt();
        int B;
        for (int i = 1; i < 4; i++) {
            B = getInt();
            put( GCD2(A, B) );
        }
        return false;
    }
    
    public static int GCD2(int A, int B) {
        int temp;
        while (B != 0) {
            temp = A % B;
            A = B;
            B = temp;
        }
        return A;
    }

    private boolean proc47() {
        int a = getInt();
        int b = getInt();
        int c, d, maxraj, surat;
        TInt p = new TInt(), q = new TInt();
        for (int i = 1; i <= 3; i++) {
            c = getInt();
            d = getInt();
            maxraj = b * d;
            surat = d * a + b * c;
            Fraq1(surat, maxraj, p, q);
            put(p.value);
            put(q.value);
        }
        return false;
    }
    
    public static int GCD2(int A, int B) {
        int temp;
        while (B != 0) {
            temp = A % B;
            A = B;
            B = temp;
        }
        return A;
    }
    
    static void Fraq1(int a, int b, TInt p, TInt q) {
        int gcd = GCD2(a, b);
        p.value = a / gcd;
        q.value = b / gcd;
        if (q.value < 0) {
            q.value *= -1;
            p.value *= -1;
        }
    }

    private boolean proc48() {
        int A = getInt();
        int B;
        for (int i = 1; i < 4; i++) {
            B = getInt();
            put( LCM2(A, B) );
        }
        return false;
    }
    
    public static int LCM2(int A, int B) {
        return B / GCD2(A, B) * A;
    }
    
    public static int GCD2(int A, int B) {
        int temp;
        while (B != 0) {
            temp = A % B;
            A = B;
            B = temp;
        }
        return A;
    }

    private boolean proc49() {
        int A = getInt();
        int B = getInt();
        int C = getInt();
        int D = getInt();
        put( GCD3(A, B, C) );
        put( GCD3(A, C, D) );
        put( GCD3(B, C, D) );
        return false;
    }
    
    public static int GCD2(int A, int B) {
        int temp;
        while (B != 0) {
            temp = A % B;
            A = B;
            B = temp;
        }
        return A;
    }
    
    public static int GCD3(int A, int B, int C) {
        return GCD2(GCD2(A, B), C);
    }

    private boolean proc50() {
        int t;
        TInt h = new TInt(), m = new TInt(), s = new TInt();
        for (int i = 1; i <= 5; i++) {
            t = getInt();
            TimeToHMS(t, h, m, s);
            put(h.value);
            put(m.value);
            put(s.value);
        }
        return false;
    }
    
    static void TimeToHMS(int T, TInt H, TInt M, TInt S) {
        H.value = T / 3600;
        T %= 3600;
        M.value = T / 60;
        T %= 60;
        S.value = T;
    }

    private boolean proc51() {
        TInt h = new TInt();
        h.value = getInt();
        TInt m = new TInt();
        m.value = getInt();
        TInt s = new TInt();
        s.value = getInt();
        int t = getInt();
        IncTime(h, m, s, t);
        put(h.value);
        put(m.value);
        put(s.value);
        return false;
    }
    
    static void IncTime(TInt H, TInt M, TInt S, int T) {
        S.value += T;
        M.value += S.value / 60;
        S.value %= 60;
        H.value += M.value / 60;
        M.value %= 60;
    }

    private boolean proc52() {
        int year;
        for (int i = 1; i <= 5; i++) {
            year = getInt();
            put( IsLeapYear(year) );
        }
        return false;
    }
    
    public static boolean IsLeapYear(int Y) {
        if (Y % 400 == 0) {
            return true;
        } else if (Y % 100 == 0) {
            return false;
        } else if (Y % 4 == 0) {
            return true;
        } else {
            return false;
        }
    }

    private boolean proc53() {
        int year = getInt();
        int month;
        for (int i = 1; i <= 3; i++) {
            month = getInt();
            put( MonthDays(month, year) );
        }
        return false;
    }
    
    public static boolean IsLeapYear(int Y) {
        if (Y % 400 == 0) {
            return true;
        } else if (Y % 100 == 0) {
            return false;
        } else if (Y % 4 == 0) {
            return true;
        } else {
            return false;
        }
    }
    
    public static int MonthDays(int M, int Y) {
        switch (M) {
            case 1: case 3: case 5: case 7: case 8: case 10: case 12:
                return 31;
            case 4: case 6: case 9: case 11:
                return 30;
            case 2:
                return IsLeapYear(Y) ? 29 : 28;
            default:
                return 0;
        }
    }

    private boolean proc54() {
        TInt d = new TInt(), m = new TInt(), y = new TInt();
        for (int i = 1; i < 4; i++) {
            d.value = getInt();
            m.value = getInt();
            y.value = getInt();
            PrevDate(d, m, y);
            put(d.value);
            put(m.value);
            put(y.value);
        }
        return false;
    }
    
    static void PrevDate(TInt D, TInt M, TInt Y) {
        if ( (D.value == 1) && (M.value == 1) ) {
            D.value = 31;
            M.value = 12;
            Y.value--;
        } else if (D.value == 1) {
            M.value--;
            D.value = MonthDays(M.value, Y.value);
        } else {
            D.value--;
        }
    }
    
    public static boolean IsLeapYear(int Y) {
        if (Y % 400 == 0) {
            return true;
        } else if (Y % 100 == 0) {
            return false;
        } else if (Y % 4 == 0) {
            return true;
        } else {
            return false;
        }
    }
    
    public static int MonthDays(int M, int Y) {
        int days = 0;
        switch (M) {
            case 1: case 3: case 5: case 7: case 8: case 10: case 12:
                days = 31;
                break;
            case 4: case 6: case 9: case 11:
                days = 30;
                break;
            case 2:
                days = IsLeapYear(Y) ? 29 : 28;
                break;
        }
        return days;
    }

    private boolean proc55() {
        TInt D = new TInt(), M = new TInt(), Y = new TInt();
        for (int i = 1; i <= 3; i++) {
            D.value = getInt();
            M.value = getInt();
            Y.value = getInt();
            NextDate(D, M, Y);
            put(D.value);
            put(M.value);
            put(Y.value);
        }
        return false;
    }
    
    static void NextDate(TInt D, TInt M, TInt Y) {
        if ( (D.value == 31) && (M.value == 12) ) {
            D.value = M.value = 1;
            Y.value++;
        } else if (MonthDays(M.value, Y.value) == D.value) {
            D.value = 1;
            M.value++;
        } else {
            D.value++;
        }
    }
    
    public static int MonthDays(int M, int Y) {
        int days = 0;
        switch (M) {
            case 1: case 3: case 5: case 7: case 8: case 10: case 12:
                days = 31;
                break;
            case 4: case 6: case 9: case 11:
                days = 30;
                break;
            case 2:
                days = IsLeapYear(Y) ? 29 : 28;
                break;
        }
        return days;
    }
    
    public static boolean IsLeapYear(int Y) {
        if (Y % 400 == 0) {
            return true;
        } else if (Y % 100 == 0) {
            return false;
        } else if (Y % 4 == 0) {
            return true;
        } else {
            return false;
        }
    }

    private boolean proc56() {
        double xA = getDouble(), yA = getDouble();
        double xB, yB, leng;
        for (int i = 1; i <= 3; i++) {
            xB = getDouble();
            yB = getDouble();
            leng = Leng(xA, yA, xB, yB);
            put(leng);
        }
        return false;
    }
    
    public static double Leng(double xA, double yA, double xB, double yB) {
        double length = Math.sqrt(Math.pow(xB - xA, 2) + Math.pow(yB - yA, 2));
        return length;
    }

    private boolean proc57() {
        double xA = getDouble();
        double yA = getDouble();
        double xB = getDouble();
        double yB = getDouble();
        double xC = getDouble();
        double yC = getDouble();
        double xD = getDouble();
        double yD = getDouble();
        double PABC = Perim(xA, yA, xB, yB, xC, yC);
        double PABD = Perim(xA, yA, xB, yB, xD, yD);
        double PACD = Perim(xA, yA, xC, yC, xD, yD);
        put(PABC);
        put(PABD);
        put(PACD);
        return false;
    }
    
    public static double Perim(double xA, double yA, double xB, double yB, double xC, double yC) {
        double a = Leng(xA, yA, xB, yB);
        double b = Leng(xA, yA, xC, yC);
        double c = Leng(xB, yB, xC, yC);
        double P = a + b + c;
        return P;
    }
    
    public static double Leng(double xA, double yA, double xB, double yB) {
        double length = Math.sqrt( Math.pow(xB - xA, 2) + Math.pow(yB - yA, 2) );
        return length;
    }

    private boolean proc58() {
        double xA = getDouble();
        double yA = getDouble();
        double xB = getDouble();
        double yB = getDouble();
        double xC = getDouble();
        double yC = getDouble();
        double xD = getDouble();
        double yD = getDouble();
        double SABC = Area(xA, yA, xB, yB, xC, yC);
        double SABD = Area(xA, yA, xB, yB, xD, yD);
        double SACD = Area(xA, yA, xC, yC, xD, yD);
        put(SABC);
        put(SABD);
        put(SACD);
        return false;
    }
    
    public static double Leng(double xA, double yA, double xB, double yB) {
        double length = Math.sqrt( Math.pow(xB - xA, 2) + Math.pow(yB - yA, 2) );
        return length;
    }
    
    public static double Perim(double xA, double yA, double xB, double yB, double xC, double yC) {
        double a = Leng(xA, yA, xB, yB);
        double b = Leng(xA, yA, xC, yC);
        double c = Leng(xB, yB, xC, yC);
        double P = a + b + c;
        return P;
    }
    
    public static double Area(double xA, double yA, double xB, double yB, double xC, double yC) {
        double P = Perim(xA, yA, xB, yB, xC, yC);
        double AB = Leng(xA, yA, xB, yB);
        double AC = Leng(xA, yA, xC, yC);
        double BC = Leng(xB, yB, xC, yC);
        double p = P / 2;
        double S = Math.sqrt( p * (p - AB) * (p - AC) * (p - BC) );
        return S;
    }

    private boolean proc59() {
        double xP = getDouble();
        double yP = getDouble();
        double xA = getDouble();
        double yA = getDouble();
        double xB = getDouble();
        double yB = getDouble();
        double xC = getDouble();
        double yC = getDouble();
        double DPAB = Dist(xP, yP, xA, yA, xB, yB);
        double DPAC = Dist(xP, yP, xA, yA, xC, yC);
        double DPBC = Dist(xP, yP, xB, yB, xC, yC);
        put(DPAB);
        put(DPAC);
        put(DPBC);
        return false;
    }
    
    public static double Dist(double xP, double yP, double xA, double yA, double xB, double yB) {
        double SPAB = Area(xP, yP, xA, yA, xB, yB);
        double AB = Leng(xA, yA, xB, yB);
        double D = 2 * SPAB / AB;
        return D;
    }
    
    public static double Leng(double xA, double yA, double xB, double yB) {
        double length = Math.sqrt( Math.pow(xB - xA, 2) + Math.pow(yB - yA, 2) );
        return length;
    }
    
    public static double Area(double xA, double yA, double xB, double yB, double xC, double yC) {
        double P = Perim(xA, yA, xB, yB, xC, yC);
        double p = P / 2;
        double AB = Leng(xA, yA, xB, yB);
        double AC = Leng(xA, yA, xC, yC);
        double BC = Leng(xB, yB, xC, yC);
        double S = Math.sqrt(p * (p - AB) * (p - AC) * (p - BC));
        return S;
    }
    
    public static double Perim(double xA, double yA, double xB, double yB, double xC, double yC) {
        double a = Leng(xA, yA, xB, yB);
        double b = Leng(xA, yA, xC, yC);
        double c = Leng(xB, yB, xC, yC);
        double P = a + b + c;
        return P;
    }

    private boolean proc60() {
        double xA = getDouble();
        double yA = getDouble();
        double xB = getDouble();
        double yB = getDouble();
        double xC = getDouble();
        double yC = getDouble();
        double xD = getDouble();
        double yD = getDouble();
        TDouble ha = new TDouble(), hb = new TDouble(), hc = new TDouble();
        Altitudes(xA, yA, xB, yB, xC, yC, ha, hb, hc);
        put(ha.value); put(hb.value); put(hc.value);
        Altitudes(xA, yA, xB, yB, xD, yD, ha, hb, hc);
        put(ha.value); put(hb.value); put(hc.value);
        Altitudes(xA, yA, xC, yC, xD, yD, ha, hb, hc);
        put(ha.value); put(hb.value); put(hc.value);
        return false;
    }
    
    static void Altitudes(double xA, double yA, double xB, double yB, double xC, double yC, TDouble hA, TDouble hB, TDouble hC) {
        hA.value = Dist(xA, yA, xB, yB, xC, yC);
        hB.value = Dist(xB, yB, xA, yA, xC, yC);
        hC.value = Dist(xC, yC, xA, yA, xB, yB);
    }
    
    public static double Dist(double xP, double yP, double xA, double yA, double xB, double yB) {
        double SPAB = Area(xP, yP, xA, yA, xB, yB);
        double AB = Leng(xA, yA, xB, yB);
        double D = 2 * SPAB / AB;
        return D;
    }
    
    public static double Leng(double xA, double yA, double xB, double yB) {
        double length = Math.sqrt( Math.pow(xB - xA, 2) + Math.pow(yB - yA, 2) );
        return length;
    }
    
    public static double Area(double xA, double yA, double xB, double yB, double xC, double yC) {
        double P = Perim(xA, yA, xB, yB, xC, yC);
        double p = P / 2;
        double AB = Leng(xA, yA, xB, yB);
        double AC = Leng(xA, yA, xC, yC);
        double BC = Leng(xB, yB, xC, yC);
        double S = Math.sqrt( p * (p - AB) * (p - AC) * (p - BC) );
        return S;
    }
    
    public static double Perim(double xA, double yA, double xB, double yB, double xC, double yC) {
        double a = Leng(xA, yA, xB, yB);
        double b = Leng(xA, yA, xC, yC);
        double c = Leng(xB, yB, xC, yC);
        double P = a + b + c;
        return P;
    }
}