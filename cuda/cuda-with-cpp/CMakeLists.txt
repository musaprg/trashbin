cmake_minimum_required(VERSION 2.8)
project(CudaWithCpp)

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
if(CUDA_FOUND)
    message("CUDA version: " ${CUDA_VERSION})
    CUDA_ADD_LIBRARY(Example STATIC
        example.h
        example.cu
    )
endif()

target_link_libraries(main Example)
