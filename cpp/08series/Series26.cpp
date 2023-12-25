#include <iostream>
using namespace std;

int main()
{
	//Task("Series26");
	int K, N;
	cin >> K >> N;
	double number, degree;
	for (int i = 1; i <= N; i++)
	{
		cin >> number;
		degree = 1;
		for (int j = 1; j <= K; j++)
			degree *= number;
		cout << degree;
	}
	
	return 0;
}
