#include <stdio.h>

int main() {
	char fileName[50];
	int count = 0;
	FILE* file = NULL;
	for (int i = 0; i < 4; ++i) {
		gets(fileName);
		file = fopen(fileName, "rb");
		if (file != NULL) {
			++count;
			fclose(file);
		}
	}
	printf("%d", count);
	return 0;
}