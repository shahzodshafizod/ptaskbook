#include <iostream>
#include <fstream>
using namespace std;

int main()
{
	//Task("File54");
	char S[100], S0[100];
	cin.getline(S, 100);
	cin.getline(S0, 100);
	ifstream ifs(S0, ios::binary);
	ofstream ofs(S, ios::binary);
	int N;
	ifs.read(reinterpret_cast<char*>(&N), sizeof(N));
	int* sizes = new int [N];
	for (int i = 0; i < N; i++)
		ifs.read(reinterpret_cast<char*>(&sizes[i]), sizeof(sizes[i]));
	int number, sum;
	double amean;
	for (int i = 0; i < N; i++)
	{
		sum = 0;
		for (int j = 0; j < sizes[i]; j++)
		{
			ifs.read(reinterpret_cast<char*>(&number), sizeof(number));
			sum += number;
		}
		amean = (double)sum / sizes[i];
		ofs.write(reinterpret_cast<const char*>(&amean), sizeof(amean));
	}
	ifs.close();
	ofs.close();
	
	return 0;
}
