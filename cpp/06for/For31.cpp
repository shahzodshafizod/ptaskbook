#include <iostream>
using namespace std;

int main()
{
	//Task("For31");
	int N;
	cin >> N;
	double A = 2;
	for (int K = 1; K <= N; K++)
	{
		A = 2 + 1 / A;
		cout << A;
	}
	
	return 0;
}
