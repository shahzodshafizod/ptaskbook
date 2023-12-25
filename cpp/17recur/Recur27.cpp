#include <iostream>
#include <fstream>
using namespace std;

void Write(ofstream& Out, char* mass, int N);
int GetWeight(char* mass, int N);
bool CanContinue(char* mass, int N);

int main()
{
	//Task("Recur27");
	int N;
	cin >> N;
	cin.ignore();
	char fileName[100];
	cin.getline(fileName, 100);
	char* mass = new char [N+1];
	mass[0] = 'C';
	for (int i = 1; i <= N; i++)
		mass[i] = 'A';
	ofstream Out(fileName);
	Write(Out, mass, N);
	Out.close();
	delete [] mass;
	mass = NULL;
	return 0;
}

void Write(ofstream& Out, char* mass, int N)
{
	if (GetWeight(mass, N) == 0)
	{
		for (int i = 0; i <= N; i++)
			Out << mass[i];
		Out << endl;
	}
	if (CanContinue(mass, N))
		Write(Out, mass, N);
}

int GetWeight(char* mass, int N)
{
	int sum = 0;
	for (int i = 1; i <= N; i++)
		sum += mass[i] == 'A' ? 1 : -1;
	return sum;
}

bool CanContinue(char* mass, int N)
{
	for (int i = N; i > 0; i--)
	{
		if (mass[i] == 'A')
		{
			mass[i] = 'B';
			return true;
		}
		else
			mass[i] = 'A';
	}
	return false;
}