#include <iostream>
#include <fstream>
using namespace std;

void Write(ofstream& Out, char* mass, int N);
bool isOk(char* mass, int N);
bool CanContinue(char* mass, int N);

int main()
{
	//Task("Recur30");
	int N;
	cin >> N;
	char* mass = new char [N+1];
	mass[0] = 'D';
	for (int i = 1; i <= N; i++)
		mass[i] = 'A';
	char fileName[100];
	cin >> fileName;
	ofstream Out(fileName);
	Write(Out, mass, N);
	Out.close();
	delete [] mass;
	mass = NULL;
	return 0;
}

void Write(ofstream& Out, char* mass, int N)
{
	if (isOk(mass, N))
	{
		for (int i = 0; i <= N; i++)
			Out << mass[i];
		Out << endl;
	}
	if (CanContinue(mass, N))
		Write(Out, mass, N);
}

bool isOk(char* mass, int N)
{
	int sum = mass[1] == 'A' ? 1 : mass[1] == 'C' ? -1 : 0;
	for (int i = 2; i <= N; i++)
	{
		sum += mass[i] == 'A' ? 1 : mass[i] == 'C' ? -1 : 0;
		if (mass[i] == mass[i-1])
			return false;
	}
	return sum == 0;
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
		else if (mass[i] == 'B')
		{
			mass[i] = 'C';
			return true;
		}
		else
			mass[i] = 'A';
	}
	return false;
}