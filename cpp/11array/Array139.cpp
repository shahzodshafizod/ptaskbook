#include <iostream>
using namespace std;

struct Point
{
	int x;
	int y;
};

void Swap(Point& a, Point& b);

int main()
{
	//Task("Array139");
	
	int N;
	cin >> N;

	Point* A = new Point [N];
	for (int i = 0; i < N; i++)
		cin >> A[i].x >> A[i].y;

	bool condition;
	for (int i = 0; i < N-1; i++)
	{
		for (int j = 1; j < N-i; j++)
		{
			condition = (A[j-1].x < A[j].x) || (A[j-1].x == A[j].x) && (A[j-1].y < A[j].y);
			if ( !condition )
				Swap(A[j-1], A[j]);
		}
	}

	for (int i = 0; i < N; i++)
		cout << A[i].x << A[i].y;

	delete [] A;
	A = 0;
	
	return 0;
}

void Swap(Point& a, Point& b)
{
	Point c;
	c.x = a.x;
	c.y = a.y;

	a.x = b.x;
	a.y = b.y;

	b.x = c.x;
	b.y = c.y;
}