import java.util.Scanner;
import static java.lang.Math.pow;

class BeginTask {

	private boolean begin1() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("a = ");
		double a = scanner.nextDouble();
		double P = 4 * a;
		System.out.printf("%1$.2f\n", P);
		return false;
	}

	private boolean begin2() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("a = ");
		double a = scanner.nextDouble();
		double S = a*a;
		System.out.printf("S = %1$.2f\n", S);
		return false;
	}

	private boolean begin3() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("a = ");
		double a = scanner.nextDouble();
		System.out.print("b = ");
		double b = scanner.nextDouble();
		double S = a * b;
		double P = 2 * (a + b);
		System.out.printf("S = %1$.2f\nP = %2$.2f\n", S, P);
		return false;
	}

	private boolean begin4() {
		Scanner scanner = new Scanner(System.in);
		final double PI = 3.14;
		System.out.print("d = ");
		double d = scanner.nextDouble();
		double L = PI * d;
		System.out.printf("L = %1$.2f\n", L);
		return false;
	}

	private boolean begin5() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("a = ");
		double a = scanner.nextDouble();
		double V = pow(a, 3);
		double S = 6 * a * a;
		System.out.printf("V = %1$.2f\nS = %2$.2f\n", V, S);
		return false;
	}

	private boolean begin6() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("a = ");
		double a = scanner.nextDouble();
		System.out.print("b = ");
		double b = scanner.nextDouble();
		System.out.print("c = ");
		double c = scanner.nextDouble();
		double V = a * b * c;
		double S = 2 * (a*b + b*c + a*c);
		System.out.printf("V = %1$.2f\nS = %2$.2f\n", V, S);
		return false;
	}
}