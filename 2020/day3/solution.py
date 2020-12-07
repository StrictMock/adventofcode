
#!/usr/bin/env python

def read_input():
    with open('input.txt', 'r') as fin:
        return [x.rstrip() for x in fin.readlines()]


def count_trees(data, step_right, step_down):
    col = 0
    counter = 0
    for y in range(0, len(data), step_down):
        row = data[y]
        if '#' in row[col % len(row)]:
            counter = counter + 1
        col = col + step_right
    return counter


def part_one(data):
    counter = count_trees(data=data, step_right=3, step_down=1)
    print("part one: {}".format(counter))


def part_two(data):
    steps = [
        {"step_right": 1, "step_down": 1},
        {"step_right": 3, "step_down": 1},
        {"step_right": 5, "step_down": 1},
        {"step_right": 7, "step_down": 1},
        {"step_right": 1, "step_down": 2}]

    result = 1
    for step in steps:
        result = result * \
            count_trees(data, step["step_right"], step["step_down"])
    print("part two: {}".format(result))


def main():
    data = read_input()
    part_one(data)
    part_two(data)


if __name__ == "__main__":
    main()
