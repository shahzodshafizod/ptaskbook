#include <iostream>
using namespace std;

bool IsSquare(int K);

int main()
{
	//Task("Proc25");
	int number, squares = 0;
	for (int i = 1; i <= 10; i++)
	{
		cin >> number;
		squares += (int)IsSquare(number);
	}
	cout << squares;
	
	return 0;
}

bool IsSquare(int K)
{
	int i = 1;
	for (; i*i < K; i++);
	return i*i == K;
}