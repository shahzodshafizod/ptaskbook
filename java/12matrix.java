import java.util.Scanner;

class MatrixTask {

	private boolean matrix1() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("M = ");
		int M = scanner.nextInt();
		System.out.print("N = ");
		int N = scanner.nextInt();
		int[][] matr = new int [M][N];
		for (int i = 0; i < M; i++) {
			for (int j = 0; j < N; j++) {
				matr[i][j] = (i+1) * 10;
			}
		}
		for (int i = 0; i < M; i++) {
			for (int j = 0; j < N; j++) {
				System.out.print(matr[i][j] + "\t");
			}
			System.out.println();
		}
		return false;
	}

	private boolean matrix2() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("M = ");
		int M = scanner.nextInt();
		System.out.print("N = ");
		int N = scanner.nextInt();
		int[][] matr = new int [M][N];
		for (int i = 0; i < M; i++) {
			for (int j = 0; j < N; j++) {
				matr[i][j] = 5 * (j+1);
			}
		}
		for (int i = 0; i < M; i++) {
			for (int j = 0; j < N; j++) {
				System.out.print(matr[i][j] + "\t");
			}
			System.out.println();
		}
		return false;
	}

	private boolean matrix3() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("M = ");
		int M = scanner.nextInt();
		System.out.print("N = ");
		int N = scanner.nextInt();
		double[] arr = new double [M];
		for (int i = 0; i < M; i++) {
			arr[i] = scanner.nextDouble();
		}
		double[][] matr = new double [M][N];
		for (int i = 0; i < M; i++) {
			for (int j = 0; j < N; j++) {
				matr[i][j] = arr[i];
			}
		}
		for (int i = 0; i < M; i++) {
			for (int j = 0; j < N; j++) {
				System.out.printf("%1$.2f\t", matr[i][j]);
			}
			System.out.println();
		}
		return false;
	}

	private boolean matrix4() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("M = ");
		int M = scanner.nextInt();
		System.out.print("N = ");
		int N = scanner.nextInt();
		double[] arr = new double [N];
		for (int i = 0; i < N; i++) {
			arr[i] = scanner.nextDouble();
		}
		double[][] matr = new double [M][N];
		for (int i = 0; i < M; i++) {
			for (int j = 0; j < N; j++) {
				matr[i][j] = arr[j];
			}
		}
		for (int i = 0; i < M; i++) {
			for (int j = 0; j < N; j++) {
				System.out.printf("%1$.2f\t", matr[i][j]);
			}
			System.out.println();
		}
		return false;
	}

	private boolean matrix5() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("M = ");
		int M = scanner.nextInt();
		System.out.print("N = ");
		int N = scanner.nextInt();
		System.out.print("D = ");
		double D = scanner.nextDouble();
		double[] arr = new double [M];
		for (int i = 0; i < M; i++) {
			arr[i] = scanner.nextDouble();
		}
		double[][] matr = new double [M][N];
		for (int i = 0; i < M; i++) {
			for (int j = 0; j < N; j++) {
				matr[i][j] = j == 0 ? arr[i] : matr[i][j - 1] + D;
			}
		}
		for (int i = 0; i < M; i++) {
			for (int j = 0; j < N; j++) {
				System.out.printf("%1$.2f\t", matr[i][j]);
			}
			System.out.println();
		}
		return false;
	}

	private boolean matrix6() {
		Scanner scanner = new Scanner(System.in);
		System.out.print("M = ");
		int M = scanner.nextInt();
		System.out.print("N = ");
		int N = scanner.nextInt();
		System.out.print("D = ");
		double D = scanner.nextDouble();
		double[] arr = new double [N];
		for (int i = 0; i < N; i++) {
			arr[i] = scanner.nextDouble();
		}
		double[][] matr = new double [M][N];
		for (int i = 0; i < M; i++) {
			for (int j = 0; j < N; j++) {
				if (i == 0) {
					matr[i][j] = arr[j];
				} else {
					matr[i][j] = matr[i-1][j] * D;
				}
			}
		}
		for (int i = 0; i < M; i++) {
			for (int j = 0; j < N; j++) {
				System.out.printf("%1$.2f\t", matr[i][j]);
			}
			System.out.println();
		}
		return false;
	}

	private boolean matrix7() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        int K = getInt();
        for (int j = 0; j < N; j++) {
            put(matr[K-1][j]);
        }
        return false;
    }

	private boolean matrix8() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        int K = getInt();
        for (int i = 0; i < M; i++) {
            put(matr[i][K-1]);
        }
        return false;
    }

	private boolean matrix9() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        for (int i = 1; i < M; i += 2) {
            for (int j = 0; j < N; j++) {
                put(matr[i][j]);
            }
        }
        return false;
    }

	private boolean matrix10() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        for (int j = 0; j < N; j += 2) {
            for (int i = 0; i < M; i++) {
                put(matr[i][j]);
            }
        }
        return false;
    }

	private boolean matrix11() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        int start = -1, finish = N, incrementer = 1;
        int temp;
        for (int i = 0; i < M; i++) {
            for (int j = start + incrementer; j != finish; j += incrementer) {
                put(matr[i][j]);
            }
            temp = start;
            start = finish;
            finish = temp;
            incrementer *= -1;
        }
        return false;
    }

	private boolean matrix12() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        int start = -1, finish = M, incrementer = 1;
        int temp;
        for (int j = 0; j < N; j++) {
            for (int i = start + incrementer; i != finish; i += incrementer) {
                put(matr[i][j]);
            }
            temp = start;
            start = finish;
            finish = temp;
            incrementer *= -1;
        }
        return false;
    }

	private boolean matrix13() {
        int M = getInt();
        double[][] A = new double [M][M];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < M; j++) {
                A[i][j] = getDouble();
            }
        }
        int col;
        for (int i = 0; i < M; i++) {
            col = M - 1 - i;
            for (int j = 0; j <= col; j++) {
                put(A[i][j]);
            }
            for (int j = i+1; j < M; j++) {
                put(A[j][col]);
            }
        }
        return false;
    }

	private boolean matrix14() {
        int M = getInt();
        double[][] A = new double [M][M];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < M; j++) {
                A[i][j] = getDouble();
            }
        }
        int row;
        for (int j = 0; j < M; j++) {
            row = M - 1 - j;
            for (int i = 0; i <= row; i++) {
                put(A[i][j]);
            }
            for (int i = j+1; i < M; i++) {
                put(A[row][i]);
            }
        }
        return false;
    }

	private boolean matrix15() {
        int M = getInt();
        double[][] A = new double [M][M];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < M; j++) {
                A[i][j] = getDouble();
            }
        }
        int index, limit = M / 2 + M % 2;
        for (int i = 0; i < limit; i++) {
            index = M - i - 1;
            for (int j = i; j <= index; j++) {
                put(A[i][j]);
            }
            for (int j = i+1; j <= index; j++) {
                put(A[j][index]);
            }
            for (int j = index-1; j >= i; j--) {
                put(A[index][j]);
            }
            for (int j = index-1; j > i; j--) {
                put(A[j][i]);
            }
        }
        return false;
    }

	private boolean matrix16() {
        int M = getInt();
        double[][] A = new double [M][M];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < M; j++) {
                A[i][j] = getDouble();
            }
        }
        int index, limit = M / 2 + M % 2;
        for (int j = 0; j < limit; j++) {
            index = M - 1 - j;
            for (int i = j; i <= index; i++) {
                put(A[i][j]);
            }
            for (int i = j+1; i <= index; i++) {
                put(A[index][i]);
            }
            for (int i = index-1; i >= j; i--) {
                put(A[i][index]);
            }
            for (int i = index-1; i > j; i--) {
                put(A[j][i]);
            }
        }
        return false;
    }

	private boolean matrix17() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        int K = getInt();
        double sum = 0, zarb = 1;
        for (int j = 0; j < N; j++) {
            sum += matr[K-1][j];
            zarb *= matr[K-1][j];
        }
        put(sum);
        put(zarb);
        return false;
    }

	private boolean matrix18() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++)  {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        int K = getInt();
        double sum = 0, zarb = 1;
        for (int i = 0; i < M; i++) {
            sum += matr[i][K-1];
            zarb *= matr[i][K-1];
        }
        put(sum);
        put(zarb);
        return false;
    }

	private boolean matrix19() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        double sum;
        for (int i = 0; i < M; i++) {
            sum = 0;
            for (int j = 0; j < N; j++) {
                sum += matr[i][j];
            }
            put(sum);
        }
        return false;
    }

	private boolean matrix20() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        double zarb;
        for (int j = 0; j < N; j++) {
            zarb = 1;
            for (int i = 0; i < M; i++) {
                zarb *= matr[i][j];
            }
            put(zarb);
        }
        return false;
    }

	private boolean matrix21() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        double sum, amean;
        for (int i = 0; i < M; i += 2) {
            sum = 0;
            for (int j = 0; j < N; j++) {
                sum += matr[i][j];
            }
            amean = sum / N;
            put(amean);
        }
        return false;
    }

	private boolean matrix22() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        double sum;
        for (int j = 1; j < N; j += 2) {
            sum = 0;
            for (int i = 0; i < M; i++) {
                sum += matr[i][j];
            }
            put(sum);
        }
        return false;
    }

	private boolean matrix23() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        double minimal;
        for (int i = 0; i < M; i++) {
            minimal = matr[i][0];
            for (int j = 1; j < N; j++) {
                if (matr[i][j] < minimal) {
                    minimal = matr[i][j];
                }
            }
            put(minimal);
        }
        return false;
    }

	private boolean matrix24() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        double maximal;
        for (int j = 0; j < N; j++) {
            maximal = matr[0][j];
            for (int i = 1; i < M; i++) {
                if (matr[i][j] > maximal) {
                    maximal = matr[i][j];
                }
            }
            put(maximal);
        }
        return false;
    }

	private boolean matrix25() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        double sum, maxSum = 0;
        int maxSumRow = -1;
        for (int i = 0; i < M; i++) {
            sum = 0;
            for (int j = 0; j < N; j++) {
                sum += matr[i][j];
            }
            if (maxSumRow < 0) {
                maxSum = sum;
                maxSumRow = i;
            } else if (sum > maxSum) {
                maxSum = sum;
                maxSumRow = i;
            }
        }
        put(maxSumRow+1);
        put(maxSum);
        return false;
    }

	private boolean matrix26() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        double zarb, minZarb = 0;
        int minZarbCol = -1;
        for (int j = 0; j < N; j++) {
            zarb = 1;
            for (int i = 0; i < M; i++) {
                zarb *= matr[i][j];
            }
            if (minZarbCol < 0) {
                minZarbCol = j;
                minZarb = zarb;
            } else if (zarb < minZarb) {
                minZarbCol = j;
                minZarb = zarb;
            }
        }
        put(minZarbCol+1);
        put(minZarb);
        return false;
    }

	private boolean matrix27() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        double minimal, maximalMin = 0;
        for (int i = 0; i < M; i++) {
            minimal = matr[i][0];
            for (int j = 1; j < N; j++) {
                if (matr[i][j] < minimal) {
                    minimal = matr[i][j];
                }
            }
            if (maximalMin == 0) {
                maximalMin = minimal;
            } else if (minimal > maximalMin) {
                maximalMin = minimal;
            }
        }
        put(maximalMin);
        return false;
    }

	private boolean matrix28() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        double maximal, minimalMax = 0;
        for (int j = 0; j < N; j++) {
            maximal = matr[0][j];
            for (int i = 1; i < M; i++) {
                if (matr[i][j] > maximal) {
                    maximal = matr[i][j];
                }
            }
            if (minimalMax == 0) {
                minimalMax = maximal;
            } else if (maximal < minimalMax) {
                minimalMax = maximal;
            }
        }
        put(minimalMax);
        return false;
    }

	private boolean matrix29() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        double sum, amean;
        int count;
        for (int i = 0; i < M; i++) {
            sum = 0;
            for (int j = 0; j < N; j++) {
                sum += matr[i][j];
            }
            amean = sum / N;
            count = 0;
            for (int j = 0; j < N; j++) {
                if (matr[i][j] < amean) {
                    count++;
                }
            }
            put(count);
        }
        return false;
    }

	private boolean matrix30() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        double sum, amean;
        int count;
        for (int j = 0; j < N; j++) {
            sum = 0;
            for (int i = 0; i < M; i++) {
                sum += matr[i][j];
            }
            amean = sum / M;
            count = 0;
            for (int i = 0; i < M; i++) {
                if (matr[i][j] > amean) {
                    count++;
                }
            }
            put(count);
        }
        return false;
    }

	private boolean matrix31() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        double sum = 0;
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
                sum += matr[i][j];
            }
        }
        double amean = sum / (M * N);
        double distance, closer = -1;
        int count = 0, row = 0, col = 0;
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                distance = Math.abs(matr[i][j] - amean);
                if (closer < 0) {
                    closer = distance;
                    row = i;
                    col = j;
                } else if (distance < closer) {
                    closer = distance;
                    row = i;
                    col = j;
                }
            }
        }
        put(row+1);
        put(col+1);
        return false;
    }

	private boolean matrix32() {
        int M = getInt();
        int N = getInt();
        int[][] matr = new int [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getInt();
            }
        }
        int positives, negatives, row = -1;
        for (int i = 0; i < M; i++) {
            positives = negatives = 0;
            for (int j = 0; j < N; j++) {
                if (matr[i][j] > 0) {
                    positives++;
                } else if (matr[i][j] < 0) {
                    negatives++;
                }
            }
            if (positives == negatives) {
                row = i;
                break;
            }
        }
        put(row+1);
        return false;
    }

	private boolean matrix33() {
        int M = getInt();
        int N = getInt();
        int[][] matr = new int [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getInt();
            }
        }
        int positives, negatives, col = -1;
        for (int j = 0; j < N; j++) {
            positives = negatives = 0;
            for (int i = 0; i < M; i++) {
                if (matr[i][j] > 0) {
                    positives++;
                } else if (matr[i][j] < 0) {
                    negatives++;
                }
            }
            if (positives == negatives) {
                col = j;
            }
        }
        put(col+1);
        return false;
    }

	private boolean matrix34() {
        int M = getInt();
        int N = getInt();
        int[][] matr = new int [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getInt();
            }
        }
        int row = -1;
        boolean hasOdd;
        for (int i = 0; i < M; i++) {
            hasOdd = false;
            for (int j = 0; j < N; j++) {
                if (matr[i][j] % 2 != 0) {
                    hasOdd = true;
                    break;
                }
            }
            if (!hasOdd) {
                row = i;
            }
        }
        put(row+1);
        return false;
    }

	private boolean matrix35() {
        int M = getInt();
        int N = getInt();
        int[][] matr = new int [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getInt();
            }
        }
        int col = -1;
        boolean hasEven;
        for (int j = 0; j < N; j++) {
            hasEven = false;
            for (int i = 0; i < M; i++) {
                if (matr[i][j] % 2 == 0) {
                    hasEven = true;
                    break;
                }
            }
            if (!hasEven) {
                col = j;
                break;
            }
        }
        put(col+1);
        return false;
    }

	private boolean matrix36() {
        int M = getInt();
        int N = getInt();
        int[][] matr = new int [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getInt();
            }
        }
        boolean[] first = new boolean [101];
        boolean[] other = new boolean [101];
        for (int i = 0; i <= 100; i++) {
            first[i] = other[i] = false;
        }
        for (int j = 0; j < N; j++) {
            first[matr[0][j]] = true;
        }
        boolean same;
        int rows = 0;
        for (int i = 1; i < M; i++) {
            for (int k = 0; k <= 100; k++) {
                other[k] = false;
            }
            for (int j = 0; j < N; j++) {
                other[matr[i][j]] = true;
            }
            same = true;
            for (int k = 0; k <= 100; k++) {
                if (first[k] != other[k]) {
                    same = false;
                    break;
                }
            }
            if (same) {
                rows++;
            }
        }
        put(rows);
        return false;
    }

	private boolean matrix37() {
        int M = getInt();
        int N = getInt();
        int[][] matr = new int [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getInt();
            }
        }
        boolean[] last = new boolean [101];
        boolean[] other = new boolean [101];
        for (int k = 0; k < 101; k++) {
            last[k] = false;
        }
        for (int i = 0; i < M; i++) {
            last[matr[i][N-1]] = true;
        }
        boolean same;
        int cols = 0;
        for (int j = N-2; j >= 0; j--) {
            for (int k = 0; k < 101; k++) {
                other[k] = false;
            }
            for (int i = 0; i < M; i++) {
                other[matr[i][j]] = true;
            }
            same = true;
            for (int k = 0; k < 101; k++) {
                if (last[k] != other[k]) {
                    same = false;
                    break;
                }
            }
            if (same) {
                cols++;
            }
        }
        put(cols);
        return false;
    }

	private boolean matrix38() {
        int M = getInt();
        int N = getInt();
        int[][] matr = new int [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getInt();
            }
        }
        int rows = 0;
        boolean different;
        boolean[] arr = new boolean [101];
        for (int i = 0; i < M; i++) {
            for (int k = 0; k < 101; k++) {
                arr[k] = false;
            }
            different = true;
            for (int j = 0; j < N; j++) {
                if (arr[matr[i][j]]) {
                    different = false;
                    break;
                }
                arr[matr[i][j]] = true;
            }
            if (different) {
                rows++;
            }
        }
        put(rows);
        return false;
    }

	private boolean matrix39() {
        int M = getInt();
        int N = getInt();
        int[][] matr = new int [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getInt();
            }
        }
        int cols = 0;
        boolean[] arr = new boolean [101];
        boolean different;
        for (int j = 0; j < N; j++) {
            for (int k = 0; k < 101; k++) {
                arr[k] = false;
            }
            different = true;
            for (int i = 0; i < M; i++) {
                if (arr[matr[i][j]]) {
                    different = false;
                    break;
                }
                arr[matr[i][j]] = true;
            }
            if (different) {
                cols++;
            }
        }
        put(cols);
        return false;
    }

	private boolean matrix40() {
        int M = getInt();
        int N = getInt();
        int[][] matr = new int [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getInt();
            }
        }
        int[] arr = new int [101];
        int maximal, MAXIMAL = 0, rowIndex = -1;
        for (int i = 0; i < M; i++) {
            for (int k = 0; k < 101; k++) {
                arr[k] = 0;
            }
            for (int j = 0; j < N; j++) {
                arr[matr[i][j]]++;
            }
            maximal = arr[0];
            for (int k = 1; k < 101; k++) {
                if (arr[k] > maximal) {
                    maximal = arr[k];
                }
            }
            if (rowIndex < 0) {
                rowIndex = i;
                MAXIMAL = maximal;
            } else if (maximal >= MAXIMAL) {
                MAXIMAL = maximal;
                rowIndex = i;
            }
        }
        put(rowIndex+1);
        return false;
    }

	private boolean matrix41() {
        int M = getInt();
        int N = getInt();
        int[][] matr = new int [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getInt();
            }
        }
        int colIndex = -1;
        int[] arr = new int [101];
        int maximal, MAXIMAL = 0;
        for (int j = 0; j < N; j++) {
            for (int k = 0; k < 101; k++) {
                arr[k] = 0;
            }
            for (int i = 0; i < M; i++) {
                arr[matr[i][j]]++;
            }
            maximal = arr[0];
            for (int k = 1; k < 101; k++) {
                if (arr[k] > maximal) {
                    maximal = arr[k];
                }
            }
            if (colIndex < 0) {
                colIndex = j;
                MAXIMAL = maximal;
            } else if (maximal > MAXIMAL) {
                MAXIMAL = maximal;
                colIndex = j;
            }
        }
        put(colIndex+1);
        return false;
    }

	private boolean matrix42() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        boolean progress;
        int sortedRows = 0;
        for (int i = 0; i < M; i++) {
            progress = true;
            for (int j = 1; j < N; j++) {
                if (matr[i][j-1] > matr[i][j]) {
                    progress = false;
                    break;
                }
            }
            if (progress) {
                sortedRows++;
            }
        }
        put(sortedRows);
        return false;
    }

	private boolean matrix43() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        int sortedCols = 0;
        boolean regress;
        for (int j = 0; j < N; j++) {
            regress = true;
            for (int i = 1; i < M; i++) {
                if (matr[i-1][j] < matr[i][j]) {
                    regress = false;
                    break;
                }
            }
            if (regress) {
                sortedCols++;
            }
        }
        put(sortedCols);
        return false;
    }

	private boolean matrix44() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        double minimal = 0, MINIMAL = 0;
        boolean progress, regress;
        for (int i = 0; i < M; i++) {
            progress = regress = true;
            minimal = matr[i][0];
            for (int j = 1; j < N; j++) {
                if (matr[i][j] < minimal) {
                    minimal = matr[i][j];
                }
                if (matr[i][j-1] > matr[i][j]) {
                    progress = false;
                }
                if (matr[i][j-1] < matr[i][j]) {
                    regress = false;
                }
            }
            if (progress || regress) {
                if (MINIMAL == 0) {
                    MINIMAL = minimal;
                } else if (minimal < MINIMAL) {
                    MINIMAL = minimal;
                }
            }
        }
        put(MINIMAL);
        return false;
    }

	private boolean matrix45() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        boolean progress, regress;
        double maximal = 0, MAXIMAL = 0;
        for (int j = 0; j < N; j++) {
            progress = regress = true;
            maximal = matr[0][j];
            for (int i = 1; i < M; i++) {
                if (matr[i][j] > maximal) {
                    maximal = matr[i][j];
                }
                if (matr[i-1][j] < matr[i][j]) {
                    regress = false;
                }
                if (matr[i-1][j] > matr[i][j]) {
                    progress = false;
                }
            }
            if (progress || regress) {
                if (MAXIMAL == 0) {
                    MAXIMAL = maximal;
                } else if (maximal > MAXIMAL) {
                    MAXIMAL = maximal;
                }
            }
        }
        put(MAXIMAL);
        return false;
    }

	private boolean matrix46() {
        int M = getInt();
        int N = getInt();
        int[][] matr = new int [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getInt();
            }
        }
        int[] maximal = new int [M];
        for (int i = 0; i < M; i++) {
            maximal[i] = matr[i][0];
            for (int j = 1; j < N; j++) {
                if (matr[i][j] > maximal[i]) {
                    maximal[i] = matr[i][j];
                }
            }
        }
        int[] minimal = new int [N];
        for (int j = 0; j < N; j++) {
            minimal[j] = matr[0][j];
            for (int i = 1; i < M; i++) {
                if (matr[i][j] < minimal[j]) {
                    minimal[j] = matr[i][j];
                }
            }
        }
        int minMax = 0;
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                if ( (matr[i][j] == maximal[i]) && (matr[i][j] == minimal[j]) ) {
                    minMax = matr[i][j];
                }
            }
        }
        put(minMax);
        return false;
    }

	private boolean matrix47() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        int K1 = getInt();
        int K2 = getInt();
        double temp;
        for (int j = 0; j < N; j++) {
            temp = matr[K1-1][j];
            matr[K1-1][j] = matr[K2-1][j];
            matr[K2-1][j] = temp;
        }
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                put(matr[i][j]);
            }
        }
        return false;
    }

	private boolean matrix48() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        int K1 = getInt();
        int K2 = getInt();
        double temp;
        for (int i = 0; i < M; i++) {
            temp = matr[i][K1-1];
            matr[i][K1-1] = matr[i][K2-1];
            matr[i][K2-1] = temp;
        }
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                put(matr[i][j]);
            }
        }
        return false;
    }

	private boolean matrix49() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        int minIndex, maxIndex;
        double temp;
        for (int i = 0; i < M; i++) {
            minIndex = maxIndex = 0;
            for (int j = 1; j < N; j++) {
                if (matr[i][j] > matr[i][maxIndex]) {
                    maxIndex = j;
                }
                if (matr[i][j] < matr[i][minIndex]) {
                    minIndex = j;
                }
            }
            temp = matr[i][minIndex];
            matr[i][minIndex] = matr[i][maxIndex];
            matr[i][maxIndex] = temp;
        }
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                put(matr[i][j]);
            }
        }
        return false;
    }

	private boolean matrix50() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        int minIndex, maxIndex;
        double temp;
        for (int j = 0; j < N; j++) {
            minIndex = maxIndex = 0;
            for (int i = 1; i < M; i++) {
                if (matr[i][j] > matr[maxIndex][j]) {
                    maxIndex = i;
                }
                if (matr[i][j] < matr[minIndex][j]) {
                    minIndex = i;
                }
            }
            temp = matr[minIndex][j];
            matr[minIndex][j] = matr[maxIndex][j];
            matr[maxIndex][j] = temp;
        }
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                put(matr[i][j]);
            }
        }
        return false;
    }

	private boolean matrix51() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        int minRow, maxRow;
        double minimal, maximal;
        minRow = maxRow = 0;
        minimal = maximal = matr[0][0];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                if (matr[i][j] > maximal) {
                    maximal = matr[i][j];
                    maxRow = i;
                }
                if (matr[i][j] < minimal) {
                    minimal = matr[i][j];
                    minRow = i;
                }
            }
        }
        double temp;
        for (int j = 0; j < N; j++) {
            temp = matr[minRow][j];
            matr[minRow][j] = matr[maxRow][j];
            matr[maxRow][j] = temp;
        }
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                put(matr[i][j]);
            }
        }
        return false;
    }

	private boolean matrix52() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        double maximal, minimal;
        int minCol, maxCol;
        minCol = maxCol = 0;
        minimal = maximal = matr[0][0];
        for (int j = 0; j < N; j++) {
            for (int i = 0; i < M; i++) {
                if (matr[i][j] > maximal) {
                    maximal = matr[i][j];
                    maxCol = j;
                }
                if (matr[i][j] < minimal) {
                    minimal = matr[i][j];
                    minCol = j;
                }
            }
        }
        double temp;
        for (int i = 0; i < M; i++) {
            temp = matr[i][minCol];
            matr[i][minCol] = matr[i][maxCol];
            matr[i][maxCol] = temp;
        }
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                put(matr[i][j]);
            }
        }
        return false;
    }

	private boolean matrix53() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        int positiveCol = 0;
        boolean hasNegative;
        for (int j = 1; j < N; j++) {
            hasNegative = false;
            for (int i = 0; i < M; i++) {
                if (matr[i][j] < 0) {
                    hasNegative = true;
                    break;
                }
            }
            if (!hasNegative) {
                positiveCol = j;
            }
        }
        if (positiveCol > 0) {
            double temp;
            for (int i = 0; i < M; i++) {
                temp = matr[i][0];
                matr[i][0] = matr[i][positiveCol];
                matr[i][positiveCol] = temp;
            }
        }
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                put(matr[i][j]);
            }
        }
        return false;
    }

	private boolean matrix54() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        int positiveCol = N-1;
        boolean hasPositive;
        for (int j = 0; j < N-1; j++) {
            hasPositive = false;
            for (int i = 0; i < M; i++) {
                if (matr[i][j] > 0) {
                    hasPositive = true;
                    break;
                }
            }
            if (!hasPositive) {
                positiveCol = j;
                break;
            }
        }
        if (positiveCol < N-1) {
            double temp;
            for (int i = 0; i < M; i++) {
                temp = matr[i][N-1];
                matr[i][N-1] = matr[i][positiveCol];
                matr[i][positiveCol] = temp;
            }
        }
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                put(matr[i][j]);
            }
        }
        return false;
    }

	private boolean matrix55() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        double temp;
        for (int a = 0, b = M/2; b < M; a++, b++) {
            for (int j = 0; j < N; j++) {
                temp = matr[a][j];
                matr[a][j] = matr[b][j];
                matr[b][j] = temp;
            }
        }
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                put(matr[i][j]);
            }
        }
        return false;
    }

	private boolean matrix56() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        double temp;
        for (int a = 0, b = N/2; b < N; a++, b++) {
            for (int i = 0; i < M; i++) {
                temp = matr[i][a];
                matr[i][a] = matr[i][b];
                matr[i][b] = temp;
            }
        }
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                put(matr[i][j]);
            }
        }
        return false;
    }

	private boolean matrix57() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        double temp = 0;
        for (int i = 0, k = M/2; k < M; i++, k++) {
            for (int j = 0, l = N/2; l < N; j++, l++) {
                temp = matr[i][j];
                matr[i][j] = matr[k][l];
                matr[k][l] = temp;
            }
        }
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                put(matr[i][j]);
            }
        }
        return false;
    }

	private boolean matrix58() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        double temp;
        for (int i = M/2, k = 0; i < M; i++, k++) {
            for (int j = 0, l = N/2; l < N; j++, l++) {
                temp = matr[i][j];
                matr[i][j] = matr[k][l];
                matr[k][l] = temp;
            }
        }
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                put(matr[i][j]);
            }
        }
        return false;
    }

	private boolean matrix59() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        double temp;
        for (int start = 0, finish = M-1; start < finish; start++, finish--) {
            for (int j = 0; j < N; j++) {
                temp = matr[start][j];
                matr[start][j] = matr[finish][j];
                matr[finish][j] = temp;
            }
        }
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                put(matr[i][j]);
            }
        }
        return false;
    }

	private boolean matrix60() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        double temp;
        for (int start = 0, finish = N-1; start < finish; start++, finish--) {
            for (int i = 0; i < M; i++) {
                temp = matr[i][start];
                matr[i][start] = matr[i][finish];
                matr[i][finish] = temp;
            }
        }
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                put(matr[i][j]);
            }
        }
        return false;
    }

	private boolean matrix61() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        int K = getInt();
        double[][] temp = new double [M-1][N];
        for (int i = 0, l = 0; i < M; i++) {
            if (i == K-1) {
                continue;
            }
            for (int j = 0; j < N; j++) {
                temp[l][j] = matr[i][j];
            }
            l++;
        }
        matr = temp;
        temp = null;
        M--;
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                put(matr[i][j]);
            }
        }
        return false;
    }

	private boolean matrix62() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        int K = getInt();
        double[][] temp = new double [M][N-1];
        for (int j = 0, l = 0; j < N; j++) {
            if (j == K-1) {
                continue;
            }
            for (int i = 0; i < M; i++) {
                temp[i][l] = matr[i][j];
            }
            l++;
        }
        matr = temp;
        temp = null;
        N--;
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                put(matr[i][j]);
            }
        }
        return false;
    }

	private boolean matrix63() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        double minimal = matr[0][0];
        int minimalRow = 0;
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                if (matr[i][j] < minimal) {
                    minimal = matr[i][j];
                    minimalRow = i;
                }
            }
        }
        double[][] temp = new double [M-1][N];
        for (int i = 0, l = 0; i < M; i++) {
            if (i == minimalRow) {
                continue;
            }
            for (int j = 0; j < N; j++) {
                temp[l][j] = matr[i][j];
            }
            l++;
        }
        matr = temp;
        temp = null;
        M--;
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                put(matr[i][j]);
            }
        }
        return false;
    }

	private boolean matrix64() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        double maximal = matr[0][0];
        int maximalCol = 0;
        for (int j = 0; j < N; j++) {
            for (int i = 0; i < M; i++) {
                if (matr[i][j] > maximal) {
                    maximal = matr[i][j];
                    maximalCol = j;
                }
            }
        }
        double[][] temp = new double [M][N-1];
        for (int j = 0, k = 0; j < N; j++) {
            if (j == maximalCol) {
                continue;
            }
            for (int i = 0; i < M; i++) {
                temp[i][k] = matr[i][j];
            }
            k++;
        }
        matr = temp;
        temp = null;
        N--;
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                put(matr[i][j]);
            }
        }
        return false;
    }

	private boolean matrix65() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        boolean isPositiveCol;
        int positiveCol = -1;
        for (int j = 0; j < N; j++) {
            isPositiveCol = true;
            for (int i = 0; i < M; i++) {
                if (matr[i][j] < 0) {
                    isPositiveCol = false;
                    break;
                }
            }
            if (isPositiveCol) {
                positiveCol = j;
                break;
            }
        }
        if (positiveCol >= 0) {
            double[][] temp = new double [M][N-1];
            for (int j = 0, k = 0; j < N; j++) {
                if (j == positiveCol) {
                    continue;
                }
                for (int i = 0; i < M; i++) {
                    temp[i][k] = matr[i][j];
                }
                k++;
            }
            matr = temp;
            temp = null;
            N--;
        }
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                put(matr[i][j]);
            }
        }
        return false;
    }

	private boolean matrix66() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        boolean isNegativeCol;
        int negativeCol = -1;
        for (int j = 0; j < N; j++) {
            isNegativeCol = true;
            for (int i = 0; i < M; i++) {
                if (matr[i][j] > 0) {
                    isNegativeCol = false;
                    break;
                }
            }
            if (isNegativeCol) {
                negativeCol = j;
            }
        }
        if (negativeCol >= 0) {
            double[][] temp = new double [M][N-1];
            for (int j = 0, k = 0; j < N; j++) {
                if (j == negativeCol) {
                    continue;
                }
                for (int i = 0; i < M; i++) {
                    temp[i][k] = matr[i][j];
                }
                k++;
            }
            matr = temp;
            temp = null;
            N--;
        }
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                put(matr[i][j]);
            }
        }
        return false;
    }

	private boolean matrix67() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        boolean isPositiveCol;
        for (int j = 0; j < N; ) {
            isPositiveCol = true;
            for (int i = 0; i < M; i++) {
                if (matr[i][j] < 0) {
                    isPositiveCol = false;
                    break;
                }
            }
            if (isPositiveCol) {
                double[][] temp = new double [M][N-1];
                for (int col1 = 0, col2 = 0; col1 < N; col1++) {
                    if (col1 == j) {
                        continue;
                    }
                    for (int row = 0; row < M; row++) {
                        temp[row][col2] = matr[row][col1];
                    }
                    col2++;
                }
                matr = temp;
                temp = null;
                N--;
            } else {
                j++;
            }
        }
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                put(matr[i][j]);
            }
        }
        return false;
    }

	private boolean matrix68() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        int K = getInt();
        double[][] temp = new double [M+1][N];
        for (int i = 0, k = 0; i < M; i++) {
            if (i == K-1) {
                for (int j = 0; j < N; j++) {
                    temp[k][j] = 0;
                }
                k++;
            }
            for (int j = 0; j < N; j++) {
                temp[k][j] = matr[i][j];
            }
            k++;
        }
        matr = temp;
        temp = null;
        M++;
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                put(matr[i][j]);
            }
        }
        return false;
    }

	private boolean matrix69() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        int K = getInt();
        double[][] temp = new double [M][N+1];
        for (int j = 0, l = 0; j < N; j++) {
            for (int i = 0; i < M; i++) {
                temp[i][l] = matr[i][j];
            }
            l++;
            if (j == K-1) {
                for (int i = 0; i < M; i++) {
                    temp[i][l] = 1;
                }
                l++;
            }
        }
        matr = temp;
        temp = null;
        N++;
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                put(matr[i][j]);
            }
        }
        return false;
    }

	private boolean matrix70() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        double maximal = matr[0][0];
        int maximalRow = 0;
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                if (matr[i][j] > maximal) {
                    maximal = matr[i][j];
                    maximalRow = i;
                }
            }
        }
        double[][] temp = new double [M+1][N];
        for (int i = 0, k = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                temp[k][j] = matr[i][j];
            }
            k++;
            if (i == maximalRow) {
                for (int j = 0; j < N; j++) {
                    temp[k][j] = matr[i][j];
                }
                k++;
            }
        }
        matr = temp;
        temp = null;
        M++;
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                put(matr[i][j]);
            }
        }
        return false;
    }

	private boolean matrix71() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        double minimal = matr[0][0];
        int minimalCol = 0;
        for (int j = 0; j < N; j++) {
            for (int i = 0; i < M; i++) {
                if (matr[i][j] < minimal) {
                    minimal = matr[i][j];
                    minimalCol = j;
                }
            }
        }
        double[][] temp = new double [M][N+1];
        for (int j = 0, k = 0; j < N; j++) {
            if (j == minimalCol) {
                for (int i = 0; i < M; i++) {
                    temp[i][k] = matr[i][j];
                }
                k++;
            }
            for (int i = 0; i < M; i++) {
                temp[i][k] = matr[i][j];
            }
            k++;
        }
        matr = temp;
        temp = null;
        N++;
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                put(matr[i][j]);
            }
        }
        return false;
    }

	private boolean matrix72() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        boolean isPositiveCol;
        int positiveCol = -1;
        for (int j = 0; j < N; j++) {
            isPositiveCol = true;
            for (int i = 0; i < M; i++) {
                if (matr[i][j] < 0) {
                    isPositiveCol = false;
                    break;
                }
            }
            if (isPositiveCol) {
                positiveCol = j;
                break;
            }
        }
        if (positiveCol >= 0) {
            double[][] temp = new double [M][N+1];
            for (int j = 0, k = 0; j < N; j++) {
                if (positiveCol == j) {
                    for (int i = 0; i < M; i++) {
                        temp[i][k] = 1;
                    }
                    k++;
                }
                for (int i = 0; i < M; i++) {
                    temp[i][k] = matr[i][j];
                }
                k++;
            }
            matr = temp;
            temp = null;
            N++;
        }
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                put(matr[i][j]);
            }
        }
        return false;
    }

	private boolean matrix73() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        boolean isNegativeCol;
        int negativeCol = -1;
        for (int j = 0; j < N; j++) {
            isNegativeCol = true;
            for (int i = 0; i < M; i++) {
                if (matr[i][j] > 0) {
                    isNegativeCol = false;
                    break;
                }
            }
            if (isNegativeCol) {
                negativeCol = j;
            }
        }
        if (negativeCol >= 0) {
            double[][] temp = new double [M][N+1];
            for (int j = 0, k = 0; j < N; j++) {
                for (int i = 0; i < M; i++) {
                    temp[i][k] = matr[i][j];
                }
                k++;
                if (negativeCol == j) {
                    for (int i = 0; i < M; i++) {
                        temp[i][k] = 0;
                    }
                    k++;
                }
            }
            matr = temp;
            temp = null;
            N++;
        }
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                put(matr[i][j]);
            }
        }
        return false;
    }

	private boolean matrix74() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        double[][] temp = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
                temp[i][j] = matr[i][j];
            }
        }
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                if ( (i > 0) || (j > 0) ) {
                    if ( (i > 0) && (temp[i][j] > temp[i-1][j]) ) {
                        continue;
                    }
                    if ( (j > 0) && (temp[i][j] > temp[i][j-1]) ) {
                        continue;
                    }
                    if ( (i > 0) && (j > 0) && (temp[i][j] > temp[i-1][j-1]) ) {
                        continue;
                    }
                    if ( (j > 0) && (i < M-1) && (temp[i][j] > temp[i+1][j-1]) ) {
                        continue;
                    }
                }
                if ( (i < M-1) || (j < N-1) ) {
                    if ( (i < M-1) && (temp[i][j] > temp[i+1][j]) ) {
                        continue;
                    }
                    if ( (j < N-1) && (temp[i][j] > temp[i][j+1]) ) {
                        continue;
                    }
                    if ( (i < M-1) && (j < N-1) && (temp[i][j] > temp[i+1][j+1]) ) {
                        continue;
                    }
                    if ( (i > 0) && (j < N-1) && (temp[i][j] > temp[i-1][j+1]) ) {
                        continue;
                    }
                }
                matr[i][j] = 0;
            }
        }
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                put(matr[i][j]);
            }
        }
		return false;
    }

	private boolean matrix75() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        double[][] temp = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
                temp[i][j] = matr[i][j];
            }
        }
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                if ( (i > 0) || (j > 0) ) {
                    if ( (i > 0) && (temp[i][j] < temp[i-1][j]) ) {
                        continue;
                    }
                    if ( (j > 0) && (temp[i][j] < temp[i][j-1]) ) {
                        continue;
                    }
                    if ( (i > 0) && (j > 0) && (temp[i][j] < temp[i-1][j-1]) ) {
                        continue;
                    }
                    if ( (j > 0) && (i < M-1) && (temp[i][j] < temp[i+1][j-1]) ) {
                        continue;
                    }
                }
                if ( (i < M-1) || (j < N-1) ) {
                    if ( (i < M-1) && (temp[i][j] < temp[i+1][j]) ) {
                        continue;
                    }
                    if ( (j < N-1) && (temp[i][j] < temp[i][j+1]) ) {
                        continue;
                    }
                    if ( (i < M-1) && (j < N-1) && (temp[i][j] < temp[i+1][j+1]) ) {
                        continue;
                    }
                    if ( (i > 0) && (j < N-1) && (temp[i][j] < temp[i-1][j+1]) ) {
                        continue;
                    }
                }
                matr[i][j] *= -1;
            }
        }
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                put(matr[i][j]);
            }
        }
        return false;
    }

	private boolean matrix76() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        double temp;
        for (int k = 0; k < M-1; k++) {
            for (int i = 1; i < M-k; i++) {
                if (matr[i][0] < matr[i-1][0]) {
                    for (int j = 0; j < N; j++) {
                        temp = matr[i][j];
                        matr[i][j] = matr[i-1][j];
                        matr[i-1][j] = temp;
                    }
                }
            }
        }
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                put(matr[i][j]);
            }
        }
        return false;
    }

	private boolean matrix77() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        double temp;
        for (int k = 0; k < N-1; k++) {
            for (int j = 1; j < N-k; j++) {
                if (matr[M-1][j] > matr[M-1][j-1]) {
                    for (int i = 0; i < M; i++) {
                        temp = matr[i][j];
                        matr[i][j] = matr[i][j-1];
                        matr[i][j-1] = temp;
                    }
                }
            }
        }
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                put(matr[i][j]);
            }
        }
        return false;
    }

	private boolean matrix78() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        double[] minimal = new double [M];
        for (int i = 0; i < M; i++) {
            minimal[i] = matr[i][0];
            for (int j = 1; j < N; j++) {
                if (matr[i][j] < minimal[i]) {
                    minimal[i] = matr[i][j];
                }
            }
        }
        double temp;
        for (int k = 0; k < M-1; k++) {
            for (int i = 1; i < M-k; i++) {
                if (minimal[i-1] < minimal[i]) {
                    for (int j = 0; j < N; j++) {
                        temp = matr[i][j];
                        matr[i][j] = matr[i-1][j];
                        matr[i-1][j] = temp;
                    }
                    temp = minimal[i-1];
                    minimal[i-1] = minimal[i];
                    minimal[i] = temp;
                }
            }
        }
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                put(matr[i][j]);
            }
        }
        return false;
    }

	private boolean matrix79() {
        int M = getInt();
        int N = getInt();
        double[][] matr = new double [M][N];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                matr[i][j] = getDouble();
            }
        }
        double[] maximal = new double [N];
        for (int j = 0; j < N; j++) {
            maximal[j] = matr[0][j];
            for (int i = 1; i < M; i++) {
                if (matr[i][j] > maximal[j]) {
                    maximal[j] = matr[i][j];
                }
            }
        }
        double temp;
        for (int k = 0; k < N-1; k++) {
            for (int j = 1; j < N-k; j++) {
                if (maximal[j-1] > maximal[j]) {
                    for (int i = 0; i < M; i++) {
                        temp = matr[i][j];
                        matr[i][j] = matr[i][j-1];
                        matr[i][j-1] = temp;
                    }
                    temp = maximal[j];
                    maximal[j] = maximal[j-1];
                    maximal[j-1] = temp;
                }
            }
        }
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < N; j++) {
                put(matr[i][j]);
            }
        }
        return false;
    }

	private boolean matrix80() {
        int M = getInt();
        double[][] A = new double [M][M];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < M; j++) {
                A[i][j] = getDouble();
            }
        }
        double sum = 0;
        for (int row = 0, col = 0; row < M; row++, col++) {
            sum += A[row][col];
        }
        put(sum);
        return false;
    }

	private boolean matrix81() {
        int M = getInt();
        double[][] A = new double [M][M];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < M; j++) {
                A[i][j] = getDouble();
            }
        }
        double sum = 0;
        int count = 0;
        for (int row = 0, col = M-1; row < M; row++, col--) {
            sum += A[row][col];
            count++;
        }
        double amean = sum / count;
        put(amean);
        return false;
    }

	private boolean matrix82() {
        int M = getInt();
        double[][] A = new double [M][M];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < M; j++) {
                A[i][j] = getDouble();
            }
        }
        double sum;
        for (int col = M-1; col >= 0; col--) {
            sum = 0;
            for (int i = 0, j = col; j < M; i++, j++) {
                sum += A[i][j];
            }
            put(sum);
        }
        for (int row = 1; row < M; row++) {
            sum = 0;
            for (int i = row, j = 0; i < M; i++, j++) {
                sum += A[i][j];
            }
            put(sum);
        }
        return false;
    }

	private boolean matrix83() {
        int M = getInt();
        double[][] A = new double [M][M];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < M; j++) {
                A[i][j] = getDouble();
            }
        }
        double sum;
        for (int col = 0; col < M; col++) {
            sum = 0;
            for (int i = 0, j = col; j >= 0; i++, j--) {
                sum += A[i][j];
            }
            put(sum);
        }
        for (int row = 1; row < M; row++) {
            sum = 0;
            for (int i = row, j = M-1; i < M; i++, j--) {
                sum += A[i][j];
            }
            put(sum);
        }
        return false;
    }

	private boolean matrix84() {
        int M = getInt();
        double[][] A = new double [M][M];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < M; j++) {
                A[i][j] = getDouble();
            }
        }
        double sum, amean;
        int count;
        for (int col = M-1; col >= 0; col--) {
            sum = 0;
            count = 0;
            for (int i = 0, j = col; j < M; i++, j++) {
                sum += A[i][j];
                count++;
            }
            amean = sum / count;
            put(amean);
        }
        for (int row = 1; row < M; row++) {
            sum = 0;
            count = 0;
            for (int i = row, j = 0; i < M; i++, j++) {
                sum += A[i][j];
                count++;
            }
            amean = sum / count;
            put(amean);
        }
        return false;
    }

	private boolean matrix85() {
        int M = getInt();
        double[][] A = new double [M][M];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < M; j++) {
                A[i][j] = getDouble();
            }
        }
        double sum, amean;
        int count;
        for (int col = 0; col < M; col++) {
            sum = 0;
            count = 0;
            for (int i = 0, j = col; j >= 0; i++, j--) {
                sum += A[i][j];
                count++;
            }
            amean = sum / count;
            put(amean);
        }
        for (int row = 1; row < M; row++) {
            sum = 0;
            count = 0;
            for (int i = row, j = M-1; i < M; i++, j--) {
                sum += A[i][j];
                count++;
            }
            amean = sum / count;
            put(amean);
        }
        return false;
    }

	private boolean matrix86() {
        int M = getInt();
        double[][] A = new double [M][M];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < M; j++) {
                A[i][j] = getDouble();
            }
        }
        double minimal;
        for (int col = M-1; col >= 0; col--) {
            minimal = A[0][col];
            for (int i = 0, j = col; j < M; i++, j++) {
                if (A[i][j] < minimal) {
                    minimal = A[i][j];
                }
            }
            put(minimal);
        }
        for (int row = 1; row < M; row++) {
            minimal = A[row][0];
            for (int i = row, j = 0; i < M; i++, j++) {
                if (A[i][j] < minimal) {
                    minimal = A[i][j];
                }
            }
            put(minimal);
        }
        return false;
    }

	private boolean matrix87() {
        int M = getInt();
        double[][] A = new double [M][M];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < M; j++) {
                A[i][j] = getDouble();
            }
        }
        double maximal;
        for (int col = 0; col < M; col++) {
            maximal = A[0][col];
            for (int i = 0, j = col; j >= 0; i++, j--) {
                if (A[i][j] > maximal) {
                    maximal = A[i][j];
                }
            }
            put(maximal);
        }
        for (int row = 1; row < M; row++) {
            maximal = A[row][M-1];
            for (int i = row, j = M-1; i < M; i++, j--) {
                if (A[i][j] > maximal) {
                    maximal = A[i][j];
                }
            }
            put(maximal);
        }
        return false;
    }

	private boolean matrix88() {
        int M = getInt();
        double[][] matr = new double [M][M];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < M; j++) {
                matr[i][j] = getDouble();
            }
        }
        for (int row = 1; row < M; row++) {
            for (int i = row, j = 0; i < M; i++, j++) {
                matr[i][j] = 0;
            }
        }
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < M; j++) {
                put(matr[i][j]);
            }
        }
        return false;
    }

	private boolean matrix89() {
        int M = getInt();
        double[][] matr = new double [M][M];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < M; j++) {
                matr[i][j] = getDouble();
            }
        }
        for (int col = M-2; col >= 0; col--) {
            for (int i = 0, j = col; j >= 0; i++, j--) {
                matr[i][j] = 0;
            }
        }
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < M; j++) {
                put(matr[i][j]);
            }
        }
        return false;
    }

	private boolean matrix90() {
        int M = getInt();
        double[][] matr = new double [M][M];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < M; j++) {
                matr[i][j] = getDouble();
            }
        }
        for (int row = 0; row < M; row++) {
            for (int i = row, j = M-1; i < M; i++, j--) {
                matr[i][j] = 0;
            }
        }
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < M; j++) {
                put(matr[i][j]);
            }
        }
        return false;
    }

	private boolean matrix91() {
        int M = getInt();
        double[][] matr = new double [M][M];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < M; j++) {
                matr[i][j] = getDouble();
            }
        }
        for (int col = 0; col < M; col++) {
            for (int i = 0, j = col; j < M; i++, j++) {
                matr[i][j] = 0;
            }
        }
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < M; j++) {
                put(matr[i][j]);
            }
        }
        return false;
    }

	private boolean matrix92() {
        int M = getInt();
        double[][] matr = new double [M][M];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < M; j++) {
                matr[i][j] = getDouble();
            }
        }
        for (int row = 0; row < M/2; row++) {
            for (int col = row+1; col < M-row-1; col++) {
                matr[row][col] = 0;
            }
        }
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < M; j++) {
                put(matr[i][j]);
            }
        }
        return false;
    }

	private boolean matrix93() {
        int M = getInt();
        double[][] matr = new double [M][M];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < M; j++) {
                matr[i][j] = getDouble();
            }
        }
        for (int col = M-1; col > M/2; col--) {
            for (int row = M-col; row < col; row++) {
                matr[row][col] = 0;
            }
        }
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < M; j++) {
                put(matr[i][j]);
            }
        }
        return false;
    }

	private boolean matrix94() {
        int M = getInt();
        double[][] matr = new double [M][M];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < M; j++) {
                matr[i][j] = getDouble();
            }
        }
        for (int col = 0; col <= M/2; col++) {
            for (int row = col; row < M-col; row++) {
                matr[row][col] = 0;
            }
        }
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < M; j++) {
                put(matr[i][j]);
            }
        }
        return false;
    }

	private boolean matrix95() {
        int M = getInt();
        double[][] matr = new double [M][M];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < M; j++) {
                matr[i][j] = getDouble();
            }
        }
        for (int row = M-1; row >= M/2; row--) {
            for (int col = M - row - 1; col <= row; col++) {
                matr[row][col] = 0;
            }
        }
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < M; j++) {
                put(matr[i][j]);
            }
        }
        return false;
    }

	private boolean matrix96() {
        int M = getInt();
        double[][] A = new double [M][M];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < M; j++) {
                A[i][j] = getDouble();
            }
        }
        double temp;
        for (int col = 1; col < M; col++) {
            for (int i = 0, j = col; j < M; i++, j++) {
                temp = A[i][j];
                A[i][j] = A[j][i];
                A[j][i] = temp;
            }
        }
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < M; j++) {
                put(A[i][j]);
            }
        }
        return false;
    }

	private boolean matrix97() {
        int M = getInt();
        double[][] A = new double [M][M];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < M; j++) {
                A[i][j] = getDouble();
            }
        }
        double temp;
        for (int col = M-2; col >= 0; col--) {
            for (int i = 0, j = col; j >= 0; i++, j--) {
                temp = A[i][j];
                A[i][j] = A[M-1-j][M-1-i];
                A[M-1-j][M-1-i] = temp;
            }
        }
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < M; j++) {
                put(A[i][j]);
            }
        }
        return false;
    }

	private boolean matrix98() {
        int M = getInt();
        double[][] A = new double [M][M];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < M; j++) {
                A[i][j] = getDouble();
            }
        }
        double temp;
        int limit;
        for (int i = M/2+M%2-1; i >= 0; i--) {
            limit = i == M-1-i ? M/2 : M;
            for (int j = 0; j < limit; j++) {
                temp = A[i][j];
                A[i][j] = A[M-1-i][M-1-j];
                A[M-1-i][M-1-j] = temp;
            }
        }
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < M; j++) {
                put(A[i][j]);
            }
        }
        return false;
    }

	private boolean matrix99() {
        int M = getInt();
        double[][] A = new double [M][M];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < M; j++) {
                A[i][j] = getDouble();
            }
        }
        double temp;
        for (int i = 0; i < M/2; i++) {
            for (int j = i; j < M-1-i; j++) {
                temp = A[i][j];
                A[i][j] = A[j][M-1-i];
                A[j][M-1-i] = A[M-1-i][M-1-j];
                A[M-1-i][M-1-j] = A[M-1-j][i];
                A[M-1-j][i] = temp;
            }
        }
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < M; j++) {
                put(A[i][j]);
            }
        }
        return false;
    }

	private boolean matrix100() {
        int M = getInt();
        double[][] A = new double [M][M];
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < M; j++) {
                A[i][j] = getDouble();
            }
        }
        double temp;
        for (int i = 0; i < M/2; i++) {
            for (int j = i; j < M-1-i; j++) {
                temp = A[i][j];
                A[i][j] = A[M-1-j][i];
                A[M-1-j][i] = A[M-1-i][M-1-j];
                A[M-1-i][M-1-j] = A[j][M-1-i];
                A[j][M-1-i] = temp;
            }
        }
        for (int i = 0; i < M; i++) {
            for (int j = 0; j < M; j++) {
                put(A[i][j]);
            }
        }
        return false;
    }
}