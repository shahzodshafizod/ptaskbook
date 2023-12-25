#include <iostream>
using namespace std;

int main()
{
	//Task("String5");
	int N;
	cin >> N;

	int code = int('z');
	for (int i = 1; i <= N; i++)
	{
		cout << char(code);
		code--;
	}
	
	return 0;
}
