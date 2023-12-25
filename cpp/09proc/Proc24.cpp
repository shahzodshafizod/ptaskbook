#include <iostream>
using namespace std;

bool Even(int K);

int main()
{
	//Task("Proc24");
	int number, evens = 0;
	for (int i = 1; i <= 10; i++)
	{
		cin >> number;
		evens += (int)Even(number);
	}
	cout << evens;
	
	return 0;
}

bool Even(int K)
{
	return ( K % 2 == 0 );
}