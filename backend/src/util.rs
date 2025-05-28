use std::{collections::HashSet, hash};

pub fn contains_duplicates<T, I>(iterator: I) -> bool
    where
        T: Eq + hash::Hash,
        I: Iterator<Item = T> {
    let mut values = HashSet::new();
    for value in iterator {
        if values.contains(&value) {
            return true;
        }

        values.insert(value);
    }

    false
}
