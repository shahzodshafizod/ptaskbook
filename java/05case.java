import java.util.Scanner;

class CaseTask {

	private boolean case1() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("number [1-7] = ");
		int number = scanner.nextInt();
		switch (number) {
			case 1: System.out.println("Monday"); break;
			case 2: System.out.println("Tuesday"); break;
			case 3: System.out.println("Wednesday"); break;
			case 4: System.out.println("Thursday"); break;
			case 5: System.out.println("Friday"); break;
			case 6: System.out.println("Saturday"); break;
			case 7: System.out.println("Sunday"); break;
		}
		return false;
	}

	private boolean case2() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("K = ");
		int K = scanner.nextInt();
		switch (K) {
			case 1: System.out.println("Bad"); break;
			case 2: System.out.println("Ghayriqanoatbaxsh"); break;
			case 3: System.out.println("Qanoatbaxsh"); break;
			case 4: System.out.println("Xub"); break;
			case 5: System.out.println("A'lo"); break;
		}
		return false;
	}

	private boolean case3() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("monthNo = ");
		int monthNo = scanner.nextInt();
		switch (monthNo) {
			case 1:
			case 2:
			case 12:
				System.out.println("Winter");
				break;
			case 3:
			case 4:
			case 5:
				System.out.println("Spring");
				break;
			case 6:
			case 7:
			case 8:
				System.out.println("Summer");
				break;
			case 9:
			case 10:
			case 11:
				System.out.println("Autumn");
				break;
		}
		return false;
	}

	private boolean case4() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("monthNo = ");
		int monthNo = scanner.nextInt();
		int days = 0;
		switch (monthNo) {
			case 1:
			case 3:
			case 5:
			case 7:
			case 8:
			case 10:
			case 12:
				days = 31;
				break;
			case 4:
			case 6:
			case 9:
			case 11:
				days = 30;
				break;
			case 2:
				days = 28;
				break;
		}
		System.out.println("days = " + days);
		return false;
	}

	private boolean case5() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("N = ");
		int N = scanner.nextInt();
		System.out.print("A = ");
		double A = scanner.nextDouble();
		System.out.print("B = ");
		double B = scanner.nextDouble();
		double result = 0;
		char sign = '\0';
		switch (N) {
			case 1:
				result = A + B;
				sign = '+';
				break;
			case 2:
				result = A - B;
				sign = '-';
				break;
			case 3:
				result = A * B;
				sign = '*';
				break;
			case 4:
				result = B != 0 ? A / B : 0;
				sign = '/';
				break;
		}
		System.out.printf("%1$.2f %2$c %3$.2f = %4$.2f\n", A, sign, B, result);
		return false;
	}

	private boolean case6() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("length no:\t");
		int lengthNo = scanner.nextInt();
		System.out.print("length = ");
		double length = scanner.nextDouble();
		double meters = 0;
		switch (lengthNo) {
			case 1: meters = length / 10; break;
			case 2: meters = length * 1000; break;
			case 3: meters = length; break;
			case 4: meters = length / 1000; break;
			case 5: meters = length / 100; break;
		}
		System.out.println("meters = " + meters);
		return false;
	}

	private boolean case7() {
        int no = getInt();
        double kg = 0, value = getDouble();
        switch (no) {
            case 1: kg = value; break;
            case 2: kg = value / 1000000; break;
            case 3: kg = value / 1000; break;
            case 4: kg = value * 1000; break;
            case 5: kg = value * 100; break;
        }
        put(kg);
        return false;
    }

	private boolean case8() {
        int D = getInt(), M = getInt();
        switch (D) {
            case 1:
                switch (M) {
                    case 1: D = 31; M = 12; break;
                    default:
                        switch (M) {
                            case 3: D = 28; break;
                            case 2: case 4: case 6: case 8: case 9: case 11: D = 31; break;
                            default: D = 30; break;
                        }
                        M--;
                        break;
                }
                break;
            default:
                D--;
        }            
        put(D);
        put(M);
        return false;
    }

	private boolean case9() {
        int D = getInt(), M = getInt();
        switch (M) {
            case 12:
                if (D == 31) {
                    D = M = 1;
                } else {
                    D++;
                }
                break;
            case 2:
                if (D == 28) {
                    D = 1;
                    M++;
                } else {
                    D++;
                }
                break;
            case 1: case 3: case 5: case 7: case 8: case 10:
                if (D == 31) {
                    D = 1;
                    M++;
                } else {
                    D++;
                }
                break;
            case 4: case 6: case 9: case 11:
                if (D == 30) {
                    D = 1;
                    M++;
                } else {
                    D++;
                }
                break;
        }
        put(D);
        put(M);
        return false;
    }

	private boolean case10() {
        char C = getChar();
        int N = getInt();
        switch (N) {
            case 1:
                //to left;
                switch (C) {
                    case 'С': C = 'З'; break;
                    case 'Ю': C = 'В'; break;
                    case 'В': C = 'С'; break;
                    case 'З': C = 'Ю'; break;
                }
                break;
            case -1:
                //to right;
                switch (C) {
                    case 'С': C = 'В'; break;
                    case 'Ю': C = 'З'; break;
                    case 'В': C = 'Ю'; break;
                    case 'З': C = 'С'; break;
                }
                break;
        }
        
        put(C);
        return false;
    }

	private boolean case11() {
        char C = getChar();
        int N1 = getInt();
        int N2 = getInt();
        
        //1+1 = 2   180
        //1-1 = 0   continue
        //1+2 = 3   to right

        //-1+1 = 0  continue
        //-1-1 = -2 180
        //-1+2 = 1  to left

        //2+1 = 3   to right
        //2-1 = 1   to left
        //2+2 = 4   continue
        
        int N = Math.abs(N1 + N2);
        switch (N) {
            case 1:
                //to left;
                switch (C) {
                    case 'С': C = 'З'; break;
                    case 'Ю': C = 'В'; break;
                    case 'В': C = 'С'; break;
                    case 'З': C = 'Ю'; break;
                }
                break;
            case 2:
                //turn 180
                switch (C) {
                    case 'С': C = 'Ю'; break;
                    case 'Ю': C = 'С'; break;
                    case 'В': C = 'З'; break;
                    case 'З': C = 'В'; break;
                }
                break;
            case 3:
                //to right
                switch (C) {
                    case 'С': C = 'В'; break;
                    case 'Ю': C = 'З'; break;
                    case 'В': C = 'Ю'; break;
                    case 'З': C = 'С'; break;
                }
                break;
        }
        
        put(C);
        return false;
    }

	private boolean case12() {
        int no = getInt();
        double value = getDouble();
        final double PI = 3.14;
        double R, D, L, S;
        switch (no) {
            case 1:
                R = value;
                D = 2 * R;
                L = 2 * PI * R;
                S = PI * R * R;
                put(D);
                put(L);
                put(S);
                break;
            case 2:
                D = value;
                R = D / 2;
                L = 2 * PI * R;
                S = PI * R * R;
                put(R);
                put(L);
                put(S);
                break;
            case 3:
                L = value;
                R = L / (2 * PI);
                D = 2 * R;
                S = PI * R * R;
                put(R);
                put(D);
                put(S);
                break;
            case 4:
                S = value;
                R = Math.sqrt(S / PI);
                D = 2 * R;
                L = 2 * PI * R;
                put(R);
                put(D);
                put(L);
                break;
        }
        return false;
    }

	private boolean case13() {
        int no = getInt();
        double value = getDouble();
        double a, c, h, S;
        switch (no) {
            case 1:
                a = value;
                c = a * Math.sqrt(2);
                h = c / 2;
                S = c * h / 2;
                put(c);
                put(h);
                put(S);
                break;
            case 2:
                c = value;
                a = c / Math.sqrt(2);
                h = c / 2;
                S = c * h / 2;
                put(a);
                put(h);
                put(S);
                break;
            case 3:
                h = value;
                a = 2 * h / Math.sqrt(2);
                c = a * Math.sqrt(2);
                S = c * h / 2;
                put(a);
                put(c);
                put(S);
                break;
            case 4:
                S = value;
                h = Math.sqrt(S);
                c = 2 * h;
                a = c / Math.sqrt(2);
                put(a);
                put(c);
                put(h);
                break;
        }
        return false;
    }

	private boolean case14() {
        int no = getInt();
        double value = getDouble();
        double a, R1, R2, S;
        switch (no) {
            case 1:
                a = value;
                R1 = a * Math.sqrt(3) / 6;
                R2 = 2 * R1;
                S = a*a * Math.sqrt(3) / 4;
                put(R1);
                put(R2);
                put(S);
                break;
            case 2:
                R1 = value;
                R2 = 2 * R1;
                a = 6 * R1 / Math.sqrt(3);
                S = a*a * Math.sqrt(3) / 4;
                put(a);
                put(R2);
                put(S);
                break;
            case 3:
                R2 = value;
                R1 = R2 / 2;
                a = 6 * R1 / Math.sqrt(3);
                S = a*a * Math.sqrt(3) / 4;
                put(a);
                put(R1);
                put(S);
                break;
            case 4:
                S = value;
                a = Math.sqrt( 4 * S / Math.sqrt(3) );
                R1 = a * Math.sqrt(3) / 6;
                R2 = 2 * R1;
                put(a);
                put(R1);
                put(R2);
                break;
        }
        return false;
    }

	private boolean case15() {
        int N = getInt(), M = getInt();
        String cardName = "";
        switch (N) {
            case 6: cardName += "шестерка"; break;
            case 7: cardName += "семерка"; break;
            case 8: cardName += "восьмерка"; break;
            case 9: cardName += "девятка"; break;
            case 10: cardName += "десятка"; break;
            case 11: cardName += "валет"; break;
            case 12: cardName += "дама"; break;
            case 13: cardName += "король"; break;
            case 14: cardName += "туз"; break;
        }
        cardName += ' ';
        switch (M) {
            case 1: cardName += "пик"; break;
            case 2: cardName += "треф"; break;
            case 3: cardName += "бубен"; break;
            case 4: cardName += "червей"; break;
        }
        put(cardName);
        return false;
    }

	private boolean case16() {
        int number = getInt();
        String descr = "";
        int dahi = number / 10;
        int vohid = number % 10;
        switch (dahi) {
            case 2: descr += "двадцать"; break;
            case 3: descr += "тридцать"; break;
            case 4: descr += "сорок"; break;
            case 5: descr += "пятьдесят"; break;
            case 6: descr += "шестьдесят"; break;
        }
        if (vohid != 0) {
            descr += ' ';
            switch (vohid) {
                case 1: descr += "один"; break;
                case 2: descr += "два"; break;
                case 3: descr += "три"; break;
                case 4: descr += "четыре"; break;
                case 5: descr += "пять"; break;
                case 6: descr += "шесть"; break;
                case 7: descr += "семь"; break;
                case 8: descr += "восемь"; break;
                case 9: descr += "девять"; break;
            }
        }
        descr += ' ';
        switch (vohid) {
            case 1: descr += "год"; break;
            case 2: case 3: case 4: descr += "года"; break;
            default: descr += "лет"; break;
        }
        put(descr);
        return false;
    }

	private boolean case17() {
        int number = getInt();
        int dahi = number / 10;
        int vohid = number % 10;
        String descr = "";
        switch (dahi) {
            case 1:
                switch (vohid) {
                    case 0: descr += "десять"; break;
                    case 1: descr += "одиннадцать"; break;
                    case 2: descr += "двенадцать"; break;
                    case 3: descr += "тринадцать"; break;
                    case 4: descr += "четырнадцать"; break;
                    case 5: descr += "пятнадцать"; break;
                    case 6: descr += "шестнадцать"; break;
                    case 7: descr += "семнадцать"; break;
                    case 8: descr += "восемнадцать"; break;
                    case 9: descr += "девятнадцать"; break;
                }
                break;
            case 2: descr += "двадцать"; break;
            case 3: descr += "тридцать"; break;
            case 4: descr += "сорок"; break;
        }
        if ( (vohid != 0) && (dahi != 1) ) {
            descr += ' ';
            switch (vohid) {
                case 1: descr += "одно"; break;
                case 2: descr += "два"; break;
                case 3: descr += "три"; break;
                case 4: descr += "четыре"; break;
                case 5: descr += "пять"; break;
                case 6: descr += "шесть"; break;
                case 7: descr += "семь"; break;
                case 8: descr += "восемь"; break;
                case 9: descr += "девять"; break;
            }
        }
        descr += " учебн";
        if ( (dahi != 1) && (vohid == 1) ) {
            descr += "ое задание";
        } else {
            descr += "ых задани";
            if ( (dahi != 1) && (vohid != 0) && (vohid < 5) ) {
                descr += 'я';
            } else {
                descr += 'й';
            }
        }
        
        put(descr);
        return false;
    }

	private boolean case18() {
        int number = getInt();
        int sadi = number / 100;
        int dahi = number / 10 % 10;
        int vohid = number % 10;
        String descr = "";
        switch (sadi) {
            case 1: descr += "сто"; break;
            case 2: descr += "двести"; break;
            case 3: descr += "триста"; break;
            case 4: descr += "четыреста"; break;
            case 5: descr += "пятьсот"; break;
            case 6: descr += "шестьсот"; break;
            case 7: descr += "семьсот"; break;
            case 8: descr += "восемьсот"; break;
            case 9: descr += "девятьсот"; break;
        }
        if (dahi != 0) {
            if (sadi != 0) {
                descr += ' ';
            }
            switch (dahi) {
                case 1:
                    switch (vohid) {
                        case 0: descr += "десять"; break;
                        case 1: descr += "одиннадцать"; break;
                        case 2: descr += "двенадцать"; break;
                        case 3: descr += "тринадцать"; break;
                        case 4: descr += "четырнадцать"; break;
                        case 5: descr += "пятнадцать"; break;
                        case 6: descr += "шестнадцать"; break;
                        case 7: descr += "семнадцать"; break;
                        case 8: descr += "восемнадцать"; break;
                        case 9: descr += "девятнадцать"; break;
                    }
                    break;
                case 2: descr += "двадцать"; break;
                case 3: descr += "тридцать"; break;
                case 4: descr += "сорок"; break;
                case 5: descr += "пятьдесят"; break;
                case 6: descr += "шестьдесят"; break;
                case 7: descr += "семьдесят"; break;
                case 8: descr += "восемьдесят"; break;
                case 9: descr += "девяносто"; break;
            }
        }
        if ( (dahi != 1) && (vohid != 0) ) {
            if ( (sadi != 0) || (dahi != 0) ) {
                descr += ' ';
            }
            switch (vohid) {
                case 1: descr += "один"; break;
                case 2: descr += "два"; break;
                case 3: descr += "три"; break;
                case 4: descr += "четыре"; break;
                case 5: descr += "пять"; break;
                case 6: descr += "шесть"; break;
                case 7: descr += "семь"; break;
                case 8: descr += "восемь"; break;
                case 9: descr += "девять"; break;
            }
        }
        put(descr);
        return false;
    }

	private boolean case19() {
        int year = getInt();
        String yearName = "год ";
        
        int color = year % 60;
        if ( (color >= 52) || (color <= 3) ) {
            yearName += "черно";
        } else if (color <= 15) {
            yearName += "зелено";
        } else if (color <= 27) {
            yearName += "красно";
        } else if (color <= 39) {
            yearName += "желто";
        } else if (color <= 51) {
            yearName += "бело";
        }
        int animal = year % 60 % 12;
        switch (animal) {
            case 4: yearName += "й крысы"; break;
            case 5: yearName += "й коровы"; break;
            case 6: yearName += "го тигра"; break;
            case 7: yearName += "го зайца"; break;
            case 8: yearName += "го дракона"; break;
            case 9: yearName += "й змеи"; break;
            case 10: yearName += "й лошади"; break;
            case 11: yearName += "й овцы"; break;
            case 0: yearName += "й обезьяны"; break;
            case 1: yearName += "й курицы"; break;
            case 2: yearName += "й собаки"; break;
            case 3: yearName += "й свиньи"; break;
        }
        put(yearName);
        return false;
    }

	private boolean case20() {
        int D = getInt(), M = getInt();
        String zodiak = "";
        switch (M) {
            case 1:
                if (D < 20) {
                    zodiak = "Козерог";
                } else {
                    zodiak = "Водолей";
                }
                break;
            case 2:
                if (D < 19) {
                    zodiak = "Водолей";
                } else {
                    zodiak = "Рыбы";
                }
                break;
            case 3:
                if (D < 21) {
                    zodiak = "Рыбы";
                } else {
                    zodiak = "Овен";
                }
                break;
            case 4:
                if (D < 20) {
                    zodiak = "Овен";
                } else {
                    zodiak = "Телец";
                }
                break;
            case 5:
                if (D < 21) {
                    zodiak = "Телец";
                } else {
                    zodiak = "Близнецы";
                }
                break;
            case 6:
                if (D < 22) {
                    zodiak = "Близнецы";
                } else {
                    zodiak = "Рак";
                }
                break;
            case 7:
                if (D < 23) {
                    zodiak = "Рак";
                } else {
                    zodiak = "Лев";
                }
                break;
            case 8:
                if (D < 23) {
                    zodiak = "Лев";
                } else {
                    zodiak = "Дева";
                }
                break;
            case 9:
                if (D < 23) {
                    zodiak = "Дева";
                } else {
                    zodiak = "Весы";
                }
                break;
            case 10:
                if (D < 23) {
                    zodiak = "Весы";
                } else {
                    zodiak = "Скорпион";
                }
                break;
            case 11:
                if (D < 23) {
                    zodiak = "Скорпион";
                } else {
                    zodiak = "Стрелец";
                }
                break;
            case 12:
                if (D < 23) {
                    zodiak = "Стрелец";
                } else {
                    zodiak = "Козерог";
                }
                break;
        }
        put(zodiak);
        return false;
    }
}