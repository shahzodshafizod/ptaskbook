#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File26");
	char fileName[100];
	cin.getline(fileName, 100);
	fstream fs(fileName, ios::binary | ios::out | ios::in);
	double number, minimal, maximal;
	fs.read(reinterpret_cast<char*>(&number), sizeof(number));
	minimal = maximal = number;
	int minInd = 0, maxInd = 0, index = 0;
	while (true)
	{
		fs.read(reinterpret_cast<char*>(&number), sizeof(number));
		if (fs.eof())
		{
			fs.clear();
			break;
		}
		index++;
		if (number < minimal)
		{
			minimal = number;
			minInd = index;
		}
		if (number > maximal)
		{
			maximal = number;
			maxInd = index;
		}
	}
	
	fs.seekg(0);
	fs.seekp(minInd*sizeof(double));
	fs.write(reinterpret_cast<const char*>(&maximal), sizeof(maximal));

	fs.seekp(maxInd*sizeof(double));
	fs.write(reinterpret_cast<const char*>(&minimal), sizeof(minimal));

	fs.close();
	
	return 0;
}
