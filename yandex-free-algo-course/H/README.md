# H. Сизиф

В этой задаче вы будете перекладывать камни. Изначально есть n кучек камней. Кучка i весит a<sub>i</sub> килограммов. Кучки можно объединять. При объединении кучек i и j затрачивается a<sub>i</sub>+a<sub>j</sub> единиц энергии, при этом две исходные кучки пропадают и появляется кучка весом a<sub>i</sub>+a<sub>j</sub>. Определите наименьшее количество энергии, которое надо затратить для объединения всех кучек в одну.

## Формат ввода

В первой строке дано число n (1≤n≤10<sup>5</sup>)

В следующей строке записаны массы кучек через пробел — a<sub>i</sub> (1≤a<sub>i</sub>≤10<sup>6</sup>)

## Формат вывода

Выведите единственное число — минимальную энергию, которую надо затратить на объединение всех кучек.

## Пример 1
<table>
<tr><td><b>Ввод</b></td><td><b>Вывод</b></td></tr>
<tr><td>2<br>
2 6
</td><td>8</td></tr>
</table>

## Пример 2
<table>
<tr><td><b>Ввод</b></td><td><b>Вывод</b></td></tr>
<tr><td>3<br>
6 2 4
</td><td>18</td></tr>
</table>

---
# Решение

Базовый алгоритм понятен — на каждом шаге нужно складывать две самых лёгких кучки, и тогда мы получим правильны ответ.

Сложность в строгих требованиях к производительности, в которые не укладывается ни одно из решений, использующее идеологию "Sorted Array", при которой мы изначально сортируем кучи, а затем вставляем новую кучу из 0 и 1 элемента масиива в нужное место. Самое медленное в этом подходе - вставка нового элемента в массив со сдвигом других элементов.

Таких реализаций, пожалуй, 3, вот они в порядке увеличения эффективности (могу быть неточен в индексах):
* append(arr[:i], arr[i+1]...)
* copy(arr[1:], arr[1:i+1])
* slices.Insert(arr[1:], i+1, value)

Соответственно, решением является использование такой структуры данных, как heap (куча), реализация которой уже есть в страндартной библиотеке golang
