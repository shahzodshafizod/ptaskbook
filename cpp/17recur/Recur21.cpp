#include <iostream>
using namespace std;

bool LogicResult(const char* S, int& index);

int main()
{
	//Task("Recur21");
	char str[100];
	cin >> str;
	int index = 0;
	bool result = LogicResult(str, index);
	cout << result;
	return 0;
}

bool LogicResult(const char* S, int& index)
{
	if ((S[index] == 'T') || (S[index] == 'F'))
		return S[index] == 'T';
	bool And = S[index] == 'A';
	index += 3;
	index += (int)And;
	bool first = LogicResult(S, index);
	index += 2;
	bool second = LogicResult(S, index);
	index++;
	return ( And ? first && second : first || second );
}