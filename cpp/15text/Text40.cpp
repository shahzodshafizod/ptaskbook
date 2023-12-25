#include <iostream>
#include <fstream>
//#include <iomanip>
using namespace std;

void write2file(ofstream& fout, int number1, int width);

int main()
{
	//Task("Text40");
	char file1Name[100], file2Name[100], textFileName[100];
	cin.getline(file1Name, 100);
	cin.getline(file2Name, 100);
	cin.getline(textFileName, 100);

	ifstream file1(file1Name, ios::binary);
	ifstream file2(file2Name, ios::binary);

	ofstream textFile(textFileName);

	int number1, number2;
	while (file1.peek() != -1)
	{
		file1.read(reinterpret_cast<char*>(&number1), sizeof(number1));
		file2.read(reinterpret_cast<char*>(&number2), sizeof(number2));

		//textFile << '|' << setw(30) << number1 << setw(30) << number2 << "|\n";
		
		textFile << '|';
		
		write2file(textFile, number1, 30);
		write2file(textFile, number2, 30);
		
		textFile << "|\n";
	}

	file1.close();
	file2.close();
	textFile.close();

	return 0;
}

void write2file(ofstream& fout, int number, int width)
{
	int length = 0;
	bool isNegative = false;
	if (number == 0)
		length++;
	else
	{
		if (number < 0)
		{
			isNegative = true;
			length++;
			number *= -1;
		}
		int temp = number;
		while (temp > 0)
		{
			temp /= 10;
			length++;
		}
	}
	for (int i = width-length; i >= 1; i--)
		fout << ' ';
	
	if (isNegative)
		fout << '-';
	fout << number;
}