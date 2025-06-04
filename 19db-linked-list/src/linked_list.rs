use std::collections::LinkedList;

#[derive(Clone, Debug)]
struct Node {
    num: i32,
    name: String,
}

pub struct DbLinkedList {
    list: LinkedList<Node>,
}

impl DbLinkedList {
    pub fn new() -> Self {
        DbLinkedList {
            list: LinkedList::new(),
        }
    }

    pub fn add_node(&mut self, num: i32, name: &str) {
        self.list.push_back(Node { num, name: name.to_string() });
    }

    pub fn remove_node(&mut self, num: i32) {
        let mut new_list = LinkedList::new();
        while let Some(node) = self.list.pop_front() {
            if node.num != num {
                new_list.push_back(node);
            }
        }
        self.list = new_list;
    }

    pub fn print(&self) {
        for node in self.list.iter() {
            println!("num=[{}], name=[{}]", node.num, node.name);
        }
    }

    pub fn print_reverse(&self) {
        for node in self.list.iter().rev() {
            println!("num=[{}], name=[{}]", node.num, node.name);
        }
    }
}
