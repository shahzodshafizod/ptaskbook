#include <iostream>
using namespace std;

int Combin2(int N, int K);
int matr[20][20];

int main()
{
	//Task("Recur7");
	int N;
	cin >> N;
	int K;
	for (int i = 1; i < 6; i++)
	{
		cin >> K;
		
		for (int k = 0; k < 20; k++)
			for (int l = 0; l < 20; ++l)
				matr[k][l] = 0;

		cout << Combin2(N, K);
	}
	return 0;
}

int Combin2(int N, int K)
{
	if ( (N == K) || (K == 0) )
		return 1;

	matr[N-1][K-1] = ( matr[N-2][K-1] > 0 ? matr[N-2][K-1] : Combin2(N-1, K) ) +
					 ( matr[N-2][K-2] > 0 ? matr[N-2][K-2] : Combin2(N-1, K-1) );

	return matr[N-1][K-1];
}