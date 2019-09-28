#include "example.h"

__global__ void test_kernel(void) {
}

namespace Example {
    void example(void)
    {
        test_kernel <<<1, 1>>> ();
        cout << "Hello, world"s << endl;
    }
}
