#include <iostream>
#include <fstream>
#include <cmath>
#include <iomanip>
using namespace std;

int main()
{
	//Task("Text43");
	double A, B;
	cin >> A >> B;

	int N;
	cin >> N;

	char fileName[100];
	cin.ignore();
	cin.getline(fileName, 100);

	ofstream file(fileName);

	for (double i = A; i <= B+0.0001; i += (B-A)/N)
	{
		file << setw(8) << fixed << setprecision(4) << i;
		file << fixed << setprecision(8);
		file << setw(12) << sin(i);
		file << setw(12) << cos(i);
		file << endl;
	}
	file.close();
	
	return 0;
}
