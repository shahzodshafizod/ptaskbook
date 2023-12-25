#include <iostream>
#include <fstream>
#include <iomanip>
#include <cmath>
using namespace std;

void write2file(ofstream& fout, double number, int intWidth, int precisionWidth, int width);

int main()
{
	//Task("Text42");
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
		//write2file(file, i, 5, 4, 10);
		file << setw(10) << fixed << setprecision(4) << i;
		file << setw(15) << fixed << setprecision(8) << sqrt(i) << endl;
	}
	file.close();
	
	return 0;
}

/*
void write2file(ofstream& fout, double number, int intWidth, int precisionWidth, int width)
{
	int intPart = (int)number;
	double precisionPart = number - intPart;

	int temp = intPart;
	int length = 0;
	if (intPart == 0)
		length++;
	else
		while (temp > 0)
		{
			temp /= 10;
			length++;
		}
	for (int i = intWidth-length; i >= 1; i--)
		fout << ' ';
	fout << intPart << '.';
	int digit;
	for (int i = 1; i <= precisionWidth; i++)
	{
		precisionPart *= 10;
		temp = (int)precisionPart;
		digit = temp % 10;
		fout << digit;
	}
}
*/