
#!/usr/bin/env python

def read_input():
    with open('input.txt', 'r') as fin:
        return [x.rstrip() for x in fin.readlines()]

def part_one(data):
    unique_questions = set()
    sum = 0
    for answer in data:
        if not answer:
            sum = sum + len(unique_questions)
            unique_questions = set()
            continue
        for question in answer:
            unique_questions.add(question)

    return sum


def part_two(data):
    return 0


def main():
    data = read_input()
    print("part one: {}".format(part_one(data)))
    print("part two: {}".format(part_two(data)))


if __name__ == "__main__":
    main()
