#include <iostream>
#include <cmath>
using namespace std;

struct TPoint
{
	double X, Y;
};

struct TTriangle
{
	TPoint A, B, C;
};

double AreaN(TPoint P[], int N);
double Area(TTriangle T);
double Leng(TPoint A, TPoint B);

int main()
{
	//Task("Param70");
	int N;
	TPoint* mass = NULL;
	for (int i = 1; i <= 3; i++)
	{
		cin >> N;
		mass = new TPoint [N];
		for (int j = 0; j < N; j++)
			cin >> mass[j].X >> mass[j].Y;
		
		cout << AreaN(mass, N);

		delete [] mass;
		mass = NULL;
	}
	
	return 0;
}

//double AreaN(TPoint* P, int N)
double AreaN(TPoint P[], int N)
{
	TTriangle sekunja;
	sekunja.A = P[0];
	double result = 0;
	for (int i = 2; i < N; i++)
	{
		sekunja.B = P[i-1];
		sekunja.C = P[i];
		result += Area(sekunja);
	}

	return result;
}

double Area(TTriangle T)
{
	double a = Leng(T.A, T.B);
	double b = Leng(T.B, T.C);
	double c = Leng(T.A, T.C);
	double p = (a+b+c) / 2;

	return sqrt( p * (p-a) * (p-b) * (p-c) );
}

double Leng(TPoint A, TPoint B)
{
	return sqrt( pow(A.X - B.X, 2) + pow(A.Y - B.Y, 2) );
}