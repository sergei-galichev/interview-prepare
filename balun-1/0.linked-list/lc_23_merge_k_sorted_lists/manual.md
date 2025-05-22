[23. Merge k Sorted Lists](https://leetcode.com/problems/merge-k-sorted-lists/)

```go
package main

type ListNode struct {
    Val  int
    Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    dummyNode := &ListNode{}

    curr := dummyNode

    for {
        minNodeIdx := minNodeIndex(lists)

        if minNodeIdx == nil {
            curr.Next = nil

            break
        }

        curr.Next = lists[*minNodeIdx]
        curr = curr.Next
        lists[*minNodeIdx] = lists[*minNodeIdx].Next
    }

    return dummyNode.Next
}

func minNodeIndex(lists []*ListNode) *int {
    minListNodeVal, minListNodeIndex := math.Inf(1), math.Inf(1)

    for i := 0; i < len(lists); i++ {
        if lists[i] == nil {
            continue
        }

        nodeValue := float64(lists[i].Val)

        if nodeValue < minListNodeVal {
            minListNodeVal = nodeValue
            minListNodeIndex = float64(i)
        }
    }

    if math.IsInf(minListNodeVal, 1) {
        return nil
    }

    index := int(minListNodeIndex)

    return &index
}
```

***Оценка по времени:*** `O(n)`, где `n` - количество элементов `ListNode`

*Объяснения:* делаем 1 проход по всем элементам, который образует `n` итераций

***Оценка по памяти:*** `O(1)`

*Объяснения:* количество элементов неизменнно

**Описание решения**

Пишем вспомогательную функцию minNodeIndex для поиска ноды с минимальным значением.

Решение в лоб.

Вычисляем индекс ноды с минимальным значением. Если он равен nil, то возврат. Создаем указатели: curr - указывает на 
ноду с минимальным значением; head - указывает на curr. Указатель на ноду по найденному индексу переставляем на 
curr.Next. В цикле идем по остальным нодам.

Решение лучше.

Используем DummyNode. Создаем указатель curr на dummyNode. В цикле идем по нодам. Вычисляем индекс ноды с минимальным 
значением. Если он равен nil, то curr.Next устанавливаем в nil и прерываем цикл. Если нет, то curr.Next устанавливаем
на ноду с минимальным значением. Указатель curr устанавливаем на curr.Next. Указатель на ноду с минимальным значением 
меняем на Next.