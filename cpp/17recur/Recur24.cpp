#include <iostream>
using namespace std;

bool Logic(const char* S, int& index);
int main()
{
	//Task("Recur24");
	char str[100];
	cin >> str;
	int index = 0;
	bool result = Logic(str, index);
	cout << result;
	return 0;
}

bool Logic(const char* S, int& index)
{
	if ( (S[index] == 'T') || (S[index] == 'F') )
		return S[index] == 'T';
	bool And = S[index] == 'A';
	bool Not = S[index] == 'N';
	bool Or = S[index] == 'O';
	index += 4;
	index -= (int)Or;
	bool first = Logic(S, index);
	if (Not) first = !first;
	else
		while (S[index+1] != ')')
		{
			index += 2;
			bool second = Logic(S, index);
			first = And ? first && second : first || second;
		}
	index++;
	return first;
}