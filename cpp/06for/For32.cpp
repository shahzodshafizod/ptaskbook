#include <iostream>
using namespace std;

int main()
{
	//Task("For32");
	int N;
	cin >> N;
	double A = 1;
	for (int K = 1; K <= N; K++)
	{
		A = (A + 1) / K;
		cout << A;
	}
	
	return 0;
}
