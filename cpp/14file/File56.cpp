#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File56");
	char S[100], S0[100];
	int N, n;
	cin.getline(S, 100);
	cin >> N;
	cin.ignore();
	cin.getline(S0, 100);
	ifstream ifs(S0, ios::binary);
	ofstream ofs(S, ios::binary);
	int number, elements;
	for (int i = 1; true; i++)
	{
		ifs.read(reinterpret_cast<char*>(&elements), sizeof(elements));
		if (ifs.eof())
		{
			ifs.clear();
			break;
		}
		if (i < N)
		{
			ifs.seekg(elements*sizeof(int), ios::cur);
			continue;
		}
		for (int j = 0; j < elements; j++)
		{
			ifs.read(reinterpret_cast<char*>(&number), sizeof(number));
			ofs.write(reinterpret_cast<const char*>(&number), sizeof(number));
		}
		break;
	}
	ifs.close();
	ofs.close();
	
	return 0;
}
