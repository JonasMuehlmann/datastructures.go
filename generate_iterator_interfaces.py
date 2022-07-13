#!/usr/bin/env python3

import asyncio
import re
import signal
from dataclasses import dataclass
from typing import Dict, List, Optional, Tuple

from matplotlib._api import itertools

FILE_HEADER: str = """// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ds"""

BOOL: str = "bool"
INT: str = "int"
T: str = "T"

ITERATOR: str = "Z"
ORDERED_ITERATOR: str = "Ordered_"
COMPARABLE_ITERATOR: str = "Comparable_"
SIZED_ITERATOR: str = "Sized_"
WRITEABLE_ITERATOR: str = "Writeable_"
READABLE_ITERATOR: str = "Readable_"
FOReadWriteARD_ITERATOR: str = "Forward_"
REVERSED_ITERATOR: str = "Reversed_"
BACKWARD_ITERATOR: str = "Backward_"
UNORDERED_FOReadWriteARD_ITERATOR: str = "UnorderedForward_"
UNORDERED_REVERSED_ITERATOR: str = "UnorderedReversed_"
UNORDERED_BACKWARD_ITERATOR: str = "UnorderedBackward_"
BIDIRECTIONAL_ITERATOR: str = "Bidirectional_"

RANDOM_ACCESS_ITERATOR: str = "RandomAccess_"

GENERATED_INTERFACES: Dict[str, "Interface"] = {}
PERMUTATION_LENGHT: int = 0
PREVIOUS_PERMUTATION_LENGTH: int = -1


def split_name(names: List[str]) -> List[str]:
    new_names: List[str] = []

    for name in names:
        tmp = name.split("_")
        try:
            tmp.remove("")
        except ValueError:
            pass

        new_names += tmp

    new_names = [name for name in new_names if name != ITERATOR]

    return new_names


def does_hierarchy_contain_interface(
    target: "Interface", hierarchy: "Interface"
) -> bool:
    if hierarchy.name == target.name:
        return True

    if hierarchy.inherited_interfaces is None:
        return False

    return any(
        does_hierarchy_contain_interface(target, GENERATED_INTERFACES[interface])
        for interface in hierarchy.inherited_interfaces
    )


def interface_contained_in_combinations_hierarchy(combination: List[str]) -> bool:

    for interface in combination:
        hierarchies_names: List[str] = combination.copy()
        hierarchies_names.remove(interface)

        for hierarchy_name in hierarchies_names:
            if does_hierarchy_contain_interface(
                GENERATED_INTERFACES[interface], GENERATED_INTERFACES[hierarchy_name]
            ):
                return True

    return False


def is_method_generic(method: "Method") -> bool:
    if method.parameters is None:
        return False

    return any(p.type == T for p in method.parameters)


def is_interface_generic(interface: "Interface") -> bool:
    has_generic_methods: bool = False
    has_generic_inherited_interface: bool = False

    if interface.methods is not None:
        has_generic_methods = any(map(is_method_generic, interface.methods))

    if interface.inherited_interfaces is not None:
        has_generic_inherited_interface = any(
            is_interface_generic(GENERATED_INTERFACES[name])
            for name in interface.inherited_interfaces
        )

    return has_generic_methods or has_generic_inherited_interface


def make_inherited_interface_name(interface: "Interface") -> str:
    name: str = interface.name

    if is_interface_generic(interface):
        name += "[T any]"

    return name


@dataclass
class Parameter:
    name: str
    type: str

    def __str__(self):
        return f"{self.name} {self.type}"


@dataclass
class Method:
    name: str
    parameters: Optional[List[Parameter]] = None
    return_value: str = ""

    def __str__(self):
        return f"{self.name}({', '.join(map(str, self.parameters if self.parameters is not None else ''))}) {self.return_value}"


@dataclass
class Interface:
    name: str
    methods: Optional[List[Method]] = None
    inherited_interfaces: Optional[List[str]] = None

    def __str__(self):
        inherited_methods_fragment: str = ""

        if self.inherited_interfaces is not None:
            inherited_methods_fragment = "\n    ".join(
                make_inherited_interface_name(GENERATED_INTERFACES[name])
                for name in self.inherited_interfaces
            )

        methods_fragment: str = "\n    ".join(
            map(str, self.methods if self.methods else "")
        )

        type_param_fragment: str = "[T any]" if is_interface_generic(self) else ""

        return f"type {self.name}{type_param_fragment} interface {{\n    {inherited_methods_fragment}\n\n    {methods_fragment}\n}}"


iterator: Interface = Interface(
    name="Iterator",
    methods=[
        Method(name="Begin"),
        Method(name="End"),
        Method(name="IsBegin", return_value=BOOL),
        Method(name="IsEnd", return_value=BOOL),
        Method(name="First", return_value=BOOL),
        Method(name="Last", return_value=BOOL),
        Method(name="IsValid", return_value=BOOL),
    ],
)

