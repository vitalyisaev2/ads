#include <iostream>

//Template class for Singly Linked List Element
template <typename T>
class SinglyLinkedListElement {
    public:
        SinglyLinkedListElement( const T& value ): data(value), next( nullptr ) {}
        ~SinglyLinkedListElement() {}
        const T& getValue() const { return data; }
        SinglyLinkedListElement<T>* getNext() const { return next; }
        void setValue(const T& value) { data = value; }
        void setNext(SinglyLinkedListElement* element) { next = element; }
        template <typename V> friend std::ostream& operator<<(std::ostream& os, const SinglyLinkedListElement<V>& element);

    private:
        T data;
        SinglyLinkedListElement* next;
};

template <typename T>
std::ostream& operator<<(std::ostream& os, const SinglyLinkedListElement<T>& element) {
    os << "Own address: " << &element << "; next address: " << element.getNext() << "; own value: " << element.getValue();
    return os;
}


//Template class for Singly Linked List Element
template <typename T>
class SinglyLinkedList {

    public:
        SinglyLinkedList(): head( nullptr ), length( 0 ) {}
        ~SinglyLinkedList() {};
        void append(const T& value);
        const T remove();
        SinglyLinkedListElement<T>* getHead() { return head; };
        unsigned int getLength() { return length; };

    private:
        SinglyLinkedListElement<T> *head;
        unsigned int length;
};

template <class T>
void SinglyLinkedList<T>::append(const T& value) {

    SinglyLinkedListElement<T>* element = new SinglyLinkedListElement<T>(value);

    if (head == nullptr) {
        head = element;
    } else {
        element->setNext(head);
        head = element;
    }

    length++;
    return;
}

template <class T>
const T SinglyLinkedList<T>::remove() {
    if (head != nullptr) {
        SinglyLinkedListElement<T> *removedElement = head;
        auto removedElementValue = removedElement->getValue();
        std::cout << "Removing: " << *removedElement << std::endl;

        head = removedElement->getNext();
        delete removedElement;
        length--;

        std::cout << "Return: " << removedElementValue << std::endl;
        return removedElementValue;
    } else {
        throw "Tried to remove element from an empty list";
    }
}
