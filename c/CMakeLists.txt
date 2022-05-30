cmake_minimum_required(VERSION 3.21)
project(demo C)

set(CMAKE_C_STANDARD 99)

add_executable(demo datatest/datatest.c ptrtest/ptrtest.c functest/functest.c arraytest/arraytest.c strtest/strtest.c memtest/memtest.c structtest/structtest.c functest/functest.h arraytest/arraytest.h datatest/datatest.h memtest/memtest.h ptrtest/ptrtest.h strtest/strtest.h structtest/structtest.h main.c uniontest/uniontest.c uniontest/uniontest.h pretest/pretest.c pretest/pretest.h iotest/iotest.c iotest/iotest.h explaintest/explain.c explaintest/explain.h explaintest/a.c explaintest/a.h)
