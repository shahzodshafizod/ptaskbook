#include <iostream>
using namespace std;

int main()
{
	//Task("For3");
	int A, B;
	cin >> A >> B;
	int N = 0;
	for (int i = B-1; i > A; i--)
	{
		cout << i;
		N++;
	}
	cout << N;
	
	return 0;
}
