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
	//Task("Array134");
	
	int N;
	cin >> N;

	Point* A = new Point[N];
	for (int i = 0; i < N; i++)
		cin >> A[i].x >> A[i].y;

	double length, maxLength = sqrt( pow(A[0].x - A[1].x, 2) + pow(A[0].y - A[1].y, 2) );
	int index1 = 0, index2 = 1;
	for (int i = 0; i < N; i++)
	{
		for (int j = i+1; j < N; j++)
		{
			length = sqrt( pow(A[i].x - A[j].x, 2) + pow(A[i].y - A[j].y, 2) );

			if (length > maxLength)
			{
				maxLength = length;
				index1 = i;
				index2 = j;
			}
		}
	}

	cout << A[index1].x << A[index1].y;
	cout << A[index2].x << A[index2].y;
	cout << maxLength;

	delete [] A;
	A = 0;
	
	return 0;
}
