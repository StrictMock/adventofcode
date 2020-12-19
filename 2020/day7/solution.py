
#!/usr/bin/env python

import re


def parse_input():
    raw_data = []
    with open('input.txt', 'r') as fin:
        raw_data.extend([x.rstrip() for x in fin.readlines()])
    tree = {}
    for rule in raw_data:
        first_split = re.split(" bags contain ", rule)
        second_split = [x for x in re.split(
            " bags, | bag, | bags.| bag.", first_split[1]) if (x and "no other" not in x)]
        subtree = {" ".join(x.split()[1:]): int(
            x.split()[0]) for x in second_split}
        tree[first_split[0]] = subtree
    return tree


def dfs(visited, graph, node):
    if node not in visited:
        visited.add(node)
        for neighbour in graph[node]:
            dfs(visited, graph, neighbour)


def deep_first_search(graph, node):
    visited = set()
    dfs(visited, graph, node)
    return "shiny gold" in visited


def part_one(tree):
    count = 0
    keys_other_than_shiny_gold = [k for k in tree.keys() if "shiny gold" not in k]
    for key in keys_other_than_shiny_gold:
        if deep_first_search(tree, key):
            count += 1

    return count


def part_two(tree):
    return 0


def main():
    tree = parse_input()
    print("part one: {}".format(part_one(tree)))
    print("part two: {}".format(part_two(tree)))


if __name__ == "__main__":
    main()
