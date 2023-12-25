#include <iostream>
using namespace std;

char* FillStr(const char* S, int N);

int main()
{
	//Task("Param31");
	const int stringSize = 100;
	char str[stringSize+1];
	char* newStr = NULL;
	int N;
	cin >> N;
	cin.ignore();
	for (int i = 0; i < 5; i++)
	{
		cin.getline(str, stringSize+1);
		newStr = FillStr(str, N);
		cout << newStr;
		
		delete [] newStr;
		newStr = NULL;
	}
	
	return 0;
}

char* FillStr(const char* S, int N)
{
	char* newS = new char [N+1];
	int len = strlen(S);
	int index = 0;
	for (int i = 0; i < N; i++)
	{
		newS[i] = S[index];
		if (index+1 < len)
			index++;
		else
			index = 0;
	}
	newS[N] = '\0';
	return newS;
}