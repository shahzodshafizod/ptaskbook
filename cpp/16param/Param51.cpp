#include <iostream>
#include <fstream>
#include <iomanip>
using namespace std;

void AddLineNumbers(const char* S, int N, int K, int L);

int main()
{
	//Task("Param51");
	char fileName[100];
	cin.getline(fileName, 100);
	int N, K, L;
	cin >> N >> K >> L;
	AddLineNumbers(fileName, N, K, L);
	
	return 0;
}

void AddLineNumbers(const char* S, int N, int K, int L)
{
	ifstream In(S);
	if (In.is_open())
	{
		char tempFileName[] = "tempFileName.txt";
		ofstream Out(tempFileName);
		const int stringSize = 100;
		char str[stringSize+1];
		int spaces;
		while (In.peek() != -1)
		{
			In.getline(str, stringSize+1);			
			spaces = strlen(str) == 0 ? 0 : L;
			Out << setw(K) << N++;
			for (int i = 0; i < spaces; i++)
				Out << ' ';
			Out << str << endl;
		}
		In.close();
		Out.close();
		
		remove(S);
		rename(tempFileName, S);
	}
}