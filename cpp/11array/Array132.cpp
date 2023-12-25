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
	//Task("Array132");
	
	int N;
	cin >> N;

	Point* A = new Point [N];
	int index = -1;
	for (int i = 0; i < N; i++)
	{
		cin >> A[i].x >> A[i].y;

		if ( (index < 0) && (A[i].x < 0) && (A[i].y > 0) )
			index = i;
	}

	Point point;
	point.x = 0;
	point.y = 0;

	if (index >= 0)
	{
		double length, maxLength = sqrt( pow(A[index].x, 2) + pow(A[index].y, 2) );
		for (int i = index+1; i < N; i++)
		{
			if ( (A[i].x >= 0) || (A[i].y <= 0) )
				continue;

			length = sqrt( pow(A[i].x, 2) + pow(A[i].y, 2) );

			if (length > maxLength)
			{
				maxLength = length;
				index = i;
			}
		}
		point.x = A[index].x;
		point.y = A[index].y;
	}

	cout << point.x << point.y;

	delete [] A;
	A = 0;
	
	return 0;
}
