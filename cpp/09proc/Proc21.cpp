#include <iostream>
using namespace std;

int SumRange(int A, int B);

int main()
{
	//Task("Proc21");
	int A, B, C;
	cin >> A >> B >> C;

	cout << SumRange(A, B);
	cout << SumRange(B, C);
	
	return 0;
}

int SumRange(int A, int B)
{
	int sum = 0;
	for (int i = A; i <= B; i++)
		sum += i;
	
	return sum;
}