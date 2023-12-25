#include <iostream>
using namespace std;

int main()
{
	//Task("String27");
	unsigned int N1, N2;
	cin >> N1 >> N2;
	cin.ignore();

	char S1[1000], S2[1000];
	cin.getline(S1, 1000);
	cin.getline(S2, 1000);

	int newLen = N1+N2;
	char* newStr = new char [newLen+1];
	
	int index = 0;
	for (; index < N1; index++)
		newStr[index] = S1[index];
	
	int S2Len = 0;
	while (S2[S2Len] != '\0')
		S2Len++;

	int S2index = S2Len-N2;
	while (S2index < S2Len)
		newStr[index++] = S2[S2index++];
	
	newStr[newLen] = '\0';

	cout << newStr;

	delete [] newStr;
	newStr = NULL;
	
	return 0;
}
