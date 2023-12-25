#include <iostream>
#include <fstream>
using namespace std;

void write2file(ofstream& fout, int number, int width);

int main()
{
	//Task("Text41");
	char file1Name[100], file2Name[100], file3Name[100], textFileName[100];
	cin.getline(file1Name, 100);
	cin.getline(file2Name, 100);
	cin.getline(file3Name, 100);
	cin.getline(textFileName, 100);

	ifstream file1(file1Name, ios::binary);
	ifstream file2(file2Name, ios::binary);
	ifstream file3(file3Name, ios::binary);

	ofstream textFile(textFileName);

	int number1, number2, number3;
	while (file1.peek() != -1)
	{
		file1.read(reinterpret_cast<char*>(&number1), sizeof(number1));
		file2.read(reinterpret_cast<char*>(&number2), sizeof(number2));
		file3.read(reinterpret_cast<char*>(&number3), sizeof(number3));

		textFile << '|';

		write2file(textFile, number1, 20);
		write2file(textFile, number2, 20);
		write2file(textFile, number3, 20);

		textFile << "|\n";
	}
	file1.close();
	file2.close();
	file3.close();
	textFile.close();

	return 0;
}

void write2file(ofstream& fout, int number, int width)
{
	int length = 0;
	if (number == 0)
		length++;
	else
	{
		bool isNegative = false;
		if (number < 0)
		{
			length++;
			isNegative = true;
			number *= -1;
		}
		int temp = number;
		while (temp > 0)
		{
			temp /= 10;
			length++;
		}
		if (isNegative)
			fout << '-';
	}
	fout << number;
	for (int i = width-length; i >= 1; i--)
		fout << ' ';
	
}