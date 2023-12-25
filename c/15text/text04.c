#include <stdio.h>
#include <string.h>

int main() {
	char fileName[50];
	gets(fileName);
	FILE* file = fopen(fileName, "r");
	char str[100];
	int strings = 0, chars = 0;
	while (fgets(str, 100, file) != NULL) {
		++strings;
		chars += strlen(str)-1;
	}
	fclose(file);
	printf("chars: %d\n", chars);
	printf("strings: %d\n", strings);
	return 0;
}