#include <iostream>
#include <iterator>
#include <algorithm>
using namespace std;

class F {
public:
	double A;
	double D;
	int index;
	F(double _A, double _D, int _index) {
		A = _A;
		D = _D;
		index = _index;
	}
	double operator() () {
		return A + index++ * D;
	}
};



int main() {
	double A, D;
	cin >> A >> D;
	int N;
	cin >> N;
	generate_n(ostream_iterator<double>(cout, " "), N, F(A, D, 0));
	return 0;
}