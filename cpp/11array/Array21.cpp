#include <iostream>
using namespace std;

int main()
{
	//Task("Array21");
	int N;
	cin >> N;
	double* mass = new double [N];
	for (int i = 0; i < N; i++)
		cin >> mass[i];
	int K, L;
	cin >> K >> L;
	double sum = 0;
	int count = 0;
	for (int i = K-1; i < L; i++)
	{
		sum += mass[i];
		count++;
	}
	if (count == 0)
		cout << 0;
	else
		cout << sum / count;
	delete [] mass;
	mass = 0;
	
	return 0;
}
