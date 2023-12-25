#include <iostream>
using namespace std;

int main()
{
	//Task("While15");
	double deposite = 1000;
	double P;
	cin >> P;
	int month = 0;
	while (deposite <= 1100)
	{
		deposite += deposite * P / 100;
		month++;
	}
	cout << month << deposite;
	
	return 0;
}
