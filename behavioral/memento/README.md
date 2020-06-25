# Memento

Memento design pattern is a behavioral design pattern. It allows us to save checkpoints for an object and thus allow an object to revert to its previous state. Basically it helps in undo-redo operations on an object. Below are the design components of the Memento Design Pattern.

- **Originator**: It is the actual object whose state is saved as a memento. 
- **Memento**: This is the object which saves the state of the originator
- **Caretaker**: This is the object that saves multiple mementos. Given an index, it returns the corresponding memento. 

