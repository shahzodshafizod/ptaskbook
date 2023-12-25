#include <iostream>
#include <cmath>
using namespace std;

struct TPoint
{
	double X, Y;
};

double Leng(TPoint A, TPoint B);
double PerimN(TPoint P[], int N);

int main()
{
	//Task("Param69");
	int N;
	TPoint* mass = NULL;
	for (int i = 0; i < 3; i++)
	{
		cin >> N;
		mass = new TPoint [N];
		for (int j = 0; j < N; j++)
			cin >> mass[j].X >> mass[j].Y;
		
		cout << PerimN(mass, N);
		
		delete [] mass;
		mass = NULL;
	}
	
	return 0;
}

double Leng(TPoint A, TPoint B)
{
	return sqrt( pow(A.X - B.X, 2) + pow(A.Y - B.Y, 2) );
}

double PerimN(TPoint P[], int N)
{
	double result = Leng(P[0], P[N-1]);
	
	for (int i = 1; i < N; i++)
		result += Leng(P[i-1], P[i]);
	
	return result;
}