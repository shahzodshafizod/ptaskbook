#include <iostream>
using namespace std;
bool Logic(const char* S, int& index);
int main()
{
	//Task("Recur23");
	char str[100];
	cin >> str;
	int index = 0;
	bool logic = Logic(str, index);
	cout << logic;
	return 0;
}

bool Logic(const char* S, int& index)
{
	if ( (S[index] == 'T') || (S[index] == 'F') )
		return S[index] == 'T';
	bool And = S[index] == 'A';
	index += 3;
	index += (int)And;
	bool first = Logic(S, index);
	while (S[index+1] != ')')
	{
		index += 2;
		bool second = Logic(S, index);
		first = And ? first && second : first || second;
	}
	index++;
	return first;
}