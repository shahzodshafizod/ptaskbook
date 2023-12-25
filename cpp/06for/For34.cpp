#include <iostream>
using namespace std;

int main()
{
	//Task("For34");
	int N;
	cin >> N;
	double Ak, A1 = 1, A2 = 2;
	cout << A1 << A2;
	for (int K = 3; K <= N; K++)
	{
		Ak = (A1 + 2 * A2) / 3;
		cout << Ak;
		A1 = A2;
		A2 = Ak;
	}
	
	return 0;
}
