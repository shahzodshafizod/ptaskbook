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
	//Task("Array135");
	
	int N1;
	cin >> N1;

	Point* A = new Point [N1];
	for (int i = 0; i < N1; i++)
		cin >> A[i].x >> A[i].y;

	int N2;
	cin >> N2;
	Point* B = new Point [N2];
	for (int i = 0; i < N2; i++)
		cin >> B[i].x >> B[i].y;

	double length, minLength = sqrt( pow(A[0].x - B[0].x, 2) + pow(A[0].y - B[0].y, 2) );
	int indexA = 0, indexB = 0;
	for (int i = 0; i < N1; i++)
	{
		for (int j = 0; j < N2; j++)
		{
			length = sqrt( pow(A[i].x - B[j].x, 2) + pow(A[i].y - B[j].y, 2) );
			
			if (length < minLength)
			{
				minLength = length;
				indexA = i;
				indexB = j;
			}
		}
	}

	cout << minLength;
	cout << A[indexA].x << A[indexA].y;
	cout << B[indexB].x << B[indexB].y;

	delete [] A;
	delete [] B;
	A = B = 0;
	
	return 0;
}
