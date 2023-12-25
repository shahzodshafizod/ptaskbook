#include <iostream>
#include <cmath>
using namespace std;

struct Point
{
	double x;
	double y;
};

int main()
{
	//Task("Array136");
	
	int N;
	cin >> N;

	Point* A = new Point [N];
	for (int i = 0; i < N; i++)
		cin >> A[i].x >> A[i].y;

	double sum, minSum = 0;
	int index = 0;
	for (int i = 0; i < N; i++)
	{
		sum = 0;
		for (int j = 0; j < N; j++)
			sum += sqrt( pow(A[i].x - A[j].x, 2) + pow(A[i].y - A[j].y, 2) );
		
		if (minSum == 0)
		{
			minSum = sum;
			index = i;
		}
		else if (sum < minSum)
		{
			minSum = sum;
			index = i;
		}
	}

	cout << A[index].x << A[index].y;
	cout << minSum;

	delete [] A;
	A = 0;
	
	return 0;
}
