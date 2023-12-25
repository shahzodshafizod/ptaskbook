#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File14");
	char fileName[100];
	cin.getline(fileName, 100);
	ifstream ifs(fileName, ios::binary);
	double number, sum = 0;
	int elements = 0;
	while (true)
	{
		ifs.read(reinterpret_cast<char*>(&number), sizeof(number));
		if (ifs.eof())
		{
			ifs.clear();
			break;
		}
		elements++;
		sum += number;
	}
	cout << sum / elements;
	
	return 0;
}
