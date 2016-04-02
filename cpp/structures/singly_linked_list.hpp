//Template class for Singly Linked List Element
template <typename T>
class SinglyLinkedListElement {
    public:
        SinglyLinkedListElement( const T& value ): data(value), next( nullptr ) {}
        ~SinglyLinkedListElement() {}
        const T& getValue() { return data; }
        SinglyLinkedListElement* getNext() { return next; }
        void setValue(const T& value) { data = value; }
        void setNext(SinglyLinkedListElement* element) { next = element; }

    private:
        T data;
        SinglyLinkedListElement* next;
};


//Template class for Singly Linked List Element
template <typename T>
class SinglyLinkedList {

    public:
        SinglyLinkedList(): head( nullptr ), length( 0 ) {}
        ~SinglyLinkedList() {};
        void append(const T& value);
        const T& remove();
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
const T& SinglyLinkedList<T>::remove() {
    if (head != nullptr) {
        SinglyLinkedListElement<T> *removedElement = head;
        const T& removedElementValue = removedElement->getValue();
        head = removedElement->getNext();

        delete removedElement;

        length--;

        return removedElementValue;
    } else {
        throw "Tried to remove element from an empty list";
    }
}
