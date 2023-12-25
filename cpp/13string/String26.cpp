#include <iostream>
using namespace std;

int main()
{
	//Task("String26");
	int N;
	cin >> N;
	cin.ignore();
	char S[1000];
	cin.getline(S, 1000);

	int len = 0;
	while (S[len] != '\0')
		len++;

	char* newStr = new char [N+1];
	newStr[N] = '\0';
	
	int i = N-1, j = len-1;
	while (i >= 0)
	{
		if (j < 0)
			break;
		newStr[i--] = S[j--];
	}
	while (i >= 0)
		newStr[i--] = '.';

	cout << newStr;

	delete [] newStr;
	newStr = NULL;
	
	return 0;
}
