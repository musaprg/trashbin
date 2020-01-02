cmake_minimum_required(VERSION 2.8)
project(CudaWithCpp CUDA CXX)

set(SOURCES
    main.cc
)

set(HEADERS
    # main.h
)

add_executable(main
    ${SOURCES}
    ${HEADERS}
)

find_package(CUDA REQUIRED)
add_library(Example
example.h
example.cu
)
target_link_libraries(Example PUBLIC ${CUDA_LIBRARIES})

target_link_libraries(main Example)
