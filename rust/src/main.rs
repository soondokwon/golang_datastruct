mod linked_list;
use linked_list::DbLinkedList;

fn main() {
    let mut list = DbLinkedList::new();
    list.add_node(1, "benny");
    list.add_node(2, "john");
    list.add_node(3, "tommy");
    println!("--------------------------");
    list.print();
    list.remove_node(2);
    println!("--------------------------");
    list.print();
    println!("--------------------------");
    list.print_reverse();
}
