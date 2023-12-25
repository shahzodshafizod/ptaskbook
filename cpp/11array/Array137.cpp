#include <iostream>
#include <cmath>
using namespace std;

struct Point
{
	double x;
	double y;
};

double GetPerim(Point mass[], int index1, int index2, int index3);

int main()
{
	//Task("Array137");

	int N;
	cin >> N;

	Point* A = new Point [N];
	for (int i = 0; i < N; i++)
		cin >> A[i].x >> A[i].y;

	int index1 = 0, index2 = 1, index3 = 2;
	double P, maxP = GetPerim(A, index1, index2, index3);
	for (int i = 0; i < N; i++)
	{
		for (int j = i+1; j < N; j++)
		{
			for (int k = j+1; k < N; k++)
			{
				P = GetPerim(A, i, j, k);
				
				if (P > maxP)
				{
					maxP = P;
					index1 = i;
					index2 = j;
					index3 = k;
				}
			}
		}
	}

	cout << maxP;
	cout << A[index1].x << A[index1].y;
	cout << A[index2].x << A[index2].y;
	cout << A[index3].x << A[index3].y;

	delete [] A;
	A = 0;
	
	return 0;
}

double GetPerim(Point mass[], int index1, int index2, int index3)
{
	double a = sqrt( pow(mass[index1].x - mass[index2].x, 2) + pow(mass[index1].y - mass[index2].y, 2) );
	double b = sqrt( pow(mass[index1].x - mass[index3].x, 2) + pow(mass[index1].y - mass[index3].y, 2) );
	double c = sqrt( pow(mass[index2].x - mass[index3].x, 2) + pow(mass[index2].y - mass[index3].y, 2) );

	return a + b + c;
}