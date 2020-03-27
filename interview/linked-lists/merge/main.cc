#include <bits/stdc++.h>

using namespace std;

class SinglyLinkedListNode {
    public:
        int data;
        SinglyLinkedListNode *next;

        SinglyLinkedListNode(int node_data) {
            this->data = node_data;
            this->next = nullptr;
        }
};

class SinglyLinkedList {
    public:
        SinglyLinkedListNode *head;
        SinglyLinkedListNode *tail;

        SinglyLinkedList() {
            this->head = nullptr;
            this->tail = nullptr;
        }

        void insert_node(int node_data) {
            SinglyLinkedListNode* node = new SinglyLinkedListNode(node_data);

            if (!this->head) {
                this->head = node;
            } else {
                this->tail->next = node;
            }

            this->tail = node;
        }
};

void print_singly_linked_list(SinglyLinkedListNode* node, string sep, ofstream& fout) {
    while (node) {
        fout << node->data;

        node = node->next;

        if (node) {
            fout << sep;
        }
    }
}

void free_singly_linked_list(SinglyLinkedListNode* node) {
    while (node) {
        SinglyLinkedListNode* temp = node;
        node = node->next;

        free(temp);
    }
}

// Complete the mergeLists function below.

/*
 * For your reference:
 *
 * SinglyLinkedListNode {
 *     int data;
 *     SinglyLinkedListNode* next;
 * };
 *
 */
SinglyLinkedListNode* mergeLists(SinglyLinkedListNode* head1, SinglyLinkedListNode* head2) {
  SinglyLinkedListNode* m = nullptr;
  SinglyLinkedListNode* curr;
  cout << "m:" << m << " curr:" << curr << "\n" << std::flush;
  do {
      if (head1 == nullptr) {
          cout << "head1 empty " << "m:" << m << " curr:" << curr << "\n" << std::flush;
          if (m == nullptr) {
              return head2;
          } else {
            curr->next = head2;
          }
          return m;
      }
      if (head2 == nullptr) {
          cout << "head2 empty " << "m:" << m << " curr:" << curr << "\n" << std::flush;
          if (m == nullptr) {
              return head1;
          } else {
            curr->next = head1;
          }
        return m;
      }
      cout << head1->data << " " << head2->data << "\n";
      if (head1->data <= head2->data) {
          cout << "head1 < " << "m:" << m << " curr:" << curr << "\n" << std::flush;
          SinglyLinkedListNode* temp = new SinglyLinkedListNode(head1->data);
          if (m == nullptr) {
              m = temp;
              curr = m;
          } else {
            curr->next = temp;
            curr = temp;
          }
          head1 = head1->next;
      } else {
          cout << "head2 < " << "m:" << m << " curr:" << curr << "\n" << std::flush;
          SinglyLinkedListNode* temp = new SinglyLinkedListNode(head2->data);
          if (m == nullptr) {
              m = temp;
              curr = m;
          } else {
            curr->next = temp;
            curr = temp;
          }
          head2 = head2->next;
      }
      cout << curr;
  } while (head1 != nullptr || head2 != nullptr);
  return m;
}

int main()
{
    ofstream fout(getenv("OUTPUT_PATH"));

    int tests;
    cin >> tests;
    cin.ignore(numeric_limits<streamsize>::max(), '\n');

    for (int tests_itr = 0; tests_itr < tests; tests_itr++) {
        SinglyLinkedList* llist1 = new SinglyLinkedList();

        int llist1_count;
        cin >> llist1_count;
        cin.ignore(numeric_limits<streamsize>::max(), '\n');

        for (int i = 0; i < llist1_count; i++) {
            int llist1_item;
            cin >> llist1_item;
            cin.ignore(numeric_limits<streamsize>::max(), '\n');

            llist1->insert_node(llist1_item);
        }
      
      	SinglyLinkedList* llist2 = new SinglyLinkedList();

        int llist2_count;
        cin >> llist2_count;
        cin.ignore(numeric_limits<streamsize>::max(), '\n');

        for (int i = 0; i < llist2_count; i++) {
            int llist2_item;
            cin >> llist2_item;
            cin.ignore(numeric_limits<streamsize>::max(), '\n');

            llist2->insert_node(llist2_item);
        }

        SinglyLinkedListNode* llist3 = mergeLists(llist1->head, llist2->head);

        print_singly_linked_list(llist3, " ", fout);
        fout << "\n";

        free_singly_linked_list(llist3);
    }

    fout.close();

    return 0;
}
