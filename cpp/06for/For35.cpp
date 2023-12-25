#include <iostream>
using namespace std;

int main()
{
	//Task("For35");
	int N;
	cin >> N;
	int Ak, A1 = 1, A2 = 2, A3 = 3;
	cout << A1 << A2 << A3;
	for (int K = 4; K <= N; K++)
	{
		Ak = A3 + A2 - 2 * A1;
		cout << Ak;
		A1 = A2;
		A2 = A3;
		A3 = Ak;
	}
	
	return 0;
}
