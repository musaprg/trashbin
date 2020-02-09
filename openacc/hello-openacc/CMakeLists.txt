cmake_minimum_required(VERSION 3.10)

project(hello-openacc LANGUAGES CXX)

find_package(OpenACC REQUIRED)

add_executable(hello-openacc
    main.cpp
    )
