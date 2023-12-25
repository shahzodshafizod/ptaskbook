#include <iostream>
using namespace std;

int main()
{
	//Task("String9");
	int N;
	cin >> N;
	cin.ignore();
	char C1, C2;
	cin >> C1;
	cin.ignore();
	cin >> C2;

	char* str = new char [N+1];
	for (int i = 0; i < N; i += 2)
	{
		str[i] = C1;
		str[i+1] = C2;
	}
	str[N] = '\0';

	cout << str;

	delete [] str;
	str = NULL;
	
	return 0;
}
