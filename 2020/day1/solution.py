#!/usr/bin/env python

import itertools


def read_input():
    with open('input.txt', 'r') as fin:
        return [int(x) for x in fin.readlines()]


def calculate(numbers_array):
    combined = list(itertools.product(*numbers_array))
    for t in combined:
        sum = 0
        mul = 1
        for num in t:
            sum = sum + num
            mul = mul * num
        if sum == 2020:
            print("part one: {}".format(mul))
            return


def part_one(data):
    nums = [data, data]
    calculate(nums)


def part_two(data):
    nums = [data, data, data]
    calculate(nums)


def main():
    input = read_input()
    part_one(input)
    part_two(input)


if __name__ == "__main__":
    main()
