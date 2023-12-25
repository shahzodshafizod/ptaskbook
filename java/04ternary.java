import java.util.Scanner;

class TernaryTask {

    private boolean ternary2() {
        int number = getInt();
        number += number > 0 ? -8 : 6;
        put(number);
        return false;
    }

    private boolean ternary3() {
        int number = getInt();
        number += number > 0 ? -8 : number < 0 ? 6 : -number + 10;
        put(number);
        return false;
    }

    private boolean ternary4() {
        int a = getInt();
        int b = getInt();
        int c = getInt();
        int pluses = a > 0 ? 1 : 0;
        pluses += b > 0 ? 1 : 0;
        pluses += c > 0 ? 1 : 0;
        put(pluses);
        return false;
    }

    private boolean ternary5() {
        int a = getInt();
        int b = getInt();
        int c = getInt();
        int pluses = a > 0 ? 1 : 0;
        pluses += b > 0 ? 1 : 0;
        pluses += c > 0 ? 1 : 0;
        int minuses = a < 0 ? 1 : 0;
        minuses += b < 0 ? 1 : 0;
        minuses += c < 0 ? 1 : 0;
        put(pluses);
        put(minuses);
        return false;
    }

    private boolean ternary6() {
        double a = getDouble();
        double b = getDouble();
        put( a > b ? a : b );
        return false;
    }

    private boolean ternary7() {
        double a = getDouble();
        double b = getDouble();
        put( a < b ? 1 : 2 );
        return false;
    }

    private boolean ternary8() {
        double a = getDouble();
        double b = getDouble();
        double kalon = a > b ? a : b;
        double xurd = a < b ? a : b;
        put(kalon);
        put(xurd);
        return false;
    }

    private boolean ternary9() {
        double A = getDouble();
        double B = getDouble();
        double kalon = A > B ? A : B;
        double xurd = A < B ? A : B;
        A = xurd;
        B = kalon;
        put(A);
        put(B);
        return false;
    }

    private boolean ternary10() {
        int A = getInt();
        int B = getInt();
        A = B = A != B ? A + B : 0;
        put(A);
        put(B);
        return false;
    }

    private boolean ternary11() {
        int A = getInt();
        int B = getInt();
        A = B = A != B ? A > B ? A : B : 0;
        put(A);
        put(B);
        return false;
    }

    private boolean ternary12() {
        double a = getDouble();
        double b = getDouble();
        double c = getDouble();
        double xurd = (a < b) && (a < c) ? a : b < c ? b : c;
        put(xurd);
        return false;
    }

    private boolean ternary13() {
        double a = getDouble();
        double b = getDouble();
        double c = getDouble();
        double kalon = (a > b) && (a > c) ? a : b > c ? b : c;
        double xurd = (a < b) && (a < c) ? a : b < c ? b : c;
        double middle = a + b + c - kalon - xurd;
        put(middle);
        return false;
    }

    private boolean ternary14() {
        double a = getDouble();
        double b = getDouble();
        double c = getDouble();
        double kalon = (a > b) && (a > c) ? a : b > c ? b : c;
        double xurd = (a < b) && (a < c) ? a : b < c ? b : c;
        put(xurd);
        put(kalon);
        return false;
    }

    private boolean ternary15() {
        double a = getDouble();
        double b = getDouble();
        double c = getDouble();
        double xurd = (a < b) && (a < c) ? a : b < c ? b : c;
        double sum = a + b + c - xurd;
        put(sum);
        return false;
    }

    private boolean ternary16() {
        double A = getDouble();
        double B = getDouble();
        double C = getDouble();
        int factor = (A < B) && (B < C) ? 2 : -1;
        A *= factor;
        B *= factor;
        C *= factor;
        put(A);
        put(B);
        put(C);
        return false;
    }

    private boolean ternary17() {
        double A = getDouble();
        double B = getDouble();
        double C = getDouble();
        int factor = (A < B) && (B < C) || (A > B) && (B > C) ? 2 : -1;
        A *= factor;
        B *= factor;
        C *= factor;
        put(A);
        put(B);
        put(C);
        return false;
    }

    private boolean ternary18() {
        int a = getInt();
        int b = getInt();
        int c = getInt();
        int differentNo = a == b ? 3 : a == c ? 2 : 1;
        put(differentNo);
        return false;
    }

    private boolean ternary19() {
        int a = getInt();
        int b = getInt();
        int c = getInt();
        int d = getInt();
        int differentNo = a == b ? a == c ? 4 : 3 : a == c ? 2 : 1;
        put(differentNo);
        return false;
    }

    private boolean ternary20() {
        double A = getDouble();
        double B = getDouble();
        double C = getDouble();
        double AB = Math.abs(B - A);
        double AC = Math.abs(C - A);
        put( AB < AC ? B : C );
        put( AB < AC ? AB : AC );
        return false;
    }

    private boolean ternary21() {
        int x = getInt();
        int y = getInt();
        int answer = (x == 0) && (y == 0) ? 0 : y == 0 ? 1 : x == 0 ? 2 : 3;
        put(answer);
        return false;
    }

    private boolean ternary22() {
        double x = getDouble();
        double y = getDouble();
        int quarterNo = x > 0 ? y > 0 ? 1 : 4 : y > 0 ? 2 : 3;
        put(quarterNo);
        return false;
    }

    private boolean ternary23() {
        int x1 = getInt();
        int y1 = getInt();
        int x2 = getInt();
        int y2 = getInt();
        int x3 = getInt();
        int y3 = getInt();
        int x4 = x1 == x2 ? x3 : x1 == x3 ? x2 : x1;
        int y4 = y1 == y2 ? y3 : y1 == y3 ? y2 : y1;
        put(x4);
        put(y4);
        return false;
    }

    private boolean ternary24() {
        double x = getDouble();
        double y = x > 0 ? 2 * Math.sin(x) : 6 - x;
        put(y);
        return false;
    }

    private boolean ternary25() {
        int x = getInt();
        int y = (x < -2) || (x > 2) ? 2*x : -3*x;
        put(y);
        return false;
    }

    private boolean ternary26() {
        double x = getDouble();
        double y = x <= 0 ? -x : x >= 2 ? 4 : x*x;
        put(y);
        return false;
    }

    private boolean ternary27() {
        double x = getDouble();
        int y = x < 0 ? 0 : (int)x % 2 == 0 ? 1 : -1;
        put(y);
        return false;
    }

    private boolean ternary28() {
        int year = getInt();
        int days = year % 400 == 0 ? 366 : year % 100 == 0 ? 365 : year % 4 == 0 ? 366 : 365;
        put(days);
        return false;
    }

    private boolean ternary29() {
        int number = getInt();
        String descr = "";
        descr += number == 0 ? "нулевое " : number > 0 ? "положительное " : "отрицательное ";
        descr += (number != 0) && (number % 2 != 0) ? "не" : "";
        descr += number != 0 ? "четное " : "";
        descr += "число";
        put(descr);
        return false;
    }

    private boolean ternary30() {
        int number = getInt();
        String descr = "";
        descr += number % 2 != 0 ? "не" : "";
        descr += "четное ";
        descr += number <= 9 ? "одно" : number <= 99 ? "дву" : "трех";
        descr += "значное число";
        put(descr);
        return false;
    }
}
