#define CATCH_CONFIG_MAIN
#include "catch.hpp"
#include "../singly_linked_list.hpp"

TEST_CASE("Singly linked list basic", "[SINGLY_LINKED_LIST]") {
    SinglyLinkedList<int> list;

    list.append(0);
    list.append(1);
    list.append(2);

    REQUIRE(list.getHead()->getValue() == 2);
    REQUIRE(list.getLength() == 3);
}