BASE_INTERFACES: Dict[str, Interface] = {
    ORDERED_ITERATOR: Interface(
        name=ORDERED_ITERATOR,
        inherited_interfaces=[ITERATOR],
        methods=[
            Method(
                name="IsBefore",
                parameters=[Parameter("other", ORDERED_ITERATOR)],
                return_value=BOOL,
            ),
            Method(
                name="After",
                parameters=[Parameter("other", ORDERED_ITERATOR)],
                return_value=BOOL,
            ),
            Method(
                name="DistanceTo",
                parameters=[Parameter("other", ORDERED_ITERATOR)],
                return_value=INT,
            ),
        ],
    ),
    COMPARABLE_ITERATOR: Interface(
        name=COMPARABLE_ITERATOR,
        inherited_interfaces=[ITERATOR],
        methods=[
            Method(
                name="IsEqual",
                parameters=[Parameter("other", COMPARABLE_ITERATOR)],
                return_value=BOOL,
            ),
        ],
    ),
    SIZED_ITERATOR: Interface(
        name=SIZED_ITERATOR,
        inherited_interfaces=[ITERATOR],
        methods=[
            Method(
                name="Size",
                return_value=INT,
            ),
        ],
    ),
    WRITEABLE_ITERATOR: Interface(
        name=WRITEABLE_ITERATOR,
        inherited_interfaces=[ITERATOR],
        methods=[
            Method(
                name="SET",
                parameters=[Parameter("value", T)],
            ),
        ],
    ),
    READABLE_ITERATOR: Interface(
        name=READABLE_ITERATOR,
        inherited_interfaces=[ITERATOR],
        methods=[
            Method(
                name="Get",
                parameters=[Parameter("value", T)],
            ),
        ],
    ),
    FOReadWriteARD_ITERATOR: Interface(
        name=FOReadWriteARD_ITERATOR,
        inherited_interfaces=[ITERATOR],
        methods=[
            Method(
                name="Advance",
                return_value=BOOL,
            ),
            Method(
                name="Next",
                return_value=FOReadWriteARD_ITERATOR,
            ),
            Method(
                name="AdvanceN",
                parameters=[Parameter("n", INT)],
                return_value=BOOL,
            ),
            Method(
                name="NextN",
                parameters=[Parameter("n", INT)],
                return_value=FOReadWriteARD_ITERATOR,
            ),
        ],
    ),
    REVERSED_ITERATOR: Interface(
        name=REVERSED_ITERATOR, inherited_interfaces=[FOReadWriteARD_ITERATOR]
    ),
    BACKWARD_ITERATOR: Interface(
        name=BACKWARD_ITERATOR,
        inherited_interfaces=[ITERATOR],
        methods=[
            Method(
                name="Recede",
                return_value=BOOL,
            ),
            Method(
                name="Previous",
                return_value=BACKWARD_ITERATOR,
            ),
            Method(
                name="RecedeN",
                parameters=[Parameter("n", INT)],
                return_value=BOOL,
            ),
            Method(
                name="PreviousN",
                parameters=[Parameter("n", INT)],
                return_value=BACKWARD_ITERATOR,
            ),
        ],
    ),
    UNORDERED_FOReadWriteARD_ITERATOR: Interface(
        name=UNORDERED_FOReadWriteARD_ITERATOR,
        inherited_interfaces=[FOReadWriteARD_ITERATOR],
        methods=[
            Method(
                name="AdvanceOrdered",
                return_value=BOOL,
            ),
            Method(
                name="NextOrdered",
                return_value=UNORDERED_FOReadWriteARD_ITERATOR,
            ),
            Method(
                name="AdvanceOrderedN",
                parameters=[Parameter("n", INT)],
                return_value=BOOL,
            ),
            Method(
                name="NextOrderedN",
                parameters=[Parameter("n", INT)],
                return_value=UNORDERED_FOReadWriteARD_ITERATOR,
            ),
        ],
    ),
    UNORDERED_REVERSED_ITERATOR: Interface(
        name=UNORDERED_REVERSED_ITERATOR,
        inherited_interfaces=[REVERSED_ITERATOR],
    ),
    UNORDERED_BACKWARD_ITERATOR: Interface(
        name=UNORDERED_BACKWARD_ITERATOR,
        inherited_interfaces=[BACKWARD_ITERATOR],
        methods=[
            Method(
                name="RecedeOrdered",
                return_value=BOOL,
            ),
            Method(
                name="PreviousOrdered",
                return_value=UNORDERED_BACKWARD_ITERATOR,
            ),
            Method(
                name="RecedeOrderedN",
                parameters=[Parameter("n", INT)],
                return_value=BOOL,
            ),
            Method(
                name="PreviousOrderedN",
                parameters=[Parameter("n", INT)],
                return_value=UNORDERED_BACKWARD_ITERATOR,
            ),
        ],
    ),
    BIDIRECTIONAL_ITERATOR: Interface(
        name=BIDIRECTIONAL_ITERATOR,
        inherited_interfaces=[FOReadWriteARD_ITERATOR, BACKWARD_ITERATOR],
        methods=[
            Method(
                name="MoveBy",
                parameters=[Parameter("n", INT)],
                return_value=BOOL,
            ),
            Method(
                name="Nth",
                parameters=[Parameter("n", INT)],
                return_value=BIDIRECTIONAL_ITERATOR,
            ),
        ],
    ),
    RANDOM_ACCESS_ITERATOR: Interface(
        name=RANDOM_ACCESS_ITERATOR,
        inherited_interfaces=[ITERATOR],
        methods=[
            Method(
                name="MoveTo",
                parameters=[Parameter("i", INT)],
                return_value=BOOL,
            ),
            Method(
                name="GetAt",
                parameters=[Parameter("i", INT)],
                return_value=BOOL,
            ),
            Method(
                name="SetAt",
                parameters=[Parameter("i", INT), Parameter("value", T)],
                return_value=BOOL,
            ),
            Method(
                name="Index",
                return_value=T,
            ),
        ],
    ),
}


