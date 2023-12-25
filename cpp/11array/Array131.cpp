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
	//Task("Array131");
	
	int N;
	cin >> N;

	Point* A = new Point [N];
	for (int i = 0; i < N; i++)
		cin >> A[i].x >> A[i].y;

	Point B;
	cin >> B.x >> B.y;

	int index = 0;
	double length, minLength = sqrt( pow(A[index].x - B.x, 2) + pow(A[index].y - B.y, 2) );
	for (int i = 1; i < N; i++)
	{
		length = sqrt( pow(A[i].x - B.x, 2) + pow(A[i].y - B.y, 2) );
		
		if (length < minLength)
		{
			minLength = length;
			index = i;
		}
	}

	cout << A[index].x << A[index].y;

	delete [] A;
	A = 0;
	
	return 0;
}
