#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File53");
	char S[100], S0[100];
	int N;
	cin.getline(S, 100);
	cin >> N;
	cin.ignore();
	cin.getline(S0, 100);
	ofstream ofs(S, ios::binary);
	ifstream ifs(S0, ios::binary);
	int number, n;
	ifs.read(reinterpret_cast<char*>(&n), sizeof(n));
	if (N <= n)
	{
		int* sizes = new int [n];
		for (int i = 0; i < n; i++)
			ifs.read(reinterpret_cast<char*>(&sizes[i]), sizeof(sizes[i]));
		for (int i = 0; i < n; i++)
		{
			if (i+1 < N)
			{
				ifs.seekg(sizes[i]*sizeof(int), ios::cur);
				continue;
			}
			for (int j = 0; j < sizes[i]; j++)
			{
				ifs.read(reinterpret_cast<char*>(&number), sizeof(number));
				ofs.write(reinterpret_cast<const char*>(&number), sizeof(number));
			}
			break;
		}
	}
	ifs.close();
	ofs.close();
	
	return 0;
}
