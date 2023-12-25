#include <iostream>
using namespace std;

void InvDigits(int& K);

int main()
{
    //Task("Proc7");
	int N;
	for (int i = 1; i < 6; i++)
	{
		cin >> N;
		InvDigits(N);
		cout << N;
	}
	
	return 0;
}

void InvDigits(int& K)
{
	int temp = 0;
	while (K > 0)
	{
		temp = temp * 10 + K % 10;
		K /= 10;
	}
	K = temp;
}