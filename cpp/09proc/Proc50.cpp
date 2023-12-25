#include <iostream>
using namespace std;

void TimeToHMS(int T, int& H, int& M, int& S);
int main()
{
	//Task("Proc50");
	int T, H, M, S;
	for (int i = 1; i < 6; i++)
	{
		cin >> T;
		TimeToHMS(T, H, M, S);
		cout << H << M << S;
	}
	
	return 0;
}

void TimeToHMS(int T, int& H, int& M, int& S)
{
	H = T / 3600;
	T -= H * 3600;

	M = T / 60;
	T -= M * 60;

	S = T;
}