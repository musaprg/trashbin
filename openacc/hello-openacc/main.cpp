#include <iostream>
#include <vector>
#include <random>

using namespace std;

template<typename T>
vector<T> vector_add(const vector<T> &a, const vector<T> &b)
{
    const auto min_size = min(a.size(), b.size());

    vector<T> result;

    for (size_t i = 0 ; i < min_size; i++){
        result.emplace_back(a.at(i) + b.at(i));
    }

    return result;
}

int main() {
    vector<int> a;
    vector<int> b;

    random_device seed_gen;
    mt19937 engine(seed_gen());
    uniform_real_distribution<> dist1(-100, 100);

    for(size_t i = 0 ; i < (1<<20); i++){
        a.emplace_back(dist1(engine));
        b.emplace_back(dist1(engine));
    }

    auto result = vector_add(a,b);

    for(auto &&item : result){
        cout << item << endl;
    }

    return 0;
}
