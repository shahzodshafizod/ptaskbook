#include <iostream>
using namespace std;

int main()
{
	//Task("For2");
	int A, B;
	cin >> A >> B;
	int N = 0;
	for (int i = A; i <= B; i++)
	{
		cout << i;
		N++;
	}
	cout << N;
	
	return 0;
}
