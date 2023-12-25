#include <iostream>
#include <fstream>
using namespace std;

void Write(ofstream& Out, int* mass, int N, int K);
bool CanIncrement(int* mass, int N, int K);

int main()
{
	//Task("Recur25");
	int N, K;
	cin >> N >> K;
	cin.ignore();
	char fileName[100];
	cin.getline(fileName, 100);
	int* mass = new int [N+1];
	mass[0] = 0;
	for (int i = 1; i <= N; i++)
		mass[i] = 1;
	ofstream Out(fileName);
	Write(Out, mass, N, K);
	Out.close();
	delete [] mass;
	mass = NULL;
	return 0;
}

void Write(ofstream& Out, int* mass, int N, int K)
{
	for (int i = 0; i <= N; i++)
		Out << mass[i];
	Out << endl;
	
	if (CanIncrement(mass, N, K))
		Write(Out, mass, N, K);
}

bool CanIncrement(int* mass, int N, int K)
{
	for (int i = N; i > 0; i--)
	{
		if (mass[i] + 1 > K)
			mass[i] = 1;
		else
		{
			mass[i]++;
			return true;
		}
	}
	return false;
}