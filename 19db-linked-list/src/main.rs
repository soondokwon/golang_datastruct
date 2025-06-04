mod linked_list;
use linked_list::DbLinkedList;

fn main() {
    let mut linked_list = DbLinkedList::new();
    linked_list.add_node(1, "benny");
    linked_list.add_node(2, "john");
    linked_list.add_node(3, "tommy");
    linked_list.add_node(4, "jack");
    linked_list.add_node(5, "allice");
    linked_list.add_node(6, "jordan");
    linked_list.add_node(7, "kwon");
    linked_list.add_node(8, "kim");

    println!("--------------------------");
    linked_list.print();

    linked_list.remove_node(5);
    println!("--------------------------");
    linked_list.print();

    linked_list.remove_node(8);
    println!("--------------------------");
    linked_list.print();

    linked_list.remove_node(1);
    println!("--------------------------");
    linked_list.print();

    println!("--------------------------");
    linked_list.print_reverse();
}
