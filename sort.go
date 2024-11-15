package sort


import (
    "slices"
)

type CompareFunc func(int, int) bool

// バブルソート
func BubbleSort(array []int, compare CompareFunc) []int {
    for i := range array {
        for j := i+1; j < len(array); j++ {
            if compare(array[i], array[j]) {
                swap(&array[i], &array[j])
            }
        }
    }
    return array
}

// 選択ソート
func SelectSort(array []int, compare CompareFunc) []int {
    for i := range array {
        tmp := array[i]
        idx := i
        for j:= i+1; j < len(array); j++ {
            if compare(tmp, array[j]) {
                tmp = array[j]
                idx = j
            }
        }
        array[idx] = array[i]
        array[i]   = tmp
    }
    return array
}

// クイックソート
func QuickSort(array []int, compare CompareFunc) []int {
    count := len(array)
    if count <= 1 {
        return array
    }
    // 境界は先頭の要素を使用
    pivot := array[0]
    left  := []int{}
    right := []int{}
    for i:= 1; i < count; i++ {
        if compare(pivot, array[i]) {
            left = append(left, array[i])
        } else {
            right = append(right, array[i]) 
        }
    }
    SortedLeft  := QuickSort(left, compare)
    SortedRight := QuickSort(right, compare)
    return slices.Concat(append(SortedLeft, pivot), SortedRight)
}


// マージソート
func MergeSort(array []int, compare CompareFunc) []int {
    count := len(array)
    if count <= 1 {
        return array
    }
    divNum := count/2
    left   := array[:divNum]
    right  := array[divNum:]

    sortedLeft  := MergeSort(left, compare)
    sortedRight := MergeSort(right, compare)
    sorted := make([]int, 0, count)

    // 右の要素番号保存
    j := 0
    for i := range sortedLeft {
        for ;j < len(sortedRight); j++ {
            if compare(sortedLeft[i], sortedRight[j]) {
                sorted = append(sorted, sortedRight[j])
                // 右の要素が最後だった場合、残る左をすべて入れてreturn
                if j+1 == len(sortedRight) {
                    return slices.Concat(sorted, sortedLeft[i:])
                }
            } else {
                sorted = append(sorted, sortedLeft[i])
                // 左の要素が最後だった場合、残る右をすべて入れてreturn
                if i+1 == len(sortedLeft) {
                    return slices.Concat(sorted, sortedRight[j:])
                }
                // 次の右の要素も見るのをやめて、左の要素を進める
                break
            }
        }
    }
    return sorted
}


// ヒープソート
func HeapSort(array []int, compare CompareFunc) []int {
    for i := range array {
        current_idx := i
        for current_idx > 0 {
            parent_idx := (current_idx -1) / 2
            // 親子の入れ替わりがあった場合、入れ替え後の親要素を確認する
            if compare(array[current_idx], array[parent_idx]) {
                swap(&array[current_idx], &array[parent_idx])
                current_idx = parent_idx
            } else {
                // 親子の入れ替わりがない場合、次の要素を確認する
                break
            }
        }
    }

    // すでに先頭(最後尾に移動)はソート済みのため、sliceの長さ-1回数分ループ。最後に残った要素もソート不要なので i > 0
    for i := len(array) - 1; i > 0; i-- {
        current_idx := 0
        swap(&array[i], &array[current_idx])
        left_idx := getLeftIdx(current_idx)
        // iはすでに確定している要素のインデックスのためそれよりも左の要素のインデックスが小さい場合に入れ替えが発生する
        // 右の要素のインデックスは左の要素のインデックス+1のため、十分条件は i > left_idx
        for i > left_idx {
            right_idx := left_idx + 1;
            child_idx := left_idx
            // 右の要素も生存していた場合、左の要素と比べて入れ替える要素を定める
            if i > right_idx && compare(array[right_idx], array[child_idx]) {
                child_idx = right_idx
            }
            // 親子の入れ替わりがあった場合、入れ替え後の子要素を確認する
            if compare(array[child_idx], array[current_idx]) {
                swap(&array[child_idx], &array[current_idx])
                current_idx = child_idx
                left_idx = getLeftIdx(current_idx)
            } else {
                // 親子の入れ替わりがない場合、次の要素を確認する
                break
            }
        }
    }
    return array
}

// 昇順用の関数
func Desc(num1 int, num2 int) bool {
    return num1 < num2
}

// 降順用の関数
func Asc(num1 int, num2 int) bool {
    return num1 > num2
}

// 値の入れ替えを行う
func swap(num1 *int, num2 *int) {
    *num1, *num2 = *num2, *num1
}

// ヒープソート用。引数の左の要素のインデックスを返す
func getLeftIdx(current_idx int) int {
    return 2 * current_idx + 1
}