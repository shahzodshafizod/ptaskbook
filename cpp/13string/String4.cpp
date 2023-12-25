#include <iostream>
using namespace std;

int main()
{
	//Task("String4");
	int N;
	cin >> N;

	int code = int('A');
	for (int i = 1; i <= N; i++)
	{
		cout << char(code);
		code++;
	}
	
	return 0;
}
