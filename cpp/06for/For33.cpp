#include <iostream>
using namespace std;

int main()
{
	//Task("For33");
	int N;
	cin >> N;
	int F1, F2, Fk;
	F1 = F2 = 1;
	cout << F1 << F2;
	for (int K = 3; K <= N; K++)
	{
		Fk = F1 + F2;
		cout << Fk;
		F1 = F2;
		F2 = Fk;
	}
	
	return 0;
}
