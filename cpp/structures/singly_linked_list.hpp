//Template class for Singly Linked List Element
template <typename T>
class SinglyLinkedListElement {
    public:
        SinglyLinkedListElement( const T& value ): data(value), next( nullptr ) {}
        ~SinglyLinkedListElement() {}
        const T& getValue() { return this->data; }
        SinglyLinkedListElement* getNext() { return this->next; }
        void setValue(const T& value) { this->data = value; }
        void setNext(SinglyLinkedListElement* element) { this->next = element; }

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
        SinglyLinkedListElement<T>* getHead() { return this->head; };
        unsigned int getLength() { return this->length; };

    private:
        SinglyLinkedListElement<T> *head;
        unsigned int length;
};

template <class T>
void SinglyLinkedList<T>::append(const T& value) {

    SinglyLinkedListElement<T>* element = new SinglyLinkedListElement<T>(value);

    if (this->head == nullptr) {
        this->head = element;
    } else {
        element->setNext(this->head);
        this->head = element;
    }

    length++;
    return;
}

template <class T>
const T& SinglyLinkedList<T>::remove() {
    SinglyLinkedListElement<T> *removedElement = this->head;
    const T& removedElementValue = removedElement->getValue();
    this->head = removedElement->getNext();

    delete removedElement;
    length--;

    return removedElementValue;
}