def add_interface(combination: List[str]):
    new_interface_name: str = ""
    # print(combination)

    for interface in combination:
        new_interface_name += GENERATED_INTERFACES[interface].name.split(ITERATOR)[0]

    new_interface_name += ITERATOR
    new_interface = Interface(
        name=new_interface_name,
        inherited_interfaces=[
            GENERATED_INTERFACES[interface].name for interface in combination
        ],
    )

    GENERATED_INTERFACES[new_interface_name] = new_interface


def generate_interfaces():
    global GENERATED_INTERFACES
    global PERMUTATION_LENGHT
    global PREVIOUS_PERMUTATION_LENGTH

    GENERATED_INTERFACES = BASE_INTERFACES
    GENERATED_INTERFACES[ITERATOR] = iterator

    PERMUTATION_LENGHT = len(BASE_INTERFACES) + 1

    while PREVIOUS_PERMUTATION_LENGTH != PERMUTATION_LENGHT:
        for next_combination_length in range(2, PERMUTATION_LENGHT):
            next_combinations: List[Tuple[str]] = list(
                itertools.combinations(
                    GENERATED_INTERFACES.keys(), next_combination_length
                )
            )

            for combination in next_combinations:
                combination = list(combination)

                if ITERATOR in combination:
                    continue

                split: List[str] = sorted(split_name(combination)) + [ITERATOR]
                split_set: List[str] = sorted(list(set(split_name(combination)))) + [
                    ITERATOR
                ]

                if len(split) != len(split_set):
                    continue

                if (
                    (FOReadWriteARD_ITERATOR[:-1] in split and BACKWARD_ITERATOR[:-1] in split)
                    or (
                        FOReadWriteARD_ITERATOR[:-1] in split
                        and BIDIRECTIONAL_ITERATOR[:-1] in split
                    )
                    or (
                        FOReadWriteARD_ITERATOR[:-1] in split
                        and REVERSED_ITERATOR[:-1] in split
                    )
                    or (
                        BACKWARD_ITERATOR[:-1] in split
                        and REVERSED_ITERATOR[:-1] in split
                    )
                    or (
                        BACKWARD_ITERATOR[:-1] in split
                        and BIDIRECTIONAL_ITERATOR[:-1] in split
                    )
                ):

                    continue

                if ORDERED_ITERATOR[:-1] in split and (
                    UNORDERED_BACKWARD_ITERATOR[:-1] in split
                    or UNORDERED_FOReadWriteARD_ITERATOR[:-1]
                    or UNORDERED_REVERSED_ITERATOR[:-1] in split
                ):
                    continue

                if interface_contained_in_combinations_hierarchy(combination):
                    continue

                print(split)
                add_interface(combination)
                print("Generated interfaces: ", len(GENERATED_INTERFACES))

        PREVIOUS_PERMUTATION_LENGTH = PERMUTATION_LENGHT
        PERMUTATION_LENGHT = len(GENERATED_INTERFACES)
        print(PERMUTATION_LENGHT)


def signal_handler(signum, frame):
    print("Generated interfaces: ", len(GENERATED_INTERFACES))
    write_interfaces()
    raise Exception("Timed out or cancelled!")


def write_interfaces():
    with open("interfaces.go", "w+", encoding="utf-8") as f:
        f.write(
            "\n\n".join([FILE_HEADER] + list(map(str, GENERATED_INTERFACES.values())))
        )


def main():
    signal.signal(signal.SIGALRM, signal_handler)
    signal.signal(signal.SIGINT, signal_handler)
    signal.signal(signal.SIGHUP, signal_handler)
    signal.signal(signal.SIGTERM, signal_handler)
    # signal.alarm(5)  # Ten seconds

    generate_interfaces()

    write_interfaces()


# TODO: I might need to reimplement this in go

if __name__ == "__main__":
    main()
    print("Done")
