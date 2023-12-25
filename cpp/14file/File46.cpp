#include <iostream>
#include <fstream>
using namespace std;

void AddDatas(const char* from, const char* to);

int main()
{
	//Task("File46");
	char S0[100], S[100];
	cin.getline(S0, 100);
	int N;
	cin >> N;
	cin.ignore();
	for (int i = 0; i < N; i++)
	{
		cin.getline(S, 100);
		AddDatas(S, S0);
	}
}

void AddDatas(const char* from, const char* to)
{
	ifstream ifs(from, ios::binary);
	ofstream ofs(to, ios::binary | ios::app);
	char C;
	while (true)
	{
		ifs.read(reinterpret_cast<char*>(&C), sizeof(C));
		if (ifs.eof())
		{
			ifs.clear();
			break;
		}
		ofs.write(reinterpret_cast<const char*>(&C), sizeof(C));
	}
	ifs.close();
	ofs.close();
}