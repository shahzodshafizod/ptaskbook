#include <stdio.h>

int main() {
	int monthNo, days = 0;
	scanf("%d", &monthNo);
	switch (monthNo) {
		case 1: case 3: case 5: case 7: case 8: case 10: case 12:
			days = 31;
			break;
		case 4: case 6: case 9: case 11:
			days = 30;
			break;
		case 2:
			days = 28;
			break;
	}
	printf("%d", days);
	return 0;
}