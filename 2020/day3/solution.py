
#!/usr/bin/env python

def read_input():
    with open('input.txt', 'r') as fin:
        return [x.rstrip() for x in fin.readlines()]


def part_one(data):
    x = 0
    counter = 0
    for y in data:
        if '#' in y[x % len(y)]:
            counter = counter + 1
        x = x + 3

    print("part one: {}".format(counter))


def part_two(data):
    pass


def main():
    data = read_input()
    part_one(data)
    part_two(data)


if __name__ == "__main__":
    main()
