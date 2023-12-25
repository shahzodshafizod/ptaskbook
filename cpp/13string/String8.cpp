#include <iostream>
using namespace std;

int main()
{
	//Task("String8");
	int N;
	cin >> N;
	cin.ignore();
	char C;
	cin >> C;
	
	char* str = new char [N+1];
	for (int i = 0; i < N; i++)
		str[i] = C;
	str[N] = '\0';

	cout << str;

	delete [] str;
	str = NULL;
	
	return 0;
}
