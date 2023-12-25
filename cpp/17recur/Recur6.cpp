#include <iostream>
using namespace std;

int Combin1(int N, int K);
int counts;

int main()
{
	//Task("Recur6");
	int N;
	cin >> N;
	int K;
	for (int i = 1; i < 6; i++)
	{
		cin >> K;
		counts = 0;
		cout << Combin1(N, K);
		cout << counts;
	}
	return 0;
}

int Combin1(int N, int K)
{
	counts++;
	if ( (N == K) || (K == 0) )
		return 1;

	return Combin1(N-1, K) + Combin1(N-1, K-1);
}